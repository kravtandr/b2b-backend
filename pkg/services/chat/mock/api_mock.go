// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/services/chat/api_grpc.pb.go

// Package mock_chat_service is a generated GoMock package.
package mock_chat_service

import (
	chat_service "b2b/m/pkg/services/chat"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockChatServiceClient is a mock of ChatServiceClient interface.
type MockChatServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockChatServiceClientMockRecorder
}

// MockChatServiceClientMockRecorder is the mock recorder for MockChatServiceClient.
type MockChatServiceClientMockRecorder struct {
	mock *MockChatServiceClient
}

// NewMockChatServiceClient creates a new mock instance.
func NewMockChatServiceClient(ctrl *gomock.Controller) *MockChatServiceClient {
	mock := &MockChatServiceClient{ctrl: ctrl}
	mock.recorder = &MockChatServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatServiceClient) EXPECT() *MockChatServiceClientMockRecorder {
	return m.recorder
}

// LoginUser mocks base method.
func (m *MockChatServiceClient) LoginUser(ctx context.Context, in *chat_service.LoginRequest, opts ...grpc.CallOption) (*chat_service.LoginResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LoginUser", varargs...)
	ret0, _ := ret[0].(*chat_service.LoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginUser indicates an expected call of LoginUser.
func (mr *MockChatServiceClientMockRecorder) LoginUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUser", reflect.TypeOf((*MockChatServiceClient)(nil).LoginUser), varargs...)
}

// MockChatServiceServer is a mock of ChatServiceServer interface.
type MockChatServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockChatServiceServerMockRecorder
}

// MockChatServiceServerMockRecorder is the mock recorder for MockChatServiceServer.
type MockChatServiceServerMockRecorder struct {
	mock *MockChatServiceServer
}

// NewMockChatServiceServer creates a new mock instance.
func NewMockChatServiceServer(ctrl *gomock.Controller) *MockChatServiceServer {
	mock := &MockChatServiceServer{ctrl: ctrl}
	mock.recorder = &MockChatServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatServiceServer) EXPECT() *MockChatServiceServerMockRecorder {
	return m.recorder
}

// LoginUser mocks base method.
func (m *MockChatServiceServer) LoginUser(arg0 context.Context, arg1 *chat_service.LoginRequest) (*chat_service.LoginResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUser", arg0, arg1)
	ret0, _ := ret[0].(*chat_service.LoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginUser indicates an expected call of LoginUser.
func (mr *MockChatServiceServerMockRecorder) LoginUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUser", reflect.TypeOf((*MockChatServiceServer)(nil).LoginUser), arg0, arg1)
}

// mustEmbedUnimplementedChatServiceServer mocks base method.
func (m *MockChatServiceServer) mustEmbedUnimplementedChatServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedChatServiceServer")
}

// mustEmbedUnimplementedChatServiceServer indicates an expected call of mustEmbedUnimplementedChatServiceServer.
func (mr *MockChatServiceServerMockRecorder) mustEmbedUnimplementedChatServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedChatServiceServer", reflect.TypeOf((*MockChatServiceServer)(nil).mustEmbedUnimplementedChatServiceServer))
}

// MockUnsafeChatServiceServer is a mock of UnsafeChatServiceServer interface.
type MockUnsafeChatServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeChatServiceServerMockRecorder
}

// MockUnsafeChatServiceServerMockRecorder is the mock recorder for MockUnsafeChatServiceServer.
type MockUnsafeChatServiceServerMockRecorder struct {
	mock *MockUnsafeChatServiceServer
}

// NewMockUnsafeChatServiceServer creates a new mock instance.
func NewMockUnsafeChatServiceServer(ctrl *gomock.Controller) *MockUnsafeChatServiceServer {
	mock := &MockUnsafeChatServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeChatServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeChatServiceServer) EXPECT() *MockUnsafeChatServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedChatServiceServer mocks base method.
func (m *MockUnsafeChatServiceServer) mustEmbedUnimplementedChatServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedChatServiceServer")
}

// mustEmbedUnimplementedChatServiceServer indicates an expected call of mustEmbedUnimplementedChatServiceServer.
func (mr *MockUnsafeChatServiceServerMockRecorder) mustEmbedUnimplementedChatServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedChatServiceServer", reflect.TypeOf((*MockUnsafeChatServiceServer)(nil).mustEmbedUnimplementedChatServiceServer))
}
