// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.11.4
// source: internal/pb/sensors/sensors.proto

package sensors_pb

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetVoltageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceID  string               `protobuf:"bytes,1,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
	StartTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime   *timestamp.Timestamp `protobuf:"bytes,3,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
}

func (x *GetVoltageRequest) Reset() {
	*x = GetVoltageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_sensors_sensors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVoltageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVoltageRequest) ProtoMessage() {}

func (x *GetVoltageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_sensors_sensors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVoltageRequest.ProtoReflect.Descriptor instead.
func (*GetVoltageRequest) Descriptor() ([]byte, []int) {
	return file_internal_pb_sensors_sensors_proto_rawDescGZIP(), []int{0}
}

func (x *GetVoltageRequest) GetDeviceID() string {
	if x != nil {
		return x.DeviceID
	}
	return ""
}

func (x *GetVoltageRequest) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *GetVoltageRequest) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type GetVoltageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Voltage []float32 `protobuf:"fixed32,1,rep,packed,name=Voltage,proto3" json:"Voltage,omitempty"`
}

func (x *GetVoltageResponse) Reset() {
	*x = GetVoltageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_sensors_sensors_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVoltageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVoltageResponse) ProtoMessage() {}

func (x *GetVoltageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_sensors_sensors_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVoltageResponse.ProtoReflect.Descriptor instead.
func (*GetVoltageResponse) Descriptor() ([]byte, []int) {
	return file_internal_pb_sensors_sensors_proto_rawDescGZIP(), []int{1}
}

func (x *GetVoltageResponse) GetVoltage() []float32 {
	if x != nil {
		return x.Voltage
	}
	return nil
}

type GetHumidityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceID  string               `protobuf:"bytes,1,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
	StartTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime   *timestamp.Timestamp `protobuf:"bytes,3,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
}

func (x *GetHumidityRequest) Reset() {
	*x = GetHumidityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_sensors_sensors_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHumidityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHumidityRequest) ProtoMessage() {}

func (x *GetHumidityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_sensors_sensors_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHumidityRequest.ProtoReflect.Descriptor instead.
func (*GetHumidityRequest) Descriptor() ([]byte, []int) {
	return file_internal_pb_sensors_sensors_proto_rawDescGZIP(), []int{2}
}

func (x *GetHumidityRequest) GetDeviceID() string {
	if x != nil {
		return x.DeviceID
	}
	return ""
}

func (x *GetHumidityRequest) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *GetHumidityRequest) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type GetHumidityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Humidity []float32 `protobuf:"fixed32,1,rep,packed,name=Humidity,proto3" json:"Humidity,omitempty"`
}

func (x *GetHumidityResponse) Reset() {
	*x = GetHumidityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_sensors_sensors_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHumidityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHumidityResponse) ProtoMessage() {}

func (x *GetHumidityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_sensors_sensors_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHumidityResponse.ProtoReflect.Descriptor instead.
func (*GetHumidityResponse) Descriptor() ([]byte, []int) {
	return file_internal_pb_sensors_sensors_proto_rawDescGZIP(), []int{3}
}

func (x *GetHumidityResponse) GetHumidity() []float32 {
	if x != nil {
		return x.Humidity
	}
	return nil
}

type GetTemperatureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceID  string               `protobuf:"bytes,1,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
	StartTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime   *timestamp.Timestamp `protobuf:"bytes,3,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
}

func (x *GetTemperatureRequest) Reset() {
	*x = GetTemperatureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_sensors_sensors_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemperatureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemperatureRequest) ProtoMessage() {}

func (x *GetTemperatureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_sensors_sensors_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemperatureRequest.ProtoReflect.Descriptor instead.
func (*GetTemperatureRequest) Descriptor() ([]byte, []int) {
	return file_internal_pb_sensors_sensors_proto_rawDescGZIP(), []int{4}
}

func (x *GetTemperatureRequest) GetDeviceID() string {
	if x != nil {
		return x.DeviceID
	}
	return ""
}

func (x *GetTemperatureRequest) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *GetTemperatureRequest) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type GetTemperatureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temperature []float32 `protobuf:"fixed32,1,rep,packed,name=Temperature,proto3" json:"Temperature,omitempty"`
}

func (x *GetTemperatureResponse) Reset() {
	*x = GetTemperatureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_sensors_sensors_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemperatureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemperatureResponse) ProtoMessage() {}

