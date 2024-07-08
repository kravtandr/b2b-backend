package setup

import (
	zap_middleware "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"

	"b2b/m/internal/services/company/config"
	"b2b/m/internal/services/company/delivery"
	"b2b/m/internal/services/company/repository"
	company_usecase "b2b/m/internal/services/company/usecase"
	"b2b/m/pkg/error_adapter"
	"b2b/m/pkg/grpc_errors"
	"b2b/m/pkg/helpers"
	company_service "b2b/m/pkg/services/company"
)

func SetupServer(cfg config.Config) (server *grpc.Server, cancel func(), err error) {
	conn, err := helpers.CreatePGXPool(cfg.Ctx, cfg.DBUrl, cfg.Logger)
	if err != nil {
		return server, cancel, err
	}

	companyRepo := repository.NewLoggingMiddleware(
		cfg.Logger.Sugar(),
		repository.NewCompanyRepository(
			repository.NewQueryFactory(), conn,
		),
	)
	companyUseCase := company_usecase.NewCompanyUseCase(companyRepo)
	companyDelivery := delivery.NewCompanyDelivery(companyUseCase, error_adapter.NewErrorAdapter(grpc_errors.PreparedCompanyServiceErrorMap))

	server = grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			zap_middleware.UnaryServerInterceptor(cfg.Logger),
			grpc_validator.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
		),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
	)
	company_service.RegisterCompanyServiceServer(server, companyDelivery)
	grpc_prometheus.Register(server)

	cancel = func() {
		conn.Close()
	}
	return server, cancel, nil
}
