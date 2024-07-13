package usecase

import (
	"context"

	auth_service "b2b/m/pkg/services/auth"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthGRPC interface {
	GetUserInfo(ctx context.Context, in *auth_service.GetUserRequest, opts ...grpc.CallOption) (*auth_service.UserInfo, error)
	LogoutUser(ctx context.Context, in *auth_service.Session, opts ...grpc.CallOption) (*emptypb.Empty, error)
	LoginUser(ctx context.Context, in *auth_service.LoginRequest, opts ...grpc.CallOption) (*auth_service.LoginResponse, error)
	RegisterUser(ctx context.Context, in *auth_service.RegisterRequest, opts ...grpc.CallOption) (*auth_service.RegisterResponse, error)
	GetUser(ctx context.Context, in *auth_service.GetUserRequest, opts ...grpc.CallOption) (*auth_service.GetUserResponse, error)
	UpdateUser(ctx context.Context, in *auth_service.UpdateUserRequest, opts ...grpc.CallOption) (*auth_service.GetPublicUserResponse, error)
	FastRegister(ctx context.Context, in *auth_service.FastRegisterRequest, opts ...grpc.CallOption) (*auth_service.LoginResponse, error)
	GetUserIdByCookie(ctx context.Context, in *auth_service.GetUserIdByCookieRequest, opts ...grpc.CallOption) (*auth_service.UserId, error)
	CheckEmail(ctx context.Context, in *auth_service.CheckEmailRequest, opts ...grpc.CallOption) (*auth_service.GetPublicUserResponse, error)
	GetUsersCompany(ctx context.Context, in *auth_service.UserIdRequest, opts ...grpc.CallOption) (*auth_service.GetPrivateCompanyResponse, error)
	GetCompanyUserLink(ctx context.Context, in *auth_service.UserAndCompanyIdsRequest, opts ...grpc.CallOption) (*auth_service.GetCompanyUserLinkResponse, error)
	UpdateUserBalance(ctx context.Context, in *auth_service.UpdateUserBalanceRequest, opts ...grpc.CallOption) (*auth_service.GetPublicUserResponse, error)
	AddPayment(ctx context.Context, in *auth_service.AddPaymentRequest, opts ...grpc.CallOption) (*auth_service.PaymentResponse, error)
	UpdatePayment(ctx context.Context, in *auth_service.UpdatePaymentRequest, opts ...grpc.CallOption) (*auth_service.PaymentResponse, error)
	GetPayment(ctx context.Context, in *auth_service.GetPaymentRequest, opts ...grpc.CallOption) (*auth_service.PaymentResponse, error)
	GetUsersPayments(ctx context.Context, in *auth_service.UserIdRequest, opts ...grpc.CallOption) (*auth_service.PaymentsResponse, error)
	HandlePaidPayments(ctx context.Context, in *auth_service.HandlePaidPaymentsRequest, opts ...grpc.CallOption) (*auth_service.HandlePaidPaymentsResponse, error)
}
