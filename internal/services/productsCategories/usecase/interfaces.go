package usecase

import (
	"context"

	"b2b/m/internal/services/productsCategories/models"
)

type productsCategoriesRepository interface {
	GetCategoryById(ctx context.Context, CategoryId *models.CategorieId) (*models.Category, error)
}
