package repository

import (
	"b2b/m/internal/services/company/models"
	"b2b/m/pkg/errors"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

type CompanyRepository interface {
	GetCompanyById(ctx context.Context, ID int64) (*models.Company, error)
	GetCompanyByOwnerIdAndItn(ctx context.Context, company models.Company) (*models.Company, error)
	UpdateCompanyById(ctx context.Context, newCompany models.Company) (*models.Company, error)
	UpdateCompanyUsersLink(ctx context.Context, companyId int64, userId int64, post string) (string, error)
	GetCompanyUserLinkByOwnerIdAndItn(ctx context.Context, id int64, itn string) (*models.CompaniesUsersLink, error)
	GetCompanyByProductId(ctx context.Context, ID int64) (*models.Company, error)
	GetProductsCompaniesLink(ctx context.Context, productId int64) (*models.ProductsCompaniesLink, error)
}

type companyRepository struct {
	queryFactory QueryFactory
	conn         *pgxpool.Pool
}

func (a *companyRepository) UpdateCompanyById(ctx context.Context, newCompany models.Company) (*models.Company, error) {
	query := a.queryFactory.CreateUpdateCompanyById(newCompany)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	repoCompany := &models.Company{}
	if err := row.Scan(
		&repoCompany.Id, &repoCompany.Name, &repoCompany.Description, &repoCompany.LegalName, &repoCompany.Itn, &repoCompany.Psrn, &repoCompany.Address, &repoCompany.LegalAddress, &repoCompany.Email, &repoCompany.Phone, &repoCompany.Link, &repoCompany.Activity, &repoCompany.OwnerId, &repoCompany.Rating, &repoCompany.Verified,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.UserDoesNotExist
		}

		return nil, err
	}

	return repoCompany, nil
}

func (a *companyRepository) GetCompanyByOwnerIdAndItn(ctx context.Context, company models.Company) (*models.Company, error) {
	query := a.queryFactory.GetCompanyByOwnerIdAndItn(company.OwnerId, company.Itn)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	repoCompany := &models.Company{}
	if err := row.Scan(
		&repoCompany.Id, &repoCompany.Name, &repoCompany.Description, &repoCompany.LegalName, &repoCompany.Itn, &repoCompany.Psrn, &repoCompany.Address, &repoCompany.LegalAddress, &repoCompany.Email, &repoCompany.Phone, &repoCompany.Link, &repoCompany.Activity, &repoCompany.OwnerId, &repoCompany.Rating, &repoCompany.Verified,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.UserDoesNotExist
		}

		return nil, err
	}

	return repoCompany, nil
}

func (a *companyRepository) GetCompanyUserLinkByOwnerIdAndItn(ctx context.Context, id int64, itn string) (*models.CompaniesUsersLink, error) {
	query := a.queryFactory.CreateGetCompanyUserLinkByOwnerIdAndItn(id, itn)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	companiesUsersLink := &models.CompaniesUsersLink{}
	if err := row.Scan(
		&companiesUsersLink.Id, &companiesUsersLink.Post, &companiesUsersLink.CompanyId, &companiesUsersLink.UserId, &companiesUsersLink.Itn,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.CompanyUsersLinkNotExist
		}
		return nil, err
	}
	return companiesUsersLink, nil
}
func (a *companyRepository) UpdateCompanyUsersLink(ctx context.Context, companyId int64, userId int64, post string) (string, error) {
	query := a.queryFactory.CreateUpdateCompanyUsersLink(companyId, userId, post)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	companiesUsersLink := &models.CompaniesUsersLink{}
	if err := row.Scan(
		&companiesUsersLink.Id, &companiesUsersLink.Post, &companiesUsersLink.CompanyId, &companiesUsersLink.UserId, &companiesUsersLink.Itn,
	); err != nil {
		if err == pgx.ErrNoRows {
			return "no rows", errors.CompanyUsersLinkNotExist
		}

		return fmt.Sprint(err), err
	}
	return companiesUsersLink.Post, nil
}

func (a *companyRepository) GetCompanyById(ctx context.Context, ID int64) (*models.Company, error) {
	query := a.queryFactory.CreateGetCompanyByID(ID)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	company := &models.Company{}
	if err := row.Scan(
		&company.Id, &company.Name, &company.Description, &company.LegalName, &company.Itn, &company.Psrn, &company.Address, &company.LegalAddress, &company.Email, &company.Phone, &company.Link, &company.Activity, &company.OwnerId, &company.Rating, &company.Verified,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.CompanyDoesNotExist
		}

		return nil, err
	}

	return company, nil
}

func (a *companyRepository) GetProductsCompaniesLink(ctx context.Context, productId int64) (*models.ProductsCompaniesLink, error) {
	query := a.queryFactory.CreateGetProductsCompaniesLink(productId)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	productsCompaniesLink := &models.ProductsCompaniesLink{}
	if err := row.Scan(
		&productsCompaniesLink.Id, &productsCompaniesLink.CompanyId, &productsCompaniesLink.ProductId, &productsCompaniesLink.Amount,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.CompanyUsersLinkNotExist
		}

		return nil, err
	}
	return productsCompaniesLink, nil
}

func (a *companyRepository) GetCompanyByProductId(ctx context.Context, ID int64) (*models.Company, error) {
	productsCompaniesLink, err := a.GetProductsCompaniesLink(ctx, ID)
	if err != nil {
		return nil, err
	}
	company, err := a.GetCompanyById(ctx, productsCompaniesLink.CompanyId)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func NewCompanyRepository(
	queryFactory QueryFactory,
	conn *pgxpool.Pool,
) CompanyRepository {
	return &companyRepository{
		queryFactory: queryFactory,
		conn:         conn,
	}
}
