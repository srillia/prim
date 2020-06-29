// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: im_protobuf.proto

package protobuf

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

// 查询用户是否在线
type QueryUsersOnlineReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SysAccount  string `protobuf:"bytes,1,opt,name=sysAccount,proto3" json:"sysAccount,omitempty"`   // AppPlatform
	AppPlatform string `protobuf:"bytes,2,opt,name=appPlatform,proto3" json:"appPlatform,omitempty"` // AppPlatform
	UserId      string `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`           // 用户ID
}

func (x *QueryUsersOnlineReq) Reset() {
	*x = QueryUsersOnlineReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_protobuf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUsersOnlineReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUsersOnlineReq) ProtoMessage() {}

func (x *QueryUsersOnlineReq) ProtoReflect() protoreflect.Message {
	mi := &file_im_protobuf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUsersOnlineReq.ProtoReflect.Descriptor instead.
func (*QueryUsersOnlineReq) Descriptor() ([]byte, []int) {
	return file_im_protobuf_proto_rawDescGZIP(), []int{0}
}

func (x *QueryUsersOnlineReq) GetSysAccount() string {
	if x != nil {
		return x.SysAccount
	}
	return ""
}

func (x *QueryUsersOnlineReq) GetAppPlatform() string {
	if x != nil {
		return x.AppPlatform
	}
	return ""
}

func (x *QueryUsersOnlineReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type QueryUsersOnlineRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetCode uint32 `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode,omitempty"`
	ErrMsg  string `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	Online  bool   `protobuf:"varint,3,opt,name=online,proto3" json:"online,omitempty"`
}

func (x *QueryUsersOnlineRsp) Reset() {
	*x = QueryUsersOnlineRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_protobuf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUsersOnlineRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUsersOnlineRsp) ProtoMessage() {}

func (x *QueryUsersOnlineRsp) ProtoReflect() protoreflect.Message {
	mi := &file_im_protobuf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUsersOnlineRsp.ProtoReflect.Descriptor instead.
func (*QueryUsersOnlineRsp) Descriptor() ([]byte, []int) {
	return file_im_protobuf_proto_rawDescGZIP(), []int{1}
}

func (x *QueryUsersOnlineRsp) GetRetCode() uint32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

func (x *QueryUsersOnlineRsp) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *QueryUsersOnlineRsp) GetOnline() bool {
	if x != nil {
		return x.Online
	}
	return false
}

// 发送消息
type SendMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq         string   `protobuf:"bytes,1,opt,name=seq,proto3" json:"seq,omitempty"`                 // 序列号
	SysAccount  string   `protobuf:"bytes,2,opt,name=sysAccount,proto3" json:"sysAccount,omitempty"`   // 用户ID
	AppPlatform string   `protobuf:"bytes,3,opt,name=appPlatform,proto3" json:"appPlatform,omitempty"` // 用户ID
	UserIds     []string `protobuf:"bytes,4,rep,name=userIds,proto3" json:"userIds,omitempty"`
	Msg         []byte   `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg,omitempty"` // 用户ID
}

func (x *SendMsgReq) Reset() {
	*x = SendMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_protobuf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMsgReq) ProtoMessage() {}

func (x *SendMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_im_protobuf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMsgReq.ProtoReflect.Descriptor instead.
func (*SendMsgReq) Descriptor() ([]byte, []int) {
	return file_im_protobuf_proto_rawDescGZIP(), []int{2}
}

func (x *SendMsgReq) GetSeq() string {
	if x != nil {
		return x.Seq
	}
	return ""
}

func (x *SendMsgReq) GetSysAccount() string {
	if x != nil {
		return x.SysAccount
	}
	return ""
}

func (x *SendMsgReq) GetAppPlatform() string {
	if x != nil {
		return x.AppPlatform
	}
	return ""
}

func (x *SendMsgReq) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *SendMsgReq) GetMsg() []byte {
	if x != nil {
		return x.Msg
	}
	return nil
}

type SendMsgRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetCode   uint32 `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode,omitempty"`
	ErrMsg    string `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	SendMsgId string `protobuf:"bytes,3,opt,name=sendMsgId,proto3" json:"sendMsgId,omitempty"`
}

func (x *SendMsgRsp) Reset() {
	*x = SendMsgRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_protobuf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMsgRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMsgRsp) ProtoMessage() {}

func (x *SendMsgRsp) ProtoReflect() protoreflect.Message {
	mi := &file_im_protobuf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMsgRsp.ProtoReflect.Descriptor instead.
func (*SendMsgRsp) Descriptor() ([]byte, []int) {
	return file_im_protobuf_proto_rawDescGZIP(), []int{3}
}

func (x *SendMsgRsp) GetRetCode() uint32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

func (x *SendMsgRsp) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *SendMsgRsp) GetSendMsgId() string {
	if x != nil {
		return x.SendMsgId
	}
	return ""
}

// 发送消息
type SendMsgAllReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq         string `protobuf:"bytes,1,opt,name=seq,proto3" json:"seq,omitempty"`                 // 序列号
	SysAccount  string `protobuf:"bytes,2,opt,name=sysAccount,proto3" json:"sysAccount,omitempty"`   // 用户ID
	AppPlatform string `protobuf:"bytes,3,opt,name=appPlatform,proto3" json:"appPlatform,omitempty"` // 用户ID
	UserId      string `protobuf:"bytes,4,opt,name=userId,proto3" json:"userId,omitempty"`           // 用户ID
	Cms         string `protobuf:"bytes,5,opt,name=cms,proto3" json:"cms,omitempty"`                 // cms 动作: msg/enter/exit
	Type        string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`               // type 消息类型，默认是 text
	Msg         string `protobuf:"bytes,7,opt,name=msg,proto3" json:"msg,omitempty"`                 // msg
}

