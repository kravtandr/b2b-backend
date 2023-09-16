package usecase

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	fastOrder_service "b2b/m/pkg/services/fastOrder"

	"google.golang.org/grpc"
)

type fastOrderGRPC interface {
	FastOrder(ctx context.Context, in *fastOrder_service.FastOrderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	LandingOrder(ctx context.Context, in *fastOrder_service.LandingOrderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}
