package repository

import (
	"b2b/m/internal/services/auth/models"
	company_models "b2b/m/internal/services/company/models"
	"b2b/m/pkg/errors"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	FastRegistration(ctx context.Context, newCompany *company_models.Company, user *models.User, post string) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, ID int64) (*models.User, error)
	GetUsersCompany(ctx context.Context, userId int64) (*company_models.Company, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserInfo(ctx context.Context, id int64) (*models.User, error)

	CreateUserSession(ctx context.Context, userID int64, hash string) error
	ValidateUserSession(ctx context.Context, hash string) (int64, error)
	RemoveUserSession(ctx context.Context, hash string) error

	GetCompanyUserLink(ctx context.Context, userId int64, companyId int64) (*company_models.CompaniesUsersLink, error)

	UpdateUserBalance(ctx context.Context, userID int64, newBalance int64) (*models.User, error)

	AddPayment(ctx context.Context, payment *models.Payment) (*models.Payment, error)
	GetPayment(ctx context.Context, paymentID string) (*models.Payment, error)
	GetUsersPayments(ctx context.Context, userID int64) (*models.Payments, error)
	CountUsersPayments(ctx context.Context, userID int64) (int, error)
	UpdatePayment(ctx context.Context, payment *models.Payment) (*models.Payment, error)
}

type authRepository struct {
	queryFactory QueryFactory
	conn         *pgxpool.Pool
}

func (a *authRepository) AddPayment(ctx context.Context, payment *models.Payment) (*models.Payment, error) {

	query := a.queryFactory.CreateAddPayment(payment)
	if _, err := a.conn.Exec(ctx, query.Request, query.Params...); err != nil {
		log.Println("ERROR: authRepository->AddPayment", err)
		return nil, err
	}
	payment, err := a.GetPayment(ctx, payment.PaymentId)
	if err != nil {
		log.Println("ERROR: authRepository->AddPayment, GetPayment err: ", err)
		return nil, err
	}

	return payment, nil
}

func (a *authRepository) UpdatePayment(ctx context.Context, payment *models.Payment) (*models.Payment, error) {
	query := a.queryFactory.CreateUpdatePayment(payment)
	if _, err := a.conn.Exec(ctx, query.Request, query.Params...); err != nil {
		log.Println("ERROR: authRepository->UpdatePayment", err)
		return nil, err
	}
	payment, err := a.GetPayment(ctx, payment.PaymentId)
	if err != nil {
		log.Println("ERROR: authRepository->UpdatePayment, GetPayment err: ", err)
		return nil, err
	}

	return payment, nil
}

func (a *authRepository) GetPayment(ctx context.Context, paymentID string) (*models.Payment, error) {
	query := a.queryFactory.CreateGetPayment(paymentID)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)
	payment := &models.Payment{}
	if err := row.Scan(
		&payment.Id, &payment.UserId, &payment.PaymentId, &payment.Amount, &payment.Status, &payment.Paid, &payment.Type, &payment.Credited, &payment.Time,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.UserDoesNotExist
		}

		return nil, err
	}

	return payment, nil
}

func (a *authRepository) CountUsersPayments(ctx context.Context, userID int64) (int, error) {
	query := a.queryFactory.CreateCountUsersPayments(userID)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)
	var count int
	if err := row.Scan(
		&count,
	); err != nil {
		if err == pgx.ErrNoRows {
			return 0, nil
		}
		log.Println("ERROR ", err)
		return -1, err
	}
	return count, nil
}

func (a *authRepository) GetUsersPayments(ctx context.Context, userID int64) (*models.Payments, error) {

	payments_amount, err := a.CountUsersPayments(ctx, userID)
	if err != nil {
		log.Println("ERROR ", err)
		return nil, err
	}
	if payments_amount == 0 {
		log.Println("No payments found for user ID: ", userID)
		return nil, nil
	}
	log.Println("payments_amount: ", payments_amount, " userID: ", userID)

	query := a.queryFactory.CreateGetUsersPayments(userID)
	var payment models.Payment
	var payments models.Payments
	// var payments = make(models.Payments, payments_amount)
	log.Println("query.Request: ", query.Request)
	rows, err := a.conn.Query(ctx, query.Request, query.Params...)
	if err != nil {
		log.Println("ERROR ", err)
		return &payments, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&payment.Id, &payment.UserId, &payment.PaymentId, &payment.Amount, &payment.Status, &payment.Paid, &payment.Type, &payment.Credited, &payment.Time)
		payments = append(payments, payment)
	}
	if rows.Err() != nil {
		log.Println("ERROR ", err)
		return &payments, err
	}
	log.Println("status: ", rows.Err())
	log.Println("Repo result OK Result: ", payments)
	return &payments, nil

}

func (a *authRepository) FastRegistration(ctx context.Context, newCompany *company_models.Company, user *models.User, post string) error {
	query := a.queryFactory.CreateCreateCompany(user, newCompany)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)
	if err := row.Scan(&newCompany.Id); err != nil {
		log.Println("ERROR: authRepository->FastRegistration->CreateCreateCompany")
		return err
	}

	query = a.queryFactory.CreateCreateUserCompanyLink(user, newCompany, post)
	_, err := a.conn.Exec(ctx, query.Request, query.Params...)
	if err != nil {
		log.Println("ERROR: authRepository->FastRegistration->CreateCreateUserCompanyLink")
		return nil
	}
	return nil
}

func (a *authRepository) GetUserByID(ctx context.Context, ID int64) (*models.User, error) {
	query := a.queryFactory.CreateGetUserByID(ID)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	user := &models.User{}
	if err := row.Scan(
		&user.Id, &user.Name, &user.Surname, &user.Patronymic, &user.Country, &user.Email, &user.Password, &user.GroupId, &user.Balance,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.UserDoesNotExist
		}

		return nil, err
	}

	return user, nil
}

