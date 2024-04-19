// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: chat.proto

package chat_v1

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

// ChatV1Client is the client API for ChatV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatV1Client interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (ChatV1_ConnectClient, error)
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type chatV1Client struct {
	cc grpc.ClientConnInterface
}

func NewChatV1Client(cc grpc.ClientConnInterface) ChatV1Client {
	return &chatV1Client{cc}
}

func (c *chatV1Client) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/ChatV1/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatV1Client) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (ChatV1_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatV1_ServiceDesc.Streams[0], "/ChatV1/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatV1ConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatV1_ConnectClient interface {
	Recv() (*MessageInfo, error)
	grpc.ClientStream
}

type chatV1ConnectClient struct {
	grpc.ClientStream
}

func (x *chatV1ConnectClient) Recv() (*MessageInfo, error) {
	m := new(MessageInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatV1Client) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ChatV1/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatV1Client) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ChatV1/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatV1Server is the server API for ChatV1 service.
// All implementations must embed UnimplementedChatV1Server
// for forward compatibility
type ChatV1Server interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Connect(*ConnectRequest, ChatV1_ConnectServer) error
	SendMessage(context.Context, *SendMessageRequest) (*emptypb.Empty, error)
	Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedChatV1Server()
}

// UnimplementedChatV1Server must be embedded to have forward compatible implementations.
type UnimplementedChatV1Server struct {
}

func (UnimplementedChatV1Server) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedChatV1Server) Connect(*ConnectRequest, ChatV1_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedChatV1Server) SendMessage(context.Context, *SendMessageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatV1Server) Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedChatV1Server) mustEmbedUnimplementedChatV1Server() {}

// UnsafeChatV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatV1Server will
// result in compilation errors.
type UnsafeChatV1Server interface {
	mustEmbedUnimplementedChatV1Server()
}

func RegisterChatV1Server(s grpc.ServiceRegistrar, srv ChatV1Server) {
	s.RegisterService(&ChatV1_ServiceDesc, srv)
}

func _ChatV1_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatV1Server).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatV1/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatV1Server).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatV1_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConnectRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatV1Server).Connect(m, &chatV1ConnectServer{stream})
}

type ChatV1_ConnectServer interface {
	Send(*MessageInfo) error
	grpc.ServerStream
}

type chatV1ConnectServer struct {
	grpc.ServerStream
}

func (x *chatV1ConnectServer) Send(m *MessageInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatV1_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatV1Server).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatV1/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatV1Server).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatV1_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatV1Server).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatV1/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatV1Server).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatV1_ServiceDesc is the grpc.ServiceDesc for ChatV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ChatV1",
	HandlerType: (*ChatV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ChatV1_Create_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _ChatV1_SendMessage_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ChatV1_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _ChatV1_Connect_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat.proto",
}
