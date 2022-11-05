package usecase

import (
	"context"

	"b2b/m/internal/services/company/models"
)

type CompanyUseCase interface {
	GetCompanyById(ctx context.Context, ID int64) (*models.Company, error)
}

type companyUseCase struct {
	repo companyRepository
}

func (a *companyUseCase) GetCompanyById(ctx context.Context, ID int64) (*models.Company, error) {
	user, err := a.repo.GetCompanyById(ctx, ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewCompanyUseCase(
	repo companyRepository,
) CompanyUseCase {
	return &companyUseCase{
		repo: repo,
	}
}
