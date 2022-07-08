// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.17.3
// source: profile/v1/stream.proto

package profilev1

import (
	proto "github.com/golang/protobuf/proto"
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

// Event sent when a user's profile is updated.
//
// Servers should sent this event only to users that can "see" (eg. they are
// in the same guild) the user this event was triggered by.
type ProfileUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User ID of the user that had it's profile updated.
	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// New username for this user.
	NewUsername *string `protobuf:"bytes,2,opt,name=new_username,json=newUsername,proto3,oneof" json:"new_username,omitempty"`
	// New avatar for this user.
	NewAvatar *string `protobuf:"bytes,3,opt,name=new_avatar,json=newAvatar,proto3,oneof" json:"new_avatar,omitempty"`
	// New status for this user.
	NewStatus *UserStatus `protobuf:"varint,4,opt,name=new_status,json=newStatus,proto3,enum=protocol.profile.v1.UserStatus,oneof" json:"new_status,omitempty"`
	// New is bot or not for this user.
	NewIsBot *bool `protobuf:"varint,5,opt,name=new_is_bot,json=newIsBot,proto3,oneof" json:"new_is_bot,omitempty"`
}

func (x *ProfileUpdated) Reset() {
	*x = ProfileUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_v1_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileUpdated) ProtoMessage() {}

func (x *ProfileUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileUpdated.ProtoReflect.Descriptor instead.
func (*ProfileUpdated) Descriptor() ([]byte, []int) {
	return file_profile_v1_stream_proto_rawDescGZIP(), []int{0}
}

func (x *ProfileUpdated) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ProfileUpdated) GetNewUsername() string {
	if x != nil && x.NewUsername != nil {
		return *x.NewUsername
	}
	return ""
}

func (x *ProfileUpdated) GetNewAvatar() string {
	if x != nil && x.NewAvatar != nil {
		return *x.NewAvatar
	}
	return ""
}

func (x *ProfileUpdated) GetNewStatus() UserStatus {
	if x != nil && x.NewStatus != nil {
		return *x.NewStatus
	}
	return UserStatus_USER_STATUS_OFFLINE_UNSPECIFIED
}

func (x *ProfileUpdated) GetNewIsBot() bool {
	if x != nil && x.NewIsBot != nil {
		return *x.NewIsBot
	}
	return false
}

// Describes an emote service event.
type StreamEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The event type.
	//
	// Types that are assignable to Event:
	//	*StreamEvent_ProfileUpdated
	Event isStreamEvent_Event `protobuf_oneof:"event"`
}

func (x *StreamEvent) Reset() {
	*x = StreamEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_v1_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamEvent) ProtoMessage() {}

func (x *StreamEvent) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamEvent.ProtoReflect.Descriptor instead.
func (*StreamEvent) Descriptor() ([]byte, []int) {
	return file_profile_v1_stream_proto_rawDescGZIP(), []int{1}
}

func (m *StreamEvent) GetEvent() isStreamEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *StreamEvent) GetProfileUpdated() *ProfileUpdated {
	if x, ok := x.GetEvent().(*StreamEvent_ProfileUpdated); ok {
		return x.ProfileUpdated
	}
	return nil
}

type isStreamEvent_Event interface {
	isStreamEvent_Event()
}

type StreamEvent_ProfileUpdated struct {
	// Send the profile updated event.
	ProfileUpdated *ProfileUpdated `protobuf:"bytes,14,opt,name=profile_updated,json=profileUpdated,proto3,oneof"`
}

func (*StreamEvent_ProfileUpdated) isStreamEvent_Event() {}

var File_profile_v1_stream_proto protoreflect.FileDescriptor

var file_profile_v1_stream_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x16,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x26, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x6e, 0x65,
	0x77, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x09, 0x6e, 0x65, 0x77, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x88, 0x01, 0x01, 0x12, 0x43,
	0x0a, 0x0a, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x48, 0x02, 0x52, 0x09, 0x6e, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x0a, 0x6e, 0x65, 0x77, 0x5f, 0x69, 0x73, 0x5f, 0x62, 0x6f,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x03, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x49, 0x73,
	0x42, 0x6f, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6e, 0x65, 0x77, 0x5f,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x69, 0x73,
	0x5f, 0x62, 0x6f, 0x74, 0x22, 0x66, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x4e, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x48, 0x00, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0xd6, 0x01, 0x0a,
	0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x2d, 0x64, 0x65, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x68, 0x69, 0x62, 0x73, 0x68, 0x69, 0x62,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x50, 0x58, 0xaa,
	0x02, 0x13, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x13, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x5c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1f, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_profile_v1_stream_proto_rawDescOnce sync.Once
	file_profile_v1_stream_proto_rawDescData = file_profile_v1_stream_proto_rawDesc
)

func file_profile_v1_stream_proto_rawDescGZIP() []byte {
	file_profile_v1_stream_proto_rawDescOnce.Do(func() {
		file_profile_v1_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_profile_v1_stream_proto_rawDescData)
	})
	return file_profile_v1_stream_proto_rawDescData
}

var file_profile_v1_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_profile_v1_stream_proto_goTypes = []interface{}{
	(*ProfileUpdated)(nil), // 0: protocol.profile.v1.ProfileUpdated
	(*StreamEvent)(nil),    // 1: protocol.profile.v1.StreamEvent
	(UserStatus)(0),        // 2: protocol.profile.v1.UserStatus
}
var file_profile_v1_stream_proto_depIdxs = []int32{
	2, // 0: protocol.profile.v1.ProfileUpdated.new_status:type_name -> protocol.profile.v1.UserStatus
	0, // 1: protocol.profile.v1.StreamEvent.profile_updated:type_name -> protocol.profile.v1.ProfileUpdated
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_profile_v1_stream_proto_init() }
func file_profile_v1_stream_proto_init() {
	if File_profile_v1_stream_proto != nil {
		return
	}
	file_profile_v1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_profile_v1_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileUpdated); i {
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
		file_profile_v1_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamEvent); i {
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
	file_profile_v1_stream_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_profile_v1_stream_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*StreamEvent_ProfileUpdated)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_profile_v1_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_profile_v1_stream_proto_goTypes,
		DependencyIndexes: file_profile_v1_stream_proto_depIdxs,
		MessageInfos:      file_profile_v1_stream_proto_msgTypes,
	}.Build()
	File_profile_v1_stream_proto = out.File
	file_profile_v1_stream_proto_rawDesc = nil
	file_profile_v1_stream_proto_goTypes = nil
	file_profile_v1_stream_proto_depIdxs = nil
}
