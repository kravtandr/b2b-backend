package delivery

import (
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

func NewCompanyDelivery(
	companyUseCase usecase.CompanyUseCase,
	errorAdapter error_adapter.ErrorAdapter,
) company_service.CompanyServiceServer {
	return &companyDelivery{
		companyUseCase: companyUseCase,
		errorAdapter:   errorAdapter,
	}
}
