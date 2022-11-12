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
func (l *loggingMiddleware) GetCompanyByOwnerIdAndItn(ctx context.Context, company models.Company) (c *models.Company, err error) {
	l.logger.Infow(module,
		"Action", "GetCompanyByOwnerIdAndItn",
		"Request", company,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetCompanyByOwnerIdItnAndLegalName",
				"Request", company,
				"Error", err,
			)
		}
	}()

	return l.next.GetCompanyByOwnerIdAndItn(ctx, company)
}

func (l *loggingMiddleware) GetCompanyUserLinkByOwnerIdAndItn(ctx context.Context, id int64, itn string) (c *models.CompaniesUsersLink, err error) {
	l.logger.Infow(module,
		"Action", "GetCompanyUserLinkByOwnerIdAndItn",
		"Request", id, itn,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetCompanyUserLinkByOwnerIdAndItn",
				"Request", id, itn,
				"Error", err,
			)
		}
	}()

	return l.next.GetCompanyUserLinkByOwnerIdAndItn(ctx, id, itn)
}

func (l *loggingMiddleware) UpdateCompanyById(ctx context.Context, newCompany models.Company) (c *models.Company, err error) {
	l.logger.Infow(module,
		"Action", "UpdateCompanyById",
		"Request", newCompany,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "UpdateCompanyById",
				"Request", newCompany,
				"Error", err,
			)
		}
	}()

	return l.next.UpdateCompanyById(ctx, newCompany)
}

func (l *loggingMiddleware) UpdateCompanyUsersLink(ctx context.Context, companyId int64, userId int64, post string) (newPost string, err error) {
	l.logger.Infow(module,
		"Action", "UpdateCompanyUsersLink",
		"Request", companyId, userId, post,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "UpdateCompanyUsersLink",
				"Request", companyId, userId, post,
				"Error", err,
			)
		}
	}()

	return l.next.UpdateCompanyUsersLink(ctx, companyId, userId, post)
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
