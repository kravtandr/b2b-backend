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
	company_service "b2b/m/pkg/services/company"
	product_service "b2b/m/pkg/services/productsCategories"
	"context"
)

var NEW_CHAT_PRICE int64 = 100

type ChatUsecase interface {
	InitChat(ctx context.Context, userId int64, productId int64) (bool, int64, error)
	UpdateChatStatus(ctx context.Context, chatId int64, status string, blured bool) (*models.Chat, error)
	DeleteChat(ctx context.Context, request *models.DeleteChatRequest) (bool, error)
	CheckIfUniqChat(ctx context.Context, userId int64, productId int64) (bool, error)
	NewChat(ctx context.Context, userId int64, productId int64) (*models.Chat, error)
	GetChat(ctx context.Context, userId int64, productId int64) (*models.Chat, error)
	GetMsgsFromChat(ctx context.Context, chatId int64, userId int64) (*models.Msgs, error)
	GetAllUserChats(ctx context.Context, userId int64, cookie string) (*models.Chats, error)
	GetAllChatsAndLastMsg(ctx context.Context, userId int64) (*models.ChatsAndLastMsg, error)
	CheckAndUnblurChats(ctx context.Context, userId int64) error
	CheckOwnerBalanceAndReduce(ctx context.Context, productId int64) (bool, error)
	BlurChatLastMsg(ctx context.Context, chat models.ChatAndLastMsg) models.ChatAndLastMsg
	BlurChatMsg(ctx context.Context, msg models.Msg) models.Msg
}

type chatUsecase struct {
	chatGRPC    chatGRPC
	companyGRPC company_usecase.CompanyGRPC
	productGRPC product_usecase.ProductsCategoriesGRPC
	authGRPC    auth_usecase.AuthGRPC
	authUsecase auth_usecase.UserUsecase
}

