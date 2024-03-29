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
	init_msg := &models.InitMsg{}
	// WSClients = append(WSClients, models.Client{Id: userID, Ws: ws})
	defer ws.Close()
	knownConn := false
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("ERROR: ws.ReadMessage():", err)
			break
		}
		if !knownConn {
			if err := json.Unmarshal(message, init_msg); err != nil {
				log.Println("Unmarshal err:", err)
			}
			WSClients[init_msg.SenderId] = ws
			knownConn = true

		} else {
			if err != nil {
				log.Println("WSClients: ", WSClients)
				log.Println("read:", err)
			} else if err := json.Unmarshal(message, msg); err != nil {
				log.Println("Unmarshal err:", err)
			} else {
				ctx := context.Background()
				newMsg_struct := &models.Msg{Text: msg.Text, SenderId: msg.SenderId, ReceiverId: msg.ReceiverId, ChatId: msg.ChatId, Type: msg.Type}
				wsmsg := &models.WsMsg{Text: newMsg_struct.Text, ChatId: newMsg_struct.ChatId, SenderId: newMsg_struct.SenderId, ReceiverId: newMsg_struct.ReceiverId, Type: newMsg_struct.Type}
				_, err := u.manager.WriteNewMsg(ctx, newMsg_struct)
				if err != nil {
					log.Println("ERROR: WSChatLoop->WriteNewMsg", err)
				}
				if (WSClients[msg.ReceiverId] == &websocket.Conn{}) {
					log.Println("WARN: Reciever WS client status offline")
				} else {
					log.Println("WS_INFO: Reciever WS client status online")
					WSClients[msg.ReceiverId].WriteJSON(wsmsg)
				}
				if err != nil {
					log.Println("ERROR: WSChatLoop->GetMsgsFromChat", err)
				}
			}

		}
	}
}

func (u *chatDelivery) WSUpgradeRequest(ctx *fasthttp.RequestCtx) {
	// надо получить айди пользователя
	//userID, err := u.auth_manager.GetUserIdByCookie(ctx, string(ctx.Request.Header.Cookie(cnst.CookieName)))
	// if err != nil {
	// 	log.Println("Error: WSUpgradeRequest -> GetUserIdByCookie ", err)
	// 	return
	// }

	//log.Println("USER ID TO WS CONN", userID)
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
