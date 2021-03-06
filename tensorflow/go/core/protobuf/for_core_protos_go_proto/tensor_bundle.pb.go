// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/tensor_bundle.proto

package for_core_protos_go_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	tensor_shape_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_shape_go_proto"
	tensor_slice_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_slice_go_proto"
	types_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/types_go_proto"
	versions_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/versions_go_proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// An enum indicating the endianness of the platform that produced this
// bundle.  A bundle can only be read by a platform with matching endianness.
// Defaults to LITTLE, as most modern platforms are little-endian.
//
// Affects the binary tensor data bytes only, not the metadata in protobufs.
type BundleHeaderProto_Endianness int32

const (
	BundleHeaderProto_LITTLE BundleHeaderProto_Endianness = 0
	BundleHeaderProto_BIG    BundleHeaderProto_Endianness = 1
)

var BundleHeaderProto_Endianness_name = map[int32]string{
	0: "LITTLE",
	1: "BIG",
}

var BundleHeaderProto_Endianness_value = map[string]int32{
	"LITTLE": 0,
	"BIG":    1,
}

func (x BundleHeaderProto_Endianness) String() string {
	return proto.EnumName(BundleHeaderProto_Endianness_name, int32(x))
}

func (BundleHeaderProto_Endianness) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9ab648e6509929dc, []int{0, 0}
}

// Special header that is associated with a bundle.
//
// TODO(zongheng,zhifengc): maybe in the future, we can add information about
// which binary produced this checkpoint, timestamp, etc. Sometime, these can be
// valuable debugging information. And if needed, these can be used as defensive
// information ensuring reader (binary version) of the checkpoint and the writer
// (binary version) must match within certain range, etc.
type BundleHeaderProto struct {
	// Number of data files in the bundle.
	NumShards  int32                        `protobuf:"varint,1,opt,name=num_shards,json=numShards,proto3" json:"num_shards,omitempty"`
	Endianness BundleHeaderProto_Endianness `protobuf:"varint,2,opt,name=endianness,proto3,enum=tensorflow.BundleHeaderProto_Endianness" json:"endianness,omitempty"`
	// Versioning of the tensor bundle format.
	Version              *versions_go_proto.VersionDef `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *BundleHeaderProto) Reset()         { *m = BundleHeaderProto{} }
func (m *BundleHeaderProto) String() string { return proto.CompactTextString(m) }
func (*BundleHeaderProto) ProtoMessage()    {}
func (*BundleHeaderProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ab648e6509929dc, []int{0}
}

func (m *BundleHeaderProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BundleHeaderProto.Unmarshal(m, b)
}
func (m *BundleHeaderProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BundleHeaderProto.Marshal(b, m, deterministic)
}
func (m *BundleHeaderProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BundleHeaderProto.Merge(m, src)
}
func (m *BundleHeaderProto) XXX_Size() int {
	return xxx_messageInfo_BundleHeaderProto.Size(m)
}
func (m *BundleHeaderProto) XXX_DiscardUnknown() {
	xxx_messageInfo_BundleHeaderProto.DiscardUnknown(m)
}

var xxx_messageInfo_BundleHeaderProto proto.InternalMessageInfo

func (m *BundleHeaderProto) GetNumShards() int32 {
	if m != nil {
		return m.NumShards
	}
	return 0
}

func (m *BundleHeaderProto) GetEndianness() BundleHeaderProto_Endianness {
	if m != nil {
		return m.Endianness
	}
	return BundleHeaderProto_LITTLE
}

func (m *BundleHeaderProto) GetVersion() *versions_go_proto.VersionDef {
	if m != nil {
		return m.Version
	}
	return nil
}

// Describes the metadata related to a checkpointed tensor.
type BundleEntryProto struct {
	// The tensor dtype and shape.
	Dtype types_go_proto.DataType                 `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Shape *tensor_shape_go_proto.TensorShapeProto `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	// The binary content of the tensor lies in:
	//   File "shard_id": bytes [offset, offset + size).
	ShardId int32 `protobuf:"varint,3,opt,name=shard_id,json=shardId,proto3" json:"shard_id,omitempty"`
	Offset  int64 `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Size    int64 `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	// The CRC32C checksum of the tensor bytes.
	Crc32C uint32 `protobuf:"fixed32,6,opt,name=crc32c,proto3" json:"crc32c,omitempty"`
	// Iff present, this entry represents a partitioned tensor.  The previous
	// fields are interpreted as follows:
	//
	//   "dtype", "shape": describe the full tensor.
	//   "shard_id", "offset", "size", "crc32c": all IGNORED.
	//      These information for each slice can be looked up in their own
	//      BundleEntryProto, keyed by each "slice_name".
	Slices               []*tensor_slice_go_proto.TensorSliceProto `protobuf:"bytes,7,rep,name=slices,proto3" json:"slices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *BundleEntryProto) Reset()         { *m = BundleEntryProto{} }
func (m *BundleEntryProto) String() string { return proto.CompactTextString(m) }
func (*BundleEntryProto) ProtoMessage()    {}
func (*BundleEntryProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ab648e6509929dc, []int{1}
}

func (m *BundleEntryProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BundleEntryProto.Unmarshal(m, b)
}
func (m *BundleEntryProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BundleEntryProto.Marshal(b, m, deterministic)
}
func (m *BundleEntryProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BundleEntryProto.Merge(m, src)
}
func (m *BundleEntryProto) XXX_Size() int {
	return xxx_messageInfo_BundleEntryProto.Size(m)
}
func (m *BundleEntryProto) XXX_DiscardUnknown() {
	xxx_messageInfo_BundleEntryProto.DiscardUnknown(m)
}

var xxx_messageInfo_BundleEntryProto proto.InternalMessageInfo

func (m *BundleEntryProto) GetDtype() types_go_proto.DataType {
	if m != nil {
		return m.Dtype
	}
	return types_go_proto.DataType_DT_INVALID
}

func (m *BundleEntryProto) GetShape() *tensor_shape_go_proto.TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *BundleEntryProto) GetShardId() int32 {
	if m != nil {
		return m.ShardId
	}
	return 0
}

