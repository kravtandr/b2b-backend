package repository

import (
	"b2b/m/internal/services/company/models"
	"b2b/m/pkg/errors"
	"context"
	"github.com/jackc/pgx/v4"
	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

type CompanyRepository interface {
	GetCompanyById(ctx context.Context, ID int64) (*models.Company, error)
}

type companyRepository struct {
	queryFactory QueryFactory
	conn         *pgxpool.Pool
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
