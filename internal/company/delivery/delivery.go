package companyDelivery

import (
	cnst "b2b/m/pkg/constants"
	"b2b/m/pkg/domain"
	"encoding/json"
	"fmt"
	"log"

	cr "b2b/m/internal/company/repository"
	cu "b2b/m/internal/company/usecase"
	ccd "b2b/m/internal/cookie/delivery"

	"github.com/fasthttp/router"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	"gopkg.in/webdeskltd/dadata.v2"
)

type CompanyHandler interface {
	Add(ctx *fasthttp.RequestCtx)
	GetCompanyEmployees(ctx *fasthttp.RequestCtx)
	GetCompanyById(ctx *fasthttp.RequestCtx)
	GetCompanyFullInfo(ctx *fasthttp.RequestCtx)
	GetCompanyByItnDaData(ctx *fasthttp.RequestCtx)
	// SearchCompanies(ctx *fasthttp.RequestCtx)
}

type companyHandler struct {
	CompanyUseCase domain.CompanyUseCase
	CookieHandler  ccd.CookieHandler
}

func NewCompanyHandler(CompanyUseCase domain.CompanyUseCase, CookieHandler ccd.CookieHandler) CompanyHandler {
	return &companyHandler{
		CompanyUseCase: CompanyUseCase,
		CookieHandler:  CookieHandler,
	}
}

func CreateDelivery(db *pgxpool.Pool, daData *dadata.DaData) CompanyHandler {
	cookieLayer := ccd.CreateDelivery(db)
	userLayer := NewCompanyHandler(cu.NewCompanyUseCase(cr.NewCompanyStorage(db, daData)), cookieLayer)
	return userLayer
}

func SetUpCompanyRouter(db *pgxpool.Pool, daData *dadata.DaData, r *router.Router) *router.Router {
	companyHandler := CreateDelivery(db, daData)
	r.POST(cnst.RegisterCompanyURL, companyHandler.Add)
	r.GET(cnst.CompanyEmployeesURL, companyHandler.GetCompanyEmployees)
	r.GET(cnst.CompanyURL, companyHandler.GetCompanyById)
	r.GET(cnst.CompanyInfoURL, companyHandler.GetCompanyFullInfo)
	r.POST(cnst.CompanyFindByItnURL, companyHandler.GetCompanyByItnDaData)
	return r
}

func (s *companyHandler) GetCompanyById(ctx *fasthttp.RequestCtx) {
	param, _ := ctx.UserValue("id").(string)
	bytes, err := s.CompanyUseCase.GetCompanyById(param)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		log.Printf(": %s", err)
		ctx.Write([]byte("{}"))
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(bytes)
}

func (s *companyHandler) GetCompanyEmployees(ctx *fasthttp.RequestCtx) {
	param, _ := ctx.UserValue("id").(string)
	bytes, err := s.CompanyUseCase.GetCompanyEmployees(param)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		log.Printf(": %s", err)
		ctx.Write(bytes)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(bytes)
}

func (s *companyHandler) GetCompanyFullInfo(ctx *fasthttp.RequestCtx) {
	param, _ := ctx.UserValue("id").(string)
	bytes, err := s.CompanyUseCase.GetCompanyFullInfo(param)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		log.Printf(": %s", err)
		ctx.Write(bytes)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(bytes)
}

func (s *companyHandler) Add(ctx *fasthttp.RequestCtx) {
	company := new(domain.Company)

	if err := json.Unmarshal(ctx.PostBody(), &company); err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	err := s.CompanyUseCase.Add(company)

	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		log.Printf("error while registering user in", err)
		return
	}

	newCompany, _ := s.CompanyUseCase.GetByEmail(company.Email)
	bytes, err := s.CompanyUseCase.GetCompanyById(fmt.Sprint(newCompany.Id))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		log.Printf(": %s", err)
		ctx.Write([]byte("{}"))
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(bytes)

}

func (s *companyHandler) GetCompanyByItnDaData(ctx *fasthttp.RequestCtx) {
	param, _ := ctx.UserValue("itn").(string)
	bytes, err := s.CompanyUseCase.GetCompanyByItnDaData(param)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		log.Printf(": %s", err)
		ctx.Write(bytes)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(bytes)
}

// func (s *companyHandler) SearchCompanies(ctx *fasthttp.RequestCtx) {
// 	//var request = &domain.CompanySearch{}
// 	var request domain.CompanySearch
// 	if err := json.Unmarshal(ctx.Request.Body(), &request); err != nil {
// 		ctx.SetStatusCode(fasthttp.StatusBadRequest)
// 		ctx.SetBody([]byte{})
// 		return
// 	}

// 	bytes, err := s.CompanyUseCase.SearchCompanies(request)
// 	if err != nil {
// 		ctx.SetStatusCode(fasthttp.StatusBadRequest)
// 		ctx.SetBody([]byte{})
// 		return
// 	}

// 	ctx.SetStatusCode(fasthttp.StatusOK)
// 	ctx.Write(bytes)
// }

// func (s *companyHandler) GetCompaniesByCategoryId(ctx *fasthttp.RequestCtx) {
// 	param, _ := ctx.UserValue("id").(string)
// 	bytes, err := s.CompanyUseCase.GetCompaniesByCategoryId(param)
// 	if err != nil {
// 		ctx.SetStatusCode(fasthttp.StatusNotFound)
// 		log.Printf("companyHandler error while getting list: %s", err)
// 		ctx.Write([]byte("{}"))
// 		return
// 	}
// 	ctx.SetStatusCode(fasthttp.StatusOK)
// 	ctx.Write(bytes)
// }
