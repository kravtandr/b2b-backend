package delivery

import (
	"b2b/m/internal/services/auth/models"
	"b2b/m/internal/services/auth/usecase"
	"b2b/m/pkg/error_adapter"
	auth_service "b2b/m/pkg/services/auth"
	"context"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
)

type authDelivery struct {
	authUsecase  usecase.AuthUseCase
	errorAdapter error_adapter.ErrorAdapter
	auth_service.UnimplementedAuthServiceServer
}

func (a *authDelivery) ValidateSession(ctx context.Context, session *auth_service.Session) (*auth_service.ValidateSessionResponse, error) {
	response, err := a.authUsecase.ValidateSession(ctx, session.Cookie)
	if err != nil {
		return &auth_service.ValidateSessionResponse{}, a.errorAdapter.AdaptError(err)
	}

	return &auth_service.ValidateSessionResponse{UserId: response}, nil
}

func (a *authDelivery) LogoutUser(ctx context.Context, session *auth_service.Session) (*empty.Empty, error) {
	return &empty.Empty{}, a.authUsecase.LogoutUser(ctx, session.Cookie)
}

func (a *authDelivery) CheckEmail(ctx context.Context, request *auth_service.CheckEmailRequest) (*auth_service.GetPublicUserResponse, error) {
	response, err := a.authUsecase.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return &auth_service.GetPublicUserResponse{}, a.errorAdapter.AdaptError(err)
	}

	return &auth_service.GetPublicUserResponse{
		Name:       response.Name,
		Surname:    response.Surname,
		Patronymic: response.Patronymic,
		Email:      response.Email,
	}, nil
}

