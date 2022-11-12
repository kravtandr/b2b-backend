package repository

import (
	"context"

	"b2b/m/internal/services/productsCategories/models"

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

func (l *loggingMiddleware) SearchCategories(ctx context.Context, name string) (c *[]models.Category, err error) {
	l.logger.Infow(module,
		"Action", "SearchCategories",
		"Request", name,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "SearchCategories",
				"Request", name,
				"Error", err,
			)
		}
	}()

	return l.next.SearchCategories(ctx, name)
}
