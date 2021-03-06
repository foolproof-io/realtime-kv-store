// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kvstore.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	kvstore.proto

It has these top-level messages:
	GetValueRequest
	GetValueReply
	SetValueRequest
	SetValueReply
	ListenRequest
	ListenReply
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type GetValueRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *GetValueRequest) Reset()                    { *m = GetValueRequest{} }
func (m *GetValueRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetValueRequest) ProtoMessage()               {}
func (*GetValueRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetValueRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetValueReply struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *GetValueReply) Reset()                    { *m = GetValueReply{} }
func (m *GetValueReply) String() string            { return proto1.CompactTextString(m) }
func (*GetValueReply) ProtoMessage()               {}
func (*GetValueReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetValueReply) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SetValueRequest struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *SetValueRequest) Reset()                    { *m = SetValueRequest{} }
func (m *SetValueRequest) String() string            { return proto1.CompactTextString(m) }
func (*SetValueRequest) ProtoMessage()               {}
func (*SetValueRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SetValueRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetValueRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SetValueReply struct {
}

func (m *SetValueReply) Reset()                    { *m = SetValueReply{} }
func (m *SetValueReply) String() string            { return proto1.CompactTextString(m) }
func (*SetValueReply) ProtoMessage()               {}
func (*SetValueReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ListenRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *ListenRequest) Reset()                    { *m = ListenRequest{} }
func (m *ListenRequest) String() string            { return proto1.CompactTextString(m) }
func (*ListenRequest) ProtoMessage()               {}
func (*ListenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListenRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type ListenReply struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *ListenReply) Reset()                    { *m = ListenReply{} }
func (m *ListenReply) String() string            { return proto1.CompactTextString(m) }
func (*ListenReply) ProtoMessage()               {}
func (*ListenReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ListenReply) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ListenReply) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto1.RegisterType((*GetValueRequest)(nil), "proto.GetValueRequest")
	proto1.RegisterType((*GetValueReply)(nil), "proto.GetValueReply")
	proto1.RegisterType((*SetValueRequest)(nil), "proto.SetValueRequest")
	proto1.RegisterType((*SetValueReply)(nil), "proto.SetValueReply")
	proto1.RegisterType((*ListenRequest)(nil), "proto.ListenRequest")
	proto1.RegisterType((*ListenReply)(nil), "proto.ListenReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for KvStore service

type KvStoreClient interface {
	// Gets a value, by key
	GetValue(ctx context.Context, in *GetValueRequest, opts ...grpc.CallOption) (*GetValueReply, error)
	SetValue(ctx context.Context, in *SetValueRequest, opts ...grpc.CallOption) (*SetValueReply, error)
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (KvStore_ListenClient, error)
}

type kvStoreClient struct {
	cc *grpc.ClientConn
}

func NewKvStoreClient(cc *grpc.ClientConn) KvStoreClient {
	return &kvStoreClient{cc}
}

func (c *kvStoreClient) GetValue(ctx context.Context, in *GetValueRequest, opts ...grpc.CallOption) (*GetValueReply, error) {
	out := new(GetValueReply)
	err := grpc.Invoke(ctx, "/proto.KvStore/GetValue", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kvStoreClient) SetValue(ctx context.Context, in *SetValueRequest, opts ...grpc.CallOption) (*SetValueReply, error) {
	out := new(SetValueReply)
	err := grpc.Invoke(ctx, "/proto.KvStore/SetValue", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kvStoreClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (KvStore_ListenClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_KvStore_serviceDesc.Streams[0], c.cc, "/proto.KvStore/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &kvStoreListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KvStore_ListenClient interface {
	Recv() (*ListenReply, error)
	grpc.ClientStream
}

type kvStoreListenClient struct {
	grpc.ClientStream
}

func (x *kvStoreListenClient) Recv() (*ListenReply, error) {
	m := new(ListenReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for KvStore service

type KvStoreServer interface {
	// Gets a value, by key
	GetValue(context.Context, *GetValueRequest) (*GetValueReply, error)
	SetValue(context.Context, *SetValueRequest) (*SetValueReply, error)
	Listen(*ListenRequest, KvStore_ListenServer) error
}

func RegisterKvStoreServer(s *grpc.Server, srv KvStoreServer) {
	s.RegisterService(&_KvStore_serviceDesc, srv)
}

func _KvStore_GetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvStoreServer).GetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KvStore/GetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvStoreServer).GetValue(ctx, req.(*GetValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KvStore_SetValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvStoreServer).SetValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KvStore/SetValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvStoreServer).SetValue(ctx, req.(*SetValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KvStore_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KvStoreServer).Listen(m, &kvStoreListenServer{stream})
}

type KvStore_ListenServer interface {
	Send(*ListenReply) error
	grpc.ServerStream
}

type kvStoreListenServer struct {
	grpc.ServerStream
}

func (x *kvStoreListenServer) Send(m *ListenReply) error {
	return x.ServerStream.SendMsg(m)
}

var _KvStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.KvStore",
	HandlerType: (*KvStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetValue",
			Handler:    _KvStore_GetValue_Handler,
		},
		{
			MethodName: "SetValue",
			Handler:    _KvStore_SetValue_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listen",
			Handler:       _KvStore_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "kvstore.proto",
}

func init() { proto1.RegisterFile("kvstore.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x2e, 0x2b, 0x2e,
	0xc9, 0x2f, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xca, 0x5c,
	0xfc, 0xee, 0xa9, 0x25, 0x61, 0x89, 0x39, 0xa5, 0xa9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25,
	0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6,
	0x92, 0x2a, 0x17, 0x2f, 0x42, 0x51, 0x41, 0x4e, 0xa5, 0x90, 0x08, 0x17, 0x6b, 0x19, 0x88, 0x07,
	0x55, 0x04, 0xe1, 0x28, 0x59, 0x72, 0xf1, 0x07, 0x13, 0x32, 0x0b, 0xa1, 0x95, 0x09, 0x59, 0x2b,
	0x3f, 0x17, 0x6f, 0x30, 0xb2, 0x0d, 0x4a, 0x8a, 0x5c, 0xbc, 0x3e, 0x99, 0xc5, 0x25, 0xa9, 0x79,
	0xb8, 0x5d, 0x65, 0xca, 0xc5, 0x0d, 0x53, 0x02, 0x72, 0x13, 0x91, 0x56, 0x19, 0xed, 0x64, 0xe4,
	0x62, 0xf7, 0x2e, 0x0b, 0x06, 0x05, 0x85, 0x90, 0x15, 0x17, 0x07, 0xcc, 0x63, 0x42, 0x62, 0x90,
	0x80, 0xd1, 0x43, 0x0b, 0x0e, 0x29, 0x11, 0x0c, 0x71, 0x90, 0xfb, 0x18, 0x40, 0x7a, 0x83, 0xd1,
	0xf5, 0x06, 0xe3, 0xd0, 0x1b, 0x8c, 0xa6, 0xd7, 0x8c, 0x8b, 0x0d, 0xe2, 0x74, 0x21, 0x98, 0x0a,
	0x14, 0xcf, 0x4a, 0x09, 0xa1, 0x89, 0x82, 0x75, 0x19, 0x30, 0x26, 0xb1, 0x81, 0x85, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x78, 0x80, 0x54, 0x8f, 0xcc, 0x01, 0x00, 0x00,
}
