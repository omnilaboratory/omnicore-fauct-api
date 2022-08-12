// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: toolrpc/proto.proto

package toolrpc

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

// ToolsClient is the client API for Tools service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ToolsClient interface {
	GetBalance(ctx context.Context, in *OmniGetbalanceReq, opts ...grpc.CallOption) (*OmniGetbalanceRes, error)
	// this command will send user    one bitcoin  and 100 asset, and mine 3 blocks
	SendCoin(ctx context.Context, in *OmniSendCoinReq, opts ...grpc.CallOption) (*OmniSendCoinRes, error)
	//mine three block
	Mine(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*OmniMineCoinRes, error)
	//query  Property
	GetProperty(ctx context.Context, in *OmniGetPropertyReq, opts ...grpc.CallOption) (*OmniGetPropertyRes, error)
	//list all asset
	ListProperties(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListPropertiesRes, error)
	//create asset  cmd: omni_sendissuancefixed "mtowceAw2yeftR1pPg15QcsDqsnSik7Spz" 2 2 0 "t1" "" "ftoken" "baidu.com" "" 10000000000
	CreateProperty(ctx context.Context, in *CreatePropertyReq, opts ...grpc.CallOption) (*CreatePropertyRes, error)
}

type toolsClient struct {
	cc grpc.ClientConnInterface
}

func NewToolsClient(cc grpc.ClientConnInterface) ToolsClient {
	return &toolsClient{cc}
}

func (c *toolsClient) GetBalance(ctx context.Context, in *OmniGetbalanceReq, opts ...grpc.CallOption) (*OmniGetbalanceRes, error) {
	out := new(OmniGetbalanceRes)
	err := c.cc.Invoke(ctx, "/toolrpc.tools/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toolsClient) SendCoin(ctx context.Context, in *OmniSendCoinReq, opts ...grpc.CallOption) (*OmniSendCoinRes, error) {
	out := new(OmniSendCoinRes)
	err := c.cc.Invoke(ctx, "/toolrpc.tools/SendCoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toolsClient) Mine(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*OmniMineCoinRes, error) {
	out := new(OmniMineCoinRes)
	err := c.cc.Invoke(ctx, "/toolrpc.tools/Mine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toolsClient) GetProperty(ctx context.Context, in *OmniGetPropertyReq, opts ...grpc.CallOption) (*OmniGetPropertyRes, error) {
	out := new(OmniGetPropertyRes)
	err := c.cc.Invoke(ctx, "/toolrpc.tools/GetProperty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toolsClient) ListProperties(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListPropertiesRes, error) {
	out := new(ListPropertiesRes)
	err := c.cc.Invoke(ctx, "/toolrpc.tools/ListProperties", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toolsClient) CreateProperty(ctx context.Context, in *CreatePropertyReq, opts ...grpc.CallOption) (*CreatePropertyRes, error) {
	out := new(CreatePropertyRes)
	err := c.cc.Invoke(ctx, "/toolrpc.tools/CreateProperty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToolsServer is the server API for Tools service.
// All implementations must embed UnimplementedToolsServer
// for forward compatibility
type ToolsServer interface {
	GetBalance(context.Context, *OmniGetbalanceReq) (*OmniGetbalanceRes, error)
	// this command will send user    one bitcoin  and 100 asset, and mine 3 blocks
	SendCoin(context.Context, *OmniSendCoinReq) (*OmniSendCoinRes, error)
	//mine three block
	Mine(context.Context, *emptypb.Empty) (*OmniMineCoinRes, error)
	//query  Property
	GetProperty(context.Context, *OmniGetPropertyReq) (*OmniGetPropertyRes, error)
	//list all asset
	ListProperties(context.Context, *emptypb.Empty) (*ListPropertiesRes, error)
	//create asset  cmd: omni_sendissuancefixed "mtowceAw2yeftR1pPg15QcsDqsnSik7Spz" 2 2 0 "t1" "" "ftoken" "baidu.com" "" 10000000000
	CreateProperty(context.Context, *CreatePropertyReq) (*CreatePropertyRes, error)
	mustEmbedUnimplementedToolsServer()
}

// UnimplementedToolsServer must be embedded to have forward compatible implementations.
type UnimplementedToolsServer struct {
}

func (UnimplementedToolsServer) GetBalance(context.Context, *OmniGetbalanceReq) (*OmniGetbalanceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedToolsServer) SendCoin(context.Context, *OmniSendCoinReq) (*OmniSendCoinRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCoin not implemented")
}
func (UnimplementedToolsServer) Mine(context.Context, *emptypb.Empty) (*OmniMineCoinRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mine not implemented")
}
func (UnimplementedToolsServer) GetProperty(context.Context, *OmniGetPropertyReq) (*OmniGetPropertyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProperty not implemented")
}
func (UnimplementedToolsServer) ListProperties(context.Context, *emptypb.Empty) (*ListPropertiesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProperties not implemented")
}
func (UnimplementedToolsServer) CreateProperty(context.Context, *CreatePropertyReq) (*CreatePropertyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProperty not implemented")
}
func (UnimplementedToolsServer) mustEmbedUnimplementedToolsServer() {}

// UnsafeToolsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ToolsServer will
// result in compilation errors.
type UnsafeToolsServer interface {
	mustEmbedUnimplementedToolsServer()
}

func RegisterToolsServer(s grpc.ServiceRegistrar, srv ToolsServer) {
	s.RegisterService(&Tools_ServiceDesc, srv)
}

func _Tools_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OmniGetbalanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToolsServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toolrpc.tools/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToolsServer).GetBalance(ctx, req.(*OmniGetbalanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tools_SendCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OmniSendCoinReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToolsServer).SendCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toolrpc.tools/SendCoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToolsServer).SendCoin(ctx, req.(*OmniSendCoinReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tools_Mine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToolsServer).Mine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toolrpc.tools/Mine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToolsServer).Mine(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tools_GetProperty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OmniGetPropertyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToolsServer).GetProperty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toolrpc.tools/GetProperty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToolsServer).GetProperty(ctx, req.(*OmniGetPropertyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tools_ListProperties_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToolsServer).ListProperties(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toolrpc.tools/ListProperties",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToolsServer).ListProperties(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tools_CreateProperty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePropertyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToolsServer).CreateProperty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toolrpc.tools/CreateProperty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToolsServer).CreateProperty(ctx, req.(*CreatePropertyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Tools_ServiceDesc is the grpc.ServiceDesc for Tools service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tools_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "toolrpc.tools",
	HandlerType: (*ToolsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBalance",
			Handler:    _Tools_GetBalance_Handler,
		},
		{
			MethodName: "SendCoin",
			Handler:    _Tools_SendCoin_Handler,
		},
		{
			MethodName: "Mine",
			Handler:    _Tools_Mine_Handler,
		},
		{
			MethodName: "GetProperty",
			Handler:    _Tools_GetProperty_Handler,
		},
		{
			MethodName: "ListProperties",
			Handler:    _Tools_ListProperties_Handler,
		},
		{
			MethodName: "CreateProperty",
			Handler:    _Tools_CreateProperty_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "toolrpc/proto.proto",
}
