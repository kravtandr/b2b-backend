package usecase

import (
	"context"

	"b2b/m/internal/models"
	auth_service "b2b/m/pkg/services/auth"
)

type UserUsecase interface {
	Login(ctx context.Context, request *models.LoginUserRequest) (*models.CompanyWithCookie, error)
	Register(ctx context.Context, request *models.RegisterUserRequest) (*models.Session, error)
	Logout(ctx context.Context, cookie string) error
	GetUserInfo(ctx context.Context, id int) (*models.Profile, error)
	FastRegister(ctx context.Context, request *models.FastRegistrationForm) (*models.CompanyWithCookie, error)
	Profile(ctx context.Context, userID int) (*models.Profile, error)
	UpdateProfile(ctx context.Context, userID int, request *models.UpdateProfileRequest) (*models.Profile, error)
}

type userUsecase struct {
	authGRPC authGRPC
}

func (u *userUsecase) Login(ctx context.Context, request *models.LoginUserRequest) (*models.CompanyWithCookie, error) {
	response, err := u.authGRPC.LoginUser(ctx, &auth_service.LoginRequest{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}

	return &models.CompanyWithCookie{
		Token:        response.Token,
		Cookie:       response.Cookie,
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
	}, nil
}

func (u *userUsecase) Register(ctx context.Context, request *models.RegisterUserRequest) (*models.Session, error) {
	response, err := u.authGRPC.RegisterUser(ctx, &auth_service.RegisterRequest{
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

func (u *userUsecase) FastRegister(ctx context.Context, request *models.FastRegistrationForm) (*models.CompanyWithCookie, error) {
	response, err := u.authGRPC.FastRegister(ctx, &auth_service.FastRegisterRequest{
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
		Cookie: response.Cookie,
		Token:  response.Token,
		Name:   response.Name,
		//Description:  userCompany.Description,
		LegalName: response.LegalName,
		Itn:       response.Itn,
		//Psrn:         userCompany.Psrn,
		//Address:      userCompany.Address,
		//LegalAddress: userCompany.LegalAddress,
		//Email:        userCompany.Email,
		//Phone:        userCompany.Phone,
		//Link:         userCompany.Link,
		//Activity:     userCompany.Activity,
		OwnerId: response.OwnerId,
		//Rating:       userCompany.Rating,
		//Verified:     userCompany.Verified,
	}, nil
}

func (u *userUsecase) Logout(ctx context.Context, cookie string) error {
	_, err := u.authGRPC.LogoutUser(ctx, &auth_service.Session{
		Token:  "??",
		Cookie: cookie,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) Profile(ctx context.Context, userID int) (*models.Profile, error) {
	response, err := u.authGRPC.GetUser(ctx, &auth_service.GetUserRequest{Id: int64(userID)})
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

func (u *userUsecase) UpdateProfile(ctx context.Context, userID int, request *models.UpdateProfileRequest) (*models.Profile, error) {
	response, err := u.authGRPC.UpdateUser(ctx, &auth_service.UpdateUserRequest{
		Id:          int64(userID),
		Name:        request.Name,
		Surname:     request.Surname,
		Email:       request.Email,
		Description: request.Description,
		Password:    request.Password,
		Image:       request.Avatar,
	})
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

func (u *userUsecase) GetUserInfo(ctx context.Context, id int) (*models.Profile, error) {
	responce, err := u.authGRPC.GetUserInfo(ctx, &auth_service.GetUserRequest{Id: int64(id)})
	if err != nil {
		return nil, err
	}

	return &models.Profile{
		Id:      int(responce.UserId),
		Name:    responce.Name,
		Surname: responce.Surname,
		Avatar:  responce.Image,
	}, nil
}

func NewUserUsecase(grpc authGRPC) UserUsecase {
	return &userUsecase{authGRPC: grpc}
}
