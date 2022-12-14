// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: face.proto

package pbface2

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

type EmptyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMessage) Reset() {
	*x = EmptyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_face_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMessage) ProtoMessage() {}

func (x *EmptyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_face_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMessage.ProtoReflect.Descriptor instead.
func (*EmptyMessage) Descriptor() ([]byte, []int) {
	return file_face_proto_rawDescGZIP(), []int{0}
}

type ErrMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrString string `protobuf:"bytes,1,opt,name=errString,proto3" json:"errString,omitempty"`
}

func (x *ErrMessage) Reset() {
	*x = ErrMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_face_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrMessage) ProtoMessage() {}

func (x *ErrMessage) ProtoReflect() protoreflect.Message {
	mi := &file_face_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrMessage.ProtoReflect.Descriptor instead.
func (*ErrMessage) Descriptor() ([]byte, []int) {
	return file_face_proto_rawDescGZIP(), []int{1}
}

func (x *ErrMessage) GetErrString() string {
	if x != nil {
		return x.ErrString
	}
	return ""
}

type ClusterIDMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterID string `protobuf:"bytes,1,opt,name=clusterID,proto3" json:"clusterID,omitempty"`
}

func (x *ClusterIDMessage) Reset() {
	*x = ClusterIDMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_face_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterIDMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterIDMessage) ProtoMessage() {}

func (x *ClusterIDMessage) ProtoReflect() protoreflect.Message {
	mi := &file_face_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterIDMessage.ProtoReflect.Descriptor instead.
func (*ClusterIDMessage) Descriptor() ([]byte, []int) {
	return file_face_proto_rawDescGZIP(), []int{2}
}

func (x *ClusterIDMessage) GetClusterID() string {
	if x != nil {
		return x.ClusterID
	}
	return ""
}

type FaceIDMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FaceID int32 `protobuf:"varint,1,opt,name=faceID,proto3" json:"faceID,omitempty"`
}

func (x *FaceIDMessage) Reset() {
	*x = FaceIDMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_face_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaceIDMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaceIDMessage) ProtoMessage() {}

func (x *FaceIDMessage) ProtoReflect() protoreflect.Message {
	mi := &file_face_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaceIDMessage.ProtoReflect.Descriptor instead.
func (*FaceIDMessage) Descriptor() ([]byte, []int) {
	return file_face_proto_rawDescGZIP(), []int{3}
}

func (x *FaceIDMessage) GetFaceID() int32 {
	if x != nil {
		return x.FaceID
	}
	return 0
}

type FileIDMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileID int32 `protobuf:"varint,1,opt,name=fileID,proto3" json:"fileID,omitempty"`
}

func (x *FileIDMessage) Reset() {
	*x = FileIDMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_face_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileIDMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileIDMessage) ProtoMessage() {}

func (x *FileIDMessage) ProtoReflect() protoreflect.Message {
	mi := &file_face_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileIDMessage.ProtoReflect.Descriptor instead.
func (*FileIDMessage) Descriptor() ([]byte, []int) {
	return file_face_proto_rawDescGZIP(), []int{4}
}

func (x *FileIDMessage) GetFileID() int32 {
	if x != nil {
		return x.FileID
	}
	return 0
}

type RenameClusterMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterID     string `protobuf:"bytes,1,opt,name=clusterID,proto3" json:"clusterID,omitempty"`
	NewPersonName string `protobuf:"bytes,2,opt,name=newPersonName,proto3" json:"newPersonName,omitempty"`
}

func (x *RenameClusterMessage) Reset() {
	*x = RenameClusterMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_face_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenameClusterMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameClusterMessage) ProtoMessage() {}

func (x *RenameClusterMessage) ProtoReflect() protoreflect.Message {
	mi := &file_face_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameClusterMessage.ProtoReflect.Descriptor instead.
func (*RenameClusterMessage) Descriptor() ([]byte, []int) {
	return file_face_proto_rawDescGZIP(), []int{5}
}

func (x *RenameClusterMessage) GetClusterID() string {
	if x != nil {
		return x.ClusterID
	}
	return ""
}

func (x *RenameClusterMessage) GetNewPersonName() string {
	if x != nil {
		return x.NewPersonName
	}
	return ""
}

type MergeClustersMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterIDs []*ClusterIDMessage `protobuf:"bytes,1,rep,name=clusterIDs,proto3" json:"clusterIDs,omitempty"`
}

