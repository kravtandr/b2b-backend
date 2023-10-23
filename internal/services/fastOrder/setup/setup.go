package setup

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	zap_middleware "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"

	"b2b/m/internal/services/fastOrder/config"
	"b2b/m/internal/services/fastOrder/delivery"
	"b2b/m/internal/services/fastOrder/repository"
	fastOrder_usecase "b2b/m/internal/services/fastOrder/usecase"
	"b2b/m/pkg/error_adapter"
	"b2b/m/pkg/grpc_errors"
	"b2b/m/pkg/helpers"
	fastOrder_service "b2b/m/pkg/services/fastOrder"
)

func SetupServer(cfg config.Config) (server *grpc.Server, cancel func(), err error) {
	conn, err := helpers.CreatePGXPool(cfg.Ctx, cfg.DBUrl, cfg.Logger)
	if err != nil {
		return server, cancel, err
	}

	fastOrderRepo := repository.NewLoggingMiddleware(
		cfg.Logger.Sugar(),
		repository.NewFastOrderRepository(
			repository.NewQueryFactory(), conn,
		),
	)
	fastOrderUseCase := fastOrder_usecase.NewFastOrderUseCase(fastOrderRepo)
	fastOrderDelivery := delivery.NewFastOrderDelivery(fastOrderUseCase, error_adapter.NewErrorAdapter(grpc_errors.PreparedFastOrderServiceErrorMap))

	server = grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			zap_middleware.UnaryServerInterceptor(cfg.Logger),
			grpc_validator.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
		),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
	)
	fastOrder_service.RegisterFastOrderServiceServer(server, fastOrderDelivery)
	grpc_prometheus.Register(server)

	cancel = func() {
		conn.Close()
	}
	return server, cancel, nil
}
