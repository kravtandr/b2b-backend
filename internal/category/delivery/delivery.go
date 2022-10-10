package categoryDelivery

import (
	cnst "b2b/m/pkg/constants"
	"b2b/m/pkg/domain"
	"encoding/json"
	"log"

	cr "b2b/m/internal/category/repository"
	cu "b2b/m/internal/category/usecase"
	ccd "b2b/m/internal/cookie/delivery"

	"github.com/fasthttp/router"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

type CategoryHandler interface {
	// GetCategoryById(ctx *fasthttp.RequestCtx)
	// GetCategoriesInIndustry(ctx *fasthttp.RequestCtx)
	GetAllCategories(ctx *fasthttp.RequestCtx)
	SearchCategories(ctx *fasthttp.RequestCtx)
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
	// r.GET(cnst.CategoryTestURL, categoryHandler.GetCategoryById)
	// r.GET(cnst.IndustryURL, categoryHandler.GetCategoriesInIndustry)
	r.GET(cnst.CategoriesURL, categoryHandler.GetAllCategories)
	r.POST(cnst.CategoriesSearchURL, categoryHandler.SearchCategories)
	return r
}

// func (s *categoryHandler) GetCategoryById(ctx *fasthttp.RequestCtx) {
// 	param, _ := ctx.UserValue("id").(string)
// 	bytes, err := s.CategoryUseCase.GetCategoryById(param)
// 	if err != nil {
// 		ctx.SetStatusCode(fasthttp.StatusNotFound)
// 		log.Printf("error while getting list: %s", err)
// 		ctx.Write([]byte("{}"))
// 		return
// 	}
// 	ctx.SetStatusCode(fasthttp.StatusOK)
// 	ctx.Write(bytes)
// }

// func (s *categoryHandler) GetCategoriesInIndustry(ctx *fasthttp.RequestCtx) {
// 	param, _ := ctx.UserValue("id").(string)
// 	bytes, err := s.CategoryUseCase.GetCategoriesInIndustry(param)
// 	if err != nil {
// 		ctx.SetStatusCode(fasthttp.StatusNotFound)
// 		log.Printf("error while getting list: %s", err)
// 		ctx.Write([]byte("{}"))
// 		return
// 	}
// 	ctx.SetStatusCode(fasthttp.StatusOK)
// 	ctx.Write(bytes)
// }

func (s *categoryHandler) GetAllCategories(ctx *fasthttp.RequestCtx) {
	bytes, err := s.CategoryUseCase.GetAllCategories()
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		log.Printf("error while GetAllCategories: %s", err)
		ctx.Write([]byte("{}"))
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(bytes)
}

func (s *categoryHandler) SearchCategories(ctx *fasthttp.RequestCtx) {
	searchBody := new(domain.Search)
	if err := json.Unmarshal(ctx.PostBody(), &searchBody); err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	bytes, err := s.CategoryUseCase.SearchCategories(searchBody.Name)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		log.Printf("error while SearchCategories: %s", err)
		ctx.Write([]byte("{}"))
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(bytes)
}
