package repository

import (
	"context"

	"b2b/m/internal/services/chat/models"

	"go.uber.org/zap"
)

const (
	module = "auth_repo"
)

type loggingMiddleware struct {
	logger *zap.SugaredLogger

	next ChatRepository
}

func NewLoggingMiddleware(logger *zap.SugaredLogger, next ChatRepository) ChatRepository {
	return &loggingMiddleware{
		logger: logger,
		next:   next,
	}
}

func (l *loggingMiddleware) UpdateMsg(ctx context.Context, msg *models.Msg) (err error) {
	l.logger.Infow(module,
		"Action", "UpdateMsg",
		"Request", msg,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "UpdateMsg",
				"Request", msg,
				"Error", err,
			)
		}
	}()
	return l.next.UpdateMsg(ctx, msg)
}

func (l *loggingMiddleware) GetMsgById(ctx context.Context, msgId int64) (msg *models.Msg, err error) {
	l.logger.Infow(module,
		"Action", "GetMsgById",
		"Request", msgId,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetMsgById",
				"Request", msgId,
				"Error", err,
			)
		}
	}()
	return l.next.GetMsgById(ctx, msgId)
}

func (l *loggingMiddleware) UpdateChat(ctx context.Context, chat *models.Chat) (err error) {
	l.logger.Infow(module,
		"Action", "UpdateChat",
		"Request", chat,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "UpdateChat",
				"Request", chat,
				"Error", err,
			)
		}
	}()
	return l.next.UpdateChat(ctx, chat)
}

func (l *loggingMiddleware) DeleteChat(ctx context.Context, chat_id int64) (deleted bool, err error) {
	l.logger.Infow(module,
		"Action", "DeleteChat",
		"Request", chat_id,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "DeleteChat",
				"Request", chat_id,
				"Error", err,
			)
		}
	}()
	return l.next.DeleteChat(ctx, chat_id)
}

func (l *loggingMiddleware) CheckIfUniqChat(ctx context.Context, productId int64, userId int64) (unique bool, err error) {
	l.logger.Infow(module,
		"Action", "CheckIfUniqChat",
		"Request", productId, userId,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "CheckIfUniqChat",
				"Request", productId, userId,
				"Error", err,
			)
		}
	}()

	return l.next.CheckIfUniqChat(ctx, productId, userId)
}

func (l *loggingMiddleware) CombineChatsWithAndWithoutMsgs(ctx context.Context, onlyChats *models.Chats, chatsAndLM *models.ChatsAndLastMsg) *models.ChatsAndLastMsg {
	l.logger.Infow(module,
		"Action", "CombineChatsWithAndWithoutMsgs",
		"Request", onlyChats, chatsAndLM,
	)
	defer func() {
		l.logger.Infow(module,
			"Action", "CombineChatsWithAndWithoutMsgs",
			"Request", onlyChats, chatsAndLM,
		)
	}()

	return l.next.CombineChatsWithAndWithoutMsgs(ctx, onlyChats, chatsAndLM)
}
func (l *loggingMiddleware) GetUserCreatedChats(ctx context.Context, userId int64) (c *models.Chats, err error) {
	l.logger.Infow(module,
		"Action", "GetUserCreatedChats",
		"Request", userId,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetUserCreatedChats",
				"Request", userId,
				"Error", err,
			)
		}
	}()

	return l.next.GetUserCreatedChats(ctx, userId)
}

func (l *loggingMiddleware) GetAllChatsAndLastMsg(ctx context.Context, userId int64) (c *models.ChatsAndLastMsg, err error) {
	l.logger.Infow(module,
		"Action", "GetAllChatsAndLastMsg",
		"Request", userId,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetAllChatsAndLastMsg",
				"Request", userId,
				"Error", err,
			)
		}
	}()

	return l.next.GetAllChatsAndLastMsg(ctx, userId)
}

func (l *loggingMiddleware) NewChat(ctx context.Context, newChat *models.Chat) (chat *models.Chat, err error) {
	l.logger.Infow(module,
		"Action", "NewChat",
		"Request", newChat,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "NewChat",
				"Request", newChat,
				"Error", err,
			)
		}
	}()

	return l.next.NewChat(ctx, newChat)
}

func (l *loggingMiddleware) GetChat(ctx context.Context, chat *models.Chat) (getChat *models.Chat, err error) {
	l.logger.Infow(module,
		"Action", "GetChat",
		"Request", chat,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetChat",
				"Request", chat,
				"Error", err,
			)
		}
	}()

	return l.next.GetChat(ctx, chat)
}

func (l *loggingMiddleware) GetChatById(ctx context.Context, chatId int64) (chat *models.Chat, err error) {
	l.logger.Infow(module,
		"Action", "GetChatById",
		"Request", chatId,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetChatById",
				"Request", chatId,
				"Error", err,
			)
		}
	}()

	return l.next.GetChatById(ctx, chatId)
}

func (l *loggingMiddleware) WriteNewMsg(ctx context.Context, newMsg *models.Msg) (id int64, err error) {
	l.logger.Infow(module,
		"Action", "WriteNewMsg",
		"Request", newMsg,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "WriteNewMsg",
				"Request", newMsg,
				"Error", err,
			)
		}
	}()

	return l.next.WriteNewMsg(ctx, newMsg)
}

func (l *loggingMiddleware) GetMsgsFromChat(ctx context.Context, chatId int64, userId int64) (msgs *models.Msgs, err error) {
	l.logger.Infow(module,
		"Action", "GetMsgsFromChat",
		"Request", chatId, userId,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetMsgsFromChat",
				"Request", chatId, userId,
				"Error", err,
			)
		}
	}()

	return l.next.GetMsgsFromChat(ctx, chatId, userId)
}

func (l *loggingMiddleware) GetUserLastMsgs(ctx context.Context, userId int64) (msgs *models.Msgs, err error) {
	l.logger.Infow(module,
		"Action", "GetUserLastMsgs",
		"Request", userId,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetUserLastMsgs",
				"Request", userId,
				"Error", err,
			)
		}
	}()

	return l.next.GetUserLastMsgs(ctx, userId)
}
