package repository

import (
	"b2b/m/internal/services/productsCategories/models"
	chttp "b2b/m/pkg/customhttp"
	"b2b/m/pkg/query"
)

type QueryFactory interface {
	CreateGetCategoryById(id int64) *query.Query
	CreateGetProductById(id int64) *query.Query
	CreateGetAllCategories() *query.Query
	CreateSearchCategories(SearchBody *chttp.SearchItemNameWithSkipLimit) *query.Query
	CreateGetProductsList(SkipLimit *chttp.QueryParam) *query.Query
	CreateSearchProducts(SearchBody *chttp.SearchItemNameWithSkipLimit) *query.Query
	CreateAddProduct(Product *models.Product) *query.Query
	CreateAddProductsCategoriesLink(productId int64, categoryId int64) *query.Query
	CreateAddCompaniesProductsLink(CompaniesProducts *models.CompaniesProducts) *query.Query
}

type queryFactory struct{}

func (q *queryFactory) CreateAddProduct(Product *models.Product) *query.Query {
	return &query.Query{
		Request: createAddProduct,
		Params:  []interface{}{Product.Name, Product.Description, Product.Price, Product.Photo, Product.Docs}, //5
	}
}

func (q *queryFactory) CreateAddProductsCategoriesLink(productId int64, categoryId int64) *query.Query {
	return &query.Query{
		Request: createAddProductsCategoriesLink,
		Params:  []interface{}{productId, categoryId},
	}
}

func (q *queryFactory) CreateAddCompaniesProductsLink(CompaniesProducts *models.CompaniesProducts) *query.Query {
	return &query.Query{
		Request: createAddCompaniesProductsLink,
		Params:  []interface{}{CompaniesProducts.CompanyId, CompaniesProducts.ProductId, CompaniesProducts.AddedBy, CompaniesProducts.Amount, CompaniesProducts.PayWay, CompaniesProducts.DeliveryWay, CompaniesProducts.Adress}, //7
	}
}

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

func (q *queryFactory) CreateSearchCategories(SearchBody *chttp.SearchItemNameWithSkipLimit) *query.Query {
	return &query.Query{
		Request: createSearchCategories,
		Params:  []interface{}{SearchBody.Name, SearchBody.Skip, SearchBody.Limit},
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
