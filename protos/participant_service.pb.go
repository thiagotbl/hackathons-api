// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.3
// source: protos/participant_service.proto

package protos

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

type ListHackathonsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListHackathonsRequest) Reset() {
	*x = ListHackathonsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_participant_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHackathonsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHackathonsRequest) ProtoMessage() {}

func (x *ListHackathonsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_participant_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHackathonsRequest.ProtoReflect.Descriptor instead.
func (*ListHackathonsRequest) Descriptor() ([]byte, []int) {
	return file_protos_participant_service_proto_rawDescGZIP(), []int{0}
}

type ListHackathonsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hackathons []*Hackathon `protobuf:"bytes,1,rep,name=hackathons,proto3" json:"hackathons,omitempty"`
}

func (x *ListHackathonsResponse) Reset() {
	*x = ListHackathonsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_participant_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHackathonsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHackathonsResponse) ProtoMessage() {}

func (x *ListHackathonsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_participant_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHackathonsResponse.ProtoReflect.Descriptor instead.
func (*ListHackathonsResponse) Descriptor() ([]byte, []int) {
	return file_protos_participant_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListHackathonsResponse) GetHackathons() []*Hackathon {
	if x != nil {
		return x.Hackathons
	}
	return nil
}

type SubscribeTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Team *Team `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
}

func (x *SubscribeTeamRequest) Reset() {
	*x = SubscribeTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_participant_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeTeamRequest) ProtoMessage() {}

func (x *SubscribeTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_participant_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeTeamRequest.ProtoReflect.Descriptor instead.
func (*SubscribeTeamRequest) Descriptor() ([]byte, []int) {
	return file_protos_participant_service_proto_rawDescGZIP(), []int{2}
}

func (x *SubscribeTeamRequest) GetTeam() *Team {
	if x != nil {
		return x.Team
	}
	return nil
}

type SubscribeTeamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *SubscribeTeamResponse) Reset() {
	*x = SubscribeTeamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_participant_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeTeamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeTeamResponse) ProtoMessage() {}

func (x *SubscribeTeamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_participant_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeTeamResponse.ProtoReflect.Descriptor instead.
func (*SubscribeTeamResponse) Descriptor() ([]byte, []int) {
	return file_protos_participant_service_proto_rawDescGZIP(), []int{3}
}

func (x *SubscribeTeamResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type CancelTeamSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *CancelTeamSubscriptionRequest) Reset() {
	*x = CancelTeamSubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_participant_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelTeamSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelTeamSubscriptionRequest) ProtoMessage() {}

func (x *CancelTeamSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_participant_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelTeamSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*CancelTeamSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_protos_participant_service_proto_rawDescGZIP(), []int{4}
}

func (x *CancelTeamSubscriptionRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type CancelTeamSubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CancelTeamSubscriptionResponse) Reset() {
	*x = CancelTeamSubscriptionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_participant_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelTeamSubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelTeamSubscriptionResponse) ProtoMessage() {}

func (x *CancelTeamSubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_participant_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelTeamSubscriptionResponse.ProtoReflect.Descriptor instead.
func (*CancelTeamSubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_protos_participant_service_proto_rawDescGZIP(), []int{5}
}

func (x *CancelTeamSubscriptionResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_protos_participant_service_proto protoreflect.FileDescriptor

var file_protos_participant_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x68, 0x61,
	0x63, 0x6b, 0x61, 0x74, 0x68, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x17, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x61, 0x63, 0x6b, 0x61, 0x74, 0x68,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x44, 0x0a, 0x16, 0x4c, 0x69,
	0x73, 0x74, 0x48, 0x61, 0x63, 0x6b, 0x61, 0x74, 0x68, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x0a, 0x68, 0x61, 0x63, 0x6b, 0x61, 0x74, 0x68, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x48, 0x61, 0x63, 0x6b, 0x61,
	0x74, 0x68, 0x6f, 0x6e, 0x52, 0x0a, 0x68, 0x61, 0x63, 0x6b, 0x61, 0x74, 0x68, 0x6f, 0x6e, 0x73,
	0x22, 0x31, 0x0a, 0x14, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x04, 0x74,
	0x65, 0x61, 0x6d, 0x22, 0x35, 0x0a, 0x15, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x2f, 0x0a, 0x1d, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x3e, 0x0a, 0x1e, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xf8, 0x01, 0x0a, 0x12,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x43, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x61, 0x63, 0x6b, 0x61, 0x74,
	0x68, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x61, 0x63, 0x6b, 0x61,
	0x74, 0x68, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x48, 0x61, 0x63, 0x6b, 0x61, 0x74, 0x68, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x16, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x65, 0x61, 0x6d,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x65, 0x61, 0x6d,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x69, 0x61, 0x67, 0x6f, 0x74, 0x62, 0x6c, 0x2f, 0x68,
	0x61, 0x63, 0x6b, 0x61, 0x74, 0x68, 0x6f, 0x6e, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_participant_service_proto_rawDescOnce sync.Once
	file_protos_participant_service_proto_rawDescData = file_protos_participant_service_proto_rawDesc
)

func file_protos_participant_service_proto_rawDescGZIP() []byte {
	file_protos_participant_service_proto_rawDescOnce.Do(func() {
		file_protos_participant_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_participant_service_proto_rawDescData)
	})
	return file_protos_participant_service_proto_rawDescData
}

var file_protos_participant_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_participant_service_proto_goTypes = []interface{}{
	(*ListHackathonsRequest)(nil),          // 0: ListHackathonsRequest
	(*ListHackathonsResponse)(nil),         // 1: ListHackathonsResponse
	(*SubscribeTeamRequest)(nil),           // 2: SubscribeTeamRequest
	(*SubscribeTeamResponse)(nil),          // 3: SubscribeTeamResponse
	(*CancelTeamSubscriptionRequest)(nil),  // 4: CancelTeamSubscriptionRequest
	(*CancelTeamSubscriptionResponse)(nil), // 5: CancelTeamSubscriptionResponse
	(*Hackathon)(nil),                      // 6: Hackathon
	(*Team)(nil),                           // 7: Team
	(*Error)(nil),                          // 8: Error
}
var file_protos_participant_service_proto_depIdxs = []int32{
	6, // 0: ListHackathonsResponse.hackathons:type_name -> Hackathon
	7, // 1: SubscribeTeamRequest.team:type_name -> Team
	8, // 2: SubscribeTeamResponse.error:type_name -> Error
	8, // 3: CancelTeamSubscriptionResponse.error:type_name -> Error
	0, // 4: ParticipantService.ListHackathons:input_type -> ListHackathonsRequest
	2, // 5: ParticipantService.SubscribeTeam:input_type -> SubscribeTeamRequest
	4, // 6: ParticipantService.CancelTeamSubscription:input_type -> CancelTeamSubscriptionRequest
	1, // 7: ParticipantService.ListHackathons:output_type -> ListHackathonsResponse
	3, // 8: ParticipantService.SubscribeTeam:output_type -> SubscribeTeamResponse
	5, // 9: ParticipantService.CancelTeamSubscription:output_type -> CancelTeamSubscriptionResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protos_participant_service_proto_init() }
func file_protos_participant_service_proto_init() {
	if File_protos_participant_service_proto != nil {
		return
	}
	file_protos_team_proto_init()
	file_protos_hackathon_proto_init()
	file_protos_error_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_participant_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHackathonsRequest); i {
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
		file_protos_participant_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHackathonsResponse); i {
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
		file_protos_participant_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeTeamRequest); i {
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
		file_protos_participant_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeTeamResponse); i {
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
		file_protos_participant_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelTeamSubscriptionRequest); i {
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
		file_protos_participant_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelTeamSubscriptionResponse); i {
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
			RawDescriptor: file_protos_participant_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_participant_service_proto_goTypes,
		DependencyIndexes: file_protos_participant_service_proto_depIdxs,
		MessageInfos:      file_protos_participant_service_proto_msgTypes,
	}.Build()
	File_protos_participant_service_proto = out.File
	file_protos_participant_service_proto_rawDesc = nil
	file_protos_participant_service_proto_goTypes = nil
	file_protos_participant_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ParticipantServiceClient is the client API for ParticipantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ParticipantServiceClient interface {
	ListHackathons(ctx context.Context, in *ListHackathonsRequest, opts ...grpc.CallOption) (*ListHackathonsResponse, error)
	SubscribeTeam(ctx context.Context, in *SubscribeTeamRequest, opts ...grpc.CallOption) (*SubscribeTeamResponse, error)
	CancelTeamSubscription(ctx context.Context, in *CancelTeamSubscriptionRequest, opts ...grpc.CallOption) (*CancelTeamSubscriptionResponse, error)
}

type participantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewParticipantServiceClient(cc grpc.ClientConnInterface) ParticipantServiceClient {
	return &participantServiceClient{cc}
}

func (c *participantServiceClient) ListHackathons(ctx context.Context, in *ListHackathonsRequest, opts ...grpc.CallOption) (*ListHackathonsResponse, error) {
	out := new(ListHackathonsResponse)
	err := c.cc.Invoke(ctx, "/ParticipantService/ListHackathons", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participantServiceClient) SubscribeTeam(ctx context.Context, in *SubscribeTeamRequest, opts ...grpc.CallOption) (*SubscribeTeamResponse, error) {
	out := new(SubscribeTeamResponse)
	err := c.cc.Invoke(ctx, "/ParticipantService/SubscribeTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participantServiceClient) CancelTeamSubscription(ctx context.Context, in *CancelTeamSubscriptionRequest, opts ...grpc.CallOption) (*CancelTeamSubscriptionResponse, error) {
	out := new(CancelTeamSubscriptionResponse)
	err := c.cc.Invoke(ctx, "/ParticipantService/CancelTeamSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParticipantServiceServer is the server API for ParticipantService service.
type ParticipantServiceServer interface {
	ListHackathons(context.Context, *ListHackathonsRequest) (*ListHackathonsResponse, error)
	SubscribeTeam(context.Context, *SubscribeTeamRequest) (*SubscribeTeamResponse, error)
	CancelTeamSubscription(context.Context, *CancelTeamSubscriptionRequest) (*CancelTeamSubscriptionResponse, error)
}

// UnimplementedParticipantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedParticipantServiceServer struct {
}

func (*UnimplementedParticipantServiceServer) ListHackathons(context.Context, *ListHackathonsRequest) (*ListHackathonsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHackathons not implemented")
}
func (*UnimplementedParticipantServiceServer) SubscribeTeam(context.Context, *SubscribeTeamRequest) (*SubscribeTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeTeam not implemented")
}
func (*UnimplementedParticipantServiceServer) CancelTeamSubscription(context.Context, *CancelTeamSubscriptionRequest) (*CancelTeamSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTeamSubscription not implemented")
}

func RegisterParticipantServiceServer(s *grpc.Server, srv ParticipantServiceServer) {
	s.RegisterService(&_ParticipantService_serviceDesc, srv)
}

func _ParticipantService_ListHackathons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHackathonsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipantServiceServer).ListHackathons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ParticipantService/ListHackathons",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipantServiceServer).ListHackathons(ctx, req.(*ListHackathonsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipantService_SubscribeTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipantServiceServer).SubscribeTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ParticipantService/SubscribeTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipantServiceServer).SubscribeTeam(ctx, req.(*SubscribeTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipantService_CancelTeamSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelTeamSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipantServiceServer).CancelTeamSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ParticipantService/CancelTeamSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipantServiceServer).CancelTeamSubscription(ctx, req.(*CancelTeamSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ParticipantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ParticipantService",
	HandlerType: (*ParticipantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListHackathons",
			Handler:    _ParticipantService_ListHackathons_Handler,
		},
		{
			MethodName: "SubscribeTeam",
			Handler:    _ParticipantService_SubscribeTeam_Handler,
		},
		{
			MethodName: "CancelTeamSubscription",
			Handler:    _ParticipantService_CancelTeamSubscription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/participant_service.proto",
}