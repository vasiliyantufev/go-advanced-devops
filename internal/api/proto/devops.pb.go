// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: internal/api/proto/devops.proto

package devops

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

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_devops_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_devops_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_devops_proto_rawDescGZIP(), []int{0}
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_devops_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_devops_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_devops_proto_rawDescGZIP(), []int{1}
}

type GetMetricRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *GetMetricRequest) Reset() {
	*x = GetMetricRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_devops_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMetricRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetricRequest) ProtoMessage() {}

func (x *GetMetricRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_devops_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetricRequest.ProtoReflect.Descriptor instead.
func (*GetMetricRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_devops_proto_rawDescGZIP(), []int{2}
}

func (x *GetMetricRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetMetricRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type GetMetricResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMetricResponse) Reset() {
	*x = GetMetricResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_devops_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMetricResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetricResponse) ProtoMessage() {}

func (x *GetMetricResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_devops_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetricResponse.ProtoReflect.Descriptor instead.
func (*GetMetricResponse) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_devops_proto_rawDescGZIP(), []int{3}
}

type CreateMetricRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Mtype string `protobuf:"bytes,2,opt,name=mtype,proto3" json:"mtype,omitempty"`
}

func (x *CreateMetricRequest) Reset() {
	*x = CreateMetricRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_devops_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMetricRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMetricRequest) ProtoMessage() {}

func (x *CreateMetricRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_devops_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMetricRequest.ProtoReflect.Descriptor instead.
func (*CreateMetricRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_devops_proto_rawDescGZIP(), []int{4}
}

func (x *CreateMetricRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateMetricRequest) GetMtype() string {
	if x != nil {
		return x.Mtype
	}
	return ""
}

type CreateMetricResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metric []*Metric `protobuf:"bytes,1,rep,name=metric,proto3" json:"metric,omitempty"`
}

func (x *CreateMetricResponse) Reset() {
	*x = CreateMetricResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_devops_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMetricResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMetricResponse) ProtoMessage() {}

func (x *CreateMetricResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_devops_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMetricResponse.ProtoReflect.Descriptor instead.
func (*CreateMetricResponse) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_devops_proto_rawDescGZIP(), []int{5}
}

func (x *CreateMetricResponse) GetMetric() []*Metric {
	if x != nil {
		return x.Metric
	}
	return nil
}

type CreateMetricsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metric []*MetricRequest `protobuf:"bytes,1,rep,name=metric,proto3" json:"metric,omitempty"`
}

func (x *CreateMetricsRequest) Reset() {
	*x = CreateMetricsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_devops_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMetricsRequest) ProtoMessage() {}

func (x *CreateMetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_devops_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMetricsRequest.ProtoReflect.Descriptor instead.
func (*CreateMetricsRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_devops_proto_rawDescGZIP(), []int{6}
}

func (x *CreateMetricsRequest) GetMetric() []*MetricRequest {
	if x != nil {
		return x.Metric
	}
	return nil
}

type CreateMetricsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metric []*Metric `protobuf:"bytes,1,rep,name=metric,proto3" json:"metric,omitempty"`
}

func (x *CreateMetricsResponse) Reset() {
	*x = CreateMetricsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_devops_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMetricsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMetricsResponse) ProtoMessage() {}

func (x *CreateMetricsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_devops_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMetricsResponse.ProtoReflect.Descriptor instead.
func (*CreateMetricsResponse) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_devops_proto_rawDescGZIP(), []int{7}
}

func (x *CreateMetricsResponse) GetMetric() []*Metric {
	if x != nil {
		return x.Metric
	}
	return nil
}

type Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Mtype string `protobuf:"bytes,2,opt,name=mtype,proto3" json:"mtype,omitempty"`
	Delta string `protobuf:"bytes,3,opt,name=delta,proto3" json:"delta,omitempty"`
	Value string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Hash  string `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *Metric) Reset() {
	*x = Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_devops_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric) ProtoMessage() {}

func (x *Metric) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_devops_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metric.ProtoReflect.Descriptor instead.
func (*Metric) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_devops_proto_rawDescGZIP(), []int{8}
}

func (x *Metric) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Metric) GetMtype() string {
	if x != nil {
		return x.Mtype
	}
	return ""
}

func (x *Metric) GetDelta() string {
	if x != nil {
		return x.Delta
	}
	return ""
}

func (x *Metric) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Metric) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type MetricRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Mtype string `protobuf:"bytes,2,opt,name=mtype,proto3" json:"mtype,omitempty"`
}

func (x *MetricRequest) Reset() {
	*x = MetricRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_proto_devops_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricRequest) ProtoMessage() {}

func (x *MetricRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_proto_devops_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricRequest.ProtoReflect.Descriptor instead.
func (*MetricRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_proto_devops_proto_rawDescGZIP(), []int{9}
}

func (x *MetricRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MetricRequest) GetMtype() string {
	if x != nil {
		return x.Mtype
	}
	return ""
}

var File_internal_api_proto_devops_proto protoreflect.FileDescriptor

var file_internal_api_proto_devops_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x0d, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x13, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x3b, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x3b, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x22, 0x42, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x22, 0x3c, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x22, 0x6e,
	0x0a, 0x06, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64,
	0x65, 0x6c, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x35,
	0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6d, 0x74, 0x79, 0x70, 0x65, 0x32, 0x86, 0x02, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x6f, 0x70, 0x73,
	0x12, 0x3c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x15, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x18,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x2d, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_internal_api_proto_devops_proto_rawDescOnce sync.Once
	file_internal_api_proto_devops_proto_rawDescData = file_internal_api_proto_devops_proto_rawDesc
)

func file_internal_api_proto_devops_proto_rawDescGZIP() []byte {
	file_internal_api_proto_devops_proto_rawDescOnce.Do(func() {
		file_internal_api_proto_devops_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_api_proto_devops_proto_rawDescData)
	})
	return file_internal_api_proto_devops_proto_rawDescData
}

var file_internal_api_proto_devops_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_internal_api_proto_devops_proto_goTypes = []interface{}{
	(*PingRequest)(nil),           // 0: api.PingRequest
	(*PingResponse)(nil),          // 1: api.PingResponse
	(*GetMetricRequest)(nil),      // 2: api.GetMetricRequest
	(*GetMetricResponse)(nil),     // 3: api.GetMetricResponse
	(*CreateMetricRequest)(nil),   // 4: api.CreateMetricRequest
	(*CreateMetricResponse)(nil),  // 5: api.CreateMetricResponse
	(*CreateMetricsRequest)(nil),  // 6: api.CreateMetricsRequest
	(*CreateMetricsResponse)(nil), // 7: api.CreateMetricsResponse
	(*Metric)(nil),                // 8: api.Metric
	(*MetricRequest)(nil),         // 9: api.MetricRequest
}
var file_internal_api_proto_devops_proto_depIdxs = []int32{
	8, // 0: api.CreateMetricResponse.metric:type_name -> api.Metric
	9, // 1: api.CreateMetricsRequest.metric:type_name -> api.MetricRequest
	8, // 2: api.CreateMetricsResponse.metric:type_name -> api.Metric
	2, // 3: api.Devops.GetMetric:input_type -> api.GetMetricRequest
	4, // 4: api.Devops.CreateMetric:input_type -> api.CreateMetricRequest
	6, // 5: api.Devops.CreateMetrics:input_type -> api.CreateMetricsRequest
	0, // 6: api.Devops.Ping:input_type -> api.PingRequest
	3, // 7: api.Devops.GetMetric:output_type -> api.GetMetricResponse
	5, // 8: api.Devops.CreateMetric:output_type -> api.CreateMetricResponse
	7, // 9: api.Devops.CreateMetrics:output_type -> api.CreateMetricsResponse
	1, // 10: api.Devops.Ping:output_type -> api.PingResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internal_api_proto_devops_proto_init() }
func file_internal_api_proto_devops_proto_init() {
	if File_internal_api_proto_devops_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_api_proto_devops_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_internal_api_proto_devops_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_internal_api_proto_devops_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMetricRequest); i {
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
		file_internal_api_proto_devops_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMetricResponse); i {
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
		file_internal_api_proto_devops_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMetricRequest); i {
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
		file_internal_api_proto_devops_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMetricResponse); i {
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
		file_internal_api_proto_devops_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMetricsRequest); i {
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
		file_internal_api_proto_devops_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMetricsResponse); i {
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
		file_internal_api_proto_devops_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metric); i {
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
		file_internal_api_proto_devops_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricRequest); i {
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
			RawDescriptor: file_internal_api_proto_devops_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_api_proto_devops_proto_goTypes,
		DependencyIndexes: file_internal_api_proto_devops_proto_depIdxs,
		MessageInfos:      file_internal_api_proto_devops_proto_msgTypes,
	}.Build()
	File_internal_api_proto_devops_proto = out.File
	file_internal_api_proto_devops_proto_rawDesc = nil
	file_internal_api_proto_devops_proto_goTypes = nil
	file_internal_api_proto_devops_proto_depIdxs = nil
}
