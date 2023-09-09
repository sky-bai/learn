// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: 2pb/user/message.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that 2_runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserInfoGender int32

const (
	UserInfo_MALE   UserInfoGender = 0
	UserInfo_FEMALE UserInfoGender = 1
)

// Enum value maps for UserInfoGender.
var (
	UserInfoGender_name = map[int32]string{
		0: "MALE",
		1: "FEMALE",
	}
	UserInfoGender_value = map[string]int32{
		"MALE":   0,
		"FEMALE": 1,
	}
)

func (x UserInfoGender) Enum() *UserInfoGender {
	p := new(UserInfoGender)
	*p = x
	return p
}

func (x UserInfoGender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserInfoGender) Descriptor() protoreflect.EnumDescriptor {
	return file__2pb_user_message_proto_enumTypes[0].Descriptor()
}

func (UserInfoGender) Type() protoreflect.EnumType {
	return &file__2pb_user_message_proto_enumTypes[0]
}

func (x UserInfoGender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserInfoGender.Descriptor instead.
func (UserInfoGender) EnumDescriptor() ([]byte, []int) {
	return file__2pb_user_message_proto_rawDescGZIP(), []int{1, 0}
}

type UserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *UserID) Reset() {
	*x = UserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file__2pb_user_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserID) ProtoMessage() {}

func (x *UserID) ProtoReflect() protoreflect.Message {
	mi := &file__2pb_user_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserID.ProtoReflect.Descriptor instead.
func (*UserID) Descriptor() ([]byte, []int) {
	return file__2pb_user_message_proto_rawDescGZIP(), []int{0}
}

func (x *UserID) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     int64          `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name   string         `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Age    int32          `protobuf:"varint,3,opt,name=Age,proto3" json:"Age,omitempty"`
	Gender UserInfoGender `protobuf:"varint,4,opt,name=Gender,proto3,enum=user.UserInfoGender" json:"Gender,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file__2pb_user_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file__2pb_user_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file__2pb_user_message_proto_rawDescGZIP(), []int{1}
}

func (x *UserInfo) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfo) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UserInfo) GetGender() UserInfoGender {
	if x != nil {
		return x.Gender
	}
	return UserInfo_MALE
}

var File__2pb_user_message_proto protoreflect.FileDescriptor

var file__2pb_user_message_proto_rawDesc = []byte{
	0x0a, 0x16, 0x32, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x18,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x8f, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x41, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x47,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x1e, 0x0a, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x32, 0x70,
	0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file__2pb_user_message_proto_rawDescOnce sync.Once
	file__2pb_user_message_proto_rawDescData = file__2pb_user_message_proto_rawDesc
)

func file__2pb_user_message_proto_rawDescGZIP() []byte {
	file__2pb_user_message_proto_rawDescOnce.Do(func() {
		file__2pb_user_message_proto_rawDescData = protoimpl.X.CompressGZIP(file__2pb_user_message_proto_rawDescData)
	})
	return file__2pb_user_message_proto_rawDescData
}

var file__2pb_user_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file__2pb_user_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file__2pb_user_message_proto_goTypes = []interface{}{
	(UserInfoGender)(0), // 0: user.UserInfo.gender
	(*UserID)(nil),      // 1: user.UserID
	(*UserInfo)(nil),    // 2: user.UserInfo
}
var file__2pb_user_message_proto_depIdxs = []int32{
	0, // 0: user.UserInfo.Gender:type_name -> user.UserInfo.gender
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file__2pb_user_message_proto_init() }
func file__2pb_user_message_proto_init() {
	if File__2pb_user_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file__2pb_user_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserID); i {
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
		file__2pb_user_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
			RawDescriptor: file__2pb_user_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file__2pb_user_message_proto_goTypes,
		DependencyIndexes: file__2pb_user_message_proto_depIdxs,
		EnumInfos:         file__2pb_user_message_proto_enumTypes,
		MessageInfos:      file__2pb_user_message_proto_msgTypes,
	}.Build()
	File__2pb_user_message_proto = out.File
	file__2pb_user_message_proto_rawDesc = nil
	file__2pb_user_message_proto_goTypes = nil
	file__2pb_user_message_proto_depIdxs = nil
}
