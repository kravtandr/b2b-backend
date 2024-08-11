package usecase

import (
	fastOrder_service "b2b/m/pkg/services/fastOrder"
	"context"

	"b2b/m/internal/models"
)

type FastOrderUseCase interface {
	FastOrder(ctx context.Context, request *models.FastOrderRequest) error
	LandingOrder(ctx context.Context, request *models.LandingOrderRequest) error
}

type fastOrderUseCase struct {
	fastOrderGRPC fastOrderGRPC
}

func (u *fastOrderUseCase) FastOrder(ctx context.Context, request *models.FastOrderRequest) error {
	_, err := u.fastOrderGRPC.FastOrder(ctx, &fastOrder_service.FastOrderRequest{
		Role:            request.Role,
		ProductCategory: request.Product_category,
		ProductName:     request.Product_name,
		OrderText:       request.Order_text,
		OrderComments:   request.Order_comments,
		Fio:             request.Fio,
		Email:           request.Email,
		Phone:           request.Phone,
		CompanyName:     request.Company_name,
		Itn:             request.Itn,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *fastOrderUseCase) LandingOrder(ctx context.Context, request *models.LandingOrderRequest) error {
	_, err := u.fastOrderGRPC.LandingOrder(ctx, &fastOrder_service.LandingOrderRequest{
		ProductCategory: request.Product_category,
		DeliveryAddress: request.Delivery_address,
		DeliveryDate:    request.Delivery_date,
		OrderText:       request.Order_text,
		Email:           request.Email,
		Itn:             request.Itn,
	})
	if err != nil {
		return err
	}

	return nil
}

func NewFastOrderUseCase(grpc fastOrderGRPC) FastOrderUseCase {
	return &fastOrderUseCase{fastOrderGRPC: grpc}
}
