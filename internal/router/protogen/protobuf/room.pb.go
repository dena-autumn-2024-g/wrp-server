// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: protobuf/room.proto

package protogen

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

// 部屋の作成リクエストとレスポンス
type CreateRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateRoomRequest) Reset() {
	*x = CreateRoomRequest{}
	mi := &file_protobuf_room_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomRequest) ProtoMessage() {}

func (x *CreateRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_room_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomRequest.ProtoReflect.Descriptor instead.
func (*CreateRoomRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_room_proto_rawDescGZIP(), []int{0}
}

type CreateRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomUrl string `protobuf:"bytes,1,opt,name=room_url,json=roomUrl,proto3" json:"room_url,omitempty"` // 入室用のURL
	RoomId  string `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`    // (UUID)
}

func (x *CreateRoomResponse) Reset() {
	*x = CreateRoomResponse{}
	mi := &file_protobuf_room_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomResponse) ProtoMessage() {}

func (x *CreateRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_room_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomResponse.ProtoReflect.Descriptor instead.
func (*CreateRoomResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_room_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRoomResponse) GetRoomUrl() string {
	if x != nil {
		return x.RoomUrl
	}
	return ""
}

func (x *CreateRoomResponse) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

// 部屋を閉じるリクエスト
type CloseRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId string `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"` // ルームID
}

func (x *CloseRoomRequest) Reset() {
	*x = CloseRoomRequest{}
	mi := &file_protobuf_room_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CloseRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseRoomRequest) ProtoMessage() {}

