package delivery

import (
	"b2b/m/internal/gateway_chat/chat/usecase"
	"b2b/m/internal/models"
	"b2b/m/pkg/error_adapter"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/fasthttp/websocket"

	"github.com/valyala/fasthttp"
)

type ChatDelivery interface {
	WSUpgradeRequest(ctx *fasthttp.RequestCtx)
	TestCh(ctx *fasthttp.RequestCtx)
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

func (u *chatDelivery) TestCh(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody([]byte("PASS TEST chat gateway router"))
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
	//timeout := time.Now()
	// timeout ws connection 30 sec
	//timeout = timeout.Add(time.Second * 30)
	//ws.SetReadDeadline(timeout)

	msg := &models.Msg{}
	defer ws.Close()

	//первое сообщение приходит с фронта
	firstMsg := Msg{Text: "Сколько единиц в комлекте?", SenderID: 1, RecieverID: 1}
	bytes, _ := json.Marshal(firstMsg)
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
		}
		if err := json.Unmarshal(message, msg); err != nil {
			log.Println("Unmarshal err:", err)
		}
		log.Println("read:", msg)
		u.manager.WriteNewMsg(context.Background(), &models.Msg{Text: msg.Text, SenderId: msg.SenderId, ReceiverId: msg.ReceiverId, ChatId: msg.ChatId})
		log.Println("recv msg:", msg)
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
	log.Println("+++++ WSUpgradeRequest ++++++")
	upgrader.CheckOrigin = func(r *fasthttp.RequestCtx) bool { return true }
	// epic bullshit ^^^^^^^^^^^^^^^^^^^^^^^^^^^
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
