package router

import (
	restchatd "b2b/m/internal/gateway/chat/delivery"
	cd "b2b/m/internal/gateway/company/delivery"
	fod "b2b/m/internal/gateway/fastOrder/delivery"
	pcd "b2b/m/internal/gateway/productsCategories/delivery"
	ud "b2b/m/internal/gateway/user/delivery"
	chatd "b2b/m/internal/gateway_chat/chat/delivery"
	cnst "b2b/m/pkg/constants"
	fasthttpprom "b2b/m/pkg/fasthttp_prom"
	"b2b/m/pkg/middlewares/http"
	auth_service "b2b/m/pkg/services/auth"

	"github.com/fasthttp/router"
	"go.uber.org/zap"
)

type RouterConfig struct {
	AuthGRPC auth_service.AuthServiceClient

	UserDelivery               ud.UserDelivery
	FastOrderDelivery          fod.FastOrderDelivery
	CompanyDelivery            cd.CompanyDelivery
	ProductsCategoriesDelivery pcd.ProductsCategoriesDelivery
	ChatDelivery               chatd.ChatDelivery
	RestChatDelivery           restchatd.ChatDelivery

	Logger *zap.Logger
}

func SetupRouter(cfg RouterConfig) (p fasthttpprom.Router) {
	p = fasthttpprom.NewRouter("gateway_chat")
	p.Use(router.New())

	lgrMw := http.NewLoggingMiddleware(cfg.Logger)
	// authMw := http.NewSessionValidatorMiddleware(
	// 	cfg.AuthGRPC,
	// 	error_adapter.NewGrpcToHttpAdapter(grpc_errors.PreparedAuthErrors, grpc_errors.CommonAuthError),
	// )

	p.GET(cnst.ProductChatURL, lgrMw(cfg.ChatDelivery.WSUpgradeRequest))
	p.GET("/testch", lgrMw(cfg.ChatDelivery.TestCh))

	return
}
