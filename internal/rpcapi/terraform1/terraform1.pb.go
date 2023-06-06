// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.6
// source: terraform1.proto

package terraform1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Handshake struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Handshake) Reset() {
	*x = Handshake{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Handshake) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Handshake) ProtoMessage() {}

func (x *Handshake) ProtoReflect() protoreflect.Message {
	mi := &file_terraform1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Handshake.ProtoReflect.Descriptor instead.
func (*Handshake) Descriptor() ([]byte, []int) {
	return file_terraform1_proto_rawDescGZIP(), []int{0}
}

type OpenSourceBundle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OpenSourceBundle) Reset() {
	*x = OpenSourceBundle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenSourceBundle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenSourceBundle) ProtoMessage() {}

func (x *OpenSourceBundle) ProtoReflect() protoreflect.Message {
	mi := &file_terraform1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenSourceBundle.ProtoReflect.Descriptor instead.
func (*OpenSourceBundle) Descriptor() ([]byte, []int) {
	return file_terraform1_proto_rawDescGZIP(), []int{1}
}

type CloseSourceBundle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CloseSourceBundle) Reset() {
	*x = CloseSourceBundle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseSourceBundle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseSourceBundle) ProtoMessage() {}

func (x *CloseSourceBundle) ProtoReflect() protoreflect.Message {
	mi := &file_terraform1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseSourceBundle.ProtoReflect.Descriptor instead.
func (*CloseSourceBundle) Descriptor() ([]byte, []int) {
	return file_terraform1_proto_rawDescGZIP(), []int{2}
}

// The capabilities that the client wishes to advertise to the server during
// handshake.
type ClientCapabilities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClientCapabilities) Reset() {
	*x = ClientCapabilities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientCapabilities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientCapabilities) ProtoMessage() {}

func (x *ClientCapabilities) ProtoReflect() protoreflect.Message {
	mi := &file_terraform1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientCapabilities.ProtoReflect.Descriptor instead.
func (*ClientCapabilities) Descriptor() ([]byte, []int) {
	return file_terraform1_proto_rawDescGZIP(), []int{3}
}

// The capabilities that the server wishes to advertise to the client during
// handshake. Fields in this message can also be used to acknowledge and
// confirm support for client capabilities advertised in ClientCapabilities,
// in situations where the client must vary its behavior based on the server's
// level of support.
type ServerCapabilities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServerCapabilities) Reset() {
	*x = ServerCapabilities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerCapabilities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerCapabilities) ProtoMessage() {}

func (x *ServerCapabilities) ProtoReflect() protoreflect.Message {
	mi := &file_terraform1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerCapabilities.ProtoReflect.Descriptor instead.
func (*ServerCapabilities) Descriptor() ([]byte, []int) {
	return file_terraform1_proto_rawDescGZIP(), []int{4}
}

type Handshake_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Capabilities *ClientCapabilities `protobuf:"bytes,1,opt,name=capabilities,proto3" json:"capabilities,omitempty"`
}

func (x *Handshake_Request) Reset() {
	*x = Handshake_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Handshake_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Handshake_Request) ProtoMessage() {}

