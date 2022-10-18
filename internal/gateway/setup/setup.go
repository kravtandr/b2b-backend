package setup

import (
	"b2b/m/internal/gateway/config"
	"b2b/m/internal/gateway/router"
	ud "b2b/m/internal/gateway/user/delivery"
	uu "b2b/m/internal/gateway/user/usecase"
	"b2b/m/pkg/error_adapter"
	fasthttpprom "b2b/m/pkg/fasthttp_prom"
	"b2b/m/pkg/grpc_errors"
	"b2b/m/pkg/helpers"
	auth_service "b2b/m/pkg/services/auth"

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

	p = router.SetupRouter(router.RouterConfig{
		AuthGRPC:     userGRPC,
		UserDelivery: userDelivery,
		Logger:       cfg.Logger,
	})

	stopFunc = func() {
		conn.Close()
		pgxConn.Close()
	}

	return
}
