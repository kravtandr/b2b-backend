package usecase

import (
	company_usecase "b2b/m/internal/gateway/company/usecase"
	product_usecase "b2b/m/internal/gateway/productsCategories/usecase"
	auth_usecase "b2b/m/internal/gateway/user/usecase"
	"b2b/m/internal/models"

	chat_service "b2b/m/pkg/services/chat"
	product_service "b2b/m/pkg/services/productsCategories"
	"context"
)

type ChatUsecase interface {
	CheckIfUniqChat(ctx context.Context, userId int64, productId int64) (bool, error)
	NewChat(ctx context.Context, userId int64, productId int64) (*models.Chat, error)
	WriteNewMsg(ctx context.Context, request *models.Msg) error
	GetMsgsFromChat(ctx context.Context, chatId int64) (*models.Msgs, error)
	GetAllUserChats(ctx context.Context, userId int64) (*models.Chats, error)
	GetAllChatsAndLastMsg(ctx context.Context, userId int64) (*models.ChatsAndLastMsg, error)
}

type chatUsecase struct {
	chatGRPC    chatGRPC
	companyGRPC company_usecase.CompanyGRPC
	productGRPC product_usecase.ProductsCategoriesGRPC
	AuthGRPC    auth_usecase.AuthGRPC
}

func (u *chatUsecase) GetAllUserChats(ctx context.Context, userId int64) (*models.Chats, error) {
	response, err := u.chatGRPC.GetAllUserChats(ctx, &chat_service.IdRequest{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}

	var chat models.Chat
	var chats models.Chats
	for _, result := range response.Chats {
		chat = models.Chat{
			Id:        result.Id,
			Name:      result.Name,
			CreatorId: result.CreatorId,
			ProductId: result.ProductId,
		}
		chats = append(chats, chat)
	}
	return &chats, nil
}

func (u *chatUsecase) GetAllChatsAndLastMsg(ctx context.Context, userId int64) (*models.ChatsAndLastMsg, error) {
	//userId, err := u.AuthGRPC.GetUserIdByCookie(ctx)
	//userId := ctx.UserValue(cnst.UserIDContextKey).(int)
	response, err := u.chatGRPC.GetAllChatsAndLastMsg(ctx, &chat_service.IdRequest{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}

	var chat models.Chat
	var msg models.Msg
	var chatAndLastMsg models.ChatAndLastMsg
	var chatsAndLastMsg models.ChatsAndLastMsg
	for _, result := range response.Chats {
		chat = models.Chat{
			Id:        result.Id,
			Name:      result.Name,
			CreatorId: result.CreatorId,
			ProductId: result.ProductId,
		}
		msg = models.Msg{
			Id:         result.Msg.Id,
			ChatId:     result.Msg.ChatId,
			SenderId:   result.Msg.SenderId,
			ReceiverId: result.Msg.ReceiverId,
			Checked:    result.Msg.Checked,
			Text:       result.Msg.Text,
			Type:       result.Msg.Type,
			Time:       result.Msg.Time,
		}
		chatAndLastMsg.Chat = chat
		chatAndLastMsg.LastMsg = msg
		chatsAndLastMsg = append(chatsAndLastMsg, chatAndLastMsg)
	}
	return &chatsAndLastMsg, nil
}

func (u *chatUsecase) GetMsgsFromChat(ctx context.Context, chatId int64) (*models.Msgs, error) {
	response, err := u.chatGRPC.GetMsgsFromChat(ctx, &chat_service.IdRequest{
		Id: chatId,
	})
	if err != nil {
		return nil, err
	}

	var msg models.Msg
	var msgs models.Msgs
	for _, result := range response.Msgs {
		msg = models.Msg{
			Id:      result.Id,
			ChatId:  result.ChatId,
			Checked: result.Checked,
			Text:    result.Text,
			Type:    result.Type,
			Time:    result.Time,
		}
		msgs = append(msgs, msg)
	}
	return &msgs, nil
}

func (u *chatUsecase) NewChat(ctx context.Context, userId int64, productId int64) (*models.Chat, error) {
	product, err := u.productGRPC.GetProductById(ctx, &product_service.GetProductByID{
		Id: productId,
	})
	response, err := u.chatGRPC.NewChat(ctx, &chat_service.NewChatRequest{
		Name:      product.Name,
		CreatorId: userId,
		ProductId: productId,
	})
	if err != nil {
		return nil, err
	}

	return &models.Chat{
		Id:        response.Id,
		Name:      response.Name,
		CreatorId: response.CreatorId,
		ProductId: response.ProductId,
	}, nil
}

func (u *chatUsecase) CheckIfUniqChat(ctx context.Context, userId int64, productId int64) (bool, error) {
	response, err := u.chatGRPC.CheckIfUniqChat(ctx, &chat_service.CheckIfUniqChatRequest{
		UserId:    userId,
		ProductId: productId,
	})
	if err != nil {
		return false, err
	}

	return response.Unique, nil
}

func (u *chatUsecase) WriteNewMsg(ctx context.Context, request *models.Msg) error {
	_, err := u.chatGRPC.WriteNewMsg(ctx, &chat_service.WriteNewMsgRequest{
		ChatId:     request.ChatId,
		SenderId:   request.SenderId,
		ReceiverId: request.ReceiverId,
		Checked:    request.Checked,
		Text:       request.Text,
		Type:       request.Type,
		Time:       request.Time,
	})
	if err != nil {
		return err
	}
	return nil
}

func NewChatUsecase(chatGRPC chatGRPC, companyGRPC company_usecase.CompanyGRPC, productGRPC product_usecase.ProductsCategoriesGRPC) ChatUsecase {
	return &chatUsecase{chatGRPC: chatGRPC, companyGRPC: companyGRPC, productGRPC: productGRPC}
}
