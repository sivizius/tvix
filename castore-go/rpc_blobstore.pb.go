// SPDX-License-Identifier: MIT
// Copyright © 2022 The Tvix Authors

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: tvix/castore/protos/rpc_blobstore.proto

package castorev1

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

type StatBlobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The blake3 digest of the blob requested
	Digest []byte `protobuf:"bytes,1,opt,name=digest,proto3" json:"digest,omitempty"`
	// Whether the server should reply with a list of more granular chunks.
	SendChunks bool `protobuf:"varint,2,opt,name=send_chunks,json=sendChunks,proto3" json:"send_chunks,omitempty"`
	// Whether the server should reply with a bao.
	SendBao bool `protobuf:"varint,3,opt,name=send_bao,json=sendBao,proto3" json:"send_bao,omitempty"`
}

func (x *StatBlobRequest) Reset() {
	*x = StatBlobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tvix_castore_protos_rpc_blobstore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatBlobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatBlobRequest) ProtoMessage() {}

func (x *StatBlobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tvix_castore_protos_rpc_blobstore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatBlobRequest.ProtoReflect.Descriptor instead.
func (*StatBlobRequest) Descriptor() ([]byte, []int) {
	return file_tvix_castore_protos_rpc_blobstore_proto_rawDescGZIP(), []int{0}
}

func (x *StatBlobRequest) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

func (x *StatBlobRequest) GetSendChunks() bool {
	if x != nil {
		return x.SendChunks
	}
	return false
}

func (x *StatBlobRequest) GetSendBao() bool {
	if x != nil {
		return x.SendBao
	}
	return false
}

type StatBlobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If `send_chunks` was set to true, this MAY contain a list of more
	// granular chunks, which then may be read individually via the `Read`
	// method.
	Chunks []*StatBlobResponse_ChunkMeta `protobuf:"bytes,2,rep,name=chunks,proto3" json:"chunks,omitempty"`
	// If `send_bao` was set to true, this MAY contain a outboard bao.
	// The exact format and message types here will still be fleshed out.
	Bao []byte `protobuf:"bytes,3,opt,name=bao,proto3" json:"bao,omitempty"`
}

func (x *StatBlobResponse) Reset() {
	*x = StatBlobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tvix_castore_protos_rpc_blobstore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatBlobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatBlobResponse) ProtoMessage() {}

func (x *StatBlobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tvix_castore_protos_rpc_blobstore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatBlobResponse.ProtoReflect.Descriptor instead.
func (*StatBlobResponse) Descriptor() ([]byte, []int) {
	return file_tvix_castore_protos_rpc_blobstore_proto_rawDescGZIP(), []int{1}
}

func (x *StatBlobResponse) GetChunks() []*StatBlobResponse_ChunkMeta {
	if x != nil {
		return x.Chunks
	}
	return nil
}

func (x *StatBlobResponse) GetBao() []byte {
	if x != nil {
		return x.Bao
	}
	return nil
}

type ReadBlobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The blake3 digest of the blob or chunk requested
	Digest []byte `protobuf:"bytes,1,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (x *ReadBlobRequest) Reset() {
	*x = ReadBlobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tvix_castore_protos_rpc_blobstore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadBlobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadBlobRequest) ProtoMessage() {}

func (x *ReadBlobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tvix_castore_protos_rpc_blobstore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadBlobRequest.ProtoReflect.Descriptor instead.
func (*ReadBlobRequest) Descriptor() ([]byte, []int) {
	return file_tvix_castore_protos_rpc_blobstore_proto_rawDescGZIP(), []int{2}
}

func (x *ReadBlobRequest) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

// This represents some bytes of a blob.
// Blobs are sent in smaller chunks to keep message sizes manageable.
type BlobChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BlobChunk) Reset() {
	*x = BlobChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tvix_castore_protos_rpc_blobstore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlobChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlobChunk) ProtoMessage() {}

func (x *BlobChunk) ProtoReflect() protoreflect.Message {
	mi := &file_tvix_castore_protos_rpc_blobstore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlobChunk.ProtoReflect.Descriptor instead.
func (*BlobChunk) Descriptor() ([]byte, []int) {
	return file_tvix_castore_protos_rpc_blobstore_proto_rawDescGZIP(), []int{3}
}

func (x *BlobChunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type PutBlobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The blake3 digest of the data that was sent.
	Digest []byte `protobuf:"bytes,1,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (x *PutBlobResponse) Reset() {
	*x = PutBlobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tvix_castore_protos_rpc_blobstore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutBlobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutBlobResponse) ProtoMessage() {}

