//推荐系统所有公共定义

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.5.0
// source: comm.proto

package common

import (
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

//feed id
type FEED_ID int32

const (
	FEED_ID_FEED_ID_NIL             FEED_ID = 0 //预留
	FEED_ID_FEED_ID_HOME_FEED       FEED_ID = 1 //首页
	FEED_ID_FEED_ID_PRODUCT_DETAIL  FEED_ID = 2 //商品详情页
	FEED_ID_FEED_ID_Personal_CENTER FEED_ID = 3 //个人中心
)

// Enum value maps for FEED_ID.
var (
	FEED_ID_name = map[int32]string{
		0: "FEED_ID_NIL",
		1: "FEED_ID_HOME_FEED",
		2: "FEED_ID_PRODUCT_DETAIL",
		3: "FEED_ID_Personal_CENTER",
	}
	FEED_ID_value = map[string]int32{
		"FEED_ID_NIL":             0,
		"FEED_ID_HOME_FEED":       1,
		"FEED_ID_PRODUCT_DETAIL":  2,
		"FEED_ID_Personal_CENTER": 3,
	}
)

func (x FEED_ID) Enum() *FEED_ID {
	p := new(FEED_ID)
	*p = x
	return p
}

func (x FEED_ID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FEED_ID) Descriptor() protoreflect.EnumDescriptor {
	return file_comm_proto_enumTypes[0].Descriptor()
}

func (FEED_ID) Type() protoreflect.EnumType {
	return &file_comm_proto_enumTypes[0]
}

func (x FEED_ID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FEED_ID.Descriptor instead.
func (FEED_ID) EnumDescriptor() ([]byte, []int) {
	return file_comm_proto_rawDescGZIP(), []int{0}
}

//action
type ACTION int32

const (
	ACTION_ACTION_NIL        ACTION = 0 //预留
	ACTION_COLD_START_ACTION ACTION = 1 //冷启动
	ACTION_HOME_ACTION       ACTION = 2 //首页
	ACTION_VARIETY_ACTION    ACTION = 3 //多样性
	ACTION_DEFAULT_ACTION    ACTION = 4 //默认
)

// Enum value maps for ACTION.
var (
	ACTION_name = map[int32]string{
		0: "ACTION_NIL",
		1: "COLD_START_ACTION",
		2: "HOME_ACTION",
		3: "VARIETY_ACTION",
		4: "DEFAULT_ACTION",
	}
	ACTION_value = map[string]int32{
		"ACTION_NIL":        0,
		"COLD_START_ACTION": 1,
		"HOME_ACTION":       2,
		"VARIETY_ACTION":    3,
		"DEFAULT_ACTION":    4,
	}
)

func (x ACTION) Enum() *ACTION {
	p := new(ACTION)
	*p = x
	return p
}

func (x ACTION) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ACTION) Descriptor() protoreflect.EnumDescriptor {
	return file_comm_proto_enumTypes[1].Descriptor()
}

func (ACTION) Type() protoreflect.EnumType {
	return &file_comm_proto_enumTypes[1]
}

func (x ACTION) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ACTION.Descriptor instead.
func (ACTION) EnumDescriptor() ([]byte, []int) {
	return file_comm_proto_rawDescGZIP(), []int{1}
}

//policy
type POLICY int32

const (
	POLICY_POLICY_DEFAULT        POLICY = 0  //预留
	POLICY_USER_AGE_RELATED      POLICY = 1  //用户年龄关联
	POLICY_USER_GENDER_RELATED   POLICY = 2  //用户性别关联
	POLICY_PRODUCT_CATE_RELATED  POLICY = 3  //商品类目关联
	POLICY_USER_RELATED          POLICY = 4  //用户关联
	POLICY_PRODUCT_RELATED       POLICY = 5  //商品关联
	POLICY_LIKE_RELATED          POLICY = 6  //用户喜欢关联
	POLICY_CLICK_RELATED         POLICY = 7  //过往点击相似
	POLICY_BUY_RELATED           POLICY = 8  //过往购买相似
	POLICY_SHOPPING_CART_RELATED POLICY = 9  //购物车相似
	POLICY_HOT_SEARCH            POLICY = 10 //热搜
	POLICY_HOT_LIKE              POLICY = 11 //热门喜欢
	POLICY_HOT_BUY               POLICY = 12 //热门购买
	POLICY_USER_SHARE            POLICY = 13 //用户分享关联
	POLICY_FRIEND_LIKE           POLICY = 14 //朋友喜欢
	POLICY_FRIEND_CLICK          POLICY = 15 //朋友点击
	POLICY_FRIEND_BUY            POLICY = 16 //朋友购买
)

// Enum value maps for POLICY.
var (
	POLICY_name = map[int32]string{
		0:  "POLICY_DEFAULT",
		1:  "USER_AGE_RELATED",
		2:  "USER_GENDER_RELATED",
		3:  "PRODUCT_CATE_RELATED",
		4:  "USER_RELATED",
		5:  "PRODUCT_RELATED",
		6:  "LIKE_RELATED",
		7:  "CLICK_RELATED",
		8:  "BUY_RELATED",
		9:  "SHOPPING_CART_RELATED",
		10: "HOT_SEARCH",
		11: "HOT_LIKE",
		12: "HOT_BUY",
		13: "USER_SHARE",
		14: "FRIEND_LIKE",
		15: "FRIEND_CLICK",
		16: "FRIEND_BUY",
	}
	POLICY_value = map[string]int32{
		"POLICY_DEFAULT":        0,
		"USER_AGE_RELATED":      1,
		"USER_GENDER_RELATED":   2,
		"PRODUCT_CATE_RELATED":  3,
		"USER_RELATED":          4,
		"PRODUCT_RELATED":       5,
		"LIKE_RELATED":          6,
		"CLICK_RELATED":         7,
		"BUY_RELATED":           8,
		"SHOPPING_CART_RELATED": 9,
		"HOT_SEARCH":            10,
		"HOT_LIKE":              11,
		"HOT_BUY":               12,
		"USER_SHARE":            13,
		"FRIEND_LIKE":           14,
		"FRIEND_CLICK":          15,
		"FRIEND_BUY":            16,
	}
)

func (x POLICY) Enum() *POLICY {
	p := new(POLICY)
	*p = x
	return p
}

func (x POLICY) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (POLICY) Descriptor() protoreflect.EnumDescriptor {
	return file_comm_proto_enumTypes[2].Descriptor()
}

func (POLICY) Type() protoreflect.EnumType {
	return &file_comm_proto_enumTypes[2]
}

func (x POLICY) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use POLICY.Descriptor instead.
func (POLICY) EnumDescriptor() ([]byte, []int) {
	return file_comm_proto_rawDescGZIP(), []int{2}
}

//index type
type INDEX_TYPE int32

const (
	INDEX_TYPE_INDEX_TYPE_NIL          INDEX_TYPE = 0  //预留
	INDEX_TYPE_USER_AGE                INDEX_TYPE = 1  //用户年龄
	INDEX_TYPE_USER_GENDER             INDEX_TYPE = 2  //用户性别
	INDEX_TYPE_USER_CLICK              INDEX_TYPE = 3  //用户点击商品
	INDEX_TYPE_USER_LIKE_SHARE         INDEX_TYPE = 4  //用户分享
	INDEX_TYPE_PRODUCT_FIRST_CATE      INDEX_TYPE = 5  //一级类目
	INDEX_TYPE_PRODUCT_SECOND_CATE     INDEX_TYPE = 6  //二级类目
	INDEX_TYPE_HOT_SEARCH_WORD         INDEX_TYPE = 7  //热搜分割词
	INDEX_TYPE_USER_BUY                INDEX_TYPE = 8  //用户曾经购买
	INDEX_TYPE_USER_SHOPPING_CART      INDEX_TYPE = 9  //用户加入购物车
	INDEX_TYPE_HOT_PRODUCT_LIKE        INDEX_TYPE = 10 //商品被喜欢
	INDEX_TYPE_HOT_PRODUCT_BUY         INDEX_TYPE = 11 //商品被购买
	INDEX_TYPE_HOT_PRODUCT_SHARE       INDEX_TYPE = 12 //商品被分享
	INDEX_TYPE_USER_LIKE               INDEX_TYPE = 13 //朋友列表
	INDEX_TYPE_OPERATION_RESOURCE_     INDEX_TYPE = 14 //运营资源位
	INDEX_TYPE_BACK_UP                 INDEX_TYPE = 15 //兜底池
	INDEX_TYPE_TOP_FIRST_CATE_PRODUCT  INDEX_TYPE = 16 //各个一级品类销量榜单
	INDEX_TYPE_TOP_SECOND_CATE_PRODUCT INDEX_TYPE = 17 //各个二级品类销量榜单
)

// Enum value maps for INDEX_TYPE.
var (
	INDEX_TYPE_name = map[int32]string{
		0:  "INDEX_TYPE_NIL",
		1:  "USER_AGE",
		2:  "USER_GENDER",
		3:  "USER_CLICK",
		4:  "USER_LIKE_SHARE",
		5:  "PRODUCT_FIRST_CATE",
		6:  "PRODUCT_SECOND_CATE",
		7:  "HOT_SEARCH_WORD",
		8:  "USER_BUY",
		9:  "USER_SHOPPING_CART",
		10: "HOT_PRODUCT_LIKE",
		11: "HOT_PRODUCT_BUY",
		12: "HOT_PRODUCT_SHARE",
		13: "USER_LIKE",
		14: "OPERATION_RESOURCE_",
		15: "BACK_UP",
		16: "TOP_FIRST_CATE_PRODUCT",
		17: "TOP_SECOND_CATE_PRODUCT",
	}
	INDEX_TYPE_value = map[string]int32{
		"INDEX_TYPE_NIL":          0,
		"USER_AGE":                1,
		"USER_GENDER":             2,
		"USER_CLICK":              3,
		"USER_LIKE_SHARE":         4,
		"PRODUCT_FIRST_CATE":      5,
		"PRODUCT_SECOND_CATE":     6,
		"HOT_SEARCH_WORD":         7,
		"USER_BUY":                8,
		"USER_SHOPPING_CART":      9,
		"HOT_PRODUCT_LIKE":        10,
		"HOT_PRODUCT_BUY":         11,
		"HOT_PRODUCT_SHARE":       12,
		"USER_LIKE":               13,
		"OPERATION_RESOURCE_":     14,
		"BACK_UP":                 15,
		"TOP_FIRST_CATE_PRODUCT":  16,
		"TOP_SECOND_CATE_PRODUCT": 17,
	}
)

func (x INDEX_TYPE) Enum() *INDEX_TYPE {
	p := new(INDEX_TYPE)
	*p = x
	return p
}

func (x INDEX_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (INDEX_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_comm_proto_enumTypes[3].Descriptor()
}

func (INDEX_TYPE) Type() protoreflect.EnumType {
	return &file_comm_proto_enumTypes[3]
}

func (x INDEX_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use INDEX_TYPE.Descriptor instead.
func (INDEX_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_comm_proto_rawDescGZIP(), []int{3}
}

var File_comm_proto protoreflect.FileDescriptor

var file_comm_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x6a, 0x0a, 0x07,
	0x46, 0x45, 0x45, 0x44, 0x5f, 0x49, 0x44, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x45, 0x45, 0x44, 0x5f,
	0x49, 0x44, 0x5f, 0x4e, 0x49, 0x4c, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x45, 0x45, 0x44,
	0x5f, 0x49, 0x44, 0x5f, 0x48, 0x4f, 0x4d, 0x45, 0x5f, 0x46, 0x45, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x1a, 0x0a, 0x16, 0x46, 0x45, 0x45, 0x44, 0x5f, 0x49, 0x44, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55,
	0x43, 0x54, 0x5f, 0x44, 0x45, 0x54, 0x41, 0x49, 0x4c, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x46,
	0x45, 0x45, 0x44, 0x5f, 0x49, 0x44, 0x5f, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x43, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x10, 0x03, 0x2a, 0x68, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x49, 0x4c,
	0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4f, 0x4c, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54,
	0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x48, 0x4f, 0x4d,
	0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x56, 0x41,
	0x52, 0x49, 0x45, 0x54, 0x59, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12, 0x12,
	0x0a, 0x0e, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x04, 0x2a, 0xcb, 0x02, 0x0a, 0x06, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x12, 0x12, 0x0a,
	0x0e, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10,
	0x00, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x45,
	0x4c, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x43, 0x41, 0x54, 0x45,
	0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f,
	0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x45, 0x44, 0x10,
	0x05, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x49, 0x4b, 0x45, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x45,
	0x44, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x5f, 0x52, 0x45, 0x4c,
	0x41, 0x54, 0x45, 0x44, 0x10, 0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x55, 0x59, 0x5f, 0x52, 0x45,
	0x4c, 0x41, 0x54, 0x45, 0x44, 0x10, 0x08, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x48, 0x4f, 0x50, 0x50,
	0x49, 0x4e, 0x47, 0x5f, 0x43, 0x41, 0x52, 0x54, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x45, 0x44,
	0x10, 0x09, 0x12, 0x0e, 0x0a, 0x0a, 0x48, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48,
	0x10, 0x0a, 0x12, 0x0c, 0x0a, 0x08, 0x48, 0x4f, 0x54, 0x5f, 0x4c, 0x49, 0x4b, 0x45, 0x10, 0x0b,
	0x12, 0x0b, 0x0a, 0x07, 0x48, 0x4f, 0x54, 0x5f, 0x42, 0x55, 0x59, 0x10, 0x0c, 0x12, 0x0e, 0x0a,
	0x0a, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x48, 0x41, 0x52, 0x45, 0x10, 0x0d, 0x12, 0x0f, 0x0a,
	0x0b, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x5f, 0x4c, 0x49, 0x4b, 0x45, 0x10, 0x0e, 0x12, 0x10,
	0x0a, 0x0c, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x5f, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x10, 0x0f,
	0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x5f, 0x42, 0x55, 0x59, 0x10, 0x10,
	0x2a, 0x80, 0x03, 0x0a, 0x0a, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x12,
	0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x49,
	0x4c, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x47, 0x45, 0x10,
	0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52,
	0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x43, 0x4c, 0x49, 0x43, 0x4b,
	0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4c, 0x49, 0x4b, 0x45, 0x5f,
	0x53, 0x48, 0x41, 0x52, 0x45, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x52, 0x4f, 0x44, 0x55,
	0x43, 0x54, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x10, 0x05, 0x12,
	0x17, 0x0a, 0x13, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x53, 0x45, 0x43, 0x4f, 0x4e,
	0x44, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x48, 0x4f, 0x54, 0x5f,
	0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x07, 0x12, 0x0c, 0x0a,
	0x08, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x42, 0x55, 0x59, 0x10, 0x08, 0x12, 0x16, 0x0a, 0x12, 0x55,
	0x53, 0x45, 0x52, 0x5f, 0x53, 0x48, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x41, 0x52,
	0x54, 0x10, 0x09, 0x12, 0x14, 0x0a, 0x10, 0x48, 0x4f, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55,
	0x43, 0x54, 0x5f, 0x4c, 0x49, 0x4b, 0x45, 0x10, 0x0a, 0x12, 0x13, 0x0a, 0x0f, 0x48, 0x4f, 0x54,
	0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x42, 0x55, 0x59, 0x10, 0x0b, 0x12, 0x15,
	0x0a, 0x11, 0x48, 0x4f, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x53, 0x48,
	0x41, 0x52, 0x45, 0x10, 0x0c, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4c, 0x49,
	0x4b, 0x45, 0x10, 0x0d, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x10, 0x0e, 0x12, 0x0b, 0x0a,
	0x07, 0x42, 0x41, 0x43, 0x4b, 0x5f, 0x55, 0x50, 0x10, 0x0f, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x4f,
	0x50, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x52, 0x4f,
	0x44, 0x55, 0x43, 0x54, 0x10, 0x10, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x4f, 0x50, 0x5f, 0x53, 0x45,
	0x43, 0x4f, 0x4e, 0x44, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43,
	0x54, 0x10, 0x11, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comm_proto_rawDescOnce sync.Once
	file_comm_proto_rawDescData = file_comm_proto_rawDesc
)

func file_comm_proto_rawDescGZIP() []byte {
	file_comm_proto_rawDescOnce.Do(func() {
		file_comm_proto_rawDescData = protoimpl.X.CompressGZIP(file_comm_proto_rawDescData)
	})
	return file_comm_proto_rawDescData
}

var file_comm_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_comm_proto_goTypes = []interface{}{
	(FEED_ID)(0),    // 0: FEED_ID
	(ACTION)(0),     // 1: ACTION
	(POLICY)(0),     // 2: POLICY
	(INDEX_TYPE)(0), // 3: INDEX_TYPE
}
var file_comm_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_comm_proto_init() }
func file_comm_proto_init() {
	if File_comm_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_comm_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_comm_proto_goTypes,
		DependencyIndexes: file_comm_proto_depIdxs,
		EnumInfos:         file_comm_proto_enumTypes,
	}.Build()
	File_comm_proto = out.File
	file_comm_proto_rawDesc = nil
	file_comm_proto_goTypes = nil
	file_comm_proto_depIdxs = nil
}
