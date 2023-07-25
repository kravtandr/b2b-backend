package delivery

import (
	"b2b/m/internal/gateway/chat/usecase"
	"b2b/m/internal/models"
	cnst "b2b/m/pkg/constants"
	chttp "b2b/m/pkg/customhttp"
	"b2b/m/pkg/error_adapter"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/fasthttp/websocket"

	"github.com/valyala/fasthttp"
)

type ChatDelivery interface {
	WSUpgradeRequest(ctx *fasthttp.RequestCtx)
	GetAllChatsAndLastMsg(ctx *fasthttp.RequestCtx)
	GetMsgsFromChat(ctx *fasthttp.RequestCtx)
	GetAllChats(ctx *fasthttp.RequestCtx)
	InitChat(ctx *fasthttp.RequestCtx)
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

func (u *chatDelivery) ChatLogic(ctx *fasthttp.RequestCtx) {
	//when new chat init save chat id
	//get userId, productId in first msg
	//   userId from cookie
	//	 producId from front
	// get receiverId from company that owns product
	// if userId, productId, receiverId unique create new chat
	// else open old chat

}

func (u *chatDelivery) WSChatLoop(ws *websocket.Conn) {
	defer ws.Close()
	//первое сообщение приходит с фронта
	msg := Msg{Text: "Сколько единиц в комлекте?", SenderID: 1, RecieverID: 1}
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
		u.manager.WriteNewMsg(context.Background(), &models.Msg{Text: string(message), SenderId: 1, ReceiverId: 2, ChatId: 1})
		log.Printf("recv: %s", message)
		//когда отправляю сообщение записываю его в бд
		//echo
		err = ws.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func (u *chatDelivery) WSUpgradeRequest(ctx *fasthttp.RequestCtx) {
	var upgrader = websocket.FastHTTPUpgrader{}
	err := upgrader.Upgrade(ctx, u.WSChatLoop)

	if err != nil {
		if _, ok := err.(websocket.HandshakeError); ok {
			log.Println(err)
		}
		return
	}
}

func NewChatDelivery(errorAdapter error_adapter.HttpAdapter, manager usecase.ChatUsecase) ChatDelivery {
	return &chatDelivery{
		errorAdapter: errorAdapter,
		manager:      manager,
	}
}
