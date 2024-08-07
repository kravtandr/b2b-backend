package repository

import (
	"b2b/m/internal/services/productsCategories/models"
	chttp "b2b/m/pkg/customhttp"
	"b2b/m/pkg/query"

	pq_type "github.com/lib/pq"
)

type QueryFactory interface {
	CreateAddProduct(Product *models.Product) *query.Query
	CreateAddProductDocuments(productId int64, objName string) *query.Query
	CreateAddProductPhotos(productId int64, objName string) *query.Query
	CreateGetProductById(id int64) *query.Query
	CreateGetProductPhotos(productId int64) *query.Query
	CreateGetProductDocuments(productId int64) *query.Query
	CreateUpdateProduct(Product *models.Product) *query.Query

	CreateGetProductsList(SkipLimit *chttp.QueryParam) *query.Query
	CreateGetProductsListByFilters(filters *models.ProductsFilters) *query.Query
	CreateSearchProducts(SearchBody *chttp.SearchItemNameWithSkipLimit) *query.Query
	CreateGetCompanyProducts(CompanyId int64, SkipLimit *chttp.QueryParam) *query.Query

	CreateGetCategoryById(id int64) *query.Query
	CreateGetAllCategories() *query.Query

	CreateSearchCategories(SearchBody *chttp.SearchItemNameWithSkipLimit) *query.Query

	CreateAddProductsCategoriesLink(productId int64, categoryId int64) *query.Query
	CreateGetProductsCategoriesLink(productId int64) *query.Query
	CreateUpdateProductsCategoriesLink(productId int64, newCategoryId int64) *query.Query

	CreateAddCompaniesProductsLink(CompaniesProducts *models.CompaniesProducts) *query.Query
}

type queryFactory struct{}

func (q *queryFactory) CreateUpdateProductsCategoriesLink(productId int64, newCategoryId int64) *query.Query {
	return &query.Query{
		Request: createUpdateProductsCategoriesLink,
		Params:  []interface{}{productId, newCategoryId},
	}
}

func (q *queryFactory) CreateUpdateProduct(Product *models.Product) *query.Query {
	return &query.Query{
		Request: createUpdateProduct,
		Params:  []interface{}{Product.Id, Product.Name, Product.Description, Product.Price},
	}
}

func (q *queryFactory) CreateGetProductsCategoriesLink(productId int64) *query.Query {
	return &query.Query{
		Request: createGetProductsCategoriesLink,
		Params:  []interface{}{productId},
	}
}

func (q *queryFactory) CreateGetProductPhotos(productId int64) *query.Query {
	return &query.Query{
		Request: createGetProductPhotos,
		Params:  []interface{}{productId},
	}
}

func (q *queryFactory) CreateGetProductDocuments(productId int64) *query.Query {
	return &query.Query{
		Request: createGetProductPhotos,
		Params:  []interface{}{productId},
	}
}

func (q *queryFactory) CreateAddProductDocuments(productId int64, objName string) *query.Query {
	return &query.Query{
		Request: createAddProductDocuments,
		Params:  []interface{}{productId, objName},
	}
}

func (q *queryFactory) CreateAddProductPhotos(productId int64, objName string) *query.Query {
	return &query.Query{
		Request: createAddProductPhotos,
		Params:  []interface{}{productId, objName},
	}
}

func (q *queryFactory) CreateAddProduct(Product *models.Product) *query.Query {
	return &query.Query{
		Request: createAddProduct,
		Params:  []interface{}{Product.Name, Product.Description, Product.Price},
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

func (q *queryFactory) CreateGetProductsListByFilters(filters *models.ProductsFilters) *query.Query {
	return &query.Query{
		Request: createGetProductsListByFilters,
		Params:  []interface{}{filters.Product_name, filters.Category_name, pq_type.Array(filters.Categories_ids), filters.Price_lower_limit, filters.Price_higher_limit, filters.QueryParam.Skip, filters.QueryParam.Limit},
	}
}

func (q *queryFactory) CreateSearchProducts(SearchBody *chttp.SearchItemNameWithSkipLimit) *query.Query {
	return &query.Query{
		Request: createSearchProducts,
		Params:  []interface{}{SearchBody.Name, SearchBody.Skip, SearchBody.Limit},
	}
}
func (q *queryFactory) CreateGetCompanyProducts(CompanyId int64, SkipLimit *chttp.QueryParam) *query.Query {
	return &query.Query{
		Request: createGetCompanyProducts,
		Params:  []interface{}{CompanyId, SkipLimit.Skip, SkipLimit.Limit},
	}
}

func NewQueryFactory() QueryFactory {
	return &queryFactory{}
}
