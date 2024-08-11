package setup

import (
	zap_middleware "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"google.golang.org/grpc"

	"b2b/m/internal/services/auth/config"
	"b2b/m/internal/services/auth/delivery"
	"b2b/m/internal/services/auth/repository"
	auth_usecase "b2b/m/internal/services/auth/usecase"
	"b2b/m/pkg/error_adapter"
	"b2b/m/pkg/grpc_errors"
	"b2b/m/pkg/hasher"
	"b2b/m/pkg/helpers"
	auth_service "b2b/m/pkg/services/auth"
)

func SetupServer(cfg config.Config) (server *grpc.Server, cancel func(), err error) {
	conn, err := helpers.CreatePGXPool(cfg.Ctx, cfg.DBUrl, cfg.Logger)
	if err != nil {
		return server, cancel, err
	}

	authRepo := repository.NewLoggingMiddleware(
		cfg.Logger.Sugar(),
		repository.NewAuthRepository(
			repository.NewQueryFactory(), conn,
		),
	)
	authUsecase := auth_usecase.NewAuthUseCase(hasher.NewHasher(cfg.PrefixLen), authRepo)
	authDelivery := delivery.NewAuthDelivery(authUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	server = grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			zap_middleware.UnaryServerInterceptor(cfg.Logger),
			grpc_validator.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
		),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
	)
	auth_service.RegisterAuthServiceServer(server, authDelivery)
	grpc_prometheus.Register(server)

	cancel = func() {
		conn.Close()
	}
	return server, cancel, nil
}
