package setup

import (
	restchatd "b2b/m/internal/gateway/chat/delivery"
	restchatu "b2b/m/internal/gateway/chat/usecase"
	cd "b2b/m/internal/gateway/company/delivery"
	company_usecase "b2b/m/internal/gateway/company/usecase"
	fod "b2b/m/internal/gateway/fastOrder/delivery"
	fastOrder_usecase "b2b/m/internal/gateway/fastOrder/usecase"
	pcd "b2b/m/internal/gateway/productsCategories/delivery"
	productsCategories_usecase "b2b/m/internal/gateway/productsCategories/usecase"
	ud "b2b/m/internal/gateway/user/delivery"
	uu "b2b/m/internal/gateway/user/usecase"
	chatd "b2b/m/internal/gateway_chat/chat/delivery"
	chatu "b2b/m/internal/gateway_chat/chat/usecase"
	"b2b/m/internal/gateway_chat/config"
	"b2b/m/internal/gateway_chat/router"
	"b2b/m/pkg/error_adapter"
	fasthttpprom "b2b/m/pkg/fasthttp_prom"
	"b2b/m/pkg/grpc_errors"
	"b2b/m/pkg/helpers"
	auth_service "b2b/m/pkg/services/auth"
	chat_service "b2b/m/pkg/services/chat"
	company_service "b2b/m/pkg/services/company"
	fastOrder_service "b2b/m/pkg/services/fastOrder"
	productsCategories_service "b2b/m/pkg/services/productsCategories"
	"time"

	yookassa "github.com/rvinnie/yookassa-sdk-go/yookassa"
	"gopkg.in/webdeskltd/dadata.v2"

	"google.golang.org/grpc"
)

func Setup(cfg config.Config) (p fasthttpprom.Router, stopFunc func(), err error) {
	pgxConn, err := helpers.CreatePGXPool(cfg.Ctx, cfg.DBUrl, cfg.Logger)
	if err != nil {
		return p, stopFunc, err
	}

	companyConn, err := grpc.Dial(cfg.CompanyServiceEndpoint, grpc.WithInsecure(), grpc.WithBlock())
	daData := dadata.NewDaData("42e877cc6e66e3cc70c47a2f42966120cfcea751", "984e0c50d52dd2611b98609eaa7c82268e46297e")
	if err != nil {
		return p, stopFunc, err
	}
	CompanyGRPC := company_service.NewCompanyServiceClient(companyConn)
	companyUseCase := company_usecase.NewCompanyUseCase(CompanyGRPC, daData)
	companyDelivery := cd.NewCompanyDelivery(
		error_adapter.NewGrpcToHttpAdapter(
			grpc_errors.UserGatewayError, grpc_errors.Fail,
		),
		companyUseCase,
	)

	conn, err := grpc.Dial(cfg.AuthServiceEndpoint, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return p, stopFunc, err
	}
	var UKASSA_SECRET_KEY = "test_on5XfensBJfl8nY63uzUp3CBTW5YE0w6SVhr4VxAdH4"
	var ACCIUNT_ID = "415910"
	yooclient := yookassa.NewClient(ACCIUNT_ID, UKASSA_SECRET_KEY)
	// Создаем обработчик платежей
	paymentHandler := yookassa.NewPaymentHandler(yooclient)
	userGRPC := auth_service.NewAuthServiceClient(conn)
	userUsecase := uu.NewUserUsecase(userGRPC, CompanyGRPC, paymentHandler)
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

	productsCategoriesConn, err := grpc.Dial(cfg.ProductsCategoriesServiceEndpoint, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return p, stopFunc, err
	}
	ProductsCategoriesGRPC := productsCategories_service.NewProductsCategoriesServiceClient(productsCategoriesConn)
	productsCategoriesUseCase := productsCategories_usecase.NewProductsCategoriesUseCase(ProductsCategoriesGRPC, userUsecase, companyUseCase)
	productsCategoriesDelivery := pcd.NewProductsCategoriesDelivery(
		error_adapter.NewGrpcToHttpAdapter(
			grpc_errors.UserGatewayError, grpc_errors.CommonError,
		),
		productsCategoriesUseCase,
	)
	RestChatConn, err := grpc.Dial(cfg.ChatServiceEndpoint, grpc.WithTimeout(5*time.Second), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return p, stopFunc, err
	}
	RestchatGRPC := chat_service.NewChatServiceClient(RestChatConn)
	RestchatUsecase := restchatu.NewChatUsecase(RestchatGRPC, CompanyGRPC, ProductsCategoriesGRPC, userGRPC, userUsecase)
	RestchatDelivery := restchatd.NewChatDelivery(
		error_adapter.NewGrpcToHttpAdapter(
			grpc_errors.UserGatewayError, grpc_errors.CommonError,
		),
		RestchatUsecase,
	)

	ChatConn, err := grpc.Dial(cfg.ChatServiceEndpoint, grpc.WithTimeout(5*time.Second), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return p, stopFunc, err
	}
	chatGRPC := chat_service.NewChatServiceClient(ChatConn)
	chatUsecase := chatu.NewChatUsecase(chatGRPC, CompanyGRPC, ProductsCategoriesGRPC)
	chatDelivery := chatd.NewChatDelivery(
		error_adapter.NewGrpcToHttpAdapter(
			grpc_errors.UserGatewayError, grpc_errors.CommonError,
		),
		chatUsecase, RestchatUsecase, userUsecase,
	)

	p = router.SetupRouter(router.RouterConfig{
		AuthGRPC:                   userGRPC,
		UserDelivery:               userDelivery,
		CompanyDelivery:            companyDelivery,
		FastOrderDelivery:          fastOrderDelivery,
		ProductsCategoriesDelivery: productsCategoriesDelivery,
		ChatDelivery:               chatDelivery,
		RestChatDelivery:           RestchatDelivery,
		Logger:                     cfg.Logger,
	})

	stopFunc = func() {
		conn.Close()
		pgxConn.Close()
	}

	return
}
