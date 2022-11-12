package repository

import (
	"b2b/m/pkg/query"
)

type QueryFactory interface {
	CreateGetCategoryById(id int64) *query.Query
	CreateGetAllCategories() *query.Query
}

type queryFactory struct{}

func (q *queryFactory) CreateGetCategoryById(id int64) *query.Query {
	return &query.Query{
		Request: createGetCategoryById,
		Params:  []interface{}{id},
	}
}

func (q *queryFactory) CreateGetAllCategories() *query.Query {
	return &query.Query{
		Request: createGetAllCategories,
	}
}

func NewQueryFactory() QueryFactory {
	return &queryFactory{}
}