func (a *authDelivery) LoginUser(ctx context.Context, request *auth_service.LoginRequest) (*auth_service.LoginResponse, error) {
	response, err := a.authUsecase.LoginUser(ctx, &models.User{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return &auth_service.LoginResponse{}, a.errorAdapter.AdaptError(err)
	}

	return &auth_service.LoginResponse{
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

func (a *authDelivery) FastRegister(ctx context.Context, request *auth_service.FastRegisterRequest) (*auth_service.LoginResponse, error) {
	response, err := a.authUsecase.FastRegistration(ctx, &models.FastRegistrationForm{
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
		return &auth_service.LoginResponse{}, a.errorAdapter.AdaptError(err)
	}
	return &auth_service.LoginResponse{
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

func (a *authDelivery) RegisterUser(ctx context.Context, request *auth_service.RegisterRequest) (*auth_service.RegisterResponse, error) {
	response, err := a.authUsecase.RegisterUser(ctx, &models.User{
		Name:     request.Name,
		Surname:  request.Surname,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return &auth_service.RegisterResponse{}, a.errorAdapter.AdaptError(err)
	}

	return &auth_service.RegisterResponse{
		Cookie: response.Cookie,
		Token:  response.Token,
	}, nil
}

func (a *authDelivery) GetUser(ctx context.Context, request *auth_service.GetUserRequest) (*auth_service.GetUserResponse, error) {
	log.Println("authDelivery -> GetUser", request.Id)
	response, err := a.authUsecase.GetUser(ctx, request.Id)
	if err != nil {
		return &auth_service.GetUserResponse{}, a.errorAdapter.AdaptError(err)
	}

	return &auth_service.GetUserResponse{
		Name:       response.Name,
		Surname:    response.Surname,
		Patronymic: response.Patronymic,
		Country:    response.Country,
		Email:      response.Email,
		Balance:    response.Balance,
	}, nil
}

func (a *authDelivery) UpdateUser(ctx context.Context, request *auth_service.UpdateUserRequest) (*auth_service.GetPublicUserResponse, error) {
	response, err := a.authUsecase.UpdateUser(ctx, &models.User{
		Id:         request.Id,
		Name:       request.Name,
		Surname:    request.Surname,
		Patronymic: request.Patronymic,
		Email:      request.Email,
		Password:   request.Password,
	})
	if err != nil {
		return &auth_service.GetPublicUserResponse{}, a.errorAdapter.AdaptError(err)
	}

	return &auth_service.GetPublicUserResponse{
		Name:       response.Name,
		Surname:    response.Surname,
		Patronymic: response.Patronymic,
		Email:      response.Email,
	}, nil
}

func (a *authDelivery) GetUserInfo(ctx context.Context, request *auth_service.GetUserRequest) (*auth_service.UserInfo, error) {
	log.Println("authDelivery -> GetUserInfo", request.Id)
	response, err := a.authUsecase.GetUserInfo(ctx, request.Id)
	if err != nil {
		log.Println("Error: authDelivery -> GetUserInfo", err)
		return &auth_service.UserInfo{}, a.errorAdapter.AdaptError(err)
	}

	return &auth_service.UserInfo{
		UserId:  response.Id,
		Name:    response.Name,
		Surname: response.Surname,
	}, nil
}

func (a *authDelivery) GetUserByEmail(ctx context.Context, request *auth_service.UserEmailRequest) (*auth_service.UserId, error) {
	response, err := a.authUsecase.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return &auth_service.UserId{}, a.errorAdapter.AdaptError(err)
	}

	return &auth_service.UserId{Id: response.Id}, nil
}

func (a *authDelivery) GetUserIdByCookie(ctx context.Context, request *auth_service.GetUserIdByCookieRequest) (*auth_service.UserId, error) {
	response, err := a.authUsecase.ValidateSession(ctx, request.Hash)
	if err != nil {
		return &auth_service.UserId{}, a.errorAdapter.AdaptError(err)
	}

	return &auth_service.UserId{Id: response}, nil
}

func (a *authDelivery) GetUsersCompany(ctx context.Context, request *auth_service.UserIdRequest) (*auth_service.GetPrivateCompanyResponse, error) {
	response, err := a.authUsecase.GetUsersCompany(ctx, request.Id)
	if err != nil {
		return &auth_service.GetPrivateCompanyResponse{}, a.errorAdapter.AdaptError(err)
	}

	return &auth_service.GetPrivateCompanyResponse{
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
	}, nil
}

func (a *authDelivery) GetCompanyUserLink(ctx context.Context, request *auth_service.UserAndCompanyIdsRequest) (*auth_service.GetCompanyUserLinkResponse, error) {
	response, err := a.authUsecase.GetCompanyUserLink(ctx, request.UserId, request.CompanyId)
	if err != nil {
		return &auth_service.GetCompanyUserLinkResponse{}, a.errorAdapter.AdaptError(err)
	}

	return &auth_service.GetCompanyUserLinkResponse{
		Post:      response.Post,
		CompanyId: response.CompanyId,
		UserId:    response.UserId,
		Itn:       response.Itn,
	}, nil
}

func (a *authDelivery) UpdateUserBalance(ctx context.Context, request *auth_service.UpdateUserBalanceRequest) (*auth_service.GetPublicUserResponse, error) {
	response, err := a.authUsecase.UpdateUserBalance(ctx, request.UserId, request.Balance)
	if err != nil {
		return &auth_service.GetPublicUserResponse{}, a.errorAdapter.AdaptError(err)
	}

	return &auth_service.GetPublicUserResponse{
		Name:       response.Name,
		Surname:    response.Surname,
		Patronymic: response.Patronymic,
		Email:      response.Email,
	}, nil
}

func (a *authDelivery) AddPayment(ctx context.Context, request *auth_service.AddPaymentRequest) (*auth_service.PaymentResponse, error) {
	response, err := a.authUsecase.AddPayment(ctx, &models.Payment{
		UserId:    request.UserId,
		PaymentId: request.PaymentId,
		Amount:    request.Amount,
	})
	if err != nil {
		return &auth_service.PaymentResponse{}, a.errorAdapter.AdaptError(err)
	}

	return &auth_service.PaymentResponse{
		UserId:    response.UserId,
		PaymentId: response.PaymentId,
		Amount:    response.Amount,
		Status:    response.Status,
		Paid:      response.Paid,
		Type:      response.Type,
		Time:      fmt.Sprint(response.Time),
	}, nil
}

func (a *authDelivery) UpdatePayment(ctx context.Context, request *auth_service.UpdatePaymentRequest) (*auth_service.PaymentResponse, error) {
	response, err := a.authUsecase.UpdatePayment(ctx, &models.Payment{
		UserId:    request.UserId,
		PaymentId: request.PaymentId,
		Amount:    request.Amount,
		Status:    request.Status,
		Paid:      request.Paid,
		Type:      request.Type,
	})
	if err != nil {
		return &auth_service.PaymentResponse{}, a.errorAdapter.AdaptError(err)
	}

	return &auth_service.PaymentResponse{
		UserId:    response.UserId,
		PaymentId: response.PaymentId,
		Amount:    response.Amount,
		Status:    response.Status,
		Paid:      response.Paid,
		Type:      response.Type,
		Time:      fmt.Sprint(response.Time),
	}, nil
}

func (a *authDelivery) GetPayment(ctx context.Context, request *auth_service.GetPaymentRequest) (*auth_service.PaymentResponse, error) {
	response, err := a.authUsecase.GetPayment(ctx, request.PaymentId)
	if err != nil {
		return &auth_service.PaymentResponse{}, a.errorAdapter.AdaptError(err)
	}

	return &auth_service.PaymentResponse{
		UserId:    response.UserId,
		PaymentId: response.PaymentId,
		Amount:    response.Amount,
		Status:    response.Status,
		Paid:      response.Paid,
		Type:      response.Type,
		Time:      fmt.Sprint(response.Time),
	}, nil
}

func (a *authDelivery) GetUsersPayments(ctx context.Context, request *auth_service.UserIdRequest) (*auth_service.PaymentsResponse, error) {
	response, err := a.authUsecase.GetUsersPayments(ctx, request.Id)
	if err != nil {
		return &auth_service.PaymentsResponse{}, a.errorAdapter.AdaptError(err)
	}
	var respPayment *auth_service.PaymentResponse
	var paymnets auth_service.PaymentsResponse
	for _, payment := range *response {
		respPayment = &auth_service.PaymentResponse{
			UserId:    payment.UserId,
			PaymentId: payment.PaymentId,
			Amount:    payment.Amount,
			Status:    payment.Status,
			Paid:      payment.Paid,
			Type:      payment.Type,
			Time:      fmt.Sprint(payment.Time),
		}

		paymnets.Payments = append(paymnets.Payments, respPayment)
	}
	return &paymnets, nil
}

func (a *authDelivery) HandlePaidPayments(ctx context.Context, request *auth_service.HandlePaidPaymentsRequest) (*auth_service.HandlePaidPaymentsResponse, error) {
	response, err := a.authUsecase.HandlePaidPayments(ctx, request.UserId)
	if err != nil {
		return &auth_service.HandlePaidPaymentsResponse{}, a.errorAdapter.AdaptError(err)
	}

	return &auth_service.HandlePaidPaymentsResponse{
		Credited: response,
	}, nil
}

func NewAuthDelivery(
	authUsecase usecase.AuthUseCase,
	errorAdapter error_adapter.ErrorAdapter,
) auth_service.AuthServiceServer {
	return &authDelivery{
		authUsecase:  authUsecase,
		errorAdapter: errorAdapter,
	}
}
