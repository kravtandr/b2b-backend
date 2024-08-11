package models

import (
	gateway_models "b2b/m/internal/models"
	chttp "b2b/m/pkg/customhttp"
	"database/sql"
)

type CategoryId struct {
	Id int64
}

type ProductId struct {
	Id int64
}

type Category struct {
	Id          int64
	Name        string
	Description sql.NullString
}

type Categories []Category

type Product struct {
	Id          int64
	Name        string
	Description sql.NullString
	Price       int64
	Photo       []string
	Docs        []string
	CreatedAt   string
	UpdatedAt   string
}

type ProductsFilters struct {
	Product_name       string
	Category_name      string
	Categories_ids     []int64
	Price_lower_limit  int64
	Price_higher_limit int64
	QueryParam         chttp.QueryParam
}

type CompaniesProducts struct {
	Id          int64
	CompanyId   int64
	ProductId   int64
	AddedBy     int64
	Amount      int64
	PayWay      string
	DeliveryWay string
	Adress      string
}
type Products []Product

type ProductWithCategory struct {
	Product  Product
	Category Category
}

type ProductsWithCategory []ProductWithCategory

type AddProductByFormRequest struct {
	Name        string
	Description sql.NullString
	Price       int64
	Photo       []string
	Docs        []string
	CompanyId   int64
	CategoryId  int64
	Amount      int64
	PayWay      string
	DeliveryWay string
	Adress      string
}

type UserInfoAndAddProductByFormRequest struct {
	Product     AddProductByFormRequest
	UserProfile gateway_models.Profile
}
