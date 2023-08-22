package models

import (
	gateway_models "b2b/m/internal/models"
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
	Docs        string
	CreatedAt   string
	UpdatedAt   string
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

type ProductsList []Product

type AddProductByFormRequest struct {
	Name        string         `json:"name"`
	Description sql.NullString `json:"info"`
	Price       int64          `json:"price"`
	Photo       []string       `json:"product_photo"`
	Docs        string         `json:"docs"`
	CompanyId   int64          `json:"company_id"`
	CategoryId  int64          `json:"category_id"`
	Amount      int64          `json:"amount"`
	PayWay      string         `json:"payWay"`
	DeliveryWay string         `json:"deliveryWay"`
	Adress      string         `json:"adress"`
}

type UserInfoAndAddProductByFormRequest struct {
	Product     AddProductByFormRequest
	UserProfile gateway_models.Profile
}
