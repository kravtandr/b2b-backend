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

// AddProduct mocks base method.
func (m *MockProductsCategoriesServiceClient) AddProduct(ctx context.Context, in *productsCategories_service.AddProductRequest, opts ...grpc.CallOption) (*productsCategories_service.GetProduct, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddProduct", varargs...)
	ret0, _ := ret[0].(*productsCategories_service.GetProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddProduct indicates an expected call of AddProduct.
func (mr *MockProductsCategoriesServiceClientMockRecorder) AddProduct(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProduct", reflect.TypeOf((*MockProductsCategoriesServiceClient)(nil).AddProduct), varargs...)
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

// GetCompanyProducts mocks base method.
func (m *MockProductsCategoriesServiceClient) GetCompanyProducts(ctx context.Context, in *productsCategories_service.GetCompanyProductsRequest, opts ...grpc.CallOption) (*productsCategories_service.GetProductsListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCompanyProducts", varargs...)
	ret0, _ := ret[0].(*productsCategories_service.GetProductsListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompanyProducts indicates an expected call of GetCompanyProducts.
func (mr *MockProductsCategoriesServiceClientMockRecorder) GetCompanyProducts(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompanyProducts", reflect.TypeOf((*MockProductsCategoriesServiceClient)(nil).GetCompanyProducts), varargs...)
}

// GetProductById mocks base method.
func (m *MockProductsCategoriesServiceClient) GetProductById(ctx context.Context, in *productsCategories_service.GetProductByID, opts ...grpc.CallOption) (*productsCategories_service.GetProduct, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProductById", varargs...)
	ret0, _ := ret[0].(*productsCategories_service.GetProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductById indicates an expected call of GetProductById.
func (mr *MockProductsCategoriesServiceClientMockRecorder) GetProductById(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductById", reflect.TypeOf((*MockProductsCategoriesServiceClient)(nil).GetProductById), varargs...)
}

// GetProductsList mocks base method.
func (m *MockProductsCategoriesServiceClient) GetProductsList(ctx context.Context, in *productsCategories_service.GetProductsListRequest, opts ...grpc.CallOption) (*productsCategories_service.GetProductsListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProductsList", varargs...)
	ret0, _ := ret[0].(*productsCategories_service.GetProductsListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductsList indicates an expected call of GetProductsList.
func (mr *MockProductsCategoriesServiceClientMockRecorder) GetProductsList(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsList", reflect.TypeOf((*MockProductsCategoriesServiceClient)(nil).GetProductsList), varargs...)
}

// GetProductsListByFilters mocks base method.
func (m *MockProductsCategoriesServiceClient) GetProductsListByFilters(ctx context.Context, in *productsCategories_service.GetProductsListByFiltersRequest, opts ...grpc.CallOption) (*productsCategories_service.GetProductsByFiltersResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProductsListByFilters", varargs...)
	ret0, _ := ret[0].(*productsCategories_service.GetProductsByFiltersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductsListByFilters indicates an expected call of GetProductsListByFilters.
func (mr *MockProductsCategoriesServiceClientMockRecorder) GetProductsListByFilters(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsListByFilters", reflect.TypeOf((*MockProductsCategoriesServiceClient)(nil).GetProductsListByFilters), varargs...)
}

// SearchCategories mocks base method.
func (m *MockProductsCategoriesServiceClient) SearchCategories(ctx context.Context, in *productsCategories_service.SearchItemNameWithSkipLimitRequest, opts ...grpc.CallOption) (*productsCategories_service.GetCategories, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchCategories", varargs...)
	ret0, _ := ret[0].(*productsCategories_service.GetCategories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchCategories indicates an expected call of SearchCategories.
func (mr *MockProductsCategoriesServiceClientMockRecorder) SearchCategories(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchCategories", reflect.TypeOf((*MockProductsCategoriesServiceClient)(nil).SearchCategories), varargs...)
}

// SearchProducts mocks base method.
func (m *MockProductsCategoriesServiceClient) SearchProducts(ctx context.Context, in *productsCategories_service.SearchItemNameWithSkipLimitRequest, opts ...grpc.CallOption) (*productsCategories_service.GetProductsListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchProducts", varargs...)
	ret0, _ := ret[0].(*productsCategories_service.GetProductsListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchProducts indicates an expected call of SearchProducts.
func (mr *MockProductsCategoriesServiceClientMockRecorder) SearchProducts(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchProducts", reflect.TypeOf((*MockProductsCategoriesServiceClient)(nil).SearchProducts), varargs...)
}

// UpdateProduct mocks base method.
func (m *MockProductsCategoriesServiceClient) UpdateProduct(ctx context.Context, in *productsCategories_service.UpdateProductRequest, opts ...grpc.CallOption) (*productsCategories_service.GetProduct, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateProduct", varargs...)
	ret0, _ := ret[0].(*productsCategories_service.GetProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockProductsCategoriesServiceClientMockRecorder) UpdateProduct(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockProductsCategoriesServiceClient)(nil).UpdateProduct), varargs...)
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

// AddProduct mocks base method.
func (m *MockProductsCategoriesServiceServer) AddProduct(arg0 context.Context, arg1 *productsCategories_service.AddProductRequest) (*productsCategories_service.GetProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProduct", arg0, arg1)
	ret0, _ := ret[0].(*productsCategories_service.GetProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddProduct indicates an expected call of AddProduct.
func (mr *MockProductsCategoriesServiceServerMockRecorder) AddProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProduct", reflect.TypeOf((*MockProductsCategoriesServiceServer)(nil).AddProduct), arg0, arg1)
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

// GetCompanyProducts mocks base method.
func (m *MockProductsCategoriesServiceServer) GetCompanyProducts(arg0 context.Context, arg1 *productsCategories_service.GetCompanyProductsRequest) (*productsCategories_service.GetProductsListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompanyProducts", arg0, arg1)
	ret0, _ := ret[0].(*productsCategories_service.GetProductsListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompanyProducts indicates an expected call of GetCompanyProducts.
func (mr *MockProductsCategoriesServiceServerMockRecorder) GetCompanyProducts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompanyProducts", reflect.TypeOf((*MockProductsCategoriesServiceServer)(nil).GetCompanyProducts), arg0, arg1)
}

// GetProductById mocks base method.
func (m *MockProductsCategoriesServiceServer) GetProductById(arg0 context.Context, arg1 *productsCategories_service.GetProductByID) (*productsCategories_service.GetProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductById", arg0, arg1)
	ret0, _ := ret[0].(*productsCategories_service.GetProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductById indicates an expected call of GetProductById.
func (mr *MockProductsCategoriesServiceServerMockRecorder) GetProductById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductById", reflect.TypeOf((*MockProductsCategoriesServiceServer)(nil).GetProductById), arg0, arg1)
}

// GetProductsList mocks base method.
func (m *MockProductsCategoriesServiceServer) GetProductsList(arg0 context.Context, arg1 *productsCategories_service.GetProductsListRequest) (*productsCategories_service.GetProductsListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsList", arg0, arg1)
	ret0, _ := ret[0].(*productsCategories_service.GetProductsListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductsList indicates an expected call of GetProductsList.
func (mr *MockProductsCategoriesServiceServerMockRecorder) GetProductsList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsList", reflect.TypeOf((*MockProductsCategoriesServiceServer)(nil).GetProductsList), arg0, arg1)
}

// GetProductsListByFilters mocks base method.
func (m *MockProductsCategoriesServiceServer) GetProductsListByFilters(arg0 context.Context, arg1 *productsCategories_service.GetProductsListByFiltersRequest) (*productsCategories_service.GetProductsByFiltersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsListByFilters", arg0, arg1)
	ret0, _ := ret[0].(*productsCategories_service.GetProductsByFiltersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductsListByFilters indicates an expected call of GetProductsListByFilters.
func (mr *MockProductsCategoriesServiceServerMockRecorder) GetProductsListByFilters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsListByFilters", reflect.TypeOf((*MockProductsCategoriesServiceServer)(nil).GetProductsListByFilters), arg0, arg1)
}

// SearchCategories mocks base method.
func (m *MockProductsCategoriesServiceServer) SearchCategories(arg0 context.Context, arg1 *productsCategories_service.SearchItemNameWithSkipLimitRequest) (*productsCategories_service.GetCategories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchCategories", arg0, arg1)
	ret0, _ := ret[0].(*productsCategories_service.GetCategories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchCategories indicates an expected call of SearchCategories.
func (mr *MockProductsCategoriesServiceServerMockRecorder) SearchCategories(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchCategories", reflect.TypeOf((*MockProductsCategoriesServiceServer)(nil).SearchCategories), arg0, arg1)
}

// SearchProducts mocks base method.
func (m *MockProductsCategoriesServiceServer) SearchProducts(arg0 context.Context, arg1 *productsCategories_service.SearchItemNameWithSkipLimitRequest) (*productsCategories_service.GetProductsListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchProducts", arg0, arg1)
	ret0, _ := ret[0].(*productsCategories_service.GetProductsListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchProducts indicates an expected call of SearchProducts.
func (mr *MockProductsCategoriesServiceServerMockRecorder) SearchProducts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchProducts", reflect.TypeOf((*MockProductsCategoriesServiceServer)(nil).SearchProducts), arg0, arg1)
}

// UpdateProduct mocks base method.
func (m *MockProductsCategoriesServiceServer) UpdateProduct(arg0 context.Context, arg1 *productsCategories_service.UpdateProductRequest) (*productsCategories_service.GetProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", arg0, arg1)
	ret0, _ := ret[0].(*productsCategories_service.GetProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockProductsCategoriesServiceServerMockRecorder) UpdateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockProductsCategoriesServiceServer)(nil).UpdateProduct), arg0, arg1)
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
