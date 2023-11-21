package usecase

import (
	"context"

	"b2b/m/internal/services/productsCategories/models"
	chttp "b2b/m/pkg/customhttp"
)

type productsCategoriesRepository interface {
	AddProduct(ctx context.Context, Product *models.Product) (*models.Product, error)
	AddProductsCategoriesLink(ctx context.Context, productId int64, categoryId int64) error
	AddCompaniesProductsLink(ctx context.Context, CompaniesProducts *models.CompaniesProducts) error
	GetCategoryById(ctx context.Context, CategoryId *models.CategoryId) (*models.Category, error)
	GetProductById(ctx context.Context, ProductId *models.ProductId) (*models.Product, error)
	SearchCategories(ctx context.Context, SearchBody *chttp.SearchItemNameWithSkipLimit) (*[]models.Category, error)
	GetProductsList(ctx context.Context, SkipLimit *chttp.QueryParam) (*models.Products, error)
	GetProductsListByFilters(ctx context.Context, filters *models.ProductsFilters) (*models.ProductsWithCategory, error)
	SearchProducts(ctx context.Context, SearchBody *chttp.SearchItemNameWithSkipLimit) (*models.Products, error)
}
