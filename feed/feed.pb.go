// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.5.0
// source: proto/feed.proto

package proto

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

//推荐理由这块先不开发
type RecommendReasonType int32

const (
	RecommendReasonType_Friend_Like     RecommendReasonType = 0 //朋友喜欢/点赞行为
	RecommendReasonType_Friend_Buy      RecommendReasonType = 1 //朋友购买
	RecommendReasonType_Friend_Share    RecommendReasonType = 2 //朋友分享
	RecommendReasonType_Friend_Click    RecommendReasonType = 3 //朋友看过/点击行为
	RecommendReasonType_Myself_Like_Sim RecommendReasonType = 4 //自己喜欢/点赞相似物品
)

// Enum value maps for RecommendReasonType.
var (
	RecommendReasonType_name = map[int32]string{
		0: "Friend_Like",
		1: "Friend_Buy",
		2: "Friend_Share",
		3: "Friend_Click",
		4: "Myself_Like_Sim",
	}
	RecommendReasonType_value = map[string]int32{
		"Friend_Like":     0,
		"Friend_Buy":      1,
		"Friend_Share":    2,
		"Friend_Click":    3,
		"Myself_Like_Sim": 4,
	}
)

func (x RecommendReasonType) Enum() *RecommendReasonType {
	p := new(RecommendReasonType)
	*p = x
	return p
}

func (x RecommendReasonType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RecommendReasonType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_feed_proto_enumTypes[0].Descriptor()
}

func (RecommendReasonType) Type() protoreflect.EnumType {
	return &file_proto_feed_proto_enumTypes[0]
}

func (x RecommendReasonType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecommendReasonType.Descriptor instead.
func (RecommendReasonType) EnumDescriptor() ([]byte, []int) {
	return file_proto_feed_proto_rawDescGZIP(), []int{0}
}

type FeedIDType int32

const (
	FeedIDType_Home_Page_Feed       FeedIDType = 0 //首页推荐页
	FeedIDType_Personal_Center_Feed FeedIDType = 1 //个人中心页
)

// Enum value maps for FeedIDType.
var (
	FeedIDType_name = map[int32]string{
		0: "Home_Page_Feed",
		1: "Personal_Center_Feed",
	}
	FeedIDType_value = map[string]int32{
		"Home_Page_Feed":       0,
		"Personal_Center_Feed": 1,
	}
)

func (x FeedIDType) Enum() *FeedIDType {
	p := new(FeedIDType)
	*p = x
	return p
}

func (x FeedIDType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeedIDType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_feed_proto_enumTypes[1].Descriptor()
}

func (FeedIDType) Type() protoreflect.EnumType {
	return &file_proto_feed_proto_enumTypes[1]
}

func (x FeedIDType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeedIDType.Descriptor instead.
func (FeedIDType) EnumDescriptor() ([]byte, []int) {
	return file_proto_feed_proto_rawDescGZIP(), []int{1}
}

type QueryFeedItemsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID        uint64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`              //用户ID
	FeedID        string `protobuf:"bytes,2,opt,name=FeedID,proto3" json:"FeedID,omitempty"`               //org.js.client.grpc.Feed id【场景】
	PageNum       int32  `protobuf:"varint,3,opt,name=PageNum,proto3" json:"PageNum,omitempty"`            //页码
	Limit         int32  `protobuf:"varint,4,opt,name=Limit,proto3" json:"Limit,omitempty"`                //分页大小
	RequestOption string `protobuf:"bytes,5,opt,name=RequestOption,proto3" json:"RequestOption,omitempty"` //扩展参数【json格式】
}

func (x *QueryFeedItemsReq) Reset() {
	*x = QueryFeedItemsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_feed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryFeedItemsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryFeedItemsReq) ProtoMessage() {}

func (x *QueryFeedItemsReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_feed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryFeedItemsReq.ProtoReflect.Descriptor instead.
func (*QueryFeedItemsReq) Descriptor() ([]byte, []int) {
	return file_proto_feed_proto_rawDescGZIP(), []int{0}
}

func (x *QueryFeedItemsReq) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *QueryFeedItemsReq) GetFeedID() string {
	if x != nil {
		return x.FeedID
	}
	return ""
}

func (x *QueryFeedItemsReq) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *QueryFeedItemsReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *QueryFeedItemsReq) GetRequestOption() string {
	if x != nil {
		return x.RequestOption
	}
	return ""
}

type QueryFeedItemsRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items          []*FeedItem `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
	ResponseOption string      `protobuf:"bytes,2,opt,name=ResponseOption,proto3" json:"ResponseOption,omitempty"` //扩展参数
}

func (x *QueryFeedItemsRsp) Reset() {
	*x = QueryFeedItemsRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_feed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryFeedItemsRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryFeedItemsRsp) ProtoMessage() {}

func (x *QueryFeedItemsRsp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_feed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryFeedItemsRsp.ProtoReflect.Descriptor instead.
func (*QueryFeedItemsRsp) Descriptor() ([]byte, []int) {
	return file_proto_feed_proto_rawDescGZIP(), []int{1}
}

func (x *QueryFeedItemsRsp) GetItems() []*FeedItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *QueryFeedItemsRsp) GetResponseOption() string {
	if x != nil {
		return x.ResponseOption
	}
	return ""
}

//Item
type FeedItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemID uint64              `protobuf:"varint,1,opt,name=ItemID,proto3" json:"ItemID,omitempty"`                      //item id
	Score  float64             `protobuf:"fixed64,2,opt,name=Score,proto3" json:"Score,omitempty"`                       //打分值
	Type   RecommendReasonType `protobuf:"varint,3,opt,name=type,proto3,enum=RecommendReasonType" json:"type,omitempty"` //推荐理由
}

