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
	// Login(ctx *fasthttp.RequestCtx)
	// Registration(ctx *fasthttp.RequestCtx)
	FastOrder(ctx *fasthttp.RequestCtx)
	// Logout(ctx *fasthttp.RequestCtx)
	// GetPublicUserByEmail(ctx *fasthttp.RequestCtx)
	// GetPublicUserById(ctx *fasthttp.RequestCtx)
	// GetCompanyById(ctx *fasthttp.RequestCtx)
	// GetCompaniesByCategoryId(ctx *fasthttp.RequestCtx)
	// SearchCompanies(ctx *fasthttp.RequestCtx)
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
	// r.POST(cnst.LoginURL, userHandler.Login)
	// r.POST(cnst.RegisterURL, userHandler.Registration)
	// r.POST(cnst.CheckEmailURL, userHandler.GetPublicUserByEmail)
	r.POST(cnst.FastOrderURL, OrderFormHandler.FastOrder)
	// r.GET(cnst.UserURL, userHandler.GetPublicUserById)
	// r.GET(cnst.CompanyURL, companyHandler.GetCompanyById)
	// r.GET(cnst.CategoryURL, companyHandler.GetCompaniesByCategoryId)
	// r.POST(cnst.CompanySearchURL, companyHandler.SearchCompanies)
	return r
}

// func (s *userHandler) Login(ctx *fasthttp.RequestCtx) {
// 	user := new(domain.User)
// 	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
// 		log.Printf("error while unmarshalling JSON: %s", err)
// 		ctx.SetStatusCode(fasthttp.StatusBadRequest)
// 		return
// 	}
// 	valid := s.UserUseCase.Validate(user)
// 	if !valid {
// 		log.Printf("error while validating user")
// 		ctx.SetStatusCode(fasthttp.StatusBadRequest)
// 		return
// 	}

// 	code, err := s.UserUseCase.Login(user)
// 	ctx.SetStatusCode(code)
// 	if err != nil {
// 		log.Printf("error while logging user in")
// 		return
// 	}

// 	с := fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(user.Email)))
// 	foundUser, _ := s.UserUseCase.GetByEmail(user.Email)
// 	s.CookieHandler.SetCookieAndToken(ctx, с, foundUser.Id)
// }

// func (s *userHandler) Registration(ctx *fasthttp.RequestCtx) {
// 	user := new(domain.User)
// 	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
// 		log.Printf("error while unmarshalling JSON: %s", err)
// 		ctx.SetStatusCode(fasthttp.StatusBadRequest)
// 		return
// 	}

// 	code, err := s.UserUseCase.Registration(user)
// 	ctx.SetStatusCode(code)
// 	if err != nil {
// 		log.Printf("error while registering user in", err)
// 		return
// 	}

// 	с := fmt.Sprint(uuid.NewMD5(uuid.UUID{}, []byte(user.Email)))
// 	newUser, _ := s.UserUseCase.GetByEmail(user.Email)
// 	s.CookieHandler.SetCookieAndToken(ctx, с, newUser.Id)
// }

func (s *orderFormHandler) FastOrder(ctx *fasthttp.RequestCtx) {
	form := new(domain.OrderForm)
	if err := json.Unmarshal(ctx.PostBody(), &form); err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	code, err := s.OrderFormUseCase.Add(*form)
	if err != nil {
		ctx.SetStatusCode(code)
		log.Printf("error while AddBaseCompany", err)
		return
	}
}

// func (s *userHandler) Logout(ctx *fasthttp.RequestCtx) {
// 	ctx.SetStatusCode(fasthttp.StatusOK)
// 	s.CookieHandler.DeleteCookie(ctx, string(ctx.Request.Header.Cookie(cnst.CookieName)))
// }

// func (s *userHandler) GetPublicUserByEmail(ctx *fasthttp.RequestCtx) {
// 	request := new(domain.UserEmail)
// 	if err := json.Unmarshal(ctx.Request.Body(), &request); err != nil {
// 		ctx.SetStatusCode(fasthttp.StatusBadRequest)
// 		ctx.SetBody([]byte{})
// 		return
// 	}
// 	bytes, err := s.UserUseCase.GetPublicUserByEmail(request.Email)
// 	if err != nil {
// 		ctx.SetStatusCode(fasthttp.StatusBadRequest)
// 		ctx.SetBody([]byte{})
// 		return
// 	}
// 	ctx.SetStatusCode(fasthttp.StatusOK)
// 	ctx.Write(bytes)
// }

// func (s *userHandler) GetPublicUserById(ctx *fasthttp.RequestCtx) {
// 	param, _ := ctx.UserValue("id").(string)
// 	bytes, err := s.UserUseCase.GetPublicUserById(param)
// 	if err != nil {
// 		ctx.SetStatusCode(fasthttp.StatusNotFound)
// 		ctx.SetBody([]byte{})
// 		return
// 	}
// 	ctx.SetStatusCode(fasthttp.StatusOK)
// 	ctx.Write(bytes)
// }

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
