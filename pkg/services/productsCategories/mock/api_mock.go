// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/services/productsCategories/api_grpc.pb.go

// Package mock_productsCategories_service is a generated GoMock package.
package mock_productsCategories_service

import (
	productsCategories_service "b2b/m/pkg/services/productsCategories"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockProductsCategoriesServiceClient is a mock of ProductsCategoriesServiceClient interface.
type MockProductsCategoriesServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockProductsCategoriesServiceClientMockRecorder
}

// MockProductsCategoriesServiceClientMockRecorder is the mock recorder for MockProductsCategoriesServiceClient.
type MockProductsCategoriesServiceClientMockRecorder struct {
	mock *MockProductsCategoriesServiceClient
}

// NewMockProductsCategoriesServiceClient creates a new mock instance.
func NewMockProductsCategoriesServiceClient(ctrl *gomock.Controller) *MockProductsCategoriesServiceClient {
	mock := &MockProductsCategoriesServiceClient{ctrl: ctrl}
	mock.recorder = &MockProductsCategoriesServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductsCategoriesServiceClient) EXPECT() *MockProductsCategoriesServiceClientMockRecorder {
	return m.recorder
}

// GetCategoryById mocks base method.
func (m *MockProductsCategoriesServiceClient) GetCategoryById(ctx context.Context, in *productsCategories_service.GetCategoryByID, opts ...grpc.CallOption) (*productsCategories_service.GetCategory, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCategoryById", varargs...)
	ret0, _ := ret[0].(*productsCategories_service.GetCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoryById indicates an expected call of GetCategoryById.
func (mr *MockProductsCategoriesServiceClientMockRecorder) GetCategoryById(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryById", reflect.TypeOf((*MockProductsCategoriesServiceClient)(nil).GetCategoryById), varargs...)
}

// MockProductsCategoriesServiceServer is a mock of ProductsCategoriesServiceServer interface.
type MockProductsCategoriesServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockProductsCategoriesServiceServerMockRecorder
}

// MockProductsCategoriesServiceServerMockRecorder is the mock recorder for MockProductsCategoriesServiceServer.
type MockProductsCategoriesServiceServerMockRecorder struct {
	mock *MockProductsCategoriesServiceServer
}

// NewMockProductsCategoriesServiceServer creates a new mock instance.
func NewMockProductsCategoriesServiceServer(ctrl *gomock.Controller) *MockProductsCategoriesServiceServer {
	mock := &MockProductsCategoriesServiceServer{ctrl: ctrl}
	mock.recorder = &MockProductsCategoriesServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductsCategoriesServiceServer) EXPECT() *MockProductsCategoriesServiceServerMockRecorder {
	return m.recorder
}

// GetCategoryById mocks base method.
func (m *MockProductsCategoriesServiceServer) GetCategoryById(arg0 context.Context, arg1 *productsCategories_service.GetCategoryByID) (*productsCategories_service.GetCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoryById", arg0, arg1)
	ret0, _ := ret[0].(*productsCategories_service.GetCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoryById indicates an expected call of GetCategoryById.
func (mr *MockProductsCategoriesServiceServerMockRecorder) GetCategoryById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryById", reflect.TypeOf((*MockProductsCategoriesServiceServer)(nil).GetCategoryById), arg0, arg1)
}

// mustEmbedUnimplementedProductsCategoriesServiceServer mocks base method.
func (m *MockProductsCategoriesServiceServer) mustEmbedUnimplementedProductsCategoriesServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedProductsCategoriesServiceServer")
}

// mustEmbedUnimplementedProductsCategoriesServiceServer indicates an expected call of mustEmbedUnimplementedProductsCategoriesServiceServer.
func (mr *MockProductsCategoriesServiceServerMockRecorder) mustEmbedUnimplementedProductsCategoriesServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedProductsCategoriesServiceServer", reflect.TypeOf((*MockProductsCategoriesServiceServer)(nil).mustEmbedUnimplementedProductsCategoriesServiceServer))
}

// MockUnsafeProductsCategoriesServiceServer is a mock of UnsafeProductsCategoriesServiceServer interface.
type MockUnsafeProductsCategoriesServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeProductsCategoriesServiceServerMockRecorder
}

// MockUnsafeProductsCategoriesServiceServerMockRecorder is the mock recorder for MockUnsafeProductsCategoriesServiceServer.
type MockUnsafeProductsCategoriesServiceServerMockRecorder struct {
	mock *MockUnsafeProductsCategoriesServiceServer
}

// NewMockUnsafeProductsCategoriesServiceServer creates a new mock instance.
func NewMockUnsafeProductsCategoriesServiceServer(ctrl *gomock.Controller) *MockUnsafeProductsCategoriesServiceServer {
	mock := &MockUnsafeProductsCategoriesServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeProductsCategoriesServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeProductsCategoriesServiceServer) EXPECT() *MockUnsafeProductsCategoriesServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedProductsCategoriesServiceServer mocks base method.
func (m *MockUnsafeProductsCategoriesServiceServer) mustEmbedUnimplementedProductsCategoriesServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedProductsCategoriesServiceServer")
}

// mustEmbedUnimplementedProductsCategoriesServiceServer indicates an expected call of mustEmbedUnimplementedProductsCategoriesServiceServer.
func (mr *MockUnsafeProductsCategoriesServiceServerMockRecorder) mustEmbedUnimplementedProductsCategoriesServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedProductsCategoriesServiceServer", reflect.TypeOf((*MockUnsafeProductsCategoriesServiceServer)(nil).mustEmbedUnimplementedProductsCategoriesServiceServer))
}
