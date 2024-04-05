// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: api/menu/menu.proto

package menu

import (
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetMenuReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secret     string `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	PlatformID int32  `protobuf:"varint,2,opt,name=platformID,proto3" json:"platformID,omitempty"`
	UserID     string `protobuf:"bytes,3,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *GetMenuReq) Reset() {
	*x = GetMenuReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_menu_menu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMenuReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMenuReq) ProtoMessage() {}

func (x *GetMenuReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_menu_menu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMenuReq.ProtoReflect.Descriptor instead.
func (*GetMenuReq) Descriptor() ([]byte, []int) {
	return file_api_menu_menu_proto_rawDescGZIP(), []int{0}
}

func (x *GetMenuReq) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *GetMenuReq) GetPlatformID() int32 {
	if x != nil {
		return x.PlatformID
	}
	return 0
}

func (x *GetMenuReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type GetMenuResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Menus             []*Menu `protobuf:"bytes,1,rep,name=menus,proto3" json:"menus,omitempty"`
	ExpireTimeSeconds int64   `protobuf:"varint,3,opt,name=expireTimeSeconds,proto3" json:"expireTimeSeconds,omitempty"`
}

func (x *GetMenuResp) Reset() {
	*x = GetMenuResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_menu_menu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMenuResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMenuResp) ProtoMessage() {}

func (x *GetMenuResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_menu_menu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMenuResp.ProtoReflect.Descriptor instead.
func (*GetMenuResp) Descriptor() ([]byte, []int) {
	return file_api_menu_menu_proto_rawDescGZIP(), []int{1}
}

func (x *GetMenuResp) GetMenus() []*Menu {
	if x != nil {
		return x.Menus
	}
	return nil
}

func (x *GetMenuResp) GetExpireTimeSeconds() int64 {
	if x != nil {
		return x.ExpireTimeSeconds
	}
	return 0
}

type SysMenu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 菜单ID
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 父级ID
	ParentId int32 `protobuf:"varint,2,opt,name=parentId,proto3" json:"parentId,omitempty"`
	// 菜单名称
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// 菜单名称
	I18NTitle string `protobuf:"bytes,4,opt,name=i18nTitle,proto3" json:"i18nTitle,omitempty"`
	// 菜单图标
	Icon string `protobuf:"bytes,5,opt,name=icon,proto3" json:"icon,omitempty"`
	// 授权标识
	Permission string `protobuf:"bytes,6,opt,name=permission,proto3" json:"permission,omitempty"`
	// 路由地址
	Path string `protobuf:"bytes,7,opt,name=path,proto3" json:"path,omitempty"`
	// 打开方式 (1组件 2内链 3外链)
	TargetType int32 `protobuf:"varint,8,opt,name=targetType,proto3" json:"targetType,omitempty"`
	// 定位标识 (打开方式为组件时其值为组件相对路径，其他为URL地址)
	Uri string `protobuf:"bytes,9,opt,name=uri,proto3" json:"uri,omitempty"`
	// 显示排序
	Sort int32 `protobuf:"varint,10,opt,name=sort,proto3" json:"sort,omitempty"`
	// 组件缓存：0-开启，1-关闭
	KeepAlive int32 `protobuf:"varint,11,opt,name=keepAlive,proto3" json:"keepAlive,omitempty"`
	// 隐藏菜单: 0-否，1-是
	Hidden int32 `protobuf:"varint,12,opt,name=hidden,proto3" json:"hidden,omitempty"`
	// 菜单类型 （0目录，1菜单，2按钮）
	Type int32 `protobuf:"varint,13,opt,name=type,proto3" json:"type,omitempty"`
	// 备注信息
	Remarks string `protobuf:"bytes,14,opt,name=remarks,proto3" json:"remarks,omitempty"`
	// 创建时间
	CreateTime string `protobuf:"bytes,15,opt,name=createTime,proto3" json:"createTime,omitempty"`
	// 更新时间
	UpdateTime string `protobuf:"bytes,16,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	// 子级
	Children []*SysMenu `protobuf:"bytes,17,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *SysMenu) Reset() {
	*x = SysMenu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_menu_menu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysMenu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysMenu) ProtoMessage() {}

