package usecase

import (
	chat_service "b2b/m/pkg/services/chat"
	"context"

	company_usecase "b2b/m/internal/gateway/company/usecase"
	"b2b/m/internal/models"
)

type ChatUsecase interface {
	Login(ctx context.Context, request *models.LoginUserRequest) (*models.CompanyWithCookie, error)
}

type chatUsecase struct {
	chatGRPC    chatGRPC
	companyGRPC company_usecase.CompanyGRPC
}

func (u *chatUsecase) Login(ctx context.Context, request *models.LoginUserRequest) (*models.CompanyWithCookie, error) {
	response, err := u.chatGRPC.LoginUser(ctx, &chat_service.LoginRequest{
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

// func (u *userUsecase) Register(ctx context.Context, request *models.RegisterUserRequest) (*models.Session, error) {
// 	response, err := u.authGRPC.RegisterUser(ctx, &auth_service.RegisterRequest{
// 		Email:    request.Email,
// 		Password: request.Password,
// 		Name:     request.Name,
// 		Surname:  request.Surname,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &models.Session{
// 		Token:  response.Token,
// 		Cookie: response.Cookie,
// 	}, nil
// }

// func (u *userUsecase) CheckEmail(ctx context.Context, request *models.Email) (*models.PublicUser, error) {
// 	response, err := u.authGRPC.CheckEmail(ctx, &auth_service.CheckEmailRequest{Email: request.Email})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &models.PublicUser{
// 		Name:       response.Name,
// 		Surname:    response.Surname,
// 		Patronymic: response.Patronymic,
// 		Email:      response.Email,
// 	}, nil
// }

// func (u *userUsecase) FastRegister(ctx context.Context, request *models.FastRegistrationForm) (*models.CompanyWithCookie, error) {
// 	response, err := u.authGRPC.FastRegister(ctx, &auth_service.FastRegisterRequest{
// 		Name:       request.Name,
// 		LegalName:  request.LegalName,
// 		Itn:        request.Itn,
// 		Email:      request.Email,
// 		Password:   request.Password,
// 		OwnerName:  request.OwnerName,
// 		Surname:    request.Surname,
// 		Patronymic: request.Patronymic,
// 		Country:    request.Country,
// 		Post:       request.Post,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &models.CompanyWithCookie{
// 		Cookie:    response.Cookie,
// 		Token:     response.Token,
// 		Name:      response.Name,
// 		Email:     response.Email,
// 		LegalName: response.LegalName,
// 		Itn:       response.Itn,
// 		OwnerId:   response.OwnerId,
// 	}, nil
// }

// func (u *userUsecase) Logout(ctx context.Context, cookie string) error {
// 	_, err := u.authGRPC.LogoutUser(ctx, &auth_service.Session{
// 		Token:  "??",
// 		Cookie: cookie,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (u *userUsecase) Profile(ctx context.Context, userID int) (*models.Profile, error) {
// 	response, err := u.authGRPC.GetUser(ctx, &auth_service.GetUserRequest{Id: int64(userID)})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &models.Profile{
// 		Id:          userID,
// 		Name:        response.Name,
// 		Surname:     response.Surname,
// 		Avatar:      response.Image,
// 		Email:       response.Email,
// 		Description: response.Description,
// 	}, nil
// }

// func (u *userUsecase) GetUserIdByCookie(ctx context.Context, hash string) (int64, error) {
// 	response, err := u.authGRPC.GetUserIdByCookie(ctx, &auth_service.GetUserIdByCookieRequest{Hash: hash})
// 	if err != nil {
// 		return -1, err
// 	}

// 	return response.Id, nil
// }

// func (u *userUsecase) UpdateProfile(ctx context.Context, userID int64, request *models.PublicCompanyAndOwnerRequest) (*models.PublicCompanyAndOwnerResponse, error) {
// 	updatedUser, err := u.authGRPC.UpdateUser(ctx, &auth_service.UpdateUserRequest{
// 		Id:         userID,
// 		Name:       request.Owner.Name,
// 		Surname:    request.Owner.Surname,
// 		Patronymic: request.Owner.Patronymic,
// 		Email:      request.Owner.Email,
// 		Password:   request.Owner.Password,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	updatedCompany, err := u.companyGRPC.UpdateCompanyByOwnerId(ctx, &company_service.UpdateCompanyRequest{
// 		Name:         request.Company.Name,
// 		Description:  request.Company.Description,
// 		Address:      request.Company.Address,
// 		LegalAddress: request.Company.LegalAddress,
// 		Itn:          request.Company.Itn,
// 		Phone:        request.Company.Phone,
// 		Link:         request.Company.Link,
// 		Activity:     request.Company.Activity,
// 		OwnerId:      userID,
// 		Post:         request.Post,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &models.PublicCompanyAndOwnerResponse{
// 		Owner: models.UpdateUserResponse{
// 			Name:       updatedUser.Name,
// 			Surname:    updatedUser.Surname,
// 			Patronymic: updatedUser.Patronymic,
// 			Email:      updatedUser.Email,
// 		},
// 		Company: models.CompanyUpdateProfileResponse{
// 			Name:         request.Company.Name,
// 			Description:  request.Company.Description,
// 			Address:      request.Company.Address,
// 			LegalAddress: request.Company.LegalAddress,
// 			Phone:        request.Company.Phone,
// 			Link:         request.Company.Link,
// 			Activity:     request.Company.Activity,
// 		},
// 		Post: updatedCompany.Post,
// 	}, nil
// }

//func (u *userUsecase) UpdateProfile(ctx context.Context, userID int, request *models.UpdateProfileRequest) (*models.Profile, error) {
//	response, err := u.authGRPC.UpdateUser(ctx, &auth_service.UpdateUserRequest{
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

//func (u *userUsecase) GetUserInfo(ctx context.Context, id int) (*models.Profile, error) {
// 	responce, err := u.authGRPC.GetUserInfo(ctx, &auth_service.GetUserRequest{Id: int64(id)})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &models.Profile{
// 		Id:      int(responce.UserId),
// 		Name:    responce.Name,
// 		Surname: responce.Surname,
// 		Avatar:  responce.Image,
// 	}, nil
// }

func NewChatUsecase(chatGRPC chatGRPC, companyGRPC company_usecase.CompanyGRPC) ChatUsecase {
	return &chatUsecase{chatGRPC: chatGRPC, companyGRPC: companyGRPC}
}
