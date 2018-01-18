// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config.proto

/*
Package config is a generated protocol buffer package.

It is generated from these files:
	config.proto

It has these top-level messages:
	Global
	Local
	Ack
	Empty
*/
package config

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

type Global struct {
	A     string   `protobuf:"bytes,1,opt,name=a" json:"a,omitempty"`
	B     string   `protobuf:"bytes,2,opt,name=b" json:"b,omitempty"`
	C     bool     `protobuf:"varint,3,opt,name=c" json:"c,omitempty"`
	Local []*Local `protobuf:"bytes,4,rep,name=local" json:"local,omitempty"`
}

func (m *Global) Reset()                    { *m = Global{} }
func (m *Global) String() string            { return proto.CompactTextString(m) }
func (*Global) ProtoMessage()               {}
func (*Global) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Global) GetA() string {
	if m != nil {
		return m.A
	}
	return ""
}

func (m *Global) GetB() string {
	if m != nil {
		return m.B
	}
	return ""
}

func (m *Global) GetC() bool {
	if m != nil {
		return m.C
	}
	return false
}

func (m *Global) GetLocal() []*Local {
	if m != nil {
		return m.Local
	}
	return nil
}

type Local struct {
	D string `protobuf:"bytes,1,opt,name=d" json:"d,omitempty"`
	E int32  `protobuf:"varint,2,opt,name=e" json:"e,omitempty"`
}

func (m *Local) Reset()                    { *m = Local{} }
func (m *Local) String() string            { return proto.CompactTextString(m) }
func (*Local) ProtoMessage()               {}
func (*Local) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Local) GetD() string {
	if m != nil {
		return m.D
	}
	return ""
}

func (m *Local) GetE() int32 {
	if m != nil {
		return m.E
	}
	return 0
}

type Ack struct {
	N int64 `protobuf:"varint,1,opt,name=n" json:"n,omitempty"`
}

func (m *Ack) Reset()                    { *m = Ack{} }
func (m *Ack) String() string            { return proto.CompactTextString(m) }
func (*Ack) ProtoMessage()               {}
func (*Ack) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ack) GetN() int64 {
	if m != nil {
		return m.N
	}
	return 0
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Global)(nil), "config.Global")
	proto.RegisterType((*Local)(nil), "config.Local")
	proto.RegisterType((*Ack)(nil), "config.Ack")
	proto.RegisterType((*Empty)(nil), "config.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Messenger service

type MessengerClient interface {
	Get(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Global, error)
	Set(ctx context.Context, in *Global, opts ...grpc.CallOption) (*Ack, error)
}

type messengerClient struct {
	cc *grpc.ClientConn
}

func NewMessengerClient(cc *grpc.ClientConn) MessengerClient {
	return &messengerClient{cc}
}

func (c *messengerClient) Get(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Global, error) {
	out := new(Global)
	err := grpc.Invoke(ctx, "/config.Messenger/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messengerClient) Set(ctx context.Context, in *Global, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := grpc.Invoke(ctx, "/config.Messenger/Set", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Messenger service

type MessengerServer interface {
	Get(context.Context, *Empty) (*Global, error)
	Set(context.Context, *Global) (*Ack, error)
}

func RegisterMessengerServer(s *grpc.Server, srv MessengerServer) {
	s.RegisterService(&_Messenger_serviceDesc, srv)
}

func _Messenger_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.Messenger/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServer).Get(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messenger_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Global)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.Messenger/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServer).Set(ctx, req.(*Global))
	}
	return interceptor(ctx, in, info, handler)
}

var _Messenger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "config.Messenger",
	HandlerType: (*MessengerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Messenger_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _Messenger_Set_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "config.proto",
}

func init() { proto.RegisterFile("config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xd0, 0x31, 0x4f, 0xc4, 0x20,
	0x14, 0x07, 0xf0, 0x7b, 0x22, 0xd5, 0x7b, 0x77, 0x3a, 0xe0, 0x42, 0x6e, 0x6a, 0x38, 0x63, 0x3a,
	0xdd, 0x70, 0x7e, 0x82, 0x0e, 0xa6, 0x8b, 0x2e, 0x38, 0x18, 0x47, 0xa0, 0xd8, 0x98, 0x22, 0x34,
	0x2d, 0x8b, 0xdf, 0xde, 0x00, 0xd6, 0xc1, 0xf1, 0xf7, 0xde, 0x9f, 0x3f, 0x01, 0xdc, 0x9b, 0xe0,
	0x3f, 0x3e, 0x87, 0xd3, 0x34, 0x87, 0x18, 0x58, 0x55, 0x24, 0xde, 0xb0, 0xea, 0x5c, 0xd0, 0xca,
	0xb1, 0x3d, 0x82, 0xe2, 0x50, 0x43, 0xb3, 0x95, 0xa0, 0x92, 0x34, 0xbf, 0x28, 0xd2, 0x49, 0x86,
	0x93, 0x1a, 0x9a, 0x6b, 0x09, 0x86, 0x1d, 0x91, 0xba, 0x60, 0x94, 0xe3, 0x97, 0x35, 0x69, 0x76,
	0xe7, 0x9b, 0xd3, 0x6f, 0xf3, 0x73, 0x1a, 0xca, 0xb2, 0x13, 0x47, 0xa4, 0xd9, 0xe9, 0x6c, 0xbf,
	0xf6, 0xf6, 0x49, 0x36, 0xf7, 0x52, 0x09, 0x56, 0xdc, 0x21, 0x69, 0xcd, 0x98, 0x86, 0x3e, 0x47,
	0x88, 0x04, 0x2f, 0xae, 0x90, 0x3e, 0x7d, 0x4d, 0xf1, 0xfb, 0xfc, 0x8e, 0xdb, 0x17, 0xbb, 0x2c,
	0xd6, 0x0f, 0x76, 0x66, 0x0f, 0x48, 0x3a, 0x1b, 0xd9, 0xdf, 0x65, 0x39, 0x72, 0xb8, 0x5d, 0x59,
	0x1e, 0x21, 0x36, 0xec, 0x1e, 0xc9, 0xab, 0x8d, 0xec, 0xdf, 0xe2, 0xb0, 0x5b, 0xdd, 0x9a, 0x51,
	0x6c, 0x74, 0x95, 0x7f, 0xe1, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0x49, 0x79, 0x3f, 0x49, 0x15,
	0x01, 0x00, 0x00,
}