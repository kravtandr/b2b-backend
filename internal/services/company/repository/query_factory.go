package repository

import (
	"b2b/m/pkg/query"
)

type QueryFactory interface {
	CreateGetCompanyByID(ID int64) *query.Query
}

type queryFactory struct{}

func (q *queryFactory) CreateGetCompanyByID(ID int64) *query.Query {
	return &query.Query{
		Request: getCompanyByIDRequest,
		Params:  []interface{}{ID},
	}
}

func NewQueryFactory() QueryFactory {
	return &queryFactory{}
}
