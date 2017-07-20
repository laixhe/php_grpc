package xuexitest

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

// TestRequest 请求结构
type TestRequest struct {
	Typeid int64 `protobuf:"varint,1,opt,name=typeid" json:"typeid,omitempty"`
}

func (m *TestRequest) Reset()                    { *m = TestRequest{} }
func (m *TestRequest) String() string            { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()               {}
func (*TestRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TestRequest) GetTypeid() int64 {
	if m != nil {
		return m.Typeid
	}
	return 0
}

// TestReply 响应结构
type TestReply struct {
	Getdataarr []*TestReply_GetData `protobuf:"bytes,1,rep,name=getdataarr" json:"getdataarr,omitempty"`
}

func (m *TestReply) Reset()                    { *m = TestReply{} }
func (m *TestReply) String() string            { return proto.CompactTextString(m) }
func (*TestReply) ProtoMessage()               {}
func (*TestReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TestReply) GetGetdataarr() []*TestReply_GetData {
	if m != nil {
		return m.Getdataarr
	}
	return nil
}

// 返回数据类型
type TestReply_GetData struct {
	Id   int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *TestReply_GetData) Reset()                    { *m = TestReply_GetData{} }
func (m *TestReply_GetData) String() string            { return proto.CompactTextString(m) }
func (*TestReply_GetData) ProtoMessage()               {}
func (*TestReply_GetData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *TestReply_GetData) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TestReply_GetData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*TestRequest)(nil), "xuexitest.TestRequest")
	proto.RegisterType((*TestReply)(nil), "xuexitest.TestReply")
	proto.RegisterType((*TestReply_GetData)(nil), "xuexitest.TestReply.GetData")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Xuexitest service

type XuexitestClient interface {
	// 定义 SayTest 方法
	SayTest(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestReply, error)
}

type xuexitestClient struct {
	cc *grpc.ClientConn
}

func NewXuexitestClient(cc *grpc.ClientConn) XuexitestClient {
	return &xuexitestClient{cc}
}

func (c *xuexitestClient) SayTest(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestReply, error) {
	out := new(TestReply)
	err := grpc.Invoke(ctx, "/xuexitest.Xuexitest/SayTest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Xuexitest service

type XuexitestServer interface {
	// 定义 SayTest 方法
	SayTest(context.Context, *TestRequest) (*TestReply, error)
}

func RegisterXuexitestServer(s *grpc.Server, srv XuexitestServer) {
	s.RegisterService(&_Xuexitest_serviceDesc, srv)
}

func _Xuexitest_SayTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XuexitestServer).SayTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xuexitest.Xuexitest/SayTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XuexitestServer).SayTest(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Xuexitest_serviceDesc = grpc.ServiceDesc{
	ServiceName: "xuexitest.Xuexitest",
	HandlerType: (*XuexitestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayTest",
			Handler:    _Xuexitest_SayTest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "xuexitest.proto",
}

func init() { proto.RegisterFile("xuexitest.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xaf, 0x28, 0x4d, 0xad,
	0xc8, 0x2c, 0x49, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0xa9, 0x72, 0x71, 0x87, 0xa4, 0x16, 0x97, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x89,
	0x71, 0xb1, 0x95, 0x54, 0x16, 0xa4, 0x66, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x41,
	0x79, 0x4a, 0x15, 0x5c, 0x9c, 0x10, 0x65, 0x05, 0x39, 0x95, 0x42, 0x36, 0x5c, 0x5c, 0xe9, 0xa9,
	0x25, 0x29, 0x89, 0x25, 0x89, 0x89, 0x45, 0x45, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x32,
	0x7a, 0x08, 0x4b, 0xe0, 0x2a, 0xf5, 0xdc, 0x53, 0x4b, 0x5c, 0x12, 0x4b, 0x12, 0x83, 0x90, 0xd4,
	0x4b, 0xe9, 0x72, 0xb1, 0x43, 0x85, 0x85, 0xf8, 0xb8, 0x98, 0xe0, 0x36, 0x31, 0x65, 0xa6, 0x08,
	0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9,
	0x46, 0x6e, 0x5c, 0x9c, 0x11, 0x30, 0x93, 0x85, 0x2c, 0xb9, 0xd8, 0x83, 0x13, 0x2b, 0x43, 0xc0,
	0x2e, 0xc5, 0xb0, 0x10, 0xec, 0x03, 0x29, 0x11, 0x6c, 0x0e, 0x51, 0x62, 0x48, 0x62, 0x03, 0x7b,
	0xdd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xb3, 0x29, 0x17, 0x0d, 0x01, 0x00, 0x00,
}
