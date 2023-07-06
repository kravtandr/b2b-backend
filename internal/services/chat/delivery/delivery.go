package delivery

import (
	"b2b/m/internal/services/chat/models"
	"b2b/m/internal/services/chat/usecase"
	"b2b/m/pkg/error_adapter"
	chat_service "b2b/m/pkg/services/chat"
	"context"
)

type chatDelivery struct {
	chatUsecase  usecase.ChatUseCase
	errorAdapter error_adapter.ErrorAdapter
	chat_service.UnimplementedChatServiceServer
}

// func (a *authDelivery) ValidateSession(ctx context.Context, session *auth_service.Session) (*auth_service.ValidateSessionResponse, error) {
// 	response, err := a.authUsecase.ValidateSession(ctx, session.Cookie)
// 	if err != nil {
// 		return &auth_service.ValidateSessionResponse{}, a.errorAdapter.AdaptError(err)
// 	}

// 	return &auth_service.ValidateSessionResponse{UserId: response}, nil
// }

// func (a *authDelivery) LogoutUser(ctx context.Context, session *auth_service.Session) (*empty.Empty, error) {
// 	return &empty.Empty{}, a.authUsecase.LogoutUser(ctx, session.Cookie)
// }

// func (a *authDelivery) CheckEmail(ctx context.Context, request *auth_service.CheckEmailRequest) (*auth_service.GetPublicUserResponse, error) {
// 	response, err := a.authUsecase.GetUserByEmail(ctx, request.Email)
// 	if err != nil {
// 		return &auth_service.GetPublicUserResponse{}, a.errorAdapter.AdaptError(err)
// 	}

// 	return &auth_service.GetPublicUserResponse{
// 		Name:       response.Name,
// 		Surname:    response.Surname,
// 		Patronymic: response.Patronymic,
// 		Email:      response.Email,
// 	}, nil
// }

func (a *chatDelivery) LoginUser(ctx context.Context, request *chat_service.LoginRequest) (*chat_service.LoginResponse, error) {
	response, err := a.chatUsecase.LoginUser(ctx, &models.User{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return &chat_service.LoginResponse{}, a.errorAdapter.AdaptError(err)
	}

	return &chat_service.LoginResponse{
		Cookie:       response.Cookie,
		Token:        response.Token,
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

// func (a *authDelivery) FastRegister(ctx context.Context, request *auth_service.FastRegisterRequest) (*auth_service.LoginResponse, error) {
// 	response, err := a.authUsecase.FastRegistration(ctx, &models.FastRegistrationForm{
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
// 		return &auth_service.LoginResponse{}, a.errorAdapter.AdaptError(err)
// 	}
// 	return &auth_service.LoginResponse{
// 		Cookie:       response.Cookie,
// 		Token:        response.Token,
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

// func (a *authDelivery) RegisterUser(ctx context.Context, request *auth_service.RegisterRequest) (*auth_service.RegisterResponse, error) {
// 	response, err := a.authUsecase.RegisterUser(ctx, &models.User{
// 		Name:     request.Name,
// 		Surname:  request.Surname,
// 		Email:    request.Email,
// 		Password: request.Password,
// 	})
// 	if err != nil {
// 		return &auth_service.RegisterResponse{}, a.errorAdapter.AdaptError(err)
// 	}

// 	return &auth_service.RegisterResponse{
// 		Cookie: response.Cookie,
// 		Token:  response.Token,
// 	}, nil
// }

// func (a *authDelivery) GetUser(ctx context.Context, request *auth_service.GetUserRequest) (*auth_service.GetUserResponse, error) {
// 	response, err := a.authUsecase.GetUser(ctx, request.Id)
// 	if err != nil {
// 		return &auth_service.GetUserResponse{}, a.errorAdapter.AdaptError(err)
// 	}

// 	return &auth_service.GetUserResponse{
// 		Name:    response.Name,
// 		Surname: response.Surname,
// 		Email:   response.Email,
// 	}, nil
// }

// func (a *authDelivery) UpdateUser(ctx context.Context, request *auth_service.UpdateUserRequest) (*auth_service.GetPublicUserResponse, error) {
// 	response, err := a.authUsecase.UpdateUser(ctx, &models.User{
// 		Id:         request.Id,
// 		Name:       request.Name,
// 		Surname:    request.Surname,
// 		Patronymic: request.Patronymic,
// 		Email:      request.Email,
// 		Password:   request.Password,
// 	})
// 	if err != nil {
// 		return &auth_service.GetPublicUserResponse{}, a.errorAdapter.AdaptError(err)
// 	}

// 	return &auth_service.GetPublicUserResponse{
// 		Name:       response.Name,
// 		Surname:    response.Surname,
// 		Patronymic: response.Patronymic,
// 		Email:      response.Email,
// 	}, nil
// }

// func (a *authDelivery) GetUserInfo(ctx context.Context, request *auth_service.GetUserRequest) (*auth_service.UserInfo, error) {
// 	responce, err := a.authUsecase.GetUserInfo(ctx, int(request.Id))
// 	if err != nil {
// 		return &auth_service.UserInfo{}, a.errorAdapter.AdaptError(err)
// 	}

// 	return &auth_service.UserInfo{
// 		UserId:  responce.Id,
// 		Name:    responce.Name,
// 		Surname: responce.Surname,
// 	}, nil
// }

// func (a *authDelivery) GetUserByEmail(ctx context.Context, request *auth_service.UserEmailRequest) (*auth_service.UserId, error) {
// 	responce, err := a.authUsecase.GetUserByEmail(ctx, request.Email)
// 	if err != nil {
// 		return &auth_service.UserId{}, a.errorAdapter.AdaptError(err)
// 	}

// 	return &auth_service.UserId{Id: responce.Id}, nil
// }

// func (a *authDelivery) GetUserIdByCookie(ctx context.Context, request *auth_service.GetUserIdByCookieRequest) (*auth_service.UserId, error) {
// 	responce, err := a.authUsecase.ValidateSession(ctx, request.Hash)
// 	if err != nil {
// 		return &auth_service.UserId{}, a.errorAdapter.AdaptError(err)
// 	}

// 	return &auth_service.UserId{Id: responce}, nil
// }

func NewChatDelivery(
	chatUsecase usecase.ChatUseCase,
	errorAdapter error_adapter.ErrorAdapter,
) chat_service.ChatServiceServer {
	return &chatDelivery{
		chatUsecase:  chatUsecase,
		errorAdapter: errorAdapter,
	}
}
