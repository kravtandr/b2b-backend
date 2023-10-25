package delivery

import (
	rest_chat "b2b/m/internal/gateway/chat/usecase"
	auth_usecase "b2b/m/internal/gateway/user/usecase"
	"b2b/m/internal/gateway_chat/chat/usecase"
	"b2b/m/internal/models"
	"b2b/m/pkg/error_adapter"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/fasthttp/websocket"

	cnst "b2b/m/pkg/constants"

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

// var WSClients = make(map()models.Clients, 1)
var WSClients = make(map[int64]*websocket.Conn)

type chatDelivery struct {
	errorAdapter      error_adapter.HttpAdapter
	manager           usecase.ChatUsecase
	rest_chat_manager rest_chat.ChatUsecase
	auth_manager      auth_usecase.UserUsecase
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

	// WSClients = append(WSClients, models.Client{Id: userID, Ws: ws})
	defer ws.Close()

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("WSClients: ", WSClients)
			//когда приходит сообщение записываю его в бд
			log.Println("read:", err)
		} else if err := json.Unmarshal(message, msg); err != nil {
			log.Println("Unmarshal err:", err)
		} else {
			ctx := context.Background()
			newMsg_id, err := u.manager.WriteNewMsg(ctx, &models.Msg{Text: msg.Text, SenderId: msg.SenderId, ReceiverId: msg.ReceiverId, ChatId: msg.ChatId, Type: msg.Type})
			if err != nil {
				log.Println("ERROR: WSChatLoop->WriteNewMsg", err)
			}
			WSClients[msg.SenderId] = ws
			// далее нужно прочичать все сообщения из этого чата, которые идут после только что отправленно и отправить их по ws
			// 1 получить все сообщения из чата
			msgs, err := u.rest_chat_manager.GetMsgsFromChat(ctx, msg.ChatId, msg.SenderId)
			var reducedMsgs models.Msgs
			for _, item := range *msgs {
				if item.Id > newMsg_id {
					reducedMsgs = append(reducedMsgs, item)
					ws.WriteJSON(item)
				}
			}
			WSClients[msg.ReceiverId].WriteJSON(newMsg_id)
			log.Println("reduced msgs: ", reducedMsgs)
			if err != nil {
				log.Println("ERROR: WSChatLoop->GetMsgsFromChat", err)
			}
			// 2 выделить те, которые идут после отправленного
			// 3 отправить сообщения по ws
		}
	}
}

func (u *chatDelivery) WSUpgradeRequest(ctx *fasthttp.RequestCtx) {
	// надо получить айди пользователя
	userID := ctx.UserValue(cnst.UserIDContextKey).(int64)
	log.Println("USER ID TO WS CONN", userID)
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

func NewChatDelivery(errorAdapter error_adapter.HttpAdapter, manager usecase.ChatUsecase, rest_chat_manager rest_chat.ChatUsecase, auth_manager auth_usecase.UserUsecase) ChatDelivery {
	return &chatDelivery{
		errorAdapter:      errorAdapter,
		manager:           manager,
		rest_chat_manager: rest_chat_manager,
		auth_manager:      auth_manager,
	}
}
