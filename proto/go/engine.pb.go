// Copyright 2016-2018, KhulnaSoft, Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: khulnasoft/engine.proto

package khulnasoftrpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// LogSeverity is the severity level of a log message.  Errors are fatal; all others are informational.
type LogSeverity int32

const (
	LogSeverity_DEBUG   LogSeverity = 0 // a debug-level message not displayed to end-users (the default).
	LogSeverity_INFO    LogSeverity = 1 // an informational message printed to output during resource operations.
	LogSeverity_WARNING LogSeverity = 2 // a warning to indicate that something went wrong.
	LogSeverity_ERROR   LogSeverity = 3 // a fatal error indicating that the tool should stop processing subsequent resource operations.
)

// Enum value maps for LogSeverity.
var (
	LogSeverity_name = map[int32]string{
		0: "DEBUG",
		1: "INFO",
		2: "WARNING",
		3: "ERROR",
	}
	LogSeverity_value = map[string]int32{
		"DEBUG":   0,
		"INFO":    1,
		"WARNING": 2,
		"ERROR":   3,
	}
)

func (x LogSeverity) Enum() *LogSeverity {
	p := new(LogSeverity)
	*p = x
	return p
}

func (x LogSeverity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogSeverity) Descriptor() protoreflect.EnumDescriptor {
	return file_khulnasoft_engine_proto_enumTypes[0].Descriptor()
}

func (LogSeverity) Type() protoreflect.EnumType {
	return &file_khulnasoft_engine_proto_enumTypes[0]
}

func (x LogSeverity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogSeverity.Descriptor instead.
func (LogSeverity) EnumDescriptor() ([]byte, []int) {
	return file_khulnasoft_engine_proto_rawDescGZIP(), []int{0}
}

type LogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the logging level of this message.
	Severity LogSeverity `protobuf:"varint,1,opt,name=severity,proto3,enum=khulnasoftrpc.LogSeverity" json:"severity,omitempty"`
	// the contents of the logged message.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// the (optional) resource urn this log is associated with.
	Urn string `protobuf:"bytes,3,opt,name=urn,proto3" json:"urn,omitempty"`
	// the (optional) stream id that a stream of log messages can be associated with. This allows
	// clients to not have to buffer a large set of log messages that they all want to be
	// conceptually connected.  Instead the messages can be sent as chunks (with the same stream id)
	// and the end display can show the messages as they arrive, while still stitching them together
	// into one total log message.
	//
	// 0/not-given means: do not associate with any stream.
	StreamId int32 `protobuf:"varint,4,opt,name=streamId,proto3" json:"streamId,omitempty"`
	// Optional value indicating whether this is a status message.
	Ephemeral bool `protobuf:"varint,5,opt,name=ephemeral,proto3" json:"ephemeral,omitempty"`
}

func (x *LogRequest) Reset() {
	*x = LogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_khulnasoft_engine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRequest) ProtoMessage() {}

func (x *LogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_khulnasoft_engine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRequest.ProtoReflect.Descriptor instead.
func (*LogRequest) Descriptor() ([]byte, []int) {
	return file_khulnasoft_engine_proto_rawDescGZIP(), []int{0}
}

func (x *LogRequest) GetSeverity() LogSeverity {
	if x != nil {
		return x.Severity
	}
	return LogSeverity_DEBUG
}

func (x *LogRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LogRequest) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

func (x *LogRequest) GetStreamId() int32 {
	if x != nil {
		return x.StreamId
	}
	return 0
}

func (x *LogRequest) GetEphemeral() bool {
	if x != nil {
		return x.Ephemeral
	}
	return false
}

type GetRootResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRootResourceRequest) Reset() {
	*x = GetRootResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_khulnasoft_engine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRootResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRootResourceRequest) ProtoMessage() {}