func (x *PutBlobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tvix_castore_protos_rpc_blobstore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutBlobResponse.ProtoReflect.Descriptor instead.
func (*PutBlobResponse) Descriptor() ([]byte, []int) {
	return file_tvix_castore_protos_rpc_blobstore_proto_rawDescGZIP(), []int{4}
}

func (x *PutBlobResponse) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

type StatBlobResponse_ChunkMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Digest of that specific chunk
	Digest []byte `protobuf:"bytes,1,opt,name=digest,proto3" json:"digest,omitempty"`
	// Length of that chunk, in bytes.
	Size uint64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *StatBlobResponse_ChunkMeta) Reset() {
	*x = StatBlobResponse_ChunkMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tvix_castore_protos_rpc_blobstore_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatBlobResponse_ChunkMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatBlobResponse_ChunkMeta) ProtoMessage() {}

func (x *StatBlobResponse_ChunkMeta) ProtoReflect() protoreflect.Message {
	mi := &file_tvix_castore_protos_rpc_blobstore_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatBlobResponse_ChunkMeta.ProtoReflect.Descriptor instead.
func (*StatBlobResponse_ChunkMeta) Descriptor() ([]byte, []int) {
	return file_tvix_castore_protos_rpc_blobstore_proto_rawDescGZIP(), []int{1, 0}
}

func (x *StatBlobResponse_ChunkMeta) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

func (x *StatBlobResponse_ChunkMeta) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

var File_tvix_castore_protos_rpc_blobstore_proto protoreflect.FileDescriptor

var file_tvix_castore_protos_rpc_blobstore_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x76, 0x69, 0x78, 0x2f, 0x63, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74, 0x76, 0x69, 0x78, 0x2e,
	0x63, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x65, 0x0a, 0x0f, 0x53, 0x74,
	0x61, 0x74, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x64,
	0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x62,
	0x61, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x65, 0x6e, 0x64, 0x42, 0x61,
	0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x74, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x74, 0x76, 0x69, 0x78, 0x2e, 0x63, 0x61,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x42, 0x6c, 0x6f,
	0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x06, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x62,
	0x61, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x62, 0x61, 0x6f, 0x1a, 0x37, 0x0a,
	0x09, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x29, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6c,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67,
	0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73,
	0x74, 0x22, 0x1f, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x62, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x29, 0x0a, 0x0f, 0x50, 0x75, 0x74, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x32, 0xe9, 0x01,
	0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a,
	0x04, 0x53, 0x74, 0x61, 0x74, 0x12, 0x20, 0x2e, 0x74, 0x76, 0x69, 0x78, 0x2e, 0x63, 0x61, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x42, 0x6c, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x76, 0x69, 0x78, 0x2e, 0x63,
	0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x42, 0x6c,
	0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x04, 0x52, 0x65,
	0x61, 0x64, 0x12, 0x20, 0x2e, 0x74, 0x76, 0x69, 0x78, 0x2e, 0x63, 0x61, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x76, 0x69, 0x78, 0x2e, 0x63, 0x61, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x30, 0x01, 0x12, 0x45, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x1a, 0x2e, 0x74, 0x76, 0x69, 0x78,
	0x2e, 0x63, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x62,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x20, 0x2e, 0x74, 0x76, 0x69, 0x78, 0x2e, 0x63, 0x61, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x74, 0x42, 0x6c, 0x6f, 0x62, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x63, 0x6f, 0x64,
	0x65, 0x2e, 0x74, 0x76, 0x6c, 0x2e, 0x66, 0x79, 0x69, 0x2f, 0x74, 0x76, 0x69, 0x78, 0x2f, 0x63,
	0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2d, 0x67, 0x6f, 0x3b, 0x63, 0x61, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tvix_castore_protos_rpc_blobstore_proto_rawDescOnce sync.Once
	file_tvix_castore_protos_rpc_blobstore_proto_rawDescData = file_tvix_castore_protos_rpc_blobstore_proto_rawDesc
)

