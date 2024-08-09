package usecase

import (
	"b2b/m/internal/services/chat/models"
	"context"
)

type chatRepository interface {
	NewChat(ctx context.Context, newChat *models.Chat) (*models.Chat, error)
	GetChat(ctx context.Context, chat *models.Chat) (*models.Chat, error)
	GetChatById(ctx context.Context, id int64) (*models.Chat, error)
	CheckIfUniqChat(ctx context.Context, productId int64, userId int64) (bool, error)
	UpdateChat(ctx context.Context, chat *models.Chat) error
	DeleteChat(ctx context.Context, chat_id int64) (bool, error)

	GetUserCreatedChats(ctx context.Context, userId int64) (*models.Chats, error)

	WriteNewMsg(ctx context.Context, newMsg *models.Msg) (int64, error)
	GetMsgsFromChat(ctx context.Context, chatId int64, userId int64) (*models.Msgs, error)
	GetMsgById(ctx context.Context, msgId int64) (*models.Msg, error)
	GetUserLastMsgs(ctx context.Context, userId int64) (*models.Msgs, error)
	UpdateMsg(ctx context.Context, msg *models.Msg) error

	GetAllChatsAndLastMsg(ctx context.Context, userId int64) (*models.ChatsAndLastMsg, error)
	CombineChatsWithAndWithoutMsgs(ctx context.Context, onlyChats *models.Chats, chatsAndLM *models.ChatsAndLastMsg) *models.ChatsAndLastMsg
}

type hasher interface {
	EncodeString(value string) string
	DecodeString(value string) (string, error)
}