func (a *authRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	log.Println("in repo GetUserByEmail")
	query := a.queryFactory.CreateGetUserByEmail(email)
	log.Println("in repo CreateGetUserByEmail")
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)
	log.Println("in repo GetUserByEmail QueryRow")
	user := &models.User{}
	log.Println("repo GetUserByEmail Scan start", user)
	if err := row.Scan(&user.Id, &user.Name, &user.Surname, &user.Patronymic, &user.Email, &user.Password); err != nil {
		if err == pgx.ErrNoRows {
			log.Println("ERROR: authRepository->GetUserByEmail, UserDoesNotExist err: ", err)
			return nil, errors.UserDoesNotExist
		}
		log.Println("ERROR: authRepository->GetUserByEmail, err: ", err)
		return nil, err
	}
	log.Println("repo GetUserByEmail return - OK User: ", user)
	return user, nil
}

func (a *authRepository) GetUserInfo(ctx context.Context, id int64) (*models.User, error) {
	log.Println("authRepository -> GetUserInfo", id)
	query := a.queryFactory.CreateGetUserByID(id)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	user := &models.User{}
	if err := row.Scan(&user.Id, &user.Name, &user.Surname, &user.Patronymic, &user.Email, &user.Password); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.UserDoesNotExist
		}

		return nil, err
	}

	return user, nil
}

func (a *authRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	query := a.queryFactory.CreateCreateUser(user)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	if err := row.Scan(&user.Id); err != nil {
		return nil, err
	}

	return user, nil
}

func (a *authRepository) CreateUserSession(ctx context.Context, userID int64, hash string) error {
	query := a.queryFactory.CreateCreateUserSession(userID, hash)
	if _, err := a.conn.Exec(ctx, query.Request, query.Params...); err != nil {
		return err
	}

	return nil
}

func (a *authRepository) GetUsersCompany(ctx context.Context, userId int64) (*company_models.Company, error) {
	// conn, err := a.conn.Acquire(context.Background())
	// if err != nil {
	// 	return nil, err
	// }
	// defer conn.Release()
	var company company_models.Company
	err := a.conn.QueryRow(context.Background(),
		createCompanyByUserId,
		userId,
	).Scan(&company.Id, &company.Name, &company.Description, &company.LegalName, &company.Itn, &company.Psrn, &company.Address, &company.LegalAddress, &company.Email, &company.Phone, &company.Link, &company.Activity, &company.OwnerId, &company.Rating, &company.Verified)
	if err != nil {
		return nil, err
	}

	return &company, nil
}

func (a *authRepository) GetCompanyUserLink(ctx context.Context, userId int64, companyId int64) (*company_models.CompaniesUsersLink, error) {
	// conn, err := a.conn.Acquire(context.Background())
	// if err != nil {
	// 	return nil, err
	// }
	// defer conn.Release()

	var cULink company_models.CompaniesUsersLink
	err := a.conn.QueryRow(context.Background(),
		createGetCompanyUserLink,
		userId, companyId,
	).Scan(&cULink.Id, &cULink.Post, &cULink.CompanyId, &cULink.UserId, &cULink.Itn)
	if err != nil {
		log.Println("ERROR: authRepository->GetCompanyUserLink", err)
		return nil, err
	}

	return &cULink, nil
}

func (a *authRepository) ValidateUserSession(ctx context.Context, hash string) (int64, error) {
	userID := int64(0)
	query := a.queryFactory.CreateValidateUserSession(hash)
	if err := a.conn.QueryRow(ctx, query.Request, query.Params...).Scan(&userID); err != nil {
		if err == pgx.ErrNoRows {
			return userID, errors.SessionDoesNotExist
		}

		return userID, err
	}

	return userID, nil
}

func (a *authRepository) RemoveUserSession(ctx context.Context, hash string) error {
	query := a.queryFactory.CreateRemoveUserSession(hash)
	if _, err := a.conn.Exec(ctx, query.Request, query.Params...); err != nil {
		log.Println("ERROR: authRepository->RemoveUserSession", err)
		return err
	}

	return nil
}

func (a *authRepository) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	query := a.queryFactory.CreateUpdateUser(user)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	updatedUser := &models.User{}
	if err := row.Scan(
		&updatedUser.Id, &updatedUser.Name, &updatedUser.Surname, &updatedUser.Patronymic, &updatedUser.Email, &updatedUser.Password,
	); err != nil {
		if err == pgx.ErrNoRows {
			fmt.Println("Error", err)
			return &models.User{}, errors.UserDoesNotExist
		}

		return &models.User{}, err
	}
	return updatedUser, nil
}

func (a *authRepository) UpdateUserBalance(ctx context.Context, userID int64, newBalance int64) (*models.User, error) {
	query := a.queryFactory.CreateUpdateUserBalance(userID, newBalance)

	if _, err := a.conn.Exec(ctx, query.Request, query.Params...); err != nil {
		log.Println("ERROR: authRepository->RemoveUserSession", err)
		return nil, err
	}
	updatedUser, err := a.GetUserByID(ctx, userID)
	if err != nil {
		log.Println("ERROR: authRepository->UpdateUserBalance, GetUserByID err: ", err)
		return nil, err
	}
	return updatedUser, nil
}

func NewAuthRepository(
	queryFactory QueryFactory,
	conn *pgxpool.Pool,
) AuthRepository {
	return &authRepository{
		queryFactory: queryFactory,
		conn:         conn,
	}
}
