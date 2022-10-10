package domain

import "database/sql"

type Category struct {
	Id          int            `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
}
type Categories []Category

type CategoryStorage interface {
	// GetCategoryById(key string) (value Category, err error)
	// GetCategoriesInIndustry(key string) (value Categories, err error)
	GetAllCategories() (value Categories, err error)
	SearchCategories(param string) (value Categories, err error)
}

type CategoryUseCase interface {
	// GetCategoryById(key string) (value []byte, err error)
	// GetCategoriesInIndustry(key string) (value []byte, err error)
	GetAllCategories() (value []byte, err error)
	SearchCategories(param string) (value []byte, err error)
}
