// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/services/fastOrder/api_grpc.pb.go

// Package mock_fastOrder_service is a generated GoMock package.
package mock_fastOrder_service

import (
	fastOrder_service "b2b/m/pkg/services/fastOrder"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockFastOrderServiceClient is a mock of FastOrderServiceClient interface.
type MockFastOrderServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockFastOrderServiceClientMockRecorder
}

// MockFastOrderServiceClientMockRecorder is the mock recorder for MockFastOrderServiceClient.
type MockFastOrderServiceClientMockRecorder struct {
	mock *MockFastOrderServiceClient
}

// NewMockFastOrderServiceClient creates a new mock instance.
func NewMockFastOrderServiceClient(ctrl *gomock.Controller) *MockFastOrderServiceClient {
	mock := &MockFastOrderServiceClient{ctrl: ctrl}
	mock.recorder = &MockFastOrderServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFastOrderServiceClient) EXPECT() *MockFastOrderServiceClientMockRecorder {
	return m.recorder
}

// FastOrder mocks base method.
func (m *MockFastOrderServiceClient) FastOrder(ctx context.Context, in *fastOrder_service.FastOrderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FastOrder", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FastOrder indicates an expected call of FastOrder.
func (mr *MockFastOrderServiceClientMockRecorder) FastOrder(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FastOrder", reflect.TypeOf((*MockFastOrderServiceClient)(nil).FastOrder), varargs...)
}

// MockFastOrderServiceServer is a mock of FastOrderServiceServer interface.
type MockFastOrderServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockFastOrderServiceServerMockRecorder
}

// MockFastOrderServiceServerMockRecorder is the mock recorder for MockFastOrderServiceServer.
type MockFastOrderServiceServerMockRecorder struct {
	mock *MockFastOrderServiceServer
}

// NewMockFastOrderServiceServer creates a new mock instance.
func NewMockFastOrderServiceServer(ctrl *gomock.Controller) *MockFastOrderServiceServer {
	mock := &MockFastOrderServiceServer{ctrl: ctrl}
	mock.recorder = &MockFastOrderServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFastOrderServiceServer) EXPECT() *MockFastOrderServiceServerMockRecorder {
	return m.recorder
}

// FastOrder mocks base method.
func (m *MockFastOrderServiceServer) FastOrder(arg0 context.Context, arg1 *fastOrder_service.FastOrderRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FastOrder", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FastOrder indicates an expected call of FastOrder.
func (mr *MockFastOrderServiceServerMockRecorder) FastOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FastOrder", reflect.TypeOf((*MockFastOrderServiceServer)(nil).FastOrder), arg0, arg1)
}

// mustEmbedUnimplementedFastOrderServiceServer mocks base method.
func (m *MockFastOrderServiceServer) mustEmbedUnimplementedFastOrderServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedFastOrderServiceServer")
}

// mustEmbedUnimplementedFastOrderServiceServer indicates an expected call of mustEmbedUnimplementedFastOrderServiceServer.
func (mr *MockFastOrderServiceServerMockRecorder) mustEmbedUnimplementedFastOrderServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedFastOrderServiceServer", reflect.TypeOf((*MockFastOrderServiceServer)(nil).mustEmbedUnimplementedFastOrderServiceServer))
}

// MockUnsafeFastOrderServiceServer is a mock of UnsafeFastOrderServiceServer interface.
type MockUnsafeFastOrderServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeFastOrderServiceServerMockRecorder
}

// MockUnsafeFastOrderServiceServerMockRecorder is the mock recorder for MockUnsafeFastOrderServiceServer.
type MockUnsafeFastOrderServiceServerMockRecorder struct {
	mock *MockUnsafeFastOrderServiceServer
}

// NewMockUnsafeFastOrderServiceServer creates a new mock instance.
func NewMockUnsafeFastOrderServiceServer(ctrl *gomock.Controller) *MockUnsafeFastOrderServiceServer {
	mock := &MockUnsafeFastOrderServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeFastOrderServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeFastOrderServiceServer) EXPECT() *MockUnsafeFastOrderServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedFastOrderServiceServer mocks base method.
func (m *MockUnsafeFastOrderServiceServer) mustEmbedUnimplementedFastOrderServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedFastOrderServiceServer")
}

// mustEmbedUnimplementedFastOrderServiceServer indicates an expected call of mustEmbedUnimplementedFastOrderServiceServer.
func (mr *MockUnsafeFastOrderServiceServerMockRecorder) mustEmbedUnimplementedFastOrderServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedFastOrderServiceServer", reflect.TypeOf((*MockUnsafeFastOrderServiceServer)(nil).mustEmbedUnimplementedFastOrderServiceServer))
}
