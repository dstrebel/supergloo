// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ratelimit.proto

package v1 // import "github.com/solo-io/supergloo/pkg/api/external/gloo/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
import _ "github.com/gogo/protobuf/gogoproto"
import ratelimit "github.com/solo-io/supergloo/pkg/api/external/gloo/v1/plugins/ratelimit"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import bytes "bytes"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Constraint struct {
	Key                  string               `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string               `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	RateLimit            *ratelimit.RateLimit `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit" json:"rate_limit,omitempty"`
	Constraints          []*Constraint        `protobuf:"bytes,4,rep,name=constraints" json:"constraints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Constraint) Reset()         { *m = Constraint{} }
func (m *Constraint) String() string { return proto.CompactTextString(m) }
func (*Constraint) ProtoMessage()    {}
func (*Constraint) Descriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_6f3d7c62e1963865, []int{0}
}
func (m *Constraint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Constraint.Unmarshal(m, b)
}
func (m *Constraint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Constraint.Marshal(b, m, deterministic)
}
func (dst *Constraint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Constraint.Merge(dst, src)
}
func (m *Constraint) XXX_Size() int {
	return xxx_messageInfo_Constraint.Size(m)
}
func (m *Constraint) XXX_DiscardUnknown() {
	xxx_messageInfo_Constraint.DiscardUnknown(m)
}

var xxx_messageInfo_Constraint proto.InternalMessageInfo

func (m *Constraint) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Constraint) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Constraint) GetRateLimit() *ratelimit.RateLimit {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

func (m *Constraint) GetConstraints() []*Constraint {
	if m != nil {
		return m.Constraints
	}
	return nil
}

//
// @solo-kit:xds-service=RateLimitDiscoveryService
// @solo-kit:resource.no_references
type RateLimitConfig struct {
	// @solo-kit:resource.name
	Domain               string        `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Constraints          []*Constraint `protobuf:"bytes,2,rep,name=constraints" json:"constraints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RateLimitConfig) Reset()         { *m = RateLimitConfig{} }
func (m *RateLimitConfig) String() string { return proto.CompactTextString(m) }
func (*RateLimitConfig) ProtoMessage()    {}
func (*RateLimitConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ratelimit_6f3d7c62e1963865, []int{1}
}
func (m *RateLimitConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitConfig.Unmarshal(m, b)
}
func (m *RateLimitConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitConfig.Marshal(b, m, deterministic)
}
func (dst *RateLimitConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitConfig.Merge(dst, src)
}
func (m *RateLimitConfig) XXX_Size() int {
	return xxx_messageInfo_RateLimitConfig.Size(m)
}
func (m *RateLimitConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitConfig proto.InternalMessageInfo

func (m *RateLimitConfig) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimitConfig) GetConstraints() []*Constraint {
	if m != nil {
		return m.Constraints
	}
	return nil
}

func init() {
	proto.RegisterType((*Constraint)(nil), "gloo.solo.io.Constraint")
	proto.RegisterType((*RateLimitConfig)(nil), "gloo.solo.io.RateLimitConfig")
}
func (this *Constraint) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Constraint)
	if !ok {
		that2, ok := that.(Constraint)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if !this.RateLimit.Equal(that1.RateLimit) {
		return false
	}
	if len(this.Constraints) != len(that1.Constraints) {
		return false
	}
	for i := range this.Constraints {
		if !this.Constraints[i].Equal(that1.Constraints[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RateLimitConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RateLimitConfig)
	if !ok {
		that2, ok := that.(RateLimitConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Domain != that1.Domain {
		return false
	}
	if len(this.Constraints) != len(that1.Constraints) {
		return false
	}
	for i := range this.Constraints {
		if !this.Constraints[i].Equal(that1.Constraints[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RateLimitDiscoveryService service

type RateLimitDiscoveryServiceClient interface {
	StreamRateLimitConfig(ctx context.Context, opts ...grpc.CallOption) (RateLimitDiscoveryService_StreamRateLimitConfigClient, error)
	IncrementalRateLimitConfig(ctx context.Context, opts ...grpc.CallOption) (RateLimitDiscoveryService_IncrementalRateLimitConfigClient, error)
	FetchRateLimitConfig(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error)
}

type rateLimitDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRateLimitDiscoveryServiceClient(cc *grpc.ClientConn) RateLimitDiscoveryServiceClient {
	return &rateLimitDiscoveryServiceClient{cc}
}

func (c *rateLimitDiscoveryServiceClient) StreamRateLimitConfig(ctx context.Context, opts ...grpc.CallOption) (RateLimitDiscoveryService_StreamRateLimitConfigClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RateLimitDiscoveryService_serviceDesc.Streams[0], "/gloo.solo.io.RateLimitDiscoveryService/StreamRateLimitConfig", opts...)
	if err != nil {
		return nil, err
	}
	x := &rateLimitDiscoveryServiceStreamRateLimitConfigClient{stream}
	return x, nil
}

type RateLimitDiscoveryService_StreamRateLimitConfigClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type rateLimitDiscoveryServiceStreamRateLimitConfigClient struct {
	grpc.ClientStream
}

func (x *rateLimitDiscoveryServiceStreamRateLimitConfigClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rateLimitDiscoveryServiceStreamRateLimitConfigClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rateLimitDiscoveryServiceClient) IncrementalRateLimitConfig(ctx context.Context, opts ...grpc.CallOption) (RateLimitDiscoveryService_IncrementalRateLimitConfigClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RateLimitDiscoveryService_serviceDesc.Streams[1], "/gloo.solo.io.RateLimitDiscoveryService/IncrementalRateLimitConfig", opts...)
	if err != nil {
		return nil, err
	}
	x := &rateLimitDiscoveryServiceIncrementalRateLimitConfigClient{stream}
	return x, nil
}

type RateLimitDiscoveryService_IncrementalRateLimitConfigClient interface {
	Send(*v2.IncrementalDiscoveryRequest) error
	Recv() (*v2.IncrementalDiscoveryResponse, error)
	grpc.ClientStream
}

type rateLimitDiscoveryServiceIncrementalRateLimitConfigClient struct {
	grpc.ClientStream
}

func (x *rateLimitDiscoveryServiceIncrementalRateLimitConfigClient) Send(m *v2.IncrementalDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rateLimitDiscoveryServiceIncrementalRateLimitConfigClient) Recv() (*v2.IncrementalDiscoveryResponse, error) {
	m := new(v2.IncrementalDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rateLimitDiscoveryServiceClient) FetchRateLimitConfig(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error) {
	out := new(v2.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/gloo.solo.io.RateLimitDiscoveryService/FetchRateLimitConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RateLimitDiscoveryService service

type RateLimitDiscoveryServiceServer interface {
	StreamRateLimitConfig(RateLimitDiscoveryService_StreamRateLimitConfigServer) error
	IncrementalRateLimitConfig(RateLimitDiscoveryService_IncrementalRateLimitConfigServer) error
	FetchRateLimitConfig(context.Context, *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error)
}

func RegisterRateLimitDiscoveryServiceServer(s *grpc.Server, srv RateLimitDiscoveryServiceServer) {
	s.RegisterService(&_RateLimitDiscoveryService_serviceDesc, srv)
}

func _RateLimitDiscoveryService_StreamRateLimitConfig_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RateLimitDiscoveryServiceServer).StreamRateLimitConfig(&rateLimitDiscoveryServiceStreamRateLimitConfigServer{stream})
}

type RateLimitDiscoveryService_StreamRateLimitConfigServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type rateLimitDiscoveryServiceStreamRateLimitConfigServer struct {
	grpc.ServerStream
}

func (x *rateLimitDiscoveryServiceStreamRateLimitConfigServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rateLimitDiscoveryServiceStreamRateLimitConfigServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RateLimitDiscoveryService_IncrementalRateLimitConfig_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RateLimitDiscoveryServiceServer).IncrementalRateLimitConfig(&rateLimitDiscoveryServiceIncrementalRateLimitConfigServer{stream})
}

type RateLimitDiscoveryService_IncrementalRateLimitConfigServer interface {
	Send(*v2.IncrementalDiscoveryResponse) error
	Recv() (*v2.IncrementalDiscoveryRequest, error)
	grpc.ServerStream
}

type rateLimitDiscoveryServiceIncrementalRateLimitConfigServer struct {
	grpc.ServerStream
}

func (x *rateLimitDiscoveryServiceIncrementalRateLimitConfigServer) Send(m *v2.IncrementalDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rateLimitDiscoveryServiceIncrementalRateLimitConfigServer) Recv() (*v2.IncrementalDiscoveryRequest, error) {
	m := new(v2.IncrementalDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RateLimitDiscoveryService_FetchRateLimitConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v2.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimitDiscoveryServiceServer).FetchRateLimitConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gloo.solo.io.RateLimitDiscoveryService/FetchRateLimitConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimitDiscoveryServiceServer).FetchRateLimitConfig(ctx, req.(*v2.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RateLimitDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gloo.solo.io.RateLimitDiscoveryService",
	HandlerType: (*RateLimitDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchRateLimitConfig",
			Handler:    _RateLimitDiscoveryService_FetchRateLimitConfig_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRateLimitConfig",
			Handler:       _RateLimitDiscoveryService_StreamRateLimitConfig_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "IncrementalRateLimitConfig",
			Handler:       _RateLimitDiscoveryService_IncrementalRateLimitConfig_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ratelimit.proto",
}

func init() { proto.RegisterFile("ratelimit.proto", fileDescriptor_ratelimit_6f3d7c62e1963865) }

var fileDescriptor_ratelimit_6f3d7c62e1963865 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcd, 0x6e, 0xd4, 0x30,
	0x10, 0xc7, 0xf1, 0x2e, 0x54, 0xaa, 0x17, 0xa9, 0x28, 0x5a, 0x50, 0x88, 0x4a, 0x59, 0xe5, 0x94,
	0x56, 0xc2, 0x86, 0x20, 0x2e, 0xcb, 0x8d, 0x22, 0x04, 0x12, 0xa7, 0xec, 0x8d, 0x03, 0xc8, 0x4d,
	0x07, 0xd7, 0xaa, 0xe3, 0x09, 0xb6, 0x13, 0xb1, 0x47, 0x78, 0x05, 0x5e, 0x82, 0x67, 0xe0, 0x35,
	0x78, 0x03, 0xc4, 0x53, 0x70, 0x42, 0xf1, 0x6e, 0xc3, 0x6e, 0xf8, 0x10, 0x48, 0xdc, 0xc6, 0xf3,
	0xf5, 0xfb, 0x4f, 0x32, 0x43, 0xf7, 0xac, 0xf0, 0xa0, 0x55, 0xa5, 0x3c, 0xab, 0x2d, 0x7a, 0x8c,
	0xae, 0x4a, 0x8d, 0xc8, 0x1c, 0x6a, 0x64, 0x0a, 0x93, 0x7d, 0x30, 0x2d, 0x2e, 0xb9, 0xa8, 0x15,
	0x6f, 0x73, 0x7e, 0xaa, 0x5c, 0x89, 0x2d, 0xd8, 0xe5, 0x2a, 0x37, 0xd9, 0x97, 0x88, 0x52, 0x43,
	0x08, 0x0b, 0x63, 0xd0, 0x0b, 0xaf, 0xd0, 0xb8, 0x75, 0x74, 0x21, 0x95, 0x3f, 0x6b, 0x4e, 0x58,
	0x89, 0x15, 0xef, 0xfa, 0xdd, 0x51, 0xc8, 0x5d, 0x53, 0x83, 0xed, 0x08, 0xa1, 0x06, 0xde, 0x7a,
	0xb0, 0x46, 0x68, 0x1e, 0x3c, 0xed, 0x3d, 0x5e, 0xeb, 0x46, 0x2a, 0xe3, 0x78, 0x2f, 0x8a, 0x0f,
	0xe4, 0x25, 0x53, 0x89, 0x12, 0x83, 0xc9, 0x3b, 0x6b, 0xe5, 0x4d, 0x3f, 0x11, 0x4a, 0x8f, 0xd1,
	0x38, 0x6f, 0x85, 0x32, 0x3e, 0xba, 0x46, 0xc7, 0xe7, 0xb0, 0x8c, 0xc9, 0x8c, 0x64, 0xbb, 0x45,
	0x67, 0x46, 0x53, 0x7a, 0xa5, 0x15, 0xba, 0x81, 0x78, 0x14, 0x7c, 0xab, 0x47, 0xf4, 0x94, 0xd2,
	0xae, 0xff, 0xab, 0x00, 0x88, 0xc7, 0x33, 0x92, 0x4d, 0xf2, 0x43, 0xb6, 0x81, 0x5c, 0xc9, 0x61,
	0x9b, 0x9f, 0x84, 0x15, 0xc2, 0xc3, 0xf3, 0x2e, 0x5c, 0xec, 0xda, 0x0b, 0x33, 0x9a, 0xd3, 0x49,
	0xd9, 0xf3, 0x5d, 0x7c, 0x79, 0x36, 0xce, 0x26, 0x79, 0xbc, 0x5d, 0xf8, 0x43, 0x60, 0xb1, 0x99,
	0x9c, 0x02, 0xdd, 0xeb, 0x7b, 0x1e, 0xa3, 0x79, 0xad, 0x64, 0x74, 0x83, 0xee, 0x9c, 0x62, 0x25,
	0x94, 0x59, 0xcf, 0xb0, 0x7e, 0x0d, 0x31, 0xa3, 0x7f, 0xc0, 0xe4, 0xdf, 0x46, 0xf4, 0x66, 0xcf,
	0x79, 0x7c, 0xf1, 0x27, 0x17, 0x60, 0x5b, 0x55, 0x42, 0xf4, 0x92, 0x5e, 0x5f, 0x78, 0x0b, 0xa2,
	0x1a, 0x4a, 0x39, 0x60, 0x61, 0x05, 0x98, 0xa8, 0x15, 0x6b, 0x73, 0xd6, 0x17, 0x16, 0xf0, 0xa6,
	0x01, 0xe7, 0x93, 0xdb, 0xbf, 0x8d, 0xbb, 0x1a, 0x8d, 0x83, 0xf4, 0x52, 0x46, 0xee, 0x92, 0x68,
	0x49, 0x93, 0x67, 0xa6, 0xb4, 0x50, 0x81, 0xf1, 0x42, 0x0f, 0x21, 0x87, 0xdb, 0x4d, 0x36, 0x32,
	0x7f, 0xe2, 0x1d, 0xfd, 0x4d, 0xea, 0x16, 0xfa, 0x1d, 0xa1, 0xd3, 0x27, 0xe0, 0xcb, 0xb3, 0xff,
	0x3e, 0x5a, 0xf6, 0xfe, 0xf3, 0xd7, 0x0f, 0xa3, 0x34, 0xbd, 0xb5, 0x75, 0x1c, 0xf3, 0x7e, 0x81,
	0xca, 0xc0, 0x99, 0x93, 0xa3, 0x47, 0x0f, 0x3f, 0x7e, 0x39, 0x20, 0x2f, 0x1e, 0xfc, 0xf1, 0x22,
	0xea, 0x73, 0xf9, 0xcb, 0xab, 0x38, 0xd9, 0x09, 0x4b, 0x7e, 0xff, 0x7b, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xaa, 0x1a, 0x82, 0x59, 0xac, 0x03, 0x00, 0x00,
}