// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: demo.proto

/*
Package demo is a generated protocol buffer package.

It is generated from these files:
	demo.proto

It has these top-level messages:
	Request
	Reply
*/
package demo

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the user's name.
type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptorDemo, []int{0} }

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type Reply struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptorDemo, []int{1} }

func (m *Reply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "demo.Request")
	proto.RegisterType((*Reply)(nil), "demo.Reply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Demo service

type DemoClient interface {
	// 单次调用
	SayHello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	// 客户端流
	SayHelloCliStream(ctx context.Context, opts ...grpc.CallOption) (Demo_SayHelloCliStreamClient, error)
	// 服务端流
	SayHelloServStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (Demo_SayHelloServStreamClient, error)
	// 双向流
	SayHelloBidiStream(ctx context.Context, opts ...grpc.CallOption) (Demo_SayHelloBidiStreamClient, error)
}

type demoClient struct {
	cc *grpc.ClientConn
}

func NewDemoClient(cc *grpc.ClientConn) DemoClient {
	return &demoClient{cc}
}

func (c *demoClient) SayHello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/demo.demo/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoClient) SayHelloCliStream(ctx context.Context, opts ...grpc.CallOption) (Demo_SayHelloCliStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Demo_serviceDesc.Streams[0], c.cc, "/demo.demo/SayHelloCliStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &demoSayHelloCliStreamClient{stream}
	return x, nil
}

type Demo_SayHelloCliStreamClient interface {
	Send(*Request) error
	CloseAndRecv() (*Reply, error)
	grpc.ClientStream
}

type demoSayHelloCliStreamClient struct {
	grpc.ClientStream
}

func (x *demoSayHelloCliStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *demoSayHelloCliStreamClient) CloseAndRecv() (*Reply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Reply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *demoClient) SayHelloServStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (Demo_SayHelloServStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Demo_serviceDesc.Streams[1], c.cc, "/demo.demo/SayHelloServStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &demoSayHelloServStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Demo_SayHelloServStreamClient interface {
	Recv() (*Reply, error)
	grpc.ClientStream
}

type demoSayHelloServStreamClient struct {
	grpc.ClientStream
}

func (x *demoSayHelloServStreamClient) Recv() (*Reply, error) {
	m := new(Reply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *demoClient) SayHelloBidiStream(ctx context.Context, opts ...grpc.CallOption) (Demo_SayHelloBidiStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Demo_serviceDesc.Streams[2], c.cc, "/demo.demo/SayHelloBidiStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &demoSayHelloBidiStreamClient{stream}
	return x, nil
}

type Demo_SayHelloBidiStreamClient interface {
	Send(*Request) error
	Recv() (*Reply, error)
	grpc.ClientStream
}

type demoSayHelloBidiStreamClient struct {
	grpc.ClientStream
}

func (x *demoSayHelloBidiStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *demoSayHelloBidiStreamClient) Recv() (*Reply, error) {
	m := new(Reply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Demo service

type DemoServer interface {
	// 单次调用
	SayHello(context.Context, *Request) (*Reply, error)
	// 客户端流
	SayHelloCliStream(Demo_SayHelloCliStreamServer) error
	// 服务端流
	SayHelloServStream(*Request, Demo_SayHelloServStreamServer) error
	// 双向流
	SayHelloBidiStream(Demo_SayHelloBidiStreamServer) error
}

func RegisterDemoServer(s *grpc.Server, srv DemoServer) {
	s.RegisterService(&_Demo_serviceDesc, srv)
}

func _Demo_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.demo/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServer).SayHello(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Demo_SayHelloCliStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DemoServer).SayHelloCliStream(&demoSayHelloCliStreamServer{stream})
}

type Demo_SayHelloCliStreamServer interface {
	SendAndClose(*Reply) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type demoSayHelloCliStreamServer struct {
	grpc.ServerStream
}

func (x *demoSayHelloCliStreamServer) SendAndClose(m *Reply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *demoSayHelloCliStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Demo_SayHelloServStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DemoServer).SayHelloServStream(m, &demoSayHelloServStreamServer{stream})
}

type Demo_SayHelloServStreamServer interface {
	Send(*Reply) error
	grpc.ServerStream
}

type demoSayHelloServStreamServer struct {
	grpc.ServerStream
}

func (x *demoSayHelloServStreamServer) Send(m *Reply) error {
	return x.ServerStream.SendMsg(m)
}

func _Demo_SayHelloBidiStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DemoServer).SayHelloBidiStream(&demoSayHelloBidiStreamServer{stream})
}

type Demo_SayHelloBidiStreamServer interface {
	Send(*Reply) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type demoSayHelloBidiStreamServer struct {
	grpc.ServerStream
}

func (x *demoSayHelloBidiStreamServer) Send(m *Reply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *demoSayHelloBidiStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Demo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "demo.demo",
	HandlerType: (*DemoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Demo_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloCliStream",
			Handler:       _Demo_SayHelloCliStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SayHelloServStream",
			Handler:       _Demo_SayHelloServStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayHelloBidiStream",
			Handler:       _Demo_SayHelloBidiStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "demo.proto",
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDemo(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func (m *Reply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Reply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDemo(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	return i, nil
}

func encodeVarintDemo(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Request) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDemo(uint64(l))
	}
	return n
}

func (m *Reply) Size() (n int) {
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovDemo(uint64(l))
	}
	return n
}

func sovDemo(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDemo(x uint64) (n int) {
	return sovDemo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDemo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDemo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDemo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDemo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Reply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDemo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Reply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Reply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDemo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDemo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDemo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDemo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDemo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDemo
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDemo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDemo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthDemo
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDemo
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipDemo(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthDemo = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDemo   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("demo.proto", fileDescriptorDemo) }

var fileDescriptorDemo = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x49, 0xcd, 0xcd,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x64, 0xb9, 0xd8, 0x83, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x45, 0x2e, 0xd6, 0xa0, 0xd4, 0x82, 0x9c, 0x4a, 0x21,
	0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0x74, 0x98, 0x3c, 0x8c, 0x6b, 0x74, 0x99, 0x91,
	0x0b, 0x6c, 0x94, 0x90, 0x06, 0x17, 0x47, 0x70, 0x62, 0xa5, 0x47, 0x6a, 0x4e, 0x4e, 0xbe, 0x10,
	0xaf, 0x1e, 0xd8, 0x26, 0xa8, 0xd1, 0x52, 0xdc, 0x30, 0x6e, 0x41, 0x4e, 0xa5, 0x12, 0x83, 0x90,
	0x31, 0x97, 0x20, 0x4c, 0xa5, 0x73, 0x4e, 0x66, 0x70, 0x49, 0x51, 0x6a, 0x62, 0x2e, 0x7e, 0x2d,
	0x1a, 0x8c, 0x42, 0x26, 0x5c, 0x42, 0x30, 0x4d, 0xc1, 0xa9, 0x45, 0x65, 0xc4, 0xe8, 0x32, 0x60,
	0x14, 0x32, 0x43, 0xe8, 0x72, 0xca, 0x4c, 0x21, 0xd2, 0x2e, 0x03, 0x46, 0x27, 0xa1, 0x13, 0x8f,
	0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x06, 0x0f,
	0xe6, 0x24, 0x36, 0x70, 0xc0, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x7e, 0x35, 0x65,
	0x46, 0x01, 0x00, 0x00,
}
