// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: api/dict/dict.proto

package dict

import (
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SysDict struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// 编号
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// *
	// 标识
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	// *
	// 名称
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// *
	// Hash值
	HashCode string `protobuf:"bytes,4,opt,name=hashCode,proto3" json:"hashCode,omitempty"`
	// *
	// 备注
	Remarks string `protobuf:"bytes,5,opt,name=remarks,proto3" json:"remarks,omitempty"`
	// *
	// 状态,1：启用 0：禁用
	Status int32 `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	// *
	// 数据类型,1:int32 2:String 3:Boolean
	ValueType int32 `protobuf:"varint,7,opt,name=valueType,proto3" json:"valueType,omitempty"`
	// *
	// 逻辑删除标识，已删除:0，未删除：删除时间戳
	Deleted string `protobuf:"bytes,8,opt,name=deleted,proto3" json:"deleted,omitempty"`
	// *
	// 创建时间
	CreatedAt string `protobuf:"bytes,9,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// *
	// 更新时间
	UpdatedAt string `protobuf:"bytes,10,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *SysDict) Reset() {
	*x = SysDict{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dict_dict_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysDict) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysDict) ProtoMessage() {}

func (x *SysDict) ProtoReflect() protoreflect.Message {
	mi := &file_api_dict_dict_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysDict.ProtoReflect.Descriptor instead.
func (*SysDict) Descriptor() ([]byte, []int) {
	return file_api_dict_dict_proto_rawDescGZIP(), []int{0}
}

func (x *SysDict) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SysDict) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SysDict) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SysDict) GetHashCode() string {
	if x != nil {
		return x.HashCode
	}
	return ""
}

func (x *SysDict) GetRemarks() string {
	if x != nil {
		return x.Remarks
	}
	return ""
}

func (x *SysDict) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SysDict) GetValueType() int32 {
	if x != nil {
		return x.ValueType
	}
	return 0
}

func (x *SysDict) GetDeleted() string {
	if x != nil {
		return x.Deleted
	}
	return ""
}

func (x *SysDict) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *SysDict) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type ListSysDictResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*SysDict `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ListSysDictResp) Reset() {
	*x = ListSysDictResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dict_dict_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSysDictResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSysDictResp) ProtoMessage() {}

