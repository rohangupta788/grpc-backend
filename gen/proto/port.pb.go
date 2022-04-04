// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: port.proto

package proto

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

type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Code    string `protobuf:"bytes,3,opt,name=code,proto3" json:"code"`
	City    string `protobuf:"bytes,4,opt,name=city,proto3" json:"city"`
	State   string `protobuf:"bytes,5,opt,name=state,proto3" json:"state"`
	Country string `protobuf:"bytes,6,opt,name=country,proto3" json:"country"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_port_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_port_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_port_proto_rawDescGZIP(), []int{0}
}

func (x *Port) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Port) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Port) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Port) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Port) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Port) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type CreatePortRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port *Port `protobuf:"bytes,1,opt,name=port,proto3" json:"port"`
}

func (x *CreatePortRequest) Reset() {
	*x = CreatePortRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_port_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePortRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePortRequest) ProtoMessage() {}

func (x *CreatePortRequest) ProtoReflect() protoreflect.Message {
	mi := &file_port_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePortRequest.ProtoReflect.Descriptor instead.
func (*CreatePortRequest) Descriptor() ([]byte, []int) {
	return file_port_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePortRequest) GetPort() *Port {
	if x != nil {
		return x.Port
	}
	return nil
}

type CreatePortResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result"`
}

func (x *CreatePortResponse) Reset() {
	*x = CreatePortResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_port_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePortResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePortResponse) ProtoMessage() {}

func (x *CreatePortResponse) ProtoReflect() protoreflect.Message {
	mi := &file_port_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePortResponse.ProtoReflect.Descriptor instead.
func (*CreatePortResponse) Descriptor() ([]byte, []int) {
	return file_port_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePortResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type UpdatePortRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port *Port `protobuf:"bytes,1,opt,name=port,proto3" json:"port"`
}

func (x *UpdatePortRequest) Reset() {
	*x = UpdatePortRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_port_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePortRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePortRequest) ProtoMessage() {}

func (x *UpdatePortRequest) ProtoReflect() protoreflect.Message {
	mi := &file_port_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePortRequest.ProtoReflect.Descriptor instead.
func (*UpdatePortRequest) Descriptor() ([]byte, []int) {
	return file_port_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePortRequest) GetPort() *Port {
	if x != nil {
		return x.Port
	}
	return nil
}

type UpdatePortResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result"`
}

func (x *UpdatePortResponse) Reset() {
	*x = UpdatePortResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_port_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePortResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePortResponse) ProtoMessage() {}

func (x *UpdatePortResponse) ProtoReflect() protoreflect.Message {
	mi := &file_port_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePortResponse.ProtoReflect.Descriptor instead.
func (*UpdatePortResponse) Descriptor() ([]byte, []int) {
	return file_port_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePortResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type RetrievePortRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PortId string `protobuf:"bytes,1,opt,name=port_id,json=portId,proto3" json:"port_id"`
}

func (x *RetrievePortRequest) Reset() {
	*x = RetrievePortRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_port_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrievePortRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrievePortRequest) ProtoMessage() {}

func (x *RetrievePortRequest) ProtoReflect() protoreflect.Message {
	mi := &file_port_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrievePortRequest.ProtoReflect.Descriptor instead.
func (*RetrievePortRequest) Descriptor() ([]byte, []int) {
	return file_port_proto_rawDescGZIP(), []int{5}
}

func (x *RetrievePortRequest) GetPortId() string {
	if x != nil {
		return x.PortId
	}
	return ""
}

type RetrievePortResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Code    string `protobuf:"bytes,3,opt,name=code,proto3" json:"code"`
	City    string `protobuf:"bytes,4,opt,name=city,proto3" json:"city"`
	State   string `protobuf:"bytes,5,opt,name=state,proto3" json:"state"`
	Country string `protobuf:"bytes,6,opt,name=country,proto3" json:"country"`
}

func (x *RetrievePortResponse) Reset() {
	*x = RetrievePortResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_port_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrievePortResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrievePortResponse) ProtoMessage() {}

func (x *RetrievePortResponse) ProtoReflect() protoreflect.Message {
	mi := &file_port_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrievePortResponse.ProtoReflect.Descriptor instead.
func (*RetrievePortResponse) Descriptor() ([]byte, []int) {
	return file_port_proto_rawDescGZIP(), []int{6}
}

func (x *RetrievePortResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RetrievePortResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RetrievePortResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *RetrievePortResponse) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *RetrievePortResponse) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *RetrievePortResponse) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type DeletePortResquest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PortId string `protobuf:"bytes,1,opt,name=port_id,json=portId,proto3" json:"port_id"`
}

func (x *DeletePortResquest) Reset() {
	*x = DeletePortResquest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_port_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePortResquest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePortResquest) ProtoMessage() {}

func (x *DeletePortResquest) ProtoReflect() protoreflect.Message {
	mi := &file_port_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePortResquest.ProtoReflect.Descriptor instead.
func (*DeletePortResquest) Descriptor() ([]byte, []int) {
	return file_port_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePortResquest) GetPortId() string {
	if x != nil {
		return x.PortId
	}
	return ""
}

type DeletePortResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result"`
}

func (x *DeletePortResponse) Reset() {
	*x = DeletePortResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_port_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePortResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePortResponse) ProtoMessage() {}

