package usecase

import (
	"context"

	chat_service "b2b/m/pkg/services/chat"

	"google.golang.org/grpc"
)

type chatGRPC interface {
	NewChat(ctx context.Context, in *chat_service.NewChatRequest, opts ...grpc.CallOption) (*chat_service.ChatResponse, error)
	GetChatById(ctx context.Context, in *chat_service.IdRequest, opts ...grpc.CallOption) (*chat_service.ChatResponse, error)
	GetChat(ctx context.Context, in *chat_service.GetChatRequest, opts ...grpc.CallOption) (*chat_service.ChatResponse, error)
	GetAllUserChats(ctx context.Context, in *chat_service.IdRequest, opts ...grpc.CallOption) (*chat_service.GetAllUserChatsResponse, error)
	UpdateChatStatus(ctx context.Context, in *chat_service.UpdateChatStatusRequest, opts ...grpc.CallOption) (*chat_service.ChatResponse, error)
	DeleteChat(ctx context.Context, in *chat_service.IdRequest, opts ...grpc.CallOption) (*chat_service.Bool, error)

	GetMsgsFromChat(ctx context.Context, in *chat_service.ChatAndUserIdRequest, opts ...grpc.CallOption) (*chat_service.MsgsResponse, error)
	GetAllChatsAndLastMsg(ctx context.Context, in *chat_service.IdRequest, opts ...grpc.CallOption) (*chat_service.GetAllChatsAndLastMsgResponse, error)

	CheckIfUniqChat(ctx context.Context, in *chat_service.CheckIfUniqChatRequest, opts ...grpc.CallOption) (*chat_service.CheckIfUniqChatResponse, error)
}