func (x *ListSysDictResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_dict_dict_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSysDictResp.ProtoReflect.Descriptor instead.
func (*ListSysDictResp) Descriptor() ([]byte, []int) {
	return file_api_dict_dict_proto_rawDescGZIP(), []int{1}
}

func (x *ListSysDictResp) GetResults() []*SysDict {
	if x != nil {
		return x.Results
	}
	return nil
}

type ValidHashResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []string `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ValidHashResp) Reset() {
	*x = ValidHashResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dict_dict_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidHashResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidHashResp) ProtoMessage() {}

func (x *ValidHashResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_dict_dict_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidHashResp.ProtoReflect.Descriptor instead.
func (*ValidHashResp) Descriptor() ([]byte, []int) {
	return file_api_dict_dict_proto_rawDescGZIP(), []int{2}
}

func (x *ValidHashResp) GetResults() []string {
	if x != nil {
		return x.Results
	}
	return nil
}

type DictDataReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DictCodes string `protobuf:"bytes,1,opt,name=dictCodes,proto3" json:"dictCodes,omitempty"`
}

func (x *DictDataReq) Reset() {
	*x = DictDataReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dict_dict_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DictDataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DictDataReq) ProtoMessage() {}

func (x *DictDataReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_dict_dict_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DictDataReq.ProtoReflect.Descriptor instead.
func (*DictDataReq) Descriptor() ([]byte, []int) {
	return file_api_dict_dict_proto_rawDescGZIP(), []int{3}
}

func (x *DictDataReq) GetDictCodes() string {
	if x != nil {
		return x.DictCodes
	}
	return ""
}

type DictData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DictCode string `protobuf:"bytes,1,opt,name=dictCode,proto3" json:"dictCode,omitempty"`
	HashCode string `protobuf:"bytes,2,opt,name=hashCode,proto3" json:"hashCode,omitempty"`
	// 1: 数字; 2: 字符串; 3: 布尔
	ValueType int32              `protobuf:"varint,3,opt,name=valueType,proto3" json:"valueType,omitempty"`
	DictItems []*SysDictDataItem `protobuf:"bytes,4,rep,name=dictItems,proto3" json:"dictItems,omitempty"`
	Loading   bool               `protobuf:"varint,5,opt,name=loading,proto3" json:"loading,omitempty"`
}

func (x *DictData) Reset() {
	*x = DictData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dict_dict_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DictData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DictData) ProtoMessage() {}

func (x *DictData) ProtoReflect() protoreflect.Message {
	mi := &file_api_dict_dict_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DictData.ProtoReflect.Descriptor instead.
func (*DictData) Descriptor() ([]byte, []int) {
	return file_api_dict_dict_proto_rawDescGZIP(), []int{4}
}

func (x *DictData) GetDictCode() string {
	if x != nil {
		return x.DictCode
	}
	return ""
}

func (x *DictData) GetHashCode() string {
	if x != nil {
		return x.HashCode
	}
	return ""
}

func (x *DictData) GetValueType() int32 {
	if x != nil {
		return x.ValueType
	}
	return 0
}

func (x *DictData) GetDictItems() []*SysDictDataItem {
	if x != nil {
		return x.DictItems
	}
	return nil
}

func (x *DictData) GetLoading() bool {
	if x != nil {
		return x.Loading
	}
	return false
}

type SysDictItemAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TagColor  string `protobuf:"bytes,1,opt,name=tagColor,proto3" json:"tagColor,omitempty"`
	TextColor string `protobuf:"bytes,2,opt,name=textColor,proto3" json:"textColor,omitempty"`
}

func (x *SysDictItemAttributes) Reset() {
	*x = SysDictItemAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dict_dict_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysDictItemAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysDictItemAttributes) ProtoMessage() {}

func (x *SysDictItemAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_api_dict_dict_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysDictItemAttributes.ProtoReflect.Descriptor instead.
func (*SysDictItemAttributes) Descriptor() ([]byte, []int) {
	return file_api_dict_dict_proto_rawDescGZIP(), []int{5}
}

func (x *SysDictItemAttributes) GetTagColor() string {
	if x != nil {
		return x.TagColor
	}
	return ""
}

func (x *SysDictItemAttributes) GetTextColor() string {
	if x != nil {
		return x.TextColor
	}
	return ""
}

type SysDictDataItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 文本值
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 数据值
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	// 附加属性值
	Attributes *SysDictItemAttributes `protobuf:"bytes,4,opt,name=attributes,proto3" json:"attributes,omitempty"`
	// 真实数据
	RealVal []byte `protobuf:"bytes,5,opt,name=realVal,proto3" json:"realVal,omitempty"`
}

func (x *SysDictDataItem) Reset() {
	*x = SysDictDataItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dict_dict_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysDictDataItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysDictDataItem) ProtoMessage() {}

func (x *SysDictDataItem) ProtoReflect() protoreflect.Message {
	mi := &file_api_dict_dict_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysDictDataItem.ProtoReflect.Descriptor instead.
func (*SysDictDataItem) Descriptor() ([]byte, []int) {
	return file_api_dict_dict_proto_rawDescGZIP(), []int{6}
}

func (x *SysDictDataItem) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SysDictDataItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SysDictDataItem) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *SysDictDataItem) GetAttributes() *SysDictItemAttributes {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *SysDictDataItem) GetRealVal() []byte {
	if x != nil {
		return x.RealVal
	}
	return nil
}

type DictDataResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*DictData `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *DictDataResp) Reset() {
	*x = DictDataResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dict_dict_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DictDataResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DictDataResp) ProtoMessage() {}

