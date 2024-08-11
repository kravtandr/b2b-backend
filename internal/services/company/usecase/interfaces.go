package usecase

import (
	"context"

	"b2b/m/internal/services/company/models"
)

type companyRepository interface {
	GetCompanyById(ctx context.Context, ID int64) (*models.Company, error)
	GetCompanyByOwnerIdAndItn(ctx context.Context, company models.Company) (*models.Company, error)
	UpdateCompanyById(ctx context.Context, newCompany models.Company) (*models.Company, error)
	UpdateCompanyUsersLink(ctx context.Context, companyId int64, userId int64, post string) (string, error)
	GetCompanyUserLinkByOwnerIdAndItn(ctx context.Context, id int64, itn string) (*models.CompaniesUsersLink, error)
	GetCompanyByProductId(ctx context.Context, ID int64) (*models.Company, error)
}
