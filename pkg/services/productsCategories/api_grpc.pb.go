// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: pkg/services/productsCategories/api.proto

package productsCategories_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProductsCategoriesServiceClient is the client API for ProductsCategoriesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductsCategoriesServiceClient interface {
	AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*GetProduct, error)
	GetCategoryById(ctx context.Context, in *GetCategoryByID, opts ...grpc.CallOption) (*GetCategory, error)
	GetProductById(ctx context.Context, in *GetProductByID, opts ...grpc.CallOption) (*GetProduct, error)
	SearchCategories(ctx context.Context, in *SearchItemNameWithSkipLimitRequest, opts ...grpc.CallOption) (*GetCategories, error)
	SearchProducts(ctx context.Context, in *SearchItemNameWithSkipLimitRequest, opts ...grpc.CallOption) (*GetProductsListResponse, error)
	GetProductsList(ctx context.Context, in *GetProductsListRequest, opts ...grpc.CallOption) (*GetProductsListResponse, error)
	GetProductsListByFilters(ctx context.Context, in *GetProductsListByFiltersRequest, opts ...grpc.CallOption) (*GetProductsByFiltersResponse, error)
	GetCompanyProducts(ctx context.Context, in *GetCompanyProductsRequest, opts ...grpc.CallOption) (*GetProductsListResponse, error)
}

type productsCategoriesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductsCategoriesServiceClient(cc grpc.ClientConnInterface) ProductsCategoriesServiceClient {
	return &productsCategoriesServiceClient{cc}
}

