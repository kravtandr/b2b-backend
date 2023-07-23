package repository

import (
	"b2b/m/internal/services/company/models"
	"b2b/m/pkg/query"
)

type QueryFactory interface {
	CreateGetCompanyByID(ID int64) *query.Query
	GetCompanyByOwnerIdAndItn(ownerId int64, itn string) *query.Query
	CreateUpdateCompanyById(newCompany models.Company) *query.Query
	CreateUpdateCompanyUsersLink(companyId int64, userId int64, post string) *query.Query
	CreateGetCompanyUserLinkByOwnerIdAndItn(userId int64, itn string) *query.Query
	CreateGetProductsCompaniesLink(ProductId int64) *query.Query
}

type queryFactory struct{}

func (q *queryFactory) CreateGetCompanyByID(ID int64) *query.Query {
	return &query.Query{
		Request: getCompanyByIDRequest,
		Params:  []interface{}{ID},
	}
}

func (q *queryFactory) CreateGetProductsCompaniesLink(ProductId int64) *query.Query {
	return &query.Query{
		Request: createGetProductsCompaniesLink,
		Params:  []interface{}{ProductId},
	}
}

func (q *queryFactory) CreateGetCompanyUserLinkByOwnerIdAndItn(userId int64, itn string) *query.Query {
	return &query.Query{
		Request: createGetCompanyUserLinkByOwnerIdAndItn,
		Params:  []interface{}{userId, itn},
	}
}
func (q *queryFactory) GetCompanyByOwnerIdAndItn(ownerId int64, itn string) *query.Query {
	return &query.Query{
		Request: createGetCompanyByOwnerIdAndItn,
		Params:  []interface{}{ownerId, itn},
	}
}
func (q *queryFactory) CreateUpdateCompanyById(newCompany models.Company) *query.Query {
	return &query.Query{
		Request: createUpdateCompanyById,
		Params:  []interface{}{newCompany.Id, newCompany.Name, newCompany.Description, newCompany.Address, newCompany.LegalAddress, newCompany.Phone, newCompany.Link, newCompany.Activity},
	}
}
func (q *queryFactory) CreateUpdateCompanyUsersLink(companyId int64, userId int64, post string) *query.Query {
	return &query.Query{
		Request: createUpdateCompanyUsersLink,
		Params:  []interface{}{companyId, userId, post},
	}
}

func NewQueryFactory() QueryFactory {
	return &queryFactory{}
}
