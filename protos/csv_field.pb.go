// Code generated by protoc-gen-go. DO NOT EDIT.
// source: csv_field.proto

package protos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
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

// Message containing options related to CSV schema generation and management via Protobuf.
type CSVFieldOptions struct {
	// Flag to specify that a field should be marked as 'REQUIRED' when used to generate schema for CSV.
	Require bool `protobuf:"varint,1,opt,name=require,proto3" json:"require,omitempty"`
	// Optionally omit a field from CSV schema.
	Ignore bool `protobuf:"varint,2,opt,name=ignore,proto3" json:"ignore,omitempty"`
	// Optionally override whatever type is resolved by the schema generator.
	// This is useful, for example, to store epoch timestamps with the underlying 'TIMESTAMP' type,
	// when normally, they would be structured as 'INTEGER' fields.
	TypeOverride string `protobuf:"bytes,3,opt,name=type_override,json=typeOverride,proto3" json:"type_override,omitempty"`
	// Optionally set a format string for parsing the value, eg timestamp format
	Format string `protobuf:"bytes,4,opt,name=format,proto3" json:"format,omitempty"`
	// Optional setting for parsing int, float, and complex numbers
	BitSize int32 `protobuf:"varint,5,opt,name=bit_size,json=bitSize,proto3" json:"bit_size,omitempty"`
	// Optional setting for parsing int, float, and complex numbers
	Base int32 `protobuf:"varint,6,opt,name=base,proto3" json:"base,omitempty"`
	// Customize the name of the field in the CSV schema.
	// If this is not set we use the name of the field in the proto instead.
	Name                 string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CSVFieldOptions) Reset()         { *m = CSVFieldOptions{} }
func (m *CSVFieldOptions) String() string { return proto.CompactTextString(m) }
func (*CSVFieldOptions) ProtoMessage()    {}
func (*CSVFieldOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_161ab9af91c6dea4, []int{0}
}

func (m *CSVFieldOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CSVFieldOptions.Unmarshal(m, b)
}
func (m *CSVFieldOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CSVFieldOptions.Marshal(b, m, deterministic)
}
func (m *CSVFieldOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CSVFieldOptions.Merge(m, src)
}
func (m *CSVFieldOptions) XXX_Size() int {
	return xxx_messageInfo_CSVFieldOptions.Size(m)
}
func (m *CSVFieldOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_CSVFieldOptions.DiscardUnknown(m)
}

var xxx_messageInfo_CSVFieldOptions proto.InternalMessageInfo

func (m *CSVFieldOptions) GetRequire() bool {
	if m != nil {
		return m.Require
	}
	return false
}

func (m *CSVFieldOptions) GetIgnore() bool {
	if m != nil {
		return m.Ignore
	}
	return false
}

func (m *CSVFieldOptions) GetTypeOverride() string {
	if m != nil {
		return m.TypeOverride
	}
	return ""
}

func (m *CSVFieldOptions) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *CSVFieldOptions) GetBitSize() int32 {
	if m != nil {
		return m.BitSize
	}
	return 0
}

func (m *CSVFieldOptions) GetBase() int32 {
	if m != nil {
		return m.Base
	}
	return 0
}

func (m *CSVFieldOptions) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

var E_Csv = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.FieldOptions)(nil),
	ExtensionType: (*CSVFieldOptions)(nil),
	Field:         7777,
	Name:          "gen_bq_schema.csv",
	Tag:           "bytes,7777,opt,name=csv",
	Filename:      "csv_field.proto",
}

func init() {
	proto.RegisterType((*CSVFieldOptions)(nil), "gen_bq_schema.CSVFieldOptions")
	proto.RegisterExtension(E_Csv)
}

func init() { proto.RegisterFile("csv_field.proto", fileDescriptor_161ab9af91c6dea4) }

var fileDescriptor_161ab9af91c6dea4 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0x89, 0xfb, 0xd7, 0xe8, 0xb2, 0x90, 0x83, 0x44, 0x41, 0x29, 0x7a, 0x29, 0xe2, 0x66,
	0x41, 0x6f, 0xe2, 0x49, 0xc1, 0xeb, 0x4a, 0x17, 0x3c, 0x78, 0x29, 0x4d, 0x76, 0x9a, 0x0d, 0xb4,
	0x49, 0x37, 0x49, 0x0b, 0xee, 0xd7, 0xf3, 0x93, 0xf8, 0x4d, 0xa4, 0x69, 0x17, 0xd4, 0xd3, 0xcc,
	0xfb, 0x3d, 0xe6, 0xc1, 0x1b, 0x3c, 0x17, 0xae, 0x49, 0x73, 0x05, 0xc5, 0x86, 0x55, 0xd6, 0x78,
	0x43, 0x66, 0x12, 0x74, 0xca, 0x77, 0xa9, 0x13, 0x5b, 0x28, 0xb3, 0x8b, 0x48, 0x1a, 0x23, 0x0b,
	0x58, 0x06, 0x93, 0xd7, 0xf9, 0x72, 0x03, 0x4e, 0x58, 0x55, 0x79, 0x63, 0xbb, 0x83, 0xeb, 0x2f,
	0x84, 0xe7, 0x2f, 0xeb, 0xf7, 0xd7, 0x36, 0x63, 0x55, 0x79, 0x65, 0xb4, 0x23, 0x14, 0x4f, 0x2c,
	0xec, 0x6a, 0x65, 0x81, 0xa2, 0x08, 0xc5, 0xd3, 0xe4, 0x20, 0xc9, 0x19, 0x1e, 0x2b, 0xa9, 0x8d,
	0x05, 0x7a, 0x14, 0x8c, 0x5e, 0x91, 0x1b, 0x3c, 0xf3, 0x9f, 0x15, 0xa4, 0xa6, 0x01, 0x6b, 0xd5,
	0x06, 0xe8, 0x20, 0x42, 0xf1, 0x71, 0x72, 0xda, 0xc2, 0x55, 0xcf, 0xda, 0xe3, 0xdc, 0xd8, 0x32,
	0xf3, 0x74, 0x18, 0xdc, 0x5e, 0x91, 0x73, 0x3c, 0xe5, 0xca, 0xa7, 0x4e, 0xed, 0x81, 0x8e, 0x22,
	0x14, 0x8f, 0x92, 0x09, 0x57, 0x7e, 0xad, 0xf6, 0x40, 0x08, 0x1e, 0xf2, 0xcc, 0x01, 0x1d, 0x07,
	0x1c, 0xf6, 0x96, 0xe9, 0xac, 0x04, 0x3a, 0x09, 0x21, 0x61, 0x7f, 0x7c, 0xc3, 0x03, 0xe1, 0x1a,
	0x72, 0xc9, 0xba, 0xbe, 0xec, 0xd0, 0x97, 0xfd, 0xee, 0x45, 0xbf, 0x9f, 0x22, 0x14, 0x9f, 0xdc,
	0x5f, 0xb1, 0x3f, 0x4f, 0x62, 0xff, 0xea, 0x27, 0x6d, 0xd4, 0xf3, 0xdd, 0xc7, 0xad, 0x54, 0x7e,
	0x5b, 0x73, 0x26, 0x4c, 0xb9, 0x94, 0x3a, 0x93, 0x50, 0x74, 0x6f, 0x14, 0x0b, 0x09, 0x7a, 0x21,
	0xcd, 0x42, 0xb8, 0xa6, 0x23, 0x8e, 0x8f, 0xc3, 0x7c, 0xf8, 0x09, 0x00, 0x00, 0xff, 0xff, 0x6c,
	0xef, 0x95, 0xf1, 0x90, 0x01, 0x00, 0x00,
}