func (x *GetRootResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_khulnasoft_engine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRootResourceRequest.ProtoReflect.Descriptor instead.
func (*GetRootResourceRequest) Descriptor() ([]byte, []int) {
	return file_khulnasoft_engine_proto_rawDescGZIP(), []int{1}
}

type GetRootResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the URN of the root resource, or the empty string if one was not set.
	Urn string `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
}

func (x *GetRootResourceResponse) Reset() {
	*x = GetRootResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_khulnasoft_engine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRootResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRootResourceResponse) ProtoMessage() {}

func (x *GetRootResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_khulnasoft_engine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRootResourceResponse.ProtoReflect.Descriptor instead.
func (*GetRootResourceResponse) Descriptor() ([]byte, []int) {
	return file_khulnasoft_engine_proto_rawDescGZIP(), []int{2}
}

func (x *GetRootResourceResponse) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

type SetRootResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the URN of the root resource, or the empty string.
	Urn string `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
}

func (x *SetRootResourceRequest) Reset() {
	*x = SetRootResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_khulnasoft_engine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRootResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRootResourceRequest) ProtoMessage() {}

func (x *SetRootResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_khulnasoft_engine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRootResourceRequest.ProtoReflect.Descriptor instead.
func (*SetRootResourceRequest) Descriptor() ([]byte, []int) {
	return file_khulnasoft_engine_proto_rawDescGZIP(), []int{3}
}

func (x *SetRootResourceRequest) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

type SetRootResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetRootResourceResponse) Reset() {
	*x = SetRootResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_khulnasoft_engine_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRootResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRootResourceResponse) ProtoMessage() {}

