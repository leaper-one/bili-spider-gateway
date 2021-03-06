// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: grpc/rank.proto

package rank

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Symbols defined in public import of google/protobuf/timestamp.proto.

type Timestamp = timestamppb.Timestamp

type GetRankRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 房间号
	RoomId int64 `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	// 时间戳 起点
	Offset *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetRankRequest) Reset() {
	*x = GetRankRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_rank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRankRequest) ProtoMessage() {}

func (x *GetRankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_rank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRankRequest.ProtoReflect.Descriptor instead.
func (*GetRankRequest) Descriptor() ([]byte, []int) {
	return file_grpc_rank_proto_rawDescGZIP(), []int{0}
}

func (x *GetRankRequest) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *GetRankRequest) GetOffset() *timestamppb.Timestamp {
	if x != nil {
		return x.Offset
	}
	return nil
}

type GetRankResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 房间号
	RoomId int64 `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	// 投稿的时间
	NoteTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=note_time,json=noteTime,proto3" json:"note_time,omitempty"`
	// 排名
	Rank int32 `protobuf:"varint,3,opt,name=rank,proto3" json:"rank,omitempty"`
}

func (x *GetRankResponse) Reset() {
	*x = GetRankResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_rank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRankResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRankResponse) ProtoMessage() {}

func (x *GetRankResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_rank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRankResponse.ProtoReflect.Descriptor instead.
func (*GetRankResponse) Descriptor() ([]byte, []int) {
	return file_grpc_rank_proto_rawDescGZIP(), []int{1}
}

func (x *GetRankResponse) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *GetRankResponse) GetNoteTime() *timestamppb.Timestamp {
	if x != nil {
		return x.NoteTime
	}
	return nil
}

func (x *GetRankResponse) GetRank() int32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

type HasNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 房间号
	RoomId int64 `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	// B 站 ID
	Buid int64 `protobuf:"varint,2,opt,name=buid,proto3" json:"buid,omitempty"`
}

func (x *HasNoteRequest) Reset() {
	*x = HasNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_rank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HasNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HasNoteRequest) ProtoMessage() {}

func (x *HasNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_rank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HasNoteRequest.ProtoReflect.Descriptor instead.
func (*HasNoteRequest) Descriptor() ([]byte, []int) {
	return file_grpc_rank_proto_rawDescGZIP(), []int{2}
}

func (x *HasNoteRequest) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *HasNoteRequest) GetBuid() int64 {
	if x != nil {
		return x.Buid
	}
	return 0
}

type HasNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 状态码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 提示信息
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HasNoteResponse) Reset() {
	*x = HasNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_rank_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HasNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HasNoteResponse) ProtoMessage() {}

func (x *HasNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_rank_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HasNoteResponse.ProtoReflect.Descriptor instead.
func (*HasNoteResponse) Descriptor() ([]byte, []int) {
	return file_grpc_rank_proto_rawDescGZIP(), []int{3}
}

func (x *HasNoteResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *HasNoteResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type HandInRankInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 房间号
	RoomId int64 `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	// B 站 ID
	Buid int64 `protobuf:"varint,2,opt,name=buid,proto3" json:"buid,omitempty"`
	// 礼物权重
	GiftWeight int32 `protobuf:"varint,3,opt,name=gift_weight,json=giftWeight,proto3" json:"gift_weight,omitempty"`
}

func (x *HandInRankInfoRequest) Reset() {
	*x = HandInRankInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_rank_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandInRankInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandInRankInfoRequest) ProtoMessage() {}

func (x *HandInRankInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_rank_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandInRankInfoRequest.ProtoReflect.Descriptor instead.
func (*HandInRankInfoRequest) Descriptor() ([]byte, []int) {
	return file_grpc_rank_proto_rawDescGZIP(), []int{4}
}

func (x *HandInRankInfoRequest) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *HandInRankInfoRequest) GetBuid() int64 {
	if x != nil {
		return x.Buid
	}
	return 0
}

func (x *HandInRankInfoRequest) GetGiftWeight() int32 {
	if x != nil {
		return x.GiftWeight
	}
	return 0
}

type HandInRankInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 状态码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 提示信息
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HandInRankInfoResponse) Reset() {
	*x = HandInRankInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_rank_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandInRankInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandInRankInfoResponse) ProtoMessage() {}

