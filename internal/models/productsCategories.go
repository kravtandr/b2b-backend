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

type SearchCategory struct {
	Name string `json:"name"`
}

type SearchCategoriesResponse struct {
	SearchResults []GetCategoryByIdResponse `json:"searchresults"`
}
