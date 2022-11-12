package usecase

import (
	"b2b/m/internal/services/company/models"
	"context"
)

type CompanyUseCase interface {
	GetCompanyById(ctx context.Context, ID int64) (*models.Company, error)
	UpdateCompanyById(ctx context.Context, newCompany models.Company, post string) (*models.Company, string, error)
}

type companyUseCase struct {
	repo companyRepository
}

func (a *companyUseCase) UpdateCompanyById(ctx context.Context, newCompany models.Company, post string) (*models.Company, string, error) {
	link, err := a.repo.GetCompanyUserLinkByOwnerIdAndItn(ctx, newCompany.OwnerId, newCompany.Itn)
	currentCompany, err := a.repo.GetCompanyById(ctx, link.CompanyId)
	newCompany.Id = currentCompany.Id
	updatedCompany, err := a.repo.UpdateCompanyById(ctx, newCompany)
	if err != nil {
		return nil, "error UpdateCompanyByOwnerId", err
	}
	post, err = a.repo.UpdateCompanyUsersLink(ctx, updatedCompany.Id, updatedCompany.OwnerId, post)
	if err != nil {
		return nil, "error UpdateCompanyUsersLink", err
	}
	return updatedCompany, post, nil
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
