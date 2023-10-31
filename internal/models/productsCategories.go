package models

import "database/sql"

type GetProductsByFilters struct {
	Product_name       string  `json:"product_name"`
	Category_name      string  `json:"category_name"`
	Categories_ids     []int64 `json:"categories_ids"`
	Price_lower_limit  int64   `json:"price_lower_limit"`
	Price_higher_limit int64   `json:"price_higher_limit"`
}

type GetCategoryByIdRequest struct {
	Id int64 `json:"id"`
}

type GetProductByIdRequest struct {
	Id int64 `json:"id"`
}

type GetCategoryByIdResponse struct {
	Id          int64          `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
}

type AddProductByFormRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"info"`
	Price       int64    `json:"price"`
	Photo       []string `json:"product_photo"`
	Docs        []string `json:"docs"`
	CategoryId  int64    `json:"category_id"`
	Amount      int64    `json:"amount"`
	PayWay      string   `json:"payWay"`
	DeliveryWay string   `json:"deliveryWay"`
	Adress      string   `json:"adress"`
}

type UserInfoAndAddProductByFormRequest struct {
	Product     AddProductByFormRequest
	UserProfile Profile
}

type GetProductByIdResponse struct {
	Id          int64          `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Price       int64          `json:"price"`
	Photo       []string       `json:"photo"`
	Docs        []string       `json:"docs"`
	Company     Company        `json:"company"`
}

type GetProduct struct {
	Id          int64          `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Price       int64          `json:"price"`
	Photo       []string       `json:"photo"`
	Docs        []string       `json:"docs"`
}

type GetProductsList []GetProduct

type SearchCategoriesResponse struct {
	SearchResults []GetCategoryByIdResponse `json:"searchresults"`
}

type SearchProductsResponse struct {
	SearchResults []GetCategoryByIdResponse `json:"searchresults"`
}
