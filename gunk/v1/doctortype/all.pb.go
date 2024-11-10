// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: main.go/gunk/v1/doctortype/all.proto

package doctortypepb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// doctor type struct
type DoctorType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	DoctorType string `protobuf:"bytes,2,opt,name=DoctorType,proto3" json:"DoctorType,omitempty"`
}

func (x *DoctorType) Reset() {
	*x = DoctorType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_go_gunk_v1_doctortype_all_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoctorType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorType) ProtoMessage() {}

func (x *DoctorType) ProtoReflect() protoreflect.Message {
	mi := &file_main_go_gunk_v1_doctortype_all_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorType.ProtoReflect.Descriptor instead.
func (*DoctorType) Descriptor() ([]byte, []int) {
	return file_main_go_gunk_v1_doctortype_all_proto_rawDescGZIP(), []int{0}
}

func (x *DoctorType) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *DoctorType) GetDoctorType() string {
	if x != nil {
		return x.DoctorType
	}
	return ""
}

// Doctor type register request
type RegisterDoctorTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DoctorType string `protobuf:"bytes,1,opt,name=DoctorType,proto3" json:"DoctorType,omitempty"`
}

func (x *RegisterDoctorTypeRequest) Reset() {
	*x = RegisterDoctorTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_go_gunk_v1_doctortype_all_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterDoctorTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterDoctorTypeRequest) ProtoMessage() {}

func (x *RegisterDoctorTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_go_gunk_v1_doctortype_all_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterDoctorTypeRequest.ProtoReflect.Descriptor instead.
func (*RegisterDoctorTypeRequest) Descriptor() ([]byte, []int) {
	return file_main_go_gunk_v1_doctortype_all_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterDoctorTypeRequest) GetDoctorType() string {
	if x != nil {
		return x.DoctorType
	}
	return ""
}

// Doctor type register response
type RegisterDoctorTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *DoctorType `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *RegisterDoctorTypeResponse) Reset() {
	*x = RegisterDoctorTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_go_gunk_v1_doctortype_all_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterDoctorTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterDoctorTypeResponse) ProtoMessage() {}

func (x *RegisterDoctorTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_main_go_gunk_v1_doctortype_all_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterDoctorTypeResponse.ProtoReflect.Descriptor instead.
func (*RegisterDoctorTypeResponse) Descriptor() ([]byte, []int) {
	return file_main_go_gunk_v1_doctortype_all_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterDoctorTypeResponse) GetUser() *DoctorType {
	if x != nil {
		return x.User
	}
	return nil
}

// edit doctor type request
type EditDoctorTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EditDoctorTypeRequest) Reset() {
	*x = EditDoctorTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_go_gunk_v1_doctortype_all_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditDoctorTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditDoctorTypeRequest) ProtoMessage() {}

func (x *EditDoctorTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_go_gunk_v1_doctortype_all_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditDoctorTypeRequest.ProtoReflect.Descriptor instead.
func (*EditDoctorTypeRequest) Descriptor() ([]byte, []int) {
	return file_main_go_gunk_v1_doctortype_all_proto_rawDescGZIP(), []int{3}
}

func (x *EditDoctorTypeRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// edit doctor type response
type EditDoctorTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *DoctorType `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *EditDoctorTypeResponse) Reset() {
	*x = EditDoctorTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_go_gunk_v1_doctortype_all_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditDoctorTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditDoctorTypeResponse) ProtoMessage() {}

func (x *EditDoctorTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_main_go_gunk_v1_doctortype_all_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditDoctorTypeResponse.ProtoReflect.Descriptor instead.
func (*EditDoctorTypeResponse) Descriptor() ([]byte, []int) {
	return file_main_go_gunk_v1_doctortype_all_proto_rawDescGZIP(), []int{4}
}

func (x *EditDoctorTypeResponse) GetUser() *DoctorType {
	if x != nil {
		return x.User
	}
	return nil
}

// update doctor type request
type UpdateDoctorTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	DcotorType string `protobuf:"bytes,2,opt,name=DcotorType,proto3" json:"DcotorType,omitempty"`
}

func (x *UpdateDoctorTypeRequest) Reset() {
	*x = UpdateDoctorTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_go_gunk_v1_doctortype_all_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDoctorTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDoctorTypeRequest) ProtoMessage() {}

func (x *UpdateDoctorTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_go_gunk_v1_doctortype_all_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDoctorTypeRequest.ProtoReflect.Descriptor instead.
func (*UpdateDoctorTypeRequest) Descriptor() ([]byte, []int) {
	return file_main_go_gunk_v1_doctortype_all_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateDoctorTypeRequest) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateDoctorTypeRequest) GetDcotorType() string {
	if x != nil {
		return x.DcotorType
	}
	return ""
}

