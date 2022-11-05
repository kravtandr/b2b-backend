package repository

import (
	"context"

	"b2b/m/internal/services/company/models"

	"go.uber.org/zap"
)

const (
	module = "company_repo"
)

type loggingMiddleware struct {
	logger *zap.SugaredLogger

	next CompanyRepository
}

func NewLoggingMiddleware(logger *zap.SugaredLogger, next CompanyRepository) CompanyRepository {
	return &loggingMiddleware{
		logger: logger,
		next:   next,
	}
}

func (l *loggingMiddleware) GetCompanyById(ctx context.Context, ID int64) (c *models.Company, err error) {
	l.logger.Infow(module,
		"Action", "GetCompanyById",
		"Request", ID,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetCompanyById",
				"Request", ID,
				"Error", err,
			)
		}
	}()

	return l.next.GetCompanyById(ctx, ID)
}