func (x *SendMsgAllReq) Reset() {
	*x = SendMsgAllReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_protobuf_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMsgAllReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMsgAllReq) ProtoMessage() {}

func (x *SendMsgAllReq) ProtoReflect() protoreflect.Message {
	mi := &file_im_protobuf_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMsgAllReq.ProtoReflect.Descriptor instead.
func (*SendMsgAllReq) Descriptor() ([]byte, []int) {
	return file_im_protobuf_proto_rawDescGZIP(), []int{4}
}

func (x *SendMsgAllReq) GetSeq() string {
	if x != nil {
		return x.Seq
	}
	return ""
}

func (x *SendMsgAllReq) GetSysAccount() string {
	if x != nil {
		return x.SysAccount
	}
	return ""
}

func (x *SendMsgAllReq) GetAppPlatform() string {
	if x != nil {
		return x.AppPlatform
	}
	return ""
}

func (x *SendMsgAllReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SendMsgAllReq) GetCms() string {
	if x != nil {
		return x.Cms
	}
	return ""
}

func (x *SendMsgAllReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SendMsgAllReq) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SendMsgAllRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetCode   uint32 `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode,omitempty"`
	ErrMsg    string `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	SendMsgId string `protobuf:"bytes,3,opt,name=sendMsgId,proto3" json:"sendMsgId,omitempty"`
}

func (x *SendMsgAllRsp) Reset() {
	*x = SendMsgAllRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_protobuf_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMsgAllRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMsgAllRsp) ProtoMessage() {}

func (x *SendMsgAllRsp) ProtoReflect() protoreflect.Message {
	mi := &file_im_protobuf_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMsgAllRsp.ProtoReflect.Descriptor instead.
func (*SendMsgAllRsp) Descriptor() ([]byte, []int) {
	return file_im_protobuf_proto_rawDescGZIP(), []int{5}
}

func (x *SendMsgAllRsp) GetRetCode() uint32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

func (x *SendMsgAllRsp) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *SendMsgAllRsp) GetSendMsgId() string {
	if x != nil {
		return x.SendMsgId
	}
	return ""
}

// 获取用户列表
type GetUserListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUserListReq) Reset() {
	*x = GetUserListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_protobuf_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserListReq) ProtoMessage() {}

func (x *GetUserListReq) ProtoReflect() protoreflect.Message {
	mi := &file_im_protobuf_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserListReq.ProtoReflect.Descriptor instead.
func (*GetUserListReq) Descriptor() ([]byte, []int) {
	return file_im_protobuf_proto_rawDescGZIP(), []int{6}
}

type GetUserListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetCode uint32   `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode,omitempty"`
	ErrMsg  string   `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	UserId  []string `protobuf:"bytes,3,rep,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetUserListRsp) Reset() {
	*x = GetUserListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_protobuf_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserListRsp) ProtoMessage() {}

func (x *GetUserListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_im_protobuf_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserListRsp.ProtoReflect.Descriptor instead.
func (*GetUserListRsp) Descriptor() ([]byte, []int) {
	return file_im_protobuf_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserListRsp) GetRetCode() uint32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

func (x *GetUserListRsp) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *GetUserListRsp) GetUserId() []string {
	if x != nil {
		return x.UserId
	}
	return nil
}

