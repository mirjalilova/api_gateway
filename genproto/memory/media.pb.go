// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: media.proto

package memory

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

type Media struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MemoryId  string `protobuf:"bytes,2,opt,name=memory_id,json=memoryId,proto3" json:"memory_id,omitempty"`
	Type      string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Url       string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Media) Reset() {
	*x = Media{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Media) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Media) ProtoMessage() {}

func (x *Media) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Media.ProtoReflect.Descriptor instead.
func (*Media) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{0}
}

func (x *Media) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Media) GetMemoryId() string {
	if x != nil {
		return x.MemoryId
	}
	return ""
}

func (x *Media) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Media) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Media) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type MediaCreateBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemoryId string `protobuf:"bytes,1,opt,name=memory_id,json=memoryId,proto3" json:"memory_id,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Url      string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *MediaCreateBody) Reset() {
	*x = MediaCreateBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaCreateBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaCreateBody) ProtoMessage() {}

func (x *MediaCreateBody) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaCreateBody.ProtoReflect.Descriptor instead.
func (*MediaCreateBody) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{1}
}

func (x *MediaCreateBody) GetMemoryId() string {
	if x != nil {
		return x.MemoryId
	}
	return ""
}

func (x *MediaCreateBody) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MediaCreateBody) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type MediaCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemoryId string `protobuf:"bytes,1,opt,name=memory_id,json=memoryId,proto3" json:"memory_id,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Url      string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	UserId   string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *MediaCreate) Reset() {
	*x = MediaCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaCreate) ProtoMessage() {}

func (x *MediaCreate) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaCreate.ProtoReflect.Descriptor instead.
func (*MediaCreate) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{2}
}

func (x *MediaCreate) GetMemoryId() string {
	if x != nil {
		return x.MemoryId
	}
	return ""
}

func (x *MediaCreate) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MediaCreate) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *MediaCreate) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type MediaRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Media []*Media `protobuf:"bytes,1,rep,name=media,proto3" json:"media,omitempty"`
}

func (x *MediaRes) Reset() {
	*x = MediaRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaRes) ProtoMessage() {}

func (x *MediaRes) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaRes.ProtoReflect.Descriptor instead.
func (*MediaRes) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{3}
}

func (x *MediaRes) GetMedia() []*Media {
	if x != nil {
		return x.Media
	}
	return nil
}

var File_media_proto protoreflect.FileDescriptor

var file_media_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x54,
	0x0a, 0x0f, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x64,
	0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0x69, 0x0a, 0x0b, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x2f, 0x0a, 0x08, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x32, 0x94, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x1a, 0x0c, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00,
	0x12, 0x2a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0f, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_media_proto_rawDescOnce sync.Once
	file_media_proto_rawDescData = file_media_proto_rawDesc
)

func file_media_proto_rawDescGZIP() []byte {
	file_media_proto_rawDescOnce.Do(func() {
		file_media_proto_rawDescData = protoimpl.X.CompressGZIP(file_media_proto_rawDescData)
	})
	return file_media_proto_rawDescData
}

var file_media_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_media_proto_goTypes = []any{
	(*Media)(nil),           // 0: memory.Media
	(*MediaCreateBody)(nil), // 1: memory.MediaCreateBody
	(*MediaCreate)(nil),     // 2: memory.MediaCreate
	(*MediaRes)(nil),        // 3: memory.MediaRes
	(*GetById)(nil),         // 4: memory.GetById
	(*Void)(nil),            // 5: memory.Void
}
var file_media_proto_depIdxs = []int32{
	0, // 0: memory.MediaRes.media:type_name -> memory.Media
	2, // 1: memory.MediaService.Create:input_type -> memory.MediaCreate
	4, // 2: memory.MediaService.Get:input_type -> memory.GetById
	4, // 3: memory.MediaService.Delete:input_type -> memory.GetById
	5, // 4: memory.MediaService.Create:output_type -> memory.Void
	3, // 5: memory.MediaService.Get:output_type -> memory.MediaRes
	5, // 6: memory.MediaService.Delete:output_type -> memory.Void
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_media_proto_init() }
func file_media_proto_init() {
	if File_media_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_media_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Media); i {
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
		file_media_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*MediaCreateBody); i {
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
		file_media_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*MediaCreate); i {
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
		file_media_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*MediaRes); i {
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
			RawDescriptor: file_media_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_media_proto_goTypes,
		DependencyIndexes: file_media_proto_depIdxs,
		MessageInfos:      file_media_proto_msgTypes,
	}.Build()
	File_media_proto = out.File
	file_media_proto_rawDesc = nil
	file_media_proto_goTypes = nil
	file_media_proto_depIdxs = nil
}
