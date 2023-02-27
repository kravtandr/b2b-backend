package models

import "database/sql"

type GetCategoryByIdRequest struct {
	Id int64 `json:"id"`
}

type GetCategoryByIdResponse struct {
	Id          int64          `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
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
