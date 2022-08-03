package industryDelivery

import (
	cnst "b2b/m/pkg/constants"
	"b2b/m/pkg/domain"
	"log"

	ccd "b2b/m/internal/cookie/delivery"
	ir "b2b/m/internal/industry/repository"
	iu "b2b/m/internal/industry/usecase"

	"github.com/fasthttp/router"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

type IndustryHandler interface {
	GetAllIndustries(ctx *fasthttp.RequestCtx)
}

type industryHandler struct {
	IndustryUseCase domain.IndustryUseCase
	CookieHandler   ccd.CookieHandler
}

func NewIndustryHandler(IndustryUseCase domain.IndustryUseCase, CookieHandler ccd.CookieHandler) IndustryHandler {
	return &industryHandler{
		IndustryUseCase: IndustryUseCase,
		CookieHandler:   CookieHandler,
	}
}

func CreateDelivery(db *pgxpool.Pool) IndustryHandler {
	cookieLayer := ccd.CreateDelivery(db)
	userLayer := NewIndustryHandler(iu.NewIndustryUseCase(ir.NewIndustryStorage(db)), cookieLayer)
	return userLayer
}

func SetUpIndustryRouter(db *pgxpool.Pool, r *router.Router) *router.Router {
	industryHandler := CreateDelivery(db)
	r.GET(cnst.IndustriesURL, industryHandler.GetAllIndustries)
	return r
}

func (s *industryHandler) GetAllIndustries(ctx *fasthttp.RequestCtx) {
	log.Printf("HERE")
	bytes, err := s.IndustryUseCase.GetAllIndustries()
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		log.Printf("error while getting list: %s", err)
		ctx.Write([]byte("{}"))
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(bytes)
}