func (x *GetTemperatureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_sensors_sensors_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemperatureResponse.ProtoReflect.Descriptor instead.
func (*GetTemperatureResponse) Descriptor() ([]byte, []int) {
	return file_internal_pb_sensors_sensors_proto_rawDescGZIP(), []int{5}
}

func (x *GetTemperatureResponse) GetTemperature() []float32 {
	if x != nil {
		return x.Temperature
	}
	return nil
}

var File_internal_pb_sensors_sensors_proto protoreflect.FileDescriptor

var file_internal_pb_sensors_sensors_proto_rawDesc = []byte{
	0x0a, 0x21, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x73, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x5f, 0x70, 0x62, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9f, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x44, 0x12, 0x38, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07,
	0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x6f, 0x6c, 0x74,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x02, 0x52, 0x07, 0x56, 0x6f, 0x6c, 0x74, 0x61,
	0x67, 0x65, 0x22, 0xa0, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x48, 0x75, 0x6d, 0x69, 0x64, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x38, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x34, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x45, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x31, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x48, 0x75, 0x6d, 0x69,
	0x64, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x48, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x02, 0x52, 0x08,
	0x48, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x22, 0xa3, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x38,
	0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x3a,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x65, 0x6d, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x02, 0x52, 0x0b, 0x54,
	0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x32, 0x85, 0x02, 0x0a, 0x0d, 0x53,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x73, 0x5f, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x74, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x73, 0x5f, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x74, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x48, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x73, 0x5f, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x73, 0x5f, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x21, 0x2e, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x73, 0x5f, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x5f, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x3b, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x5f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_pb_sensors_sensors_proto_rawDescOnce sync.Once
	file_internal_pb_sensors_sensors_proto_rawDescData = file_internal_pb_sensors_sensors_proto_rawDesc
)

func file_internal_pb_sensors_sensors_proto_rawDescGZIP() []byte {
	file_internal_pb_sensors_sensors_proto_rawDescOnce.Do(func() {
		file_internal_pb_sensors_sensors_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pb_sensors_sensors_proto_rawDescData)
	})
	return file_internal_pb_sensors_sensors_proto_rawDescData
}

var file_internal_pb_sensors_sensors_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_pb_sensors_sensors_proto_goTypes = []interface{}{
	(*GetVoltageRequest)(nil),      // 0: sensors_pb.GetVoltageRequest
	(*GetVoltageResponse)(nil),     // 1: sensors_pb.GetVoltageResponse
	(*GetHumidityRequest)(nil),     // 2: sensors_pb.GetHumidityRequest
	(*GetHumidityResponse)(nil),    // 3: sensors_pb.GetHumidityResponse
	(*GetTemperatureRequest)(nil),  // 4: sensors_pb.GetTemperatureRequest
	(*GetTemperatureResponse)(nil), // 5: sensors_pb.GetTemperatureResponse
	(*timestamp.Timestamp)(nil),    // 6: google.protobuf.Timestamp
}
var file_internal_pb_sensors_sensors_proto_depIdxs = []int32{
	6, // 0: sensors_pb.GetVoltageRequest.StartTime:type_name -> google.protobuf.Timestamp
	6, // 1: sensors_pb.GetVoltageRequest.EndTime:type_name -> google.protobuf.Timestamp
	6, // 2: sensors_pb.GetHumidityRequest.StartTime:type_name -> google.protobuf.Timestamp
	6, // 3: sensors_pb.GetHumidityRequest.EndTime:type_name -> google.protobuf.Timestamp
	6, // 4: sensors_pb.GetTemperatureRequest.StartTime:type_name -> google.protobuf.Timestamp
	6, // 5: sensors_pb.GetTemperatureRequest.EndTime:type_name -> google.protobuf.Timestamp
	0, // 6: sensors_pb.SensorService.GetVoltage:input_type -> sensors_pb.GetVoltageRequest
	2, // 7: sensors_pb.SensorService.GetHumidity:input_type -> sensors_pb.GetHumidityRequest
	4, // 8: sensors_pb.SensorService.GetTemperature:input_type -> sensors_pb.GetTemperatureRequest
	1, // 9: sensors_pb.SensorService.GetVoltage:output_type -> sensors_pb.GetVoltageResponse
	3, // 10: sensors_pb.SensorService.GetHumidity:output_type -> sensors_pb.GetHumidityResponse
	5, // 11: sensors_pb.SensorService.GetTemperature:output_type -> sensors_pb.GetTemperatureResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_internal_pb_sensors_sensors_proto_init() }
func file_internal_pb_sensors_sensors_proto_init() {
	if File_internal_pb_sensors_sensors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_pb_sensors_sensors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVoltageRequest); i {
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
		file_internal_pb_sensors_sensors_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVoltageResponse); i {
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
		file_internal_pb_sensors_sensors_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHumidityRequest); i {
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
		file_internal_pb_sensors_sensors_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHumidityResponse); i {
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
		file_internal_pb_sensors_sensors_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemperatureRequest); i {
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
		file_internal_pb_sensors_sensors_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemperatureResponse); i {
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
			RawDescriptor: file_internal_pb_sensors_sensors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_pb_sensors_sensors_proto_goTypes,
		DependencyIndexes: file_internal_pb_sensors_sensors_proto_depIdxs,
		MessageInfos:      file_internal_pb_sensors_sensors_proto_msgTypes,
	}.Build()
	File_internal_pb_sensors_sensors_proto = out.File
	file_internal_pb_sensors_sensors_proto_rawDesc = nil
	file_internal_pb_sensors_sensors_proto_goTypes = nil
	file_internal_pb_sensors_sensors_proto_depIdxs = nil
}
