package models

import "time"

type UniqueCheck struct {
	UserId    int64
	ProductId int64
}

type Chat struct {
	Id        int64
	Name      string
	CreatorId int64
	ProductId int64
	Status    string
	Blured    bool
}

type Chats []Chat

type Msg struct {
	Id           int64
	ChatId       int64
	SenderId     int64
	ReceiverId   int64
	SenderName   string
	ReceiverName string
	Checked      bool
	Text         string
	Type         string
	Time         time.Time
}

type Msgs []Msg

type ChatAndLastMsg struct {
	Chat    Chat
	LastMsg Msg
}

type ChatsAndLastMsg []ChatAndLastMsg
