package categoryDelivery

import (
	cnst "b2b/m/pkg/constants"
	"b2b/m/pkg/domain"
	"log"

	cr "b2b/m/internal/category/repository"
	cu "b2b/m/internal/category/usecase"
	ccd "b2b/m/internal/cookie/delivery"

	"github.com/fasthttp/router"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

type CategoryHandler interface {
	GetCategoryById(ctx *fasthttp.RequestCtx)
	GetCategoriesInIndustry(ctx *fasthttp.RequestCtx)
}

type categoryHandler struct {
	CategoryUseCase domain.CategoryUseCase
	CookieHandler   ccd.CookieHandler
}

func NewCategoryHandler(CategoryUseCase domain.CategoryUseCase, CookieHandler ccd.CookieHandler) CategoryHandler {
	return &categoryHandler{
		CategoryUseCase: CategoryUseCase,
		CookieHandler:   CookieHandler,
	}
}

func CreateDelivery(db *pgxpool.Pool) CategoryHandler {
	cookieLayer := ccd.CreateDelivery(db)
	userLayer := NewCategoryHandler(cu.NewCategoryUseCase(cr.NewCategoryStorage(db)), cookieLayer)
	return userLayer
}

func SetUpCategoryRouter(db *pgxpool.Pool, r *router.Router) *router.Router {
	categoryHandler := CreateDelivery(db)
	r.GET(cnst.CategoryTestURL, categoryHandler.GetCategoryById)
	r.GET(cnst.IndustryURL, categoryHandler.GetCategoriesInIndustry)
	return r
}

func (s *categoryHandler) GetCategoryById(ctx *fasthttp.RequestCtx) {
	param, _ := ctx.UserValue("id").(string)
	bytes, err := s.CategoryUseCase.GetCategoryById(param)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		log.Printf("error while getting list: %s", err)
		ctx.Write([]byte("{}"))
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(bytes)
}

func (s *categoryHandler) GetCategoriesInIndustry(ctx *fasthttp.RequestCtx) {
	param, _ := ctx.UserValue("id").(string)
	bytes, err := s.CategoryUseCase.GetCategoriesInIndustry(param)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		log.Printf("error while getting list: %s", err)
		ctx.Write([]byte("{}"))
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(bytes)
}
