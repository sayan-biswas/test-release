// Copyright 2020 The Tekton Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: resources.proto

package results_go_proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RecordSummary_Status int32

const (
	RecordSummary_UNKNOWN   RecordSummary_Status = 0
	RecordSummary_SUCCESS   RecordSummary_Status = 1
	RecordSummary_FAILURE   RecordSummary_Status = 2
	RecordSummary_TIMEOUT   RecordSummary_Status = 3
	RecordSummary_CANCELLED RecordSummary_Status = 4
)

// Enum value maps for RecordSummary_Status.
var (
	RecordSummary_Status_name = map[int32]string{
		0: "UNKNOWN",
		1: "SUCCESS",
		2: "FAILURE",
		3: "TIMEOUT",
		4: "CANCELLED",
	}
	RecordSummary_Status_value = map[string]int32{
		"UNKNOWN":   0,
		"SUCCESS":   1,
		"FAILURE":   2,
		"TIMEOUT":   3,
		"CANCELLED": 4,
	}
)

func (x RecordSummary_Status) Enum() *RecordSummary_Status {
	p := new(RecordSummary_Status)
	*p = x
	return p
}

func (x RecordSummary_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RecordSummary_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_resources_proto_enumTypes[0].Descriptor()
}

func (RecordSummary_Status) Type() protoreflect.EnumType {
	return &file_resources_proto_enumTypes[0]
}

func (x RecordSummary_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecordSummary_Status.Descriptor instead.
func (RecordSummary_Status) EnumDescriptor() ([]byte, []int) {
	return file_resources_proto_rawDescGZIP(), []int{3, 0}
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User assigned identifier of the Result.
	// Encodes parent information.
	//
	// Examples: namespace/default/results/1234
	//           cluster/<cluster-id>/namespace/tekton/results/1234
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Server assigned identifier of the Result.
	// DEPRECATED: use uid instead.
	//
	// Deprecated: Do not use.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Server assigned identifier of the Result.
	Uid string `protobuf:"bytes,7,opt,name=uid,proto3" json:"uid,omitempty"`
	// Server assigned timestamp for when the result was created.
	// DEPRECATED: use create_time instead.
	//
	// Deprecated: Do not use.
	CreatedTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// Server assigned timestamp for when the result was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Server assigned timestamp for when the results was updated.
	// DEPRECATED: use update_time instead.
	//
	// Deprecated: Do not use.
	UpdatedTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	// Server assigned timestamp for when the results was updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Arbitrary user provided labels for the result.
	Annotations map[string]string `protobuf:"bytes,4,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The etag for this result.
	// If this is provided on update, it must match the server's etag.
	Etag string `protobuf:"bytes,5,opt,name=etag,proto3" json:"etag,omitempty"`
	// High level overview of the root record for the Result. This is provided
	// as a convinence for clients to query Record state without needing to make
	// multiple calls to fetch the underlying Records.
	Summary *RecordSummary `protobuf:"bytes,10,opt,name=summary,proto3" json:"summary,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_resources_proto_rawDescGZIP(), []int{0}
}

func (x *Result) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Deprecated: Do not use.
func (x *Result) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Result) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

// Deprecated: Do not use.
func (x *Result) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *Result) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

// Deprecated: Do not use.
func (x *Result) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *Result) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Result) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *Result) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *Result) GetSummary() *RecordSummary {
	if x != nil {
		return x.Summary
	}
	return nil
}

// Record belonging to a Result. Typically will be Tekton
// Task/PipelineRuns, but may also include other execution information
// (e.g. alternative configs, DSLs, input payloads, post-execution actions, etc.)
type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name, must be rooted in parent result.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Server assigned identifier of the Result.
	// DEPRECATED: use uid instead.
	//
	// Deprecated: Do not use.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Server assigned identifier of the Result.
	Uid  string `protobuf:"bytes,7,opt,name=uid,proto3" json:"uid,omitempty"`
	Data *Any   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// The etag for this record.
	// If this is provided on update, it must match the server's etag.
	Etag string `protobuf:"bytes,4,opt,name=etag,proto3" json:"etag,omitempty"`
	// Server assigned timestamp for when the result was created.
	// DEPRECATED: use create_time instead.
	//
	// Deprecated: Do not use.
	CreatedTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// Server assigned timestamp for when the result was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Server assigned timestamp for when the results was updated.
	// DEPRECATED: use update_time instead.
	//
	// Deprecated: Do not use.
	UpdatedTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	// Server assigned timestamp for when the results was updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_resources_proto_rawDescGZIP(), []int{1}
}

func (x *Record) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Deprecated: Do not use.
func (x *Record) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Record) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Record) GetData() *Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Record) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

// Deprecated: Do not use.
func (x *Record) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *Record) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

