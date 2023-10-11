package usecase

import (
	company_usecase "b2b/m/internal/gateway/company/usecase"
	product_usecase "b2b/m/internal/gateway/productsCategories/usecase"
	auth_usecase "b2b/m/internal/gateway/user/usecase"
	"b2b/m/internal/models"
	"errors"
	"log"

	auth_service "b2b/m/pkg/services/auth"
	chat_service "b2b/m/pkg/services/chat"
	product_service "b2b/m/pkg/services/productsCategories"
	"context"
)

type ChatUsecase interface {
	InitChat(ctx context.Context, userId int64, productId int64) (bool, int64, error)
	CheckIfUniqChat(ctx context.Context, userId int64, productId int64) (bool, error)
	NewChat(ctx context.Context, userId int64, productId int64) (*models.Chat, error)
	GetChat(ctx context.Context, userId int64, productId int64) (*models.Chat, error)
	WriteNewMsg(ctx context.Context, request *models.Msg) error
	GetMsgsFromChat(ctx context.Context, chatId int64, userId int64) (*models.Msgs, error)
	GetAllUserChats(ctx context.Context, userId int64, cookie string) (*models.Chats, error)
	GetAllChatsAndLastMsg(ctx context.Context, userId int64) (*models.ChatsAndLastMsg, error)
}

type chatUsecase struct {
	chatGRPC    chatGRPC
	companyGRPC company_usecase.CompanyGRPC
	productGRPC product_usecase.ProductsCategoriesGRPC
	authGRPC    auth_usecase.AuthGRPC
}

func (u *chatUsecase) GetAllUserChats(ctx context.Context, userId int64, cookie string) (*models.Chats, error) {

	userId_test, err := u.authGRPC.GetUserIdByCookie(ctx, &auth_service.GetUserIdByCookieRequest{
		Hash: cookie,
	})
	log.Println("======== USECASE userId_test:", userId_test)
	if err != nil {
		return nil, err
	}

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
	log.Println("start GetAllChatsAndLastMsg")
	response, err := u.chatGRPC.GetAllChatsAndLastMsg(ctx, &chat_service.IdRequest{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}
	// log.Println("got GetAllChatsAndLastMsg", response)
	// sender, err := u.AuthUsecase.Profile(ctx, response.Chats[0].Msg.SenderId)
	// if err != nil {
	// 	log.Println("ERROR: GetAllChatsAndLastMsg -> GetUserInfo for sender ", err)
	// 	return nil, err
	// }
	// log.Println("------")
	// reciever, err := u.AuthUsecase.Profile(ctx, response.Chats[0].Msg.ReceiverId)
	// if err != nil {
	// 	log.Println("ERROR: GetAllChatsAndLastMsg -> GetUserInfo for reciever ", err)
	// 	return nil, err
	// }
	// log.Println("got sender and reciever", sender)

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
			Id:           result.Msg.Id,
			ChatId:       result.Msg.ChatId,
			SenderId:     result.Msg.SenderId,
			ReceiverId:   result.Msg.ReceiverId,
			SenderName:   result.Msg.SenderName,
			ReceiverName: result.Msg.ReceiverName,
			Checked:      result.Msg.Checked,
			Text:         result.Msg.Text,
			Type:         result.Msg.Type,
			Time:         result.Msg.Time,
		}
		chatAndLastMsg.Chat = chat
		chatAndLastMsg.LastMsg = msg
		chatsAndLastMsg = append(chatsAndLastMsg, chatAndLastMsg)
	}
	return &chatsAndLastMsg, nil
}

func (u *chatUsecase) GetMsgsFromChat(ctx context.Context, chatId int64, userId int64) (*models.Msgs, error) {
	log.Println("start GetMsgsFromChat")
	response, err := u.chatGRPC.GetMsgsFromChat(ctx, &chat_service.ChatAndUserIdRequest{
		ChatId: chatId,
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	//log.Println("---Get names---")
	// log.Println(response.Msgs[0].SenderId)
	// log.Println(response.Msgs[0].ReceiverId)
	var msg models.Msg
	var msgs models.Msgs
	log.Println("---names---")
	if len(response.Msgs) > 0 {
		sender, err := u.authGRPC.GetUser(ctx, &auth_service.GetUserRequest{Id: response.Msgs[0].SenderId})
		if err != nil {
			log.Println("ERROR: GetMsgsFromChat -> GetUserInfo for sender ", err)
			return nil, err
		}
		log.Println("------")
		reciever, err := u.authGRPC.GetUser(ctx, &auth_service.GetUserRequest{Id: response.Msgs[0].ReceiverId})
		if err != nil {
			log.Println("ERROR: GetMsgsFromChat -> GetUserInfo for reciever ", err)
			return nil, err
		}
		log.Println(sender, reciever)
		for _, result := range response.Msgs {
			msg = models.Msg{
				Id:           result.Id,
				ChatId:       result.ChatId,
				SenderId:     result.SenderId,
				ReceiverId:   result.ReceiverId,
				SenderName:   sender.Name,
				ReceiverName: reciever.Name,
				Checked:      result.Checked,
				Text:         result.Text,
				Type:         result.Type,
				Time:         result.Time,
			}
			msgs = append(msgs, msg)
		}
		return &msgs, nil
	} else {
		log.Println("No msgs in chat")
		return &msgs, errors.New("No msgs in chat")
	}
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
		Status:    response.Status,
	}, nil
}

func (u *chatUsecase) InitChat(ctx context.Context, userId int64, productId int64) (bool, int64, error) {
	response, err := u.chatGRPC.CheckIfUniqChat(ctx, &chat_service.CheckIfUniqChatRequest{
		UserId:    userId,
		ProductId: productId,
	})
	if err != nil {
		return false, -1, err
	}
	chat := &models.Chat{}
	if response.Unique {
		chat, err = u.NewChat(ctx, userId, productId)
		if err != nil {
			return false, -1, err
		}
	} else {
		chat, err = u.GetChat(ctx, userId, productId)
		log.Println("CHATTTTTT", chat)
		if err != nil {
			return false, -1, err
		}
	}

	return response.Unique, chat.Id, nil
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

func (u *chatUsecase) GetChat(ctx context.Context, userId int64, productId int64) (*models.Chat, error) {
	response, err := u.chatGRPC.GetChat(ctx, &chat_service.GetChatRequest{
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
		Status:    response.Status,
	}, nil
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
	log.Println("WriteNewMsg:", request)
	if err != nil {
		return err
	}
	return nil
}

func NewChatUsecase(chatGRPC chatGRPC, companyGRPC company_usecase.CompanyGRPC, productGRPC product_usecase.ProductsCategoriesGRPC, authGRPC auth_usecase.AuthGRPC) ChatUsecase {
	return &chatUsecase{chatGRPC: chatGRPC, companyGRPC: companyGRPC, productGRPC: productGRPC, authGRPC: authGRPC}
}
