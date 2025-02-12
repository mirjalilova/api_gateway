// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: share.proto

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

type ShareCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemoryId   string   `protobuf:"bytes,1,opt,name=memory_id,json=memoryId,proto3" json:"memory_id,omitempty"`
	SharedWith []string `protobuf:"bytes,2,rep,name=shared_with,json=sharedWith,proto3" json:"shared_with,omitempty"`
	UserId     string   `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ShareCreate) Reset() {
	*x = ShareCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_share_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareCreate) ProtoMessage() {}

func (x *ShareCreate) ProtoReflect() protoreflect.Message {
	mi := &file_share_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareCreate.ProtoReflect.Descriptor instead.
func (*ShareCreate) Descriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{0}
}

func (x *ShareCreate) GetMemoryId() string {
	if x != nil {
		return x.MemoryId
	}
	return ""
}

func (x *ShareCreate) GetSharedWith() []string {
	if x != nil {
		return x.SharedWith
	}
	return nil
}

func (x *ShareCreate) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ShareCreateBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemoryId   string   `protobuf:"bytes,1,opt,name=memory_id,json=memoryId,proto3" json:"memory_id,omitempty"`
	SharedWith []string `protobuf:"bytes,2,rep,name=shared_with,json=sharedWith,proto3" json:"shared_with,omitempty"`
}

func (x *ShareCreateBody) Reset() {
	*x = ShareCreateBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_share_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareCreateBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareCreateBody) ProtoMessage() {}

func (x *ShareCreateBody) ProtoReflect() protoreflect.Message {
	mi := &file_share_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareCreateBody.ProtoReflect.Descriptor instead.
func (*ShareCreateBody) Descriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{1}
}

func (x *ShareCreateBody) GetMemoryId() string {
	if x != nil {
		return x.MemoryId
	}
	return ""
}

func (x *ShareCreateBody) GetSharedWith() []string {
	if x != nil {
		return x.SharedWith
	}
	return nil
}

type ShareRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemonyTitle string   `protobuf:"bytes,1,opt,name=memony_title,json=memonyTitle,proto3" json:"memony_title,omitempty"`
	SharedWith  []string `protobuf:"bytes,2,rep,name=shared_with,json=sharedWith,proto3" json:"shared_with,omitempty"`
}

func (x *ShareRes) Reset() {
	*x = ShareRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_share_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareRes) ProtoMessage() {}

func (x *ShareRes) ProtoReflect() protoreflect.Message {
	mi := &file_share_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareRes.ProtoReflect.Descriptor instead.
func (*ShareRes) Descriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{2}
}

func (x *ShareRes) GetMemonyTitle() string {
	if x != nil {
		return x.MemonyTitle
	}
	return ""
}

func (x *ShareRes) GetSharedWith() []string {
	if x != nil {
		return x.SharedWith
	}
	return nil
}

type ShareDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MemoryId     string   `protobuf:"bytes,2,opt,name=memory_id,json=memoryId,proto3" json:"memory_id,omitempty"`
	UserId       string   `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UnsharedWith []string `protobuf:"bytes,4,rep,name=unshared_with,json=unsharedWith,proto3" json:"unshared_with,omitempty"`
}

func (x *ShareDelete) Reset() {
	*x = ShareDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_share_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareDelete) ProtoMessage() {}

func (x *ShareDelete) ProtoReflect() protoreflect.Message {
	mi := &file_share_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareDelete.ProtoReflect.Descriptor instead.
func (*ShareDelete) Descriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{3}
}

func (x *ShareDelete) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ShareDelete) GetMemoryId() string {
	if x != nil {
		return x.MemoryId
	}
	return ""
}

func (x *ShareDelete) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ShareDelete) GetUnsharedWith() []string {
	if x != nil {
		return x.UnsharedWith
	}
	return nil
}

type ShareDeleteBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemoryId     string   `protobuf:"bytes,1,opt,name=memory_id,json=memoryId,proto3" json:"memory_id,omitempty"`
	UserId       string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UnsharedWith []string `protobuf:"bytes,3,rep,name=unshared_with,json=unsharedWith,proto3" json:"unshared_with,omitempty"`
}

