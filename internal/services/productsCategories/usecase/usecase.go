package usecase

import (
	"b2b/m/internal/services/productsCategories/models"
	chttp "b2b/m/pkg/customhttp"
	"context"
)

type ProductsCategoriesUseCase interface {
	AddProduct(ctx context.Context, Product *models.Product, CompaniesProducts *models.CompaniesProducts, userId int64, companyId int64, categoryId int64) (*models.Product, error)
	GetCategoryById(ctx context.Context, CategoryId *models.CategoryId) (*models.Category, error)
	GetProductById(ctx context.Context, ProductId *models.ProductId) (*models.Product, error)
	SearchCategories(ctx context.Context, SearchBody *chttp.SearchItemNameWithSkipLimit) (*[]models.Category, error)
	GetProductsList(ctx context.Context, SkipLimit *chttp.QueryParam) (*models.ProductsList, error)
	SearchProducts(ctx context.Context, SearchBody *chttp.SearchItemNameWithSkipLimit) (*models.ProductsList, error)
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

func (a *productsCategoriesUseCase) GetProductById(ctx context.Context, ProductId *models.ProductId) (*models.Product, error) {
	category, err := a.repo.GetProductById(ctx, ProductId)
	if err != nil {
		return &models.Product{}, err
	}
	return category, nil
}

func (a *productsCategoriesUseCase) AddProduct(ctx context.Context, Product *models.Product, CompaniesProducts *models.CompaniesProducts, userId int64, companyId int64, categoryId int64) (*models.Product, error) {
	product, err := a.repo.AddProduct(ctx, Product)
	if err != nil {
		return &models.Product{}, err
	}
	err = a.repo.AddProductsCategoriesLink(ctx, Product.Id, categoryId)
	if err != nil {
		return &models.Product{}, err
	}
	CompaniesProducts.ProductId = product.Id
	err = a.repo.AddCompaniesProductsLink(ctx, CompaniesProducts)
	if err != nil {
		return &models.Product{}, err
	}
	return product, nil
}

func (a *productsCategoriesUseCase) SearchCategories(ctx context.Context, SearchBody *chttp.SearchItemNameWithSkipLimit) (*[]models.Category, error) {
	categories, err := a.repo.SearchCategories(ctx, SearchBody)
	if err != nil {
		return &[]models.Category{}, err
	}
	return categories, nil
}

func (a *productsCategoriesUseCase) GetProductsList(ctx context.Context, SkipLimit *chttp.QueryParam) (*models.ProductsList, error) {
	products, err := a.repo.GetProductsList(ctx, SkipLimit)
	if err != nil {
		return &models.ProductsList{}, err
	}
	return products, nil
}

func (a *productsCategoriesUseCase) SearchProducts(ctx context.Context, SearchBody *chttp.SearchItemNameWithSkipLimit) (*models.ProductsList, error) {
	products, err := a.repo.SearchProducts(ctx, SearchBody)
	if err != nil {
		return &models.ProductsList{}, err
	}
	return products, nil
}

func NewProductsCategoriesUseCase(
	repo productsCategoriesRepository,
) ProductsCategoriesUseCase {
	return &productsCategoriesUseCase{
		repo: repo,
	}
}
