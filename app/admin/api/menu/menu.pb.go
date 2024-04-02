// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: api/menu/menu.proto

package menu

import (
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
		mi := &file_api_menu_menu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Menu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Menu) ProtoMessage() {}

func (x *Menu) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Menu.ProtoReflect.Descriptor instead.
func (*Menu) Descriptor() ([]byte, []int) {
	return file_api_menu_menu_proto_rawDescGZIP(), []int{2}
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

var File_api_menu_menu_proto protoreflect.FileDescriptor

var file_api_menu_menu_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x6d, 0x65,
	0x6e, 0x75, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5c, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x64,
	0x0a, 0x0b, 0x67, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x12, 0x27, 0x0a,
	0x05, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61,
	0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x52,
	0x05, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x11, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x22, 0x5d, 0x0a, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x69, 0x63, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e,
	0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64,
	0x72, 0x65, 0x6e, 0x32, 0xbe, 0x01, 0x0a, 0x04, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x56, 0x0a, 0x07,
	0x67, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x17, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73,
	0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x67, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71,
	0x1a, 0x18, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x67,
	0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2f, 0x67, 0x65, 0x74,
	0x4d, 0x65, 0x6e, 0x75, 0x12, 0x5e, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x6d, 0x65, 0x6e,
	0x75, 0x2e, 0x67, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x61,
	0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x67, 0x65, 0x74, 0x4d, 0x65,
	0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01,
	0x2a, 0x22, 0x11, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2f, 0x67, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x91, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x74, 0x72,
	0x65, 0x75, 0x73, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x42, 0x09, 0x4d, 0x65, 0x6e, 0x75, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0xa2, 0x02, 0x03,
	0x41, 0x4d, 0x58, 0xaa, 0x02, 0x0b, 0x41, 0x74, 0x72, 0x65, 0x75, 0x73, 0x2e, 0x4d, 0x65, 0x6e,
	0x75, 0xca, 0x02, 0x0b, 0x41, 0x74, 0x72, 0x65, 0x75, 0x73, 0x5c, 0x4d, 0x65, 0x6e, 0x75, 0xe2,
	0x02, 0x17, 0x41, 0x74, 0x72, 0x65, 0x75, 0x73, 0x5c, 0x4d, 0x65, 0x6e, 0x75, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x41, 0x74, 0x72, 0x65,
	0x75, 0x73, 0x3a, 0x3a, 0x4d, 0x65, 0x6e, 0x75, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_api_menu_menu_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_menu_menu_proto_goTypes = []interface{}{
	(*GetMenuReq)(nil),  // 0: atreus.menu.getMenuReq
	(*GetMenuResp)(nil), // 1: atreus.menu.getMenuResp
	(*Menu)(nil),        // 2: atreus.menu.menu
}
var file_api_menu_menu_proto_depIdxs = []int32{
	2, // 0: atreus.menu.getMenuResp.menus:type_name -> atreus.menu.menu
	2, // 1: atreus.menu.menu.children:type_name -> atreus.menu.menu
	0, // 2: atreus.menu.Menu.getMenu:input_type -> atreus.menu.getMenuReq
	0, // 3: atreus.menu.Menu.getMenuList:input_type -> atreus.menu.getMenuReq
	1, // 4: atreus.menu.Menu.getMenu:output_type -> atreus.menu.getMenuResp
	1, // 5: atreus.menu.Menu.getMenuList:output_type -> atreus.menu.getMenuResp
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_menu_menu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
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
