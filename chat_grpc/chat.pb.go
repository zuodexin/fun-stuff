// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: chat.proto

package main

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ChatInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ChatInfo) Reset() {
	*x = ChatInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatInfo) ProtoMessage() {}

func (x *ChatInfo) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatInfo.ProtoReflect.Descriptor instead.
func (*ChatInfo) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *ChatInfo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x08,
	0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x32, 0xc1, 0x01, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x43, 0x68,
	0x61, 0x74, 0x12, 0x09, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x09, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x30, 0x01, 0x12, 0x2f, 0x0a, 0x13,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x43,
	0x68, 0x61, 0x74, 0x12, 0x09, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x09,
	0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x28, 0x01, 0x12, 0x2e, 0x0a,
	0x10, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x68, 0x61,
	0x74, 0x12, 0x09, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x09, 0x2e, 0x43,
	0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x27, 0x0a,
	0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x68, 0x61, 0x74, 0x12, 0x09, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x09, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x6d, 0x61, 0x69, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_chat_proto_goTypes = []interface{}{
	(*ChatInfo)(nil), // 0: ChatInfo
}
var file_chat_proto_depIdxs = []int32{
	0, // 0: Chat.ServerStreamingChat:input_type -> ChatInfo
	0, // 1: Chat.ClientStreamingChat:input_type -> ChatInfo
	0, // 2: Chat.ClientServerChat:input_type -> ChatInfo
	0, // 3: Chat.GroupChat:input_type -> ChatInfo
	0, // 4: Chat.ServerStreamingChat:output_type -> ChatInfo
	0, // 5: Chat.ClientStreamingChat:output_type -> ChatInfo
	0, // 6: Chat.ClientServerChat:output_type -> ChatInfo
	0, // 7: Chat.GroupChat:output_type -> ChatInfo
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatClient interface {
	ServerStreamingChat(ctx context.Context, in *ChatInfo, opts ...grpc.CallOption) (Chat_ServerStreamingChatClient, error)
	ClientStreamingChat(ctx context.Context, opts ...grpc.CallOption) (Chat_ClientStreamingChatClient, error)
	ClientServerChat(ctx context.Context, opts ...grpc.CallOption) (Chat_ClientServerChatClient, error)
	GroupChat(ctx context.Context, opts ...grpc.CallOption) (Chat_GroupChatClient, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) ServerStreamingChat(ctx context.Context, in *ChatInfo, opts ...grpc.CallOption) (Chat_ServerStreamingChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chat_serviceDesc.Streams[0], "/Chat/ServerStreamingChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServerStreamingChatClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chat_ServerStreamingChatClient interface {
	Recv() (*ChatInfo, error)
	grpc.ClientStream
}

type chatServerStreamingChatClient struct {
	grpc.ClientStream
}

func (x *chatServerStreamingChatClient) Recv() (*ChatInfo, error) {
	m := new(ChatInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatClient) ClientStreamingChat(ctx context.Context, opts ...grpc.CallOption) (Chat_ClientStreamingChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chat_serviceDesc.Streams[1], "/Chat/ClientStreamingChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatClientStreamingChatClient{stream}
	return x, nil
}

type Chat_ClientStreamingChatClient interface {
	Send(*ChatInfo) error
	CloseAndRecv() (*ChatInfo, error)
	grpc.ClientStream
}

type chatClientStreamingChatClient struct {
	grpc.ClientStream
}

func (x *chatClientStreamingChatClient) Send(m *ChatInfo) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatClientStreamingChatClient) CloseAndRecv() (*ChatInfo, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ChatInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatClient) ClientServerChat(ctx context.Context, opts ...grpc.CallOption) (Chat_ClientServerChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chat_serviceDesc.Streams[2], "/Chat/ClientServerChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatClientServerChatClient{stream}
	return x, nil
}

type Chat_ClientServerChatClient interface {
	Send(*ChatInfo) error
	Recv() (*ChatInfo, error)
	grpc.ClientStream
}

type chatClientServerChatClient struct {
	grpc.ClientStream
}

func (x *chatClientServerChatClient) Send(m *ChatInfo) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatClientServerChatClient) Recv() (*ChatInfo, error) {
	m := new(ChatInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatClient) GroupChat(ctx context.Context, opts ...grpc.CallOption) (Chat_GroupChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chat_serviceDesc.Streams[3], "/Chat/GroupChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatGroupChatClient{stream}
	return x, nil
}

type Chat_GroupChatClient interface {
	Send(*ChatInfo) error
	Recv() (*ChatInfo, error)
	grpc.ClientStream
}

type chatGroupChatClient struct {
	grpc.ClientStream
}

func (x *chatGroupChatClient) Send(m *ChatInfo) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatGroupChatClient) Recv() (*ChatInfo, error) {
	m := new(ChatInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServer is the server API for Chat service.
type ChatServer interface {
	ServerStreamingChat(*ChatInfo, Chat_ServerStreamingChatServer) error
	ClientStreamingChat(Chat_ClientStreamingChatServer) error
	ClientServerChat(Chat_ClientServerChatServer) error
	GroupChat(Chat_GroupChatServer) error
}

// UnimplementedChatServer can be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (*UnimplementedChatServer) ServerStreamingChat(*ChatInfo, Chat_ServerStreamingChatServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreamingChat not implemented")
}
func (*UnimplementedChatServer) ClientStreamingChat(Chat_ClientStreamingChatServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreamingChat not implemented")
}
func (*UnimplementedChatServer) ClientServerChat(Chat_ClientServerChatServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientServerChat not implemented")
}
func (*UnimplementedChatServer) GroupChat(Chat_GroupChatServer) error {
	return status.Errorf(codes.Unimplemented, "method GroupChat not implemented")
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_ServerStreamingChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ChatInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServer).ServerStreamingChat(m, &chatServerStreamingChatServer{stream})
}

type Chat_ServerStreamingChatServer interface {
	Send(*ChatInfo) error
	grpc.ServerStream
}

type chatServerStreamingChatServer struct {
	grpc.ServerStream
}

func (x *chatServerStreamingChatServer) Send(m *ChatInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _Chat_ClientStreamingChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).ClientStreamingChat(&chatClientStreamingChatServer{stream})
}

