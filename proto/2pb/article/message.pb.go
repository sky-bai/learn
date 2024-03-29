// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: 2pb/article/message.proto

package article

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

type Articles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Articles []*Articles_Article `protobuf:"bytes,1,rep,name=Articles,proto3" json:"Articles,omitempty"`
}

func (x *Articles) Reset() {
	*x = Articles{}
	if protoimpl.UnsafeEnabled {
		mi := &file__2pb_article_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Articles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Articles) ProtoMessage() {}

func (x *Articles) ProtoReflect() protoreflect.Message {
	mi := &file__2pb_article_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Articles.ProtoReflect.Descriptor instead.
func (*Articles) Descriptor() ([]byte, []int) {
	return file__2pb_article_message_proto_rawDescGZIP(), []int{0}
}

func (x *Articles) GetArticles() []*Articles_Article {
	if x != nil {
		return x.Articles
	}
	return nil
}

type Articles_Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
}

func (x *Articles_Article) Reset() {
	*x = Articles_Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file__2pb_article_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Articles_Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Articles_Article) ProtoMessage() {}

func (x *Articles_Article) ProtoReflect() protoreflect.Message {
	mi := &file__2pb_article_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Articles_Article.ProtoReflect.Descriptor instead.
func (*Articles_Article) Descriptor() ([]byte, []int) {
	return file__2pb_article_message_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Articles_Article) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Articles_Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File__2pb_article_message_proto protoreflect.FileDescriptor

var file__2pb_article_message_proto_rawDesc = []byte{
	0x0a, 0x19, 0x32, 0x70, 0x62, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x22, 0x72, 0x0a, 0x08, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x12, 0x35, 0x0a, 0x08, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x08, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x1a, 0x2f, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x15, 0x5a, 0x13, 0x32, 0x70, 0x62, 0x2f,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x3b, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file__2pb_article_message_proto_rawDescOnce sync.Once
	file__2pb_article_message_proto_rawDescData = file__2pb_article_message_proto_rawDesc
)

func file__2pb_article_message_proto_rawDescGZIP() []byte {
	file__2pb_article_message_proto_rawDescOnce.Do(func() {
		file__2pb_article_message_proto_rawDescData = protoimpl.X.CompressGZIP(file__2pb_article_message_proto_rawDescData)
	})
	return file__2pb_article_message_proto_rawDescData
}

var file__2pb_article_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file__2pb_article_message_proto_goTypes = []interface{}{
	(*Articles)(nil),         // 0: article.Articles
	(*Articles_Article)(nil), // 1: article.Articles.Article
}
var file__2pb_article_message_proto_depIdxs = []int32{
	1, // 0: article.Articles.Articles:type_name -> article.Articles.Article
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file__2pb_article_message_proto_init() }
func file__2pb_article_message_proto_init() {
	if File__2pb_article_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file__2pb_article_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Articles); i {
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
		file__2pb_article_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Articles_Article); i {
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
			RawDescriptor: file__2pb_article_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file__2pb_article_message_proto_goTypes,
		DependencyIndexes: file__2pb_article_message_proto_depIdxs,
		MessageInfos:      file__2pb_article_message_proto_msgTypes,
	}.Build()
	File__2pb_article_message_proto = out.File
	file__2pb_article_message_proto_rawDesc = nil
	file__2pb_article_message_proto_goTypes = nil
	file__2pb_article_message_proto_depIdxs = nil
}