func (x *MergeClustersMessage) Reset() {
	*x = MergeClustersMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_face_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MergeClustersMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MergeClustersMessage) ProtoMessage() {}

func (x *MergeClustersMessage) ProtoReflect() protoreflect.Message {
	mi := &file_face_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MergeClustersMessage.ProtoReflect.Descriptor instead.
func (*MergeClustersMessage) Descriptor() ([]byte, []int) {
	return file_face_proto_rawDescGZIP(), []int{6}
}

func (x *MergeClustersMessage) GetClusterIDs() []*ClusterIDMessage {
	if x != nil {
		return x.ClusterIDs
	}
	return nil
}

type ManuallyMoveFacesToAnotherClusterMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FaceIDs   []*FaceIDMessage `protobuf:"bytes,1,rep,name=faceIDs,proto3" json:"faceIDs,omitempty"`
	ClusterID string           `protobuf:"bytes,2,opt,name=clusterID,proto3" json:"clusterID,omitempty"`
}

func (x *ManuallyMoveFacesToAnotherClusterMessage) Reset() {
	*x = ManuallyMoveFacesToAnotherClusterMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_face_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManuallyMoveFacesToAnotherClusterMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManuallyMoveFacesToAnotherClusterMessage) ProtoMessage() {}

func (x *ManuallyMoveFacesToAnotherClusterMessage) ProtoReflect() protoreflect.Message {
	mi := &file_face_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManuallyMoveFacesToAnotherClusterMessage.ProtoReflect.Descriptor instead.
func (*ManuallyMoveFacesToAnotherClusterMessage) Descriptor() ([]byte, []int) {
	return file_face_proto_rawDescGZIP(), []int{7}
}

func (x *ManuallyMoveFacesToAnotherClusterMessage) GetFaceIDs() []*FaceIDMessage {
	if x != nil {
		return x.FaceIDs
	}
	return nil
}

func (x *ManuallyMoveFacesToAnotherClusterMessage) GetClusterID() string {
	if x != nil {
		return x.ClusterID
	}
	return ""
}

type RemoveFacesFromDatabaseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FaceIDs []*FaceIDMessage `protobuf:"bytes,1,rep,name=faceIDs,proto3" json:"faceIDs,omitempty"`
}

func (x *RemoveFacesFromDatabaseMessage) Reset() {
	*x = RemoveFacesFromDatabaseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_face_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveFacesFromDatabaseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFacesFromDatabaseMessage) ProtoMessage() {}

func (x *RemoveFacesFromDatabaseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_face_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFacesFromDatabaseMessage.ProtoReflect.Descriptor instead.
func (*RemoveFacesFromDatabaseMessage) Descriptor() ([]byte, []int) {
	return file_face_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveFacesFromDatabaseMessage) GetFaceIDs() []*FaceIDMessage {
	if x != nil {
		return x.FaceIDs
	}
	return nil
}

type RemoveFilesFromDatabaseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileIDs []*FileIDMessage `protobuf:"bytes,1,rep,name=fileIDs,proto3" json:"fileIDs,omitempty"`
}

func (x *RemoveFilesFromDatabaseMessage) Reset() {
	*x = RemoveFilesFromDatabaseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_face_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveFilesFromDatabaseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFilesFromDatabaseMessage) ProtoMessage() {}

func (x *RemoveFilesFromDatabaseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_face_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFilesFromDatabaseMessage.ProtoReflect.Descriptor instead.
func (*RemoveFilesFromDatabaseMessage) Descriptor() ([]byte, []int) {
	return file_face_proto_rawDescGZIP(), []int{9}
}

func (x *RemoveFilesFromDatabaseMessage) GetFileIDs() []*FileIDMessage {
	if x != nil {
		return x.FileIDs
	}
	return nil
}

var File_face_proto protoreflect.FileDescriptor

var file_face_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0e, 0x0a, 0x0c,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a, 0x0a,
	0x45, 0x72, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x72,
	0x72, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65,
	0x72, 0x72, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x30, 0x0a, 0x10, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x49, 0x44, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x22, 0x27, 0x0a, 0x0d, 0x46, 0x61,
	0x63, 0x65, 0x49, 0x44, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x61, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x61, 0x63,
	0x65, 0x49, 0x44, 0x22, 0x27, 0x0a, 0x0d, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x22, 0x5a, 0x0a, 0x14,
	0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x65, 0x77, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x77, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x49, 0x0a, 0x14, 0x4d, 0x65, 0x72, 0x67,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x31, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x44, 0x73, 0x22, 0x72, 0x0a, 0x28, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x6c, 0x79, 0x4d,
	0x6f, 0x76, 0x65, 0x46, 0x61, 0x63, 0x65, 0x73, 0x54, 0x6f, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x28, 0x0a, 0x07, 0x66, 0x61, 0x63, 0x65, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x46, 0x61, 0x63, 0x65, 0x49, 0x44, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x07, 0x66, 0x61, 0x63, 0x65, 0x49, 0x44, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x22, 0x4a, 0x0a, 0x1e, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x46, 0x61, 0x63, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x66, 0x61, 0x63,
	0x65, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x46, 0x61, 0x63,
	0x65, 0x49, 0x44, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x66, 0x61, 0x63, 0x65,
	0x49, 0x44, 0x73, 0x22, 0x4a, 0x0a, 0x1e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x73, 0x32,
	0x47, 0x0a, 0x13, 0x52, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x10, 0x52, 0x65, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x75, 0x6e, 0x63, 0x12, 0x0d, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0b, 0x2e, 0x45, 0x72, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x32, 0x5b, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x46, 0x61, 0x63, 0x65, 0x73, 0x41, 0x6e, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x1a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x46, 0x61, 0x63, 0x65, 0x73, 0x41, 0x6e, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x46, 0x75, 0x6e, 0x63, 0x12, 0x0d, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0b, 0x2e, 0x45, 0x72, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x00, 0x32, 0x51, 0x0a, 0x14, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a,
	0x11, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x75,
	0x6e, 0x63, 0x12, 0x15, 0x2e, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0b, 0x2e, 0x45, 0x72, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x32, 0x57, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x46,
	0x75, 0x6e, 0x63, 0x12, 0x0d, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x0b, 0x2e, 0x45, 0x72, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x00, 0x32, 0x51, 0x0a, 0x14, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x11, 0x4d, 0x65, 0x72,
	0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x46, 0x75, 0x6e, 0x63, 0x12, 0x15,
	0x2e, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0b, 0x2e, 0x45, 0x72, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x00, 0x32, 0x8d, 0x01, 0x0a, 0x28, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x6c,
	0x79, 0x4d, 0x6f, 0x76, 0x65, 0x46, 0x61, 0x63, 0x65, 0x73, 0x54, 0x6f, 0x41, 0x6e, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x61, 0x0a, 0x25, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x6c, 0x79, 0x4d, 0x6f, 0x76,
	0x65, 0x46, 0x61, 0x63, 0x65, 0x73, 0x54, 0x6f, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x75, 0x6e, 0x63, 0x12, 0x29, 0x2e, 0x4d, 0x61, 0x6e,
	0x75, 0x61, 0x6c, 0x6c, 0x79, 0x4d, 0x6f, 0x76, 0x65, 0x46, 0x61, 0x63, 0x65, 0x73, 0x54, 0x6f,
	0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0b, 0x2e, 0x45, 0x72, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x00, 0x32, 0x6f, 0x0a, 0x1e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x61,
	0x63, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x1b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x46, 0x61, 0x63, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x46, 0x75, 0x6e, 0x63, 0x12, 0x1f, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x61,
	0x63, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0b, 0x2e, 0x45, 0x72, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x00, 0x32, 0x6f, 0x0a, 0x1e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x1b, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x46, 0x75, 0x6e, 0x63, 0x12, 0x1f, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0b, 0x2e, 0x45, 0x72, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x66,
	0x61, 0x63, 0x65, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_face_proto_rawDescOnce sync.Once
	file_face_proto_rawDescData = file_face_proto_rawDesc
)

func file_face_proto_rawDescGZIP() []byte {
	file_face_proto_rawDescOnce.Do(func() {
		file_face_proto_rawDescData = protoimpl.X.CompressGZIP(file_face_proto_rawDescData)
	})
	return file_face_proto_rawDescData
}

