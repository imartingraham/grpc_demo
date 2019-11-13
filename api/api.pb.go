// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

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

type UserMessage struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserMessage) Reset()         { *m = UserMessage{} }
func (m *UserMessage) String() string { return proto.CompactTextString(m) }
func (*UserMessage) ProtoMessage()    {}
func (*UserMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *UserMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserMessage.Unmarshal(m, b)
}
func (m *UserMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserMessage.Marshal(b, m, deterministic)
}
func (m *UserMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserMessage.Merge(m, src)
}
func (m *UserMessage) XXX_Size() int {
	return xxx_messageInfo_UserMessage.Size(m)
}
func (m *UserMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_UserMessage.DiscardUnknown(m)
}

var xxx_messageInfo_UserMessage proto.InternalMessageInfo

func (m *UserMessage) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserMessage) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func init() {
	proto.RegisterType((*UserMessage)(nil), "api.UserMessage")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0xb2, 0xe1, 0xe2, 0x0e, 0x2d,
	0x4e, 0x2d, 0xf2, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91,
	0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x62, 0xca, 0x4c, 0x11, 0x92, 0xe5, 0xe2, 0x4a, 0xcb, 0x2c,
	0x2a, 0x2e, 0x89, 0xcf, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x04,
	0x8b, 0xf8, 0x25, 0xe6, 0xa6, 0x1a, 0x59, 0x70, 0xb1, 0x80, 0x74, 0x0b, 0x19, 0x70, 0x71, 0xb8,
	0x65, 0xe6, 0xa5, 0x80, 0xd9, 0x02, 0x7a, 0x20, 0x2b, 0x90, 0x0c, 0x95, 0xc2, 0x10, 0x51, 0x62,
	0x48, 0x62, 0x03, 0xbb, 0xc1, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x22, 0xa1, 0x9f, 0x7a, 0x90,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	FindUser(ctx context.Context, in *UserMessage, opts ...grpc.CallOption) (*UserMessage, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) FindUser(ctx context.Context, in *UserMessage, opts ...grpc.CallOption) (*UserMessage, error) {
	out := new(UserMessage)
	err := c.cc.Invoke(ctx, "/api.User/FindUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	FindUser(context.Context, *UserMessage) (*UserMessage, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) FindUser(ctx context.Context, req *UserMessage) (*UserMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUser not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_FindUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).FindUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.User/FindUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).FindUser(ctx, req.(*UserMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindUser",
			Handler:    _User_FindUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}