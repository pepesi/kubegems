// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: kubegems/datas/v1/cluster.proto

package v1

import (
	_ "github.com/mwitkow/go-proto-validators"
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

type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubegems_datas_v1_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_kubegems_datas_v1_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_kubegems_datas_v1_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *Cluster) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Cluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Kubeconfig string `protobuf:"bytes,2,opt,name=kubeconfig,proto3" json:"kubeconfig,omitempty"`
}

func (x *CreateClusterRequest) Reset() {
	*x = CreateClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubegems_datas_v1_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClusterRequest) ProtoMessage() {}

func (x *CreateClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kubegems_datas_v1_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClusterRequest.ProtoReflect.Descriptor instead.
func (*CreateClusterRequest) Descriptor() ([]byte, []int) {
	return file_kubegems_datas_v1_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *CreateClusterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateClusterRequest) GetKubeconfig() string {
	if x != nil {
		return x.Kubeconfig
	}
	return ""
}

type CreateClusterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Succeed bool   `protobuf:"varint,3,opt,name=succeed,proto3" json:"succeed,omitempty"`
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateClusterResponse) Reset() {
	*x = CreateClusterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubegems_datas_v1_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClusterResponse) ProtoMessage() {}

func (x *CreateClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kubegems_datas_v1_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClusterResponse.ProtoReflect.Descriptor instead.
func (*CreateClusterResponse) Descriptor() ([]byte, []int) {
	return file_kubegems_datas_v1_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *CreateClusterResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateClusterResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateClusterResponse) GetSucceed() bool {
	if x != nil {
		return x.Succeed
	}
	return false
}

func (x *CreateClusterResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteClusterRequest) Reset() {
	*x = DeleteClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubegems_datas_v1_cluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClusterRequest) ProtoMessage() {}

func (x *DeleteClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kubegems_datas_v1_cluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClusterRequest.ProtoReflect.Descriptor instead.
func (*DeleteClusterRequest) Descriptor() ([]byte, []int) {
	return file_kubegems_datas_v1_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteClusterRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteClusterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteClusterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteClusterResponse) Reset() {
	*x = DeleteClusterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubegems_datas_v1_cluster_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClusterResponse) ProtoMessage() {}

func (x *DeleteClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kubegems_datas_v1_cluster_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClusterResponse.ProtoReflect.Descriptor instead.
func (*DeleteClusterResponse) Descriptor() ([]byte, []int) {
	return file_kubegems_datas_v1_cluster_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteClusterResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DeleteClusterResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_kubegems_datas_v1_cluster_proto protoreflect.FileDescriptor

var file_kubegems_datas_v1_cluster_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6b, 0x75, 0x62, 0x65, 0x67, 0x65, 0x6d, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x6b, 0x75, 0x62, 0x65, 0x67, 0x65, 0x6d, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x07,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x62, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x16, 0xe2, 0xdf, 0x1f, 0x12, 0x0a, 0x10, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d,
	0x39, 0x5d, 0x7b, 0x35, 0x2c, 0x33, 0x30, 0x7d, 0x24, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22,
	0x6f, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x3a, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x45, 0x0a, 0x15,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x6b, 0x75, 0x62, 0x65, 0x67, 0x65, 0x6d, 0x73, 0x2e,
	0x69, 0x6f, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x67, 0x65, 0x6d, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kubegems_datas_v1_cluster_proto_rawDescOnce sync.Once
	file_kubegems_datas_v1_cluster_proto_rawDescData = file_kubegems_datas_v1_cluster_proto_rawDesc
)

func file_kubegems_datas_v1_cluster_proto_rawDescGZIP() []byte {
	file_kubegems_datas_v1_cluster_proto_rawDescOnce.Do(func() {
		file_kubegems_datas_v1_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_kubegems_datas_v1_cluster_proto_rawDescData)
	})
	return file_kubegems_datas_v1_cluster_proto_rawDescData
}

var file_kubegems_datas_v1_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_kubegems_datas_v1_cluster_proto_goTypes = []interface{}{
	(*Cluster)(nil),               // 0: kubegems.datas.v1.Cluster
	(*CreateClusterRequest)(nil),  // 1: kubegems.datas.v1.CreateClusterRequest
	(*CreateClusterResponse)(nil), // 2: kubegems.datas.v1.CreateClusterResponse
	(*DeleteClusterRequest)(nil),  // 3: kubegems.datas.v1.DeleteClusterRequest
	(*DeleteClusterResponse)(nil), // 4: kubegems.datas.v1.DeleteClusterResponse
}
var file_kubegems_datas_v1_cluster_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kubegems_datas_v1_cluster_proto_init() }
func file_kubegems_datas_v1_cluster_proto_init() {
	if File_kubegems_datas_v1_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kubegems_datas_v1_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster); i {
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
		file_kubegems_datas_v1_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClusterRequest); i {
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
		file_kubegems_datas_v1_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClusterResponse); i {
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
		file_kubegems_datas_v1_cluster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClusterRequest); i {
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
		file_kubegems_datas_v1_cluster_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClusterResponse); i {
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
			RawDescriptor: file_kubegems_datas_v1_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kubegems_datas_v1_cluster_proto_goTypes,
		DependencyIndexes: file_kubegems_datas_v1_cluster_proto_depIdxs,
		MessageInfos:      file_kubegems_datas_v1_cluster_proto_msgTypes,
	}.Build()
	File_kubegems_datas_v1_cluster_proto = out.File
	file_kubegems_datas_v1_cluster_proto_rawDesc = nil
	file_kubegems_datas_v1_cluster_proto_goTypes = nil
	file_kubegems_datas_v1_cluster_proto_depIdxs = nil
}
