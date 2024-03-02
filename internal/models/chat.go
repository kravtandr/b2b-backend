package models

import "github.com/fasthttp/websocket"

type Chat struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CreatorId int64  `json:"creator_id"`
	ProductId int64  `json:"product_id"`
	Status    string `json:"status"`
}

type Chats []Chat

type Msg struct {
	Id           int64  `json:"id"`
	ChatId       int64  `json:"chat_id"`
	SenderId     int64  `json:"sender_id"`
	ReceiverId   int64  `json:"receiver_id"`
	SenderName   string `json:"sender_name"`
	ReceiverName string `json:"receiver_name"`
	Checked      bool   `json:"checked"`
	Text         string `json:"text"`
	Type         string `json:"type"`
	Time         string `json:"time"`
}

type WsMsg struct {
	ChatId     int64  `json:"chat_id"`
	SenderId   int64  `json:"sender_id"`
	ReceiverId int64  `json:"receiver_id"`
	Text       string `json:"text"`
	Type       string `json:"type"`
}

type InitMsg struct {
	SenderId int64 `json:"sender_id"`
}
type DeleteChatRequest struct {
	Chat_id int64 `json:"chat_id"`
}

type Msgs []Msg

type ChatAndLastMsg struct {
	Chat    Chat `json:"chat"`
	LastMsg Msg  `json:"last_msg"`
}

type ChatsAndLastMsg []ChatAndLastMsg

type CheckUnique struct {
	SenderId int64 `json:"sender_id"`
	ProducId int64 `json:"product_id"`
}

type InitChatResponce struct {
	ChatId        int64 `json:"chat_id"`
	CreateNewChat bool  `json:"create_new_chat"`
}

type Client struct {
	Id int64
	Ws *websocket.Conn
}

type Clients []Client
