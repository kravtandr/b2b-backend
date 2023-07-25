package delivery

import (
	"b2b/m/internal/services/chat/models"
	"b2b/m/internal/services/chat/usecase"
	"b2b/m/pkg/error_adapter"
	chat_service "b2b/m/pkg/services/chat"
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
)

type chatDelivery struct {
	chatUsecase  usecase.ChatUseCase
	errorAdapter error_adapter.ErrorAdapter
	chat_service.UnimplementedChatServiceServer
}

func (a *chatDelivery) CheckIfUniqChat(ctx context.Context, request *chat_service.CheckIfUniqChatRequest) (*chat_service.CheckIfUniqChatResponse, error) {
	response, err := a.chatUsecase.CheckIfUniqChat(ctx, &models.UniqueCheck{
		UserId:    request.UserId,
		ProductId: request.ProductId,
	})
	if err != nil {
		return nil, a.errorAdapter.AdaptError(err)
	}

	return &chat_service.CheckIfUniqChatResponse{
		Unique: response,
	}, nil
}

func (a *chatDelivery) NewChat(ctx context.Context, request *chat_service.NewChatRequest) (*chat_service.ChatResponse, error) {
	response, err := a.chatUsecase.NewChat(ctx, &models.Chat{
		Name:      request.Name,
		CreatorId: request.CreatorId,
		ProductId: request.ProductId,
	})
	if err != nil {
		return nil, a.errorAdapter.AdaptError(err)
	}

	return &chat_service.ChatResponse{
		Id:        response.Id,
		Name:      response.Name,
		CreatorId: response.CreatorId,
		ProductId: response.ProductId,
	}, nil
}

func (a *chatDelivery) GetChat(ctx context.Context, request *chat_service.GetChatRequest) (*chat_service.ChatResponse, error) {
	response, err := a.chatUsecase.GetChat(ctx, &models.Chat{
		CreatorId: request.CreatorId,
		ProductId: request.ProductId,
	})
	if err != nil {
		return nil, a.errorAdapter.AdaptError(err)
	}

	return &chat_service.ChatResponse{
		Id:        response.Id,
		Name:      response.Name,
		CreatorId: response.CreatorId,
		ProductId: response.ProductId,
	}, nil
}

func (a *chatDelivery) WriteNewMsg(ctx context.Context, request *chat_service.WriteNewMsgRequest) (*empty.Empty, error) {
	err := a.chatUsecase.WriteNewMsg(ctx, &models.Msg{
		ChatId:     request.ChatId,
		SenderId:   request.SenderId,
		ReceiverId: request.ReceiverId,
		Text:       request.Text,
		Type:       request.Type,
	})
	if err != nil {
		return &empty.Empty{}, a.errorAdapter.AdaptError(err)
	}

	return &empty.Empty{}, nil
}

func (a *chatDelivery) GetMsgsFromChat(ctx context.Context, request *chat_service.ChatAndUserIdRequest) (*chat_service.MsgsResponse, error) {
	resp, err := a.chatUsecase.GetMsgsFromChat(ctx, request.ChatId, request.UserId)
	if err != nil {
		return nil, a.errorAdapter.AdaptError(err)
	}
	var res chat_service.MsgsResponse
	var msg *chat_service.MsgResponse

	for _, result := range *resp {
		msg = &chat_service.MsgResponse{
			Id:      result.Id,
			ChatId:  result.ChatId,
			Checked: result.Checked,
			Text:    result.Text,
			Type:    result.Type,
			Time:    fmt.Sprint(result.Time),
		}
		res.Msgs = append(res.Msgs, msg)

	}
	return &res, nil
}

func (a *chatDelivery) GetAllChatsAndLastMsg(ctx context.Context, request *chat_service.IdRequest) (*chat_service.GetAllChatsAndLastMsgResponse, error) {
	chatsAndLastMsg, err := a.chatUsecase.GetAllChatsAndLastMsg(ctx, request.Id)
	if err != nil {
		return nil, a.errorAdapter.AdaptError(err)
	}
	//var msg chat_service.MsgResponse
	var res chat_service.GetAllChatsAndLastMsgResponse
	var chatAndLatsMsg *chat_service.ChatAndLastMsgResponse

	for _, item := range *chatsAndLastMsg {
		chatAndLatsMsg = &chat_service.ChatAndLastMsgResponse{
			Id:        item.Chat.Id,
			Name:      item.Chat.Name,
			CreatorId: item.Chat.CreatorId,
			ProductId: item.Chat.ProductId,
			Msg: &chat_service.MsgResponse{
				Id:         item.LastMsg.Id,
				ChatId:     item.LastMsg.ChatId,
				SenderId:   item.LastMsg.SenderId,
				ReceiverId: item.LastMsg.ReceiverId,
				Checked:    item.LastMsg.Checked,
				Text:       item.LastMsg.Text,
				Type:       item.LastMsg.Type,
				Time:       fmt.Sprint(item.LastMsg.Time),
			},
		}
		res.Chats = append(res.Chats, chatAndLatsMsg)

	}
	return &res, nil
}

func (a *chatDelivery) GetAllUserChats(ctx context.Context, request *chat_service.IdRequest) (*chat_service.GetAllUserChatsResponse, error) {
	chats, err := a.chatUsecase.GetAllUserChats(ctx, request.Id)
	if err != nil {
		return nil, a.errorAdapter.AdaptError(err)
	}
	var res chat_service.GetAllUserChatsResponse
	var chat *chat_service.ChatResponse

	for _, item := range *chats {
		chat = &chat_service.ChatResponse{
			Id:        item.Id,
			Name:      item.Name,
			CreatorId: item.CreatorId,
			ProductId: item.ProductId,
		}
		res.Chats = append(res.Chats, chat)

	}
	return &res, nil
}

func NewChatDelivery(
	chatUsecase usecase.ChatUseCase,
	errorAdapter error_adapter.ErrorAdapter,
) chat_service.ChatServiceServer {
	return &chatDelivery{
		chatUsecase:  chatUsecase,
		errorAdapter: errorAdapter,
	}
}
