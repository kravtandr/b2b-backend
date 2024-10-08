package usecase

import (
	"context"
	"log"

	company_usecase "b2b/m/internal/gateway/company/usecase"
	"b2b/m/internal/models"
	auth_service "b2b/m/pkg/services/auth"
	company_service "b2b/m/pkg/services/company"

	yookassa "github.com/rvinnie/yookassa-sdk-go/yookassa"
	yoopcommon "github.com/rvinnie/yookassa-sdk-go/yookassa/common"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
)

type UserUsecase interface {
	Register(ctx context.Context, request *models.RegisterUserRequest) (*models.CompanyWithCookie, error)
	FastRegister(ctx context.Context, request *models.FastRegistrationForm) (*models.CompanyWithCookie, error)
	Login(ctx context.Context, request *models.LoginUserRequest) (*models.CompanyWithCookie, error)
	Logout(ctx context.Context, cookie string) error

	GetUserInfo(ctx context.Context, id int64) (*models.Profile, error)
	Profile(ctx context.Context, userID int64) (*models.Profile, error)
	GetUserIdByCookie(ctx context.Context, hash string) (int64, error)
	UpdateUserBalance(ctx context.Context, userID int64, newBalance int64) (*models.UpdateUserResponse, error)
	UpdateProfile(ctx context.Context, userID int64, request *models.PublicCompanyAndOwnerRequest) (*models.PublicCompanyAndOwnerResponse, error)
	AddUserBalance(ctx context.Context, userID int64, add int64) (*models.UpdateUserResponse, error)

	CheckEmail(ctx context.Context, request *models.Email) (*models.PublicUser, error)

	CreatePayment(ctx context.Context, request *models.CreatePaymentRequest) (*yoopayment.Payment, error)
	GetPaymentInfo(ctx context.Context, paymentID string) (*yoopayment.Payment, error)
	ConfirmPayment(ctx context.Context, paymentID string) (*yoopayment.Payment, error)
	CancelPayment(ctx context.Context, paymentID string) (bool, error)

	GetUsersPayments(ctx context.Context, userID int64) (*models.Payments, error)
	HandlePaidPayments(ctx context.Context, userID int64) (bool, error)
	CountUsersPayments(ctx context.Context, userID int64) (*models.PaymentsAmount, error)
}

type userUsecase struct {
	AuthGRPC       AuthGRPC
	companyGRPC    company_usecase.CompanyGRPC
	paymentHandler *yookassa.PaymentHandler
}

func (u *userUsecase) CountUsersPayments(ctx context.Context, userID int64) (*models.PaymentsAmount, error) {
	response, err := u.AuthGRPC.CountUsersPayments(ctx, &auth_service.UserIdRequest{Id: userID})
	if err != nil {
		log.Println("Gateway -> Usecase -> CountUsersPayments -> u.AuthGRPC.CountUsersPayments ERROR", err)
		return &models.PaymentsAmount{}, err
	}
	return &models.PaymentsAmount{Amount: response.Amount}, nil
}

func (u *userUsecase) HandlePaidPayments(ctx context.Context, userID int64) (bool, error) {
	amount, err := u.CountUsersPayments(ctx, userID)
	if err != nil {
		log.Println("Gateway -> Usecase -> HandlePaidPayments -> u.CountUsersPayments ERROR", err)
		return false, err
	}
	if amount.Amount > 0 {
		log.Println("Gateway -> Usecase -> HandlePaidPayments -> AMOUNT", amount)
		payments, err := u.GetUsersPayments(ctx, userID)
		if err != nil {
			log.Println("Gateway -> Usecase -> HandlePaidPayments -> u.GetUsersPayments ERROR", err)
			return false, err
		}
		if len(*payments) == 0 {
			log.Println("Gateway -> Usecase -> HandlePaidPayments -> No  payments found")
			return false, nil
		}
		log.Println("Gateway -> Usecase -> HandlePaidPayments -> PAYMENTS", payments)
		for _, payment := range *payments {
			// TODO check correct work with payment status
			if payment.Status == "pending" {
				_, err := u.ConfirmPayment(ctx, payment.ID)
				if err != nil {
					log.Println("Gateway -> Usecase -> HandlePaidPayments -> u.ConfirmPayment ERROR", err)
					return false, err
				}
			}
		}
		log.Println("Gateway -> Usecase -> HandlePaidPayments -> PAYMENTS CONFIRMED")
		credited, err := u.AuthGRPC.HandlePaidPayments(ctx, &auth_service.HandlePaidPaymentsRequest{
			UserId: userID,
		})
		log.Println("Gateway -> Usecase -> HandlePaidPayments -> PAYMENTS CREDITED", credited)
		if err != nil {
			log.Println("Gateway -> Usecase -> HandlePaidPayments -> u.AuthGRPC.HandlePaidPayments ERROR", err)
			return false, err
		}
		return credited.Credited, nil
	} else {
		log.Println("Gateway -> Usecase -> HandlePaidPayments -> No payments to handle")
		return false, nil
	}

}

