package usecase

import (
	"b2b/m/internal/services/chat/models"
	"context"
)

type chatRepository interface {
	CheckIfUniqChat(ctx context.Context, productId int64, userId int64) (bool, error)
	NewChat(ctx context.Context, newChat *models.Chat) (*models.Chat, error)
	GetChat(ctx context.Context, chat *models.Chat) (*models.Chat, error)
	WriteNewMsg(ctx context.Context, newMsg *models.Msg) (int64, error)
	GetMsgsFromChat(ctx context.Context, chatId int64, userId int64) (*models.Msgs, error)
	GetAllChatsAndLastMsg(ctx context.Context, userId int64) (*models.ChatsAndLastMsg, error)
	GetAllUserChats(ctx context.Context, userId int64) (*models.Chats, error)
	GetUserLastMsgs(ctx context.Context, userId int64) (*models.Msgs, error)
}

type hasher interface {
	EncodeString(value string) string
	DecodeString(value string) (string, error)
}
