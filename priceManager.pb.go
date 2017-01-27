// Code generated by protoc-gen-go.
// source: priceManager.proto
// DO NOT EDIT!

/*
Package priceManagerClient is a generated protocol buffer package.

It is generated from these files:
	priceManager.proto

It has these top-level messages:
	Entry
	FetchRecordsRequest
	FetchRecordsResponse
	NotifyRequest
	NotifyResponse
*/
package priceManagerClient

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

type Entry struct {
	ProductId int32  `protobuf:"varint,1,opt,name=product_id,json=productId" json:"product_id,omitempty"`
	Version   string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Entry) GetProductId() int32 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *Entry) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type FetchRecordsRequest struct {
}

func (m *FetchRecordsRequest) Reset()                    { *m = FetchRecordsRequest{} }
func (m *FetchRecordsRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchRecordsRequest) ProtoMessage()               {}
func (*FetchRecordsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type FetchRecordsResponse struct {
	Entries []*Entry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
}

func (m *FetchRecordsResponse) Reset()                    { *m = FetchRecordsResponse{} }
func (m *FetchRecordsResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchRecordsResponse) ProtoMessage()               {}
func (*FetchRecordsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FetchRecordsResponse) GetEntries() []*Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type NotifyRequest struct {
	Entries []*Entry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
}

func (m *NotifyRequest) Reset()                    { *m = NotifyRequest{} }
func (m *NotifyRequest) String() string            { return proto.CompactTextString(m) }
func (*NotifyRequest) ProtoMessage()               {}
func (*NotifyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *NotifyRequest) GetEntries() []*Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type NotifyResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *NotifyResponse) Reset()                    { *m = NotifyResponse{} }
func (m *NotifyResponse) String() string            { return proto.CompactTextString(m) }
func (*NotifyResponse) ProtoMessage()               {}
func (*NotifyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *NotifyResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Entry)(nil), "priceManagerClient.Entry")
	proto.RegisterType((*FetchRecordsRequest)(nil), "priceManagerClient.FetchRecordsRequest")
	proto.RegisterType((*FetchRecordsResponse)(nil), "priceManagerClient.FetchRecordsResponse")
	proto.RegisterType((*NotifyRequest)(nil), "priceManagerClient.NotifyRequest")
	proto.RegisterType((*NotifyResponse)(nil), "priceManagerClient.NotifyResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PriceManager service

type PriceManagerClient interface {
	GetPriceUpdateRecords(ctx context.Context, in *FetchRecordsRequest, opts ...grpc.CallOption) (*FetchRecordsResponse, error)
	NotifySuccessfullyPicked(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*NotifyResponse, error)
	NotifySuccessfullyProcessed(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*NotifyResponse, error)
}

type priceManagerClient struct {
	cc *grpc.ClientConn
}

func NewPriceManagerClient(cc *grpc.ClientConn) PriceManagerClient {
	return &priceManagerClient{cc}
}

func (c *priceManagerClient) GetPriceUpdateRecords(ctx context.Context, in *FetchRecordsRequest, opts ...grpc.CallOption) (*FetchRecordsResponse, error) {
	out := new(FetchRecordsResponse)
	err := grpc.Invoke(ctx, "/priceManagerClient.PriceManager/GetPriceUpdateRecords", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *priceManagerClient) NotifySuccessfullyPicked(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*NotifyResponse, error) {
	out := new(NotifyResponse)
	err := grpc.Invoke(ctx, "/priceManagerClient.PriceManager/NotifySuccessfullyPicked", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *priceManagerClient) NotifySuccessfullyProcessed(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*NotifyResponse, error) {
	out := new(NotifyResponse)
	err := grpc.Invoke(ctx, "/priceManagerClient.PriceManager/NotifySuccessfullyProcessed", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PriceManager service

type PriceManagerServer interface {
	GetPriceUpdateRecords(context.Context, *FetchRecordsRequest) (*FetchRecordsResponse, error)
	NotifySuccessfullyPicked(context.Context, *NotifyRequest) (*NotifyResponse, error)
	NotifySuccessfullyProcessed(context.Context, *NotifyRequest) (*NotifyResponse, error)
}

func RegisterPriceManagerServer(s *grpc.Server, srv PriceManagerServer) {
	s.RegisterService(&_PriceManager_serviceDesc, srv)
}

func _PriceManager_GetPriceUpdateRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceManagerServer).GetPriceUpdateRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/priceManagerClient.PriceManager/GetPriceUpdateRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceManagerServer).GetPriceUpdateRecords(ctx, req.(*FetchRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PriceManager_NotifySuccessfullyPicked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceManagerServer).NotifySuccessfullyPicked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/priceManagerClient.PriceManager/NotifySuccessfullyPicked",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceManagerServer).NotifySuccessfullyPicked(ctx, req.(*NotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PriceManager_NotifySuccessfullyProcessed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceManagerServer).NotifySuccessfullyProcessed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/priceManagerClient.PriceManager/NotifySuccessfullyProcessed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceManagerServer).NotifySuccessfullyProcessed(ctx, req.(*NotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PriceManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "priceManagerClient.PriceManager",
	HandlerType: (*PriceManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPriceUpdateRecords",
			Handler:    _PriceManager_GetPriceUpdateRecords_Handler,
		},
		{
			MethodName: "NotifySuccessfullyPicked",
			Handler:    _PriceManager_NotifySuccessfullyPicked_Handler,
		},
		{
			MethodName: "NotifySuccessfullyProcessed",
			Handler:    _PriceManager_NotifySuccessfullyProcessed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "priceManager.proto",
}

func init() { proto.RegisterFile("priceManager.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x92, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0xed, 0xc8, 0x38, 0xf4, 0xfa, 0xb3, 0x88, 0x0e, 0xc4, 0x11, 0xa1, 0x66, 0x63, 0x71,
	0xd1, 0xc5, 0xcc, 0x0b, 0x08, 0xfe, 0x21, 0xa2, 0x0c, 0x11, 0xd7, 0x52, 0x93, 0xdb, 0x31, 0x58,
	0x93, 0x9a, 0xa4, 0x42, 0x1f, 0xd1, 0xb7, 0x12, 0x3b, 0xad, 0x58, 0x2c, 0x0a, 0xe2, 0x2e, 0xf7,
	0x70, 0x38, 0xe7, 0xbb, 0x97, 0x00, 0x29, 0xac, 0x12, 0x78, 0x9d, 0xea, 0x74, 0x81, 0x36, 0x29,
	0xac, 0xf1, 0x86, 0x74, 0xb4, 0x93, 0x5c, 0xa1, 0xf6, 0xec, 0x18, 0x86, 0x67, 0xda, 0xdb, 0x8a,
	0xec, 0x03, 0x14, 0xd6, 0xc8, 0x52, 0xf8, 0x7b, 0x25, 0x69, 0x10, 0x05, 0xf1, 0x90, 0x87, 0x8d,
	0x72, 0x29, 0x09, 0x85, 0xd1, 0x2b, 0x5a, 0xa7, 0x8c, 0xa6, 0x83, 0x28, 0x88, 0x43, 0xde, 0x8e,
	0x6c, 0x0c, 0xdb, 0xe7, 0xe8, 0xc5, 0x23, 0x47, 0x61, 0xac, 0x74, 0x1c, 0x5f, 0x4a, 0x74, 0x9e,
	0x5d, 0xc1, 0x4e, 0x57, 0x76, 0x85, 0xd1, 0x0e, 0xc9, 0x0c, 0x46, 0xa8, 0xbd, 0x55, 0xe8, 0x68,
	0x10, 0xad, 0xc6, 0xeb, 0xd3, 0xdd, 0xe4, 0x3b, 0x56, 0x52, 0x33, 0xf1, 0xd6, 0xc9, 0x4e, 0x61,
	0xf3, 0xc6, 0x78, 0x95, 0x55, 0x4d, 0xfa, 0xdf, 0x52, 0x8e, 0x60, 0xab, 0x4d, 0x69, 0x60, 0x28,
	0x8c, 0x9e, 0xd1, 0xb9, 0x74, 0x81, 0xf5, 0xc6, 0x21, 0x6f, 0xc7, 0xe9, 0xdb, 0x00, 0x36, 0xe6,
	0x5f, 0x12, 0x49, 0x0e, 0xe3, 0x0b, 0xf4, 0xb5, 0x74, 0x57, 0xc8, 0xd4, 0x63, 0xb3, 0x18, 0x39,
	0xec, 0x6b, 0xee, 0xb9, 0xc8, 0x24, 0xfe, 0xdd, 0xb8, 0xc4, 0x62, 0x2b, 0x44, 0x00, 0x5d, 0xa2,
	0xde, 0x96, 0x42, 0xa0, 0x73, 0x59, 0x99, 0xe7, 0xd5, 0x5c, 0x89, 0x27, 0x94, 0xe4, 0xa0, 0x2f,
	0xa7, 0x73, 0x9e, 0x09, 0xfb, 0xc9, 0xf2, 0x59, 0x92, 0xc1, 0x5e, 0x4f, 0x89, 0x35, 0x1f, 0xef,
	0x7f, 0xec, 0x79, 0x58, 0xab, 0xbf, 0xdf, 0xec, 0x3d, 0x00, 0x00, 0xff, 0xff, 0x17, 0xd8, 0x7a,
	0xdf, 0x94, 0x02, 0x00, 0x00,
}
