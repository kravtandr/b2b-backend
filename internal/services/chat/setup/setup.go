package setup

import (
	"b2b/m/internal/services/chat/usecase"
	chat_service "b2b/m/pkg/services/chat"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	zap_middleware "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"

	"b2b/m/internal/services/chat/config"
	"b2b/m/internal/services/chat/delivery"
	"b2b/m/internal/services/chat/repository"
	"b2b/m/pkg/error_adapter"
	"b2b/m/pkg/grpc_errors"
	"b2b/m/pkg/hasher"
	"b2b/m/pkg/helpers"
)

func SetupServer(cfg config.Config) (server *grpc.Server, cancel func(), err error) {
	conn, err := helpers.CreatePGXPool(cfg.Ctx, cfg.DBUrl, cfg.Logger)
	if err != nil {
		return server, cancel, err
	}

	chatRepo := repository.NewLoggingMiddleware(
		cfg.Logger.Sugar(),
		repository.NewChatRepository(
			repository.NewQueryFactory(), conn,
		),
	)
	chatUsecase := usecase.NewChatUseCase(hasher.NewHasher(3), chatRepo)
	//authUsecase := auth_usecase.NewCompanyUseCase(authRepo)
	//PreparedChatServiceErrorMap
	chatDelivery := delivery.NewChatDelivery(chatUsecase, error_adapter.NewErrorAdapter(grpc_errors.PreparedAuthServiceErrorMap))

	server = grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			zap_middleware.UnaryServerInterceptor(cfg.Logger),
			grpc_validator.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
		),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
	)
	chat_service.RegisterChatServiceServer(server, chatDelivery)
	grpc_prometheus.Register(server)

	cancel = func() {
		conn.Close()
	}
	return server, cancel, nil
}
