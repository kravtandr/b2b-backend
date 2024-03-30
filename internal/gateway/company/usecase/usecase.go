package usecase

import (
	"b2b/m/internal/models"
	company_service "b2b/m/pkg/services/company"
	"context"
	"fmt"
	"log"

	"gopkg.in/webdeskltd/dadata.v2"
)

type CompanyUseCase interface {
	GetCompanyById(ctx context.Context, Id int64) (*models.Company, error)
	GetCompanyByItnFromDaData(ctx context.Context, itn string) ([]dadata.ResponseParty, error)
	GetCompanyByProductId(ctx context.Context, Id int64) (*models.Company, error)
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
	response, err := u.companyGRPC.UpdateCompanyByOwnerId(ctx, &company_service.UpdateCompanyRequest{})
	if err != nil {
		return nil, err
	}

	return &models.Company{
		Name:         response.Name,
		Description:  response.Description,
		LegalName:    response.LegalName,
		Itn:          response.Itn,
		Psrn:         response.Psrn,
		Address:      response.Address,
		LegalAddress: response.LegalAddress,
		Email:        response.Email,
		Phone:        response.Phone,
		Link:         response.Link,
		Activity:     response.Activity,
		OwnerId:      response.OwnerId,
		Rating:       response.Rating,
		Verified:     response.Verified,
		Photo:        response.Photo,
	}, nil
}

func (u *companyUseCase) GetCompanyById(ctx context.Context, id int64) (*models.Company, error) {
	response, err := u.companyGRPC.GetCompanyById(ctx, &company_service.GetCompanyRequestById{Id: int64(id)})
	if err != nil {
		return nil, err
	}

	return &models.Company{
		Id:           response.Id,
		Name:         response.Name,
		Description:  response.Description,
		LegalName:    response.LegalName,
		Itn:          response.Itn,
		Psrn:         response.Psrn,
		Address:      response.Address,
		LegalAddress: response.LegalAddress,
		Email:        response.Email,
		Phone:        response.Phone,
		Link:         response.Link,
		Activity:     response.Activity,
		OwnerId:      response.OwnerId,
		Rating:       response.Rating,
		Verified:     response.Verified,
		Photo:        response.Photo,
	}, nil
}

// func (u *companyUseCase) GetCompanyByUserCookie(ctx context.Context, hash string) (*models.Company, error) {
// 	userId, err := u.authGRPC.GetUserIdByCookie(ctx, hash)
// 	if err != nil {
// 		return nil, err
// 	}

// 	response, err := u.companyGRPC.GetCompanyById(ctx, &company_service.GetCompanyRequestById{Id: int64(userId)})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &models.Company{
// 		Name:         response.Name,
// 		Description:  response.Description,
// 		LegalName:    response.LegalName,
// 		Itn:          response.Itn,
// 		Psrn:         response.Psrn,
// 		Address:      response.Address,
// 		LegalAddress: response.LegalAddress,
// 		Email:        response.Email,
// 		Phone:        response.Phone,
// 		Link:         response.Link,
// 		Activity:     response.Activity,
// 		OwnerId:      response.OwnerId,
// 		Rating:       response.Rating,
// 		Verified:     response.Verified,
// 	}, nil
// }

func (u *companyUseCase) GetCompanyByProductId(ctx context.Context, id int64) (*models.Company, error) {
	response, err := u.companyGRPC.GetCompanyByProductId(ctx, &company_service.IdRequest{Id: int64(id)})
	if err != nil {
		return nil, err
	}

	return &models.Company{
		Id:           response.Id,
		Name:         response.Name,
		Description:  response.Description,
		LegalName:    response.LegalName,
		Itn:          response.Itn,
		Psrn:         response.Psrn,
		Address:      response.Address,
		LegalAddress: response.LegalAddress,
		Email:        response.Email,
		Phone:        response.Phone,
		Link:         response.Link,
		Activity:     response.Activity,
		OwnerId:      response.OwnerId,
		Rating:       response.Rating,
		Verified:     response.Verified,
		Photo:        response.Photo,
	}, nil
}

func NewCompanyUseCase(grpc CompanyGRPC, daData *dadata.DaData) CompanyUseCase {
	return &companyUseCase{companyGRPC: grpc, daData: daData}
}