func (x *SetRootResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_khulnasoft_engine_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRootResourceResponse.ProtoReflect.Descriptor instead.
func (*SetRootResourceResponse) Descriptor() ([]byte, []int) {
	return file_khulnasoft_engine_proto_rawDescGZIP(), []int{4}
}

type StartDebuggingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the debug configuration parameters.  These are meant to be in the right format for the DAP protocol to consume.
	Config *structpb.Struct `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	// the string to display to the user with instructions on how to connect to the debugger.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StartDebuggingRequest) Reset() {
	*x = StartDebuggingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_khulnasoft_engine_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartDebuggingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartDebuggingRequest) ProtoMessage() {}

func (x *StartDebuggingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_khulnasoft_engine_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartDebuggingRequest.ProtoReflect.Descriptor instead.
func (*StartDebuggingRequest) Descriptor() ([]byte, []int) {
	return file_khulnasoft_engine_proto_rawDescGZIP(), []int{5}
}

func (x *StartDebuggingRequest) GetConfig() *structpb.Struct {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *StartDebuggingRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_khulnasoft_engine_proto protoreflect.FileDescriptor

var file_khulnasoft_engine_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x72, 0x70, 0x63,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x0a,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x73, 0x65,
	0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70,
	0x75, 0x6c, 0x75, 0x6d, 0x69, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x76, 0x65,
	0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x70, 0x68, 0x65, 0x6d, 0x65,
	0x72, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x70, 0x68, 0x65, 0x6d,
	0x65, 0x72, 0x61, 0x6c, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2b,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x22, 0x2a, 0x0a, 0x16, 0x53,
	0x65, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x22, 0x19, 0x0a, 0x17, 0x53, 0x65, 0x74, 0x52, 0x6f,
	0x6f, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x62, 0x0a, 0x15, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x65, 0x62, 0x75, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x3a, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x76,
	0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41,
	0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x03, 0x32, 0xc6, 0x02, 0x0a, 0x06, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x36, 0x0a,
	0x03, 0x4c, 0x6f, 0x67, 0x12, 0x15, 0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x72, 0x70, 0x63,
	0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x74,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x21, 0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d,
	0x69, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x75,
	0x6c, 0x75, 0x6d, 0x69, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x5a, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x21, 0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x72, 0x70, 0x63,
	0x2e, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a,
	0x0e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x65, 0x62, 0x75, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x12,
	0x20, 0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x44, 0x65, 0x62, 0x75, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69,
	0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x33, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x3b, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_khulnasoft_engine_proto_rawDescOnce sync.Once
	file_khulnasoft_engine_proto_rawDescData = file_khulnasoft_engine_proto_rawDesc
)

func file_khulnasoft_engine_proto_rawDescGZIP() []byte {
	file_khulnasoft_engine_proto_rawDescOnce.Do(func() {
		file_khulnasoft_engine_proto_rawDescData = protoimpl.X.CompressGZIP(file_khulnasoft_engine_proto_rawDescData)
	})
	return file_khulnasoft_engine_proto_rawDescData
}

var file_khulnasoft_engine_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_khulnasoft_engine_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_khulnasoft_engine_proto_goTypes = []interface{}{
	(LogSeverity)(0),                // 0: khulnasoftrpc.LogSeverity
	(*LogRequest)(nil),              // 1: khulnasoftrpc.LogRequest
	(*GetRootResourceRequest)(nil),  // 2: khulnasoftrpc.GetRootResourceRequest
	(*GetRootResourceResponse)(nil), // 3: khulnasoftrpc.GetRootResourceResponse
	(*SetRootResourceRequest)(nil),  // 4: khulnasoftrpc.SetRootResourceRequest
	(*SetRootResourceResponse)(nil), // 5: khulnasoftrpc.SetRootResourceResponse
	(*StartDebuggingRequest)(nil),   // 6: khulnasoftrpc.StartDebuggingRequest
	(*structpb.Struct)(nil),         // 7: google.protobuf.Struct
	(*emptypb.Empty)(nil),           // 8: google.protobuf.Empty
}
var file_khulnasoft_engine_proto_depIdxs = []int32{
	0, // 0: khulnasoftrpc.LogRequest.severity:type_name -> khulnasoftrpc.LogSeverity
	7, // 1: khulnasoftrpc.StartDebuggingRequest.config:type_name -> google.protobuf.Struct
	1, // 2: khulnasoftrpc.Engine.Log:input_type -> khulnasoftrpc.LogRequest
	2, // 3: khulnasoftrpc.Engine.GetRootResource:input_type -> khulnasoftrpc.GetRootResourceRequest
	4, // 4: khulnasoftrpc.Engine.SetRootResource:input_type -> khulnasoftrpc.SetRootResourceRequest
	6, // 5: khulnasoftrpc.Engine.StartDebugging:input_type -> khulnasoftrpc.StartDebuggingRequest
	8, // 6: khulnasoftrpc.Engine.Log:output_type -> google.protobuf.Empty
	3, // 7: khulnasoftrpc.Engine.GetRootResource:output_type -> khulnasoftrpc.GetRootResourceResponse
	5, // 8: khulnasoftrpc.Engine.SetRootResource:output_type -> khulnasoftrpc.SetRootResourceResponse
	8, // 9: khulnasoftrpc.Engine.StartDebugging:output_type -> google.protobuf.Empty
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_khulnasoft_engine_proto_init() }
func file_khulnasoft_engine_proto_init() {
	if File_khulnasoft_engine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_khulnasoft_engine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogRequest); i {
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
		file_khulnasoft_engine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRootResourceRequest); i {
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
		file_khulnasoft_engine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRootResourceResponse); i {
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
		file_khulnasoft_engine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRootResourceRequest); i {
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
		file_khulnasoft_engine_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRootResourceResponse); i {
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
		file_khulnasoft_engine_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartDebuggingRequest); i {
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
			RawDescriptor: file_khulnasoft_engine_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_khulnasoft_engine_proto_goTypes,
		DependencyIndexes: file_khulnasoft_engine_proto_depIdxs,
		EnumInfos:         file_khulnasoft_engine_proto_enumTypes,
		MessageInfos:      file_khulnasoft_engine_proto_msgTypes,
	}.Build()
	File_khulnasoft_engine_proto = out.File
	file_khulnasoft_engine_proto_rawDesc = nil
	file_khulnasoft_engine_proto_goTypes = nil
	file_khulnasoft_engine_proto_depIdxs = nil
}