// Deprecated: Do not use.
func (x *Record) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *Record) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// Any represents loosely typed data to be stored within a Record.
type Any struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier of the data type stored in value.
	// This is used as a type hint to determine how to unmarshal value.
	// Limited to 128 characters.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// JSON-encoded data.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Any) Reset() {
	*x = Any{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Any) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Any) ProtoMessage() {}

func (x *Any) ProtoReflect() protoreflect.Message {
	mi := &file_resources_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Any.ProtoReflect.Descriptor instead.
func (*Any) Descriptor() ([]byte, []int) {
	return file_resources_proto_rawDescGZIP(), []int{2}
}

func (x *Any) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Any) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

// RecordSummary is a high level overview of a Record, typically representing a
// "root" record for a result. It includes type agnostic information so that
// UIs and other tools do not need to be aware of underlying types.
type RecordSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the Record this summary represents.
	Record string `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	// Identifier of underlying data.
	// e.g. `pipelines.tekton.dev/PipelineRun`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Common Record agnostic fields.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Completion status of the Record.
	Status RecordSummary_Status `protobuf:"varint,5,opt,name=status,proto3,enum=tekton.results.v1alpha2.RecordSummary_Status" json:"status,omitempty"`
	// Key-value pairs representing arbitrary underlying record data that clients want to include
	// that aren't covered by the above fields.
	Annotations map[string]string `protobuf:"bytes,6,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RecordSummary) Reset() {
	*x = RecordSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordSummary) ProtoMessage() {}

func (x *RecordSummary) ProtoReflect() protoreflect.Message {
	mi := &file_resources_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordSummary.ProtoReflect.Descriptor instead.
func (*RecordSummary) Descriptor() ([]byte, []int) {
	return file_resources_proto_rawDescGZIP(), []int{3}
}

func (x *RecordSummary) GetRecord() string {
	if x != nil {
		return x.Record
	}
	return ""
}

func (x *RecordSummary) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *RecordSummary) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *RecordSummary) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *RecordSummary) GetStatus() RecordSummary_Status {
	if x != nil {
		return x.Status
	}
	return RecordSummary_UNKNOWN
}

func (x *RecordSummary) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

// Log is a chunk of a log.
type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name fo the log
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The log data
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_resources_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_resources_proto_rawDescGZIP(), []int{4}
}

