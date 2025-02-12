// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: millistones.proto

package timeline

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

type Millistone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Date        string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	UserId      string `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Category    string `protobuf:"bytes,6,opt,name=category,proto3" json:"category,omitempty"`
	CreateAt    string `protobuf:"bytes,7,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	UpdateAt    string `protobuf:"bytes,8,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
	DeletedAt   int32  `protobuf:"varint,9,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Millistone) Reset() {
	*x = Millistone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_millistones_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Millistone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Millistone) ProtoMessage() {}

func (x *Millistone) ProtoReflect() protoreflect.Message {
	mi := &file_millistones_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Millistone.ProtoReflect.Descriptor instead.
func (*Millistone) Descriptor() ([]byte, []int) {
	return file_millistones_proto_rawDescGZIP(), []int{0}
}

func (x *Millistone) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Millistone) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Millistone) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Millistone) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Millistone) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Millistone) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Millistone) GetCreateAt() string {
	if x != nil {
		return x.CreateAt
	}
	return ""
}

func (x *Millistone) GetUpdateAt() string {
	if x != nil {
		return x.UpdateAt
	}
	return ""
}

func (x *Millistone) GetDeletedAt() int32 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

type Millistones struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Date        string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Category    string `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	CreateAt    string `protobuf:"bytes,6,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
}

func (x *Millistones) Reset() {
	*x = Millistones{}
	if protoimpl.UnsafeEnabled {
		mi := &file_millistones_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Millistones) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Millistones) ProtoMessage() {}

func (x *Millistones) ProtoReflect() protoreflect.Message {
	mi := &file_millistones_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Millistones.ProtoReflect.Descriptor instead.
func (*Millistones) Descriptor() ([]byte, []int) {
	return file_millistones_proto_rawDescGZIP(), []int{1}
}

func (x *Millistones) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Millistones) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Millistones) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Millistones) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Millistones) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Millistones) GetCreateAt() string {
	if x != nil {
		return x.CreateAt
	}
	return ""
}

type MillistonesCreateBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Date        string `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Category    string `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *MillistonesCreateBody) Reset() {
	*x = MillistonesCreateBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_millistones_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MillistonesCreateBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MillistonesCreateBody) ProtoMessage() {}

func (x *MillistonesCreateBody) ProtoReflect() protoreflect.Message {
	mi := &file_millistones_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MillistonesCreateBody.ProtoReflect.Descriptor instead.
func (*MillistonesCreateBody) Descriptor() ([]byte, []int) {
	return file_millistones_proto_rawDescGZIP(), []int{2}
}

func (x *MillistonesCreateBody) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MillistonesCreateBody) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *MillistonesCreateBody) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MillistonesCreateBody) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type MillistonesCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Date        string `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Category    string `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	UserId      string `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *MillistonesCreate) Reset() {
	*x = MillistonesCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_millistones_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MillistonesCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MillistonesCreate) ProtoMessage() {}

func (x *MillistonesCreate) ProtoReflect() protoreflect.Message {
	mi := &file_millistones_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MillistonesCreate.ProtoReflect.Descriptor instead.
func (*MillistonesCreate) Descriptor() ([]byte, []int) {
	return file_millistones_proto_rawDescGZIP(), []int{3}
}

func (x *MillistonesCreate) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MillistonesCreate) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *MillistonesCreate) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MillistonesCreate) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *MillistonesCreate) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type MillistonesUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Date        string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	UserId      string `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Category    string `protobuf:"bytes,6,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *MillistonesUpdate) Reset() {
	*x = MillistonesUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_millistones_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MillistonesUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MillistonesUpdate) ProtoMessage() {}

