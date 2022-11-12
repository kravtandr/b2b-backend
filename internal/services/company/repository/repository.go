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
	fmt.Println("GetCompanyUserLinkByOwnerIdAndItn", id, itn)
	query := a.queryFactory.CreateGetCompanyUserLinkByOwnerIdAndItn(id, itn)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	companiesUsersLink := &models.CompaniesUsersLink{}
	if err := row.Scan(
		&companiesUsersLink.Id, &companiesUsersLink.CompanyId, &companiesUsersLink.UserId, &companiesUsersLink.Post, &companiesUsersLink.Itn,
	); err != nil {
		if err == pgx.ErrNoRows {
			fmt.Println("ERROR 1", err)
			return nil, errors.UserDoesNotExist
		}
		fmt.Println("ERROR 2", err)
		return nil, err
	}
	fmt.Println("GetCompanyUserLinkByOwnerIdAndItn RETURN ", companiesUsersLink)
	return companiesUsersLink, nil
}
func (a *companyRepository) UpdateCompanyUsersLink(ctx context.Context, companyId int64, userId int64, post string) (string, error) {
	query := a.queryFactory.CreateUpdateCompanyUsersLink(companyId, userId, post)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	companiesUsersLink := &models.CompaniesUsersLink{}
	if err := row.Scan(
		&companiesUsersLink.Id, &companiesUsersLink.CompanyId, &companiesUsersLink.UserId, &companiesUsersLink.Post,
	); err != nil {
		if err == pgx.ErrNoRows {
			return "error", errors.UserDoesNotExist
		}

		return "error", err
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

func NewCompanyRepository(
	queryFactory QueryFactory,
	conn *pgxpool.Pool,
) CompanyRepository {
	return &companyRepository{
		queryFactory: queryFactory,
		conn:         conn,
	}
}