func file_tvix_castore_protos_rpc_blobstore_proto_rawDescGZIP() []byte {
	file_tvix_castore_protos_rpc_blobstore_proto_rawDescOnce.Do(func() {
		file_tvix_castore_protos_rpc_blobstore_proto_rawDescData = protoimpl.X.CompressGZIP(file_tvix_castore_protos_rpc_blobstore_proto_rawDescData)
	})
	return file_tvix_castore_protos_rpc_blobstore_proto_rawDescData
}

var file_tvix_castore_protos_rpc_blobstore_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tvix_castore_protos_rpc_blobstore_proto_goTypes = []interface{}{
	(*StatBlobRequest)(nil),            // 0: tvix.castore.v1.StatBlobRequest
	(*StatBlobResponse)(nil),           // 1: tvix.castore.v1.StatBlobResponse
	(*ReadBlobRequest)(nil),            // 2: tvix.castore.v1.ReadBlobRequest
	(*BlobChunk)(nil),                  // 3: tvix.castore.v1.BlobChunk
	(*PutBlobResponse)(nil),            // 4: tvix.castore.v1.PutBlobResponse
	(*StatBlobResponse_ChunkMeta)(nil), // 5: tvix.castore.v1.StatBlobResponse.ChunkMeta
}
var file_tvix_castore_protos_rpc_blobstore_proto_depIdxs = []int32{
	5, // 0: tvix.castore.v1.StatBlobResponse.chunks:type_name -> tvix.castore.v1.StatBlobResponse.ChunkMeta
	0, // 1: tvix.castore.v1.BlobService.Stat:input_type -> tvix.castore.v1.StatBlobRequest
	2, // 2: tvix.castore.v1.BlobService.Read:input_type -> tvix.castore.v1.ReadBlobRequest
	3, // 3: tvix.castore.v1.BlobService.Put:input_type -> tvix.castore.v1.BlobChunk
	1, // 4: tvix.castore.v1.BlobService.Stat:output_type -> tvix.castore.v1.StatBlobResponse
	3, // 5: tvix.castore.v1.BlobService.Read:output_type -> tvix.castore.v1.BlobChunk
	4, // 6: tvix.castore.v1.BlobService.Put:output_type -> tvix.castore.v1.PutBlobResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tvix_castore_protos_rpc_blobstore_proto_init() }
func file_tvix_castore_protos_rpc_blobstore_proto_init() {
	if File_tvix_castore_protos_rpc_blobstore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tvix_castore_protos_rpc_blobstore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatBlobRequest); i {
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
		file_tvix_castore_protos_rpc_blobstore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatBlobResponse); i {
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
		file_tvix_castore_protos_rpc_blobstore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadBlobRequest); i {
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
		file_tvix_castore_protos_rpc_blobstore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlobChunk); i {
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
		file_tvix_castore_protos_rpc_blobstore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutBlobResponse); i {
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
		file_tvix_castore_protos_rpc_blobstore_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatBlobResponse_ChunkMeta); i {
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
			RawDescriptor: file_tvix_castore_protos_rpc_blobstore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tvix_castore_protos_rpc_blobstore_proto_goTypes,
		DependencyIndexes: file_tvix_castore_protos_rpc_blobstore_proto_depIdxs,
		MessageInfos:      file_tvix_castore_protos_rpc_blobstore_proto_msgTypes,
	}.Build()
	File_tvix_castore_protos_rpc_blobstore_proto = out.File
	file_tvix_castore_protos_rpc_blobstore_proto_rawDesc = nil
	file_tvix_castore_protos_rpc_blobstore_proto_goTypes = nil
	file_tvix_castore_protos_rpc_blobstore_proto_depIdxs = nil
}
