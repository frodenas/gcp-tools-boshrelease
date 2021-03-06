// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/product_channel_exclusivity.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing the availability of a product offer.
type ProductChannelExclusivityEnum_ProductChannelExclusivity int32

const (
	// Not specified.
	ProductChannelExclusivityEnum_UNSPECIFIED ProductChannelExclusivityEnum_ProductChannelExclusivity = 0
	// Used for return value only. Represents value unknown in this version.
	ProductChannelExclusivityEnum_UNKNOWN ProductChannelExclusivityEnum_ProductChannelExclusivity = 1
	// The item is sold through one channel only, either local stores or online
	// as indicated by its ProductChannel.
	ProductChannelExclusivityEnum_SINGLE_CHANNEL ProductChannelExclusivityEnum_ProductChannelExclusivity = 2
	// The item is matched to its online or local stores counterpart, indicating
	// it is available for purchase in both ShoppingProductChannels.
	ProductChannelExclusivityEnum_MULTI_CHANNEL ProductChannelExclusivityEnum_ProductChannelExclusivity = 3
)

var ProductChannelExclusivityEnum_ProductChannelExclusivity_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "SINGLE_CHANNEL",
	3: "MULTI_CHANNEL",
}
var ProductChannelExclusivityEnum_ProductChannelExclusivity_value = map[string]int32{
	"UNSPECIFIED":    0,
	"UNKNOWN":        1,
	"SINGLE_CHANNEL": 2,
	"MULTI_CHANNEL":  3,
}

func (x ProductChannelExclusivityEnum_ProductChannelExclusivity) String() string {
	return proto.EnumName(ProductChannelExclusivityEnum_ProductChannelExclusivity_name, int32(x))
}
func (ProductChannelExclusivityEnum_ProductChannelExclusivity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_product_channel_exclusivity_8ac2b04352a1c000, []int{0, 0}
}

// Availability of a product offer.
type ProductChannelExclusivityEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductChannelExclusivityEnum) Reset()         { *m = ProductChannelExclusivityEnum{} }
func (m *ProductChannelExclusivityEnum) String() string { return proto.CompactTextString(m) }
func (*ProductChannelExclusivityEnum) ProtoMessage()    {}
func (*ProductChannelExclusivityEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_channel_exclusivity_8ac2b04352a1c000, []int{0}
}
func (m *ProductChannelExclusivityEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductChannelExclusivityEnum.Unmarshal(m, b)
}
func (m *ProductChannelExclusivityEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductChannelExclusivityEnum.Marshal(b, m, deterministic)
}
func (dst *ProductChannelExclusivityEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductChannelExclusivityEnum.Merge(dst, src)
}
func (m *ProductChannelExclusivityEnum) XXX_Size() int {
	return xxx_messageInfo_ProductChannelExclusivityEnum.Size(m)
}
func (m *ProductChannelExclusivityEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductChannelExclusivityEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ProductChannelExclusivityEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ProductChannelExclusivityEnum)(nil), "google.ads.googleads.v0.enums.ProductChannelExclusivityEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.ProductChannelExclusivityEnum_ProductChannelExclusivity", ProductChannelExclusivityEnum_ProductChannelExclusivity_name, ProductChannelExclusivityEnum_ProductChannelExclusivity_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/product_channel_exclusivity.proto", fileDescriptor_product_channel_exclusivity_8ac2b04352a1c000)
}

var fileDescriptor_product_channel_exclusivity_8ac2b04352a1c000 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4f, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0x82, 0xa2, 0xfc, 0x94, 0xd2, 0xe4, 0x92, 0xf8, 0xe4, 0x8c, 0xc4, 0xbc,
	0xbc, 0xd4, 0x9c, 0xf8, 0xd4, 0x8a, 0xe4, 0x9c, 0xd2, 0xe2, 0xcc, 0xb2, 0xcc, 0x92, 0x4a, 0xbd,
	0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x59, 0x88, 0x2e, 0xbd, 0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x01,
	0x7a, 0x65, 0x06, 0x7a, 0x60, 0x03, 0x94, 0x1a, 0x19, 0xb9, 0x64, 0x03, 0x20, 0x86, 0x38, 0x43,
	0xcc, 0x70, 0x45, 0x18, 0xe1, 0x9a, 0x57, 0x9a, 0xab, 0x94, 0xc0, 0x25, 0x89, 0x53, 0x81, 0x10,
	0x3f, 0x17, 0x77, 0xa8, 0x5f, 0x70, 0x80, 0xab, 0xb3, 0xa7, 0x9b, 0xa7, 0xab, 0x8b, 0x00, 0x83,
	0x10, 0x37, 0x17, 0x7b, 0xa8, 0x9f, 0xb7, 0x9f, 0x7f, 0xb8, 0x9f, 0x00, 0xa3, 0x90, 0x10, 0x17,
	0x5f, 0xb0, 0xa7, 0x9f, 0xbb, 0x8f, 0x6b, 0xbc, 0xb3, 0x87, 0xa3, 0x9f, 0x9f, 0xab, 0x8f, 0x00,
	0x93, 0x90, 0x20, 0x17, 0xaf, 0x6f, 0xa8, 0x4f, 0x88, 0x27, 0x5c, 0x88, 0xd9, 0xe9, 0x3c, 0x23,
	0x97, 0x62, 0x72, 0x7e, 0xae, 0x1e, 0x5e, 0x97, 0x3a, 0xc9, 0xe1, 0x74, 0x45, 0x00, 0xc8, 0xa3,
	0x01, 0x8c, 0x51, 0x4e, 0x50, 0x03, 0xd2, 0xf3, 0x73, 0x12, 0xf3, 0xd2, 0xf5, 0xf2, 0x8b, 0xd2,
	0xf5, 0xd3, 0x53, 0xf3, 0xc0, 0xc1, 0x00, 0x0b, 0xbb, 0x82, 0xcc, 0x62, 0x1c, 0x41, 0x69, 0x0d,
	0x26, 0x17, 0x31, 0x31, 0xbb, 0x3b, 0x3a, 0xae, 0x62, 0x92, 0x75, 0x87, 0x18, 0xe5, 0x98, 0x52,
	0xac, 0x07, 0x61, 0x82, 0x58, 0x61, 0x06, 0x7a, 0xa0, 0x20, 0x29, 0x3e, 0x05, 0x93, 0x8f, 0x71,
	0x4c, 0x29, 0x8e, 0x81, 0xcb, 0xc7, 0x84, 0x19, 0xc4, 0x80, 0xe5, 0x93, 0xd8, 0xc0, 0x96, 0x1a,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7d, 0xf1, 0xef, 0x78, 0xbe, 0x01, 0x00, 0x00,
}