// update doctor type response
type UpdateDoctorTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *DoctorType `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *UpdateDoctorTypeResponse) Reset() {
	*x = UpdateDoctorTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_go_gunk_v1_doctortype_all_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDoctorTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDoctorTypeResponse) ProtoMessage() {}

func (x *UpdateDoctorTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_main_go_gunk_v1_doctortype_all_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDoctorTypeResponse.ProtoReflect.Descriptor instead.
func (*UpdateDoctorTypeResponse) Descriptor() ([]byte, []int) {
	return file_main_go_gunk_v1_doctortype_all_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateDoctorTypeResponse) GetUser() *DoctorType {
	if x != nil {
		return x.User
	}
	return nil
}

// delete doctor type request
type DeleteDoctorTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteDoctorTypeRequest) Reset() {
	*x = DeleteDoctorTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_go_gunk_v1_doctortype_all_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDoctorTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDoctorTypeRequest) ProtoMessage() {}

func (x *DeleteDoctorTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_go_gunk_v1_doctortype_all_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDoctorTypeRequest.ProtoReflect.Descriptor instead.
func (*DeleteDoctorTypeRequest) Descriptor() ([]byte, []int) {
	return file_main_go_gunk_v1_doctortype_all_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteDoctorTypeRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// delete doctor type response
type DeleteDoctorTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteDoctorTypeResponse) Reset() {
	*x = DeleteDoctorTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_go_gunk_v1_doctortype_all_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDoctorTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDoctorTypeResponse) ProtoMessage() {}

func (x *DeleteDoctorTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_main_go_gunk_v1_doctortype_all_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDoctorTypeResponse.ProtoReflect.Descriptor instead.
func (*DeleteDoctorTypeResponse) Descriptor() ([]byte, []int) {
	return file_main_go_gunk_v1_doctortype_all_proto_rawDescGZIP(), []int{8}
}

// list user request
type DoctorTypeListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SearchTerm string `protobuf:"bytes,1,opt,name=SearchTerm,proto3" json:"SearchTerm,omitempty"`
}

func (x *DoctorTypeListRequest) Reset() {
	*x = DoctorTypeListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_go_gunk_v1_doctortype_all_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoctorTypeListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorTypeListRequest) ProtoMessage() {}

func (x *DoctorTypeListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_go_gunk_v1_doctortype_all_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorTypeListRequest.ProtoReflect.Descriptor instead.
func (*DoctorTypeListRequest) Descriptor() ([]byte, []int) {
	return file_main_go_gunk_v1_doctortype_all_proto_rawDescGZIP(), []int{9}
}

func (x *DoctorTypeListRequest) GetSearchTerm() string {
	if x != nil {
		return x.SearchTerm
	}
	return ""
}

// list user response
type DoctorTypeListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DoctorType []*DoctorType `protobuf:"bytes,1,rep,name=DoctorType,proto3" json:"DoctorType,omitempty"`
}

func (x *DoctorTypeListResponse) Reset() {
	*x = DoctorTypeListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_go_gunk_v1_doctortype_all_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoctorTypeListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorTypeListResponse) ProtoMessage() {}

func (x *DoctorTypeListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_main_go_gunk_v1_doctortype_all_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoctorTypeListResponse.ProtoReflect.Descriptor instead.
func (*DoctorTypeListResponse) Descriptor() ([]byte, []int) {
	return file_main_go_gunk_v1_doctortype_all_proto_rawDescGZIP(), []int{10}
}

func (x *DoctorTypeListResponse) GetDoctorType() []*DoctorType {
	if x != nil {
		return x.DoctorType
	}
	return nil
}

var File_main_go_gunk_v1_doctortype_all_proto protoreflect.FileDescriptor

var file_main_go_gunk_v1_doctortype_all_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x6f, 0x2f, 0x67, 0x75, 0x6e, 0x6b, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x61, 0x6c, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x74, 0x79,
	0x70, 0x65, 0x70, 0x62, 0x22, 0x4c, 0x0a, 0x0a, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a,
	0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x6f,
	0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00,
	0x18, 0x00, 0x22, 0x43, 0x0a, 0x19, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x6f,
	0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a,
	0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x58, 0x0a, 0x1a, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x74, 0x79, 0x70, 0x65,
	0x70, 0x62, 0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0a, 0x08,
	0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18,
	0x00, 0x22, 0x37, 0x0a, 0x15, 0x45, 0x64, 0x69, 0x74, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00,
	0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x54, 0x0a, 0x16, 0x45, 0x64,
	0x69, 0x74, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x74, 0x79, 0x70, 0x65, 0x70,
	0x62, 0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0a, 0x08, 0x00,
	0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00,
	0x22, 0x59, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30,
	0x00, 0x50, 0x00, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x63, 0x6f, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30,
	0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x56, 0x0a, 0x18, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x74, 0x79,
	0x70, 0x65, 0x70, 0x62, 0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10,
	0x00, 0x18, 0x00, 0x22, 0x39, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00,
	0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x22,
	0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00,
	0x18, 0x00, 0x22, 0x3f, 0x0a, 0x15, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10,
	0x00, 0x18, 0x00, 0x22, 0x5a, 0x0a, 0x16, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a,
	0x0a, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x74, 0x79, 0x70, 0x65, 0x70, 0x62,
	0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0a, 0x08, 0x00, 0x18,
	0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x32,
	0xbd, 0x04, 0x0a, 0x11, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x73, 0x0a, 0x12, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x2e, 0x64, 0x6f,
	0x63, 0x74, 0x6f, 0x72, 0x74, 0x79, 0x70, 0x65, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x74, 0x79, 0x70,
	0x65, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x6f, 0x63, 0x74,
	0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06,
	0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x28, 0x00, 0x30, 0x00, 0x12, 0x67, 0x0a, 0x0e, 0x45, 0x64,
	0x69, 0x74, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x2e, 0x64,
	0x6f, 0x63, 0x74, 0x6f, 0x72, 0x74, 0x79, 0x70, 0x65, 0x70, 0x62, 0x2e, 0x45, 0x64, 0x69, 0x74,
	0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x74, 0x79, 0x70, 0x65, 0x70, 0x62,
	0x2e, 0x45, 0x64, 0x69, 0x74, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x28,
	0x00, 0x30, 0x00, 0x12, 0x6d, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72,
	0x74, 0x79, 0x70, 0x65, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x74, 0x79, 0x70, 0x65, 0x70, 0x62, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x28, 0x00,
	0x30, 0x00, 0x12, 0x6d, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x74,
	0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x74,
	0x79, 0x70, 0x65, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x74,
	0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x74, 0x79, 0x70, 0x65, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x28, 0x00, 0x30,
	0x00, 0x12, 0x67, 0x0a, 0x0e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x74, 0x79, 0x70, 0x65,
	0x70, 0x62, 0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x64, 0x6f, 0x63, 0x74, 0x6f,
	0x72, 0x74, 0x79, 0x70, 0x65, 0x70, 0x62, 0x2e, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06,
	0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x28, 0x00, 0x30, 0x00, 0x1a, 0x03, 0x88, 0x02, 0x00, 0x42,
	0x42, 0x48, 0x01, 0x50, 0x00, 0x5a, 0x27, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x6f, 0x2f, 0x67,
	0x75, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x74, 0x79, 0x70,
	0x65, 0x3b, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x74, 0x79, 0x70, 0x65, 0x70, 0x62, 0x80, 0x01,
	0x00, 0x88, 0x01, 0x00, 0x90, 0x01, 0x00, 0xb8, 0x01, 0x00, 0xd8, 0x01, 0x00, 0xf8, 0x01, 0x01,
	0xd0, 0x02, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_main_go_gunk_v1_doctortype_all_proto_rawDescOnce sync.Once
	file_main_go_gunk_v1_doctortype_all_proto_rawDescData = file_main_go_gunk_v1_doctortype_all_proto_rawDesc
)

func file_main_go_gunk_v1_doctortype_all_proto_rawDescGZIP() []byte {
	file_main_go_gunk_v1_doctortype_all_proto_rawDescOnce.Do(func() {
		file_main_go_gunk_v1_doctortype_all_proto_rawDescData = protoimpl.X.CompressGZIP(file_main_go_gunk_v1_doctortype_all_proto_rawDescData)
	})
	return file_main_go_gunk_v1_doctortype_all_proto_rawDescData
}

var (
	file_main_go_gunk_v1_doctortype_all_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
	file_main_go_gunk_v1_doctortype_all_proto_goTypes  = []interface{}{
		(*DoctorType)(nil),                 // 0: doctortypepb.DoctorType
		(*RegisterDoctorTypeRequest)(nil),  // 1: doctortypepb.RegisterDoctorTypeRequest
		(*RegisterDoctorTypeResponse)(nil), // 2: doctortypepb.RegisterDoctorTypeResponse
		(*EditDoctorTypeRequest)(nil),      // 3: doctortypepb.EditDoctorTypeRequest
		(*EditDoctorTypeResponse)(nil),     // 4: doctortypepb.EditDoctorTypeResponse
		(*UpdateDoctorTypeRequest)(nil),    // 5: doctortypepb.UpdateDoctorTypeRequest
		(*UpdateDoctorTypeResponse)(nil),   // 6: doctortypepb.UpdateDoctorTypeResponse
		(*DeleteDoctorTypeRequest)(nil),    // 7: doctortypepb.DeleteDoctorTypeRequest
		(*DeleteDoctorTypeResponse)(nil),   // 8: doctortypepb.DeleteDoctorTypeResponse
		(*DoctorTypeListRequest)(nil),      // 9: doctortypepb.DoctorTypeListRequest
		(*DoctorTypeListResponse)(nil),     // 10: doctortypepb.DoctorTypeListResponse
	}
)

var file_main_go_gunk_v1_doctortype_all_proto_depIdxs = []int32{
	0,  // 0: doctortypepb.RegisterDoctorTypeResponse.User:type_name -> doctortypepb.DoctorType
	0,  // 1: doctortypepb.EditDoctorTypeResponse.User:type_name -> doctortypepb.DoctorType
	0,  // 2: doctortypepb.UpdateDoctorTypeResponse.User:type_name -> doctortypepb.DoctorType
	0,  // 3: doctortypepb.DoctorTypeListResponse.DoctorType:type_name -> doctortypepb.DoctorType
	1,  // 4: doctortypepb.DoctorTypeService.RegisterDoctorType:input_type -> doctortypepb.RegisterDoctorTypeRequest
	3,  // 5: doctortypepb.DoctorTypeService.EditDoctorType:input_type -> doctortypepb.EditDoctorTypeRequest
	5,  // 6: doctortypepb.DoctorTypeService.UpdateDoctorType:input_type -> doctortypepb.UpdateDoctorTypeRequest
	7,  // 7: doctortypepb.DoctorTypeService.DeleteDoctorType:input_type -> doctortypepb.DeleteDoctorTypeRequest
	9,  // 8: doctortypepb.DoctorTypeService.DoctorTypeList:input_type -> doctortypepb.DoctorTypeListRequest
	2,  // 9: doctortypepb.DoctorTypeService.RegisterDoctorType:output_type -> doctortypepb.RegisterDoctorTypeResponse
	4,  // 10: doctortypepb.DoctorTypeService.EditDoctorType:output_type -> doctortypepb.EditDoctorTypeResponse
	6,  // 11: doctortypepb.DoctorTypeService.UpdateDoctorType:output_type -> doctortypepb.UpdateDoctorTypeResponse
	8,  // 12: doctortypepb.DoctorTypeService.DeleteDoctorType:output_type -> doctortypepb.DeleteDoctorTypeResponse
	10, // 13: doctortypepb.DoctorTypeService.DoctorTypeList:output_type -> doctortypepb.DoctorTypeListResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_main_go_gunk_v1_doctortype_all_proto_init() }
func file_main_go_gunk_v1_doctortype_all_proto_init() {
	if File_main_go_gunk_v1_doctortype_all_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_main_go_gunk_v1_doctortype_all_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoctorType); i {
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
		file_main_go_gunk_v1_doctortype_all_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterDoctorTypeRequest); i {
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
		file_main_go_gunk_v1_doctortype_all_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterDoctorTypeResponse); i {
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
		file_main_go_gunk_v1_doctortype_all_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditDoctorTypeRequest); i {
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
		file_main_go_gunk_v1_doctortype_all_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditDoctorTypeResponse); i {
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
		file_main_go_gunk_v1_doctortype_all_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDoctorTypeRequest); i {
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
		file_main_go_gunk_v1_doctortype_all_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDoctorTypeResponse); i {
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
		file_main_go_gunk_v1_doctortype_all_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDoctorTypeRequest); i {
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
		file_main_go_gunk_v1_doctortype_all_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDoctorTypeResponse); i {
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
		file_main_go_gunk_v1_doctortype_all_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoctorTypeListRequest); i {
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
		file_main_go_gunk_v1_doctortype_all_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoctorTypeListResponse); i {
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
			RawDescriptor: file_main_go_gunk_v1_doctortype_all_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_main_go_gunk_v1_doctortype_all_proto_goTypes,
		DependencyIndexes: file_main_go_gunk_v1_doctortype_all_proto_depIdxs,
		MessageInfos:      file_main_go_gunk_v1_doctortype_all_proto_msgTypes,
	}.Build()
	File_main_go_gunk_v1_doctortype_all_proto = out.File
	file_main_go_gunk_v1_doctortype_all_proto_rawDesc = nil
	file_main_go_gunk_v1_doctortype_all_proto_goTypes = nil
	file_main_go_gunk_v1_doctortype_all_proto_depIdxs = nil
}
