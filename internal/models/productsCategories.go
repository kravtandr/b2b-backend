package models

type GetCategorieByIdRequest struct {
	Id int64 `json:"id"`
}

type GetCategorieByIdResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
