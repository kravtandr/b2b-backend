package usecase

import (
	"context"

	"b2b/m/internal/services/fastOrder/models"
)

type FastOrderUseCase interface {
	FastOrder(ctx context.Context, OrderForm *models.OrderForm) error
	LandingOrder(ctx context.Context, LandingForm *models.LandingForm) error
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

func (a *fastOrderUseCase) LandingOrder(ctx context.Context, LandingForm *models.LandingForm) error {
	err := a.repo.LandingOrder(ctx, LandingForm)
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
