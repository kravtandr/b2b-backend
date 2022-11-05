package usecase

import (
	company_service "b2b/m/pkg/services/company"
	"context"
	"fmt"

	"b2b/m/internal/models"
)

type CompanyUseCase interface {
	GetCompanyById(ctx context.Context, Id int64) (*models.Company, error)
}

type companyUseCase struct {
	companyGRPC companyGRPC
}

func (u *companyUseCase) GetCompanyById(ctx context.Context, id int64) (*models.Company, error) {
	fmt.Println("FUCKKKKKKKK companyUseCase")
	responce, err := u.companyGRPC.GetCompanyById(ctx, &company_service.GetCompanyRequestById{Id: int64(id)})
	if err != nil {
		fmt.Println("PIZDIC  err", err)
		return nil, err
	}

	return &models.Company{
		Name:         responce.Name,
		Description:  responce.Description,
		LegalName:    responce.LegalName,
		Itn:          responce.Itn,
		Psrn:         responce.Psrn,
		Address:      responce.Address,
		LegalAddress: responce.LegalAddress,
		Email:        responce.Email,
		Phone:        responce.Phone,
		Link:         responce.Link,
		Activity:     responce.Activity,
		OwnerId:      responce.OwnerId,
		Rating:       responce.Rating,
		Verified:     responce.Verified,
	}, nil
}

func NewCompanyUseCase(grpc companyGRPC) CompanyUseCase {
	return &companyUseCase{companyGRPC: grpc}
}