func (x *Handshake_Request) ProtoReflect() protoreflect.Message {
	mi := &file_terraform1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Handshake_Request.ProtoReflect.Descriptor instead.
func (*Handshake_Request) Descriptor() ([]byte, []int) {
	return file_terraform1_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Handshake_Request) GetCapabilities() *ClientCapabilities {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

type Handshake_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Capabilities *ServerCapabilities `protobuf:"bytes,2,opt,name=capabilities,proto3" json:"capabilities,omitempty"`
}

func (x *Handshake_Response) Reset() {
	*x = Handshake_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Handshake_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Handshake_Response) ProtoMessage() {}

func (x *Handshake_Response) ProtoReflect() protoreflect.Message {
	mi := &file_terraform1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Handshake_Response.ProtoReflect.Descriptor instead.
func (*Handshake_Response) Descriptor() ([]byte, []int) {
	return file_terraform1_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Handshake_Response) GetCapabilities() *ServerCapabilities {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

type OpenSourceBundle_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalPath string `protobuf:"bytes,1,opt,name=local_path,json=localPath,proto3" json:"local_path,omitempty"`
}

func (x *OpenSourceBundle_Request) Reset() {
	*x = OpenSourceBundle_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenSourceBundle_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenSourceBundle_Request) ProtoMessage() {}

func (x *OpenSourceBundle_Request) ProtoReflect() protoreflect.Message {
	mi := &file_terraform1_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenSourceBundle_Request.ProtoReflect.Descriptor instead.
func (*OpenSourceBundle_Request) Descriptor() ([]byte, []int) {
	return file_terraform1_proto_rawDescGZIP(), []int{1, 0}
}

func (x *OpenSourceBundle_Request) GetLocalPath() string {
	if x != nil {
		return x.LocalPath
	}
	return ""
}

type OpenSourceBundle_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceBundleHandle int64 `protobuf:"varint,1,opt,name=source_bundle_handle,json=sourceBundleHandle,proto3" json:"source_bundle_handle,omitempty"`
}

func (x *OpenSourceBundle_Response) Reset() {
	*x = OpenSourceBundle_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform1_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenSourceBundle_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenSourceBundle_Response) ProtoMessage() {}

func (x *OpenSourceBundle_Response) ProtoReflect() protoreflect.Message {
	mi := &file_terraform1_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenSourceBundle_Response.ProtoReflect.Descriptor instead.
func (*OpenSourceBundle_Response) Descriptor() ([]byte, []int) {
	return file_terraform1_proto_rawDescGZIP(), []int{1, 1}
}

func (x *OpenSourceBundle_Response) GetSourceBundleHandle() int64 {
	if x != nil {
		return x.SourceBundleHandle
	}
	return 0
}

type CloseSourceBundle_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceBundleHandle int64 `protobuf:"varint,1,opt,name=source_bundle_handle,json=sourceBundleHandle,proto3" json:"source_bundle_handle,omitempty"`
}

func (x *CloseSourceBundle_Request) Reset() {
	*x = CloseSourceBundle_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform1_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseSourceBundle_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseSourceBundle_Request) ProtoMessage() {}

func (x *CloseSourceBundle_Request) ProtoReflect() protoreflect.Message {
	mi := &file_terraform1_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseSourceBundle_Request.ProtoReflect.Descriptor instead.
func (*CloseSourceBundle_Request) Descriptor() ([]byte, []int) {
	return file_terraform1_proto_rawDescGZIP(), []int{2, 0}
}

func (x *CloseSourceBundle_Request) GetSourceBundleHandle() int64 {
	if x != nil {
		return x.SourceBundleHandle
	}
	return 0
}

type CloseSourceBundle_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CloseSourceBundle_Response) Reset() {
	*x = CloseSourceBundle_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform1_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseSourceBundle_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseSourceBundle_Response) ProtoMessage() {}

func (x *CloseSourceBundle_Response) ProtoReflect() protoreflect.Message {
	mi := &file_terraform1_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseSourceBundle_Response.ProtoReflect.Descriptor instead.
func (*CloseSourceBundle_Response) Descriptor() ([]byte, []int) {
	return file_terraform1_proto_rawDescGZIP(), []int{2, 1}
}

var File_terraform1_proto protoreflect.FileDescriptor

var file_terraform1_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x31, 0x22, 0xaa,
	0x01, 0x0a, 0x09, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x1a, 0x4d, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x0c, 0x63,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x4e, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x0c, 0x63,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x7a, 0x0a, 0x10, 0x4f,
	0x70, 0x65, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x1a,
	0x28, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x1a, 0x3c, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x12, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x22, 0x5c, 0x0a, 0x11, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x1a, 0x3b, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x75, 0x6e,
	0x64, 0x6c, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x1a, 0x0a, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x14, 0x0a, 0x12, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x32, 0x53, 0x0a, 0x05, 0x53, 0x65, 0x74, 0x75, 0x70, 0x12, 0x4a, 0x0a, 0x09, 0x48, 0x61,
	0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x12, 0x1d, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66,
	0x6f, 0x72, 0x6d, 0x31, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f,
	0x72, 0x6d, 0x31, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xd3, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x70, 0x65, 0x6e,
	0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x5f, 0x0a, 0x10, 0x4f, 0x70, 0x65, 0x6e, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x24, 0x2e, 0x74, 0x65,
	0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x31, 0x2e, 0x4f,
	0x70, 0x65, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x11, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x25, 0x2e,
	0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d,
	0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x75, 0x6e,
	0x64, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_terraform1_proto_rawDescOnce sync.Once
	file_terraform1_proto_rawDescData = file_terraform1_proto_rawDesc
)

func file_terraform1_proto_rawDescGZIP() []byte {
	file_terraform1_proto_rawDescOnce.Do(func() {
		file_terraform1_proto_rawDescData = protoimpl.X.CompressGZIP(file_terraform1_proto_rawDescData)
	})
	return file_terraform1_proto_rawDescData
}

var file_terraform1_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_terraform1_proto_goTypes = []interface{}{
	(*Handshake)(nil),                  // 0: terraform1.Handshake
	(*OpenSourceBundle)(nil),           // 1: terraform1.OpenSourceBundle
	(*CloseSourceBundle)(nil),          // 2: terraform1.CloseSourceBundle
	(*ClientCapabilities)(nil),         // 3: terraform1.ClientCapabilities
	(*ServerCapabilities)(nil),         // 4: terraform1.ServerCapabilities
	(*Handshake_Request)(nil),          // 5: terraform1.Handshake.Request
	(*Handshake_Response)(nil),         // 6: terraform1.Handshake.Response
	(*OpenSourceBundle_Request)(nil),   // 7: terraform1.OpenSourceBundle.Request
	(*OpenSourceBundle_Response)(nil),  // 8: terraform1.OpenSourceBundle.Response
	(*CloseSourceBundle_Request)(nil),  // 9: terraform1.CloseSourceBundle.Request
	(*CloseSourceBundle_Response)(nil), // 10: terraform1.CloseSourceBundle.Response
}
var file_terraform1_proto_depIdxs = []int32{
	3,  // 0: terraform1.Handshake.Request.capabilities:type_name -> terraform1.ClientCapabilities
	4,  // 1: terraform1.Handshake.Response.capabilities:type_name -> terraform1.ServerCapabilities
	5,  // 2: terraform1.Setup.Handshake:input_type -> terraform1.Handshake.Request
	7,  // 3: terraform1.Dependencies.OpenSourceBundle:input_type -> terraform1.OpenSourceBundle.Request
	9,  // 4: terraform1.Dependencies.CloseSourceBundle:input_type -> terraform1.CloseSourceBundle.Request
	6,  // 5: terraform1.Setup.Handshake:output_type -> terraform1.Handshake.Response
	8,  // 6: terraform1.Dependencies.OpenSourceBundle:output_type -> terraform1.OpenSourceBundle.Response
	10, // 7: terraform1.Dependencies.CloseSourceBundle:output_type -> terraform1.CloseSourceBundle.Response
	5,  // [5:8] is the sub-list for method output_type
	2,  // [2:5] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_terraform1_proto_init() }
func file_terraform1_proto_init() {
	if File_terraform1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_terraform1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Handshake); i {
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
		file_terraform1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenSourceBundle); i {
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
		file_terraform1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseSourceBundle); i {
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
		file_terraform1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientCapabilities); i {
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
		file_terraform1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerCapabilities); i {
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
		file_terraform1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Handshake_Request); i {
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
		file_terraform1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Handshake_Response); i {
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
		file_terraform1_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenSourceBundle_Request); i {
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
		file_terraform1_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenSourceBundle_Response); i {
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
		file_terraform1_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseSourceBundle_Request); i {
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
		file_terraform1_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseSourceBundle_Response); i {
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
			RawDescriptor: file_terraform1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_terraform1_proto_goTypes,
		DependencyIndexes: file_terraform1_proto_depIdxs,
		MessageInfos:      file_terraform1_proto_msgTypes,
	}.Build()
	File_terraform1_proto = out.File
	file_terraform1_proto_rawDesc = nil
	file_terraform1_proto_goTypes = nil
	file_terraform1_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SetupClient is the client API for Setup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SetupClient interface {
	// Clients must call Handshake before any other function of any other
	// service, to complete the capability negotiation step that may
	// then affect the behaviors of subsequent operations.
	//
	// This function can be called only once per RPC server.
	Handshake(ctx context.Context, in *Handshake_Request, opts ...grpc.CallOption) (*Handshake_Response, error)
}

type setupClient struct {
	cc grpc.ClientConnInterface
}

func NewSetupClient(cc grpc.ClientConnInterface) SetupClient {
	return &setupClient{cc}
}

func (c *setupClient) Handshake(ctx context.Context, in *Handshake_Request, opts ...grpc.CallOption) (*Handshake_Response, error) {
	out := new(Handshake_Response)
	err := c.cc.Invoke(ctx, "/terraform1.Setup/Handshake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SetupServer is the server API for Setup service.
type SetupServer interface {
	// Clients must call Handshake before any other function of any other
	// service, to complete the capability negotiation step that may
	// then affect the behaviors of subsequent operations.
	//
	// This function can be called only once per RPC server.
	Handshake(context.Context, *Handshake_Request) (*Handshake_Response, error)
}

// UnimplementedSetupServer can be embedded to have forward compatible implementations.
type UnimplementedSetupServer struct {
}

func (*UnimplementedSetupServer) Handshake(context.Context, *Handshake_Request) (*Handshake_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handshake not implemented")
}

func RegisterSetupServer(s *grpc.Server, srv SetupServer) {
	s.RegisterService(&_Setup_serviceDesc, srv)
}

func _Setup_Handshake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Handshake_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SetupServer).Handshake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/terraform1.Setup/Handshake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SetupServer).Handshake(ctx, req.(*Handshake_Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Setup_serviceDesc = grpc.ServiceDesc{
	ServiceName: "terraform1.Setup",
	HandlerType: (*SetupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handshake",
			Handler:    _Setup_Handshake_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "terraform1.proto",
}

// DependenciesClient is the client API for Dependencies service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DependenciesClient interface {
	// Opens a source bundle that was already extracted into the filesystem
	// somewhere, returning an opaque source bundle handle that can be used for
	// subsequent operations.
	OpenSourceBundle(ctx context.Context, in *OpenSourceBundle_Request, opts ...grpc.CallOption) (*OpenSourceBundle_Response, error)
	// Closes a previously-opened source bundle, invalidating the given handle
	// and therefore making it safe to delete or modify the bundle directory
	// on disk.
	CloseSourceBundle(ctx context.Context, in *CloseSourceBundle_Request, opts ...grpc.CallOption) (*CloseSourceBundle_Response, error)
}

type dependenciesClient struct {
	cc grpc.ClientConnInterface
}

func NewDependenciesClient(cc grpc.ClientConnInterface) DependenciesClient {
	return &dependenciesClient{cc}
}

func (c *dependenciesClient) OpenSourceBundle(ctx context.Context, in *OpenSourceBundle_Request, opts ...grpc.CallOption) (*OpenSourceBundle_Response, error) {
	out := new(OpenSourceBundle_Response)
	err := c.cc.Invoke(ctx, "/terraform1.Dependencies/OpenSourceBundle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dependenciesClient) CloseSourceBundle(ctx context.Context, in *CloseSourceBundle_Request, opts ...grpc.CallOption) (*CloseSourceBundle_Response, error) {
	out := new(CloseSourceBundle_Response)
	err := c.cc.Invoke(ctx, "/terraform1.Dependencies/CloseSourceBundle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DependenciesServer is the server API for Dependencies service.
type DependenciesServer interface {
	// Opens a source bundle that was already extracted into the filesystem
	// somewhere, returning an opaque source bundle handle that can be used for
	// subsequent operations.
	OpenSourceBundle(context.Context, *OpenSourceBundle_Request) (*OpenSourceBundle_Response, error)
	// Closes a previously-opened source bundle, invalidating the given handle
	// and therefore making it safe to delete or modify the bundle directory
	// on disk.
	CloseSourceBundle(context.Context, *CloseSourceBundle_Request) (*CloseSourceBundle_Response, error)
}

// UnimplementedDependenciesServer can be embedded to have forward compatible implementations.
type UnimplementedDependenciesServer struct {
}

func (*UnimplementedDependenciesServer) OpenSourceBundle(context.Context, *OpenSourceBundle_Request) (*OpenSourceBundle_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenSourceBundle not implemented")
}
func (*UnimplementedDependenciesServer) CloseSourceBundle(context.Context, *CloseSourceBundle_Request) (*CloseSourceBundle_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseSourceBundle not implemented")
}

func RegisterDependenciesServer(s *grpc.Server, srv DependenciesServer) {
	s.RegisterService(&_Dependencies_serviceDesc, srv)
}

func _Dependencies_OpenSourceBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenSourceBundle_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DependenciesServer).OpenSourceBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/terraform1.Dependencies/OpenSourceBundle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DependenciesServer).OpenSourceBundle(ctx, req.(*OpenSourceBundle_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dependencies_CloseSourceBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseSourceBundle_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DependenciesServer).CloseSourceBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/terraform1.Dependencies/CloseSourceBundle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DependenciesServer).CloseSourceBundle(ctx, req.(*CloseSourceBundle_Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dependencies_serviceDesc = grpc.ServiceDesc{
	ServiceName: "terraform1.Dependencies",
	HandlerType: (*DependenciesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenSourceBundle",
			Handler:    _Dependencies_OpenSourceBundle_Handler,
		},
		{
			MethodName: "CloseSourceBundle",
			Handler:    _Dependencies_CloseSourceBundle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "terraform1.proto",
}
