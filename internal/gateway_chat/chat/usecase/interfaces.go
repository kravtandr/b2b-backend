package usecase

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	chat_service "b2b/m/pkg/services/chat"

	"google.golang.org/grpc"
)

type chatGRPC interface {
	CheckIfUniqChat(ctx context.Context, in *chat_service.CheckIfUniqChatRequest, opts ...grpc.CallOption) (*chat_service.CheckIfUniqChatResponse, error)
	NewChat(ctx context.Context, in *chat_service.NewChatRequest, opts ...grpc.CallOption) (*chat_service.ChatResponse, error)
	GetChat(ctx context.Context, in *chat_service.GetChatRequest, opts ...grpc.CallOption) (*chat_service.ChatResponse, error)
	WriteNewMsg(ctx context.Context, in *chat_service.WriteNewMsgRequest, opts ...grpc.CallOption) (*chat_service.IdResponse, error)
	GetMsgsFromChat(ctx context.Context, in *chat_service.ChatAndUserIdRequest, opts ...grpc.CallOption) (*chat_service.MsgsResponse, error)
	GetAllUserChats(ctx context.Context, in *chat_service.IdRequest, opts ...grpc.CallOption) (*chat_service.GetAllUserChatsResponse, error)
	GetAllChatsAndLastMsg(ctx context.Context, in *chat_service.IdRequest, opts ...grpc.CallOption) (*chat_service.GetAllChatsAndLastMsgResponse, error)
	ChatHealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}
