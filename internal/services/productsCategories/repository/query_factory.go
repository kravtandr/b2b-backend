package repository

import (
	"b2b/m/pkg/query"
)

type QueryFactory interface {
	CreateGetCategoryById(id int64) *query.Query
	CreateGetAllCategories() *query.Query
	CreateSearchCategories(name string) *query.Query
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

func (q *queryFactory) CreateSearchCategories(name string) *query.Query {
	return &query.Query{
		Request: createSearchCategories,
		Params:  []interface{}{name},
	}
}

func NewQueryFactory() QueryFactory {
	return &queryFactory{}
}