func (c *productsCategoriesServiceClient) AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*GetProduct, error) {
	out := new(GetProduct)
	err := c.cc.Invoke(ctx, "/services.productsCategories_service.ProductsCategoriesService/AddProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsCategoriesServiceClient) GetCategoryById(ctx context.Context, in *GetCategoryByID, opts ...grpc.CallOption) (*GetCategory, error) {
	out := new(GetCategory)
	err := c.cc.Invoke(ctx, "/services.productsCategories_service.ProductsCategoriesService/GetCategoryById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsCategoriesServiceClient) GetProductById(ctx context.Context, in *GetProductByID, opts ...grpc.CallOption) (*GetProduct, error) {
	out := new(GetProduct)
	err := c.cc.Invoke(ctx, "/services.productsCategories_service.ProductsCategoriesService/GetProductById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsCategoriesServiceClient) SearchCategories(ctx context.Context, in *SearchItemNameWithSkipLimitRequest, opts ...grpc.CallOption) (*GetCategories, error) {
	out := new(GetCategories)
	err := c.cc.Invoke(ctx, "/services.productsCategories_service.ProductsCategoriesService/SearchCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsCategoriesServiceClient) SearchProducts(ctx context.Context, in *SearchItemNameWithSkipLimitRequest, opts ...grpc.CallOption) (*GetProductsListResponse, error) {
	out := new(GetProductsListResponse)
	err := c.cc.Invoke(ctx, "/services.productsCategories_service.ProductsCategoriesService/SearchProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsCategoriesServiceClient) GetProductsList(ctx context.Context, in *GetProductsListRequest, opts ...grpc.CallOption) (*GetProductsListResponse, error) {
	out := new(GetProductsListResponse)
	err := c.cc.Invoke(ctx, "/services.productsCategories_service.ProductsCategoriesService/GetProductsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsCategoriesServiceClient) GetProductsListByFilters(ctx context.Context, in *GetProductsListByFiltersRequest, opts ...grpc.CallOption) (*GetProductsByFiltersResponse, error) {
	out := new(GetProductsByFiltersResponse)
	err := c.cc.Invoke(ctx, "/services.productsCategories_service.ProductsCategoriesService/GetProductsListByFilters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsCategoriesServiceClient) GetCompanyProducts(ctx context.Context, in *GetCompanyProductsRequest, opts ...grpc.CallOption) (*GetProductsListResponse, error) {
	out := new(GetProductsListResponse)
	err := c.cc.Invoke(ctx, "/services.productsCategories_service.ProductsCategoriesService/GetCompanyProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductsCategoriesServiceServer is the server API for ProductsCategoriesService service.
// All implementations must embed UnimplementedProductsCategoriesServiceServer
// for forward compatibility
type ProductsCategoriesServiceServer interface {
	AddProduct(context.Context, *AddProductRequest) (*GetProduct, error)
	GetCategoryById(context.Context, *GetCategoryByID) (*GetCategory, error)
	GetProductById(context.Context, *GetProductByID) (*GetProduct, error)
	SearchCategories(context.Context, *SearchItemNameWithSkipLimitRequest) (*GetCategories, error)
	SearchProducts(context.Context, *SearchItemNameWithSkipLimitRequest) (*GetProductsListResponse, error)
	GetProductsList(context.Context, *GetProductsListRequest) (*GetProductsListResponse, error)
	GetProductsListByFilters(context.Context, *GetProductsListByFiltersRequest) (*GetProductsByFiltersResponse, error)
	GetCompanyProducts(context.Context, *GetCompanyProductsRequest) (*GetProductsListResponse, error)
	mustEmbedUnimplementedProductsCategoriesServiceServer()
}

// UnimplementedProductsCategoriesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductsCategoriesServiceServer struct {
}

func (UnimplementedProductsCategoriesServiceServer) AddProduct(context.Context, *AddProductRequest) (*GetProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProduct not implemented")
}
func (UnimplementedProductsCategoriesServiceServer) GetCategoryById(context.Context, *GetCategoryByID) (*GetCategory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryById not implemented")
}
func (UnimplementedProductsCategoriesServiceServer) GetProductById(context.Context, *GetProductByID) (*GetProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductById not implemented")
}
func (UnimplementedProductsCategoriesServiceServer) SearchCategories(context.Context, *SearchItemNameWithSkipLimitRequest) (*GetCategories, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCategories not implemented")
}
func (UnimplementedProductsCategoriesServiceServer) SearchProducts(context.Context, *SearchItemNameWithSkipLimitRequest) (*GetProductsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchProducts not implemented")
}
func (UnimplementedProductsCategoriesServiceServer) GetProductsList(context.Context, *GetProductsListRequest) (*GetProductsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsList not implemented")
}
func (UnimplementedProductsCategoriesServiceServer) GetProductsListByFilters(context.Context, *GetProductsListByFiltersRequest) (*GetProductsByFiltersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsListByFilters not implemented")
}
func (UnimplementedProductsCategoriesServiceServer) GetCompanyProducts(context.Context, *GetCompanyProductsRequest) (*GetProductsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyProducts not implemented")
}
func (UnimplementedProductsCategoriesServiceServer) mustEmbedUnimplementedProductsCategoriesServiceServer() {
}

// UnsafeProductsCategoriesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductsCategoriesServiceServer will
// result in compilation errors.
type UnsafeProductsCategoriesServiceServer interface {
	mustEmbedUnimplementedProductsCategoriesServiceServer()
}

func RegisterProductsCategoriesServiceServer(s grpc.ServiceRegistrar, srv ProductsCategoriesServiceServer) {
	s.RegisterService(&ProductsCategoriesService_ServiceDesc, srv)
}

func _ProductsCategoriesService_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsCategoriesServiceServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.productsCategories_service.ProductsCategoriesService/AddProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsCategoriesServiceServer).AddProduct(ctx, req.(*AddProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductsCategoriesService_GetCategoryById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsCategoriesServiceServer).GetCategoryById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.productsCategories_service.ProductsCategoriesService/GetCategoryById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsCategoriesServiceServer).GetCategoryById(ctx, req.(*GetCategoryByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductsCategoriesService_GetProductById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsCategoriesServiceServer).GetProductById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.productsCategories_service.ProductsCategoriesService/GetProductById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsCategoriesServiceServer).GetProductById(ctx, req.(*GetProductByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductsCategoriesService_SearchCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchItemNameWithSkipLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsCategoriesServiceServer).SearchCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.productsCategories_service.ProductsCategoriesService/SearchCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsCategoriesServiceServer).SearchCategories(ctx, req.(*SearchItemNameWithSkipLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductsCategoriesService_SearchProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchItemNameWithSkipLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsCategoriesServiceServer).SearchProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.productsCategories_service.ProductsCategoriesService/SearchProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsCategoriesServiceServer).SearchProducts(ctx, req.(*SearchItemNameWithSkipLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductsCategoriesService_GetProductsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsCategoriesServiceServer).GetProductsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.productsCategories_service.ProductsCategoriesService/GetProductsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsCategoriesServiceServer).GetProductsList(ctx, req.(*GetProductsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductsCategoriesService_GetProductsListByFilters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductsListByFiltersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsCategoriesServiceServer).GetProductsListByFilters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.productsCategories_service.ProductsCategoriesService/GetProductsListByFilters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsCategoriesServiceServer).GetProductsListByFilters(ctx, req.(*GetProductsListByFiltersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductsCategoriesService_GetCompanyProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsCategoriesServiceServer).GetCompanyProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.productsCategories_service.ProductsCategoriesService/GetCompanyProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsCategoriesServiceServer).GetCompanyProducts(ctx, req.(*GetCompanyProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductsCategoriesService_ServiceDesc is the grpc.ServiceDesc for ProductsCategoriesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductsCategoriesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.productsCategories_service.ProductsCategoriesService",
	HandlerType: (*ProductsCategoriesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddProduct",
			Handler:    _ProductsCategoriesService_AddProduct_Handler,
		},
		{
			MethodName: "GetCategoryById",
			Handler:    _ProductsCategoriesService_GetCategoryById_Handler,
		},
		{
			MethodName: "GetProductById",
			Handler:    _ProductsCategoriesService_GetProductById_Handler,
		},
		{
			MethodName: "SearchCategories",
			Handler:    _ProductsCategoriesService_SearchCategories_Handler,
		},
		{
			MethodName: "SearchProducts",
			Handler:    _ProductsCategoriesService_SearchProducts_Handler,
		},
		{
			MethodName: "GetProductsList",
			Handler:    _ProductsCategoriesService_GetProductsList_Handler,
		},
		{
			MethodName: "GetProductsListByFilters",
			Handler:    _ProductsCategoriesService_GetProductsListByFilters_Handler,
		},
		{
			MethodName: "GetCompanyProducts",
			Handler:    _ProductsCategoriesService_GetCompanyProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/services/productsCategories/api.proto",
}
