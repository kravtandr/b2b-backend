package usecase

import (
	"context"

	"b2b/m/internal/services/fastOrder/models"
)

type fastOrderRepository interface {
	FastOrder(ctx context.Context, user *models.OrderForm) error
}
