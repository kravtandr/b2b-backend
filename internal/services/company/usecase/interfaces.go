package usecase

import (
	"context"

	"b2b/m/internal/services/company/models"
)

type companyRepository interface {
	GetCompanyById(ctx context.Context, ID int64) (*models.Company, error)
}