func (m *BundleEntryProto) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *BundleEntryProto) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *BundleEntryProto) GetCrc32C() uint32 {
	if m != nil {
		return m.Crc32C
	}
	return 0
}

func (m *BundleEntryProto) GetSlices() []*tensor_slice_go_proto.TensorSliceProto {
	if m != nil {
		return m.Slices
	}
	return nil
}

func init() {
	proto.RegisterEnum("tensorflow.BundleHeaderProto_Endianness", BundleHeaderProto_Endianness_name, BundleHeaderProto_Endianness_value)
	proto.RegisterType((*BundleHeaderProto)(nil), "tensorflow.BundleHeaderProto")
	proto.RegisterType((*BundleEntryProto)(nil), "tensorflow.BundleEntryProto")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/tensor_bundle.proto", fileDescriptor_9ab648e6509929dc)
}

var fileDescriptor_9ab648e6509929dc = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x6b, 0xd4, 0x40,
	0x10, 0xc7, 0xdd, 0x5e, 0x93, 0xe8, 0x14, 0xca, 0xb9, 0x4a, 0x89, 0xa2, 0x10, 0x0f, 0x84, 0x20,
	0x92, 0x93, 0xd4, 0xbf, 0xe0, 0xe8, 0x61, 0x0f, 0xfa, 0x50, 0xb6, 0xd1, 0x07, 0x5f, 0x42, 0x7e,
	0x6c, 0xd2, 0x60, 0xb2, 0x1b, 0x76, 0x13, 0xcb, 0xf9, 0x2e, 0xfe, 0x7d, 0xfe, 0x37, 0x3e, 0x4a,
	0x66, 0x73, 0x5e, 0xb0, 0x56, 0xfa, 0x36, 0x3b, 0xf3, 0xf9, 0xce, 0xce, 0x77, 0xb2, 0x81, 0xb7,
	0x1d, 0x17, 0x5a, 0xaa, 0xa2, 0x96, 0x37, 0xcb, 0x4c, 0x2a, 0xbe, 0x6c, 0x95, 0xec, 0x64, 0xda,
	0x17, 0x4b, 0x53, 0x88, 0xd3, 0x5e, 0xe4, 0x35, 0x0f, 0x30, 0x4d, 0x61, 0x4f, 0x3f, 0xbf, 0xa5,
	0x2c, 0x54, 0xd2, 0xf0, 0x1b, 0xa9, 0xbe, 0xec, 0xa4, 0xfa, 0x3a, 0x69, 0x47, 0xe5, 0x7d, 0xe8,
	0xba, 0xca, 0x76, 0xf4, 0xeb, 0xff, 0xd0, 0xdb, 0x96, 0xeb, 0x11, 0xf3, 0xef, 0xc6, 0xbe, 0x72,
	0xa5, 0x2b, 0x29, 0x46, 0x72, 0xf1, 0x93, 0xc0, 0xe3, 0x15, 0x3a, 0x39, 0xe7, 0x49, 0xce, 0xd5,
	0x25, 0xda, 0x79, 0x09, 0x20, 0xfa, 0x66, 0x98, 0x53, 0xe5, 0xda, 0x25, 0x1e, 0xf1, 0x2d, 0xf6,
	0x48, 0xf4, 0xcd, 0x15, 0x26, 0xe8, 0x39, 0x00, 0x17, 0x79, 0x95, 0x08, 0xc1, 0xb5, 0x76, 0x0f,
	0x3c, 0xe2, 0x1f, 0x87, 0x7e, 0xb0, 0xbf, 0x33, 0xb8, 0xd5, 0x31, 0x58, 0xff, 0xe1, 0xd9, 0x44,
	0x4b, 0xdf, 0x81, 0x33, 0x0e, 0xe4, 0xce, 0x3c, 0xe2, 0x1f, 0x85, 0x27, 0xd3, 0x36, 0x9f, 0x4c,
	0xe9, 0x8c, 0x17, 0x6c, 0x87, 0x2d, 0x5e, 0x01, 0xec, 0x7b, 0x51, 0x00, 0xfb, 0x62, 0x13, 0x45,
	0x17, 0xeb, 0xf9, 0x03, 0xea, 0xc0, 0x6c, 0xb5, 0xf9, 0x30, 0x27, 0x8b, 0x1f, 0x07, 0x30, 0x37,
	0x13, 0xac, 0x45, 0xa7, 0xb6, 0xc6, 0xd2, 0x1b, 0xb0, 0xf2, 0x61, 0x45, 0xe8, 0xe6, 0x38, 0x7c,
	0x3a, 0xbd, 0xe7, 0x2c, 0xe9, 0x92, 0x68, 0xdb, 0x72, 0x66, 0x10, 0x1a, 0x82, 0x85, 0x9f, 0x08,
	0xad, 0x1d, 0x85, 0x2f, 0xa6, 0x6c, 0x84, 0xe1, 0xd5, 0x50, 0xc6, 0xc6, 0xcc, 0xa0, 0xf4, 0x19,
	0x3c, 0xc4, 0x75, 0xc5, 0x55, 0x8e, 0x56, 0x2c, 0xe6, 0xe0, 0x79, 0x93, 0xd3, 0x13, 0xb0, 0x65,
	0x51, 0x68, 0xde, 0xb9, 0x87, 0x1e, 0xf1, 0x67, 0x6c, 0x3c, 0x51, 0x0a, 0x87, 0xba, 0xfa, 0xc6,
	0x5d, 0x0b, 0xb3, 0x18, 0x0f, 0x6c, 0xa6, 0xb2, 0xd3, 0x30, 0x73, 0x6d, 0x8f, 0xf8, 0x0e, 0x1b,
	0x4f, 0xf4, 0x3d, 0xd8, 0xf8, 0x0e, 0xb4, 0xeb, 0x78, 0xb3, 0x3b, 0x66, 0x1a, 0xea, 0x66, 0xa6,
	0x91, 0x5d, 0x7d, 0x27, 0xf0, 0x44, 0xaa, 0x72, 0xca, 0xf6, 0x5d, 0x55, 0xaf, 0xa8, 0x51, 0x98,
	0x25, 0xa1, 0x44, 0x5f, 0x92, 0xcf, 0x1f, 0xcb, 0xaa, 0xbb, 0xee, 0xd3, 0x20, 0x93, 0xcd, 0x72,
	0xf2, 0x80, 0xfe, 0x1d, 0x96, 0xf2, 0xaf, 0xdf, 0xa2, 0x90, 0x2a, 0x1e, 0x32, 0x31, 0x66, 0x74,
	0x5c, 0x4a, 0x13, 0xfd, 0x22, 0x24, 0xb5, 0x31, 0x3a, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x97,
	0xe0, 0x3d, 0x88, 0x55, 0x03, 0x00, 0x00,
}
