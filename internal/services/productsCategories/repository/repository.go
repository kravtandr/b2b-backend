package repository

import (
	"b2b/m/internal/services/productsCategories/models"
	"b2b/m/pkg/errors"
	"context"
	"github.com/jackc/pgx/v4"
	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

type ProductsCategoriesRepository interface {
	GetCategoryById(ctx context.Context, CategoryId *models.CategorieId) (*models.Category, error)
}

type productsCategoriesRepository struct {
	queryFactory QueryFactory
	conn         *pgxpool.Pool
}

func (a productsCategoriesRepository) GetCategoryById(ctx context.Context, CategoryId *models.CategorieId) (*models.Category, error) {
	query := a.queryFactory.CreateGetCategoryById(CategoryId.Id)
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	category := &models.Category{}
	if err := row.Scan(
		&category.Id, &category.Name,
	); err != nil {
		if err == pgx.ErrNoRows {
			return &models.Category{}, errors.UserDoesNotExist
		}
		return &models.Category{}, err
	}
	return category, nil
}

func (a productsCategoriesRepository) GetAllCategories(ctx context.Context) (*models.Categories, error) {
	query := a.queryFactory.CreateGetAllCategories()
	row := a.conn.QueryRow(ctx, query.Request, query.Params...)

	category := &models.Categories{}
	if err := row.Scan(
		&category.Id, &category.Name,
	); err != nil {
		if err == pgx.ErrNoRows {
			return &models.Category{}, errors.UserDoesNotExist
		}
		return &models.Category{}, err
	}
	return category, nil
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
