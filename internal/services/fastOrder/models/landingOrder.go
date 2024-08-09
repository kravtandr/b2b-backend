package models

type LandingForm struct {
	Product_category string `json:"product_category"`
	Delivery_address string `json:"delivery_address"`
	Delivery_date    string `json:"delivery_date"`
	Order_text       string `json:"order_text"`
	Email            string `json:"email"`
	Itn              string `json:"itn"`
}

type LandingForms []LandingForm
