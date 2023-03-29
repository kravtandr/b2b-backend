package repository

import (
	"b2b/m/internal/services/productsCategories/models"
	chttp "b2b/m/pkg/customhttp"
	"b2b/m/pkg/errors"
	"context"

	"github.com/jackc/pgx/v4"
	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

type ProductsCategoriesRepository interface {
	GetCategoryById(ctx context.Context, CategoryId *models.CategoryId) (*models.Category, error)
	GetProductById(ctx context.Context, ProductId *models.ProductId) (*models.Product, error)
	SearchCategories(ctx context.Context, SearchBody *chttp.SearchItemNameWithSkipLimit) (*[]models.Category, error)
	GetProductsList(ctx context.Context, SkipLimit *chttp.QueryParam) (*models.ProductsList, error)
	SearchProducts(ctx context.Context, SearchBody *chttp.SearchItemNameWithSkipLimit) (*models.ProductsList, error)
}

type productsCategoriesRepository struct {
	queryFactory QueryFactory
	conn         *pgxpool.Pool
}

func (a productsCategoriesRepository) GetCategoryById(ctx context.Context, CategoryId *models.CategoryId) (*models.Category, error) {
	query := a.queryFactory.CreateGetCategoryById(CategoryId.Id)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	category := &models.Category{}
	if err := row.Scan(
		&category.Id, &category.Name, &category.Description,
	); err != nil {
		if err == pgx.ErrNoRows {
			return &models.Category{}, errors.CategoryDoesNotExist
		}
		return &models.Category{}, err
	}
	return category, nil
}

func (a productsCategoriesRepository) GetProductById(ctx context.Context, ProductId *models.ProductId) (*models.Product, error) {
	query := a.queryFactory.CreateGetProductById(ProductId.Id)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	product := &models.Product{}
	if err := row.Scan(
		&product.Id, &product.Name, &product.Description, &product.Price, &product.Photo,
	); err != nil {
		if err == pgx.ErrNoRows {
			return &models.Product{}, errors.ProductDoesNotExist
		}
		return &models.Product{}, err
	}
	return product, nil
}

func (a productsCategoriesRepository) SearchCategories(ctx context.Context, SearchBody *chttp.SearchItemNameWithSkipLimit) (*[]models.Category, error) {
	query := a.queryFactory.CreateSearchCategories(SearchBody)
	var category models.Category
	var categories []models.Category
	rows, err := a.conn.Query(ctx, query.Request, query.Params...)
	if err != nil {
		return &categories, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&category.Id, &category.Name, &category.Description)
		categories = append(categories, category)
	}
	if rows.Err() != nil {
		return &categories, err
	}
	return &categories, err
}

func (a productsCategoriesRepository) GetProductsList(ctx context.Context, SkipLimit *chttp.QueryParam) (*models.ProductsList, error) {
	query := a.queryFactory.CreateGetProductsList(SkipLimit)
	var product models.Product
	var products models.ProductsList
	rows, err := a.conn.Query(ctx, query.Request, query.Params...)
	if err != nil {
		return &products, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Photo)
		products = append(products, product)
	}
	if rows.Err() != nil {
		return &products, err
	}
	return &products, err
}

func (a productsCategoriesRepository) SearchProducts(ctx context.Context, SearchBody *chttp.SearchItemNameWithSkipLimit) (*models.ProductsList, error) {
	query := a.queryFactory.CreateSearchProducts(SearchBody)
	var product models.Product
	var products models.ProductsList
	rows, err := a.conn.Query(ctx, query.Request, query.Params...)
	if err != nil {
		return &products, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Photo)
		products = append(products, product)
	}
	if rows.Err() != nil {
		return &products, err
	}
	return &products, err
}

func NewProductsCategoriesRepository(
	queryFactory QueryFactory,
	conn *pgxpool.Pool,
) ProductsCategoriesRepository {
	return &productsCategoriesRepository{
		queryFactory: queryFactory,
		conn:         conn,
	}
}
