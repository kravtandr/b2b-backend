package usecase

import (
	"context"

	"b2b/m/internal/services/productsCategories/models"
)

type productsCategoriesRepository interface {
	GetCategoryById(ctx context.Context, CategoryId *models.CategoryId) (*models.Category, error)
	SearchCategories(ctx context.Context, search string) (*[]models.Category, error)
}
