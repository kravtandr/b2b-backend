package delivery

import (
	"b2b/m/internal/gateway/chat/usecase"
	cnst "b2b/m/pkg/constants"
	chttp "b2b/m/pkg/customhttp"
	"b2b/m/pkg/error_adapter"
	"log"
	"net/http"
	"strconv"

	"github.com/valyala/fasthttp"
)

type ChatDelivery interface {
	GetAllChatsAndLastMsg(ctx *fasthttp.RequestCtx)
	GetMsgsFromChat(ctx *fasthttp.RequestCtx)
	GetAllChats(ctx *fasthttp.RequestCtx)
	InitChat(ctx *fasthttp.RequestCtx)
	TestGw(ctx *fasthttp.RequestCtx)
}
type Msg struct {
	SenderID   int64  `json:"senderID"`
	RecieverID int64  `json:"receiverID"`
	Text       string `json:"text"`
}

type chatDelivery struct {
	errorAdapter error_adapter.HttpAdapter
	manager      usecase.ChatUsecase
}

func (u *chatDelivery) TestGw(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody([]byte("PASS working gw router"))
}

func (u *chatDelivery) GetAllChatsAndLastMsg(ctx *fasthttp.RequestCtx) {
	userId := ctx.UserValue(cnst.UserIDContextKey).(int64)
	response, err := u.manager.GetAllChatsAndLastMsg(ctx, userId)
	log.Println("userId from cookie:", userId)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}
	b, err := chttp.ApiResp(response, err)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)
}

func (u *chatDelivery) GetAllChats(ctx *fasthttp.RequestCtx) {
	userId := ctx.UserValue(cnst.UserIDContextKey).(int64)
	cookie := string(ctx.Request.Header.Cookie(cnst.CookieName))
	response, err := u.manager.GetAllUserChats(ctx, userId, cookie)

	log.Println("userId from cookie:", userId)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}
	b, err := chttp.ApiResp(response, err)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)
}

func (u *chatDelivery) InitChat(ctx *fasthttp.RequestCtx) {
	userId := ctx.UserValue(cnst.UserIDContextKey).(int64)
	product_id, err := strconv.ParseInt(ctx.UserValue("id").(string), 10, 64)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}
	newChat, chat_id, err := u.manager.InitChat(ctx, userId, product_id)

	log.Println("+++++ CheckIfUniqChat userId from cookie: ", userId, " +++++")
	log.Println("+++++ CheckIfUniqChat newChat : ", newChat, " +++++")
	log.Println("+++++ CheckIfUniqChat chat_id : ", chat_id, " +++++")
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}
	b, err := chttp.ApiResp(chat_id, err)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)
}

func (u *chatDelivery) GetMsgsFromChat(ctx *fasthttp.RequestCtx) {
	log.Println("start ChatDelivery -> GetMsgsFromChat")
	userId := ctx.UserValue(cnst.UserIDContextKey).(int64)
	chat_id, _ := strconv.ParseInt(ctx.UserValue("id").(string), 10, 64)
	response, err := u.manager.GetMsgsFromChat(ctx, chat_id, userId)
	log.Println("userId from cookie:", userId)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}
	b, err := chttp.ApiResp(response, err)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)
}

func NewChatDelivery(errorAdapter error_adapter.HttpAdapter, manager usecase.ChatUsecase) ChatDelivery {
	return &chatDelivery{
		errorAdapter: errorAdapter,
		manager:      manager,
	}
}
