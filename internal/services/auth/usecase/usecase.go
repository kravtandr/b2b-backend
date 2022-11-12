package usecase

import (
	"b2b/m/internal/services/auth/models"
	company_models "b2b/m/internal/services/company/models"
	"b2b/m/pkg/errors"
	"b2b/m/pkg/generator"
	"context"

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
	GetUserInfo(ctx context.Context, id int) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type authUseCase struct {
	hashGenerator hasher
	repo          authRepository
	uuidGen       generator.UUIDGenerator
}

func (a *authUseCase) LoginUser(ctx context.Context, user *models.User) (models.CompanyWithCookie, error) {
	repoUser, err := a.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return models.CompanyWithCookie{}, err
	}

	pass, _ := a.hashGenerator.DecodeString(repoUser.Password)
	if pass != user.Password {
		return models.CompanyWithCookie{}, errors.WrongUserPassword
	}

	cookie := a.uuidGen.GenerateString()
	if err = a.repo.CreateUserSession(ctx, repoUser.Id, cookie); err != nil {
		return models.CompanyWithCookie{}, err
	}

	userCompany, err := a.repo.GetUserCompany(ctx, repoUser.Id)
	if err != nil {
		return models.CompanyWithCookie{}, err
	}

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
		//Email:        userCompany.Email,
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
	if user.Password != "" {
		user.Password = a.hashGenerator.EncodeString(user.Password)
	} else {
		user.Password = currentUser.Password
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

func (a *authUseCase) GetUserInfo(ctx context.Context, id int) (*models.User, error) {
	return a.repo.GetUserInfo(ctx, id)
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
