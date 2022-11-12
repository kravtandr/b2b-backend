package repository

import (
	"b2b/m/pkg/query"
)

type QueryFactory interface {
	CreateGetCategoryById(id int64) *query.Query
}

type queryFactory struct{}

func (q *queryFactory) CreateGetCategoryById(id int64) *query.Query {
	return &query.Query{
		Request: createGetCategoryById,
		Params:  []interface{}{id},
	}
}

func NewQueryFactory() QueryFactory {
	return &queryFactory{}
}
