// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/sharding/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("provenance/sharding/v1/query.proto", fileDescriptor_6a3659de35c7b13f)
}

var fileDescriptor_6a3659de35c7b13f = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2a, 0x28, 0xca, 0x2f,
	0x4b, 0xcd, 0x4b, 0xcc, 0x4b, 0x4e, 0xd5, 0x2f, 0xce, 0x48, 0x2c, 0x4a, 0xc9, 0xcc, 0x4b, 0xd7,
	0x2f, 0x33, 0xd4, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12,
	0x43, 0xa8, 0xd1, 0x83, 0xa9, 0xd1, 0x2b, 0x33, 0x34, 0x62, 0xe7, 0x62, 0x0d, 0x04, 0x29, 0x73,
	0xca, 0x3e, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c,
	0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x06, 0x2e, 0xc9, 0xcc, 0x7c, 0x3d,
	0xec, 0xba, 0x03, 0x18, 0xa3, 0x4c, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73,
	0xf5, 0x11, 0x8a, 0x74, 0x33, 0xf3, 0x91, 0x78, 0xfa, 0x15, 0x08, 0x67, 0x95, 0x54, 0x16, 0xa4,
	0x16, 0x27, 0xb1, 0x81, 0x1d, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x29, 0xba, 0xac, 0xbe,
	0xba, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "provenance.sharding.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "provenance/sharding/v1/query.proto",
}
