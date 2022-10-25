package models

type FastOrderRequest struct {
	Role             bool   `json:"role"`
	Product_category string `json:"product_category"`
	Product_name     string `json:"product_name"`
	Order_text       string `json:"order_text"`
	Order_comments   string `json:"order_comments"`
	Fio              string `json:"fio"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Company_name     string `json:"company_name"`
	Itn              string `json:"itn"`
}
