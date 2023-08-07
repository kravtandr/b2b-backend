package usecase

import (
	"b2b/m/internal/services/chat/models"
	"b2b/m/pkg/generator"
	"context"

	"github.com/gofrs/uuid"
)

type ChatUseCase interface {
	CheckIfUniqChat(ctx context.Context, user *models.UniqueCheck) (bool, error)
	NewChat(ctx context.Context, newChat *models.Chat) (*models.Chat, error)
	GetChat(ctx context.Context, chat *models.Chat) (*models.Chat, error)
	WriteNewMsg(ctx context.Context, newMsg *models.Msg) error
	GetMsgsFromChat(ctx context.Context, chatId int64, userId int64) (*models.Msgs, error)
	GetAllChatsAndLastMsg(ctx context.Context, userId int64) (*models.ChatsAndLastMsg, error)
	GetAllUserChats(ctx context.Context, userId int64) (*models.Chats, error)
}

type chatUseCase struct {
	hashGenerator hasher
	repo          chatRepository
	uuidGen       generator.UUIDGenerator
}

func (a *chatUseCase) CheckIfUniqChat(ctx context.Context, uniqueCheck *models.UniqueCheck) (bool, error) {
	result, err := a.repo.CheckIfUniqChat(ctx, uniqueCheck.ProductId, uniqueCheck.UserId)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (a *chatUseCase) NewChat(ctx context.Context, newChat *models.Chat) (*models.Chat, error) {
	result, err := a.repo.NewChat(ctx, newChat)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *chatUseCase) GetChat(ctx context.Context, chat *models.Chat) (*models.Chat, error) {
	result, err := a.repo.GetChat(ctx, chat)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *chatUseCase) WriteNewMsg(ctx context.Context, newMsg *models.Msg) error {
	err := a.repo.WriteNewMsg(ctx, newMsg)
	if err != nil {
		return err
	}
	return nil
}

func (a *chatUseCase) GetMsgsFromChat(ctx context.Context, chatId int64, userId int64) (*models.Msgs, error) {
	result, err := a.repo.GetMsgsFromChat(ctx, chatId, userId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *chatUseCase) GetAllChatsAndLastMsg(ctx context.Context, userId int64) (*models.ChatsAndLastMsg, error) {
	result, err := a.repo.GetAllChatsAndLastMsg(ctx, userId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *chatUseCase) GetAllUserChats(ctx context.Context, userId int64) (*models.Chats, error) {
	result, err := a.repo.GetAllUserChats(ctx, userId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func NewChatUseCase(
	hashGenerator hasher,
	repo chatRepository,
) ChatUseCase {
	return &chatUseCase{
		hashGenerator: hashGenerator,
		repo:          repo,
		uuidGen:       generator.NewUUIDGenerator(uuid.NewGen()),
	}
}
