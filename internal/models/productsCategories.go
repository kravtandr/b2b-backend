package models

import "database/sql"

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
	Name        string `json:"name"`
	Description string `json:"info"`
	Price       int64  `json:"price"`
	Photo       string `json:"product_photo"`
	Docs        string `json:"docs"`
	CategoryId  int64  `json:"category_id"`
	Amount      int64  `json:"amount"`
	PayWay      string `json:"payWay"`
	DeliveryWay string `json:"deliveryWay"`
	Adress      string `json:"adress"`
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
	Photo       string         `json:"photo"`
}

type GetProduct struct {
	Id          int64          `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Price       int64          `json:"price"`
	Photo       string         `json:"photo"`
}

type GetProductsList []GetProduct

type SearchCategoriesResponse struct {
	SearchResults []GetCategoryByIdResponse `json:"searchresults"`
}

type SearchProductsResponse struct {
	SearchResults []GetCategoryByIdResponse `json:"searchresults"`
}
