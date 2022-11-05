package setup

import (
	cd "b2b/m/internal/gateway/company/delivery"
	company_usecase "b2b/m/internal/gateway/company/usecase"
	"b2b/m/internal/gateway/config"
	fod "b2b/m/internal/gateway/fastOrder/delivery"
	fastOrder_usecase "b2b/m/internal/gateway/fastOrder/usecase"
	"b2b/m/internal/gateway/router"
	ud "b2b/m/internal/gateway/user/delivery"
	uu "b2b/m/internal/gateway/user/usecase"
	"b2b/m/pkg/error_adapter"
	fasthttpprom "b2b/m/pkg/fasthttp_prom"
	"b2b/m/pkg/grpc_errors"
	"b2b/m/pkg/helpers"
	auth_service "b2b/m/pkg/services/auth"
	company_service "b2b/m/pkg/services/company"
	fastOrder_service "b2b/m/pkg/services/fastOrder"

	"google.golang.org/grpc"
)

func Setup(cfg config.Config) (p fasthttpprom.Router, stopFunc func(), err error) {
	pgxConn, err := helpers.CreatePGXPool(cfg.Ctx, cfg.DBUrl, cfg.Logger)
	if err != nil {
		return p, stopFunc, err
	}

	conn, err := grpc.Dial(cfg.AuthServiceEndpoint, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return p, stopFunc, err
	}
	userGRPC := auth_service.NewAuthServiceClient(conn)
	userUsecase := uu.NewUserUsecase(userGRPC)
	userDelivery := ud.NewUserDelivery(
		error_adapter.NewGrpcToHttpAdapter(
			grpc_errors.UserGatewayError, grpc_errors.CommonError,
		),
		userUsecase,
	)

	fastOrderConn, err := grpc.Dial(cfg.FastOrderServiceEndpoint, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return p, stopFunc, err
	}
	fastOrderGRPC := fastOrder_service.NewFastOrderServiceClient(fastOrderConn)
	fastOrderUseCase := fastOrder_usecase.NewFastOrderUseCase(fastOrderGRPC)
	fastOrderDelivery := fod.NewFastOrderDelivery(
		error_adapter.NewGrpcToHttpAdapter(
			grpc_errors.UserGatewayError, grpc_errors.CommonError,
		),
		fastOrderUseCase,
	)

	companyConn, err := grpc.Dial(cfg.CompanyServiceEndpoint, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return p, stopFunc, err
	}
	companyGRPC := company_service.NewCompanyServiceClient(companyConn)
	companyUseCase := company_usecase.NewCompanyUseCase(companyGRPC)
	companyDelivery := cd.NewCompanyDelivery(
		error_adapter.NewGrpcToHttpAdapter(
			grpc_errors.UserGatewayError, grpc_errors.Fail,
		),
		companyUseCase,
	)

	p = router.SetupRouter(router.RouterConfig{
		AuthGRPC:          userGRPC,
		UserDelivery:      userDelivery,
		CompanyDelivery:   companyDelivery,
		FastOrderDelivery: fastOrderDelivery,
		Logger:            cfg.Logger,
	})

	stopFunc = func() {
		conn.Close()
		pgxConn.Close()
	}

	return
}