var file_face_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_face_proto_goTypes = []interface{}{
	(*EmptyMessage)(nil),                             // 0: EmptyMessage
	(*ErrMessage)(nil),                               // 1: ErrMessage
	(*ClusterIDMessage)(nil),                         // 2: ClusterIDMessage
	(*FaceIDMessage)(nil),                            // 3: FaceIDMessage
	(*FileIDMessage)(nil),                            // 4: FileIDMessage
	(*RenameClusterMessage)(nil),                     // 5: RenameClusterMessage
	(*MergeClustersMessage)(nil),                     // 6: MergeClustersMessage
	(*ManuallyMoveFacesToAnotherClusterMessage)(nil), // 7: ManuallyMoveFacesToAnotherClusterMessage
	(*RemoveFacesFromDatabaseMessage)(nil),           // 8: RemoveFacesFromDatabaseMessage
	(*RemoveFilesFromDatabaseMessage)(nil),           // 9: RemoveFilesFromDatabaseMessage
}
var file_face_proto_depIdxs = []int32{
	2,  // 0: MergeClustersMessage.clusterIDs:type_name -> ClusterIDMessage
	3,  // 1: ManuallyMoveFacesToAnotherClusterMessage.faceIDs:type_name -> FaceIDMessage
	3,  // 2: RemoveFacesFromDatabaseMessage.faceIDs:type_name -> FaceIDMessage
	4,  // 3: RemoveFilesFromDatabaseMessage.fileIDs:type_name -> FileIDMessage
	0,  // 4: ReclusteringService.ReclusteringFunc:input_type -> EmptyMessage
	0,  // 5: UpdateFacesAndClustersService.UpdateFacesAndClustersFunc:input_type -> EmptyMessage
	5,  // 6: RenameClusterService.RenameClusterFunc:input_type -> RenameClusterMessage
	0,  // 7: DeleteAllPersonNamesService.DeleteAllPersonNamesFunc:input_type -> EmptyMessage
	6,  // 8: MergeClustersService.MergeClustersFunc:input_type -> MergeClustersMessage
	7,  // 9: ManuallyMoveFacesToAnotherClusterService.ManuallyMoveFacesToAnotherClusterFunc:input_type -> ManuallyMoveFacesToAnotherClusterMessage
	8,  // 10: RemoveFacesFromDatabaseService.RemoveFacesFromDatabaseFunc:input_type -> RemoveFacesFromDatabaseMessage
	9,  // 11: RemoveFilesFromDatabaseService.RemoveFilesFromDatabaseFunc:input_type -> RemoveFilesFromDatabaseMessage
	1,  // 12: ReclusteringService.ReclusteringFunc:output_type -> ErrMessage
	1,  // 13: UpdateFacesAndClustersService.UpdateFacesAndClustersFunc:output_type -> ErrMessage
	1,  // 14: RenameClusterService.RenameClusterFunc:output_type -> ErrMessage
	1,  // 15: DeleteAllPersonNamesService.DeleteAllPersonNamesFunc:output_type -> ErrMessage
	1,  // 16: MergeClustersService.MergeClustersFunc:output_type -> ErrMessage
	1,  // 17: ManuallyMoveFacesToAnotherClusterService.ManuallyMoveFacesToAnotherClusterFunc:output_type -> ErrMessage
	1,  // 18: RemoveFacesFromDatabaseService.RemoveFacesFromDatabaseFunc:output_type -> ErrMessage
	1,  // 19: RemoveFilesFromDatabaseService.RemoveFilesFromDatabaseFunc:output_type -> ErrMessage
	12, // [12:20] is the sub-list for method output_type
	4,  // [4:12] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_face_proto_init() }
func file_face_proto_init() {
	if File_face_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_face_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMessage); i {
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
		file_face_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrMessage); i {
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
		file_face_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterIDMessage); i {
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
		file_face_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaceIDMessage); i {
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
		file_face_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileIDMessage); i {
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
		file_face_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenameClusterMessage); i {
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
		file_face_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MergeClustersMessage); i {
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
		file_face_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManuallyMoveFacesToAnotherClusterMessage); i {
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
		file_face_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveFacesFromDatabaseMessage); i {
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
		file_face_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveFilesFromDatabaseMessage); i {
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
			RawDescriptor: file_face_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   8,
		},
		GoTypes:           file_face_proto_goTypes,
		DependencyIndexes: file_face_proto_depIdxs,
		MessageInfos:      file_face_proto_msgTypes,
	}.Build()
	File_face_proto = out.File
	file_face_proto_rawDesc = nil
	file_face_proto_goTypes = nil
	file_face_proto_depIdxs = nil
}
