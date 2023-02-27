package repository

const (
	createGetCategoryById  = "SELECT id, name, description From categories WHERE id = $1"
	createGetAllCategories = "SELECT id, name From categories"
	createSearchCategories = "SELECT id, name, description FROM categories WHERE name ~ $1"
	createGetProductsList  = "SELECT id, name, description,price, photo FROM products OFFSET $1 LIMIT $2"
	createSearchProducts   = "SELECT id, name, description,price, photo FROM products WHERE name ~ $1 OFFSET $2 LIMIT $3"
)
