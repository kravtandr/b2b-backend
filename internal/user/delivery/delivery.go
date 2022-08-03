package userDelivery

import (
	cnst "b2b/m/pkg/constants"
	"b2b/m/pkg/domain"
	"encoding/json"
	"fmt"
	"log"

	ccd "b2b/m/internal/cookie/delivery"
	ur "b2b/m/internal/user/repository"
	uu "b2b/m/internal/user/usecase"

	"github.com/fasthttp/router"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

type UserHandler interface {
	Login(ctx *fasthttp.RequestCtx)
	Registration(ctx *fasthttp.RequestCtx)
	Logout(ctx *fasthttp.RequestCtx)
	GetByEmail(ctx *fasthttp.RequestCtx)
	// GetCompanyById(ctx *fasthttp.RequestCtx)
	// GetCompaniesByCategoryId(ctx *fasthttp.RequestCtx)
	// SearchCompanies(ctx *fasthttp.RequestCtx)
}

type userHandler struct {
	UserUseCase   domain.UserUseCase
	CookieHandler ccd.CookieHandler
}

func NewUserHandler(UserUseCase domain.UserUseCase, CookieHandler ccd.CookieHandler) UserHandler {
	return &userHandler{
		UserUseCase:   UserUseCase,
		CookieHandler: CookieHandler,
	}
}

func CreateDelivery(db *pgxpool.Pool) UserHandler {
	cookieLayer := ccd.CreateDelivery(db)
	userLayer := NewUserHandler(uu.NewUserUseCase(ur.NewUserStorage(db)), cookieLayer)
	return userLayer
}

func SetUpUserRouter(db *pgxpool.Pool, r *router.Router) *router.Router {
	userHandler := CreateDelivery(db)
	r.POST(cnst.LoginURL, userHandler.Login)
	r.POST(cnst.RegisterURL, userHandler.Registration)
	r.POST(cnst.CheckEmailURL, userHandler.GetByEmail)
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

func (s *userHandler) Logout(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	s.CookieHandler.DeleteCookie(ctx, string(ctx.Request.Header.Cookie(cnst.CookieName)))
}

func (s *userHandler) GetByEmail(ctx *fasthttp.RequestCtx) {
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

// func (s *userHandler) GetByEmail(ctx *fasthttp.RequestCtx) {
// 	email := new(string)
// 	if err := json.Unmarshal(ctx.PostBody(), &email); err != nil {
// 		log.Printf("error while unmarshalling JSON: %s", err)
// 		ctx.SetStatusCode(fasthttp.StatusBadRequest)
// 		return
// 	}
// 	user, err := s.UserUseCase.GetByEmail(*email)
// 	if err != nil {
// 		log.Printf("error while GetByEmail user", err)
// 		ctx.SetStatusCode(fasthttp.StatusNotFound)
// 		return
// 	}
// 	if (user == domain.User{}) {
// 		ctx.SetStatusCode(fasthttp.StatusNotFound)
// 	} else {
// 		ctx.SetStatusCode(fasthttp.StatusOK)
// 	}
// }

// func (s *companyHandler) GetCompanyById(ctx *fasthttp.RequestCtx) {
// 	param, _ := ctx.UserValue("id").(string)
// 	bytes, err := s.CompanyUseCase.GetCompanyById(param)
// 	if err != nil {
// 		ctx.SetStatusCode(fasthttp.StatusNotFound)
// 		log.Printf(": %s", err)
// 		ctx.Write([]byte("{}"))
// 		return
// 	}
// 	ctx.SetStatusCode(fasthttp.StatusOK)
// 	ctx.Write(bytes)
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
