package repository

const (
	createGetCategoryById  = "SELECT id, name, description From categories WHERE id = $1"
	createGetAllCategories = "SELECT id, name From categories"
	createSearchCategories = "SELECT id, name, description FROM categories WHERE name ~ $1"
)
