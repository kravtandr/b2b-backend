package userDelivery

import (
	cnst "b2b/m/pkg/constants"
	"b2b/m/pkg/domain"
	"encoding/json"
	"fmt"
	"log"

	cr "b2b/m/internal/company/repository"
	cu "b2b/m/internal/company/usecase"
	ccd "b2b/m/internal/cookie/delivery"
	ur "b2b/m/internal/user/repository"
	uu "b2b/m/internal/user/usecase"

	"github.com/fasthttp/router"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	"gopkg.in/webdeskltd/dadata.v2"
)

type UserHandler interface {
	Login(ctx *fasthttp.RequestCtx)
	Registration(ctx *fasthttp.RequestCtx)
	FastRegistration(ctx *fasthttp.RequestCtx)
	Logout(ctx *fasthttp.RequestCtx)
	GetPublicUserByEmail(ctx *fasthttp.RequestCtx)
	GetPublicUserById(ctx *fasthttp.RequestCtx)
	// GetCompanyById(ctx *fasthttp.RequestCtx)
	// GetCompaniesByCategoryId(ctx *fasthttp.RequestCtx)
	// SearchCompanies(ctx *fasthttp.RequestCtx)
}

type userHandler struct {
	UserUseCase    domain.UserUseCase
	CompanyUseCase domain.CompanyUseCase
	CookieHandler  ccd.CookieHandler
}

func NewUserHandler(UserUseCase domain.UserUseCase, CookieHandler ccd.CookieHandler, CompanyUseCase domain.CompanyUseCase) UserHandler {
	return &userHandler{
		UserUseCase:    UserUseCase,
		CompanyUseCase: CompanyUseCase,
		CookieHandler:  CookieHandler,
	}
}

func CreateDelivery(db *pgxpool.Pool, daData *dadata.DaData) UserHandler {
	cookieLayer := ccd.CreateDelivery(db)
	companyLayer := cu.NewCompanyUseCase(cr.NewCompanyStorage(db, daData))
	userLayer := NewUserHandler(uu.NewUserUseCase(ur.NewUserStorage(db)), cookieLayer, companyLayer)
	return userLayer
}

func SetUpUserRouter(db *pgxpool.Pool, daData *dadata.DaData, r *router.Router) *router.Router {
	userHandler := CreateDelivery(db, daData)
	r.POST(cnst.LoginURL, userHandler.Login)
	r.POST(cnst.RegisterURL, userHandler.Registration)
	r.POST(cnst.CheckEmailURL, userHandler.GetPublicUserByEmail)
	r.POST(cnst.FastRegistrationURL, userHandler.FastRegistration)
	r.GET(cnst.UserURL, userHandler.GetPublicUserById)
	// r.GET(cnst.CompanyURL, companyHandler.GetCompanyById)
	// r.GET(cnst.CategoryURL, companyHandler.GetCompaniesByCategoryId)
	// r.POST(cnst.CompanySearchURL, companyHandler.SearchCompanies)
	return r
}

func (s *userHandler) Login(ctx *fasthttp.RequestCtx) {
	user := new(domain.User)
	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	valid := s.UserUseCase.Validate(user)
	if !valid {
		log.Printf("error while validating user")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	code, err := s.UserUseCase.Login(user)
	ctx.SetStatusCode(code)
	if err != nil {
		log.Printf("error while logging user in")
		return
	}

	с := fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(user.Email)))
	foundUser, _ := s.UserUseCase.GetByEmail(user.Email)
	s.CookieHandler.SetCookieAndToken(ctx, с, foundUser.Id)
}

func (s *userHandler) Registration(ctx *fasthttp.RequestCtx) {
	user := new(domain.User)
	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	code, err := s.UserUseCase.Registration(user)
	ctx.SetStatusCode(code)
	if err != nil {
		log.Printf("error while registering user in", err)
		return
	}

	с := fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(user.Email)))
	newUser, _ := s.UserUseCase.GetByEmail(user.Email)
	s.CookieHandler.SetCookieAndToken(ctx, с, newUser.Id)
}

func (s *userHandler) FastRegistration(ctx *fasthttp.RequestCtx) {
	form := new(domain.FastRegistrationForm)
	if err := json.Unmarshal(ctx.PostBody(), &form); err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	user := new(domain.User)
	user.Email = form.Email
	user.Password = form.Password
	user.Name = form.OwnerName
	user.Surname = form.Surname
	user.Patronymic = form.Patronymic
	user.Country = form.Country

	code, err := s.UserUseCase.Registration(user)
	ctx.SetStatusCode(code)
	if err != nil {
		log.Printf("error while registering user", err)
		return
	}
	company := new(domain.Company)
	company.Name = form.Name
	company.LegalName = form.LegalName
	company.Itn = form.Itn
	company.Email = form.Email
	newUser, err := s.UserUseCase.GetByEmail(user.Email)
	if err != nil {
		log.Printf("error while UserUseCase GetByEmail", err)
		return
	}
	company.OwnerId = newUser.Id

	err = s.CompanyUseCase.AddBaseCompany(company, form.Post)
	if err != nil {
		log.Printf("error while AddBaseCompany", err)
		return
	}
}

func (s *userHandler) Logout(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	s.CookieHandler.DeleteCookie(ctx, string(ctx.Request.Header.Cookie(cnst.CookieName)))
}

func (s *userHandler) GetPublicUserByEmail(ctx *fasthttp.RequestCtx) {
	request := new(domain.UserEmail)
	if err := json.Unmarshal(ctx.Request.Body(), &request); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBody([]byte{})
		return
	}
	bytes, err := s.UserUseCase.GetPublicUserByEmail(request.Email)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBody([]byte{})
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(bytes)
}

func (s *userHandler) GetPublicUserById(ctx *fasthttp.RequestCtx) {
	param, _ := ctx.UserValue("id").(string)
	bytes, err := s.UserUseCase.GetPublicUserById(param)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.SetBody([]byte{})
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(bytes)
}

// func (s *userHandler) UpdateUserInfo(ctx *fasthttp.RequestCtx) {
// 	user := new(domain.User)
// 	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
// 		log.Printf("error while unmarshalling JSON: %s", err)
// 		ctx.SetStatusCode(fasthttp.StatusBadRequest)
// 		return
// 	}

// 	code, err := s.UserUseCase.UpdateUserInfo(user)
// 	ctx.SetStatusCode(code)
// 	if err != nil {
// 		log.Printf("error while registering user in", err)
// 		return
// 	}
// }

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
