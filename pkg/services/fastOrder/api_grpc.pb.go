// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: pkg/services/fastOrder/api.proto

package fastOrder_service

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

// FastOrderServiceClient is the client API for FastOrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FastOrderServiceClient interface {
	FastOrder(ctx context.Context, in *FastOrderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type fastOrderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFastOrderServiceClient(cc grpc.ClientConnInterface) FastOrderServiceClient {
	return &fastOrderServiceClient{cc}
}

func (c *fastOrderServiceClient) FastOrder(ctx context.Context, in *FastOrderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/services.fastOrder_service.FastOrderService/FastOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FastOrderServiceServer is the server API for FastOrderService service.
// All implementations must embed UnimplementedFastOrderServiceServer
// for forward compatibility
type FastOrderServiceServer interface {
	FastOrder(context.Context, *FastOrderRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedFastOrderServiceServer()
}

// UnimplementedFastOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFastOrderServiceServer struct {
}

func (UnimplementedFastOrderServiceServer) FastOrder(context.Context, *FastOrderRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FastOrder not implemented")
}
func (UnimplementedFastOrderServiceServer) mustEmbedUnimplementedFastOrderServiceServer() {}

// UnsafeFastOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FastOrderServiceServer will
// result in compilation errors.
type UnsafeFastOrderServiceServer interface {
	mustEmbedUnimplementedFastOrderServiceServer()
}

func RegisterFastOrderServiceServer(s grpc.ServiceRegistrar, srv FastOrderServiceServer) {
	s.RegisterService(&FastOrderService_ServiceDesc, srv)
}

func _FastOrderService_FastOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FastOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FastOrderServiceServer).FastOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.fastOrder_service.FastOrderService/FastOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FastOrderServiceServer).FastOrder(ctx, req.(*FastOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FastOrderService_ServiceDesc is the grpc.ServiceDesc for FastOrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FastOrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.fastOrder_service.FastOrderService",
	HandlerType: (*FastOrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FastOrder",
			Handler:    _FastOrderService_FastOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/services/fastOrder/api.proto",
}
