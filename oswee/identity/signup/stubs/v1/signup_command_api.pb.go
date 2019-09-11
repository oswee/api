// Code generated by protoc-gen-go. DO NOT EDIT.
// source: oswee/identity/signup/proto/v1/signup_command_api.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateRequest struct {
	Api                  string               `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Signup               *Signup              `protobuf:"bytes,2,opt,name=signup,proto3" json:"signup,omitempty"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbd1e278ba3b0982, []int{0}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateRequest) GetSignup() *Signup {
	if m != nil {
		return m.Signup
	}
	return nil
}

func (m *CreateRequest) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "oswee.identity.signup.v1.CreateRequest")
}

func init() {
	proto.RegisterFile("oswee/identity/signup/proto/v1/signup_command_api.proto", fileDescriptor_dbd1e278ba3b0982)
}

var fileDescriptor_dbd1e278ba3b0982 = []byte{
	// 631 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0x6e, 0x13, 0x3d,
	0x10, 0xd6, 0x26, 0x6d, 0x7e, 0xfd, 0xae, 0x2a, 0xa2, 0x55, 0xa9, 0x42, 0xa8, 0xc0, 0xe4, 0x42,
	0x15, 0xa8, 0x9d, 0xa4, 0x91, 0x40, 0xa9, 0x90, 0x68, 0xd3, 0xaa, 0x8a, 0x14, 0xa0, 0x4a, 0xca,
	0x85, 0x4b, 0x71, 0x36, 0xd3, 0x8d, 0x21, 0x6b, 0x2f, 0xb6, 0x77, 0xd3, 0x72, 0xe4, 0x11, 0xca,
	0x8d, 0x37, 0xe1, 0xc2, 0x43, 0x94, 0x33, 0x37, 0xc4, 0x73, 0xa0, 0xb5, 0x37, 0x2a, 0xd0, 0x16,
	0x71, 0x5a, 0xef, 0x7c, 0xdf, 0x7c, 0x33, 0xf3, 0x79, 0x8c, 0x1e, 0x49, 0x3d, 0x03, 0xa0, 0x7c,
	0x0c, 0xc2, 0x70, 0x73, 0x4a, 0x35, 0x0f, 0x45, 0x12, 0xd3, 0x58, 0x49, 0x23, 0x69, 0xda, 0xcc,
	0xff, 0x8f, 0x02, 0x19, 0x45, 0x4c, 0x8c, 0x8f, 0x58, 0xcc, 0x89, 0xc5, 0xfc, 0x8a, 0x4d, 0x24,
	0xf3, 0x44, 0xe2, 0x88, 0x24, 0x6d, 0x56, 0xef, 0x86, 0x52, 0x86, 0x53, 0x70, 0x1a, 0xa3, 0xe4,
	0x98, 0x1a, 0x1e, 0x81, 0x36, 0x2c, 0x8a, 0x5d, 0x6a, 0xf5, 0xf6, 0x9f, 0x04, 0x88, 0x62, 0x73,
	0x9a, 0x83, 0x6b, 0x39, 0xc8, 0x62, 0x4e, 0x99, 0x10, 0xd2, 0x30, 0xc3, 0xa5, 0xd0, 0x39, 0xfa,
	0xd0, 0x7e, 0x82, 0x8d, 0x10, 0xc4, 0x86, 0x9e, 0xb1, 0x30, 0x04, 0x45, 0x65, 0x6c, 0x19, 0x57,
	0xb0, 0x1f, 0xfc, 0xd3, 0x70, 0x8e, 0x5c, 0xfb, 0xe1, 0xa1, 0xe5, 0xae, 0x02, 0x66, 0x60, 0x00,
	0xef, 0x12, 0xd0, 0xc6, 0x2f, 0xa3, 0x22, 0x8b, 0x79, 0xc5, 0xc3, 0xde, 0xfa, 0xff, 0x83, 0xec,
	0xe8, 0x3f, 0x46, 0x25, 0x97, 0x53, 0x29, 0x60, 0x6f, 0x7d, 0xa9, 0x85, 0xc9, 0x75, 0x2e, 0x90,
	0xa1, 0x3d, 0x0d, 0x72, 0xbe, 0xbf, 0x85, 0x96, 0x02, 0x2b, 0x7e, 0x94, 0xb9, 0x51, 0x29, 0xda,
	0xf4, 0x2a, 0x71, 0xc3, 0x92, 0xb9, 0x13, 0xe4, 0x70, 0x6e, 0xd5, 0x00, 0x39, 0x7a, 0x16, 0xe8,
	0xf4, 0xcf, 0xb6, 0x7b, 0x68, 0xbf, 0xbe, 0x9c, 0x6b, 0xba, 0xf6, 0x5a, 0xb7, 0x76, 0xe1, 0x98,
	0x0b, 0xc0, 0x02, 0x66, 0x38, 0x43, 0x5e, 0xc6, 0x58, 0x39, 0xe8, 0xdc, 0xcb, 0xba, 0x3d, 0xf7,
	0xf2, 0xe2, 0xe7, 0xde, 0xaf, 0xb5, 0x5b, 0x27, 0x68, 0xc5, 0x09, 0x75, 0xdd, 0xa5, 0x0e, 0x41,
	0xa5, 0x3c, 0x00, 0xff, 0x35, 0x2a, 0xb9, 0xf9, 0xfd, 0xfb, 0xd7, 0x8f, 0xf5, 0x9b, 0x43, 0xd5,
	0xd5, 0x4b, 0x03, 0xec, 0x65, 0x57, 0x59, 0xbb, 0xf9, 0xe1, 0xeb, 0xf7, 0x8f, 0x85, 0x1b, 0x35,
	0x74, 0xe1, 0x72, 0xc7, 0xab, 0xef, 0x7c, 0x2b, 0x9e, 0x6d, 0x7f, 0x2e, 0xfa, 0x5f, 0x3c, 0x54,
	0x7e, 0x91, 0x15, 0xc0, 0xae, 0x0f, 0xbc, 0x7d, 0xd0, 0xf3, 0x57, 0x0f, 0x27, 0x80, 0x43, 0xab,
	0xae, 0x4d, 0x16, 0xc1, 0x90, 0x82, 0xba, 0x57, 0x7b, 0x82, 0x16, 0x2d, 0xd7, 0xaf, 0x4c, 0x8c,
	0x89, 0x75, 0x87, 0xd2, 0x90, 0x9b, 0x49, 0x32, 0x22, 0x81, 0x8c, 0xa8, 0x6d, 0xb3, 0x5a, 0x1d,
	0xbf, 0xe7, 0xc2, 0x30, 0xa5, 0xc9, 0xdb, 0x29, 0x4b, 0xb9, 0xd0, 0x4f, 0xc3, 0x88, 0xf1, 0x69,
	0xc6, 0xa8, 0xf7, 0xd1, 0xca, 0xce, 0x70, 0x17, 0x6f, 0x6e, 0x74, 0xa7, 0x2c, 0xd1, 0x80, 0xfb,
	0x3c, 0x00, 0xa1, 0xc1, 0x6f, 0x5f, 0xa7, 0x66, 0x17, 0x6f, 0x34, 0x95, 0x23, 0x1a, 0x31, 0x6d,
	0x40, 0xd1, 0x7e, 0xaf, 0xbb, 0xf7, 0x7c, 0xb8, 0x47, 0xcc, 0x89, 0x69, 0x2d, 0x36, 0x49, 0x83,
	0x34, 0xaa, 0xcb, 0xd9, 0x33, 0x70, 0x16, 0x65, 0x35, 0x0a, 0x85, 0x85, 0x56, 0x99, 0xc5, 0xf1,
	0x94, 0x07, 0x76, 0xfd, 0xe8, 0x1b, 0x2d, 0x45, 0xe7, 0x52, 0x64, 0x70, 0x80, 0x8a, 0xed, 0xc6,
	0xa6, 0xdf, 0x43, 0xfb, 0x03, 0x30, 0x89, 0x12, 0x30, 0xc6, 0xb3, 0x09, 0x08, 0x6c, 0x26, 0x80,
	0x13, 0x0d, 0x0a, 0x8f, 0x25, 0x68, 0x2c, 0xa4, 0xc1, 0x13, 0x96, 0x02, 0x8e, 0x41, 0x45, 0x5c,
	0x6b, 0x2e, 0x05, 0x36, 0x12, 0xb3, 0x20, 0x00, 0xad, 0x2d, 0x57, 0x81, 0x96, 0x89, 0x0a, 0x80,
	0x0c, 0xb6, 0x32, 0xc5, 0xb6, 0xdf, 0x46, 0xf5, 0xcb, 0x8a, 0x73, 0xd6, 0x85, 0x2a, 0x9c, 0x70,
	0x6d, 0x88, 0x5f, 0x42, 0x0b, 0x9f, 0x0a, 0xde, 0x7f, 0xaa, 0x81, 0xca, 0xcf, 0xa4, 0x02, 0xcc,
	0x46, 0x32, 0x31, 0xd8, 0x99, 0xbc, 0xf6, 0x37, 0x5b, 0x5e, 0xdd, 0xb9, 0xfa, 0x3d, 0x69, 0x93,
	0x8c, 0x34, 0x4d, 0x9b, 0xa3, 0x92, 0x5d, 0x83, 0xcd, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x64,
	0x82, 0xee, 0x41, 0x55, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SignupCommandServiceClient is the client API for SignupCommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SignupCommandServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type signupCommandServiceClient struct {
	cc *grpc.ClientConn
}

func NewSignupCommandServiceClient(cc *grpc.ClientConn) SignupCommandServiceClient {
	return &signupCommandServiceClient{cc}
}

func (c *signupCommandServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/oswee.identity.signup.v1.SignupCommandService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignupCommandServiceServer is the server API for SignupCommandService service.
type SignupCommandServiceServer interface {
	Create(context.Context, *CreateRequest) (*empty.Empty, error)
}

// UnimplementedSignupCommandServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSignupCommandServiceServer struct {
}

func (*UnimplementedSignupCommandServiceServer) Create(ctx context.Context, req *CreateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterSignupCommandServiceServer(s *grpc.Server, srv SignupCommandServiceServer) {
	s.RegisterService(&_SignupCommandService_serviceDesc, srv)
}

func _SignupCommandService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupCommandServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oswee.identity.signup.v1.SignupCommandService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupCommandServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SignupCommandService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "oswee.identity.signup.v1.SignupCommandService",
	HandlerType: (*SignupCommandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SignupCommandService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oswee/identity/signup/proto/v1/signup_command_api.proto",
}
