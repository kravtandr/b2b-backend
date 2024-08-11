package delivery

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"b2b/m/internal/gateway/user/usecase"
	"b2b/m/internal/models"
	cnst "b2b/m/pkg/constants"
	chttp "b2b/m/pkg/customhttp"
	"b2b/m/pkg/error_adapter"

	"github.com/valyala/fasthttp"
)

type UserDelivery interface {
	Register(ctx *fasthttp.RequestCtx)
	FastRegister(ctx *fasthttp.RequestCtx)
	Login(ctx *fasthttp.RequestCtx)
	Logout(ctx *fasthttp.RequestCtx)

	GetUserInfoById(ctx *fasthttp.RequestCtx)
	GetUserByCookie(ctx *fasthttp.RequestCtx)

	GetProfile(ctx *fasthttp.RequestCtx)
	UpdateProfile(ctx *fasthttp.RequestCtx)

	CheckEmail(ctx *fasthttp.RequestCtx)

	CreatePayemntAddBalance(ctx *fasthttp.RequestCtx)
	GetUsersPayments(ctx *fasthttp.RequestCtx)
	CheckPayment(ctx *fasthttp.RequestCtx)
	HandlePaidPayments(ctx *fasthttp.RequestCtx)
	CountUsersPayments(ctx *fasthttp.RequestCtx)
}

type userDelivery struct {
	errorAdapter error_adapter.HttpAdapter
	manager      usecase.UserUsecase
}

func (u *userDelivery) CountUsersPayments(ctx *fasthttp.RequestCtx) {
	userID, err := u.manager.GetUserIdByCookie(ctx, string(ctx.Request.Header.Cookie(cnst.CookieName)))
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}
	credited, err := u.manager.CountUsersPayments(ctx, userID)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}
	b, _ := chttp.ApiResp(credited, err)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)
}

func (u *userDelivery) GetUsersPayments(ctx *fasthttp.RequestCtx) {

	userID, err := u.manager.GetUserIdByCookie(ctx, string(ctx.Request.Header.Cookie(cnst.CookieName)))
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}
	payments, err := u.manager.GetUsersPayments(ctx, userID)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}
	b, _ := chttp.ApiResp(payments, err)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)
}
func (u *userDelivery) HandlePaidPayments(ctx *fasthttp.RequestCtx) {

	userID, err := u.manager.GetUserIdByCookie(ctx, string(ctx.Request.Header.Cookie(cnst.CookieName)))
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}
	// save to db
	credited, err := u.manager.HandlePaidPayments(ctx, userID)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}
	b, _ := chttp.ApiResp(credited, err)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)
}

func (u *userDelivery) CreatePayemntAddBalance(ctx *fasthttp.RequestCtx) {
	userID, err := u.manager.GetUserIdByCookie(ctx, string(ctx.Request.Header.Cookie(cnst.CookieName)))
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}
	var request = &models.CreatePaymentRequest{}
	if err := json.Unmarshal(ctx.Request.Body(), request); err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBody([]byte(cnst.WrongRequestBody))
		return
	}
	request.User_id = userID
	// save to db

	response, err := u.manager.CreatePayment(ctx, request)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}
	b, _ := chttp.ApiResp(response, err)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)
}

func (u *userDelivery) CheckPayment(ctx *fasthttp.RequestCtx) {
	var request = &models.CheckPaymentRequest{}
	if err := json.Unmarshal(ctx.Request.Body(), request); err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBody([]byte(cnst.WrongRequestBody))
		return
	}
	response, err := u.manager.GetPaymentInfo(ctx, request.PaymentID)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}
	b, _ := chttp.ApiResp(response, err)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)
}

func (u *userDelivery) Login(ctx *fasthttp.RequestCtx) {
	var request = &models.LoginUserRequest{}
	if err := json.Unmarshal(ctx.Request.Body(), request); err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBody([]byte(cnst.WrongRequestBody))
		return
	}

	response, err := u.manager.Login(ctx, request)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}

	b, _ := chttp.ApiResp(response, err)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)
	log.Println("Login SetCookie:", response.Cookie, " id ", response.Id, " email ", response.Email)

	chttp.SetCookieAndSession(ctx, response.Cookie, response.Id)
}

