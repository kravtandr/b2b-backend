package delivery

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"b2b/m/internal/services/fastOrder/models"
	"b2b/m/internal/services/fastOrder/usecase"
	"b2b/m/pkg/error_adapter"
	fastOrder_service "b2b/m/pkg/services/fastOrder"
)

type fastOrderDelivery struct {
	fastOrderUseCase usecase.FastOrderUseCase
	errorAdapter     error_adapter.ErrorAdapter
	fastOrder_service.UnimplementedFastOrderServiceServer
}

func (a *fastOrderDelivery) FastOrder(ctx context.Context, request *fastOrder_service.FastOrderRequest) (*empty.Empty, error) {
	err := a.fastOrderUseCase.FastOrder(ctx, &models.OrderForm{
		Role:             request.Role,
		Product_category: request.ProductCategory,
		Order_text:       request.OrderText,
		Order_comments:   request.OrderComments,
		Fio:              request.Fio,
		Email:            request.Email,
		Phone:            request.Phone,
		Company_name:     request.CompanyName,
		Itn:              request.Itn,
	})
	if err != nil {
		return &empty.Empty{}, a.errorAdapter.AdaptError(err)
	}

	return &empty.Empty{}, nil
}

func (a *fastOrderDelivery) LandingOrder(ctx context.Context, request *fastOrder_service.LandingOrderRequest) (*empty.Empty, error) {
	err := a.fastOrderUseCase.LandingOrder(ctx, &models.LandingForm{
		Product_category: request.ProductCategory,
		Delivery_address: request.DeliveryAddress,
		Delivery_date:    request.DeliveryDate,
		Order_text:       request.OrderText,
		Email:            request.Email,
		Itn:              request.Itn,
	})
	if err != nil {
		return &empty.Empty{}, a.errorAdapter.AdaptError(err)
	}

	return &empty.Empty{}, nil
}

func NewFastOrderDelivery(
	fastOrderUseCase usecase.FastOrderUseCase,
	errorAdapter error_adapter.ErrorAdapter,
) fastOrder_service.FastOrderServiceServer {
	return &fastOrderDelivery{
		fastOrderUseCase: fastOrderUseCase,
		errorAdapter:     errorAdapter,
	}
}