func (x *HandInRankInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_rank_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandInRankInfoResponse.ProtoReflect.Descriptor instead.
func (*HandInRankInfoResponse) Descriptor() ([]byte, []int) {
	return file_grpc_rank_proto_rawDescGZIP(), []int{5}
}

func (x *HandInRankInfoResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *HandInRankInfoResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_grpc_rank_proto protoreflect.FileDescriptor

var file_grpc_rank_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52,
	0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f,
	0x6d, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x77, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x61,
	0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f,
	0x6d, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x08, 0x6e, 0x6f, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x61, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b,
	0x22, 0x3d, 0x0a, 0x0e, 0x48, 0x61, 0x73, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x62, 0x75, 0x69, 0x64, 0x22,
	0x37, 0x0a, 0x0f, 0x48, 0x61, 0x73, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x65, 0x0a, 0x15, 0x48, 0x61, 0x6e, 0x64,
	0x49, 0x6e, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x75,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x62, 0x75, 0x69, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x67, 0x69, 0x66, 0x74, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x67, 0x69, 0x66, 0x74, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22,
	0x3e, 0x0a, 0x16, 0x48, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32,
	0xd4, 0x01, 0x0a, 0x0b, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x14, 0x2e, 0x72, 0x61, 0x6e,
	0x6b, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x38, 0x0a, 0x07, 0x48,
	0x61, 0x73, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x48, 0x61,
	0x73, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x72,
	0x61, 0x6e, 0x6b, 0x2e, 0x48, 0x61, 0x73, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0e, 0x48, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x52,
	0x61, 0x6e, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x48,
	0x61, 0x6e, 0x64, 0x49, 0x6e, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x48, 0x61, 0x6e, 0x64,
	0x49, 0x6e, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x50, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_rank_proto_rawDescOnce sync.Once
	file_grpc_rank_proto_rawDescData = file_grpc_rank_proto_rawDesc
)

func file_grpc_rank_proto_rawDescGZIP() []byte {
	file_grpc_rank_proto_rawDescOnce.Do(func() {
		file_grpc_rank_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_rank_proto_rawDescData)
	})
	return file_grpc_rank_proto_rawDescData
}

var file_grpc_rank_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_grpc_rank_proto_goTypes = []interface{}{
	(*GetRankRequest)(nil),         // 0: rank.GetRankRequest
	(*GetRankResponse)(nil),        // 1: rank.GetRankResponse
	(*HasNoteRequest)(nil),         // 2: rank.HasNoteRequest
	(*HasNoteResponse)(nil),        // 3: rank.HasNoteResponse
	(*HandInRankInfoRequest)(nil),  // 4: rank.HandInRankInfoRequest
	(*HandInRankInfoResponse)(nil), // 5: rank.HandInRankInfoResponse
	(*timestamppb.Timestamp)(nil),  // 6: google.protobuf.Timestamp
}
var file_grpc_rank_proto_depIdxs = []int32{
	6, // 0: rank.GetRankRequest.offset:type_name -> google.protobuf.Timestamp
	6, // 1: rank.GetRankResponse.note_time:type_name -> google.protobuf.Timestamp
	0, // 2: rank.RankService.GetRank:input_type -> rank.GetRankRequest
	2, // 3: rank.RankService.HasNote:input_type -> rank.HasNoteRequest
	4, // 4: rank.RankService.HandInRankInfo:input_type -> rank.HandInRankInfoRequest
	1, // 5: rank.RankService.GetRank:output_type -> rank.GetRankResponse
	3, // 6: rank.RankService.HasNote:output_type -> rank.HasNoteResponse
	5, // 7: rank.RankService.HandInRankInfo:output_type -> rank.HandInRankInfoResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_grpc_rank_proto_init() }
func file_grpc_rank_proto_init() {
	if File_grpc_rank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_rank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRankRequest); i {
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
		file_grpc_rank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRankResponse); i {
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
		file_grpc_rank_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HasNoteRequest); i {
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
		file_grpc_rank_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HasNoteResponse); i {
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
		file_grpc_rank_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandInRankInfoRequest); i {
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
		file_grpc_rank_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandInRankInfoResponse); i {
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
			RawDescriptor: file_grpc_rank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_rank_proto_goTypes,
		DependencyIndexes: file_grpc_rank_proto_depIdxs,
		MessageInfos:      file_grpc_rank_proto_msgTypes,
	}.Build()
	File_grpc_rank_proto = out.File
	file_grpc_rank_proto_rawDesc = nil
	file_grpc_rank_proto_goTypes = nil
	file_grpc_rank_proto_depIdxs = nil
}