func (x *ShareDeleteBody) Reset() {
	*x = ShareDeleteBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_share_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareDeleteBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareDeleteBody) ProtoMessage() {}

func (x *ShareDeleteBody) ProtoReflect() protoreflect.Message {
	mi := &file_share_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareDeleteBody.ProtoReflect.Descriptor instead.
func (*ShareDeleteBody) Descriptor() ([]byte, []int) {
	return file_share_proto_rawDescGZIP(), []int{4}
}

func (x *ShareDeleteBody) GetMemoryId() string {
	if x != nil {
		return x.MemoryId
	}
	return ""
}

func (x *ShareDeleteBody) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ShareDeleteBody) GetUnsharedWith() []string {
	if x != nil {
		return x.UnsharedWith
	}
	return nil
}

var File_share_proto protoreflect.FileDescriptor

var file_share_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x0b, 0x53, 0x68, 0x61, 0x72, 0x65, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x57, 0x69, 0x74, 0x68,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x0f, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x57, 0x69, 0x74, 0x68, 0x22, 0x4e, 0x0a, 0x08, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x6e, 0x79,
	0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65,
	0x6d, 0x6f, 0x6e, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x57, 0x69, 0x74, 0x68, 0x22, 0x78, 0x0a, 0x0b, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x23, 0x0a, 0x0d, 0x75, 0x6e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x77, 0x69, 0x74, 0x68,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x6e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x57, 0x69, 0x74, 0x68, 0x22, 0x6c, 0x0a, 0x0f, 0x53, 0x68, 0x61, 0x72, 0x65, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x75, 0x6e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x6e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x57, 0x69,
	0x74, 0x68, 0x32, 0x9c, 0x01, 0x0a, 0x0c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x53, 0x68, 0x61, 0x72, 0x65, 0x12, 0x13, 0x2e, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x1a, 0x0c, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22,
	0x00, 0x12, 0x32, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x12, 0x13, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x1a, 0x0c, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x56,
	0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0f, 0x2e, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x10, 0x2e,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_share_proto_rawDescOnce sync.Once
	file_share_proto_rawDescData = file_share_proto_rawDesc
)

func file_share_proto_rawDescGZIP() []byte {
	file_share_proto_rawDescOnce.Do(func() {
		file_share_proto_rawDescData = protoimpl.X.CompressGZIP(file_share_proto_rawDescData)
	})
	return file_share_proto_rawDescData
}

var file_share_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_share_proto_goTypes = []any{
	(*ShareCreate)(nil),     // 0: memory.ShareCreate
	(*ShareCreateBody)(nil), // 1: memory.ShareCreateBody
	(*ShareRes)(nil),        // 2: memory.ShareRes
	(*ShareDelete)(nil),     // 3: memory.ShareDelete
	(*ShareDeleteBody)(nil), // 4: memory.ShareDeleteBody
	(*GetById)(nil),         // 5: memory.GetById
	(*Void)(nil),            // 6: memory.Void
}
var file_share_proto_depIdxs = []int32{
	0, // 0: memory.ShareService.Share:input_type -> memory.ShareCreate
	3, // 1: memory.ShareService.Updateshare:input_type -> memory.ShareDelete
	5, // 2: memory.ShareService.Get:input_type -> memory.GetById
	6, // 3: memory.ShareService.Share:output_type -> memory.Void
	6, // 4: memory.ShareService.Updateshare:output_type -> memory.Void
	2, // 5: memory.ShareService.Get:output_type -> memory.ShareRes
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_share_proto_init() }
func file_share_proto_init() {
	if File_share_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_share_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ShareCreate); i {
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
		file_share_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ShareCreateBody); i {
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
		file_share_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ShareRes); i {
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
		file_share_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ShareDelete); i {
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
		file_share_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ShareDeleteBody); i {
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
			RawDescriptor: file_share_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_share_proto_goTypes,
		DependencyIndexes: file_share_proto_depIdxs,
		MessageInfos:      file_share_proto_msgTypes,
	}.Build()
	File_share_proto = out.File
	file_share_proto_rawDesc = nil
	file_share_proto_goTypes = nil
	file_share_proto_depIdxs = nil
}
