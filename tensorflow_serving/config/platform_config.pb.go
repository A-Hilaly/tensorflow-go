// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow_serving/config/platform_config.proto

package tensorflow_serving

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Configuration for a servable platform e.g. tensorflow or other ML systems.
type PlatformConfig struct {
	// The config proto for a SourceAdapter in the StoragePathSourceAdapter
	// registry.
	SourceAdapterConfig  *any.Any `protobuf:"bytes,1,opt,name=source_adapter_config,json=sourceAdapterConfig,proto3" json:"source_adapter_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlatformConfig) Reset()         { *m = PlatformConfig{} }
func (m *PlatformConfig) String() string { return proto.CompactTextString(m) }
func (*PlatformConfig) ProtoMessage()    {}
func (*PlatformConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b3919c9a7a426ae, []int{0}
}

func (m *PlatformConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformConfig.Unmarshal(m, b)
}
func (m *PlatformConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformConfig.Marshal(b, m, deterministic)
}
func (m *PlatformConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformConfig.Merge(m, src)
}
func (m *PlatformConfig) XXX_Size() int {
	return xxx_messageInfo_PlatformConfig.Size(m)
}
func (m *PlatformConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformConfig proto.InternalMessageInfo

func (m *PlatformConfig) GetSourceAdapterConfig() *any.Any {
	if m != nil {
		return m.SourceAdapterConfig
	}
	return nil
}

type PlatformConfigMap struct {
	// A map from a platform name to a platform config. The platform name is used
	// in ModelConfig.model_platform.
	PlatformConfigs      map[string]*PlatformConfig `protobuf:"bytes,1,rep,name=platform_configs,json=platformConfigs,proto3" json:"platform_configs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *PlatformConfigMap) Reset()         { *m = PlatformConfigMap{} }
func (m *PlatformConfigMap) String() string { return proto.CompactTextString(m) }
func (*PlatformConfigMap) ProtoMessage()    {}
func (*PlatformConfigMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b3919c9a7a426ae, []int{1}
}

func (m *PlatformConfigMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformConfigMap.Unmarshal(m, b)
}
func (m *PlatformConfigMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformConfigMap.Marshal(b, m, deterministic)
}
func (m *PlatformConfigMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformConfigMap.Merge(m, src)
}
func (m *PlatformConfigMap) XXX_Size() int {
	return xxx_messageInfo_PlatformConfigMap.Size(m)
}
func (m *PlatformConfigMap) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformConfigMap.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformConfigMap proto.InternalMessageInfo

func (m *PlatformConfigMap) GetPlatformConfigs() map[string]*PlatformConfig {
	if m != nil {
		return m.PlatformConfigs
	}
	return nil
}

func init() {
	proto.RegisterType((*PlatformConfig)(nil), "tensorflow.serving.PlatformConfig")
	proto.RegisterType((*PlatformConfigMap)(nil), "tensorflow.serving.PlatformConfigMap")
	proto.RegisterMapType((map[string]*PlatformConfig)(nil), "tensorflow.serving.PlatformConfigMap.PlatformConfigsEntry")
}

func init() {
	proto.RegisterFile("tensorflow_serving/config/platform_config.proto", fileDescriptor_7b3919c9a7a426ae)
}

var fileDescriptor_7b3919c9a7a426ae = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x49, 0x17, 0x05, 0x53, 0xd0, 0x1a, 0x2b, 0xac, 0x3d, 0x95, 0x3d, 0xf5, 0x94, 0x40,
	0xbd, 0x94, 0xde, 0xaa, 0x08, 0x5e, 0x04, 0xd9, 0xa3, 0x97, 0x25, 0xad, 0x93, 0xa5, 0xb8, 0x66,
	0x42, 0x92, 0xad, 0xe4, 0x75, 0x7d, 0x0a, 0x8f, 0x62, 0xb2, 0x22, 0xed, 0x0a, 0xbd, 0x25, 0xc3,
	0xff, 0x7d, 0xfc, 0x33, 0x54, 0x78, 0xd0, 0x0e, 0xad, 0x6a, 0xf0, 0xa3, 0x72, 0x60, 0x77, 0x5b,
	0x5d, 0x8b, 0x0d, 0x6a, 0xb5, 0xad, 0x85, 0x69, 0xa4, 0x57, 0x68, 0xdf, 0xab, 0xf4, 0xe7, 0xc6,
	0xa2, 0x47, 0xc6, 0xfe, 0x00, 0xde, 0x01, 0x93, 0x9b, 0x1a, 0xb1, 0x6e, 0x40, 0xc4, 0xc4, 0xba,
	0x55, 0x42, 0xea, 0x90, 0xe2, 0xc5, 0x0b, 0x3d, 0x7f, 0xee, 0x3c, 0xf7, 0x51, 0xc3, 0x1e, 0xe9,
	0xb5, 0xc3, 0xd6, 0x6e, 0xa0, 0x92, 0xaf, 0xd2, 0x78, 0xb0, 0x9d, 0x3f, 0x27, 0x53, 0x32, 0x1b,
	0xce, 0xc7, 0x3c, 0xc9, 0xf8, 0xaf, 0x8c, 0xaf, 0x74, 0x28, 0xaf, 0x12, 0xb2, 0x4a, 0x44, 0x32,
	0x15, 0x9f, 0x84, 0x5e, 0xee, 0xcb, 0x9f, 0xa4, 0x61, 0x40, 0x47, 0x07, 0xcd, 0x5d, 0x4e, 0xa6,
	0xd9, 0x6c, 0x38, 0x5f, 0xf2, 0x7e, 0x77, 0xde, 0x13, 0x1c, 0x4c, 0xdc, 0x83, 0xf6, 0x36, 0x94,
	0x17, 0x66, 0x7f, 0x3a, 0x51, 0x74, 0xfc, 0x5f, 0x90, 0x8d, 0x68, 0xf6, 0x06, 0x21, 0x2e, 0x73,
	0x56, 0xfe, 0x3c, 0xd9, 0x82, 0x9e, 0xec, 0x64, 0xd3, 0x42, 0x3e, 0x88, 0x0b, 0x16, 0xc7, 0x5b,
	0x94, 0x09, 0x58, 0x0e, 0x16, 0xe4, 0x2e, 0xfb, 0x22, 0x64, 0x7d, 0x1a, 0x8f, 0x71, 0xfb, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x50, 0x69, 0x59, 0xfb, 0xae, 0x01, 0x00, 0x00,
}
