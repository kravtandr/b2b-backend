package repository

const (
	CreateFastOrder = "INSERT INTO OrderForm (role, product_category, product_name, order_text, order_comments, fio, email, phone, company_name, itn) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
)