func (x *DeletePortResponse) ProtoReflect() protoreflect.Message {
	mi := &file_port_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePortResponse.ProtoReflect.Descriptor instead.
func (*DeletePortResponse) Descriptor() ([]byte, []int) {
	return file_port_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePortResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type ListPortRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page"`
	Count int32 `protobuf:"varint,2,opt,name=count,proto3" json:"count"`
}

func (x *ListPortRequest) Reset() {
	*x = ListPortRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_port_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPortRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPortRequest) ProtoMessage() {}

func (x *ListPortRequest) ProtoReflect() protoreflect.Message {
	mi := &file_port_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPortRequest.ProtoReflect.Descriptor instead.
func (*ListPortRequest) Descriptor() ([]byte, []int) {
	return file_port_proto_rawDescGZIP(), []int{9}
}

func (x *ListPortRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListPortRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ListPortResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port []*Port `protobuf:"bytes,1,rep,name=port,proto3" json:"port"`
}

func (x *ListPortResponse) Reset() {
	*x = ListPortResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_port_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPortResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPortResponse) ProtoMessage() {}

func (x *ListPortResponse) ProtoReflect() protoreflect.Message {
	mi := &file_port_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPortResponse.ProtoReflect.Descriptor instead.
func (*ListPortResponse) Descriptor() ([]byte, []int) {
	return file_port_proto_rawDescGZIP(), []int{10}
}

func (x *ListPortResponse) GetPort() []*Port {
	if x != nil {
		return x.Port
	}
	return nil
}

var File_port_proto protoreflect.FileDescriptor

var file_port_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61,
	0x69, 0x6e, 0x22, 0x82, 0x01, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x33, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x2c, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x33, 0x0a, 0x11, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22,
	0x2c, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2e, 0x0a,
	0x13, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x22, 0x92, 0x01,
	0x0a, 0x14, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x22, 0x2d, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x72, 0x74, 0x49,
	0x64, 0x22, 0x2c, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x3b, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x32, 0x0a, 0x10,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1e, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x32, 0xd3, 0x02, 0x0a, 0x0b, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x17,
	0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x45, 0x0a, 0x0c, 0x52, 0x65, 0x74, 0x72, 0x65, 0x69, 0x76, 0x65, 0x50, 0x6f, 0x72,
	0x74, 0x12, 0x19, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x65, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x50, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x17, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x15, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_port_proto_rawDescOnce sync.Once
	file_port_proto_rawDescData = file_port_proto_rawDesc
)

func file_port_proto_rawDescGZIP() []byte {
	file_port_proto_rawDescOnce.Do(func() {
		file_port_proto_rawDescData = protoimpl.X.CompressGZIP(file_port_proto_rawDescData)
	})
	return file_port_proto_rawDescData
}

var file_port_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_port_proto_goTypes = []interface{}{
	(*Port)(nil),                 // 0: main.Port
	(*CreatePortRequest)(nil),    // 1: main.CreatePortRequest
	(*CreatePortResponse)(nil),   // 2: main.CreatePortResponse
	(*UpdatePortRequest)(nil),    // 3: main.UpdatePortRequest
	(*UpdatePortResponse)(nil),   // 4: main.UpdatePortResponse
	(*RetrievePortRequest)(nil),  // 5: main.RetrievePortRequest
	(*RetrievePortResponse)(nil), // 6: main.RetrievePortResponse
	(*DeletePortResquest)(nil),   // 7: main.DeletePortResquest
	(*DeletePortResponse)(nil),   // 8: main.DeletePortResponse
	(*ListPortRequest)(nil),      // 9: main.ListPortRequest
	(*ListPortResponse)(nil),     // 10: main.ListPortResponse
}
var file_port_proto_depIdxs = []int32{
	0,  // 0: main.CreatePortRequest.port:type_name -> main.Port
	0,  // 1: main.UpdatePortRequest.port:type_name -> main.Port
	0,  // 2: main.ListPortResponse.port:type_name -> main.Port
	1,  // 3: main.PortService.CreatePort:input_type -> main.CreatePortRequest
	5,  // 4: main.PortService.RetreivePort:input_type -> main.RetrievePortRequest
	3,  // 5: main.PortService.UpdatePort:input_type -> main.UpdatePortRequest
	7,  // 6: main.PortService.DeletePort:input_type -> main.DeletePortResquest
	9,  // 7: main.PortService.ListPort:input_type -> main.ListPortRequest
	2,  // 8: main.PortService.CreatePort:output_type -> main.CreatePortResponse
	6,  // 9: main.PortService.RetreivePort:output_type -> main.RetrievePortResponse
	4,  // 10: main.PortService.UpdatePort:output_type -> main.UpdatePortResponse
	8,  // 11: main.PortService.DeletePort:output_type -> main.DeletePortResponse
	10, // 12: main.PortService.ListPort:output_type -> main.ListPortResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_port_proto_init() }
func file_port_proto_init() {
	if File_port_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_port_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
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
		file_port_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePortRequest); i {
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
		file_port_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePortResponse); i {
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
		file_port_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePortRequest); i {
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
		file_port_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePortResponse); i {
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
		file_port_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrievePortRequest); i {
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
		file_port_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrievePortResponse); i {
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
		file_port_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePortResquest); i {
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
		file_port_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePortResponse); i {
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
		file_port_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPortRequest); i {
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
		file_port_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPortResponse); i {
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
			RawDescriptor: file_port_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_port_proto_goTypes,
		DependencyIndexes: file_port_proto_depIdxs,
		MessageInfos:      file_port_proto_msgTypes,
	}.Build()
	File_port_proto = out.File
	file_port_proto_rawDesc = nil
	file_port_proto_goTypes = nil
	file_port_proto_depIdxs = nil
}