type Chat_ClientStreamingChatServer interface {
	SendAndClose(*ChatInfo) error
	Recv() (*ChatInfo, error)
	grpc.ServerStream
}

type chatClientStreamingChatServer struct {
	grpc.ServerStream
}

func (x *chatClientStreamingChatServer) SendAndClose(m *ChatInfo) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatClientStreamingChatServer) Recv() (*ChatInfo, error) {
	m := new(ChatInfo)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Chat_ClientServerChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).ClientServerChat(&chatClientServerChatServer{stream})
}

type Chat_ClientServerChatServer interface {
	Send(*ChatInfo) error
	Recv() (*ChatInfo, error)
	grpc.ServerStream
}

type chatClientServerChatServer struct {
	grpc.ServerStream
}

func (x *chatClientServerChatServer) Send(m *ChatInfo) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatClientServerChatServer) Recv() (*ChatInfo, error) {
	m := new(ChatInfo)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Chat_GroupChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).GroupChat(&chatGroupChatServer{stream})
}

type Chat_GroupChatServer interface {
	Send(*ChatInfo) error
	Recv() (*ChatInfo, error)
	grpc.ServerStream
}

type chatGroupChatServer struct {
	grpc.ServerStream
}

func (x *chatGroupChatServer) Send(m *ChatInfo) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatGroupChatServer) Recv() (*ChatInfo, error) {
	m := new(ChatInfo)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Chat",
	HandlerType: (*ChatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStreamingChat",
			Handler:       _Chat_ServerStreamingChat_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStreamingChat",
			Handler:       _Chat_ClientStreamingChat_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ClientServerChat",
			Handler:       _Chat_ClientServerChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GroupChat",
			Handler:       _Chat_GroupChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}
