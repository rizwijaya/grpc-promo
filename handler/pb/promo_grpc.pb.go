// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: promo.proto

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

// PromoServiceClient is the client API for PromoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PromoServiceClient interface {
	GetPromoByCode(ctx context.Context, in *Code, opts ...grpc.CallOption) (*Promo, error)
	GetOrderList(ctx context.Context, in *Pages, opts ...grpc.CallOption) (*Orders, error)
	GetOrderDetail(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*Order, error)
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*Order, error)
}

type promoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPromoServiceClient(cc grpc.ClientConnInterface) PromoServiceClient {
	return &promoServiceClient{cc}
}

func (c *promoServiceClient) GetPromoByCode(ctx context.Context, in *Code, opts ...grpc.CallOption) (*Promo, error) {
	out := new(Promo)
	err := c.cc.Invoke(ctx, "/promo.PromoService/GetPromoByCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promoServiceClient) GetOrderList(ctx context.Context, in *Pages, opts ...grpc.CallOption) (*Orders, error) {
	out := new(Orders)
	err := c.cc.Invoke(ctx, "/promo.PromoService/GetOrderList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promoServiceClient) GetOrderDetail(ctx context.Context, in *OrderId, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/promo.PromoService/GetOrderDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promoServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/promo.PromoService/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PromoServiceServer is the server API for PromoService service.
// All implementations must embed UnimplementedPromoServiceServer
// for forward compatibility
type PromoServiceServer interface {
	GetPromoByCode(context.Context, *Code) (*Promo, error)
	GetOrderList(context.Context, *Pages) (*Orders, error)
	GetOrderDetail(context.Context, *OrderId) (*Order, error)
	CreateOrder(context.Context, *CreateOrderRequest) (*Order, error)
	mustEmbedUnimplementedPromoServiceServer()
}

// UnimplementedPromoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPromoServiceServer struct {
}

func (UnimplementedPromoServiceServer) GetPromoByCode(context.Context, *Code) (*Promo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPromoByCode not implemented")
}
func (UnimplementedPromoServiceServer) GetOrderList(context.Context, *Pages) (*Orders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderList not implemented")
}
func (UnimplementedPromoServiceServer) GetOrderDetail(context.Context, *OrderId) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderDetail not implemented")
}
func (UnimplementedPromoServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedPromoServiceServer) mustEmbedUnimplementedPromoServiceServer() {}

// UnsafePromoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PromoServiceServer will
// result in compilation errors.
type UnsafePromoServiceServer interface {
	mustEmbedUnimplementedPromoServiceServer()
}

func RegisterPromoServiceServer(s grpc.ServiceRegistrar, srv PromoServiceServer) {
	s.RegisterService(&PromoService_ServiceDesc, srv)
}

func _PromoService_GetPromoByCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Code)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromoServiceServer).GetPromoByCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/promo.PromoService/GetPromoByCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromoServiceServer).GetPromoByCode(ctx, req.(*Code))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromoService_GetOrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pages)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromoServiceServer).GetOrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/promo.PromoService/GetOrderList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromoServiceServer).GetOrderList(ctx, req.(*Pages))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromoService_GetOrderDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromoServiceServer).GetOrderDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/promo.PromoService/GetOrderDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromoServiceServer).GetOrderDetail(ctx, req.(*OrderId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromoService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromoServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/promo.PromoService/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromoServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PromoService_ServiceDesc is the grpc.ServiceDesc for PromoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PromoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "promo.PromoService",
	HandlerType: (*PromoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPromoByCode",
			Handler:    _PromoService_GetPromoByCode_Handler,
		},
		{
			MethodName: "GetOrderList",
			Handler:    _PromoService_GetOrderList_Handler,
		},
		{
			MethodName: "GetOrderDetail",
			Handler:    _PromoService_GetOrderDetail_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _PromoService_CreateOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "promo.proto",
}
