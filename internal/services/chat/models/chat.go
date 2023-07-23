package models

type UniqueCheck struct {
	UserId    int64
	ProductId int64
}

type Chat struct {
	Id        int64
	Name      string
	CreatorId int64
	ProductId int64
}

type Chats []Chat

type Msg struct {
	Id         int64
	ChatId     int64
	SenderId   int64
	ReceiverId int64
	Checked    bool
	Text       string
	Type       string
	Time       string
}

type Msgs []Msg

type ChatAndLastMsg struct {
	Chat    Chat
	LastMsg Msg
}

type ChatsAndLastMsg []ChatAndLastMsg
