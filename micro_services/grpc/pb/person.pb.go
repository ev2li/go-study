// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: person.proto

package pb

import (
	context "context"
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

//消息体 一个package中不允许定义同名的消息体
type Teacher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Age  int32  `protobuf:"varint,1,opt,name=age,proto3" json:"age,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Teacher) Reset() {
	*x = Teacher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Teacher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Teacher) ProtoMessage() {}

func (x *Teacher) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Teacher.ProtoReflect.Descriptor instead.
func (*Teacher) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{0}
}

func (x *Teacher) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Teacher) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_person_proto protoreflect.FileDescriptor

var file_person_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x2f, 0x0a, 0x07, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x32, 0x2f, 0x0a, 0x07, 0x53, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24,
	0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e,
	0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_person_proto_rawDescOnce sync.Once
	file_person_proto_rawDescData = file_person_proto_rawDesc
)

func file_person_proto_rawDescGZIP() []byte {
	file_person_proto_rawDescOnce.Do(func() {
		file_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_person_proto_rawDescData)
	})
	return file_person_proto_rawDescData
}

var file_person_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_person_proto_goTypes = []interface{}{
	(*Teacher)(nil), // 0: pb.Teacher
}
var file_person_proto_depIdxs = []int32{
	0, // 0: pb.SayName.SayHello:input_type -> pb.Teacher
	0, // 1: pb.SayName.SayHello:output_type -> pb.Teacher
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_person_proto_init() }
func file_person_proto_init() {
	if File_person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Teacher); i {
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
			RawDescriptor: file_person_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_person_proto_goTypes,
		DependencyIndexes: file_person_proto_depIdxs,
		MessageInfos:      file_person_proto_msgTypes,
	}.Build()
	File_person_proto = out.File
	file_person_proto_rawDesc = nil
	file_person_proto_goTypes = nil
	file_person_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SayNameClient is the client API for SayName service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SayNameClient interface {
	SayHello(ctx context.Context, in *Teacher, opts ...grpc.CallOption) (*Teacher, error)
}

type sayNameClient struct {
	cc grpc.ClientConnInterface
}

func NewSayNameClient(cc grpc.ClientConnInterface) SayNameClient {
	return &sayNameClient{cc}
}

func (c *sayNameClient) SayHello(ctx context.Context, in *Teacher, opts ...grpc.CallOption) (*Teacher, error) {
	out := new(Teacher)
	err := c.cc.Invoke(ctx, "/pb.SayName/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SayNameServer is the server API for SayName service.
type SayNameServer interface {
	SayHello(context.Context, *Teacher) (*Teacher, error)
}

// UnimplementedSayNameServer can be embedded to have forward compatible implementations.
type UnimplementedSayNameServer struct {
}

func (*UnimplementedSayNameServer) SayHello(context.Context, *Teacher) (*Teacher, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterSayNameServer(s *grpc.Server, srv SayNameServer) {
	s.RegisterService(&_SayName_serviceDesc, srv)
}

func _SayName_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Teacher)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SayNameServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SayName/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SayNameServer).SayHello(ctx, req.(*Teacher))
	}
	return interceptor(ctx, in, info, handler)
}

var _SayName_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SayName",
	HandlerType: (*SayNameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _SayName_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "person.proto",
}
