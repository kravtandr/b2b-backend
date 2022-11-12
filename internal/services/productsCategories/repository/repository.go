package repository

import (
	"b2b/m/internal/services/productsCategories/models"
	"b2b/m/pkg/errors"
	"context"
	"github.com/jackc/pgx/v4"
	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

type ProductsCategoriesRepository interface {
	GetCategoryById(ctx context.Context, CategoryId *models.CategoryId) (*models.Category, error)
	SearchCategories(ctx context.Context, search string) (*[]models.Category, error)
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
			return &models.Category{}, errors.UserDoesNotExist
		}
		return &models.Category{}, err
	}
	return category, nil
}

func (a productsCategoriesRepository) SearchCategories(ctx context.Context, search string) (*[]models.Category, error) {
	query := a.queryFactory.CreateSearchCategories(search)
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

func NewProductsCategoriesRepository(
	queryFactory QueryFactory,
	conn *pgxpool.Pool,
) ProductsCategoriesRepository {
	return &productsCategoriesRepository{
		queryFactory: queryFactory,
		conn:         conn,
	}
}