func (x *SysMenu) ProtoReflect() protoreflect.Message {
	mi := &file_api_menu_menu_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysMenu.ProtoReflect.Descriptor instead.
func (*SysMenu) Descriptor() ([]byte, []int) {
	return file_api_menu_menu_proto_rawDescGZIP(), []int{2}
}

func (x *SysMenu) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SysMenu) GetParentId() int32 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *SysMenu) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SysMenu) GetI18NTitle() string {
	if x != nil {
		return x.I18NTitle
	}
	return ""
}

func (x *SysMenu) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *SysMenu) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

func (x *SysMenu) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *SysMenu) GetTargetType() int32 {
	if x != nil {
		return x.TargetType
	}
	return 0
}

func (x *SysMenu) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *SysMenu) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *SysMenu) GetKeepAlive() int32 {
	if x != nil {
		return x.KeepAlive
	}
	return 0
}

func (x *SysMenu) GetHidden() int32 {
	if x != nil {
		return x.Hidden
	}
	return 0
}

func (x *SysMenu) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *SysMenu) GetRemarks() string {
	if x != nil {
		return x.Remarks
	}
	return ""
}

func (x *SysMenu) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *SysMenu) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

func (x *SysMenu) GetChildren() []*SysMenu {
	if x != nil {
		return x.Children
	}
	return nil
}

type Menu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Icon     string  `protobuf:"bytes,2,opt,name=icon,proto3" json:"icon,omitempty"`
	Children []*Menu `protobuf:"bytes,3,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *Menu) Reset() {
	*x = Menu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_menu_menu_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Menu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Menu) ProtoMessage() {}

func (x *Menu) ProtoReflect() protoreflect.Message {
	mi := &file_api_menu_menu_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Menu.ProtoReflect.Descriptor instead.
func (*Menu) Descriptor() ([]byte, []int) {
	return file_api_menu_menu_proto_rawDescGZIP(), []int{3}
}

func (x *Menu) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Menu) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *Menu) GetChildren() []*Menu {
	if x != nil {
		return x.Children
	}
	return nil
}

type ListSysMenuResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*SysMenu `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ListSysMenuResp) Reset() {
	*x = ListSysMenuResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_menu_menu_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSysMenuResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSysMenuResp) ProtoMessage() {}

func (x *ListSysMenuResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_menu_menu_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSysMenuResp.ProtoReflect.Descriptor instead.
func (*ListSysMenuResp) Descriptor() ([]byte, []int) {
	return file_api_menu_menu_proto_rawDescGZIP(), []int{4}
}

func (x *ListSysMenuResp) GetResults() []*SysMenu {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_api_menu_menu_proto protoreflect.FileDescriptor

var file_api_menu_menu_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x6d, 0x65,
	0x6e, 0x75, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x12, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52,
	0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x22, 0x64, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x27, 0x0a, 0x05, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x6d,
	0x65, 0x6e, 0x75, 0x52, 0x05, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0xf5, 0x03, 0x0a, 0x07, 0x53, 0x79, 0x73,
	0x4d, 0x65, 0x6e, 0x75, 0x12, 0x23, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x13, 0xba, 0xb9, 0x19, 0x0f, 0x0a, 0x0d, 0x12, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65,
	0x72, 0x28, 0x01, 0x48, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69,
	0x31, 0x38, 0x6e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x69, 0x31, 0x38, 0x6e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x69, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x41,
	0x6c, 0x69, 0x76, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6b, 0x65, 0x65, 0x70,
	0x41, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x63,
	0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x53, 0x79, 0x73, 0x4d,
	0x65, 0x6e, 0x75, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x3a, 0x11, 0xba,
	0xb9, 0x19, 0x0d, 0x08, 0x01, 0x1a, 0x09, 0x73, 0x79, 0x73, 0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x73,
	0x22, 0x5d, 0x0a, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x63, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e,
	0x12, 0x2d, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x6d, 0x65, 0x6e, 0x75,
	0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22,
	0x41, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x6d, 0x65, 0x6e,
	0x75, 0x2e, 0x53, 0x79, 0x73, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x32, 0xae, 0x02, 0x0a, 0x04, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x5d, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x17, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e,
	0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x67, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x1a,
	0x18, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x67, 0x65,
	0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x6d, 0x65,
	0x6e, 0x75, 0x2f, 0x67, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x62, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x79, 0x73, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x17, 0x2e, 0x61, 0x74, 0x72, 0x65,
	0x75, 0x73, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x67, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52,
	0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x6d, 0x65, 0x6e, 0x75,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x5b,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x4d, 0x65, 0x6e, 0x75, 0x12,
	0x14, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x53, 0x79,
	0x73, 0x4d, 0x65, 0x6e, 0x75, 0x1a, 0x14, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x6d,
	0x65, 0x6e, 0x75, 0x2e, 0x53, 0x79, 0x73, 0x4d, 0x65, 0x6e, 0x75, 0x22, 0x1e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f,
	0x6d, 0x65, 0x6e, 0x75, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x06, 0xba, 0xb9, 0x19,
	0x02, 0x08, 0x01, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_menu_menu_proto_rawDescOnce sync.Once
	file_api_menu_menu_proto_rawDescData = file_api_menu_menu_proto_rawDesc
)

func file_api_menu_menu_proto_rawDescGZIP() []byte {
	file_api_menu_menu_proto_rawDescOnce.Do(func() {
		file_api_menu_menu_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_menu_menu_proto_rawDescData)
	})
	return file_api_menu_menu_proto_rawDescData
}

