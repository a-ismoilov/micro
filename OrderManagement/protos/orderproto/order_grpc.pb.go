// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: userproto/order.userproto

package orderproto

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

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	GetOrder(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Order, error)
	Choose(ctx context.Context, in *MealChoice, opts ...grpc.CallOption) (*OK, error)
	Payment(ctx context.Context, in *Request, opts ...grpc.CallOption) (*OK, error)
	MealList(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Meals, error)
	OrderList(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Orders, error)
	Cancel(ctx context.Context, in *Request, opts ...grpc.CallOption) (*OK, error)
	NewMeal(ctx context.Context, in *Meal, opts ...grpc.CallOption) (*OK, error)
	NewOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*OK, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) GetOrder(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/userproto.OrderService/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) Choose(ctx context.Context, in *MealChoice, opts ...grpc.CallOption) (*OK, error) {
	out := new(OK)
	err := c.cc.Invoke(ctx, "/userproto.OrderService/Choose", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) Payment(ctx context.Context, in *Request, opts ...grpc.CallOption) (*OK, error) {
	out := new(OK)
	err := c.cc.Invoke(ctx, "/userproto.OrderService/Payment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) MealList(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Meals, error) {
	out := new(Meals)
	err := c.cc.Invoke(ctx, "/userproto.OrderService/MealList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) OrderList(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Orders, error) {
	out := new(Orders)
	err := c.cc.Invoke(ctx, "/userproto.OrderService/OrderList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) Cancel(ctx context.Context, in *Request, opts ...grpc.CallOption) (*OK, error) {
	out := new(OK)
	err := c.cc.Invoke(ctx, "/userproto.OrderService/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) NewMeal(ctx context.Context, in *Meal, opts ...grpc.CallOption) (*OK, error) {
	out := new(OK)
	err := c.cc.Invoke(ctx, "/userproto.OrderService/NewMeal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) NewOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*OK, error) {
	out := new(OK)
	err := c.cc.Invoke(ctx, "/userproto.OrderService/NewOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	GetOrder(context.Context, *Request) (*Order, error)
	Choose(context.Context, *MealChoice) (*OK, error)
	Payment(context.Context, *Request) (*OK, error)
	MealList(context.Context, *Request) (*Meals, error)
	OrderList(context.Context, *Request) (*Orders, error)
	Cancel(context.Context, *Request) (*OK, error)
	NewMeal(context.Context, *Meal) (*OK, error)
	NewOrder(context.Context, *Order) (*OK, error)
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) GetOrder(context.Context, *Request) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrderServiceServer) Choose(context.Context, *MealChoice) (*OK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Choose not implemented")
}
func (UnimplementedOrderServiceServer) Payment(context.Context, *Request) (*OK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Payment not implemented")
}
func (UnimplementedOrderServiceServer) MealList(context.Context, *Request) (*Meals, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MealList not implemented")
}
func (UnimplementedOrderServiceServer) OrderList(context.Context, *Request) (*Orders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderList not implemented")
}
func (UnimplementedOrderServiceServer) Cancel(context.Context, *Request) (*OK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (UnimplementedOrderServiceServer) NewMeal(context.Context, *Meal) (*OK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewMeal not implemented")
}
func (UnimplementedOrderServiceServer) NewOrder(context.Context, *Order) (*OK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewOrder not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userproto.OrderService/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrder(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_Choose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MealChoice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).Choose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userproto.OrderService/Choose",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).Choose(ctx, req.(*MealChoice))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_Payment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).Payment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userproto.OrderService/Payment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).Payment(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_MealList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).MealList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userproto.OrderService/MealList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).MealList(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_OrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).OrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userproto.OrderService/OrderList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).OrderList(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userproto.OrderService/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).Cancel(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_NewMeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Meal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).NewMeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userproto.OrderService/NewMeal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).NewMeal(ctx, req.(*Meal))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_NewOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).NewOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userproto.OrderService/NewOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).NewOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userproto.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrder",
			Handler:    _OrderService_GetOrder_Handler,
		},
		{
			MethodName: "Choose",
			Handler:    _OrderService_Choose_Handler,
		},
		{
			MethodName: "Payment",
			Handler:    _OrderService_Payment_Handler,
		},
		{
			MethodName: "MealList",
			Handler:    _OrderService_MealList_Handler,
		},
		{
			MethodName: "OrderList",
			Handler:    _OrderService_OrderList_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _OrderService_Cancel_Handler,
		},
		{
			MethodName: "NewMeal",
			Handler:    _OrderService_NewMeal_Handler,
		},
		{
			MethodName: "NewOrder",
			Handler:    _OrderService_NewOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userproto/order.userproto",
}
