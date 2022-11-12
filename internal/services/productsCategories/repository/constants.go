package repository

const (
	createGetCategoryById = "SELECT id, name From categories WHERE id = $1"
)
