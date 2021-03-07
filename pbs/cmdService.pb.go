// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.5
// source: cmdService.proto

package pbs

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

type TopicMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Msg   string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *TopicMsg) Reset() {
	*x = TopicMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicMsg) ProtoMessage() {}

func (x *TopicMsg) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicMsg.ProtoReflect.Descriptor instead.
func (*TopicMsg) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{0}
}

func (x *TopicMsg) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *TopicMsg) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ShowPeer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (x *ShowPeer) Reset() {
	*x = ShowPeer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowPeer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowPeer) ProtoMessage() {}

func (x *ShowPeer) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowPeer.ProtoReflect.Descriptor instead.
func (*ShowPeer) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{1}
}

func (x *ShowPeer) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

type CommonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CommonResponse) Reset() {
	*x = CommonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResponse) ProtoMessage() {}

func (x *CommonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResponse.ProtoReflect.Descriptor instead.
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{2}
}

func (x *CommonResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_cmdService_proto protoreflect.FileDescriptor

var file_cmdService_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6d, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x70, 0x62, 0x73, 0x22, 0x32, 0x0a, 0x08, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x4d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x20, 0x0a, 0x08, 0x53,
	0x68, 0x6f, 0x77, 0x50, 0x65, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x22, 0x22, 0x0a,
	0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x32, 0x7b, 0x0a, 0x0a, 0x43, 0x6d, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x37, 0x0a, 0x0f, 0x50, 0x32, 0x70, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x4d,
	0x73, 0x67, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x4d, 0x73,
	0x67, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0c, 0x50, 0x32, 0x70, 0x53,
	0x68, 0x6f, 0x77, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x53,
	0x68, 0x6f, 0x77, 0x50, 0x65, 0x65, 0x72, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07,
	0x5a, 0x05, 0x2e, 0x3b, 0x70, 0x62, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmdService_proto_rawDescOnce sync.Once
	file_cmdService_proto_rawDescData = file_cmdService_proto_rawDesc
)

func file_cmdService_proto_rawDescGZIP() []byte {
	file_cmdService_proto_rawDescOnce.Do(func() {
		file_cmdService_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmdService_proto_rawDescData)
	})
	return file_cmdService_proto_rawDescData
}

var file_cmdService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cmdService_proto_goTypes = []interface{}{
	(*TopicMsg)(nil),       // 0: pbs.TopicMsg
	(*ShowPeer)(nil),       // 1: pbs.ShowPeer
	(*CommonResponse)(nil), // 2: pbs.CommonResponse
}
var file_cmdService_proto_depIdxs = []int32{
	0, // 0: pbs.CmdService.P2pSendTopicMsg:input_type -> pbs.TopicMsg
	1, // 1: pbs.CmdService.P2pShowPeers:input_type -> pbs.ShowPeer
	2, // 2: pbs.CmdService.P2pSendTopicMsg:output_type -> pbs.CommonResponse
	2, // 3: pbs.CmdService.P2pShowPeers:output_type -> pbs.CommonResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmdService_proto_init() }
func file_cmdService_proto_init() {
	if File_cmdService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmdService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicMsg); i {
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
		file_cmdService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowPeer); i {
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
		file_cmdService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResponse); i {
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
			RawDescriptor: file_cmdService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cmdService_proto_goTypes,
		DependencyIndexes: file_cmdService_proto_depIdxs,
		MessageInfos:      file_cmdService_proto_msgTypes,
	}.Build()
	File_cmdService_proto = out.File
	file_cmdService_proto_rawDesc = nil
	file_cmdService_proto_goTypes = nil
	file_cmdService_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CmdServiceClient is the client API for CmdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CmdServiceClient interface {
	P2PSendTopicMsg(ctx context.Context, in *TopicMsg, opts ...grpc.CallOption) (*CommonResponse, error)
	P2PShowPeers(ctx context.Context, in *ShowPeer, opts ...grpc.CallOption) (*CommonResponse, error)
}

type cmdServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCmdServiceClient(cc grpc.ClientConnInterface) CmdServiceClient {
	return &cmdServiceClient{cc}
}

func (c *cmdServiceClient) P2PSendTopicMsg(ctx context.Context, in *TopicMsg, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/pbs.CmdService/P2pSendTopicMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdServiceClient) P2PShowPeers(ctx context.Context, in *ShowPeer, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/pbs.CmdService/P2pShowPeers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CmdServiceServer is the server API for CmdService service.
type CmdServiceServer interface {
	P2PSendTopicMsg(context.Context, *TopicMsg) (*CommonResponse, error)
	P2PShowPeers(context.Context, *ShowPeer) (*CommonResponse, error)
}

// UnimplementedCmdServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCmdServiceServer struct {
}

func (*UnimplementedCmdServiceServer) P2PSendTopicMsg(context.Context, *TopicMsg) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method P2PSendTopicMsg not implemented")
}
func (*UnimplementedCmdServiceServer) P2PShowPeers(context.Context, *ShowPeer) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method P2PShowPeers not implemented")
}

func RegisterCmdServiceServer(s *grpc.Server, srv CmdServiceServer) {
	s.RegisterService(&_CmdService_serviceDesc, srv)
}

func _CmdService_P2PSendTopicMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServiceServer).P2PSendTopicMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbs.CmdService/P2PSendTopicMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServiceServer).P2PSendTopicMsg(ctx, req.(*TopicMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmdService_P2PShowPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowPeer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServiceServer).P2PShowPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbs.CmdService/P2PShowPeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServiceServer).P2PShowPeers(ctx, req.(*ShowPeer))
	}
	return interceptor(ctx, in, info, handler)
}

var _CmdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbs.CmdService",
	HandlerType: (*CmdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "P2pSendTopicMsg",
			Handler:    _CmdService_P2PSendTopicMsg_Handler,
		},
		{
			MethodName: "P2pShowPeers",
			Handler:    _CmdService_P2PShowPeers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmdService.proto",
}
