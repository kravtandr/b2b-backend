package delivery

import (
	"b2b/m/internal/services/company/models"
	"b2b/m/internal/services/company/usecase"
	"b2b/m/pkg/error_adapter"
	company_service "b2b/m/pkg/services/company"
	"context"
)

type companyDelivery struct {
	companyUseCase usecase.CompanyUseCase
	errorAdapter   error_adapter.ErrorAdapter
	company_service.UnimplementedCompanyServiceServer
}

func (a *companyDelivery) GetCompanyById(ctx context.Context, request *company_service.GetCompanyRequestById) (*company_service.GetCompanyResponse, error) {
	company, err := a.companyUseCase.GetCompanyById(ctx, request.Id)
	if err != nil {
		return nil, a.errorAdapter.AdaptError(err)
	}

	return &company_service.GetCompanyResponse{
		Name:         company.Name,
		Description:  company.Description,
		LegalName:    company.LegalName,
		Itn:          company.Itn,
		Psrn:         company.Psrn,
		Address:      company.Address,
		LegalAddress: company.LegalAddress,
		Email:        company.Email,
		Phone:        company.Phone,
		Link:         company.Link,
		Activity:     company.Activity,
		OwnerId:      company.OwnerId,
		Rating:       company.Rating,
		Verified:     company.Verified,
	}, nil
}

func (a *companyDelivery) UpdateCompanyByOwnerId(ctx context.Context, request *company_service.UpdateCompanyRequest) (*company_service.GetCompanyAndPostResponse, error) {
	company, post, err := a.companyUseCase.UpdateCompanyById(ctx, models.Company{
		Name:         request.Name,
		Description:  request.Description,
		Address:      request.Address,
		LegalAddress: request.LegalAddress,
		Itn:          request.Itn,
		Phone:        request.Phone,
		Link:         request.Link,
		Activity:     request.Activity,
		OwnerId:      request.OwnerId,
	}, request.Post)

	if err != nil {
		return nil, a.errorAdapter.AdaptError(err)
	}

	return &company_service.GetCompanyAndPostResponse{
		Name:         company.Name,
		Description:  company.Description,
		LegalName:    company.LegalName,
		Itn:          company.Itn,
		Psrn:         company.Psrn,
		Address:      company.Address,
		LegalAddress: company.LegalAddress,
		Email:        company.Email,
		Phone:        company.Phone,
		Link:         company.Link,
		Activity:     company.Activity,
		OwnerId:      company.OwnerId,
		Rating:       company.Rating,
		Verified:     company.Verified,
		Post:         post,
	}, nil
}

func NewCompanyDelivery(
	companyUseCase usecase.CompanyUseCase,
	errorAdapter error_adapter.ErrorAdapter,
) company_service.CompanyServiceServer {
	return &companyDelivery{
		companyUseCase: companyUseCase,
		errorAdapter:   errorAdapter,
	}
}
