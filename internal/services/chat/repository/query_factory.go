package repository

import (
	"b2b/m/internal/services/auth/models"
	company_models "b2b/m/internal/services/company/models"
	"b2b/m/pkg/query"
)

type QueryFactory interface {
	CreateCreateCompany(user *models.User, company *company_models.Company) *query.Query
	CreateCreateUserCompanyLink(user *models.User, company *company_models.Company, post string) *query.Query
	CreateGetUserByEmail(email string) *query.Query
	CreateGetUserByID(ID int64) *query.Query
	CreateCreateUser(user *models.User) *query.Query
	CreateCreateUserSession(userID int64, hash string) *query.Query
	CreateValidateUserSession(hash string) *query.Query
	CreateRemoveUserSession(hash string) *query.Query
	CreateUpdateUser(user *models.User) *query.Query
}

type queryFactory struct{}

func (q *queryFactory) CreateGetCompanyByUserId(Id int64) *query.Query {
	return &query.Query{
		Request: createCompanyByUserId,
		Params:  []interface{}{Id},
	}
}

func (q *queryFactory) CreateUpdateUser(user *models.User) *query.Query {
	return &query.Query{
		Request: createCreateUpdateUser,
		Params: []interface{}{
			user.Id, user.Name, user.Surname, user.Patronymic, user.Email, user.Password,
		},
	}
}

func (q *queryFactory) CreateCreateCompany(user *models.User, company *company_models.Company) *query.Query {
	return &query.Query{
		Request: createCreateCompany,
		Params: []interface{}{
			company.Name, company.LegalName, company.Itn, user.Email, user.Id,
		},
	}
}

func (q *queryFactory) CreateCreateUserCompanyLink(user *models.User, company *company_models.Company, post string) *query.Query {
	return &query.Query{
		Request: createCreateUserCompanyLink,
		Params: []interface{}{
			post, company.Id, user.Id, company.Itn,
		},
	}
}

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
