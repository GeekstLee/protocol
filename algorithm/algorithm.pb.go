//算法模型服务协议

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.5.0
// source: algorithm.proto

package algorithm

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

type AlgorithmItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    string  `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`         //item id
	Score float64 `protobuf:"fixed64,2,opt,name=Score,proto3" json:"Score,omitempty"` //打分值
}

func (x *AlgorithmItem) Reset() {
	*x = AlgorithmItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_algorithm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlgorithmItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlgorithmItem) ProtoMessage() {}

func (x *AlgorithmItem) ProtoReflect() protoreflect.Message {
	mi := &file_algorithm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlgorithmItem.ProtoReflect.Descriptor instead.
func (*AlgorithmItem) Descriptor() ([]byte, []int) {
	return file_algorithm_proto_rawDescGZIP(), []int{0}
}

func (x *AlgorithmItem) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *AlgorithmItem) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

type UserCFReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"` //用户ID
}

func (x *UserCFReq) Reset() {
	*x = UserCFReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_algorithm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCFReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCFReq) ProtoMessage() {}

func (x *UserCFReq) ProtoReflect() protoreflect.Message {
	mi := &file_algorithm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCFReq.ProtoReflect.Descriptor instead.
func (*UserCFReq) Descriptor() ([]byte, []int) {
	return file_algorithm_proto_rawDescGZIP(), []int{1}
}

func (x *UserCFReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type UserCFRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RelatedUsers []string `protobuf:"bytes,1,rep,name=relatedUsers,proto3" json:"relatedUsers,omitempty"`
}

func (x *UserCFRsp) Reset() {
	*x = UserCFRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_algorithm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCFRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCFRsp) ProtoMessage() {}

func (x *UserCFRsp) ProtoReflect() protoreflect.Message {
	mi := &file_algorithm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCFRsp.ProtoReflect.Descriptor instead.
func (*UserCFRsp) Descriptor() ([]byte, []int) {
	return file_algorithm_proto_rawDescGZIP(), []int{2}
}

func (x *UserCFRsp) GetRelatedUsers() []string {
	if x != nil {
		return x.RelatedUsers
	}
	return nil
}

type ItemCFReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*AlgorithmItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ItemCFReq) Reset() {
	*x = ItemCFReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_algorithm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemCFReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemCFReq) ProtoMessage() {}

func (x *ItemCFReq) ProtoReflect() protoreflect.Message {
	mi := &file_algorithm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemCFReq.ProtoReflect.Descriptor instead.
func (*ItemCFReq) Descriptor() ([]byte, []int) {
	return file_algorithm_proto_rawDescGZIP(), []int{3}
}

func (x *ItemCFReq) GetItems() []*AlgorithmItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type ItemCFRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RelatedItems []*AlgorithmItem `protobuf:"bytes,1,rep,name=relatedItems,proto3" json:"relatedItems,omitempty"` //相关item
}

func (x *ItemCFRsp) Reset() {
	*x = ItemCFRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_algorithm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemCFRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemCFRsp) ProtoMessage() {}

func (x *ItemCFRsp) ProtoReflect() protoreflect.Message {
	mi := &file_algorithm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemCFRsp.ProtoReflect.Descriptor instead.
func (*ItemCFRsp) Descriptor() ([]byte, []int) {
	return file_algorithm_proto_rawDescGZIP(), []int{4}
}

func (x *ItemCFRsp) GetRelatedItems() []*AlgorithmItem {
	if x != nil {
		return x.RelatedItems
	}
	return nil
}

var File_algorithm_proto protoreflect.FileDescriptor

var file_algorithm_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x35, 0x0a, 0x0d, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x23, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x46, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2f, 0x0a,
	0x09, 0x55, 0x73, 0x65, 0x72, 0x43, 0x46, 0x52, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x31,
	0x0a, 0x09, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x46, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x41, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0x3f, 0x0a, 0x09, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x46, 0x52, 0x73, 0x70, 0x12, 0x32,
	0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x32, 0x5a, 0x0a, 0x10, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x43, 0x46,
	0x12, 0x0a, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x46, 0x52, 0x65, 0x71, 0x1a, 0x0a, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x46, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x06, 0x49, 0x74,
	0x65, 0x6d, 0x43, 0x46, 0x12, 0x0a, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x46, 0x52, 0x65, 0x71,
	0x1a, 0x0a, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x46, 0x52, 0x73, 0x70, 0x22, 0x00, 0x42, 0x0d,
	0x5a, 0x0b, 0x2e, 0x2f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_algorithm_proto_rawDescOnce sync.Once
	file_algorithm_proto_rawDescData = file_algorithm_proto_rawDesc
)

func file_algorithm_proto_rawDescGZIP() []byte {
	file_algorithm_proto_rawDescOnce.Do(func() {
		file_algorithm_proto_rawDescData = protoimpl.X.CompressGZIP(file_algorithm_proto_rawDescData)
	})
	return file_algorithm_proto_rawDescData
}

var file_algorithm_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_algorithm_proto_goTypes = []interface{}{
	(*AlgorithmItem)(nil), // 0: AlgorithmItem
	(*UserCFReq)(nil),     // 1: UserCFReq
	(*UserCFRsp)(nil),     // 2: UserCFRsp
	(*ItemCFReq)(nil),     // 3: ItemCFReq
	(*ItemCFRsp)(nil),     // 4: ItemCFRsp
}
var file_algorithm_proto_depIdxs = []int32{
	0, // 0: ItemCFReq.items:type_name -> AlgorithmItem
	0, // 1: ItemCFRsp.relatedItems:type_name -> AlgorithmItem
	1, // 2: AlgorithmService.UserCF:input_type -> UserCFReq
	3, // 3: AlgorithmService.ItemCF:input_type -> ItemCFReq
	2, // 4: AlgorithmService.UserCF:output_type -> UserCFRsp
	4, // 5: AlgorithmService.ItemCF:output_type -> ItemCFRsp
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_algorithm_proto_init() }
func file_algorithm_proto_init() {
	if File_algorithm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_algorithm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlgorithmItem); i {
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
		file_algorithm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCFReq); i {
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
		file_algorithm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCFRsp); i {
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
		file_algorithm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemCFReq); i {
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
		file_algorithm_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemCFRsp); i {
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
			RawDescriptor: file_algorithm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_algorithm_proto_goTypes,
		DependencyIndexes: file_algorithm_proto_depIdxs,
		MessageInfos:      file_algorithm_proto_msgTypes,
	}.Build()
	File_algorithm_proto = out.File
	file_algorithm_proto_rawDesc = nil
	file_algorithm_proto_goTypes = nil
	file_algorithm_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AlgorithmServiceClient is the client API for AlgorithmService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlgorithmServiceClient interface {
	//基于用户的协同过滤，返回相似用户列表
	UserCF(ctx context.Context, in *UserCFReq, opts ...grpc.CallOption) (*UserCFRsp, error)
	//基于Item的协同过滤，返回相似Item列表
	ItemCF(ctx context.Context, in *ItemCFReq, opts ...grpc.CallOption) (*ItemCFRsp, error)
}

type algorithmServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAlgorithmServiceClient(cc grpc.ClientConnInterface) AlgorithmServiceClient {
	return &algorithmServiceClient{cc}
}

func (c *algorithmServiceClient) UserCF(ctx context.Context, in *UserCFReq, opts ...grpc.CallOption) (*UserCFRsp, error) {
	out := new(UserCFRsp)
	err := c.cc.Invoke(ctx, "/AlgorithmService/UserCF", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *algorithmServiceClient) ItemCF(ctx context.Context, in *ItemCFReq, opts ...grpc.CallOption) (*ItemCFRsp, error) {
	out := new(ItemCFRsp)
	err := c.cc.Invoke(ctx, "/AlgorithmService/ItemCF", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlgorithmServiceServer is the server API for AlgorithmService service.
type AlgorithmServiceServer interface {
	//基于用户的协同过滤，返回相似用户列表
	UserCF(context.Context, *UserCFReq) (*UserCFRsp, error)
	//基于Item的协同过滤，返回相似Item列表
	ItemCF(context.Context, *ItemCFReq) (*ItemCFRsp, error)
}

// UnimplementedAlgorithmServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAlgorithmServiceServer struct {
}

func (*UnimplementedAlgorithmServiceServer) UserCF(context.Context, *UserCFReq) (*UserCFRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCF not implemented")
}
func (*UnimplementedAlgorithmServiceServer) ItemCF(context.Context, *ItemCFReq) (*ItemCFRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ItemCF not implemented")
}

func RegisterAlgorithmServiceServer(s *grpc.Server, srv AlgorithmServiceServer) {
	s.RegisterService(&_AlgorithmService_serviceDesc, srv)
}

func _AlgorithmService_UserCF_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCFReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgorithmServiceServer).UserCF(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AlgorithmService/UserCF",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgorithmServiceServer).UserCF(ctx, req.(*UserCFReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlgorithmService_ItemCF_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemCFReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgorithmServiceServer).ItemCF(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AlgorithmService/ItemCF",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgorithmServiceServer).ItemCF(ctx, req.(*ItemCFReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _AlgorithmService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AlgorithmService",
	HandlerType: (*AlgorithmServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserCF",
			Handler:    _AlgorithmService_UserCF_Handler,
		},
		{
			MethodName: "ItemCF",
			Handler:    _AlgorithmService_ItemCF_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "algorithm.proto",
}
