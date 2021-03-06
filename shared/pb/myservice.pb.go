// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/rickychandra/goboil/shared/pb/myservice.proto

/*
Package myservice is a generated protocol buffer package.

It is generated from these files:
	github.com/rickychandra/goboil/shared/pb/myservice.proto

It has these top-level messages:
	HelloRequest
	HelloResponse
	HelloErrorRequest
*/
package myservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/rickychandra/goboil/common/pb"

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

type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age" json:"age,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HelloRequest) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type HelloResponse struct {
	Ok bool `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type HelloErrorRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloErrorRequest) Reset()                    { *m = HelloErrorRequest{} }
func (m *HelloErrorRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloErrorRequest) ProtoMessage()               {}
func (*HelloErrorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HelloErrorRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "myservice.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "myservice.HelloResponse")
	proto.RegisterType((*HelloErrorRequest)(nil), "myservice.HelloErrorRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MyService service

type MyServiceClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	HelloError(ctx context.Context, in *HelloErrorRequest, opts ...grpc.CallOption) (*common.Empty, error)
}

type myServiceClient struct {
	cc *grpc.ClientConn
}

func NewMyServiceClient(cc *grpc.ClientConn) MyServiceClient {
	return &myServiceClient{cc}
}

func (c *myServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := grpc.Invoke(ctx, "/myservice.MyService/Hello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myServiceClient) HelloError(ctx context.Context, in *HelloErrorRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := grpc.Invoke(ctx, "/myservice.MyService/HelloError", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MyService service

type MyServiceServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	HelloError(context.Context, *HelloErrorRequest) (*common.Empty, error)
}

func RegisterMyServiceServer(s *grpc.Server, srv MyServiceServer) {
	s.RegisterService(&_MyService_serviceDesc, srv)
}

func _MyService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myservice.MyService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyService_HelloError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloErrorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServiceServer).HelloError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myservice.MyService/HelloError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServiceServer).HelloError(ctx, req.(*HelloErrorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "myservice.MyService",
	HandlerType: (*MyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _MyService_Hello_Handler,
		},
		{
			MethodName: "HelloError",
			Handler:    _MyService_HelloError_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rickychandra/goboil/shared/pb/myservice.proto",
}

func init() {
	proto.RegisterFile("github.com/rickychandra/goboil/shared/pb/myservice.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x9b, 0x68, 0xc5, 0x0c, 0x56, 0x74, 0x2e, 0x86, 0x20, 0x58, 0xf6, 0x62, 0x4f, 0x59,
	0xf0, 0x0f, 0x08, 0x7a, 0x2d, 0x78, 0xf1, 0x12, 0x3f, 0xc1, 0x66, 0x3b, 0x24, 0x21, 0xd9, 0x4c,
	0xdc, 0x4d, 0x85, 0x7c, 0x01, 0x3f, 0xb7, 0x74, 0x53, 0x6c, 0x51, 0xb0, 0xb7, 0xc7, 0xee, 0xbc,
	0xdf, 0xbc, 0x79, 0xf0, 0x54, 0x54, 0x7d, 0xb9, 0xce, 0x53, 0xcd, 0x46, 0xda, 0x4a, 0xd7, 0x83,
	0x2e, 0x55, 0xbb, 0xb2, 0x4a, 0x16, 0x9c, 0x73, 0xd5, 0x48, 0x57, 0x2a, 0x4b, 0x2b, 0xd9, 0xe5,
	0xd2, 0x0c, 0x8e, 0xec, 0x67, 0xa5, 0x29, 0xed, 0x2c, 0xf7, 0x8c, 0xd1, 0xcf, 0x43, 0xf2, 0x78,
	0x00, 0xa2, 0xd9, 0x18, 0x6e, 0x37, 0x90, 0x51, 0x8d, 0x04, 0xf1, 0x00, 0x67, 0xaf, 0xd4, 0x34,
	0x9c, 0xd1, 0xc7, 0x9a, 0x5c, 0x8f, 0x08, 0xc7, 0xad, 0x32, 0x14, 0x07, 0xf3, 0x60, 0x11, 0x65,
	0x5e, 0xe3, 0x05, 0x1c, 0xa9, 0x82, 0xe2, 0x70, 0x1e, 0x2c, 0xa6, 0xd9, 0x46, 0x8a, 0x1b, 0x98,
	0x6d, 0x5d, 0xae, 0xe3, 0xd6, 0x11, 0x9e, 0x43, 0xc8, 0xb5, 0x37, 0x9d, 0x66, 0x21, 0xd7, 0xe2,
	0x16, 0x2e, 0xfd, 0xc0, 0xd2, 0x5a, 0xb6, 0xff, 0xb0, 0xef, 0xbe, 0x02, 0x88, 0xde, 0x86, 0xf7,
	0xf1, 0x08, 0x7c, 0x81, 0xa9, 0xb7, 0xe1, 0x55, 0xba, 0x3b, 0x75, 0x3f, 0x5f, 0x12, 0xff, 0xfd,
	0x18, 0x23, 0x88, 0x09, 0x3e, 0x03, 0xec, 0x96, 0xe2, 0xf5, 0xef, 0xc9, 0xfd, 0x2c, 0xc9, 0x2c,
	0xdd, 0xd6, 0xb0, 0x34, 0x5d, 0x3f, 0x88, 0x49, 0x7e, 0xe2, 0xfb, 0xb8, 0xff, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0x68, 0xee, 0x76, 0xc1, 0x8d, 0x01, 0x00, 0x00,
}