func (x *MillistonesUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_millistones_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MillistonesUpdate.ProtoReflect.Descriptor instead.
func (*MillistonesUpdate) Descriptor() ([]byte, []int) {
	return file_millistones_proto_rawDescGZIP(), []int{4}
}

func (x *MillistonesUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MillistonesUpdate) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MillistonesUpdate) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *MillistonesUpdate) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MillistonesUpdate) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *MillistonesUpdate) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type GetAllReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Category string  `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Date     string  `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Filter   *Filter `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *GetAllReq) Reset() {
	*x = GetAllReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_millistones_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllReq) ProtoMessage() {}

func (x *GetAllReq) ProtoReflect() protoreflect.Message {
	mi := &file_millistones_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllReq.ProtoReflect.Descriptor instead.
func (*GetAllReq) Descriptor() ([]byte, []int) {
	return file_millistones_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetAllReq) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *GetAllReq) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *GetAllReq) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type GetAllRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Millistones []*Millistones `protobuf:"bytes,1,rep,name=millistones,proto3" json:"millistones,omitempty"`
	TotalCount  int32          `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *GetAllRes) Reset() {
	*x = GetAllRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_millistones_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRes) ProtoMessage() {}

func (x *GetAllRes) ProtoReflect() protoreflect.Message {
	mi := &file_millistones_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRes.ProtoReflect.Descriptor instead.
func (*GetAllRes) Descriptor() ([]byte, []int) {
	return file_millistones_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllRes) GetMillistones() []*Millistones {
	if x != nil {
		return x.Millistones
	}
	return nil
}

func (x *GetAllRes) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type GetByDate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ToDate   string `protobuf:"bytes,1,opt,name=ToDate,proto3" json:"ToDate,omitempty"`
	FromDate string `protobuf:"bytes,2,opt,name=FromDate,proto3" json:"FromDate,omitempty"`
	UserId   string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetByDate) Reset() {
	*x = GetByDate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_millistones_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByDate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByDate) ProtoMessage() {}

func (x *GetByDate) ProtoReflect() protoreflect.Message {
	mi := &file_millistones_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByDate.ProtoReflect.Descriptor instead.
func (*GetByDate) Descriptor() ([]byte, []int) {
	return file_millistones_proto_rawDescGZIP(), []int{7}
}

func (x *GetByDate) GetToDate() string {
	if x != nil {
		return x.ToDate
	}
	return ""
}

func (x *GetByDate) GetFromDate() string {
	if x != nil {
		return x.FromDate
	}
	return ""
}

func (x *GetByDate) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_millistones_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_millistones_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_millistones_proto_rawDescGZIP(), []int{8}
}

type GetById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetById) Reset() {
	*x = GetById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_millistones_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetById) ProtoMessage() {}

func (x *GetById) ProtoReflect() protoreflect.Message {
	mi := &file_millistones_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetById.ProtoReflect.Descriptor instead.
func (*GetById) Descriptor() ([]byte, []int) {
	return file_millistones_proto_rawDescGZIP(), []int{9}
}

func (x *GetById) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetById) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_millistones_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_millistones_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_millistones_proto_rawDescGZIP(), []int{10}
}

func (x *Filter) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Filter) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_millistones_proto protoreflect.FileDescriptor

var file_millistones_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0xf6, 0x01,
	0x0a, 0x0a, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xa2, 0x01, 0x0a, 0x0b, 0x4d, 0x69, 0x6c, 0x6c, 0x69,
	0x73, 0x74, 0x6f, 0x6e, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x22, 0x7f, 0x0a, 0x15, 0x4d,
	0x69, 0x6c, 0x6c, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x6f, 0x64, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x94, 0x01, 0x0a,
	0x11, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x73, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0xa4, 0x01, 0x0a, 0x11, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x74, 0x6f,
	0x6e, 0x65, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x7e, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x28, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x65, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x0b, 0x6d, 0x69, 0x6c, 0x6c, 0x69,
	0x73, 0x74, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74,
	0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x74, 0x6f,
	0x6e, 0x65, 0x73, 0x52, 0x0b, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x58, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x54, 0x6f, 0x44, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x54, 0x6f, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x06, 0x0a, 0x04, 0x56,
	0x6f, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x36, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x32,
	0xe2, 0x02, 0x0a, 0x12, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x1b, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x69, 0x6c, 0x6c,
	0x69, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x0e, 0x2e,
	0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12,
	0x37, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x74, 0x69, 0x6d, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x73,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x0e, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x11, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x11,
	0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x1a, 0x15, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x69, 0x6c,
	0x6c, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x73, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x12, 0x13, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x74, 0x69, 0x6d, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x12, 0x42, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x42, 0x79, 0x44, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x6c,
	0x6c, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x73, 0x12, 0x13, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x44, 0x61, 0x74, 0x65, 0x1a, 0x13, 0x2e,
	0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_millistones_proto_rawDescOnce sync.Once
	file_millistones_proto_rawDescData = file_millistones_proto_rawDesc
)

