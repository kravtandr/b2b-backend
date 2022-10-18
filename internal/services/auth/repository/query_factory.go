package repository

import (
	"b2b/m/internal/services/auth/models"
	"b2b/m/pkg/query"
)

type QueryFactory interface {
	CreateGetUserByEmail(email string) *query.Query
	CreateGetUserByID(ID int64) *query.Query
	CreateCreateUser(user *models.User) *query.Query
	CreateCreateUserSession(userID int64, hash string) *query.Query
	CreateValidateUserSession(hash string) *query.Query
	CreateRemoveUserSession(hash string) *query.Query
}

type queryFactory struct{}

func (q *queryFactory) CreateCreateUser(user *models.User) *query.Query {
	return &query.Query{
		Request: createUserRequest,
		Params: []interface{}{
			user.Name, user.Surname, user.Patronymic, user.Email, user.Password, user.Country,
		},
	}
}

func (q *queryFactory) CreateGetUserByEmail(email string) *query.Query {
	return &query.Query{
		Request: getUserByEmailRequest,
		Params:  []interface{}{email},
	}
}

func (q *queryFactory) CreateGetUserByID(ID int64) *query.Query {
	return &query.Query{
		Request: getUserByIDRequest,
		Params:  []interface{}{ID},
	}
}

func (q *queryFactory) CreateCreateUserSession(userID int64, hash string) *query.Query {
	return &query.Query{
		Request: createUserSession,
		Params:  []interface{}{userID, hash},
	}
}

func (q *queryFactory) CreateValidateUserSession(hash string) *query.Query {
	return &query.Query{
		Request: validateUserSession,
		Params:  []interface{}{hash},
	}
}

func (q *queryFactory) CreateRemoveUserSession(hash string) *query.Query {
	return &query.Query{
		Request: removeUserSession,
		Params:  []interface{}{hash},
	}
}

func NewQueryFactory() QueryFactory {
	return &queryFactory{}
}
