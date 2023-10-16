// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.3
// source: api/proto/balance_ms.proto

package api

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

// BalanceMsClient is the client API for BalanceMs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BalanceMsClient interface {
	AddBalance(ctx context.Context, in *AddBalanceRequest, opts ...grpc.CallOption) (*AddBalanceResponse, error)
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	Reserve(ctx context.Context, in *ReserveRequest, opts ...grpc.CallOption) (*ReserveResponse, error)
	GetRevenue(ctx context.Context, in *GetRevenueRequest, opts ...grpc.CallOption) (*GetRevenueResponse, error)
}

type balanceMsClient struct {
	cc grpc.ClientConnInterface
}

func NewBalanceMsClient(cc grpc.ClientConnInterface) BalanceMsClient {
	return &balanceMsClient{cc}
}

func (c *balanceMsClient) AddBalance(ctx context.Context, in *AddBalanceRequest, opts ...grpc.CallOption) (*AddBalanceResponse, error) {
	out := new(AddBalanceResponse)
	err := c.cc.Invoke(ctx, "/api.Balance_ms/AddBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balanceMsClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, "/api.Balance_ms/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balanceMsClient) Reserve(ctx context.Context, in *ReserveRequest, opts ...grpc.CallOption) (*ReserveResponse, error) {
	out := new(ReserveResponse)
	err := c.cc.Invoke(ctx, "/api.Balance_ms/Reserve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balanceMsClient) GetRevenue(ctx context.Context, in *GetRevenueRequest, opts ...grpc.CallOption) (*GetRevenueResponse, error) {
	out := new(GetRevenueResponse)
	err := c.cc.Invoke(ctx, "/api.Balance_ms/GetRevenue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BalanceMsServer is the server API for BalanceMs service.
// All implementations must embed UnimplementedBalanceMsServer
// for forward compatibility
type BalanceMsServer interface {
	AddBalance(context.Context, *AddBalanceRequest) (*AddBalanceResponse, error)
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	Reserve(context.Context, *ReserveRequest) (*ReserveResponse, error)
	GetRevenue(context.Context, *GetRevenueRequest) (*GetRevenueResponse, error)
	mustEmbedUnimplementedBalanceMsServer()
}

// UnimplementedBalanceMsServer must be embedded to have forward compatible implementations.
type UnimplementedBalanceMsServer struct {
}

func (UnimplementedBalanceMsServer) AddBalance(context.Context, *AddBalanceRequest) (*AddBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBalance not implemented")
}
func (UnimplementedBalanceMsServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedBalanceMsServer) Reserve(context.Context, *ReserveRequest) (*ReserveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reserve not implemented")
}
func (UnimplementedBalanceMsServer) GetRevenue(context.Context, *GetRevenueRequest) (*GetRevenueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRevenue not implemented")
}
func (UnimplementedBalanceMsServer) mustEmbedUnimplementedBalanceMsServer() {}

// UnsafeBalanceMsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BalanceMsServer will
// result in compilation errors.
type UnsafeBalanceMsServer interface {
	mustEmbedUnimplementedBalanceMsServer()
}

func RegisterBalanceMsServer(s grpc.ServiceRegistrar, srv BalanceMsServer) {
	s.RegisterService(&BalanceMs_ServiceDesc, srv)
}

func _BalanceMs_AddBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceMsServer).AddBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Balance_ms/AddBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceMsServer).AddBalance(ctx, req.(*AddBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BalanceMs_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceMsServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Balance_ms/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceMsServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BalanceMs_Reserve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReserveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceMsServer).Reserve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Balance_ms/Reserve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceMsServer).Reserve(ctx, req.(*ReserveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BalanceMs_GetRevenue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRevenueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceMsServer).GetRevenue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Balance_ms/GetRevenue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceMsServer).GetRevenue(ctx, req.(*GetRevenueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BalanceMs_ServiceDesc is the grpc.ServiceDesc for BalanceMs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BalanceMs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Balance_ms",
	HandlerType: (*BalanceMsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBalance",
			Handler:    _BalanceMs_AddBalance_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _BalanceMs_GetBalance_Handler,
		},
		{
			MethodName: "Reserve",
			Handler:    _BalanceMs_Reserve_Handler,
		},
		{
			MethodName: "GetRevenue",
			Handler:    _BalanceMs_GetRevenue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/balance_ms.proto",
}