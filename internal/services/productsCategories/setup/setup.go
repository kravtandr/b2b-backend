package setup

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	zap_middleware "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"

	"b2b/m/internal/services/productsCategories/config"
	"b2b/m/internal/services/productsCategories/delivery"
	"b2b/m/internal/services/productsCategories/repository"
	productsCategories_usecase "b2b/m/internal/services/productsCategories/usecase"
	"b2b/m/pkg/error_adapter"
	"b2b/m/pkg/grpc_errors"
	"b2b/m/pkg/helpers"
	productsCategories_service "b2b/m/pkg/services/productsCategories"
)

func SetupServer(cfg config.Config) (server *grpc.Server, cancel func(), err error) {
	conn, err := helpers.CreatePGXPool(cfg.Ctx, cfg.DBUrl, cfg.Logger)
	if err != nil {
		return server, cancel, err
	}

	productsCategoriesRepo := repository.NewLoggingMiddleware(
		cfg.Logger.Sugar(),
		repository.NewProductsCategoriesRepository(
			repository.NewQueryFactory(), conn,
		),
	)
	productsCategoriesUseCase := productsCategories_usecase.NewProductsCategoriesUseCase(productsCategoriesRepo)
	productsCategoriesDelivery := delivery.NewProductsCategoriesDelivery(productsCategoriesUseCase, error_adapter.NewErrorAdapter(grpc_errors.PreparedFastOrderServiceErrorMap))

	server = grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			zap_middleware.UnaryServerInterceptor(cfg.Logger),
			grpc_validator.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
		),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
	)
	productsCategories_service.RegisterProductsCategoriesServiceServer(server, productsCategoriesDelivery)
	grpc_prometheus.Register(server)

	cancel = func() {
		conn.Close()
	}
	return server, cancel, nil
}