func file_millistones_proto_rawDescGZIP() []byte {
	file_millistones_proto_rawDescOnce.Do(func() {
		file_millistones_proto_rawDescData = protoimpl.X.CompressGZIP(file_millistones_proto_rawDescData)
	})
	return file_millistones_proto_rawDescData
}

var file_millistones_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_millistones_proto_goTypes = []any{
	(*Millistone)(nil),            // 0: timeline.Millistone
	(*Millistones)(nil),           // 1: timeline.Millistones
	(*MillistonesCreateBody)(nil), // 2: timeline.MillistonesCreateBody
	(*MillistonesCreate)(nil),     // 3: timeline.MillistonesCreate
	(*MillistonesUpdate)(nil),     // 4: timeline.MillistonesUpdate
	(*GetAllReq)(nil),             // 5: timeline.GetAllReq
	(*GetAllRes)(nil),             // 6: timeline.GetAllRes
	(*GetByDate)(nil),             // 7: timeline.GetByDate
	(*Void)(nil),                  // 8: timeline.Void
	(*GetById)(nil),               // 9: timeline.GetById
	(*Filter)(nil),                // 10: timeline.Filter
}
var file_millistones_proto_depIdxs = []int32{
	10, // 0: timeline.GetAllReq.filter:type_name -> timeline.Filter
	1,  // 1: timeline.GetAllRes.millistones:type_name -> timeline.Millistones
	3,  // 2: timeline.MillistonesService.Create:input_type -> timeline.MillistonesCreate
	4,  // 3: timeline.MillistonesService.Update:input_type -> timeline.MillistonesUpdate
	9,  // 4: timeline.MillistonesService.Delete:input_type -> timeline.GetById
	9,  // 5: timeline.MillistonesService.Get:input_type -> timeline.GetById
	5,  // 6: timeline.MillistonesService.GetAll:input_type -> timeline.GetAllReq
	7,  // 7: timeline.MillistonesService.GetByDateMillistones:input_type -> timeline.GetByDate
	8,  // 8: timeline.MillistonesService.Create:output_type -> timeline.Void
	8,  // 9: timeline.MillistonesService.Update:output_type -> timeline.Void
	8,  // 10: timeline.MillistonesService.Delete:output_type -> timeline.Void
	1,  // 11: timeline.MillistonesService.Get:output_type -> timeline.Millistones
	6,  // 12: timeline.MillistonesService.GetAll:output_type -> timeline.GetAllRes
	6,  // 13: timeline.MillistonesService.GetByDateMillistones:output_type -> timeline.GetAllRes
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_millistones_proto_init() }
func file_millistones_proto_init() {
	if File_millistones_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_millistones_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Millistone); i {
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
		file_millistones_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Millistones); i {
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
		file_millistones_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*MillistonesCreateBody); i {
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
		file_millistones_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*MillistonesCreate); i {
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
		file_millistones_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*MillistonesUpdate); i {
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
		file_millistones_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllReq); i {
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
		file_millistones_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllRes); i {
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
		file_millistones_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetByDate); i {
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
		file_millistones_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Void); i {
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
		file_millistones_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*GetById); i {
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
		file_millistones_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*Filter); i {
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
			RawDescriptor: file_millistones_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_millistones_proto_goTypes,
		DependencyIndexes: file_millistones_proto_depIdxs,
		MessageInfos:      file_millistones_proto_msgTypes,
	}.Build()
	File_millistones_proto = out.File
	file_millistones_proto_rawDesc = nil
	file_millistones_proto_goTypes = nil
	file_millistones_proto_depIdxs = nil
}