var file_api_menu_menu_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_menu_menu_proto_goTypes = []interface{}{
	(*GetMenuReq)(nil),      // 0: atreus.menu.getMenuReq
	(*GetMenuResp)(nil),     // 1: atreus.menu.getMenuResp
	(*SysMenu)(nil),         // 2: atreus.menu.SysMenu
	(*Menu)(nil),            // 3: atreus.menu.menu
	(*ListSysMenuResp)(nil), // 4: atreus.menu.ListSysMenuResp
}
var file_api_menu_menu_proto_depIdxs = []int32{
	3, // 0: atreus.menu.getMenuResp.menus:type_name -> atreus.menu.menu
	2, // 1: atreus.menu.SysMenu.children:type_name -> atreus.menu.SysMenu
	3, // 2: atreus.menu.menu.children:type_name -> atreus.menu.menu
	2, // 3: atreus.menu.ListSysMenuResp.results:type_name -> atreus.menu.SysMenu
	0, // 4: atreus.menu.Menu.GetMenu:input_type -> atreus.menu.getMenuReq
	0, // 5: atreus.menu.Menu.ListSysMenu:input_type -> atreus.menu.getMenuReq
	2, // 6: atreus.menu.Menu.CreateSysMenu:input_type -> atreus.menu.SysMenu
	1, // 7: atreus.menu.Menu.GetMenu:output_type -> atreus.menu.getMenuResp
	4, // 8: atreus.menu.Menu.ListSysMenu:output_type -> atreus.menu.ListSysMenuResp
	2, // 9: atreus.menu.Menu.CreateSysMenu:output_type -> atreus.menu.SysMenu
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_menu_menu_proto_init() }
func file_api_menu_menu_proto_init() {
	if File_api_menu_menu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_menu_menu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMenuReq); i {
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
		file_api_menu_menu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMenuResp); i {
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
		file_api_menu_menu_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysMenu); i {
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
		file_api_menu_menu_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Menu); i {
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
		file_api_menu_menu_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSysMenuResp); i {
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
			RawDescriptor: file_api_menu_menu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_menu_menu_proto_goTypes,
		DependencyIndexes: file_api_menu_menu_proto_depIdxs,
		MessageInfos:      file_api_menu_menu_proto_msgTypes,
	}.Build()
	File_api_menu_menu_proto = out.File
	file_api_menu_menu_proto_rawDesc = nil
	file_api_menu_menu_proto_goTypes = nil
	file_api_menu_menu_proto_depIdxs = nil
}
