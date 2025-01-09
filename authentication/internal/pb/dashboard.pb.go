// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: dashboard.proto

package pb

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

type CreateOwnerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName   string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName    string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber string `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Nik         string `protobuf:"bytes,6,opt,name=nik,proto3" json:"nik,omitempty"`
}

func (x *CreateOwnerReq) Reset() {
	*x = CreateOwnerReq{}
	mi := &file_dashboard_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOwnerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOwnerReq) ProtoMessage() {}

func (x *CreateOwnerReq) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOwnerReq.ProtoReflect.Descriptor instead.
func (*CreateOwnerReq) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOwnerReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateOwnerReq) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateOwnerReq) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreateOwnerReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateOwnerReq) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *CreateOwnerReq) GetNik() string {
	if x != nil {
		return x.Nik
	}
	return ""
}

type IsBlockedReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IsBlockedReq) Reset() {
	*x = IsBlockedReq{}
	mi := &file_dashboard_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsBlockedReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsBlockedReq) ProtoMessage() {}

func (x *IsBlockedReq) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsBlockedReq.ProtoReflect.Descriptor instead.
func (*IsBlockedReq) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{1}
}

func (x *IsBlockedReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type IsBlockedRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsBlocked bool `protobuf:"varint,1,opt,name=is_blocked,json=isBlocked,proto3" json:"is_blocked,omitempty"`
}

func (x *IsBlockedRes) Reset() {
	*x = IsBlockedRes{}
	mi := &file_dashboard_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsBlockedRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsBlockedRes) ProtoMessage() {}

func (x *IsBlockedRes) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsBlockedRes.ProtoReflect.Descriptor instead.
func (*IsBlockedRes) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{2}
}

func (x *IsBlockedRes) GetIsBlocked() bool {
	if x != nil {
		return x.IsBlocked
	}
	return false
}

var File_dashboard_proto protoreflect.FileDescriptor

var file_dashboard_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x22, 0xa7, 0x01, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x69, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6e, 0x69, 0x6b, 0x22, 0x1e, 0x0a, 0x0c, 0x49, 0x73, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x0c, 0x49, 0x73, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x64, 0x52, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x32, 0x96, 0x01, 0x0a, 0x0c, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x1a, 0x19, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x22, 0x00, 0x12, 0x3f, 0x0a,
	0x09, 0x49, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x17, 0x2e, 0x64, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x49, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x49, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x38,
	0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x61, 0x62,
	0x72, 0x69, 0x65, 0x6c, 0x4d, 0x6f, 0x6f, 0x64, 0x79, 0x2f, 0x4d, 0x69, 0x6b, 0x72, 0x6f, 0x4e,
	0x65, 0x74, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dashboard_proto_rawDescOnce sync.Once
	file_dashboard_proto_rawDescData = file_dashboard_proto_rawDesc
)

func file_dashboard_proto_rawDescGZIP() []byte {
	file_dashboard_proto_rawDescOnce.Do(func() {
		file_dashboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_dashboard_proto_rawDescData)
	})
	return file_dashboard_proto_rawDescData
}

var file_dashboard_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_dashboard_proto_goTypes = []any{
	(*CreateOwnerReq)(nil), // 0: dashboard.CreateOwnerReq
	(*IsBlockedReq)(nil),   // 1: dashboard.IsBlockedReq
	(*IsBlockedRes)(nil),   // 2: dashboard.IsBlockedRes
}
var file_dashboard_proto_depIdxs = []int32{
	0, // 0: dashboard.OwnerService.CreateOwner:input_type -> dashboard.CreateOwnerReq
	1, // 1: dashboard.OwnerService.IsBlocked:input_type -> dashboard.IsBlockedReq
	0, // 2: dashboard.OwnerService.CreateOwner:output_type -> dashboard.CreateOwnerReq
	2, // 3: dashboard.OwnerService.IsBlocked:output_type -> dashboard.IsBlockedRes
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dashboard_proto_init() }
func file_dashboard_proto_init() {
	if File_dashboard_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dashboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dashboard_proto_goTypes,
		DependencyIndexes: file_dashboard_proto_depIdxs,
		MessageInfos:      file_dashboard_proto_msgTypes,
	}.Build()
	File_dashboard_proto = out.File
	file_dashboard_proto_rawDesc = nil
	file_dashboard_proto_goTypes = nil
	file_dashboard_proto_depIdxs = nil
}