func (u *userUsecase) CreatePayment(ctx context.Context, request *models.CreatePaymentRequest) (*yoopayment.Payment, error) {
	// Создаем платеж
	payment, err := u.paymentHandler.CreatePayment(&yoopayment.Payment{
		Amount: &yoopcommon.Amount{
			Value:    request.Amount,
			Currency: "RUB",
		},
		PaymentMethod: yoopayment.PaymentMethodType("bank_card"),
		Confirmation: yoopayment.Redirect{
			Type:      "redirect",
			ReturnURL: "https://bi-tu-bi.ru/profile",
		},
		Description: "Пополннение баланса: " + request.Amount + " руб.",
	})
	if err != nil {
		return payment, err
	}

	_, err = u.AuthGRPC.AddPayment(ctx, &auth_service.AddPaymentRequest{
		PaymentId: payment.ID,
		UserId:    request.User_id,
		Amount:    payment.Amount.Value,
	})
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (u *userUsecase) ConfirmPayment(ctx context.Context, paymentID string) (*yoopayment.Payment, error) {
	getPayment, err := u.GetPaymentInfo(ctx, paymentID)
	if err != nil {
		return nil, err
	}
	if getPayment.Status == "waiting_for_capture" && getPayment.Paid {
		log.Println("confirm: ", getPayment.ID)
		// status = "waiting_for_capture"
		getPayment, err = u.paymentHandler.CapturePayment(getPayment)
		if err != nil {
			log.Println("Gateway -> Usecase -> ConfirmPayment -> u.paymentHandler.CapturePayment ERROR", getPayment.ID, err)
			return nil, err
		}
		getRepoPayment, err := u.AuthGRPC.GetPayment(ctx, &auth_service.GetPaymentRequest{
			PaymentId: paymentID,
		})
		if err != nil {
			log.Println("Gateway -> Usecase -> ConfirmPayment -> u.AuthGRPC.GetPayment ERROR", getRepoPayment.UserId, err)
			return nil, err
		}
		// status = "succeeded"
		_, err = u.AuthGRPC.UpdatePayment(ctx, &auth_service.UpdatePaymentRequest{
			UserId:    getRepoPayment.UserId,
			PaymentId: paymentID,
			Amount:    getPayment.Amount.Value,
			Status:    "succeeded",
			Paid:      getPayment.Paid,
		})
		if err != nil {
			log.Println("Gateway -> Usecase -> ConfirmPayment -> u.AuthGRPC.UpdatePayment ERROR", err)
			return nil, err
		}

	}
	return getPayment, nil
}

func (u *userUsecase) CancelPayment(ctx context.Context, paymentID string) (bool, error) {
	getPayment, err := u.paymentHandler.CancelPayment(paymentID)
	if err != nil {
		return false, err
	}

	getRepoPayment, err := u.AuthGRPC.GetPayment(ctx, &auth_service.GetPaymentRequest{
		PaymentId: paymentID,
	})
	if err != nil {
		return err == nil, err
	}

	_, err = u.AuthGRPC.UpdatePayment(ctx, &auth_service.UpdatePaymentRequest{
		UserId:    getRepoPayment.UserId,
		PaymentId: paymentID,
		Amount:    getPayment.Amount.Value,
		Status:    "cancelled",
		Paid:      getPayment.Paid,
	})
	if err != nil {
		return err == nil, err
	}

	return err == nil, err
}

func (u *userUsecase) GetPaymentInfo(ctx context.Context, paymentID string) (*yoopayment.Payment, error) {
	// get payment info
	getPayment, err := u.paymentHandler.FindPayment(paymentID)
	if err != nil {
		return nil, err
	}

	return getPayment, nil
}

func (u *userUsecase) GetUsersPayments(ctx context.Context, userID int64) (*models.Payments, error) {
	response, err := u.AuthGRPC.GetUsersPayments(ctx, &auth_service.UserIdRequest{Id: userID})
	if err != nil {
		return nil, err
	}
	var payment models.Payment
	var payments models.Payments
	for _, result := range response.Payments {
		payment = models.Payment{
			ID:     result.PaymentId,
			Status: yoopayment.Status(result.Status),
			Amount: &yoopcommon.Amount{Value: result.Amount, Currency: "RUB"},
			Paid:   result.Paid,
		}
		payments = append(payments, payment)
	}
	return &payments, nil
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

func (u *userUsecase) Register(ctx context.Context, request *models.RegisterUserRequest) (*models.CompanyWithCookie, error) {
	_, err := u.AuthGRPC.RegisterUser(ctx, &auth_service.RegisterRequest{
		Email:    request.Email,
		Password: request.Password,
		Name:     request.Name,
		Surname:  request.Surname,
	})
	if err != nil {
		return nil, err
	}

	loginResponse, err := u.AuthGRPC.LoginUser(ctx, &auth_service.LoginRequest{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}

	return &models.CompanyWithCookie{
		Cookie:    loginResponse.Cookie,
		Token:     loginResponse.Token,
		Name:      loginResponse.Name,
		Email:     loginResponse.Email,
		LegalName: loginResponse.LegalName,
		Itn:       loginResponse.Itn,
		OwnerId:   loginResponse.OwnerId,
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
	userInfo, err := u.AuthGRPC.GetUser(ctx, &auth_service.GetUserRequest{Id: int64(userID)})
	if err != nil {
		return nil, err
	}
	usersCompany, err := u.AuthGRPC.GetUsersCompany(ctx, &auth_service.UserIdRequest{Id: int64(userID)})
	if err != nil {
		return nil, err
	}
	companyUserLink, err := u.AuthGRPC.GetCompanyUserLink(ctx, &auth_service.UserAndCompanyIdsRequest{
		UserId:    userID,
		CompanyId: usersCompany.Id,
	})
	if err != nil {
		return nil, err
	}

	return &models.Profile{
		Id:          userID,
		Name:        userInfo.Name,
		Surname:     userInfo.Surname,
		Patronymic:  userInfo.Patronymic,
		Country:     userInfo.Country,
		Balance:     userInfo.Balance,
		Email:       userInfo.Email,
		Avatar:      "TODO",
		Description: "TODO",
		Company: models.Company{
			Id:           usersCompany.Id,
			Name:         usersCompany.Name,
			Description:  usersCompany.Description,
			LegalName:    usersCompany.LegalName,
			Itn:          usersCompany.Itn,
			Psrn:         usersCompany.Psrn,
			Address:      usersCompany.Address,
			LegalAddress: usersCompany.LegalAddress,
			Email:        usersCompany.Email,
			Phone:        usersCompany.Phone,
			Link:         usersCompany.Link,
			Activity:     usersCompany.Activity,
			OwnerId:      usersCompany.OwnerId,
			Rating:       usersCompany.Rating,
			Verified:     usersCompany.Verified,
			Photo:        usersCompany.Photo,
		},
		CompanyUser: models.CompanyUser{
			Post:      companyUserLink.Post,
			CompanyId: companyUserLink.CompanyId,
			UserId:    companyUserLink.UserId,
			Itn:       companyUserLink.Itn,
		},
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

func (u *userUsecase) GetUserInfo(ctx context.Context, id int64) (*models.Profile, error) {
	log.Println("start GetUserInfo")
	response, err := u.AuthGRPC.GetUserInfo(ctx, &auth_service.GetUserRequest{Id: int64(id)})
	if err != nil {
		return nil, err
	}
	log.Println("got GetUserInfo", response)

	return &models.Profile{
		Id:      response.UserId,
		Name:    response.Name,
		Surname: response.Surname,
	}, nil
}

func (u *userUsecase) UpdateUserBalance(ctx context.Context, userID int64, newBalance int64) (*models.UpdateUserResponse, error) {
	updatedUser, err := u.AuthGRPC.UpdateUserBalance(ctx, &auth_service.UpdateUserBalanceRequest{
		UserId:  userID,
		Balance: newBalance,
	})
	if err != nil {
		return &models.UpdateUserResponse{}, err
	}
	return &models.UpdateUserResponse{
		Name:       updatedUser.Name,
		Surname:    updatedUser.Surname,
		Patronymic: updatedUser.Patronymic,
		Email:      updatedUser.Email,
	}, nil
}

func (u *userUsecase) AddUserBalance(ctx context.Context, userID int64, add int64) (*models.UpdateUserResponse, error) {

	user, err := u.AuthGRPC.GetUser(ctx, &auth_service.GetUserRequest{Id: int64(userID)})
	if err != nil {
		return &models.UpdateUserResponse{}, err
	}
	updatedUser, err := u.AuthGRPC.UpdateUserBalance(ctx, &auth_service.UpdateUserBalanceRequest{
		UserId:  userID,
		Balance: user.Balance + add,
	})
	if err != nil {
		return &models.UpdateUserResponse{}, err
	}
	return &models.UpdateUserResponse{
		Name:       updatedUser.Name,
		Surname:    updatedUser.Surname,
		Patronymic: updatedUser.Patronymic,
		Email:      updatedUser.Email,
	}, nil
}

func NewUserUsecase(AuthGRPC AuthGRPC, companyGRPC company_usecase.CompanyGRPC, paymentHandler *yookassa.PaymentHandler) UserUsecase {
	return &userUsecase{AuthGRPC: AuthGRPC, companyGRPC: companyGRPC, paymentHandler: paymentHandler}
}
