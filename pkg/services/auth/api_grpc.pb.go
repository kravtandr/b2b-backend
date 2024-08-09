// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: pkg/services/auth/api.proto

package auth_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	ValidateSession(ctx context.Context, in *Session, opts ...grpc.CallOption) (*ValidateSessionResponse, error)
	LogoutUser(ctx context.Context, in *Session, opts ...grpc.CallOption) (*emptypb.Empty, error)
	LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	FastRegister(ctx context.Context, in *FastRegisterRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	RegisterUser(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*GetPublicUserResponse, error)
	GetUserInfo(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserInfo, error)
	GetUserByEmail(ctx context.Context, in *UserEmailRequest, opts ...grpc.CallOption) (*UserId, error)
	GetUserIdByCookie(ctx context.Context, in *GetUserIdByCookieRequest, opts ...grpc.CallOption) (*UserId, error)
	CheckEmail(ctx context.Context, in *CheckEmailRequest, opts ...grpc.CallOption) (*GetPublicUserResponse, error)
	GetUsersCompany(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*GetPrivateCompanyResponse, error)
	GetCompanyUserLink(ctx context.Context, in *UserAndCompanyIdsRequest, opts ...grpc.CallOption) (*GetCompanyUserLinkResponse, error)
	UpdateUserBalance(ctx context.Context, in *UpdateUserBalanceRequest, opts ...grpc.CallOption) (*GetPublicUserResponse, error)
	AddPayment(ctx context.Context, in *AddPaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error)
	UpdatePayment(ctx context.Context, in *UpdatePaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error)
	GetUsersPayments(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*PaymentsResponse, error)
	GetPayment(ctx context.Context, in *GetPaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error)
	HandlePaidPayments(ctx context.Context, in *HandlePaidPaymentsRequest, opts ...grpc.CallOption) (*HandlePaidPaymentsResponse, error)
	CountUsersPayments(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*PaymentsAmountResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) ValidateSession(ctx context.Context, in *Session, opts ...grpc.CallOption) (*ValidateSessionResponse, error) {
	out := new(ValidateSessionResponse)
	err := c.cc.Invoke(ctx, "/services.auth_service.AuthService/ValidateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) LogoutUser(ctx context.Context, in *Session, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/services.auth_service.AuthService/LogoutUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/services.auth_service.AuthService/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) FastRegister(ctx context.Context, in *FastRegisterRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/services.auth_service.AuthService/FastRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RegisterUser(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/services.auth_service.AuthService/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/services.auth_service.AuthService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*GetPublicUserResponse, error) {
	out := new(GetPublicUserResponse)
	err := c.cc.Invoke(ctx, "/services.auth_service.AuthService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUserInfo(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, "/services.auth_service.AuthService/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUserByEmail(ctx context.Context, in *UserEmailRequest, opts ...grpc.CallOption) (*UserId, error) {
	out := new(UserId)
	err := c.cc.Invoke(ctx, "/services.auth_service.AuthService/GetUserByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUserIdByCookie(ctx context.Context, in *GetUserIdByCookieRequest, opts ...grpc.CallOption) (*UserId, error) {
	out := new(UserId)
	err := c.cc.Invoke(ctx, "/services.auth_service.AuthService/GetUserIdByCookie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CheckEmail(ctx context.Context, in *CheckEmailRequest, opts ...grpc.CallOption) (*GetPublicUserResponse, error) {
	out := new(GetPublicUserResponse)
	err := c.cc.Invoke(ctx, "/services.auth_service.AuthService/CheckEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUsersCompany(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*GetPrivateCompanyResponse, error) {
	out := new(GetPrivateCompanyResponse)
	err := c.cc.Invoke(ctx, "/services.auth_service.AuthService/GetUsersCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetCompanyUserLink(ctx context.Context, in *UserAndCompanyIdsRequest, opts ...grpc.CallOption) (*GetCompanyUserLinkResponse, error) {
	out := new(GetCompanyUserLinkResponse)
	err := c.cc.Invoke(ctx, "/services.auth_service.AuthService/GetCompanyUserLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateUserBalance(ctx context.Context, in *UpdateUserBalanceRequest, opts ...grpc.CallOption) (*GetPublicUserResponse, error) {
	out := new(GetPublicUserResponse)
	err := c.cc.Invoke(ctx, "/services.auth_service.AuthService/UpdateUserBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AddPayment(ctx context.Context, in *AddPaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error) {
	out := new(PaymentResponse)
	err := c.cc.Invoke(ctx, "/services.auth_service.AuthService/AddPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdatePayment(ctx context.Context, in *UpdatePaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error) {
	out := new(PaymentResponse)
	err := c.cc.Invoke(ctx, "/services.auth_service.AuthService/UpdatePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUsersPayments(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*PaymentsResponse, error) {
	out := new(PaymentsResponse)
	err := c.cc.Invoke(ctx, "/services.auth_service.AuthService/GetUsersPayments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetPayment(ctx context.Context, in *GetPaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error) {
	out := new(PaymentResponse)
	err := c.cc.Invoke(ctx, "/services.auth_service.AuthService/GetPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) HandlePaidPayments(ctx context.Context, in *HandlePaidPaymentsRequest, opts ...grpc.CallOption) (*HandlePaidPaymentsResponse, error) {
	out := new(HandlePaidPaymentsResponse)
	err := c.cc.Invoke(ctx, "/services.auth_service.AuthService/HandlePaidPayments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CountUsersPayments(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*PaymentsAmountResponse, error) {
	out := new(PaymentsAmountResponse)
	err := c.cc.Invoke(ctx, "/services.auth_service.AuthService/CountUsersPayments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	ValidateSession(context.Context, *Session) (*ValidateSessionResponse, error)
	LogoutUser(context.Context, *Session) (*emptypb.Empty, error)
	LoginUser(context.Context, *LoginRequest) (*LoginResponse, error)
	FastRegister(context.Context, *FastRegisterRequest) (*LoginResponse, error)
	RegisterUser(context.Context, *RegisterRequest) (*RegisterResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*GetPublicUserResponse, error)
	GetUserInfo(context.Context, *GetUserRequest) (*UserInfo, error)
	GetUserByEmail(context.Context, *UserEmailRequest) (*UserId, error)
	GetUserIdByCookie(context.Context, *GetUserIdByCookieRequest) (*UserId, error)
	CheckEmail(context.Context, *CheckEmailRequest) (*GetPublicUserResponse, error)
	GetUsersCompany(context.Context, *UserIdRequest) (*GetPrivateCompanyResponse, error)
	GetCompanyUserLink(context.Context, *UserAndCompanyIdsRequest) (*GetCompanyUserLinkResponse, error)
	UpdateUserBalance(context.Context, *UpdateUserBalanceRequest) (*GetPublicUserResponse, error)
	AddPayment(context.Context, *AddPaymentRequest) (*PaymentResponse, error)
	UpdatePayment(context.Context, *UpdatePaymentRequest) (*PaymentResponse, error)
	GetUsersPayments(context.Context, *UserIdRequest) (*PaymentsResponse, error)
	GetPayment(context.Context, *GetPaymentRequest) (*PaymentResponse, error)
	HandlePaidPayments(context.Context, *HandlePaidPaymentsRequest) (*HandlePaidPaymentsResponse, error)
	CountUsersPayments(context.Context, *UserIdRequest) (*PaymentsAmountResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) ValidateSession(context.Context, *Session) (*ValidateSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateSession not implemented")
}
func (UnimplementedAuthServiceServer) LogoutUser(context.Context, *Session) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogoutUser not implemented")
}
func (UnimplementedAuthServiceServer) LoginUser(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedAuthServiceServer) FastRegister(context.Context, *FastRegisterRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FastRegister not implemented")
}
func (UnimplementedAuthServiceServer) RegisterUser(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedAuthServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedAuthServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*GetPublicUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedAuthServiceServer) GetUserInfo(context.Context, *GetUserRequest) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedAuthServiceServer) GetUserByEmail(context.Context, *UserEmailRequest) (*UserId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByEmail not implemented")
}
func (UnimplementedAuthServiceServer) GetUserIdByCookie(context.Context, *GetUserIdByCookieRequest) (*UserId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserIdByCookie not implemented")
}
func (UnimplementedAuthServiceServer) CheckEmail(context.Context, *CheckEmailRequest) (*GetPublicUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckEmail not implemented")
}
func (UnimplementedAuthServiceServer) GetUsersCompany(context.Context, *UserIdRequest) (*GetPrivateCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersCompany not implemented")
}
func (UnimplementedAuthServiceServer) GetCompanyUserLink(context.Context, *UserAndCompanyIdsRequest) (*GetCompanyUserLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyUserLink not implemented")
}
func (UnimplementedAuthServiceServer) UpdateUserBalance(context.Context, *UpdateUserBalanceRequest) (*GetPublicUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserBalance not implemented")
}
func (UnimplementedAuthServiceServer) AddPayment(context.Context, *AddPaymentRequest) (*PaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPayment not implemented")
}
func (UnimplementedAuthServiceServer) UpdatePayment(context.Context, *UpdatePaymentRequest) (*PaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePayment not implemented")
}
func (UnimplementedAuthServiceServer) GetUsersPayments(context.Context, *UserIdRequest) (*PaymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersPayments not implemented")
}
func (UnimplementedAuthServiceServer) GetPayment(context.Context, *GetPaymentRequest) (*PaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayment not implemented")
}
func (UnimplementedAuthServiceServer) HandlePaidPayments(context.Context, *HandlePaidPaymentsRequest) (*HandlePaidPaymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandlePaidPayments not implemented")
}
func (UnimplementedAuthServiceServer) CountUsersPayments(context.Context, *UserIdRequest) (*PaymentsAmountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountUsersPayments not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_ValidateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ValidateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.auth_service.AuthService/ValidateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ValidateSession(ctx, req.(*Session))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_LogoutUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).LogoutUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.auth_service.AuthService/LogoutUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).LogoutUser(ctx, req.(*Session))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.auth_service.AuthService/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).LoginUser(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_FastRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FastRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).FastRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.auth_service.AuthService/FastRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).FastRegister(ctx, req.(*FastRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.auth_service.AuthService/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RegisterUser(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.auth_service.AuthService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.auth_service.AuthService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.auth_service.AuthService/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUserInfo(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.auth_service.AuthService/GetUserByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUserByEmail(ctx, req.(*UserEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUserIdByCookie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserIdByCookieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUserIdByCookie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.auth_service.AuthService/GetUserIdByCookie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUserIdByCookie(ctx, req.(*GetUserIdByCookieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CheckEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CheckEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.auth_service.AuthService/CheckEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CheckEmail(ctx, req.(*CheckEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUsersCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUsersCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.auth_service.AuthService/GetUsersCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUsersCompany(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetCompanyUserLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAndCompanyIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetCompanyUserLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.auth_service.AuthService/GetCompanyUserLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetCompanyUserLink(ctx, req.(*UserAndCompanyIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateUserBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateUserBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.auth_service.AuthService/UpdateUserBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateUserBalance(ctx, req.(*UpdateUserBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AddPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AddPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.auth_service.AuthService/AddPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AddPayment(ctx, req.(*AddPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.auth_service.AuthService/UpdatePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdatePayment(ctx, req.(*UpdatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUsersPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUsersPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.auth_service.AuthService/GetUsersPayments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUsersPayments(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.auth_service.AuthService/GetPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetPayment(ctx, req.(*GetPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_HandlePaidPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandlePaidPaymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).HandlePaidPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.auth_service.AuthService/HandlePaidPayments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).HandlePaidPayments(ctx, req.(*HandlePaidPaymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CountUsersPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CountUsersPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.auth_service.AuthService/CountUsersPayments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CountUsersPayments(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.auth_service.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateSession",
			Handler:    _AuthService_ValidateSession_Handler,
		},
		{
			MethodName: "LogoutUser",
			Handler:    _AuthService_LogoutUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _AuthService_LoginUser_Handler,
		},
		{
			MethodName: "FastRegister",
			Handler:    _AuthService_FastRegister_Handler,
		},
		{
			MethodName: "RegisterUser",
			Handler:    _AuthService_RegisterUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _AuthService_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _AuthService_UpdateUser_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _AuthService_GetUserInfo_Handler,
		},
		{
			MethodName: "GetUserByEmail",
			Handler:    _AuthService_GetUserByEmail_Handler,
		},
		{
			MethodName: "GetUserIdByCookie",
			Handler:    _AuthService_GetUserIdByCookie_Handler,
		},
		{
			MethodName: "CheckEmail",
			Handler:    _AuthService_CheckEmail_Handler,
		},
		{
			MethodName: "GetUsersCompany",
			Handler:    _AuthService_GetUsersCompany_Handler,
		},
		{
			MethodName: "GetCompanyUserLink",
			Handler:    _AuthService_GetCompanyUserLink_Handler,
		},
		{
			MethodName: "UpdateUserBalance",
			Handler:    _AuthService_UpdateUserBalance_Handler,
		},
		{
			MethodName: "AddPayment",
			Handler:    _AuthService_AddPayment_Handler,
		},
		{
			MethodName: "UpdatePayment",
			Handler:    _AuthService_UpdatePayment_Handler,
		},
		{
			MethodName: "GetUsersPayments",
			Handler:    _AuthService_GetUsersPayments_Handler,
		},
		{
			MethodName: "GetPayment",
			Handler:    _AuthService_GetPayment_Handler,
		},
		{
			MethodName: "HandlePaidPayments",
			Handler:    _AuthService_HandlePaidPayments_Handler,
		},
		{
			MethodName: "CountUsersPayments",
			Handler:    _AuthService_CountUsersPayments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/services/auth/api.proto",
}
