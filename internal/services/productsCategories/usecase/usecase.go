package usecase

import (
	"context"

	"b2b/m/internal/services/productsCategories/models"
)

type ProductsCategoriesUseCase interface {
	GetCategoryById(ctx context.Context, CategoryId *models.CategoryId) (*models.Category, error)
	SearchCategories(ctx context.Context, search string) (*[]models.Category, error)
}

type productsCategoriesUseCase struct {
	repo productsCategoriesRepository
}

func (a *productsCategoriesUseCase) GetCategoryById(ctx context.Context, CategoryId *models.CategoryId) (*models.Category, error) {
	category, err := a.repo.GetCategoryById(ctx, CategoryId)
	if err != nil {
		return &models.Category{}, err
	}
	return category, nil
}

func (a *productsCategoriesUseCase) SearchCategories(ctx context.Context, search string) (*[]models.Category, error) {
	categories, err := a.repo.SearchCategories(ctx, search)
	if err != nil {
		return &[]models.Category{}, err
	}
	return categories, nil
}

func NewProductsCategoriesUseCase(
	repo productsCategoriesRepository,
) ProductsCategoriesUseCase {
	return &productsCategoriesUseCase{
		repo: repo,
	}
}
