// Code generated by protoc-gen-go. DO NOT EDIT.
// source: oswee/identity/signup/proto/v1/signup_command_handler.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type PublishRequest struct {
	Api                  string               `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Signup               *Signup              `protobuf:"bytes,2,opt,name=signup,proto3" json:"signup,omitempty"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PublishRequest) Reset()         { *m = PublishRequest{} }
func (m *PublishRequest) String() string { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()    {}
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3853b9605670ad17, []int{0}
}

func (m *PublishRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRequest.Unmarshal(m, b)
}
func (m *PublishRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRequest.Marshal(b, m, deterministic)
}
func (m *PublishRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRequest.Merge(m, src)
}
func (m *PublishRequest) XXX_Size() int {
	return xxx_messageInfo_PublishRequest.Size(m)
}
func (m *PublishRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRequest proto.InternalMessageInfo

func (m *PublishRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *PublishRequest) GetSignup() *Signup {
	if m != nil {
		return m.Signup
	}
	return nil
}

func (m *PublishRequest) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func init() {
	proto.RegisterType((*PublishRequest)(nil), "oswee.identity.signup.v1.PublishRequest")
}

func init() {
	proto.RegisterFile("oswee/identity/signup/proto/v1/signup_command_handler.proto", fileDescriptor_3853b9605670ad17)
}

var fileDescriptor_3853b9605670ad17 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xa9, 0x0b, 0x2b, 0x66, 0x41, 0x24, 0x07, 0x29, 0x5d, 0xd0, 0xe2, 0xa9, 0x20, 0x24,
	0x74, 0xbd, 0x08, 0x7b, 0x53, 0x04, 0xc1, 0x8b, 0x74, 0x3d, 0x79, 0x29, 0xfd, 0x33, 0x76, 0x23,
	0x6d, 0x13, 0x9b, 0xa4, 0xb2, 0x1f, 0xc5, 0x6f, 0x2b, 0xcd, 0xb4, 0x07, 0xff, 0x2c, 0xec, 0x6d,
	0xda, 0x79, 0x6f, 0xf2, 0x7b, 0x8f, 0xac, 0xa5, 0xfe, 0x04, 0xe0, 0xa2, 0x84, 0xd6, 0x08, 0xb3,
	0xe3, 0x5a, 0x54, 0xad, 0x55, 0x5c, 0x75, 0xd2, 0x48, 0xde, 0xc7, 0xe3, 0x77, 0x5a, 0xc8, 0xa6,
	0xc9, 0xda, 0x32, 0xdd, 0x66, 0x6d, 0x59, 0x43, 0xc7, 0xdc, 0x9e, 0xfa, 0xce, 0xcc, 0x26, 0x33,
	0x43, 0x31, 0xeb, 0xe3, 0xe0, 0xb2, 0x92, 0xb2, 0xaa, 0x01, 0xef, 0xe4, 0xf6, 0x8d, 0x1b, 0xd1,
	0x80, 0x36, 0x59, 0xa3, 0xd0, 0x1a, 0x2c, 0x7f, 0x0b, 0xa0, 0x51, 0x66, 0x37, 0x2e, 0xaf, 0x0f,
	0x82, 0x42, 0xf1, 0xd5, 0x97, 0x47, 0x4e, 0x9f, 0x6d, 0x5e, 0x0b, 0xbd, 0x4d, 0xe0, 0xc3, 0x82,
	0x36, 0xf4, 0x8c, 0xcc, 0x32, 0x25, 0x7c, 0x2f, 0xf4, 0xa2, 0x93, 0x64, 0x18, 0xe9, 0x2d, 0x99,
	0xa3, 0xc9, 0x3f, 0x0a, 0xbd, 0x68, 0xb1, 0x0a, 0xd9, 0x3e, 0x74, 0xb6, 0x71, 0x53, 0x32, 0xea,
	0xe9, 0x9a, 0x2c, 0x8a, 0x0e, 0x32, 0x03, 0xe9, 0x10, 0xc1, 0x9f, 0x39, 0x7b, 0xc0, 0x10, 0x9f,
	0x4d, 0xf8, 0xec, 0x65, 0xca, 0x97, 0x10, 0x94, 0x0f, 0x3f, 0x56, 0xef, 0x64, 0x89, 0xe7, 0xee,
	0xb1, 0xbf, 0x47, 0xac, 0x6f, 0x03, 0x5d, 0x2f, 0x0a, 0xa0, 0x4f, 0xe4, 0x78, 0x24, 0xa7, 0xd1,
	0x7e, 0xa0, 0x9f, 0xe1, 0x82, 0xf3, 0x3f, 0x6f, 0x3f, 0x0c, 0xd5, 0xdd, 0x85, 0xaf, 0x17, 0xff,
	0xd7, 0xa6, 0x8d, 0xcd, 0x35, 0xef, 0xe3, 0x7c, 0xee, 0x1c, 0x37, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x75, 0x05, 0xeb, 0x88, 0xf4, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SignupCommandHandlerServiceClient is the client API for SignupCommandHandlerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SignupCommandHandlerServiceClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type signupCommandHandlerServiceClient struct {
	cc *grpc.ClientConn
}

func NewSignupCommandHandlerServiceClient(cc *grpc.ClientConn) SignupCommandHandlerServiceClient {
	return &signupCommandHandlerServiceClient{cc}
}

func (c *signupCommandHandlerServiceClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/oswee.identity.signup.v1.SignupCommandHandlerService/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignupCommandHandlerServiceServer is the server API for SignupCommandHandlerService service.
type SignupCommandHandlerServiceServer interface {
	Publish(context.Context, *PublishRequest) (*empty.Empty, error)
}

// UnimplementedSignupCommandHandlerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSignupCommandHandlerServiceServer struct {
}

func (*UnimplementedSignupCommandHandlerServiceServer) Publish(ctx context.Context, req *PublishRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}

func RegisterSignupCommandHandlerServiceServer(s *grpc.Server, srv SignupCommandHandlerServiceServer) {
	s.RegisterService(&_SignupCommandHandlerService_serviceDesc, srv)
}

func _SignupCommandHandlerService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupCommandHandlerServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oswee.identity.signup.v1.SignupCommandHandlerService/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupCommandHandlerServiceServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SignupCommandHandlerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "oswee.identity.signup.v1.SignupCommandHandlerService",
	HandlerType: (*SignupCommandHandlerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _SignupCommandHandlerService_Publish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oswee/identity/signup/proto/v1/signup_command_handler.proto",
}
