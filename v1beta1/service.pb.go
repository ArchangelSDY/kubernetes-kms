// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1beta1/service.proto

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	v1beta1/service.proto

It has these top-level messages:
	VersionRequest
	VersionResponse
	DecryptRequest
	DecryptResponse
	EncryptRequest
	EncryptResponse
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type VersionRequest struct {
	// Version of the KMS plugin API.
	Version string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
}

func (m *VersionRequest) Reset()                    { *m = VersionRequest{} }
func (m *VersionRequest) String() string            { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()               {}
func (*VersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *VersionRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type VersionResponse struct {
	// Version of the KMS plugin API.
	Version string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	// Name of the KMS provider.
	RuntimeName string `protobuf:"bytes,2,opt,name=runtime_name,json=runtimeName" json:"runtime_name,omitempty"`
	// Version of the KMS provider. The string must be semver-compatible.
	RuntimeVersion string `protobuf:"bytes,3,opt,name=runtime_version,json=runtimeVersion" json:"runtime_version,omitempty"`
}

func (m *VersionResponse) Reset()                    { *m = VersionResponse{} }
func (m *VersionResponse) String() string            { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()               {}
func (*VersionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *VersionResponse) GetRuntimeName() string {
	if m != nil {
		return m.RuntimeName
	}
	return ""
}

func (m *VersionResponse) GetRuntimeVersion() string {
	if m != nil {
		return m.RuntimeVersion
	}
	return ""
}

type DecryptRequest struct {
	// Version of the KMS plugin API.
	Version string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	// The data to be decrypted.
	Cipher []byte `protobuf:"bytes,2,opt,name=cipher,proto3" json:"cipher,omitempty"`
}

func (m *DecryptRequest) Reset()                    { *m = DecryptRequest{} }
func (m *DecryptRequest) String() string            { return proto.CompactTextString(m) }
func (*DecryptRequest) ProtoMessage()               {}
func (*DecryptRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DecryptRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *DecryptRequest) GetCipher() []byte {
	if m != nil {
		return m.Cipher
	}
	return nil
}

type DecryptResponse struct {
	// The decrypted data.
	Plain []byte `protobuf:"bytes,1,opt,name=plain,proto3" json:"plain,omitempty"`
}

func (m *DecryptResponse) Reset()                    { *m = DecryptResponse{} }
func (m *DecryptResponse) String() string            { return proto.CompactTextString(m) }
func (*DecryptResponse) ProtoMessage()               {}
func (*DecryptResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DecryptResponse) GetPlain() []byte {
	if m != nil {
		return m.Plain
	}
	return nil
}

type EncryptRequest struct {
	// Version of the KMS plugin API.
	Version string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	// The data to be encrypted.
	Plain []byte `protobuf:"bytes,2,opt,name=plain,proto3" json:"plain,omitempty"`
}

func (m *EncryptRequest) Reset()                    { *m = EncryptRequest{} }
func (m *EncryptRequest) String() string            { return proto.CompactTextString(m) }
func (*EncryptRequest) ProtoMessage()               {}
func (*EncryptRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *EncryptRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *EncryptRequest) GetPlain() []byte {
	if m != nil {
		return m.Plain
	}
	return nil
}

type EncryptResponse struct {
	// The encrypted data.
	Cipher []byte `protobuf:"bytes,1,opt,name=cipher,proto3" json:"cipher,omitempty"`
}

func (m *EncryptResponse) Reset()                    { *m = EncryptResponse{} }
func (m *EncryptResponse) String() string            { return proto.CompactTextString(m) }
func (*EncryptResponse) ProtoMessage()               {}
func (*EncryptResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *EncryptResponse) GetCipher() []byte {
	if m != nil {
		return m.Cipher
	}
	return nil
}

func init() {
	proto.RegisterType((*VersionRequest)(nil), "v1beta1.VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "v1beta1.VersionResponse")
	proto.RegisterType((*DecryptRequest)(nil), "v1beta1.DecryptRequest")
	proto.RegisterType((*DecryptResponse)(nil), "v1beta1.DecryptResponse")
	proto.RegisterType((*EncryptRequest)(nil), "v1beta1.EncryptRequest")
	proto.RegisterType((*EncryptResponse)(nil), "v1beta1.EncryptResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for KMSService service

type KMSServiceClient interface {
	// Version returns the runtime name and runtime version of the KMS provider.
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	// Execute decryption operation in KMS provider.
	Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error)
	// Execute encryption operation in KMS provider.
	Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error)
}

type kMSServiceClient struct {
	cc *grpc.ClientConn
}

func NewKMSServiceClient(cc *grpc.ClientConn) KMSServiceClient {
	return &kMSServiceClient{cc}
}

func (c *kMSServiceClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := grpc.Invoke(ctx, "/v1beta1.KMSService/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kMSServiceClient) Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error) {
	out := new(DecryptResponse)
	err := grpc.Invoke(ctx, "/v1beta1.KMSService/Decrypt", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kMSServiceClient) Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error) {
	out := new(EncryptResponse)
	err := grpc.Invoke(ctx, "/v1beta1.KMSService/Encrypt", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KMSService service

type KMSServiceServer interface {
	// Version returns the runtime name and runtime version of the KMS provider.
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
	// Execute decryption operation in KMS provider.
	Decrypt(context.Context, *DecryptRequest) (*DecryptResponse, error)
	// Execute encryption operation in KMS provider.
	Encrypt(context.Context, *EncryptRequest) (*EncryptResponse, error)
}

func RegisterKMSServiceServer(s *grpc.Server, srv KMSServiceServer) {
	s.RegisterService(&_KMSService_serviceDesc, srv)
}

func _KMSService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KMSServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1beta1.KMSService/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KMSServiceServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KMSService_Decrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KMSServiceServer).Decrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1beta1.KMSService/Decrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KMSServiceServer).Decrypt(ctx, req.(*DecryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KMSService_Encrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KMSServiceServer).Encrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1beta1.KMSService/Encrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KMSServiceServer).Encrypt(ctx, req.(*EncryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KMSService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1beta1.KMSService",
	HandlerType: (*KMSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _KMSService_Version_Handler,
		},
		{
			MethodName: "Decrypt",
			Handler:    _KMSService_Decrypt_Handler,
		},
		{
			MethodName: "Encrypt",
			Handler:    _KMSService_Encrypt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1beta1/service.proto",
}

func init() { proto.RegisterFile("v1beta1/service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x4a, 0xc4, 0x30,
	0x10, 0xde, 0xae, 0xb8, 0xc5, 0xb1, 0xb4, 0x10, 0xfc, 0x29, 0x9e, 0x34, 0x97, 0x55, 0x0f, 0x95,
	0xd5, 0xbb, 0x88, 0xe8, 0x49, 0xf4, 0xd0, 0x05, 0xaf, 0xd2, 0x2d, 0x03, 0x06, 0x6c, 0x1a, 0x93,
	0x6c, 0xc5, 0x77, 0xf4, 0xa1, 0xc4, 0x66, 0x5a, 0xd3, 0x8a, 0xe8, 0x71, 0x26, 0xdf, 0xdf, 0xcc,
	0x04, 0x76, 0x9b, 0xc5, 0x0a, 0x6d, 0xb1, 0x38, 0x33, 0xa8, 0x1b, 0x51, 0x62, 0xa6, 0x74, 0x6d,
	0x6b, 0x16, 0x52, 0x9b, 0x9f, 0x42, 0xfc, 0x88, 0xda, 0x88, 0x5a, 0xe6, 0xf8, 0xba, 0x46, 0x63,
	0x59, 0x0a, 0x61, 0xe3, 0x3a, 0x69, 0x70, 0x18, 0x1c, 0x6f, 0xe5, 0x5d, 0xc9, 0xdf, 0x20, 0xe9,
	0xb1, 0x46, 0xd5, 0xd2, 0xe0, 0xef, 0x60, 0x76, 0x04, 0x91, 0x5e, 0x4b, 0x2b, 0x2a, 0x7c, 0x92,
	0x45, 0x85, 0xe9, 0xb4, 0x7d, 0xde, 0xa6, 0xde, 0x43, 0x51, 0x21, 0x9b, 0x43, 0xd2, 0x41, 0x3a,
	0x91, 0x8d, 0x16, 0x15, 0x53, 0x9b, 0xdc, 0xf8, 0x35, 0xc4, 0x37, 0x58, 0xea, 0x77, 0x65, 0xff,
	0x0c, 0xc9, 0xf6, 0x60, 0x56, 0x0a, 0xf5, 0x8c, 0xba, 0x75, 0x8c, 0x72, 0xaa, 0xf8, 0x1c, 0x92,
	0x5e, 0x83, 0xc2, 0xef, 0xc0, 0xa6, 0x7a, 0x29, 0x84, 0x93, 0x88, 0x72, 0x57, 0xf0, 0x2b, 0x88,
	0x6f, 0xe5, 0x3f, 0xcd, 0x7a, 0x85, 0xa9, 0xaf, 0x70, 0x02, 0x49, 0xaf, 0x40, 0x56, 0xdf, 0xa9,
	0x02, 0x3f, 0xd5, 0xf9, 0x47, 0x00, 0x70, 0x77, 0xbf, 0x5c, 0xba, 0xe3, 0xb0, 0x4b, 0x08, 0x69,
	0x66, 0xb6, 0x9f, 0xd1, 0x89, 0xb2, 0xe1, 0x7d, 0x0e, 0xd2, 0x9f, 0x0f, 0xce, 0x84, 0x4f, 0xbe,
	0xf8, 0x34, 0xa4, 0xc7, 0x1f, 0xae, 0xce, 0xe3, 0x8f, 0xf6, 0xe1, 0xf8, 0x94, 0xdc, 0xe3, 0x0f,
	0xb7, 0xe1, 0xf1, 0x47, 0x43, 0xf2, 0xc9, 0x6a, 0xd6, 0xfe, 0xae, 0x8b, 0xcf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x7e, 0xe1, 0x11, 0x7b, 0x76, 0x02, 0x00, 0x00,
}