var File_im_protobuf_proto protoreflect.FileDescriptor

var file_im_protobuf_proto_rawDesc = []byte{
	0x0a, 0x11, 0x69, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x6f, 0x0a,
	0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x79, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x79, 0x73, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5f,
	0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x22,
	0x8c, 0x01, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x71,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x79, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x79, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x5c,
	0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72,
	0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x22, 0xb3, 0x01, 0x0a,
	0x0d, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x71,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x79, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x79, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x22, 0x5f, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x41, 0x6c, 0x6c,
	0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65,
	0x72, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x4d, 0x73,
	0x67, 0x49, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x22, 0x5a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x32, 0x9f, 0x02, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x52, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x73,
	0x70, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x12, 0x14,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73,
	0x67, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0a,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x41, 0x6c, 0x6c, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x41, 0x6c, 0x6c, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x43,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73,
	0x70, 0x22, 0x00, 0x42, 0x2c, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x42, 0x0d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_im_protobuf_proto_rawDescOnce sync.Once
	file_im_protobuf_proto_rawDescData = file_im_protobuf_proto_rawDesc
)

func file_im_protobuf_proto_rawDescGZIP() []byte {
	file_im_protobuf_proto_rawDescOnce.Do(func() {
		file_im_protobuf_proto_rawDescData = protoimpl.X.CompressGZIP(file_im_protobuf_proto_rawDescData)
	})
	return file_im_protobuf_proto_rawDescData
}

var file_im_protobuf_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_im_protobuf_proto_goTypes = []interface{}{
	(*QueryUsersOnlineReq)(nil), // 0: protobuf.QueryUsersOnlineReq
	(*QueryUsersOnlineRsp)(nil), // 1: protobuf.QueryUsersOnlineRsp
	(*SendMsgReq)(nil),          // 2: protobuf.SendMsgReq
	(*SendMsgRsp)(nil),          // 3: protobuf.SendMsgRsp
	(*SendMsgAllReq)(nil),       // 4: protobuf.SendMsgAllReq
	(*SendMsgAllRsp)(nil),       // 5: protobuf.SendMsgAllRsp
	(*GetUserListReq)(nil),      // 6: protobuf.GetUserListReq
	(*GetUserListRsp)(nil),      // 7: protobuf.GetUserListRsp
}
var file_im_protobuf_proto_depIdxs = []int32{
	0, // 0: protobuf.AccServer.QueryUsersOnline:input_type -> protobuf.QueryUsersOnlineReq
	2, // 1: protobuf.AccServer.SendMsg:input_type -> protobuf.SendMsgReq
	4, // 2: protobuf.AccServer.SendMsgAll:input_type -> protobuf.SendMsgAllReq
	6, // 3: protobuf.AccServer.GetUserList:input_type -> protobuf.GetUserListReq
	1, // 4: protobuf.AccServer.QueryUsersOnline:output_type -> protobuf.QueryUsersOnlineRsp
	3, // 5: protobuf.AccServer.SendMsg:output_type -> protobuf.SendMsgRsp
	5, // 6: protobuf.AccServer.SendMsgAll:output_type -> protobuf.SendMsgAllRsp
	7, // 7: protobuf.AccServer.GetUserList:output_type -> protobuf.GetUserListRsp
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_im_protobuf_proto_init() }
func file_im_protobuf_proto_init() {
	if File_im_protobuf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_im_protobuf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUsersOnlineReq); i {
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
		file_im_protobuf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUsersOnlineRsp); i {
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
		file_im_protobuf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMsgReq); i {
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
		file_im_protobuf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMsgRsp); i {
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
		file_im_protobuf_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMsgAllReq); i {
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
		file_im_protobuf_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMsgAllRsp); i {
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
		file_im_protobuf_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserListReq); i {
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
		file_im_protobuf_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserListRsp); i {
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
			RawDescriptor: file_im_protobuf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_im_protobuf_proto_goTypes,
		DependencyIndexes: file_im_protobuf_proto_depIdxs,
		MessageInfos:      file_im_protobuf_proto_msgTypes,
	}.Build()
	File_im_protobuf_proto = out.File
	file_im_protobuf_proto_rawDesc = nil
	file_im_protobuf_proto_goTypes = nil
	file_im_protobuf_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AccServerClient is the client API for AccServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccServerClient interface {
	// 查询用户是否在线
	QueryUsersOnline(ctx context.Context, in *QueryUsersOnlineReq, opts ...grpc.CallOption) (*QueryUsersOnlineRsp, error)
	// 发送消息
	SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgRsp, error)
	// 发送消息
	SendMsgAll(ctx context.Context, in *SendMsgAllReq, opts ...grpc.CallOption) (*SendMsgAllRsp, error)
	// 获取用户列表
	GetUserList(ctx context.Context, in *GetUserListReq, opts ...grpc.CallOption) (*GetUserListRsp, error)
}

