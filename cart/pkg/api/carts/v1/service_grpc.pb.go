// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: api/carts/v1/service.proto

package carts

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

const (
	Carts_ItemAdd_FullMethodName    = "/gitlab.ozon.dev.go_8_middle_project.cart.pkg.api.carts.v1.Carts/ItemAdd"
	Carts_ItemDelete_FullMethodName = "/gitlab.ozon.dev.go_8_middle_project.cart.pkg.api.carts.v1.Carts/ItemDelete"
	Carts_Clear_FullMethodName      = "/gitlab.ozon.dev.go_8_middle_project.cart.pkg.api.carts.v1.Carts/Clear"
	Carts_Checkout_FullMethodName   = "/gitlab.ozon.dev.go_8_middle_project.cart.pkg.api.carts.v1.Carts/Checkout"
	Carts_List_FullMethodName       = "/gitlab.ozon.dev.go_8_middle_project.cart.pkg.api.carts.v1.Carts/List"
)

// CartsClient is the client API for Carts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartsClient interface {
	ItemAdd(ctx context.Context, in *ItemAddRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ItemDelete(ctx context.Context, in *ItemDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Clear(ctx context.Context, in *ClearRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Checkout(ctx context.Context, in *CheckoutRequest, opts ...grpc.CallOption) (*CheckoutResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type cartsClient struct {
	cc grpc.ClientConnInterface
}

func NewCartsClient(cc grpc.ClientConnInterface) CartsClient {
	return &cartsClient{cc}
}

func (c *cartsClient) ItemAdd(ctx context.Context, in *ItemAddRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Carts_ItemAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartsClient) ItemDelete(ctx context.Context, in *ItemDeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Carts_ItemDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartsClient) Clear(ctx context.Context, in *ClearRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Carts_Clear_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartsClient) Checkout(ctx context.Context, in *CheckoutRequest, opts ...grpc.CallOption) (*CheckoutResponse, error) {
	out := new(CheckoutResponse)
	err := c.cc.Invoke(ctx, Carts_Checkout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartsClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, Carts_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartsServer is the server API for Carts service.
// All implementations must embed UnimplementedCartsServer
// for forward compatibility
type CartsServer interface {
	ItemAdd(context.Context, *ItemAddRequest) (*emptypb.Empty, error)
	ItemDelete(context.Context, *ItemDeleteRequest) (*emptypb.Empty, error)
	Clear(context.Context, *ClearRequest) (*emptypb.Empty, error)
	Checkout(context.Context, *CheckoutRequest) (*CheckoutResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	mustEmbedUnimplementedCartsServer()
}

// UnimplementedCartsServer must be embedded to have forward compatible implementations.
type UnimplementedCartsServer struct {
}

func (UnimplementedCartsServer) ItemAdd(context.Context, *ItemAddRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ItemAdd not implemented")
}
func (UnimplementedCartsServer) ItemDelete(context.Context, *ItemDeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ItemDelete not implemented")
}
func (UnimplementedCartsServer) Clear(context.Context, *ClearRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clear not implemented")
}
func (UnimplementedCartsServer) Checkout(context.Context, *CheckoutRequest) (*CheckoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Checkout not implemented")
}
func (UnimplementedCartsServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCartsServer) mustEmbedUnimplementedCartsServer() {}

// UnsafeCartsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartsServer will
// result in compilation errors.
type UnsafeCartsServer interface {
	mustEmbedUnimplementedCartsServer()
}

func RegisterCartsServer(s grpc.ServiceRegistrar, srv CartsServer) {
	s.RegisterService(&Carts_ServiceDesc, srv)
}

func _Carts_ItemAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartsServer).ItemAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Carts_ItemAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartsServer).ItemAdd(ctx, req.(*ItemAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Carts_ItemDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartsServer).ItemDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Carts_ItemDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartsServer).ItemDelete(ctx, req.(*ItemDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Carts_Clear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartsServer).Clear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Carts_Clear_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartsServer).Clear(ctx, req.(*ClearRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Carts_Checkout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartsServer).Checkout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Carts_Checkout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartsServer).Checkout(ctx, req.(*CheckoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Carts_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Carts_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartsServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Carts_ServiceDesc is the grpc.ServiceDesc for Carts service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Carts_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gitlab.ozon.dev.go_8_middle_project.cart.pkg.api.carts.v1.Carts",
	HandlerType: (*CartsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ItemAdd",
			Handler:    _Carts_ItemAdd_Handler,
		},
		{
			MethodName: "ItemDelete",
			Handler:    _Carts_ItemDelete_Handler,
		},
		{
			MethodName: "Clear",
			Handler:    _Carts_Clear_Handler,
		},
		{
			MethodName: "Checkout",
			Handler:    _Carts_Checkout_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Carts_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/carts/v1/service.proto",
}
