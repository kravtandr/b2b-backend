package usecase

import (
	"context"

	company_usecase "b2b/m/internal/gateway/company/usecase"
	"b2b/m/internal/models"
	auth_service "b2b/m/pkg/services/auth"
	company_service "b2b/m/pkg/services/company"
)

type UserUsecase interface {
	Login(ctx context.Context, request *models.LoginUserRequest) (*models.CompanyWithCookie, error)
	Register(ctx context.Context, request *models.RegisterUserRequest) (*models.Session, error)
	Logout(ctx context.Context, cookie string) error
	GetUserInfo(ctx context.Context, id int64) (*models.Profile, error)
	FastRegister(ctx context.Context, request *models.FastRegistrationForm) (*models.CompanyWithCookie, error)
	Profile(ctx context.Context, userID int64) (*models.Profile, error)
	UpdateProfile(ctx context.Context, userID int64, request *models.PublicCompanyAndOwnerRequest) (*models.PublicCompanyAndOwnerResponse, error)
	GetUserIdByCookie(ctx context.Context, hash string) (int64, error)
	CheckEmail(ctx context.Context, request *models.Email) (*models.PublicUser, error)
}

type userUsecase struct {
	AuthGRPC    AuthGRPC
	companyGRPC company_usecase.CompanyGRPC
}

