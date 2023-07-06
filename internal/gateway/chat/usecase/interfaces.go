package usecase

import (
	"context"

	chat_service "b2b/m/pkg/services/chat"

	"google.golang.org/grpc"
)

type chatGRPC interface {
	LoginUser(ctx context.Context, in *chat_service.LoginRequest, opts ...grpc.CallOption) (*chat_service.LoginResponse, error)
}