func (x *Log) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Log) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type LogSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the Record this summary represents.
	Record string `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	// Number of bytes received while streaming
	BytesReceived int64 `protobuf:"varint,2,opt,name=bytesReceived,proto3" json:"bytesReceived,omitempty"`
}

func (x *LogSummary) Reset() {
	*x = LogSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogSummary) ProtoMessage() {}

func (x *LogSummary) ProtoReflect() protoreflect.Message {
	mi := &file_resources_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogSummary.ProtoReflect.Descriptor instead.
func (*LogSummary) Descriptor() ([]byte, []int) {
	return file_resources_proto_rawDescGZIP(), []int{5}
}

func (x *LogSummary) GetRecord() string {
	if x != nil {
		return x.Record
	}
	return ""
}

func (x *LogSummary) GetBytesReceived() int64 {
	if x != nil {
		return x.BytesReceived
	}
	return 0
}

var File_resources_proto protoreflect.FileDescriptor

var file_resources_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x17, 0x74, 0x65, 0x6b, 0x74, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x04, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x37, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x23, 0xfa, 0x41, 0x20, 0x12, 0x1e, 0x74, 0x65, 0x6b, 0x74, 0x6f, 0x6e, 0x2e, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0x18, 0x01, 0xe0, 0x41, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x44, 0x0a, 0x0c, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x05, 0x18, 0x01, 0xe0,
	0x41, 0x03, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x44, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x05, 0x18, 0x01, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x52, 0x0a, 0x0b, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30,
	0x2e, 0x74, 0x65, 0x6b, 0x74, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x17, 0x0a,
	0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x12, 0x40, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x6b, 0x74, 0x6f, 0x6e,
	0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52,
	0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc5, 0x03, 0x0a, 0x06, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x12, 0x37, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x23, 0xfa, 0x41, 0x20, 0x12, 0x1e, 0x74, 0x65, 0x6b, 0x74, 0x6f, 0x6e, 0x2e, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0x18, 0x01, 0xe0, 0x41, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6b, 0x74, 0x6f,
	0x6e, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04,
	0x65, 0x74, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67,
	0x12, 0x44, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x05, 0x18, 0x01, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x05, 0x18, 0x01, 0xe0, 0x41,
	0x03, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0x2f, 0x0a, 0x03, 0x41, 0x6e, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0xfc, 0x03, 0x0a, 0x0d, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x12, 0x36, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x1e, 0xfa, 0x41, 0x1b, 0x0a, 0x19, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x2e, 0x74, 0x65, 0x6b, 0x74, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x45, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2d, 0x2e, 0x74, 0x65, 0x6b, 0x74, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x59, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e,
	0x74, 0x65, 0x6b, 0x74, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x4b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55,
	0x52, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10,
	0x03, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x04,
	0x22, 0x4f, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x3a,
	0x20, 0xea, 0x41, 0x1d, 0x0a, 0x1b, 0x74, 0x65, 0x6b, 0x74, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x4c, 0x6f,
	0x67, 0x22, 0x6a, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12,
	0x36, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x1e, 0xfa, 0x41, 0x1b, 0x0a, 0x19, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x74, 0x65,
	0x6b, 0x74, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52,
	0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x42, 0x3d, 0x5a,
	0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6b, 0x74,
	0x6f, 0x6e, 0x63, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_proto_rawDescOnce sync.Once
	file_resources_proto_rawDescData = file_resources_proto_rawDesc
)

func file_resources_proto_rawDescGZIP() []byte {
	file_resources_proto_rawDescOnce.Do(func() {
		file_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_proto_rawDescData)
	})
	return file_resources_proto_rawDescData
}

var file_resources_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_resources_proto_goTypes = []interface{}{
	(RecordSummary_Status)(0),     // 0: tekton.results.v1alpha2.RecordSummary.Status
	(*Result)(nil),                // 1: tekton.results.v1alpha2.Result
	(*Record)(nil),                // 2: tekton.results.v1alpha2.Record
	(*Any)(nil),                   // 3: tekton.results.v1alpha2.Any
	(*RecordSummary)(nil),         // 4: tekton.results.v1alpha2.RecordSummary
	(*Log)(nil),                   // 5: tekton.results.v1alpha2.Log
	(*LogSummary)(nil),            // 6: tekton.results.v1alpha2.LogSummary
	nil,                           // 7: tekton.results.v1alpha2.Result.AnnotationsEntry
	nil,                           // 8: tekton.results.v1alpha2.RecordSummary.AnnotationsEntry
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_resources_proto_depIdxs = []int32{
	9,  // 0: tekton.results.v1alpha2.Result.created_time:type_name -> google.protobuf.Timestamp
	9,  // 1: tekton.results.v1alpha2.Result.create_time:type_name -> google.protobuf.Timestamp
	9,  // 2: tekton.results.v1alpha2.Result.updated_time:type_name -> google.protobuf.Timestamp
	9,  // 3: tekton.results.v1alpha2.Result.update_time:type_name -> google.protobuf.Timestamp
	7,  // 4: tekton.results.v1alpha2.Result.annotations:type_name -> tekton.results.v1alpha2.Result.AnnotationsEntry
	4,  // 5: tekton.results.v1alpha2.Result.summary:type_name -> tekton.results.v1alpha2.RecordSummary
	3,  // 6: tekton.results.v1alpha2.Record.data:type_name -> tekton.results.v1alpha2.Any
	9,  // 7: tekton.results.v1alpha2.Record.created_time:type_name -> google.protobuf.Timestamp
	9,  // 8: tekton.results.v1alpha2.Record.create_time:type_name -> google.protobuf.Timestamp
	9,  // 9: tekton.results.v1alpha2.Record.updated_time:type_name -> google.protobuf.Timestamp
	9,  // 10: tekton.results.v1alpha2.Record.update_time:type_name -> google.protobuf.Timestamp
	9,  // 11: tekton.results.v1alpha2.RecordSummary.start_time:type_name -> google.protobuf.Timestamp
	9,  // 12: tekton.results.v1alpha2.RecordSummary.end_time:type_name -> google.protobuf.Timestamp
	0,  // 13: tekton.results.v1alpha2.RecordSummary.status:type_name -> tekton.results.v1alpha2.RecordSummary.Status
	8,  // 14: tekton.results.v1alpha2.RecordSummary.annotations:type_name -> tekton.results.v1alpha2.RecordSummary.AnnotationsEntry
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_resources_proto_init() }
func file_resources_proto_init() {
	if File_resources_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
		file_resources_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Any); i {
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
		file_resources_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordSummary); i {
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
		file_resources_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_resources_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogSummary); i {
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
			RawDescriptor: file_resources_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_proto_goTypes,
		DependencyIndexes: file_resources_proto_depIdxs,
		EnumInfos:         file_resources_proto_enumTypes,
		MessageInfos:      file_resources_proto_msgTypes,
	}.Build()
	File_resources_proto = out.File
	file_resources_proto_rawDesc = nil
	file_resources_proto_goTypes = nil
	file_resources_proto_depIdxs = nil
}
