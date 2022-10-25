package repository

import (
	"context"

	"b2b/m/internal/services/fastOrder/models"

	"go.uber.org/zap"
)

const (
	module = "fastOrder_repo"
)

type loggingMiddleware struct {
	logger *zap.SugaredLogger

	next FastOrderRepository
}

func NewLoggingMiddleware(logger *zap.SugaredLogger, next FastOrderRepository) FastOrderRepository {
	return &loggingMiddleware{
		logger: logger,
		next:   next,
	}
}

func (l *loggingMiddleware) FastOrder(ctx context.Context, order *models.OrderForm) (err error) {
	l.logger.Infow(module,
		"Action", "CreateFastOrder",
		"Request", order,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "CreateFastOrder",
				"Request", order,
				"Error", err,
			)
		}
	}()

	return l.next.FastOrder(ctx, order)
}
