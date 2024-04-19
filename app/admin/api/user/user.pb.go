// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: api/user/user.proto

package user

import (
	core "github.com/go-atreus/atreus-server/app/admin/api/core"
	_ "github.com/go-atreus/atreus-server/app/admin/api/menu"
	role "github.com/go-atreus/atreus-server/app/admin/api/role"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secret     string `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	PlatformID int32  `protobuf:"varint,2,opt,name=platformID,proto3" json:"platformID,omitempty"`
	UserID     string `protobuf:"bytes,3,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *UserInfoReq) Reset() {
	*x = UserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoReq) ProtoMessage() {}

func (x *UserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoReq.ProtoReflect.Descriptor instead.
func (*UserInfoReq) Descriptor() ([]byte, []int) {
	return file_api_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfoReq) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *UserInfoReq) GetPlatformID() int32 {
	if x != nil {
		return x.PlatformID
	}
	return 0
}

func (x *UserInfoReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type SysUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username    string          `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	NickName    string          `protobuf:"bytes,3,opt,name=nickName,proto3" json:"nickName,omitempty"`
	Children    []*SysUser      `protobuf:"bytes,4,rep,name=children,proto3" json:"children,omitempty"`
	Password    string          `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	ParentId    int32           `protobuf:"varint,6,opt,name=parentId,proto3" json:"parentId,omitempty"`
	Avatar      string          `protobuf:"bytes,14,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Title       string          `protobuf:"bytes,7,opt,name=title,proto3" json:"title,omitempty"`
	Group       string          `protobuf:"bytes,8,opt,name=group,proto3" json:"group,omitempty"`
	Tags        string          `protobuf:"bytes,9,opt,name=tags,proto3" json:"tags,omitempty"`
	NotifyCount string          `protobuf:"bytes,10,opt,name=notifyCount,proto3" json:"notifyCount,omitempty"`
	UnreadCount string          `protobuf:"bytes,11,opt,name=unreadCount,proto3" json:"unreadCount,omitempty"`
	Country     string          `protobuf:"bytes,12,opt,name=country,proto3" json:"country,omitempty"`
	Access      string          `protobuf:"bytes,13,opt,name=access,proto3" json:"access,omitempty"`
	Permissions []string        `protobuf:"bytes,15,rep,name=permissions,proto3" json:"permissions,omitempty"`
	Roles       []string        `protobuf:"bytes,16,rep,name=roles,proto3" json:"roles,omitempty"`
	Role        []*role.SysRole `protobuf:"bytes,17,rep,name=role,proto3" json:"role,omitempty"`
}

func (x *SysUser) Reset() {
	*x = SysUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysUser) ProtoMessage() {}

func (x *SysUser) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysUser.ProtoReflect.Descriptor instead.
func (*SysUser) Descriptor() ([]byte, []int) {
	return file_api_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *SysUser) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SysUser) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SysUser) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *SysUser) GetChildren() []*SysUser {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *SysUser) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SysUser) GetParentId() int32 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *SysUser) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *SysUser) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SysUser) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *SysUser) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *SysUser) GetNotifyCount() string {
	if x != nil {
		return x.NotifyCount
	}
	return ""
}

func (x *SysUser) GetUnreadCount() string {
	if x != nil {
		return x.UnreadCount
	}
	return ""
}

func (x *SysUser) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *SysUser) GetAccess() string {
	if x != nil {
		return x.Access
	}
	return ""
}

func (x *SysUser) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *SysUser) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *SysUser) GetRole() []*role.SysRole {
	if x != nil {
		return x.Role
	}
	return nil
}

type ForceLogoutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlatformID int32  `protobuf:"varint,1,opt,name=platformID,proto3" json:"platformID,omitempty"`
	UserID     string `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *ForceLogoutReq) Reset() {
	*x = ForceLogoutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForceLogoutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForceLogoutReq) ProtoMessage() {}

