package usecase

import (
	"b2b/m/internal/services/auth/models"
	company_models "b2b/m/internal/services/company/models"
	"b2b/m/pkg/errors"
	"b2b/m/pkg/generator"
	"context"
	"log"
	"strconv"

	"github.com/gofrs/uuid"
)

type AuthUseCase interface {
	LoginUser(ctx context.Context, user *models.User) (models.CompanyWithCookie, error)
	LogoutUser(ctx context.Context, session string) error
	ValidateSession(ctx context.Context, session string) (int64, error)
	FastRegistration(ctx context.Context, form *models.FastRegistrationForm) (models.CompanyWithCookie, error)
	RegisterUser(ctx context.Context, user *models.User) (models.Session, error)
	GetUser(ctx context.Context, ID int64) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.PublicUser, error)
	GetUserInfo(ctx context.Context, id int64) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUsersCompany(ctx context.Context, userId int64) (*company_models.Company, error)
	GetCompanyUserLink(ctx context.Context, userId int64, companyId int64) (*company_models.CompaniesUsersLink, error)
	UpdateUserBalance(ctx context.Context, userId int64, newBalance int64) (*models.PublicUser, error)
	AddUserBalance(ctx context.Context, userId int64, amount int64) (*models.PublicUser, error)
	AddPayment(ctx context.Context, payment *models.Payment) (*models.Payment, error)
	UpdatePayment(ctx context.Context, payment *models.Payment) (*models.Payment, error)
	GetPayment(ctx context.Context, paymentID string) (*models.Payment, error)
	GetUsersPayments(ctx context.Context, userID int64) (*models.Payments, error)
	HandlePaidPayments(ctx context.Context, userID int64) (bool, error)
}

type authUseCase struct {
	hashGenerator hasher
	repo          authRepository
	uuidGen       generator.UUIDGenerator
}

func (a *authUseCase) AddUserBalance(ctx context.Context, userId int64, amount int64) (*models.PublicUser, error) {
	user, err := a.repo.GetUserByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	updatedUser, err := a.repo.UpdateUserBalance(ctx, userId, user.Balance+amount)
	if err != nil {
		return nil, err
	}
	return &models.PublicUser{
		Name:       updatedUser.Name,
		Surname:    updatedUser.Surname,
		Patronymic: updatedUser.Patronymic,
		Email:      updatedUser.Email,
	}, nil
}

func (a *authUseCase) AddPayment(ctx context.Context, payment *models.Payment) (*models.Payment, error) {
	return a.repo.AddPayment(ctx, payment)
}

func (a *authUseCase) UpdatePayment(ctx context.Context, payment *models.Payment) (*models.Payment, error) {
	return a.repo.UpdatePayment(ctx, payment)
}

func (a *authUseCase) GetPayment(ctx context.Context, paymentID string) (*models.Payment, error) {
	return a.repo.GetPayment(ctx, paymentID)
}

func (a *authUseCase) GetUsersPayments(ctx context.Context, userID int64) (*models.Payments, error) {
	return a.repo.GetUsersPayments(ctx, userID)
}

func (a *authUseCase) HandlePaidPayments(ctx context.Context, userID int64) (bool, error) {
	log.Println(" in service usecase -> authUseCase -> HandlePaidPayments", userID)
	credited := false
	payments, err := a.GetUsersPayments(ctx, userID)
	if err != nil {
		log.Println("ERROR: GetUsersPayments", err)
		return false, err
	}
	log.Println("GetUsersPayments", payments)
	// if user have payments
	if len(*payments) > 0 {
		for _, payment := range *payments {
			if payment.Status == "succeeded" && payment.Paid && !payment.Credited {
				// "1000.00" to int64 1000
				// Step 1: Parse the string to a float64
				log.Println("Step 1: Parse the string to a float64")
				floatValue, err := strconv.ParseFloat(payment.Amount, 64)
				if err != nil {
					return false, err
				}
				// Step 2: Convert the float64 to int64
				log.Println("Step 2: Convert the float64 to int64")
				amount := int64(floatValue)

				log.Println("AddUserBalance", userID, amount)
				_, err = a.AddUserBalance(ctx, userID, amount)
				if err != nil {
					log.Println("ERROR: AddUserBalance", err)
					return false, err
				}
				log.Panicln("UpdatePayment", payment.Id)
				_, err = a.repo.UpdatePayment(ctx, &models.Payment{
					Id:        payment.Id,
					UserId:    payment.UserId,
					PaymentId: payment.PaymentId,
					Amount:    payment.Amount,
					Status:    payment.Status,
					Paid:      payment.Paid,
					Credited:  true, // add balance
					Time:      payment.Time,
				})
				if err != nil {
					log.Println("ERROR: UpdatePayment", err)
					return false, err
				}
				credited = true

			}
		}

	} else {
		log.Println("No payments for user")
		return credited, nil
	}
	return credited, nil
}