func (x *CloseRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_room_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseRoomRequest.ProtoReflect.Descriptor instead.
func (*CloseRoomRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_room_proto_rawDescGZIP(), []int{2}
}

func (x *CloseRoomRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

// 部屋にユーザーが参加するのを待つリクエスト
type WaitForUserJoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId string `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"` // ルームID
}

func (x *WaitForUserJoinRequest) Reset() {
	*x = WaitForUserJoinRequest{}
	mi := &file_protobuf_room_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WaitForUserJoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaitForUserJoinRequest) ProtoMessage() {}

func (x *WaitForUserJoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_room_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaitForUserJoinRequest.ProtoReflect.Descriptor instead.
func (*WaitForUserJoinRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_room_proto_rawDescGZIP(), []int{3}
}

func (x *WaitForUserJoinRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

// 部屋にユーザーが参加した通知メッセージ
type WaitForUserJoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 参加したユーザーID
}

func (x *WaitForUserJoinResponse) Reset() {
	*x = WaitForUserJoinResponse{}
	mi := &file_protobuf_room_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WaitForUserJoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaitForUserJoinResponse) ProtoMessage() {}

func (x *WaitForUserJoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_room_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaitForUserJoinResponse.ProtoReflect.Descriptor instead.
func (*WaitForUserJoinResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_room_proto_rawDescGZIP(), []int{4}
}

func (x *WaitForUserJoinResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type JoinRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId string `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"` // ルームID
}

func (x *JoinRoomRequest) Reset() {
	*x = JoinRoomRequest{}
	mi := &file_protobuf_room_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRoomRequest) ProtoMessage() {}

func (x *JoinRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_room_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRoomRequest.ProtoReflect.Descriptor instead.
func (*JoinRoomRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_room_proto_rawDescGZIP(), []int{5}
}

func (x *JoinRoomRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type JoinRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // ユーザーID
}

func (x *JoinRoomResponse) Reset() {
	*x = JoinRoomResponse{}
	mi := &file_protobuf_room_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRoomResponse) ProtoMessage() {}

func (x *JoinRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_room_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRoomResponse.ProtoReflect.Descriptor instead.
func (*JoinRoomResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_room_proto_rawDescGZIP(), []int{6}
}

func (x *JoinRoomResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CloseRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CloseRoomResponse) Reset() {
	*x = CloseRoomResponse{}
	mi := &file_protobuf_room_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CloseRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseRoomResponse) ProtoMessage() {}

func (x *CloseRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_room_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseRoomResponse.ProtoReflect.Descriptor instead.
func (*CloseRoomResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_room_proto_rawDescGZIP(), []int{7}
}

var File_protobuf_room_proto protoreflect.FileDescriptor

var file_protobuf_room_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x77, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x69, 0x6e,
	0x67, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x48, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x72, 0x6f, 0x6f, 0x6d, 0x55, 0x72, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x22, 0x2b, 0x0a, 0x10, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0x31, 0x0a,
	0x16, 0x57, 0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x6f, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x22, 0x32, 0x0a, 0x17, 0x57, 0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4a,
	0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x0f, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x22, 0x2b, 0x0a, 0x10, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x13, 0x0a,
	0x11, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xc9, 0x02, 0x0a, 0x0b, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d,
	0x12, 0x1d, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x5c, 0x0a, 0x0f, 0x57, 0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x6f,
	0x69, 0x6e, 0x12, 0x22, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x57, 0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x6f, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x57, 0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4a,
	0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x45, 0x0a,
	0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x1b, 0x2e, 0x77, 0x61, 0x74, 0x65,
	0x72, 0x5f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x6f, 0x6f,
	0x6d, 0x12, 0x1c, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x43,
	0x6c, 0x6f, 0x73, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xb4,
	0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x69, 0x6e,
	0x67, 0x42, 0x09, 0x52, 0x6f, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x6e, 0x61, 0x2d,
	0x61, 0x75, 0x74, 0x75, 0x6d, 0x6e, 0x2d, 0x32, 0x30, 0x32, 0x34, 0x2d, 0x67, 0x2f, 0x77, 0x72,
	0x70, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x67, 0x65, 0x6e, 0xa2, 0x02, 0x03, 0x57, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x57, 0x61, 0x74, 0x65,
	0x72, 0x52, 0x69, 0x6e, 0x67, 0xca, 0x02, 0x09, 0x57, 0x61, 0x74, 0x65, 0x72, 0x52, 0x69, 0x6e,
	0x67, 0xe2, 0x02, 0x15, 0x57, 0x61, 0x74, 0x65, 0x72, 0x52, 0x69, 0x6e, 0x67, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x57, 0x61, 0x74, 0x65,
	0x72, 0x52, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_room_proto_rawDescOnce sync.Once
	file_protobuf_room_proto_rawDescData = file_protobuf_room_proto_rawDesc
)

func file_protobuf_room_proto_rawDescGZIP() []byte {
	file_protobuf_room_proto_rawDescOnce.Do(func() {
		file_protobuf_room_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_room_proto_rawDescData)
	})
	return file_protobuf_room_proto_rawDescData
}

var file_protobuf_room_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protobuf_room_proto_goTypes = []any{
	(*CreateRoomRequest)(nil),       // 0: water_ring.CreateRoomRequest
	(*CreateRoomResponse)(nil),      // 1: water_ring.CreateRoomResponse
	(*CloseRoomRequest)(nil),        // 2: water_ring.CloseRoomRequest
	(*WaitForUserJoinRequest)(nil),  // 3: water_ring.WaitForUserJoinRequest
	(*WaitForUserJoinResponse)(nil), // 4: water_ring.WaitForUserJoinResponse
	(*JoinRoomRequest)(nil),         // 5: water_ring.JoinRoomRequest
	(*JoinRoomResponse)(nil),        // 6: water_ring.JoinRoomResponse
	(*CloseRoomResponse)(nil),       // 7: water_ring.CloseRoomResponse
}
var file_protobuf_room_proto_depIdxs = []int32{
	0, // 0: water_ring.RoomService.CreateRoom:input_type -> water_ring.CreateRoomRequest
	3, // 1: water_ring.RoomService.WaitForUserJoin:input_type -> water_ring.WaitForUserJoinRequest
	5, // 2: water_ring.RoomService.JoinRoom:input_type -> water_ring.JoinRoomRequest
	2, // 3: water_ring.RoomService.CloseRoom:input_type -> water_ring.CloseRoomRequest
	1, // 4: water_ring.RoomService.CreateRoom:output_type -> water_ring.CreateRoomResponse
	4, // 5: water_ring.RoomService.WaitForUserJoin:output_type -> water_ring.WaitForUserJoinResponse
	6, // 6: water_ring.RoomService.JoinRoom:output_type -> water_ring.JoinRoomResponse
	7, // 7: water_ring.RoomService.CloseRoom:output_type -> water_ring.CloseRoomResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_room_proto_init() }
func file_protobuf_room_proto_init() {
	if File_protobuf_room_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_room_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_room_proto_goTypes,
		DependencyIndexes: file_protobuf_room_proto_depIdxs,
		MessageInfos:      file_protobuf_room_proto_msgTypes,
	}.Build()
	File_protobuf_room_proto = out.File
	file_protobuf_room_proto_rawDesc = nil
	file_protobuf_room_proto_goTypes = nil
	file_protobuf_room_proto_depIdxs = nil
}