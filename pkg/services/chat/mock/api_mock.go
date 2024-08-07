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
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

// ChatHealthCheck mocks base method.
func (m *MockChatServiceClient) ChatHealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChatHealthCheck", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChatHealthCheck indicates an expected call of ChatHealthCheck.
func (mr *MockChatServiceClientMockRecorder) ChatHealthCheck(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChatHealthCheck", reflect.TypeOf((*MockChatServiceClient)(nil).ChatHealthCheck), varargs...)
}

// CheckIfUniqChat mocks base method.
func (m *MockChatServiceClient) CheckIfUniqChat(ctx context.Context, in *chat_service.CheckIfUniqChatRequest, opts ...grpc.CallOption) (*chat_service.CheckIfUniqChatResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckIfUniqChat", varargs...)
	ret0, _ := ret[0].(*chat_service.CheckIfUniqChatResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckIfUniqChat indicates an expected call of CheckIfUniqChat.
func (mr *MockChatServiceClientMockRecorder) CheckIfUniqChat(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIfUniqChat", reflect.TypeOf((*MockChatServiceClient)(nil).CheckIfUniqChat), varargs...)
}

// DeleteChat mocks base method.
func (m *MockChatServiceClient) DeleteChat(ctx context.Context, in *chat_service.IdRequest, opts ...grpc.CallOption) (*chat_service.Bool, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteChat", varargs...)
	ret0, _ := ret[0].(*chat_service.Bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteChat indicates an expected call of DeleteChat.
func (mr *MockChatServiceClientMockRecorder) DeleteChat(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteChat", reflect.TypeOf((*MockChatServiceClient)(nil).DeleteChat), varargs...)
}

// GetAllChatsAndLastMsg mocks base method.
func (m *MockChatServiceClient) GetAllChatsAndLastMsg(ctx context.Context, in *chat_service.IdRequest, opts ...grpc.CallOption) (*chat_service.GetAllChatsAndLastMsgResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllChatsAndLastMsg", varargs...)
	ret0, _ := ret[0].(*chat_service.GetAllChatsAndLastMsgResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllChatsAndLastMsg indicates an expected call of GetAllChatsAndLastMsg.
func (mr *MockChatServiceClientMockRecorder) GetAllChatsAndLastMsg(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllChatsAndLastMsg", reflect.TypeOf((*MockChatServiceClient)(nil).GetAllChatsAndLastMsg), varargs...)
}

// GetAllUserChats mocks base method.
func (m *MockChatServiceClient) GetAllUserChats(ctx context.Context, in *chat_service.IdRequest, opts ...grpc.CallOption) (*chat_service.GetAllUserChatsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllUserChats", varargs...)
	ret0, _ := ret[0].(*chat_service.GetAllUserChatsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUserChats indicates an expected call of GetAllUserChats.
func (mr *MockChatServiceClientMockRecorder) GetAllUserChats(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUserChats", reflect.TypeOf((*MockChatServiceClient)(nil).GetAllUserChats), varargs...)
}

// GetChat mocks base method.
func (m *MockChatServiceClient) GetChat(ctx context.Context, in *chat_service.GetChatRequest, opts ...grpc.CallOption) (*chat_service.ChatResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetChat", varargs...)
	ret0, _ := ret[0].(*chat_service.ChatResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChat indicates an expected call of GetChat.
func (mr *MockChatServiceClientMockRecorder) GetChat(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChat", reflect.TypeOf((*MockChatServiceClient)(nil).GetChat), varargs...)
}

// GetChatById mocks base method.
func (m *MockChatServiceClient) GetChatById(ctx context.Context, in *chat_service.IdRequest, opts ...grpc.CallOption) (*chat_service.ChatResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetChatById", varargs...)
	ret0, _ := ret[0].(*chat_service.ChatResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChatById indicates an expected call of GetChatById.
func (mr *MockChatServiceClientMockRecorder) GetChatById(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChatById", reflect.TypeOf((*MockChatServiceClient)(nil).GetChatById), varargs...)
}

// GetMsgsFromChat mocks base method.
func (m *MockChatServiceClient) GetMsgsFromChat(ctx context.Context, in *chat_service.ChatAndUserIdRequest, opts ...grpc.CallOption) (*chat_service.MsgsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMsgsFromChat", varargs...)
	ret0, _ := ret[0].(*chat_service.MsgsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMsgsFromChat indicates an expected call of GetMsgsFromChat.
func (mr *MockChatServiceClientMockRecorder) GetMsgsFromChat(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMsgsFromChat", reflect.TypeOf((*MockChatServiceClient)(nil).GetMsgsFromChat), varargs...)
}

// NewChat mocks base method.
func (m *MockChatServiceClient) NewChat(ctx context.Context, in *chat_service.NewChatRequest, opts ...grpc.CallOption) (*chat_service.ChatResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewChat", varargs...)
	ret0, _ := ret[0].(*chat_service.ChatResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewChat indicates an expected call of NewChat.
func (mr *MockChatServiceClientMockRecorder) NewChat(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewChat", reflect.TypeOf((*MockChatServiceClient)(nil).NewChat), varargs...)
}

// UpdateChatStatus mocks base method.
func (m *MockChatServiceClient) UpdateChatStatus(ctx context.Context, in *chat_service.UpdateChatStatusRequest, opts ...grpc.CallOption) (*chat_service.ChatResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateChatStatus", varargs...)
	ret0, _ := ret[0].(*chat_service.ChatResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateChatStatus indicates an expected call of UpdateChatStatus.
func (mr *MockChatServiceClientMockRecorder) UpdateChatStatus(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateChatStatus", reflect.TypeOf((*MockChatServiceClient)(nil).UpdateChatStatus), varargs...)
}

// WriteNewMsg mocks base method.
func (m *MockChatServiceClient) WriteNewMsg(ctx context.Context, in *chat_service.WriteNewMsgRequest, opts ...grpc.CallOption) (*chat_service.IdResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WriteNewMsg", varargs...)
	ret0, _ := ret[0].(*chat_service.IdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteNewMsg indicates an expected call of WriteNewMsg.
func (mr *MockChatServiceClientMockRecorder) WriteNewMsg(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteNewMsg", reflect.TypeOf((*MockChatServiceClient)(nil).WriteNewMsg), varargs...)
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

// ChatHealthCheck mocks base method.
func (m *MockChatServiceServer) ChatHealthCheck(arg0 context.Context, arg1 *emptypb.Empty) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChatHealthCheck", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChatHealthCheck indicates an expected call of ChatHealthCheck.
func (mr *MockChatServiceServerMockRecorder) ChatHealthCheck(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChatHealthCheck", reflect.TypeOf((*MockChatServiceServer)(nil).ChatHealthCheck), arg0, arg1)
}

// CheckIfUniqChat mocks base method.
func (m *MockChatServiceServer) CheckIfUniqChat(arg0 context.Context, arg1 *chat_service.CheckIfUniqChatRequest) (*chat_service.CheckIfUniqChatResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIfUniqChat", arg0, arg1)
	ret0, _ := ret[0].(*chat_service.CheckIfUniqChatResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckIfUniqChat indicates an expected call of CheckIfUniqChat.
func (mr *MockChatServiceServerMockRecorder) CheckIfUniqChat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIfUniqChat", reflect.TypeOf((*MockChatServiceServer)(nil).CheckIfUniqChat), arg0, arg1)
}

// DeleteChat mocks base method.
func (m *MockChatServiceServer) DeleteChat(arg0 context.Context, arg1 *chat_service.IdRequest) (*chat_service.Bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteChat", arg0, arg1)
	ret0, _ := ret[0].(*chat_service.Bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteChat indicates an expected call of DeleteChat.
func (mr *MockChatServiceServerMockRecorder) DeleteChat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteChat", reflect.TypeOf((*MockChatServiceServer)(nil).DeleteChat), arg0, arg1)
}

// GetAllChatsAndLastMsg mocks base method.
func (m *MockChatServiceServer) GetAllChatsAndLastMsg(arg0 context.Context, arg1 *chat_service.IdRequest) (*chat_service.GetAllChatsAndLastMsgResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllChatsAndLastMsg", arg0, arg1)
	ret0, _ := ret[0].(*chat_service.GetAllChatsAndLastMsgResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllChatsAndLastMsg indicates an expected call of GetAllChatsAndLastMsg.
func (mr *MockChatServiceServerMockRecorder) GetAllChatsAndLastMsg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllChatsAndLastMsg", reflect.TypeOf((*MockChatServiceServer)(nil).GetAllChatsAndLastMsg), arg0, arg1)
}

// GetAllUserChats mocks base method.
func (m *MockChatServiceServer) GetAllUserChats(arg0 context.Context, arg1 *chat_service.IdRequest) (*chat_service.GetAllUserChatsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUserChats", arg0, arg1)
	ret0, _ := ret[0].(*chat_service.GetAllUserChatsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUserChats indicates an expected call of GetAllUserChats.
func (mr *MockChatServiceServerMockRecorder) GetAllUserChats(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUserChats", reflect.TypeOf((*MockChatServiceServer)(nil).GetAllUserChats), arg0, arg1)
}

// GetChat mocks base method.
func (m *MockChatServiceServer) GetChat(arg0 context.Context, arg1 *chat_service.GetChatRequest) (*chat_service.ChatResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChat", arg0, arg1)
	ret0, _ := ret[0].(*chat_service.ChatResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChat indicates an expected call of GetChat.
func (mr *MockChatServiceServerMockRecorder) GetChat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChat", reflect.TypeOf((*MockChatServiceServer)(nil).GetChat), arg0, arg1)
}

// GetChatById mocks base method.
func (m *MockChatServiceServer) GetChatById(arg0 context.Context, arg1 *chat_service.IdRequest) (*chat_service.ChatResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChatById", arg0, arg1)
	ret0, _ := ret[0].(*chat_service.ChatResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChatById indicates an expected call of GetChatById.
func (mr *MockChatServiceServerMockRecorder) GetChatById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChatById", reflect.TypeOf((*MockChatServiceServer)(nil).GetChatById), arg0, arg1)
}

// GetMsgsFromChat mocks base method.
func (m *MockChatServiceServer) GetMsgsFromChat(arg0 context.Context, arg1 *chat_service.ChatAndUserIdRequest) (*chat_service.MsgsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMsgsFromChat", arg0, arg1)
	ret0, _ := ret[0].(*chat_service.MsgsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMsgsFromChat indicates an expected call of GetMsgsFromChat.
func (mr *MockChatServiceServerMockRecorder) GetMsgsFromChat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMsgsFromChat", reflect.TypeOf((*MockChatServiceServer)(nil).GetMsgsFromChat), arg0, arg1)
}

// NewChat mocks base method.
func (m *MockChatServiceServer) NewChat(arg0 context.Context, arg1 *chat_service.NewChatRequest) (*chat_service.ChatResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewChat", arg0, arg1)
	ret0, _ := ret[0].(*chat_service.ChatResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewChat indicates an expected call of NewChat.
func (mr *MockChatServiceServerMockRecorder) NewChat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewChat", reflect.TypeOf((*MockChatServiceServer)(nil).NewChat), arg0, arg1)
}

// UpdateChatStatus mocks base method.
func (m *MockChatServiceServer) UpdateChatStatus(arg0 context.Context, arg1 *chat_service.UpdateChatStatusRequest) (*chat_service.ChatResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateChatStatus", arg0, arg1)
	ret0, _ := ret[0].(*chat_service.ChatResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateChatStatus indicates an expected call of UpdateChatStatus.
func (mr *MockChatServiceServerMockRecorder) UpdateChatStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateChatStatus", reflect.TypeOf((*MockChatServiceServer)(nil).UpdateChatStatus), arg0, arg1)
}

// WriteNewMsg mocks base method.
func (m *MockChatServiceServer) WriteNewMsg(arg0 context.Context, arg1 *chat_service.WriteNewMsgRequest) (*chat_service.IdResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteNewMsg", arg0, arg1)
	ret0, _ := ret[0].(*chat_service.IdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteNewMsg indicates an expected call of WriteNewMsg.
func (mr *MockChatServiceServerMockRecorder) WriteNewMsg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteNewMsg", reflect.TypeOf((*MockChatServiceServer)(nil).WriteNewMsg), arg0, arg1)
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
