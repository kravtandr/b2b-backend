package repository

const (
	createGetCategoryById  = "SELECT id, name From categories WHERE id = $1"
	createGetAllCategories = "SELECT id, name From categories"
)