func (u *userUsecase) Login(ctx context.Context, request *models.LoginUserRequest) (*models.CompanyWithCookie, error) {
	response, err := u.AuthGRPC.LoginUser(ctx, &auth_service.LoginRequest{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}

	return &models.CompanyWithCookie{
		Cookie:    response.Cookie,
		Token:     response.Token,
		Name:      response.Name,
		Email:     response.Email,
		LegalName: response.LegalName,
		Itn:       response.Itn,
		OwnerId:   response.OwnerId,
	}, nil
}

func (u *userUsecase) Register(ctx context.Context, request *models.RegisterUserRequest) (*models.Session, error) {
	response, err := u.AuthGRPC.RegisterUser(ctx, &auth_service.RegisterRequest{
		Email:    request.Email,
		Password: request.Password,
		Name:     request.Name,
		Surname:  request.Surname,
	})
	if err != nil {
		return nil, err
	}

	return &models.Session{
		Token:  response.Token,
		Cookie: response.Cookie,
	}, nil
}

func (u *userUsecase) CheckEmail(ctx context.Context, request *models.Email) (*models.PublicUser, error) {
	response, err := u.AuthGRPC.CheckEmail(ctx, &auth_service.CheckEmailRequest{Email: request.Email})
	if err != nil {
		return nil, err
	}

	return &models.PublicUser{
		Name:       response.Name,
		Surname:    response.Surname,
		Patronymic: response.Patronymic,
		Email:      response.Email,
	}, nil
}

func (u *userUsecase) FastRegister(ctx context.Context, request *models.FastRegistrationForm) (*models.CompanyWithCookie, error) {
	response, err := u.AuthGRPC.FastRegister(ctx, &auth_service.FastRegisterRequest{
		Name:       request.Name,
		LegalName:  request.LegalName,
		Itn:        request.Itn,
		Email:      request.Email,
		Password:   request.Password,
		OwnerName:  request.OwnerName,
		Surname:    request.Surname,
		Patronymic: request.Patronymic,
		Country:    request.Country,
		Post:       request.Post,
	})
	if err != nil {
		return nil, err
	}

	return &models.CompanyWithCookie{
		Cookie:    response.Cookie,
		Token:     response.Token,
		Name:      response.Name,
		Email:     response.Email,
		LegalName: response.LegalName,
		Itn:       response.Itn,
		OwnerId:   response.OwnerId,
	}, nil
}

func (u *userUsecase) Logout(ctx context.Context, cookie string) error {
	_, err := u.AuthGRPC.LogoutUser(ctx, &auth_service.Session{
		Token:  "??",
		Cookie: cookie,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) Profile(ctx context.Context, userID int64) (*models.Profile, error) {
	response, err := u.AuthGRPC.GetUser(ctx, &auth_service.GetUserRequest{Id: int64(userID)})
	if err != nil {
		return nil, err
	}

	return &models.Profile{
		Id:          userID,
		Name:        response.Name,
		Surname:     response.Surname,
		Avatar:      response.Image,
		Email:       response.Email,
		Description: response.Description,
	}, nil
}

func (u *userUsecase) GetUserIdByCookie(ctx context.Context, hash string) (int64, error) {
	response, err := u.AuthGRPC.GetUserIdByCookie(ctx, &auth_service.GetUserIdByCookieRequest{Hash: hash})
	if err != nil {
		return -1, err
	}

	return response.Id, nil
}

func (u *userUsecase) UpdateProfile(ctx context.Context, userID int64, request *models.PublicCompanyAndOwnerRequest) (*models.PublicCompanyAndOwnerResponse, error) {
	updatedUser, err := u.AuthGRPC.UpdateUser(ctx, &auth_service.UpdateUserRequest{
		Id:         userID,
		Name:       request.Owner.Name,
		Surname:    request.Owner.Surname,
		Patronymic: request.Owner.Patronymic,
		Email:      request.Owner.Email,
		Password:   request.Owner.Password,
	})
	if err != nil {
		return nil, err
	}

	updatedCompany, err := u.companyGRPC.UpdateCompanyByOwnerId(ctx, &company_service.UpdateCompanyRequest{
		Name:         request.Company.Name,
		Description:  request.Company.Description,
		Address:      request.Company.Address,
		LegalAddress: request.Company.LegalAddress,
		Itn:          request.Company.Itn,
		Phone:        request.Company.Phone,
		Link:         request.Company.Link,
		Activity:     request.Company.Activity,
		OwnerId:      userID,
		Post:         request.Post,
	})
	if err != nil {
		return nil, err
	}

	return &models.PublicCompanyAndOwnerResponse{
		Owner: models.UpdateUserResponse{
			Name:       updatedUser.Name,
			Surname:    updatedUser.Surname,
			Patronymic: updatedUser.Patronymic,
			Email:      updatedUser.Email,
		},
		Company: models.CompanyUpdateProfileResponse{
			Name:         request.Company.Name,
			Description:  request.Company.Description,
			Address:      request.Company.Address,
			LegalAddress: request.Company.LegalAddress,
			Phone:        request.Company.Phone,
			Link:         request.Company.Link,
			Activity:     request.Company.Activity,
		},
		Post: updatedCompany.Post,
	}, nil
}

//func (u *userUsecase) UpdateProfile(ctx context.Context, userID int, request *models.UpdateProfileRequest) (*models.Profile, error) {
//	response, err := u.AuthGRPC.UpdateUser(ctx, &auth_service.UpdateUserRequest{
//		Id:          int64(userID),
//		Name:        request.Name,
//		Surname:     request.Surname,
//		Email:       request.Email,
//		Description: request.Description,
//		Password:    request.Password,
//		Image:       request.Avatar,
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	return &models.Profile{
//		Id:          userID,
//		Name:        response.Name,
//		Surname:     response.Surname,
//		Avatar:      response.Image,
//		Email:       response.Email,
//		Description: response.Description,
//	}, nil
//}

func (u *userUsecase) GetUserInfo(ctx context.Context, id int64) (*models.Profile, error) {
	response, err := u.AuthGRPC.GetUserInfo(ctx, &auth_service.GetUserRequest{Id: int64(id)})
	if err != nil {
		return nil, err
	}

	return &models.Profile{
		Id:      response.UserId,
		Name:    response.Name,
		Surname: response.Surname,
		Avatar:  response.Image,
	}, nil
}

func NewUserUsecase(AuthGRPC AuthGRPC, companyGRPC company_usecase.CompanyGRPC) UserUsecase {
	return &userUsecase{AuthGRPC: AuthGRPC, companyGRPC: companyGRPC}
}
