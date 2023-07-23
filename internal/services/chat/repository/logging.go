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

func (l *loggingMiddleware) GetAllChatsAndLastMsg(ctx context.Context, userId int64) (c *models.ChatsAndLastMsg, err error) {
	l.logger.Infow(module,
		"Action", "CheckIfUniqChat",
		"Request", userId,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "CheckIfUniqChat",
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

func (l *loggingMiddleware) WriteNewMsg(ctx context.Context, newMsg *models.Msg) (err error) {
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

func (l *loggingMiddleware) GetMsgsFromChat(ctx context.Context, chatId int64) (msgs *models.Msgs, err error) {
	l.logger.Infow(module,
		"Action", "GetMsgsFromChat",
		"Request", chatId,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetMsgsFromChat",
				"Request", chatId,
				"Error", err,
			)
		}
	}()

	return l.next.GetMsgsFromChat(ctx, chatId)
}

func (l *loggingMiddleware) GetAllUserChats(ctx context.Context, userId int64) (chats *models.Chats, err error) {
	l.logger.Infow(module,
		"Action", "GetAllUserChats",
		"Request", userId,
	)
	defer func() {
		if err != nil {
			l.logger.Infow(module,
				"Action", "GetAllUserChats",
				"Request", userId,
				"Error", err,
			)
		}
	}()

	return l.next.GetAllUserChats(ctx, userId)
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