func (u *chatUsecase) GetAllUserChats(ctx context.Context, userId int64, cookie string) (*models.Chats, error) {

	_, err := u.authGRPC.GetUserIdByCookie(ctx, &auth_service.GetUserIdByCookieRequest{
		Hash: cookie,
	})
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

	// check if user pays and have chats to unblur
	err := u.CheckAndUnblurChats(ctx, userId)
	if err != nil {
		return nil, err
	}

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

	if len(response.Chats) > 0 {
		// TODO blur chats that created after unsufficient balance
		for _, result := range response.Chats {
			chat = models.Chat{
				Id:        result.Id,
				Name:      result.Name,
				CreatorId: result.CreatorId,
				ProductId: result.ProductId,
				Status:    result.Status,
				Blured:    result.Blured,
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
			company_response, err := u.companyGRPC.GetCompanyByProductId(ctx, &company_service.IdRequest{
				Id: result.ProductId,
			})
			if err != nil {
				return nil, err
			}
			company := models.Company{
				Id:           company_response.Id,
				Name:         company_response.Name,
				Description:  company_response.Description,
				LegalName:    company_response.LegalName,
				Itn:          company_response.Itn,
				Psrn:         company_response.Psrn,
				Address:      company_response.Address,
				LegalAddress: company_response.LegalAddress,
				Email:        company_response.Email,
				Phone:        company_response.Phone,
				Link:         company_response.Link,
				Activity:     company_response.Activity,
				OwnerId:      company_response.OwnerId,
				Rating:       company_response.Rating,
				Verified:     company_response.Verified,
				Photo:        company_response.Photo,
			}
			chatAndLastMsg.Chat = chat
			chatAndLastMsg.LastMsg = msg
			chatAndLastMsg.Company = company
			if chatAndLastMsg.Chat.Blured {
				chatAndLastMsg = u.BlurChatLastMsg(ctx, chatAndLastMsg)
			}
			chatsAndLastMsg = append(chatsAndLastMsg, chatAndLastMsg)
		}
		return &chatsAndLastMsg, nil

	} else {
		return &chatsAndLastMsg, nil
	}

}

func (u *chatUsecase) BlurChatLastMsg(ctx context.Context, chat models.ChatAndLastMsg) models.ChatAndLastMsg {
	chat.LastMsg.Text = "Last msg blurred"
	return chat
}

func (u *chatUsecase) BlurChatMsg(ctx context.Context, msg models.Msg) models.Msg {
	msg.Text = "Msg blurred"
	return msg
}

func (u *chatUsecase) CheckOwnerBalanceAndReduce(ctx context.Context, productId int64) (bool, error) {

	owner_company, err := u.companyGRPC.GetCompanyByProductId(ctx, &company_service.IdRequest{
		Id: productId,
	})
	if err != nil {
		return false, err
	}

	owner, err := u.authUsecase.Profile(ctx, owner_company.OwnerId)
	if err != nil {
		return false, err
	}

	if owner.Balance-NEW_CHAT_PRICE < 0 {
		log.Println("Not enough balance, user: ", owner.Id, " balance: ", owner.Balance)
		return false, nil
	} else {
		owner.Balance = owner.Balance - NEW_CHAT_PRICE
		_, err = u.authUsecase.UpdateUserBalance(ctx, owner.Id, owner.Balance)
		if err != nil {
			return false, err
		}
		log.Println("Payed for new chat,  user: ", owner.Id, " balance: ", owner.Balance)
		return true, nil
	}
}

func (u *chatUsecase) CheckAndUnblurChats(ctx context.Context, userId int64) error {
	// get all chats and change blur if sufficient balance

	response, err := u.chatGRPC.GetAllUserChats(ctx, &chat_service.IdRequest{
		Id: userId,
	})
	if err != nil {
		log.Println(err)
		return err
	}
	var chat models.Chat
	var chats models.Chats
	for _, result := range response.Chats {
		chat = models.Chat{
			Id:        result.Id,
			Name:      result.Name,
			CreatorId: result.CreatorId,
			ProductId: result.ProductId,
			Status:    result.Status,
			Blured:    result.Blured,
		}
		chats = append(chats, chat)
	}

	if len(chats) > 0 {
		for _, chat := range chats {
			if chat.Blured {
				isSufficientBalance, err := u.CheckOwnerBalanceAndReduce(ctx, chat.ProductId)
				if err != nil {
					log.Println(err)
					return err
				}

				if isSufficientBalance {
					_, err := u.chatGRPC.UpdateChatStatus(ctx, &chat_service.UpdateChatStatusRequest{
						ChatId: chat.Id,
						Status: chat.Status,
						Blured: false,
					})
					if err != nil {
						_, err := u.authUsecase.AddUserBalance(ctx, userId, NEW_CHAT_PRICE)
						if err != nil {
							return err
						}
						return err
					}
					log.Println("Unblurred chat: ", chat.Id)
				}
			} else {
				log.Println("Chat already unblurred: ", chat.Id, " Blured: ", chat.Blured)
			}
		}
		return nil
	} else {
		log.Println("No chats to unblur")
		return nil
	}

}

func (u *chatUsecase) GetMsgsFromChat(ctx context.Context, chatId int64, userId int64) (*models.Msgs, error) {
	// TODO check blur

	// TODO check if user pays and have chats to unblur

	response, err := u.chatGRPC.GetMsgsFromChat(ctx, &chat_service.ChatAndUserIdRequest{
		ChatId: chatId,
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	chat, err := u.chatGRPC.GetChatById(ctx, &chat_service.IdRequest{
		Id: chatId,
	})
	if err != nil {
		return nil, err
	}

	var msg models.Msg
	var msgs models.Msgs
	if len(response.Msgs) > 0 {
		sender, err := u.authGRPC.GetUser(ctx, &auth_service.GetUserRequest{Id: response.Msgs[0].SenderId})
		if err != nil {
			log.Println("ERROR: GetMsgsFromChat -> GetUserInfo for sender ", err)
			return nil, err
		}
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
			if chat.Blured {
				msg = u.BlurChatMsg(ctx, msg)
			}
			msgs = append(msgs, msg)
		}

		return &msgs, nil
	} else {
		log.Println("No msgs in chat")
		return &msgs, errors.New("no msgs in chat")
	}
}

func (u *chatUsecase) NewChat(ctx context.Context, userId int64, productId int64) (*models.Chat, error) {

	isSufficientBalance, err := u.CheckOwnerBalanceAndReduce(ctx, productId)
	if err != nil {
		return nil, err
	}

	var newChatReq chat_service.NewChatRequest

	product, err := u.productGRPC.GetProductById(ctx, &product_service.GetProductByID{
		Id: productId,
	})
	if err != nil {
		_, err := u.authUsecase.AddUserBalance(ctx, userId, NEW_CHAT_PRICE)
		if err != nil {
			return nil, err
		}

		return nil, err
	}

	if !isSufficientBalance {
		newChatReq = chat_service.NewChatRequest{
			Name:      product.Name,
			CreatorId: userId,
			ProductId: productId,
			Blured:    true,
		}
	} else {
		newChatReq = chat_service.NewChatRequest{
			Name:      product.Name,
			CreatorId: userId,
			ProductId: productId,
			Blured:    false,
		}
	}
	response, err := u.chatGRPC.NewChat(ctx, &newChatReq)
	if err != nil {
		_, err := u.authUsecase.AddUserBalance(ctx, userId, NEW_CHAT_PRICE)
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	return &models.Chat{
		Id:        response.Id,
		Name:      response.Name,
		CreatorId: response.CreatorId,
		ProductId: response.ProductId,
		Status:    response.Status,
		Blured:    response.Blured,
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
		if err != nil {
			return false, -1, err
		}
	}

	return response.Unique, chat.Id, nil
}

func (u *chatUsecase) UpdateChatStatus(ctx context.Context, chatId int64, status string, blured bool) (*models.Chat, error) {
	response, err := u.chatGRPC.UpdateChatStatus(ctx, &chat_service.UpdateChatStatusRequest{
		ChatId: chatId,
		Status: status,
		Blured: blured,
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
		Blured:    response.Blured,
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
		Blured:    response.Blured,
	}, nil
}
func (u *chatUsecase) DeleteChat(ctx context.Context, request *models.DeleteChatRequest) (bool, error) {
	_, err := u.chatGRPC.DeleteChat(ctx, &chat_service.IdRequest{
		Id: request.Chat_id,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewChatUsecase(chatGRPC chatGRPC, companyGRPC company_usecase.CompanyGRPC, productGRPC product_usecase.ProductsCategoriesGRPC, authGRPC auth_usecase.AuthGRPC, authUsecase auth_usecase.UserUsecase) ChatUsecase {
	return &chatUsecase{chatGRPC: chatGRPC, companyGRPC: companyGRPC, productGRPC: productGRPC, authGRPC: authGRPC, authUsecase: authUsecase}
}
