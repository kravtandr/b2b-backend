package delivery

import (
	"b2b/m/internal/gateway/chat/usecase"
	"b2b/m/pkg/error_adapter"
	"encoding/json"
	"github.com/fasthttp/websocket"
	"log"

	"github.com/valyala/fasthttp"
)

type ChatDelivery interface {
	WSUpgradeRequest(ctx *fasthttp.RequestCtx)
	//Login(ctx *fasthttp.RequestCtx)
	//Logout(ctx *fasthttp.RequestCtx)
	//Register(ctx *fasthttp.RequestCtx)
	//GetUserInfo(ctx *fasthttp.RequestCtx)
	//FastRegister(ctx *fasthttp.RequestCtx)
	//GetProfile(ctx *fasthttp.RequestCtx)
	//UpdateProfile(ctx *fasthttp.RequestCtx)
	//CheckEmail(ctx *fasthttp.RequestCtx)
}
type Msg struct {
	Msg       string `json:"msg"`
	UserID    int64  `json:"userID"`
	ProductID int64  `json:"productID"`
}

type chatDelivery struct {
	errorAdapter error_adapter.HttpAdapter
	manager      usecase.ChatUsecase
}

func echoView(ctx *fasthttp.RequestCtx) {
	var upgrader = websocket.FastHTTPUpgrader{}
	err := upgrader.Upgrade(ctx, func(ws *websocket.Conn) {
		defer ws.Close()
		//первое сообщение приходит с фронта
		msg := Msg{Msg: "Сколько единиц в комлекте?", UserID: 1, ProductID: 1}
		bytes, _ := json.Marshal(msg)
		// 1 - binary, 2 - text
		err := ws.WriteMessage(1, bytes)
		//initDB(ctx)
		if err != nil {
			log.Println("WS write:", err)
		}
		for {
			mt, message, err := ws.ReadMessage()
			if err != nil {
				//когда приходит сообщение записываю его в бд
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", message)
			//когда отправляю сообщение записываю его в бд
			err = ws.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	})

	if err != nil {
		if _, ok := err.(websocket.HandshakeError); ok {
			log.Println(err)
		}
		return
	}
}

func (u *chatDelivery) WSUpgradeRequest(ctx *fasthttp.RequestCtx) {

	echoView(ctx)
	//var request = &models.LoginUserRequest{}
	//if err := json.Unmarshal(ctx.Request.Body(), request); err != nil {
	//	ctx.SetStatusCode(http.StatusBadRequest)
	//	ctx.SetBody([]byte(cnst.WrongRequestBody))
	//	return
	//}
	//
	//response, err := u.manager.Login(ctx, request)
	//if err != nil {
	//	httpError := u.errorAdapter.AdaptError(err)
	//	ctx.SetStatusCode(httpError.Code)
	//	ctx.SetBody([]byte(httpError.MSG))
	//	return
	//}
	//
	//b, err := chttp.ApiResp(response, err)
	//ctx.SetStatusCode(http.StatusOK)
	//ctx.SetBody(b)
	//
	//var c fasthttp.Cookie
	//c.SetKey(cnst.CookieName)
	//c.SetValue(response.Cookie)
	//c.SetMaxAge(int(time.Hour))
	//c.SetHTTPOnly(true)
	//c.SetSameSite(fasthttp.CookieSameSiteStrictMode)
	//ctx.Response.Header.SetCookie(&c)
}

//
//func (u *userDelivery) Login(ctx *fasthttp.RequestCtx) {
//	var request = &models.LoginUserRequest{}
//	if err := json.Unmarshal(ctx.Request.Body(), request); err != nil {
//		ctx.SetStatusCode(http.StatusBadRequest)
//		ctx.SetBody([]byte(cnst.WrongRequestBody))
//		return
//	}
//
//	response, err := u.manager.Login(ctx, request)
//	if err != nil {
//		httpError := u.errorAdapter.AdaptError(err)
//		ctx.SetStatusCode(httpError.Code)
//		ctx.SetBody([]byte(httpError.MSG))
//		return
//	}
//
//	b, err := chttp.ApiResp(response, err)
//	ctx.SetStatusCode(http.StatusOK)
//	ctx.SetBody(b)
//
//	var c fasthttp.Cookie
//	c.SetKey(cnst.CookieName)
//	c.SetValue(response.Cookie)
//	c.SetMaxAge(int(time.Hour))
//	c.SetHTTPOnly(true)
//	c.SetSameSite(fasthttp.CookieSameSiteStrictMode)
//	ctx.Response.Header.SetCookie(&c)
//}
//
//func (u *userDelivery) Logout(ctx *fasthttp.RequestCtx) {
//	err := u.manager.Logout(ctx, string(ctx.Request.Header.Cookie(cnst.CookieName)))
//	if err != nil {
//		httpError := u.errorAdapter.AdaptError(err)
//		ctx.SetStatusCode(httpError.Code)
//		ctx.SetBody([]byte(httpError.MSG))
//		return
//	}
//
//	ctx.SetStatusCode(http.StatusOK)
//}
//
//func (u *userDelivery) FastRegister(ctx *fasthttp.RequestCtx) {
//	var request = &models.FastRegistrationForm{}
//	if err := json.Unmarshal(ctx.Request.Body(), request); err != nil {
//		ctx.SetStatusCode(http.StatusBadRequest)
//		ctx.SetBody([]byte(cnst.WrongRequestBody))
//		return
//	}
//
//	response, err := u.manager.FastRegister(ctx, request)
//	if err != nil {
//		httpError := u.errorAdapter.AdaptError(err)
//		ctx.SetStatusCode(httpError.Code)
//		ctx.SetBody([]byte(httpError.MSG))
//		return
//	}
//
//	b, err := chttp.ApiResp(response, err)
//	ctx.SetStatusCode(http.StatusOK)
//	ctx.SetBody(b)
//
//	var c fasthttp.Cookie
//	c.SetKey(cnst.CookieName)
//	c.SetValue(response.Cookie)
//	c.SetMaxAge(int(time.Hour))
//	c.SetHTTPOnly(true)
//	c.SetSameSite(fasthttp.CookieSameSiteStrictMode)
//	ctx.Response.Header.SetCookie(&c)
//}
//
//func (u *userDelivery) Register(ctx *fasthttp.RequestCtx) {
//	var request = &models.RegisterUserRequest{}
//	if err := json.Unmarshal(ctx.Request.Body(), request); err != nil {
//		ctx.SetStatusCode(http.StatusBadRequest)
//		ctx.SetBody([]byte(cnst.WrongRequestBody))
//		return
//	}
//
//	response, err := u.manager.Register(ctx, request)
//	if err != nil {
//		httpError := u.errorAdapter.AdaptError(err)
//		ctx.SetStatusCode(httpError.Code)
//		ctx.SetBody([]byte(httpError.MSG))
//		return
//	}
//
//	b, err := chttp.ApiResp(response, err)
//	ctx.SetStatusCode(http.StatusOK)
//	ctx.SetBody(b)
//
//	var c fasthttp.Cookie
//	c.SetKey(cnst.CookieName)
//	c.SetValue(response.Cookie)
//	c.SetMaxAge(int(time.Hour))
//	c.SetHTTPOnly(true)
//	c.SetSameSite(fasthttp.CookieSameSiteStrictMode)
//	ctx.Response.Header.SetCookie(&c)
//}
//
//func (u *userDelivery) GetProfile(ctx *fasthttp.RequestCtx) {
//	userID := ctx.UserValue(cnst.UserIDContextKey).(int)
//	response, err := u.manager.Profile(ctx, userID)
//	if err != nil {
//		httpError := u.errorAdapter.AdaptError(err)
//		ctx.SetStatusCode(httpError.Code)
//		ctx.SetBody([]byte(httpError.MSG))
//		return
//	}
//
//	b, err := chttp.ApiResp(response, err)
//	ctx.SetStatusCode(http.StatusOK)
//	ctx.SetBody(b)
//}
//
//func (u *userDelivery) UpdateProfile(ctx *fasthttp.RequestCtx) {
//	var request = &models.PublicCompanyAndOwnerRequest{}
//	if err := json.Unmarshal(ctx.Request.Body(), request); err != nil {
//		ctx.SetStatusCode(http.StatusBadRequest)
//		ctx.SetBody([]byte(cnst.WrongRequestBody))
//		return
//	}
//
//	userID, err := u.manager.GetUserIdByCookie(ctx, string(ctx.Request.Header.Cookie(cnst.CookieName)))
//	if err != nil {
//		httpError := u.errorAdapter.AdaptError(err)
//		ctx.SetStatusCode(httpError.Code)
//		ctx.SetBody([]byte(httpError.MSG))
//		return
//	}
//	response, err := u.manager.UpdateProfile(ctx, userID, request)
//
//	b, err := chttp.ApiResp(response, err)
//	ctx.SetStatusCode(http.StatusOK)
//	ctx.SetBody(b)
//}
//
//func (u *userDelivery) CheckEmail(ctx *fasthttp.RequestCtx) {
//	var request = &models.Email{}
//	if err := json.Unmarshal(ctx.Request.Body(), request); err != nil {
//		ctx.SetStatusCode(http.StatusBadRequest)
//		ctx.SetBody([]byte(cnst.WrongRequestBody))
//		return
//	}
//
//	publicUser, err := u.manager.CheckEmail(ctx, request)
//	if err != nil {
//		httpError := u.errorAdapter.AdaptError(err)
//		ctx.SetStatusCode(httpError.Code)
//		ctx.SetBody([]byte(httpError.MSG))
//		return
//	}
//
//	b, err := chttp.ApiResp(publicUser, err)
//	ctx.SetStatusCode(http.StatusOK)
//	ctx.SetBody(b)
//}
//
//func (u *userDelivery) GetUserInfo(ctx *fasthttp.RequestCtx) {
//	param, _ := strconv.Atoi(ctx.UserValue("id").(string))
//	response, err := u.manager.GetUserInfo(ctx, param)
//	if err != nil {
//		httpError := u.errorAdapter.AdaptError(err)
//		ctx.SetStatusCode(httpError.Code)
//		ctx.SetBody([]byte(httpError.MSG))
//		return
//	}
//	b, err := chttp.ApiResp(response, err)
//	ctx.SetStatusCode(http.StatusOK)
//	ctx.SetBody(b)
//}

func NewChatDelivery(errorAdapter error_adapter.HttpAdapter, manager usecase.ChatUsecase) ChatDelivery {
	return &chatDelivery{
		errorAdapter: errorAdapter,
		manager:      manager,
	}
}