func (a *authUseCase) LoginUser(ctx context.Context, user *models.User) (models.CompanyWithCookie, error) {
	log.Println("LoginUser in service usecase", user)

	repoUser, err := a.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		log.Println("ERROR GetUserByEmail  in service usecase ", err)
		return models.CompanyWithCookie{}, err
	}
	log.Println("GetUserByEmail  in service usecase ", repoUser)

	pass, _ := a.hashGenerator.DecodeString(repoUser.Password)
	if pass != user.Password {
		log.Println("WrongUserPassword  in service usecase ", err)
		return models.CompanyWithCookie{}, errors.WrongUserPassword
	}

	cookie := a.uuidGen.GenerateString()
	if err = a.repo.CreateUserSession(ctx, repoUser.Id, cookie); err != nil {
		log.Println("ERROR CreateUserSession  in service usecase ", err)
		return models.CompanyWithCookie{}, err
	}
	log.Println("CreateUserSession  in service usecase ")

	userCompany, err := a.repo.GetUsersCompany(ctx, repoUser.Id)
	if err != nil {
		log.Println("ERROR GetUsersCompany  in service usecase ", err)
		return models.CompanyWithCookie{}, err
	}
	log.Println("GetUsersCompany  in service usecase ")
	log.Println("return LoginUser  in service usecase ")

	return models.CompanyWithCookie{
		Cookie:       cookie,
		Token:        "??",
		Name:         userCompany.Name,
		Description:  userCompany.Description,
		LegalName:    userCompany.LegalName,
		Itn:          userCompany.Itn,
		Psrn:         userCompany.Psrn,
		Address:      userCompany.Address,
		LegalAddress: userCompany.LegalAddress,
		Email:        userCompany.Email,
		Phone:        userCompany.Phone,
		Link:         userCompany.Link,
		Activity:     userCompany.Activity,
		OwnerId:      userCompany.OwnerId,
		Rating:       userCompany.Rating,
		Verified:     userCompany.Verified,
	}, nil
}

func (a *authUseCase) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := a.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *authUseCase) FastRegistration(ctx context.Context, form *models.FastRegistrationForm) (models.CompanyWithCookie, error) {
	form.Password = a.hashGenerator.EncodeString(form.Password)
	newUser := models.User{
		Name:       form.OwnerName,
		Email:      form.Email,
		Password:   form.Password,
		Surname:    form.Surname,
		Patronymic: form.Patronymic,
		Country:    form.Country,
	}
	user, err := a.repo.CreateUser(ctx, &newUser)
	if err != nil {
		return models.CompanyWithCookie{}, err
	}

	cookie := a.uuidGen.GenerateString()
	if err = a.repo.CreateUserSession(ctx, user.Id, cookie); err != nil {
		return models.CompanyWithCookie{}, err
	}

	newCompany := company_models.Company{
		Name:      form.Name,
		LegalName: form.LegalName,
		Itn:       form.Itn,
		Email:     form.Email,
	}

	err = a.repo.FastRegistration(ctx, &newCompany, user, form.Post)
	if err != nil {
		return models.CompanyWithCookie{}, err
	}

	return models.CompanyWithCookie{
		Cookie: cookie,
		Token:  "??",
		Name:   newCompany.Name,
		//Description:  userCompany.Description,
		LegalName: newCompany.LegalName,
		Itn:       newCompany.Itn,
		//Psrn:         userCompany.Psrn,
		//Address:      userCompany.Address,
		//LegalAddress: userCompany.LegalAddress,
		Email: newCompany.Email,
		//Phone:        userCompany.Phone,
		//Link:         userCompany.Link,
		//Activity:     userCompany.Activity,
		OwnerId: user.Id,
		//Rating:       userCompany.Rating,
		//Verified:     userCompany.Verified,
	}, nil
}