func (x *DictDataResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_dict_dict_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DictDataResp.ProtoReflect.Descriptor instead.
func (*DictDataResp) Descriptor() ([]byte, []int) {
	return file_api_dict_dict_proto_rawDescGZIP(), []int{7}
}

func (x *DictDataResp) GetResults() []*DictData {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_api_dict_dict_proto protoreflect.FileDescriptor

var file_api_dict_dict_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x64, 0x69,
	0x63, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa4, 0x02, 0x0a, 0x07, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x12, 0x1a,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0xba, 0xb9, 0x19, 0x06,
	0x0a, 0x04, 0x28, 0x01, 0x48, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x3a, 0x11, 0xba, 0xb9, 0x19, 0x0d, 0x08, 0x01, 0x1a, 0x09,
	0x73, 0x79, 0x73, 0x5f, 0x64, 0x69, 0x63, 0x74, 0x73, 0x22, 0x41, 0x0a, 0x0f, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x44,
	0x69, 0x63, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x29, 0x0a, 0x0d,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x2b, 0x0a, 0x0b, 0x44, 0x69, 0x63, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x63, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x63, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x73, 0x22, 0xb6, 0x01, 0x0a, 0x08, 0x44, 0x69, 0x63, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x68, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x68, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x64, 0x69, 0x63, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x74, 0x72,
	0x65, 0x75, 0x73, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x64, 0x69, 0x63, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6c, 0x6f, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x51, 0x0a,
	0x15, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x67, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72,
	0x22, 0xa9, 0x01, 0x0a, 0x0f, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x42,
	0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x64, 0x69, 0x63, 0x74,
	0x2e, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x65, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x22, 0x3f, 0x0a, 0x0c,
	0x44, 0x69, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2f, 0x0a, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x44, 0x69, 0x63, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x32, 0xd4, 0x04,
	0x0a, 0x04, 0x44, 0x69, 0x63, 0x74, 0x12, 0x5a, 0x0a, 0x08, 0x44, 0x69, 0x63, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x18, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x64, 0x69, 0x63, 0x74,
	0x2e, 0x44, 0x69, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x61,
	0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12,
	0x11, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x69, 0x0a, 0x0d, 0x44, 0x69, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x61, 0x74,
	0x72, 0x65, 0x75, 0x73, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x48,
	0x61, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a,
	0x01, 0x2a, 0x22, 0x19, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x64, 0x69, 0x63, 0x74,
	0x2f, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x2d, 0x68, 0x61, 0x73, 0x68, 0x12, 0x54, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73,
	0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x1a, 0x14, 0x2e,
	0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x44,
	0x69, 0x63, 0x74, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13,
	0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x2f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x46, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e,
	0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x44,
	0x69, 0x63, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x64, 0x69, 0x63,
	0x74, 0x2e, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0a, 0x3a, 0x01, 0x2a, 0x1a, 0x05, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x12, 0x43, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x64,
	0x69, 0x63, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x74,
	0x72, 0x65, 0x75, 0x73, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63,
	0x74, 0x22, 0x0d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x2a, 0x05, 0x2f, 0x64, 0x69, 0x63, 0x74,
	0x12, 0x40, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73,
	0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x1a, 0x14, 0x2e,
	0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x44,
	0x69, 0x63, 0x74, 0x22, 0x0d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f, 0x64, 0x69,
	0x63, 0x74, 0x12, 0x58, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x61, 0x74, 0x72,
	0x65, 0x75, 0x73, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74,
	0x1a, 0x1c, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0x06, 0xba, 0xb9,
	0x19, 0x02, 0x08, 0x01, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2f, 0x61, 0x74, 0x72,
	0x65, 0x75, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x3b, 0x64, 0x69,
	0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_dict_dict_proto_rawDescOnce sync.Once
	file_api_dict_dict_proto_rawDescData = file_api_dict_dict_proto_rawDesc
)

func file_api_dict_dict_proto_rawDescGZIP() []byte {
	file_api_dict_dict_proto_rawDescOnce.Do(func() {
		file_api_dict_dict_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_dict_dict_proto_rawDescData)
	})
	return file_api_dict_dict_proto_rawDescData
}