func (x *FeedItem) Reset() {
	*x = FeedItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_feed_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedItem) ProtoMessage() {}

func (x *FeedItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_feed_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedItem.ProtoReflect.Descriptor instead.
func (*FeedItem) Descriptor() ([]byte, []int) {
	return file_proto_feed_proto_rawDescGZIP(), []int{2}
}

func (x *FeedItem) GetItemID() uint64 {
	if x != nil {
		return x.ItemID
	}
	return 0
}

func (x *FeedItem) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *FeedItem) GetType() RecommendReasonType {
	if x != nil {
		return x.Type
	}
	return RecommendReasonType_Friend_Like
}

var File_proto_feed_proto protoreflect.FileDescriptor

var file_proto_feed_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x65, 0x65, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x46, 0x65, 0x65, 0x64, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x46, 0x65, 0x65, 0x64, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61, 0x67, 0x65,
	0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x50, 0x61, 0x67, 0x65, 0x4e,
	0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5c,
	0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x52, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x62, 0x0a, 0x08,
	0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x74, 0x65, 0x6d,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x2a, 0x6f, 0x0a, 0x13, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x5f, 0x4c, 0x69, 0x6b, 0x65, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x5f, 0x42, 0x75, 0x79, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x5f, 0x53, 0x68, 0x61, 0x72, 0x65, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x5f, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f,
	0x4d, 0x79, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x4c, 0x69, 0x6b, 0x65, 0x5f, 0x53, 0x69, 0x6d, 0x10,
	0x04, 0x2a, 0x3a, 0x0a, 0x0a, 0x46, 0x65, 0x65, 0x64, 0x49, 0x44, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x0e, 0x48, 0x6f, 0x6d, 0x65, 0x5f, 0x50, 0x61, 0x67, 0x65, 0x5f, 0x46, 0x65, 0x65,
	0x64, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x46, 0x65, 0x65, 0x64, 0x10, 0x01, 0x32, 0x47, 0x0a,
	0x0b, 0x46, 0x65, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x12,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52,
	0x65, 0x71, 0x1a, 0x12, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x52, 0x73, 0x70, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_feed_proto_rawDescOnce sync.Once
	file_proto_feed_proto_rawDescData = file_proto_feed_proto_rawDesc
)

func file_proto_feed_proto_rawDescGZIP() []byte {
	file_proto_feed_proto_rawDescOnce.Do(func() {
		file_proto_feed_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_feed_proto_rawDescData)
	})
	return file_proto_feed_proto_rawDescData
}

var file_proto_feed_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_feed_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_feed_proto_goTypes = []interface{}{
	(RecommendReasonType)(0),  // 0: RecommendReasonType
	(FeedIDType)(0),           // 1: FeedIDType
	(*QueryFeedItemsReq)(nil), // 2: QueryFeedItemsReq
	(*QueryFeedItemsRsp)(nil), // 3: QueryFeedItemsRsp
	(*FeedItem)(nil),          // 4: FeedItem
}
var file_proto_feed_proto_depIdxs = []int32{
	4, // 0: QueryFeedItemsRsp.Items:type_name -> FeedItem
	0, // 1: FeedItem.type:type_name -> RecommendReasonType
	2, // 2: FeedService.QueryFeedItems:input_type -> QueryFeedItemsReq
	3, // 3: FeedService.QueryFeedItems:output_type -> QueryFeedItemsRsp
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_feed_proto_init() }
func file_proto_feed_proto_init() {
	if File_proto_feed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_feed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryFeedItemsReq); i {
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
		file_proto_feed_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryFeedItemsRsp); i {
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
		file_proto_feed_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedItem); i {
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
			RawDescriptor: file_proto_feed_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_feed_proto_goTypes,
		DependencyIndexes: file_proto_feed_proto_depIdxs,
		EnumInfos:         file_proto_feed_proto_enumTypes,
		MessageInfos:      file_proto_feed_proto_msgTypes,
	}.Build()
	File_proto_feed_proto = out.File
	file_proto_feed_proto_rawDesc = nil
	file_proto_feed_proto_goTypes = nil
	file_proto_feed_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FeedServiceClient is the client API for FeedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedServiceClient interface {
	//通用的查询Item接口
	QueryFeedItems(ctx context.Context, in *QueryFeedItemsReq, opts ...grpc.CallOption) (*QueryFeedItemsRsp, error)
}

type feedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedServiceClient(cc grpc.ClientConnInterface) FeedServiceClient {
	return &feedServiceClient{cc}
}

func (c *feedServiceClient) QueryFeedItems(ctx context.Context, in *QueryFeedItemsReq, opts ...grpc.CallOption) (*QueryFeedItemsRsp, error) {
	out := new(QueryFeedItemsRsp)
	err := c.cc.Invoke(ctx, "/FeedService/QueryFeedItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedServiceServer is the server API for FeedService service.
type FeedServiceServer interface {
	//通用的查询Item接口
	QueryFeedItems(context.Context, *QueryFeedItemsReq) (*QueryFeedItemsRsp, error)
}

// UnimplementedFeedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFeedServiceServer struct {
}

func (*UnimplementedFeedServiceServer) QueryFeedItems(context.Context, *QueryFeedItemsReq) (*QueryFeedItemsRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryFeedItems not implemented")
}

func RegisterFeedServiceServer(s *grpc.Server, srv FeedServiceServer) {
	s.RegisterService(&_FeedService_serviceDesc, srv)
}

func _FeedService_QueryFeedItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFeedItemsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).QueryFeedItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FeedService/QueryFeedItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).QueryFeedItems(ctx, req.(*QueryFeedItemsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "FeedService",
	HandlerType: (*FeedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryFeedItems",
			Handler:    _FeedService_QueryFeedItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/feed.proto",
}