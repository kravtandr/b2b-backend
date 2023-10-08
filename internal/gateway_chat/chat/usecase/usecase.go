package usecase

import (
	company_usecase "b2b/m/internal/gateway/company/usecase"
	product_usecase "b2b/m/internal/gateway/productsCategories/usecase"
	auth_usecase "b2b/m/internal/gateway/user/usecase"
	"b2b/m/internal/models"
	"log"

	chat_service "b2b/m/pkg/services/chat"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ChatUsecase interface {
	WriteNewMsg(ctx context.Context, request *models.Msg) error
	ChatHealthCheck(ctx context.Context) error
}

type chatUsecase struct {
	chatGRPC    chatGRPC
	companyGRPC company_usecase.CompanyGRPC
	productGRPC product_usecase.ProductsCategoriesGRPC
	AuthGRPC    auth_usecase.AuthGRPC
}

func (u *chatUsecase) ChatHealthCheck(ctx context.Context) error {
	_, err := u.chatGRPC.ChatHealthCheck(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}
	return nil
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

func NewChatUsecase(chatGRPC chatGRPC, companyGRPC company_usecase.CompanyGRPC, productGRPC product_usecase.ProductsCategoriesGRPC) ChatUsecase {
	return &chatUsecase{chatGRPC: chatGRPC, companyGRPC: companyGRPC, productGRPC: productGRPC}
}
