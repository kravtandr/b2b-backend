package repository

import (
	chttp "b2b/m/pkg/customhttp"
	"b2b/m/pkg/query"
)

type QueryFactory interface {
	CreateGetCategoryById(id int64) *query.Query
	CreateGetProductById(id int64) *query.Query
	CreateGetAllCategories() *query.Query
	CreateSearchCategories(name string) *query.Query
	CreateGetProductsList(SkipLimit *chttp.QueryParam) *query.Query
	CreateSearchProducts(SearchBody *chttp.SearchItemNameWithSkipLimit) *query.Query
}

type queryFactory struct{}

func (q *queryFactory) CreateGetCategoryById(id int64) *query.Query {
	return &query.Query{
		Request: createGetCategoryById,
		Params:  []interface{}{id},
	}
}

func (q *queryFactory) CreateGetProductById(id int64) *query.Query {
	return &query.Query{
		Request: createGetProductById,
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

func (q *queryFactory) CreateGetProductsList(SkipLimit *chttp.QueryParam) *query.Query {
	return &query.Query{
		Request: createGetProductsList,
		Params:  []interface{}{SkipLimit.Skip, SkipLimit.Limit},
	}
}

func (q *queryFactory) CreateSearchProducts(SearchBody *chttp.SearchItemNameWithSkipLimit) *query.Query {
	return &query.Query{
		Request: createSearchProducts,
		Params:  []interface{}{SearchBody.Name, SearchBody.Skip, SearchBody.Limit},
	}
}

func NewQueryFactory() QueryFactory {
	return &queryFactory{}
}
