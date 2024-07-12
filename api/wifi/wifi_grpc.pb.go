// Copyright 2024 TII (SSRC) and the Ghaf contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: wifi.proto

package wifi

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

const (
	WifiService_ListNetwork_FullMethodName       = "/wifimanager.WifiService/ListNetwork"
	WifiService_ConnectNetwork_FullMethodName    = "/wifimanager.WifiService/ConnectNetwork"
	WifiService_DisconnectNetwork_FullMethodName = "/wifimanager.WifiService/DisconnectNetwork"
	WifiService_TurnOn_FullMethodName            = "/wifimanager.WifiService/TurnOn"
	WifiService_TurnOff_FullMethodName           = "/wifimanager.WifiService/TurnOff"
)

// WifiServiceClient is the client API for WifiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WifiServiceClient interface {
	ListNetwork(ctx context.Context, in *WifiNetworkRequest, opts ...grpc.CallOption) (*WifiNetworkResponse, error)
	ConnectNetwork(ctx context.Context, in *WifiConnectionRequest, opts ...grpc.CallOption) (*WifiConnectionResponse, error)
	DisconnectNetwork(ctx context.Context, in *WifiNetworkRequest, opts ...grpc.CallOption) (*WifiConnectionResponse, error)
	TurnOn(ctx context.Context, in *WifiNetworkRequest, opts ...grpc.CallOption) (*WifiConnectionResponse, error)
	TurnOff(ctx context.Context, in *WifiNetworkRequest, opts ...grpc.CallOption) (*WifiConnectionResponse, error)
}

type wifiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWifiServiceClient(cc grpc.ClientConnInterface) WifiServiceClient {
	return &wifiServiceClient{cc}
}

func (c *wifiServiceClient) ListNetwork(ctx context.Context, in *WifiNetworkRequest, opts ...grpc.CallOption) (*WifiNetworkResponse, error) {
	out := new(WifiNetworkResponse)
	err := c.cc.Invoke(ctx, WifiService_ListNetwork_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wifiServiceClient) ConnectNetwork(ctx context.Context, in *WifiConnectionRequest, opts ...grpc.CallOption) (*WifiConnectionResponse, error) {
	out := new(WifiConnectionResponse)
	err := c.cc.Invoke(ctx, WifiService_ConnectNetwork_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wifiServiceClient) DisconnectNetwork(ctx context.Context, in *WifiNetworkRequest, opts ...grpc.CallOption) (*WifiConnectionResponse, error) {
	out := new(WifiConnectionResponse)
	err := c.cc.Invoke(ctx, WifiService_DisconnectNetwork_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wifiServiceClient) TurnOn(ctx context.Context, in *WifiNetworkRequest, opts ...grpc.CallOption) (*WifiConnectionResponse, error) {
	out := new(WifiConnectionResponse)
	err := c.cc.Invoke(ctx, WifiService_TurnOn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wifiServiceClient) TurnOff(ctx context.Context, in *WifiNetworkRequest, opts ...grpc.CallOption) (*WifiConnectionResponse, error) {
	out := new(WifiConnectionResponse)
	err := c.cc.Invoke(ctx, WifiService_TurnOff_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WifiServiceServer is the server API for WifiService service.
// All implementations must embed UnimplementedWifiServiceServer
// for forward compatibility
type WifiServiceServer interface {
	ListNetwork(context.Context, *WifiNetworkRequest) (*WifiNetworkResponse, error)
	ConnectNetwork(context.Context, *WifiConnectionRequest) (*WifiConnectionResponse, error)
	DisconnectNetwork(context.Context, *WifiNetworkRequest) (*WifiConnectionResponse, error)
	TurnOn(context.Context, *WifiNetworkRequest) (*WifiConnectionResponse, error)
	TurnOff(context.Context, *WifiNetworkRequest) (*WifiConnectionResponse, error)
	mustEmbedUnimplementedWifiServiceServer()
}

// UnimplementedWifiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWifiServiceServer struct {
}

func (UnimplementedWifiServiceServer) ListNetwork(context.Context, *WifiNetworkRequest) (*WifiNetworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNetwork not implemented")
}
func (UnimplementedWifiServiceServer) ConnectNetwork(context.Context, *WifiConnectionRequest) (*WifiConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectNetwork not implemented")
}
func (UnimplementedWifiServiceServer) DisconnectNetwork(context.Context, *WifiNetworkRequest) (*WifiConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisconnectNetwork not implemented")
}
func (UnimplementedWifiServiceServer) TurnOn(context.Context, *WifiNetworkRequest) (*WifiConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TurnOn not implemented")
}
func (UnimplementedWifiServiceServer) TurnOff(context.Context, *WifiNetworkRequest) (*WifiConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TurnOff not implemented")
}
func (UnimplementedWifiServiceServer) mustEmbedUnimplementedWifiServiceServer() {}

// UnsafeWifiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WifiServiceServer will
// result in compilation errors.
type UnsafeWifiServiceServer interface {
	mustEmbedUnimplementedWifiServiceServer()
}

func RegisterWifiServiceServer(s grpc.ServiceRegistrar, srv WifiServiceServer) {
	s.RegisterService(&WifiService_ServiceDesc, srv)
}

func _WifiService_ListNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WifiNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WifiServiceServer).ListNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WifiService_ListNetwork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WifiServiceServer).ListNetwork(ctx, req.(*WifiNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WifiService_ConnectNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WifiConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WifiServiceServer).ConnectNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WifiService_ConnectNetwork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WifiServiceServer).ConnectNetwork(ctx, req.(*WifiConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WifiService_DisconnectNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WifiNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WifiServiceServer).DisconnectNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WifiService_DisconnectNetwork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WifiServiceServer).DisconnectNetwork(ctx, req.(*WifiNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WifiService_TurnOn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WifiNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WifiServiceServer).TurnOn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WifiService_TurnOn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WifiServiceServer).TurnOn(ctx, req.(*WifiNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WifiService_TurnOff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WifiNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WifiServiceServer).TurnOff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WifiService_TurnOff_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WifiServiceServer).TurnOff(ctx, req.(*WifiNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WifiService_ServiceDesc is the grpc.ServiceDesc for WifiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WifiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wifimanager.WifiService",
	HandlerType: (*WifiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListNetwork",
			Handler:    _WifiService_ListNetwork_Handler,
		},
		{
			MethodName: "ConnectNetwork",
			Handler:    _WifiService_ConnectNetwork_Handler,
		},
		{
			MethodName: "DisconnectNetwork",
			Handler:    _WifiService_DisconnectNetwork_Handler,
		},
		{
			MethodName: "TurnOn",
			Handler:    _WifiService_TurnOn_Handler,
		},
		{
			MethodName: "TurnOff",
			Handler:    _WifiService_TurnOff_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wifi.proto",
}
