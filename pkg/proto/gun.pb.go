// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gun.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Hunk struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hunk) Reset()         { *m = Hunk{} }
func (m *Hunk) String() string { return proto.CompactTextString(m) }
func (*Hunk) ProtoMessage()    {}
func (*Hunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_5eb68c7936423302, []int{0}
}

func (m *Hunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hunk.Unmarshal(m, b)
}
func (m *Hunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hunk.Marshal(b, m, deterministic)
}
func (m *Hunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hunk.Merge(m, src)
}
func (m *Hunk) XXX_Size() int {
	return xxx_messageInfo_Hunk.Size(m)
}
func (m *Hunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Hunk.DiscardUnknown(m)
}

var xxx_messageInfo_Hunk proto.InternalMessageInfo

func (m *Hunk) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Hunk)(nil), "Hunk")
}

func init() { proto.RegisterFile("gun.proto", fileDescriptor_5eb68c7936423302) }

var fileDescriptor_5eb68c7936423302 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2f, 0xcd, 0xd3,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe2, 0x62, 0xf1, 0x28, 0xcd, 0xcb, 0x16, 0x12, 0xe2,
	0x62, 0x49, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x02, 0xb3, 0x8d, 0xdc,
	0xb8, 0xb8, 0xdc, 0x4b, 0xf3, 0x82, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0xc4, 0xb9, 0x98,
	0x43, 0x4a, 0xf3, 0x84, 0x58, 0xf5, 0x40, 0xea, 0xa5, 0x20, 0x94, 0x06, 0xa3, 0x01, 0xa3, 0x90,
	0x3c, 0x17, 0x77, 0x48, 0x69, 0x9e, 0x4b, 0x62, 0x49, 0x62, 0x7a, 0x51, 0x62, 0x2e, 0xa6, 0x02,
	0x27, 0xc5, 0x28, 0xf9, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xc0,
	0x32, 0xa3, 0xa2, 0xc4, 0x4a, 0xfd, 0xf4, 0xd2, 0x3c, 0xfd, 0x82, 0xec, 0x74, 0x7d, 0xb0, 0x33,
	0x92, 0xd8, 0xc0, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xae, 0x3c, 0x46, 0x28, 0x9a, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GunServiceClient is the client API for GunService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GunServiceClient interface {
	Tun(ctx context.Context, opts ...grpc.CallOption) (GunService_TunClient, error)
	TunDatagram(ctx context.Context, opts ...grpc.CallOption) (GunService_TunDatagramClient, error)
}

type gunServiceClient struct {
	cc *grpc.ClientConn
}

func NewGunServiceClient(cc *grpc.ClientConn) GunServiceClient {
	return &gunServiceClient{cc}
}

func (c *gunServiceClient) Tun(ctx context.Context, opts ...grpc.CallOption) (GunService_TunClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GunService_serviceDesc.Streams[0], "/GunService/Tun", opts...)
	if err != nil {
		return nil, err
	}
	x := &gunServiceTunClient{stream}
	return x, nil
}

type GunService_TunClient interface {
	Send(*Hunk) error
	Recv() (*Hunk, error)
	grpc.ClientStream
}

type gunServiceTunClient struct {
	grpc.ClientStream
}

func (x *gunServiceTunClient) Send(m *Hunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gunServiceTunClient) Recv() (*Hunk, error) {
	m := new(Hunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gunServiceClient) TunDatagram(ctx context.Context, opts ...grpc.CallOption) (GunService_TunDatagramClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GunService_serviceDesc.Streams[1], "/GunService/TunDatagram", opts...)
	if err != nil {
		return nil, err
	}
	x := &gunServiceTunDatagramClient{stream}
	return x, nil
}

type GunService_TunDatagramClient interface {
	Send(*Hunk) error
	Recv() (*Hunk, error)
	grpc.ClientStream
}

type gunServiceTunDatagramClient struct {
	grpc.ClientStream
}

func (x *gunServiceTunDatagramClient) Send(m *Hunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gunServiceTunDatagramClient) Recv() (*Hunk, error) {
	m := new(Hunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GunServiceServer is the server API for GunService service.
type GunServiceServer interface {
	Tun(GunService_TunServer) error
	TunDatagram(GunService_TunDatagramServer) error
}

// UnimplementedGunServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGunServiceServer struct {
}

func (*UnimplementedGunServiceServer) Tun(srv GunService_TunServer) error {
	return status.Errorf(codes.Unimplemented, "method Tun not implemented")
}
func (*UnimplementedGunServiceServer) TunDatagram(srv GunService_TunDatagramServer) error {
	return status.Errorf(codes.Unimplemented, "method TunDatagram not implemented")
}

func RegisterGunServiceServer(s *grpc.Server, srv GunServiceServer) {
	s.RegisterService(&_GunService_serviceDesc, srv)
}

func _GunService_Tun_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GunServiceServer).Tun(&gunServiceTunServer{stream})
}

type GunService_TunServer interface {
	Send(*Hunk) error
	Recv() (*Hunk, error)
	grpc.ServerStream
}

type gunServiceTunServer struct {
	grpc.ServerStream
}

func (x *gunServiceTunServer) Send(m *Hunk) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gunServiceTunServer) Recv() (*Hunk, error) {
	m := new(Hunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GunService_TunDatagram_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GunServiceServer).TunDatagram(&gunServiceTunDatagramServer{stream})
}

type GunService_TunDatagramServer interface {
	Send(*Hunk) error
	Recv() (*Hunk, error)
	grpc.ServerStream
}

type gunServiceTunDatagramServer struct {
	grpc.ServerStream
}

func (x *gunServiceTunDatagramServer) Send(m *Hunk) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gunServiceTunDatagramServer) Recv() (*Hunk, error) {
	m := new(Hunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GunService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "GunService",
	HandlerType: (*GunServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tun",
			Handler:       _GunService_Tun_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "TunDatagram",
			Handler:       _GunService_TunDatagram_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "gun.proto",
}
