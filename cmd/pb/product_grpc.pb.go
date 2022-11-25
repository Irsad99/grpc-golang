// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: product.proto

package pb

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

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	GetProduct(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ResponseProduct, error)
	CreateProduct(ctx context.Context, in *RequestProduct, opts ...grpc.CallOption) (*ResponseProduct, error)
	UpdateProduct(ctx context.Context, in *RequestProduct, opts ...grpc.CallOption) (*ResponseProduct, error)
	DeleteProduct(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ResponseProduct, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) GetProduct(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ResponseProduct, error) {
	out := new(ResponseProduct)
	err := c.cc.Invoke(ctx, "/protos.ProductService/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) CreateProduct(ctx context.Context, in *RequestProduct, opts ...grpc.CallOption) (*ResponseProduct, error) {
	out := new(ResponseProduct)
	err := c.cc.Invoke(ctx, "/protos.ProductService/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateProduct(ctx context.Context, in *RequestProduct, opts ...grpc.CallOption) (*ResponseProduct, error) {
	out := new(ResponseProduct)
	err := c.cc.Invoke(ctx, "/protos.ProductService/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteProduct(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ResponseProduct, error) {
	out := new(ResponseProduct)
	err := c.cc.Invoke(ctx, "/protos.ProductService/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	GetProduct(context.Context, *Id) (*ResponseProduct, error)
	CreateProduct(context.Context, *RequestProduct) (*ResponseProduct, error)
	UpdateProduct(context.Context, *RequestProduct) (*ResponseProduct, error)
	DeleteProduct(context.Context, *Id) (*ResponseProduct, error)
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) GetProduct(context.Context, *Id) (*ResponseProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductServiceServer) CreateProduct(context.Context, *RequestProduct) (*ResponseProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedProductServiceServer) UpdateProduct(context.Context, *RequestProduct) (*ResponseProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedProductServiceServer) DeleteProduct(context.Context, *Id) (*ResponseProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ProductService/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProduct(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestProduct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ProductService/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CreateProduct(ctx, req.(*RequestProduct))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestProduct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ProductService/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateProduct(ctx, req.(*RequestProduct))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ProductService/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteProduct(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProduct",
			Handler:    _ProductService_GetProduct_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _ProductService_CreateProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _ProductService_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _ProductService_DeleteProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
