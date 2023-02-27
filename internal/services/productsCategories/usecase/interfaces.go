package usecase

import (
	"context"

	"b2b/m/internal/services/productsCategories/models"
	chttp "b2b/m/pkg/customhttp"
)

type productsCategoriesRepository interface {
	GetCategoryById(ctx context.Context, CategoryId *models.CategoryId) (*models.Category, error)
	SearchCategories(ctx context.Context, search string) (*[]models.Category, error)
	GetProductsList(ctx context.Context, SkipLimit *chttp.QueryParam) (*models.ProductsList, error)
	SearchProducts(ctx context.Context, SearchBody *chttp.SearchItemNameWithSkipLimit) (*models.ProductsList, error)
}
