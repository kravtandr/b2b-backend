package delivery

import (
	"b2b/m/internal/gateway/chat/usecase"
	"b2b/m/internal/models"
	cnst "b2b/m/pkg/constants"
	chttp "b2b/m/pkg/customhttp"
	"b2b/m/pkg/error_adapter"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/valyala/fasthttp"
)

type ChatDelivery interface {
	CheckIfUniqChat(ctx *fasthttp.RequestCtx)
	InitChat(ctx *fasthttp.RequestCtx)
	UpdateChatStatus(ctx *fasthttp.RequestCtx)
	DeleteChat(ctx *fasthttp.RequestCtx)
	GetAllChatsAndLastMsg(ctx *fasthttp.RequestCtx)
	GetMsgsFromChat(ctx *fasthttp.RequestCtx)
	GetAllChats(ctx *fasthttp.RequestCtx)
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

func (u *chatDelivery) CheckIfUniqChat(ctx *fasthttp.RequestCtx) {
	userId := ctx.UserValue(cnst.UserIDContextKey).(int64)
	var request = &models.CheckUnique{}
	if err := json.Unmarshal(ctx.Request.Body(), request); err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBody([]byte(cnst.WrongRequestBody))
		return
	}
	if userId != request.SenderId {
		ctx.SetStatusCode(http.StatusUnauthorized)
		ctx.SetBody([]byte("userId != senderId"))
		return
	}

	response, err := u.manager.CheckIfUniqChat(ctx, userId, request.ProducId)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(httpError.MSG))
		return
	}
	b, err := chttp.ApiResp(response, err)
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		return
	}
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)
}

func (u *chatDelivery) InitChat(ctx *fasthttp.RequestCtx) {
	userId := ctx.UserValue(cnst.UserIDContextKey).(int64)
	var request = &models.CheckUnique{}
	if err := json.Unmarshal(ctx.Request.Body(), request); err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBody([]byte(cnst.WrongRequestBody))
		return
	}
	if userId != request.SenderId {
		ctx.SetStatusCode(http.StatusUnauthorized)
		ctx.SetBody([]byte("userId != senderId"))
		return
	}
	newChat, chat_id, err := u.manager.InitChat(ctx, request.SenderId, request.ProducId)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(fmt.Sprint(err)))
		return
	}
	b, err := chttp.ApiResp(models.InitChatResponce{ChatId: chat_id, CreateNewChat: newChat}, err)
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody(b)
}

func (u *chatDelivery) UpdateChatStatus(ctx *fasthttp.RequestCtx) {
	var request = &models.UpdateChatStatusRequest{}
	if err := json.Unmarshal(ctx.Request.Body(), request); err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBody([]byte(cnst.WrongRequestBody))
		return
	}
	response, err := u.manager.UpdateChatStatus(ctx, request.Chat_id, request.Status)
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

func (u *chatDelivery) DeleteChat(ctx *fasthttp.RequestCtx) {
	var request = &models.DeleteChatRequest{}
	if err := json.Unmarshal(ctx.Request.Body(), request); err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.SetBody([]byte(cnst.WrongRequestBody))
		return
	}

	response, err := u.manager.DeleteChat(ctx, request)
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

func (u *chatDelivery) GetMsgsFromChat(ctx *fasthttp.RequestCtx) {
	log.Println("start ChatDelivery -> GetMsgsFromChat")
	userId := ctx.UserValue(cnst.UserIDContextKey).(int64)
	chat_id, _ := strconv.ParseInt(ctx.UserValue("id").(string), 10, 64)
	response, err := u.manager.GetMsgsFromChat(ctx, chat_id, userId)
	log.Println("userId from cookie:", userId)
	if err != nil {
		httpError := u.errorAdapter.AdaptError(err)
		ctx.SetStatusCode(httpError.Code)
		ctx.SetBody([]byte(err.Error()))
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
