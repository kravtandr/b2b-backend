package repository

import (
	"context"

	"b2b/m/internal/services/productsCategories/models"

	chttp "b2b/m/pkg/customhttp"

	"go.uber.org/zap"
)

const (
	module = "productsCategories_repo"
)

type loggingMiddleware struct {
	logger *zap.SugaredLogger

	next ProductsCategoriesRepository
}

func NewLoggingMiddleware(logger *zap.SugaredLogger, next ProductsCategoriesRepository) ProductsCategoriesRepository {
	return &loggingMiddleware{
		logger: logger,
		next:   next,
	}
}

func (l *loggingMiddleware) UpdateProduct(ctx context.Context, Product *models.Product) (product *models.Product, err error) {
	l.logger.Infow(module,
		"Action", "UpdateProduct",
		"Request", Product,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "UpdateProduct",
				"Request", Product,
				"Error", err,
			)
		}
	}()

	return l.next.UpdateProduct(ctx, Product)
}

func (l *loggingMiddleware) UpdateProductsCategoriesLink(ctx context.Context, productId int64, categoryId int64) (err error) {
	l.logger.Infow(module,
		"Action", "UpdateProduct",
		"Request", productId, categoryId,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "UpdateProduct",
				"Request", productId, categoryId,
				"Error", err,
			)
		}
	}()

	return l.next.UpdateProductsCategoriesLink(ctx, productId, categoryId)
}

func (l *loggingMiddleware) GetProductsCategoriesLink(ctx context.Context, productId int64) (prcat *models.ProductWithCategory, err error) {
	l.logger.Infow(module,
		"Action", "GetProductsCategoriesLink",
		"Request", productId,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetProductsCategoriesLink",
				"Request", productId,
				"Error", err,
			)
		}
	}()

	return l.next.GetProductsCategoriesLink(ctx, productId)
}

func (l *loggingMiddleware) UpdateCompaniesProductsLink(ctx context.Context, CompaniesProducts *models.CompaniesProducts) (err error) {
	l.logger.Infow(module,
		"Action", "UpdateCompaniesProductsLink",
		"Request", CompaniesProducts,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "UpdateCompaniesProductsLink",
				"Request", CompaniesProducts,
				"Error", err,
			)
		}
	}()

	return l.next.UpdateCompaniesProductsLink(ctx, CompaniesProducts)
}

func (l *loggingMiddleware) GetCategoryById(ctx context.Context, CategoryId *models.CategoryId) (c *models.Category, err error) {
	l.logger.Infow(module,
		"Action", "GetCategoryById",
		"Request", CategoryId,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetCategoryById",
				"Request", CategoryId,
				"Error", err,
			)
		}
	}()

	return l.next.GetCategoryById(ctx, CategoryId)
}

func (l *loggingMiddleware) AddProduct(ctx context.Context, Product *models.Product) (c *models.Product, err error) {
	l.logger.Infow(module,
		"Action", "AddProduct",
		"Request (name)", Product.Name,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "AddProduct",
				"Request (name)", Product.Name,
				"Error", err,
			)
		}
	}()

	return l.next.AddProduct(ctx, Product)
}

func (l *loggingMiddleware) AddProductsCategoriesLink(ctx context.Context, productId int64, categoryId int64) (err error) {
	l.logger.Infow(module,
		"Action", "AddProductsCategoriesLink",
		"Request", productId,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "AddProductsCategoriesLink",
				"Request", productId,
				"Error", err,
			)
		}
	}()

	return l.next.AddProductsCategoriesLink(ctx, productId, categoryId)
}

func (l *loggingMiddleware) AddCompaniesProductsLink(ctx context.Context, CompaniesProducts *models.CompaniesProducts) (err error) {
	l.logger.Infow(module,
		"Action", "AddCompaniesProductsLink",
		"Request", CompaniesProducts,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "AddCompaniesProductsLink",
				"Request", CompaniesProducts,
				"Error", err,
			)
		}
	}()

	return l.next.AddCompaniesProductsLink(ctx, CompaniesProducts)
}

func (l *loggingMiddleware) GetProductById(ctx context.Context, ProductId *models.ProductId) (c *models.Product, err error) {
	l.logger.Infow(module,
		"Action", "GetProductById",
		"Request", ProductId,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetProductById",
				"Request", ProductId,
				"Error", err,
			)
		}
	}()

	return l.next.GetProductById(ctx, ProductId)
}

func (l *loggingMiddleware) SearchCategories(ctx context.Context, SearchBody *chttp.SearchItemNameWithSkipLimit) (c *[]models.Category, err error) {
	l.logger.Infow(module,
		"Action", "SearchCategories",
		"Request", SearchBody,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "SearchCategories",
				"Request", SearchBody,
				"Error", err,
			)
		}
	}()

	return l.next.SearchCategories(ctx, SearchBody)
}

func (l *loggingMiddleware) SearchProducts(ctx context.Context, SearchBody *chttp.SearchItemNameWithSkipLimit) (c *models.Products, err error) {
	l.logger.Infow(module,
		"Action", "SearchProducts",
		"Request", SearchBody,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "SearchProducts",
				"Request", SearchBody,
				"Error", err,
			)
		}
	}()

	return l.next.SearchProducts(ctx, SearchBody)
}

func (l *loggingMiddleware) GetProductsList(ctx context.Context, SkipLimit *chttp.QueryParam) (c *models.Products, err error) {
	l.logger.Infow(module,
		"Action", "GetProductsList",
		"Request", SkipLimit,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetProductsList",
				"Request", SkipLimit,
				"Error", err,
			)
		}
	}()

	return l.next.GetProductsList(ctx, SkipLimit)
}

func (l *loggingMiddleware) GetProductsListByFilters(ctx context.Context, filters *models.ProductsFilters) (c *models.ProductsWithCategory, err error) {
	l.logger.Infow(module,
		"Action", "GetProductsList",
		"Request", filters,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetProductsList",
				"Request", filters,
				"Error", err,
			)
		}
	}()

	return l.next.GetProductsListByFilters(ctx, filters)
}

func (l *loggingMiddleware) GetCompanyProducts(ctx context.Context, CompanyId int64, SkipLimit *chttp.QueryParam) (c *models.Products, err error) {
	l.logger.Infow(module,
		"Action", "GetCompanyProducts",
		"Request", CompanyId, SkipLimit,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetCompanyProducts",
				"Request", CompanyId, SkipLimit,
				"Error", err,
			)
		}
	}()

	return l.next.GetCompanyProducts(ctx, CompanyId, SkipLimit)
}
