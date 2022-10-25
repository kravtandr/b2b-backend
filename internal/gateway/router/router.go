package router

import (
	fod "b2b/m/internal/gateway/fastOrder/delivery"
	ud "b2b/m/internal/gateway/user/delivery"
	cnst "b2b/m/pkg/constants"
	"b2b/m/pkg/error_adapter"
	fasthttpprom "b2b/m/pkg/fasthttp_prom"
	"b2b/m/pkg/grpc_errors"
	"b2b/m/pkg/middlewares/http"
	auth_service "b2b/m/pkg/services/auth"

	"github.com/fasthttp/router"
	"go.uber.org/zap"
)

type RouterConfig struct {
	AuthGRPC auth_service.AuthServiceClient

	UserDelivery      ud.UserDelivery
	FastOrderDelivery fod.FastOrderDelivery

	Logger *zap.Logger
}

func SetupRouter(cfg RouterConfig) (p fasthttpprom.Router) {
	p = fasthttpprom.NewRouter("gateway")
	p.Use(router.New())

	lgrMw := http.NewLoggingMiddleware(cfg.Logger)
	authMw := http.NewSessionValidatorMiddleware(
		cfg.AuthGRPC,
		error_adapter.NewGrpcToHttpAdapter(grpc_errors.PreparedAuthErrors, grpc_errors.CommonAuthError),
	)

	p.POST(cnst.LoginURL, lgrMw(cfg.UserDelivery.Login))
	p.DELETE(cnst.LogoutURL, lgrMw(authMw(cfg.UserDelivery.Logout)))
	p.GET(cnst.ProfileURL, lgrMw(authMw(cfg.UserDelivery.GetProfile)))
	p.PATCH(cnst.ProfileURL, lgrMw(authMw(cfg.UserDelivery.UpdateProfile)))
	p.POST(cnst.RegisterURL, lgrMw(cfg.UserDelivery.Register))
	p.GET(cnst.UserInfoURL, lgrMw(cfg.UserDelivery.GetUserInfo))

	p.POST(cnst.FastOrderURL, lgrMw(cfg.FastOrderDelivery.FastOrder))

	return
}
