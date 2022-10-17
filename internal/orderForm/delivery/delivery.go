package OrderFormDelivery

import (
	cnst "b2b/m/pkg/constants"
	"b2b/m/pkg/domain"
	"encoding/json"
	"log"

	ccd "b2b/m/internal/cookie/delivery"
	ur "b2b/m/internal/orderForm/repository"
	uu "b2b/m/internal/orderForm/usecase"

	"github.com/fasthttp/router"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

type OrderFormHandler interface {
	FastOrder(ctx *fasthttp.RequestCtx)
}

type orderFormHandler struct {
	OrderFormUseCase domain.OrderFormUseCase
	CookieHandler    ccd.CookieHandler
}

func NewOrderFormHandler(OrderFormUseCase domain.OrderFormUseCase, CookieHandler ccd.CookieHandler) OrderFormHandler {
	return &orderFormHandler{
		OrderFormUseCase: OrderFormUseCase,
		CookieHandler:    CookieHandler,
	}
}

func CreateDelivery(db *pgxpool.Pool) OrderFormHandler {
	cookieLayer := ccd.CreateDelivery(db)
	userLayer := NewOrderFormHandler(uu.NewOrderFormUseCase(ur.NewOrderFormStorage(db)), cookieLayer)
	return userLayer
}

func SetUpOrderFormRouter(db *pgxpool.Pool, r *router.Router) *router.Router {
	OrderFormHandler := CreateDelivery(db)
	r.POST(cnst.FastOrderURL, OrderFormHandler.FastOrder)
	return r
}

func (s *orderFormHandler) FastOrder(ctx *fasthttp.RequestCtx) {
	form := new(domain.OrderForm)
	if err := json.Unmarshal(ctx.PostBody(), &form); err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	valid, bytes := s.OrderFormUseCase.Validate(form)
	if !valid {
		log.Printf("error while validating OrderForm")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.Write(bytes)
		return
	}

	code, err := s.OrderFormUseCase.Add(*form)
	if err != nil {
		ctx.SetStatusCode(code)
		ctx.Write(bytes)
		log.Printf("error while add OrderForm", err)
		return
	}
}
