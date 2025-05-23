// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: driver.proto

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

type Driver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber   string `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Sim           string `protobuf:"bytes,5,opt,name=sim,proto3" json:"sim,omitempty"`
	LicenseNumber string `protobuf:"bytes,6,opt,name=license_number,json=licenseNumber,proto3" json:"license_number,omitempty"`
	Route         string `protobuf:"bytes,7,opt,name=route,proto3" json:"route,omitempty"`
	Verified      bool   `protobuf:"varint,8,opt,name=verified,proto3" json:"verified,omitempty"`
	ImageUrl      string `protobuf:"bytes,9,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (x *Driver) Reset() {
	*x = Driver{}
	mi := &file_driver_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Driver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Driver) ProtoMessage() {}

func (x *Driver) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Driver.ProtoReflect.Descriptor instead.
func (*Driver) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{0}
}

func (x *Driver) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Driver) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Driver) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Driver) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Driver) GetSim() string {
	if x != nil {
		return x.Sim
	}
	return ""
}

func (x *Driver) GetLicenseNumber() string {
	if x != nil {
		return x.LicenseNumber
	}
	return ""
}

func (x *Driver) GetRoute() string {
	if x != nil {
		return x.Route
	}
	return ""
}

func (x *Driver) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

func (x *Driver) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

type CreateDriverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email          string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber    string `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	LicenseNumber  string `protobuf:"bytes,4,opt,name=license_number,json=licenseNumber,proto3" json:"license_number,omitempty"`
	Sim            string `protobuf:"bytes,5,opt,name=sim,proto3" json:"sim,omitempty"`
	ProfilePicture []byte `protobuf:"bytes,6,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty"`
	Id             string `protobuf:"bytes,7,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateDriverRequest) Reset() {
	*x = CreateDriverRequest{}
	mi := &file_driver_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDriverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDriverRequest) ProtoMessage() {}

func (x *CreateDriverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDriverRequest.ProtoReflect.Descriptor instead.
func (*CreateDriverRequest) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDriverRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDriverRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateDriverRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *CreateDriverRequest) GetLicenseNumber() string {
	if x != nil {
		return x.LicenseNumber
	}
	return ""
}

func (x *CreateDriverRequest) GetSim() string {
	if x != nil {
		return x.Sim
	}
	return ""
}

func (x *CreateDriverRequest) GetProfilePicture() []byte {
	if x != nil {
		return x.ProfilePicture
	}
	return nil
}

func (x *CreateDriverRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReqDrivers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Verified bool `protobuf:"varint,1,opt,name=verified,proto3" json:"verified,omitempty"`
}

func (x *ReqDrivers) Reset() {
	*x = ReqDrivers{}
	mi := &file_driver_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReqDrivers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqDrivers) ProtoMessage() {}

func (x *ReqDrivers) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqDrivers.ProtoReflect.Descriptor instead.
func (*ReqDrivers) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{2}
}

func (x *ReqDrivers) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

type ReqByID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReqByID) Reset() {
	*x = ReqByID{}
	mi := &file_driver_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReqByID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqByID) ProtoMessage() {}

func (x *ReqByID) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqByID.ProtoReflect.Descriptor instead.
func (*ReqByID) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{3}
}

func (x *ReqByID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Drivers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Drivers []*Driver `protobuf:"bytes,1,rep,name=drivers,proto3" json:"drivers,omitempty"`
}

func (x *Drivers) Reset() {
	*x = Drivers{}
	mi := &file_driver_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Drivers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Drivers) ProtoMessage() {}

func (x *Drivers) ProtoReflect() protoreflect.Message {
	mi := &file_driver_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Drivers.ProtoReflect.Descriptor instead.
func (*Drivers) Descriptor() ([]byte, []int) {
	return file_driver_proto_rawDescGZIP(), []int{4}
}

func (x *Drivers) GetDrivers() []*Driver {
	if x != nil {
		return x.Drivers
	}
	return nil
}

var File_driver_proto protoreflect.FileDescriptor

var file_driver_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x22, 0xed, 0x01, 0x0a, 0x06, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x73, 0x69, 0x6d, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x22, 0xd4, 0x01, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x25,
	0x0a, 0x0e, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x6d, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x6d, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x28, 0x0a, 0x0a, 0x52, 0x65, 0x71, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0x19, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x07, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x12, 0x2b, 0x0a, 0x07, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x52, 0x07, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x32, 0x8a, 0x02,
	0x0a, 0x0d, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x43, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12,
	0x1e, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x12, 0x15, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x52,
	0x65, 0x71, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x1a, 0x12, 0x2e, 0x64, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x22, 0x00, 0x12,
	0x3b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x12, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x52, 0x65, 0x71, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x11,
	0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x12, 0x12, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x52, 0x65,
	0x71, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x61, 0x62, 0x72, 0x69, 0x65, 0x6c,
	0x4d, 0x6f, 0x6f, 0x64, 0x79, 0x2f, 0x4d, 0x69, 0x6b, 0x72, 0x6f, 0x4e, 0x65, 0x74, 0x2f, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_driver_proto_rawDescOnce sync.Once
	file_driver_proto_rawDescData = file_driver_proto_rawDesc
)

func file_driver_proto_rawDescGZIP() []byte {
	file_driver_proto_rawDescOnce.Do(func() {
		file_driver_proto_rawDescData = protoimpl.X.CompressGZIP(file_driver_proto_rawDescData)
	})
	return file_driver_proto_rawDescData
}

var file_driver_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_driver_proto_goTypes = []any{
	(*Driver)(nil),              // 0: dashboard.Driver
	(*CreateDriverRequest)(nil), // 1: dashboard.CreateDriverRequest
	(*ReqDrivers)(nil),          // 2: dashboard.ReqDrivers
	(*ReqByID)(nil),             // 3: dashboard.ReqByID
	(*Drivers)(nil),             // 4: dashboard.Drivers
}
var file_driver_proto_depIdxs = []int32{
	0, // 0: dashboard.Drivers.drivers:type_name -> dashboard.Driver
	1, // 1: dashboard.DriverService.CreateDriver:input_type -> dashboard.CreateDriverRequest
	2, // 2: dashboard.DriverService.GetDrivers:input_type -> dashboard.ReqDrivers
	3, // 3: dashboard.DriverService.GetDriverDetails:input_type -> dashboard.ReqByID
	3, // 4: dashboard.DriverService.SetStatusVerified:input_type -> dashboard.ReqByID
	0, // 5: dashboard.DriverService.CreateDriver:output_type -> dashboard.Driver
	4, // 6: dashboard.DriverService.GetDrivers:output_type -> dashboard.Drivers
	0, // 7: dashboard.DriverService.GetDriverDetails:output_type -> dashboard.Driver
	0, // 8: dashboard.DriverService.SetStatusVerified:output_type -> dashboard.Driver
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_driver_proto_init() }
func file_driver_proto_init() {
	if File_driver_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_driver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_driver_proto_goTypes,
		DependencyIndexes: file_driver_proto_depIdxs,
		MessageInfos:      file_driver_proto_msgTypes,
	}.Build()
	File_driver_proto = out.File
	file_driver_proto_rawDesc = nil
	file_driver_proto_goTypes = nil
	file_driver_proto_depIdxs = nil
}