var file_api_dict_dict_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_dict_dict_proto_goTypes = []interface{}{
	(*SysDict)(nil),               // 0: atreus.dict.SysDict
	(*ListSysDictResp)(nil),       // 1: atreus.dict.ListSysDictResp
	(*ValidHashResp)(nil),         // 2: atreus.dict.ValidHashResp
	(*DictDataReq)(nil),           // 3: atreus.dict.DictDataReq
	(*DictData)(nil),              // 4: atreus.dict.DictData
	(*SysDictItemAttributes)(nil), // 5: atreus.dict.SysDictItemAttributes
	(*SysDictDataItem)(nil),       // 6: atreus.dict.SysDictDataItem
	(*DictDataResp)(nil),          // 7: atreus.dict.DictDataResp
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_api_dict_dict_proto_depIdxs = []int32{
	0,  // 0: atreus.dict.ListSysDictResp.results:type_name -> atreus.dict.SysDict
	6,  // 1: atreus.dict.DictData.dictItems:type_name -> atreus.dict.SysDictDataItem
	5,  // 2: atreus.dict.SysDictDataItem.attributes:type_name -> atreus.dict.SysDictItemAttributes
	4,  // 3: atreus.dict.DictDataResp.results:type_name -> atreus.dict.DictData
	3,  // 4: atreus.dict.Dict.DictData:input_type -> atreus.dict.DictDataReq
	8,  // 5: atreus.dict.Dict.DictValidHash:input_type -> google.protobuf.Empty
	0,  // 6: atreus.dict.Dict.Create:input_type -> atreus.dict.SysDict
	0,  // 7: atreus.dict.Dict.Update:input_type -> atreus.dict.SysDict
	0,  // 8: atreus.dict.Dict.Delete:input_type -> atreus.dict.SysDict
	0,  // 9: atreus.dict.Dict.Get:input_type -> atreus.dict.SysDict
	0,  // 10: atreus.dict.Dict.List:input_type -> atreus.dict.SysDict
	7,  // 11: atreus.dict.Dict.DictData:output_type -> atreus.dict.DictDataResp
	2,  // 12: atreus.dict.Dict.DictValidHash:output_type -> atreus.dict.ValidHashResp
	0,  // 13: atreus.dict.Dict.Create:output_type -> atreus.dict.SysDict
	0,  // 14: atreus.dict.Dict.Update:output_type -> atreus.dict.SysDict
	0,  // 15: atreus.dict.Dict.Delete:output_type -> atreus.dict.SysDict
	0,  // 16: atreus.dict.Dict.Get:output_type -> atreus.dict.SysDict
	1,  // 17: atreus.dict.Dict.List:output_type -> atreus.dict.ListSysDictResp
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_dict_dict_proto_init() }
func file_api_dict_dict_proto_init() {
	if File_api_dict_dict_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_dict_dict_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysDict); i {
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
		file_api_dict_dict_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSysDictResp); i {
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
		file_api_dict_dict_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidHashResp); i {
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
		file_api_dict_dict_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DictDataReq); i {
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
		file_api_dict_dict_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DictData); i {
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
		file_api_dict_dict_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysDictItemAttributes); i {
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
		file_api_dict_dict_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysDictDataItem); i {
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
		file_api_dict_dict_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DictDataResp); i {
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
			RawDescriptor: file_api_dict_dict_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_dict_dict_proto_goTypes,
		DependencyIndexes: file_api_dict_dict_proto_depIdxs,
		MessageInfos:      file_api_dict_dict_proto_msgTypes,
	}.Build()
	File_api_dict_dict_proto = out.File
	file_api_dict_dict_proto_rawDesc = nil
	file_api_dict_dict_proto_goTypes = nil
	file_api_dict_dict_proto_depIdxs = nil
}