func (x *ForceLogoutReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForceLogoutReq.ProtoReflect.Descriptor instead.
func (*ForceLogoutReq) Descriptor() ([]byte, []int) {
	return file_api_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *ForceLogoutReq) GetPlatformID() int32 {
	if x != nil {
		return x.PlatformID
	}
	return 0
}

func (x *ForceLogoutReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type ForceLogoutResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ForceLogoutResp) Reset() {
	*x = ForceLogoutResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForceLogoutResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForceLogoutResp) ProtoMessage() {}

func (x *ForceLogoutResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForceLogoutResp.ProtoReflect.Descriptor instead.
func (*ForceLogoutResp) Descriptor() ([]byte, []int) {
	return file_api_user_user_proto_rawDescGZIP(), []int{3}
}

type ParseTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ParseTokenReq) Reset() {
	*x = ParseTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseTokenReq) ProtoMessage() {}

func (x *ParseTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseTokenReq.ProtoReflect.Descriptor instead.
func (*ParseTokenReq) Descriptor() ([]byte, []int) {
	return file_api_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *ParseTokenReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ParseTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID            string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Platform          string `protobuf:"bytes,2,opt,name=platform,proto3" json:"platform,omitempty"`
	ExpireTimeSeconds int64  `protobuf:"varint,4,opt,name=expireTimeSeconds,proto3" json:"expireTimeSeconds,omitempty"`
}

func (x *ParseTokenResp) Reset() {
	*x = ParseTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseTokenResp) ProtoMessage() {}

func (x *ParseTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseTokenResp.ProtoReflect.Descriptor instead.
func (*ParseTokenResp) Descriptor() ([]byte, []int) {
	return file_api_user_user_proto_rawDescGZIP(), []int{5}
}

func (x *ParseTokenResp) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *ParseTokenResp) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *ParseTokenResp) GetExpireTimeSeconds() int64 {
	if x != nil {
		return x.ExpireTimeSeconds
	}
	return 0
}

type GetUserTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlatformID int32  `protobuf:"varint,1,opt,name=platformID,proto3" json:"platformID,omitempty"`
	UserID     string `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *GetUserTokenReq) Reset() {
	*x = GetUserTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTokenReq) ProtoMessage() {}

func (x *GetUserTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTokenReq.ProtoReflect.Descriptor instead.
func (*GetUserTokenReq) Descriptor() ([]byte, []int) {
	return file_api_user_user_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserTokenReq) GetPlatformID() int32 {
	if x != nil {
		return x.PlatformID
	}
	return 0
}

func (x *GetUserTokenReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type GetUserTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token             string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ExpireTimeSeconds int64  `protobuf:"varint,2,opt,name=expireTimeSeconds,proto3" json:"expireTimeSeconds,omitempty"`
}

func (x *GetUserTokenResp) Reset() {
	*x = GetUserTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTokenResp) ProtoMessage() {}

func (x *GetUserTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTokenResp.ProtoReflect.Descriptor instead.
func (*GetUserTokenResp) Descriptor() ([]byte, []int) {
	return file_api_user_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserTokenResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetUserTokenResp) GetExpireTimeSeconds() int64 {
	if x != nil {
		return x.ExpireTimeSeconds
	}
	return 0
}

type ListUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Note repeated field and plural name 'results'
	Results  []*SysUser     `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	PageInfo *core.PageInfo `protobuf:"bytes,2,opt,name=page_info,json=pageInfo,proto3" json:"page_info,omitempty"`
}

func (x *ListUser) Reset() {
	*x = ListUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUser) ProtoMessage() {}

func (x *ListUser) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUser.ProtoReflect.Descriptor instead.
func (*ListUser) Descriptor() ([]byte, []int) {
	return file_api_user_user_proto_rawDescGZIP(), []int{8}
}

func (x *ListUser) GetResults() []*SysUser {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ListUser) GetPageInfo() *core.PageInfo {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

var File_api_user_user_proto protoreflect.FileDescriptor

var file_api_user_user_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2f, 0x6d, 0x65, 0x6e, 0x75,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x6c, 0x65,
	0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5d, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22,
	0xcc, 0x04, 0x0a, 0x07, 0x53, 0x79, 0x73, 0x55, 0x73, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x13, 0xba, 0xb9, 0x19, 0x0f, 0x0a, 0x0d, 0x12,
	0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x28, 0x01, 0x48, 0x01, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x24, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x30, 0x01, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x53, 0x79, 0x73, 0x55, 0x73, 0x65, 0x72, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02,
	0x10, 0x01, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x28, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x09, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x10, 0x01, 0x52,
	0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x05,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x09, 0x42, 0x06, 0xba, 0xb9, 0x19,
	0x02, 0x10, 0x01, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75,
	0x73, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x15,
	0xba, 0xb9, 0x19, 0x11, 0x32, 0x0f, 0x0a, 0x0d, 0x73, 0x79, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x3a, 0x11, 0xba, 0xb9, 0x19,
	0x0d, 0x08, 0x01, 0x1a, 0x09, 0x73, 0x79, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x48,
	0x0a, 0x0e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x11, 0x0a, 0x0f, 0x66, 0x6f, 0x72, 0x63,
	0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x25, 0x0a, 0x0d, 0x70,
	0x61, 0x72, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x72, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x2c, 0x0a, 0x11, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x11, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x49, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x22, 0x56, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2c, 0x0a, 0x11, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x6e, 0x0a, 0x08, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x79, 0x73, 0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75,
	0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0x96, 0x02, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x54, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x18, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x61, 0x74,
	0x72, 0x65, 0x75, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x79, 0x73, 0x55, 0x73, 0x65,
	0x72, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x54, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x79, 0x73, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x61, 0x74, 0x72, 0x65,
	0x75, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x79, 0x73, 0x55, 0x73, 0x65, 0x72, 0x1a,
	0x14, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x79,
	0x73, 0x55, 0x73, 0x65, 0x72, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a,
	0x22, 0x0c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x5a,
	0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x22, 0x1c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0x06, 0xba, 0xb9, 0x19, 0x02,
	0x08, 0x01, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2f, 0x61, 0x74, 0x72, 0x65, 0x75,
	0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3b, 0x75, 0x73, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_user_user_proto_rawDescOnce sync.Once
	file_api_user_user_proto_rawDescData = file_api_user_user_proto_rawDesc
)