func (u *userDelivery) Logout(ctx *fasthttp.RequestCtx) {
	err := u.manager.Logout(ctx, string(ctx.Request.Header.Cookie(cnst.CookieName)))
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}
	log.Println("Logout delite cookie:", string(ctx.Request.Header.Cookie(cnst.CookieName)))

	ctx.SetStatusCode(http.StatusOK)
}

func (u *userDelivery) FastRegister(ctx *fasthttp.RequestCtx) {
	var request = &models.FastRegistrationForm{}
	if err := json.Unmarshal(ctx.Request.Body(), request); err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBody([]byte(cnst.WrongRequestBody))
		return
	}

	response, err := u.manager.FastRegister(ctx, request)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}
	log.Println("FastRegister SetCookie:", response.Cookie, " id ", response.Id, " email ", response.Email)

	b, _ := chttp.ApiResp(response, err)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)

	chttp.SetCookieAndSession(ctx, response.Cookie, response.Id)
}

func (u *userDelivery) Register(ctx *fasthttp.RequestCtx) {
	var request = &models.RegisterUserRequest{}
	if err := json.Unmarshal(ctx.Request.Body(), request); err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBody([]byte(cnst.WrongRequestBody))
		return
	}

	response, err := u.manager.Register(ctx, request)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}

	b, _ := chttp.ApiResp(response, err)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)

	chttp.SetCookieAndSession(ctx, response.Cookie, response.Id)
}

func (u *userDelivery) GetProfile(ctx *fasthttp.RequestCtx) {
	userID := ctx.UserValue(cnst.UserIDContextKey).(int64)
	response, err := u.manager.Profile(ctx, userID)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}

	b, _ := chttp.ApiResp(response, err)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)
}

func (u *userDelivery) GetUserByCookie(ctx *fasthttp.RequestCtx) {
	userID := ctx.UserValue(cnst.UserIDContextKey).(int64)
	response, err := u.manager.Profile(ctx, userID)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}

	b, _ := chttp.ApiResp(response, err)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)
	log.Println("GetUserByCookie SetCookie:", string(ctx.Request.Header.Cookie(cnst.CookieName)), " id ", response.Id, " email ", response.Email)
}

func (u *userDelivery) UpdateProfile(ctx *fasthttp.RequestCtx) {
	var request = &models.PublicCompanyAndOwnerRequest{}
	if err := json.Unmarshal(ctx.Request.Body(), request); err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBody([]byte(cnst.WrongRequestBody))
		return
	}

	userID, err := u.manager.GetUserIdByCookie(ctx, string(ctx.Request.Header.Cookie(cnst.CookieName)))
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}
	response, err := u.manager.UpdateProfile(ctx, userID, request)

	b, _ := chttp.ApiResp(response, err)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)
}

func (u *userDelivery) CheckEmail(ctx *fasthttp.RequestCtx) {
	var request = &models.Email{}
	if err := json.Unmarshal(ctx.Request.Body(), request); err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBody([]byte(cnst.WrongRequestBody))
		return
	}

	publicUser, err := u.manager.CheckEmail(ctx, request)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}

	b, _ := chttp.ApiResp(publicUser, err)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)
}

func (u *userDelivery) GetUserInfoById(ctx *fasthttp.RequestCtx) {
	param, err := strconv.ParseInt(ctx.UserValue("id").(string), 10, 64)
	if err != nil {
		log.Println("GetUserInfo ERROR ParseInt", err)
		return
	}

	response, err := u.manager.GetUserInfo(ctx, param)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}
	b, _ := chttp.ApiResp(response, err)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)
}

func NewUserDelivery(
	errorAdapter error_adapter.HttpAdapter,
	manager usecase.UserUsecase,
) UserDelivery {
	return &userDelivery{
		errorAdapter: errorAdapter,
		manager:      manager,
	}
}
