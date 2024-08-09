package usecase

import (
	"context"

	"b2b/m/internal/services/auth/models"
	company_models "b2b/m/internal/services/company/models"
)

type authRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, ID int64) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	FastRegistration(ctx context.Context, newCompany *company_models.Company, user *models.User, post string) error
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	CreateUserSession(ctx context.Context, userID int64, hash string) error
	ValidateUserSession(ctx context.Context, hash string) (int64, error)
	RemoveUserSession(ctx context.Context, hash string) error
	GetUserInfo(ctx context.Context, id int64) (*models.User, error)
	GetUsersCompany(ctx context.Context, userId int64) (*company_models.Company, error)
	GetCompanyUserLink(ctx context.Context, userId int64, companyId int64) (*company_models.CompaniesUsersLink, error)
	UpdateUserBalance(ctx context.Context, userID int64, newBalance int64) (*models.User, error)
	AddPayment(ctx context.Context, payment *models.Payment) (*models.Payment, error)
	UpdatePayment(ctx context.Context, payment *models.Payment) (*models.Payment, error)
	GetPayment(ctx context.Context, paymentID string) (*models.Payment, error)
	GetUsersPayments(ctx context.Context, userID int64) (*models.Payments, error)
	CountUsersPayments(ctx context.Context, userID int64) (int64, error)
}

type hasher interface {
	EncodeString(value string) string
	DecodeString(value string) (string, error)
}