func file_api_user_user_proto_rawDescGZIP() []byte {
	file_api_user_user_proto_rawDescOnce.Do(func() {
		file_api_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_user_user_proto_rawDescData)
	})
	return file_api_user_user_proto_rawDescData
}

var file_api_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_user_user_proto_goTypes = []interface{}{
	(*UserInfoReq)(nil),      // 0: atreus.user.userInfoReq
	(*SysUser)(nil),          // 1: atreus.user.SysUser
	(*ForceLogoutReq)(nil),   // 2: atreus.user.forceLogoutReq
	(*ForceLogoutResp)(nil),  // 3: atreus.user.forceLogoutResp
	(*ParseTokenReq)(nil),    // 4: atreus.user.parseTokenReq
	(*ParseTokenResp)(nil),   // 5: atreus.user.parseTokenResp
	(*GetUserTokenReq)(nil),  // 6: atreus.user.getUserTokenReq
	(*GetUserTokenResp)(nil), // 7: atreus.user.getUserTokenResp
	(*ListUser)(nil),         // 8: atreus.user.ListUser
	(*role.SysRole)(nil),     // 9: atreus.role.SysRole
	(*core.PageInfo)(nil),    // 10: atreus.core.PageInfo
	(*emptypb.Empty)(nil),    // 11: google.protobuf.Empty
}
var file_api_user_user_proto_depIdxs = []int32{
	1,  // 0: atreus.user.SysUser.children:type_name -> atreus.user.SysUser
	9,  // 1: atreus.user.SysUser.role:type_name -> atreus.role.SysRole
	1,  // 2: atreus.user.ListUser.results:type_name -> atreus.user.SysUser
	10, // 3: atreus.user.ListUser.page_info:type_name -> atreus.core.PageInfo
	0,  // 4: atreus.user.User.getUserInfo:input_type -> atreus.user.userInfoReq
	1,  // 5: atreus.user.User.CreateSysUser:input_type -> atreus.user.SysUser
	11, // 6: atreus.user.User.ListSysUser:input_type -> google.protobuf.Empty
	1,  // 7: atreus.user.User.getUserInfo:output_type -> atreus.user.SysUser
	1,  // 8: atreus.user.User.CreateSysUser:output_type -> atreus.user.SysUser
	8,  // 9: atreus.user.User.ListSysUser:output_type -> atreus.user.ListUser
	7,  // [7:10] is the sub-list for method output_type
	4,  // [4:7] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_user_user_proto_init() }
func file_api_user_user_proto_init() {
	if File_api_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoReq); i {
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
		file_api_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysUser); i {
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
		file_api_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForceLogoutReq); i {
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
		file_api_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForceLogoutResp); i {
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
		file_api_user_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParseTokenReq); i {
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
		file_api_user_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParseTokenResp); i {
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
		file_api_user_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserTokenReq); i {
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
		file_api_user_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserTokenResp); i {
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
		file_api_user_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUser); i {
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
			RawDescriptor: file_api_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_user_user_proto_goTypes,
		DependencyIndexes: file_api_user_user_proto_depIdxs,
		MessageInfos:      file_api_user_user_proto_msgTypes,
	}.Build()
	File_api_user_user_proto = out.File
	file_api_user_user_proto_rawDesc = nil
	file_api_user_user_proto_goTypes = nil
	file_api_user_user_proto_depIdxs = nil
}
