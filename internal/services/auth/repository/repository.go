package repository

import (
	"b2b/m/internal/services/auth/models"
	"b2b/m/pkg/errors"
	"context"

	"github.com/jackc/pgx/v4"
	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, ID int64) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	//UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	CreateUserSession(ctx context.Context, userID int64, hash string) error
	ValidateUserSession(ctx context.Context, hash string) (int64, error)
	RemoveUserSession(ctx context.Context, hash string) error
	GetUserInfo(ctx context.Context, id int) (*models.User, error)
}

//type AuthRepository interface {
//	Add(user User) error
//	//AddCompany(company Company) error
//	GetByEmail(key string) (value User, err error)
//	GetPublicUserByEmail(key string) (value PublicUser, err error)
//	GetPublicUserById(id string) (value PublicUser, err error)
//	//SearchCompanies(key string) (value Companies, err error)
//	//GetCompanyById(key string) (value Company, err error)
//	//GetCompaniesByCategoryId(key string) (value Companies, err error)
//
//}

type authRepository struct {
	queryFactory QueryFactory
	conn         *pgxpool.Pool
}

func (a *authRepository) GetUserByID(ctx context.Context, ID int64) (*models.User, error) {
	query := a.queryFactory.CreateGetUserByID(ID)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	user := &models.User{}
	if err := row.Scan(
		&user.Name, &user.Surname, &user.Email,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.UserDoesNotExist
		}

		return nil, err
	}

	return user, nil
}

func (a *authRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := a.queryFactory.CreateGetUserByEmail(email)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	user := &models.User{}
	if err := row.Scan(&user.Id, &user.Password); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.UserDoesNotExist
		}

		return nil, err
	}

	return user, nil
}

func (a *authRepository) GetUserInfo(ctx context.Context, id int) (*models.User, error) {
	conn, err := a.conn.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	var user models.User
	err = conn.QueryRow(context.Background(),
		GetUserInfoQuery,
		id,
	).Scan(&user.Id, &user.Name, &user.Surname)
	if err != nil {
		return nil, err
	}

	return &user, nil
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
		return err
	}

	return nil
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
