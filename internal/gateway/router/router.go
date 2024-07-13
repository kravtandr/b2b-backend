package router

import (
	chatd "b2b/m/internal/gateway/chat/delivery"
	cd "b2b/m/internal/gateway/company/delivery"
	fod "b2b/m/internal/gateway/fastOrder/delivery"
	pcd "b2b/m/internal/gateway/productsCategories/delivery"
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

	UserDelivery               ud.UserDelivery
	FastOrderDelivery          fod.FastOrderDelivery
	CompanyDelivery            cd.CompanyDelivery
	ProductsCategoriesDelivery pcd.ProductsCategoriesDelivery
	ChatDelivery               chatd.ChatDelivery

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

	// Auth
	p.POST(cnst.LoginURL, lgrMw(cfg.UserDelivery.Login))
	p.DELETE(cnst.LogoutURL, lgrMw(authMw(cfg.UserDelivery.Logout)))
	p.GET(cnst.ProfileURL, lgrMw(authMw(cfg.UserDelivery.GetProfile)))
	p.PATCH(cnst.ProfileURL, lgrMw(authMw(cfg.UserDelivery.UpdateProfile)))
	p.POST(cnst.RegisterURL, lgrMw(cfg.UserDelivery.Register))
	p.POST(cnst.UserCheckEmailURL, lgrMw(cfg.UserDelivery.CheckEmail))
	p.GET(cnst.UserInfoByCookieURL, lgrMw(authMw(cfg.UserDelivery.GetUserByCookie)))
	p.POST(cnst.FastRegisterURL, lgrMw(cfg.UserDelivery.FastRegister))
	//yookassa
	p.POST(cnst.CreatePaymentAddBalanceURL, lgrMw(authMw(cfg.UserDelivery.CreatePayemntAddBalance)))
	p.GET(cnst.CheckPaymentURL, lgrMw(authMw(cfg.UserDelivery.CheckPayment)))
	p.POST(cnst.HandlePaidPaymentsURL, lgrMw(authMw(cfg.UserDelivery.HandlePaidPayments)))
	p.GET(cnst.GetUsersPaymentsURL, lgrMw(authMw(cfg.UserDelivery.GetUsersPayments)))

	// Unlinked Froms
	p.POST(cnst.FastOrderURL, lgrMw(cfg.FastOrderDelivery.FastOrder))
	p.POST(cnst.LandingOrderURL, lgrMw(cfg.FastOrderDelivery.LandingOrder))

	// Company
	p.GET(cnst.CompanyURL, lgrMw(cfg.CompanyDelivery.GetCompanyById))
	p.POST(cnst.CompanyByInnFromDaDataURL, lgrMw(cfg.CompanyDelivery.GetCompanyByItnFromDaData))

	// Categories
	p.GET(cnst.CategoryByIdURL, lgrMw(cfg.ProductsCategoriesDelivery.GetCategoryById))
	p.POST(cnst.SearchCategoryURL, lgrMw(cfg.ProductsCategoriesDelivery.SearchCategories))

	// Products
	p.GET(cnst.ProductURL, lgrMw(cfg.ProductsCategoriesDelivery.GetProductById))
	p.GET(cnst.ProductsListURL, lgrMw(cfg.ProductsCategoriesDelivery.GetProductsList))
	p.POST(cnst.ProductsListByFiltersURL, lgrMw(cfg.ProductsCategoriesDelivery.GetProductsListByFilters))
	p.POST(cnst.SearchProductsURL, lgrMw(cfg.ProductsCategoriesDelivery.SearchProducts))
	p.POST(cnst.AddProductURL, lgrMw(authMw(cfg.ProductsCategoriesDelivery.AddProduct)))
	p.GET(cnst.CompanyProductsListURL, lgrMw(cfg.ProductsCategoriesDelivery.GetCompanyProducts))

	// Chat
	p.GET(cnst.CheckIfUniqChat, lgrMw(authMw(cfg.ChatDelivery.CheckIfUniqChat)))
	p.POST(cnst.InitChat, lgrMw(authMw(cfg.ChatDelivery.InitChat)))
	p.PATCH(cnst.ChatChangeStatus, lgrMw(authMw(cfg.ChatDelivery.UpdateChatStatus)))
	p.DELETE(cnst.DeleteChat, lgrMw(authMw(cfg.ChatDelivery.DeleteChat)))
	p.GET(cnst.AllChats, lgrMw(authMw(cfg.ChatDelivery.GetAllChatsAndLastMsg)))
	p.GET(cnst.AllMsgsFromChat, lgrMw(authMw(cfg.ChatDelivery.GetMsgsFromChat)))
	p.GET("/testgw", lgrMw(cfg.ChatDelivery.TestGw))
	//p.GET(cnst.ChatHealthCheck, lgrMw(cfg.ChatDelivery.ChatHealthCheck))

	return
}
