package usecase

import (
	"b2b/m/internal/services/productsCategories/models"
	chttp "b2b/m/pkg/customhttp"
	"context"
	"log"
)

type ProductsCategoriesUseCase interface {
	AddProduct(ctx context.Context, Product *models.Product, CompaniesProducts *models.CompaniesProducts, userId int64, companyId int64, categoryId int64) (*models.Product, error)
	GetProductById(ctx context.Context, ProductId *models.ProductId) (*models.Product, error)
	UpdateProduct(ctx context.Context, Product *models.Product, CompaniesProducts *models.CompaniesProducts, userId int64, companyId int64, categoryId int64) (*models.Product, error)

	GetProductsList(ctx context.Context, SkipLimit *chttp.QueryParam) (*models.Products, error)
	GetProductsListByFilters(ctx context.Context, filters *models.ProductsFilters) (*models.ProductsWithCategory, error)
	SearchProducts(ctx context.Context, SearchBody *chttp.SearchItemNameWithSkipLimit) (*models.Products, error)
	GetCompanyProducts(ctx context.Context, CompanyId int64, SkipLimit *chttp.QueryParam) (*models.Products, error)

	GetCategoryById(ctx context.Context, CategoryId *models.CategoryId) (*models.Category, error)
	SearchCategories(ctx context.Context, SearchBody *chttp.SearchItemNameWithSkipLimit) (*[]models.Category, error)
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
	log.Println("productsCategoriesUseCase -> AddProduct")
	Product, err := a.repo.AddProduct(ctx, Product)
	if err != nil {
		return &models.Product{}, err
	}
	CompaniesProducts.ProductId = Product.Id
	log.Println("productsCategoriesUseCase -> AddProductsCategoriesLink")
	err = a.repo.AddProductsCategoriesLink(ctx, Product.Id, categoryId)
	if err != nil {
		log.Println("productsCategoriesUseCase -> AddProductsCategoriesLink", err)
		return &models.Product{}, err
	}

	log.Println("productsCategoriesUseCase: GET product id from db = ", CompaniesProducts.ProductId)
	log.Println("productsCategoriesUseCase -> AddCompaniesProductsLink")
	err = a.repo.AddCompaniesProductsLink(ctx, CompaniesProducts)
	if err != nil {
		log.Println("productsCategoriesUseCase -> AddCompaniesProductsLink", err)
		return &models.Product{}, err
	}
	log.Println("OK ||| productsCategoriesUseCase -> AddProduct -> all done. New product")
	// log.Println(Product)
	return Product, nil
}

func (a *productsCategoriesUseCase) UpdateProduct(ctx context.Context, Product *models.Product, CompaniesProducts *models.CompaniesProducts, userId int64, companyId int64, categoryId int64) (*models.Product, error) {
	log.Println("productsCategoriesUseCase -> UpdateProduct")
	Product, err := a.repo.UpdateProduct(ctx, Product)
	if err != nil {
		return &models.Product{}, err
	}
	CompaniesProducts.ProductId = Product.Id
	log.Println("productsCategoriesUseCase -> UpdateProductsCategoriesLink")
	err = a.repo.UpdateProductsCategoriesLink(ctx, Product.Id, categoryId)
	if err != nil {
		log.Println("productsCategoriesUseCase -> UpdateProductsCategoriesLink", err)
		return &models.Product{}, err
	}

	log.Println("productsCategoriesUseCase: GET product id from db = ", CompaniesProducts.ProductId)
	log.Println("productsCategoriesUseCase -> UpdateCompaniesProductsLink")
	err = a.repo.UpdateCompaniesProductsLink(ctx, CompaniesProducts)
	if err != nil {
		log.Println("productsCategoriesUseCase -> UpdateCompaniesProductsLink", err)
		return &models.Product{}, err
	}
	log.Println("OK ||| productsCategoriesUseCase -> UpdateProduct -> all done. updated product")
	// log.Println(Product)
	return Product, nil
}

func (a *productsCategoriesUseCase) SearchCategories(ctx context.Context, SearchBody *chttp.SearchItemNameWithSkipLimit) (*[]models.Category, error) {
	categories, err := a.repo.SearchCategories(ctx, SearchBody)
	if err != nil {
		return &[]models.Category{}, err
	}
	return categories, nil
}
func (a *productsCategoriesUseCase) GetProductsListByFilters(ctx context.Context, filters *models.ProductsFilters) (*models.ProductsWithCategory, error) {
	if filters.Categories_ids == nil {
		filters.Categories_ids = pyrange(1, 1000, 1)
	}
	products, err := a.repo.GetProductsListByFilters(ctx, filters)
	if err != nil {
		return &models.ProductsWithCategory{}, err
	}
	return products, nil
}
func (a *productsCategoriesUseCase) GetProductsList(ctx context.Context, SkipLimit *chttp.QueryParam) (*models.Products, error) {
	products, err := a.repo.GetProductsList(ctx, SkipLimit)
	if err != nil {
		return &models.Products{}, err
	}
	return products, nil
}

func (a *productsCategoriesUseCase) SearchProducts(ctx context.Context, SearchBody *chttp.SearchItemNameWithSkipLimit) (*models.Products, error) {
	products, err := a.repo.SearchProducts(ctx, SearchBody)
	if err != nil {
		return &models.Products{}, err
	}
	return products, nil
}

func (a *productsCategoriesUseCase) GetCompanyProducts(ctx context.Context, CompanyId int64, SkipLimit *chttp.QueryParam) (*models.Products, error) {
	products, err := a.repo.GetCompanyProducts(ctx, CompanyId, SkipLimit)
	if err != nil {
		return &models.Products{}, err
	}
	return products, nil
}

func pyrange(start, end, step int64) []int64 {
	// TODO: Error checking to make sure parameters are all valid,
	// else you could get divide by zero in make and other errors.

	rtn := make([]int64, 0, (end-start)/step)
	for i := start; i < end; i += step {
		rtn = append(rtn, i)
	}
	return rtn
}

func NewProductsCategoriesUseCase(
	repo productsCategoriesRepository,
) ProductsCategoriesUseCase {
	return &productsCategoriesUseCase{
		repo: repo,
	}
}