func (a *authUseCase) LogoutUser(ctx context.Context, session string) error {
	return a.repo.RemoveUserSession(ctx, session)
}

func (a *authUseCase) ValidateSession(ctx context.Context, session string) (int64, error) {
	return a.repo.ValidateUserSession(ctx, session)
}

func (a *authUseCase) RegisterUser(ctx context.Context, user *models.User) (models.Session, error) {
	user.Password = a.hashGenerator.EncodeString(user.Password)
	user, err := a.repo.CreateUser(ctx, user)
	if err != nil {
		return models.Session{}, err
	}

	cookie := a.uuidGen.GenerateString()
	if err = a.repo.CreateUserSession(ctx, user.Id, cookie); err != nil {
		return models.Session{}, err
	}

	return models.Session{
		Cookie: cookie,
		Token:  "??",
	}, nil
}

func (a *authUseCase) GetUser(ctx context.Context, ID int64) (*models.User, error) {
	return a.repo.GetUserByID(ctx, ID)
}

func (a *authUseCase) UpdateUser(ctx context.Context, user *models.User) (*models.PublicUser, error) {
	currentUser, err := a.repo.GetUserByID(ctx, user.Id)
	if err != nil {
		return &models.PublicUser{}, errors.UserDoesNotExist
	}
	if user.Password != "" {
		user.Password = a.hashGenerator.EncodeString(user.Password)
	} else {
		user.Password = currentUser.Password
	}

	if user.Name == "" {
		user.Name = currentUser.Name
	}

	if user.Surname == "" {
		user.Surname = currentUser.Surname
	}

	if user.Patronymic == "" {
		user.Patronymic = currentUser.Patronymic
	}

	if user.Email == "" {
		user.Email = currentUser.Email
	}

	user.GroupId = currentUser.GroupId
	updatedUser, err := a.repo.UpdateUser(ctx, user)
	if err != nil {
		return &models.PublicUser{}, err
	}

	return &models.PublicUser{
		Name:       updatedUser.Name,
		Surname:    updatedUser.Surname,
		Patronymic: updatedUser.Patronymic,
		Email:      updatedUser.Email,
	}, nil
}

func (a *authUseCase) UpdateUserBalance(ctx context.Context, userId int64, newBalance int64) (*models.PublicUser, error) {
	updatedUser, err := a.repo.UpdateUserBalance(ctx, userId, newBalance)
	if err != nil {
		return &models.PublicUser{}, err
	}
	return &models.PublicUser{
		Name:       updatedUser.Name,
		Surname:    updatedUser.Surname,
		Patronymic: updatedUser.Patronymic,
		Email:      updatedUser.Email,
	}, nil
}

func (a *authUseCase) GetUserInfo(ctx context.Context, id int64) (*models.User, error) {
	log.Println("authUseCase -> GetUserInfo", id)
	return a.repo.GetUserInfo(ctx, id)
}

func (a *authUseCase) GetUsersCompany(ctx context.Context, userId int64) (*company_models.Company, error) {
	return a.repo.GetUsersCompany(ctx, userId)
}

func (a *authUseCase) GetCompanyUserLink(ctx context.Context, userId int64, companyId int64) (*company_models.CompaniesUsersLink, error) {
	return a.repo.GetCompanyUserLink(ctx, userId, companyId)
}

func NewAuthUseCase(
	hashGenerator hasher,
	repo authRepository,
) AuthUseCase {
	return &authUseCase{
		hashGenerator: hashGenerator,
		repo:          repo,
		uuidGen:       generator.NewUUIDGenerator(uuid.NewGen()),
	}
}