type accServerClient struct {
	cc grpc.ClientConnInterface
}

func NewAccServerClient(cc grpc.ClientConnInterface) AccServerClient {
	return &accServerClient{cc}
}

func (c *accServerClient) QueryUsersOnline(ctx context.Context, in *QueryUsersOnlineReq, opts ...grpc.CallOption) (*QueryUsersOnlineRsp, error) {
	out := new(QueryUsersOnlineRsp)
	err := c.cc.Invoke(ctx, "/protobuf.AccServer/QueryUsersOnline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accServerClient) SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgRsp, error) {
	out := new(SendMsgRsp)
	err := c.cc.Invoke(ctx, "/protobuf.AccServer/SendMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accServerClient) SendMsgAll(ctx context.Context, in *SendMsgAllReq, opts ...grpc.CallOption) (*SendMsgAllRsp, error) {
	out := new(SendMsgAllRsp)
	err := c.cc.Invoke(ctx, "/protobuf.AccServer/SendMsgAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accServerClient) GetUserList(ctx context.Context, in *GetUserListReq, opts ...grpc.CallOption) (*GetUserListRsp, error) {
	out := new(GetUserListRsp)
	err := c.cc.Invoke(ctx, "/protobuf.AccServer/GetUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccServerServer is the server API for AccServer service.
type AccServerServer interface {
	// 查询用户是否在线
	QueryUsersOnline(context.Context, *QueryUsersOnlineReq) (*QueryUsersOnlineRsp, error)
	// 发送消息
	SendMsg(context.Context, *SendMsgReq) (*SendMsgRsp, error)
	// 发送消息
	SendMsgAll(context.Context, *SendMsgAllReq) (*SendMsgAllRsp, error)
	// 获取用户列表
	GetUserList(context.Context, *GetUserListReq) (*GetUserListRsp, error)
}

// UnimplementedAccServerServer can be embedded to have forward compatible implementations.
type UnimplementedAccServerServer struct {
}

func (*UnimplementedAccServerServer) QueryUsersOnline(context.Context, *QueryUsersOnlineReq) (*QueryUsersOnlineRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryUsersOnline not implemented")
}
func (*UnimplementedAccServerServer) SendMsg(context.Context, *SendMsgReq) (*SendMsgRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsg not implemented")
}
func (*UnimplementedAccServerServer) SendMsgAll(context.Context, *SendMsgAllReq) (*SendMsgAllRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsgAll not implemented")
}
func (*UnimplementedAccServerServer) GetUserList(context.Context, *GetUserListReq) (*GetUserListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}

func RegisterAccServerServer(s *grpc.Server, srv AccServerServer) {
	s.RegisterService(&_AccServer_serviceDesc, srv)
}

func _AccServer_QueryUsersOnline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUsersOnlineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccServerServer).QueryUsersOnline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.AccServer/QueryUsersOnline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccServerServer).QueryUsersOnline(ctx, req.(*QueryUsersOnlineReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccServer_SendMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccServerServer).SendMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.AccServer/SendMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccServerServer).SendMsg(ctx, req.(*SendMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccServer_SendMsgAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMsgAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccServerServer).SendMsgAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.AccServer/SendMsgAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccServerServer).SendMsgAll(ctx, req.(*SendMsgAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccServer_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccServerServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.AccServer/GetUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccServerServer).GetUserList(ctx, req.(*GetUserListReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.AccServer",
	HandlerType: (*AccServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryUsersOnline",
			Handler:    _AccServer_QueryUsersOnline_Handler,
		},
		{
			MethodName: "SendMsg",
			Handler:    _AccServer_SendMsg_Handler,
		},
		{
			MethodName: "SendMsgAll",
			Handler:    _AccServer_SendMsgAll_Handler,
		},
		{
			MethodName: "GetUserList",
			Handler:    _AccServer_GetUserList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "im_protobuf.proto",
}
