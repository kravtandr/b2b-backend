package usecase

import (
	"b2b/m/internal/models"
	company_service "b2b/m/pkg/services/company"
	"context"
	"fmt"
	"gopkg.in/webdeskltd/dadata.v2"
	"log"
)

type CompanyUseCase interface {
	GetCompanyById(ctx context.Context, Id int64) (*models.Company, error)
	GetCompanyByItnFromDaData(ctx context.Context, itn string) ([]dadata.ResponseParty, error)
}

type companyUseCase struct {
	companyGRPC CompanyGRPC
	daData      *dadata.DaData
}

func (u *companyUseCase) GetCompanyByItnFromDaData(ctx context.Context, itn string) ([]dadata.ResponseParty, error) {
	var companyFullInfo []dadata.ResponseParty
	reqParams := dadata.SuggestRequestParams{Query: "\"" + fmt.Sprint(itn) + "\""}
	companyFullInfo, err := u.daData.SuggestParties(reqParams)
	if err != nil {
		log.Printf("Error while getting GetCompanyByItnDaData")
		return companyFullInfo, err
	}
	return companyFullInfo, nil
}

func (u *companyUseCase) UpdateCompanyByOwnerId(ctx context.Context) (*models.Company, error) {
	responce, err := u.companyGRPC.UpdateCompanyByOwnerId(ctx, &company_service.UpdateCompanyRequest{})
	if err != nil {
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

func (u *companyUseCase) GetCompanyById(ctx context.Context, id int64) (*models.Company, error) {
	responce, err := u.companyGRPC.GetCompanyById(ctx, &company_service.GetCompanyRequestById{Id: int64(id)})
	if err != nil {
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

func NewCompanyUseCase(grpc CompanyGRPC, daData *dadata.DaData) CompanyUseCase {
	return &companyUseCase{companyGRPC: grpc, daData: daData}
}
