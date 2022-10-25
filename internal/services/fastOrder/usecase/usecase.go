package usecase

import (
	"context"

	"b2b/m/internal/services/fastOrder/models"
)

type FastOrderUseCase interface {
	FastOrder(ctx context.Context, OrderForm *models.OrderForm) error
}

type fastOrderUseCase struct {
	repo fastOrderRepository
}

func (a *fastOrderUseCase) FastOrder(ctx context.Context, OrderForm *models.OrderForm) error {
	err := a.repo.FastOrder(ctx, OrderForm)
	if err != nil {
		return err
	}
	return nil
}

func NewFastOrderUseCase(
	repo fastOrderRepository,
) FastOrderUseCase {
	return &fastOrderUseCase{
		repo: repo,
	}
}
