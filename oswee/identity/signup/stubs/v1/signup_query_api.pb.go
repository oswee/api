// Code generated by protoc-gen-go. DO NOT EDIT.
// source: oswee/identity/signup/proto/v1/signup_query_api.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	v1 "oswee/identity/signup/proto/v1"
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

type ListSignupsRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListSignupsRequest) Reset()         { *m = ListSignupsRequest{} }
func (m *ListSignupsRequest) String() string { return proto.CompactTextString(m) }
func (*ListSignupsRequest) ProtoMessage()    {}
func (*ListSignupsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9da65fc8fde16cfb, []int{0}
}

func (m *ListSignupsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSignupsRequest.Unmarshal(m, b)
}
func (m *ListSignupsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSignupsRequest.Marshal(b, m, deterministic)
}
func (m *ListSignupsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSignupsRequest.Merge(m, src)
}
func (m *ListSignupsRequest) XXX_Size() int {
	return xxx_messageInfo_ListSignupsRequest.Size(m)
}
func (m *ListSignupsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSignupsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListSignupsRequest proto.InternalMessageInfo

func (m *ListSignupsRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

type ListSignupsResponse struct {
	Api                  string       `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Signups              []*v1.Signup `protobuf:"bytes,2,rep,name=signups,proto3" json:"signups,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListSignupsResponse) Reset()         { *m = ListSignupsResponse{} }
func (m *ListSignupsResponse) String() string { return proto.CompactTextString(m) }
func (*ListSignupsResponse) ProtoMessage()    {}
func (*ListSignupsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9da65fc8fde16cfb, []int{1}
}

func (m *ListSignupsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSignupsResponse.Unmarshal(m, b)
}
func (m *ListSignupsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSignupsResponse.Marshal(b, m, deterministic)
}
func (m *ListSignupsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSignupsResponse.Merge(m, src)
}
func (m *ListSignupsResponse) XXX_Size() int {
	return xxx_messageInfo_ListSignupsResponse.Size(m)
}
func (m *ListSignupsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSignupsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListSignupsResponse proto.InternalMessageInfo

func (m *ListSignupsResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ListSignupsResponse) GetSignups() []*v1.Signup {
	if m != nil {
		return m.Signups
	}
	return nil
}

type GetSignupRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSignupRequest) Reset()         { *m = GetSignupRequest{} }
func (m *GetSignupRequest) String() string { return proto.CompactTextString(m) }
func (*GetSignupRequest) ProtoMessage()    {}
func (*GetSignupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9da65fc8fde16cfb, []int{2}
}

func (m *GetSignupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSignupRequest.Unmarshal(m, b)
}
func (m *GetSignupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSignupRequest.Marshal(b, m, deterministic)
}
func (m *GetSignupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSignupRequest.Merge(m, src)
}
func (m *GetSignupRequest) XXX_Size() int {
	return xxx_messageInfo_GetSignupRequest.Size(m)
}
func (m *GetSignupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSignupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSignupRequest proto.InternalMessageInfo

func (m *GetSignupRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *GetSignupRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*ListSignupsRequest)(nil), "oswee.identity.signup.v1.ListSignupsRequest")
	proto.RegisterType((*ListSignupsResponse)(nil), "oswee.identity.signup.v1.ListSignupsResponse")
	proto.RegisterType((*GetSignupRequest)(nil), "oswee.identity.signup.v1.GetSignupRequest")
}

func init() {
	proto.RegisterFile("oswee/identity/signup/proto/v1/signup_query_api.proto", fileDescriptor_9da65fc8fde16cfb)
}

var fileDescriptor_9da65fc8fde16cfb = []byte{
	// 813 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0x96, 0x9d, 0xfe, 0xa8, 0xb3, 0x6d, 0x09, 0x53, 0xd4, 0x46, 0xa6, 0x82, 0x69, 0xae, 0xaa,
	0xd0, 0x8c, 0x77, 0xb3, 0x01, 0x44, 0x10, 0x12, 0xd9, 0x12, 0x56, 0x2b, 0x2d, 0xb4, 0xcd, 0x22,
	0x21, 0xf5, 0xa6, 0x9a, 0xd8, 0x67, 0xed, 0x01, 0x7b, 0xc6, 0x9d, 0x33, 0x4e, 0x1a, 0x10, 0x5c,
	0xf0, 0x08, 0xe1, 0x8e, 0x0b, 0xde, 0x03, 0x81, 0xe0, 0x1d, 0xca, 0x35, 0x77, 0x3c, 0x08, 0xf2,
	0xd8, 0x61, 0x81, 0x6c, 0xd0, 0x5e, 0xd9, 0x73, 0xe6, 0x3b, 0xdf, 0x77, 0xe6, 0x3b, 0x67, 0x86,
	0xbc, 0xad, 0x71, 0x01, 0x10, 0xca, 0x18, 0x94, 0x95, 0x76, 0x19, 0xa2, 0x4c, 0x54, 0x59, 0x84,
	0x85, 0xd1, 0x56, 0x87, 0xf3, 0xbd, 0x66, 0xfd, 0xec, 0x79, 0x09, 0x66, 0xf9, 0x4c, 0x14, 0x92,
	0xbb, 0x1d, 0xda, 0x71, 0x69, 0x7c, 0x9d, 0xc6, 0x6b, 0x18, 0x9f, 0xef, 0x05, 0xaf, 0x27, 0x5a,
	0x27, 0x19, 0xd4, 0x0c, 0xb3, 0xf2, 0x34, 0x84, 0xbc, 0xb0, 0xcb, 0x3a, 0x2d, 0xb8, 0xdb, 0x6c,
	0x8a, 0x42, 0x86, 0x42, 0x29, 0x6d, 0x85, 0x95, 0x5a, 0x61, 0xb3, 0xfb, 0xc0, 0x7d, 0xa2, 0x7e,
	0x02, 0xaa, 0x8f, 0x0b, 0x91, 0x24, 0x60, 0x42, 0x5d, 0x38, 0xc4, 0x39, 0xe8, 0xb7, 0x2e, 0x54,
	0x79, 0x0d, 0xee, 0x3e, 0x25, 0xf4, 0x58, 0xa2, 0x3d, 0x71, 0x31, 0x9c, 0xc2, 0xf3, 0x12, 0xd0,
	0xd2, 0x36, 0x69, 0x89, 0x42, 0x76, 0x3c, 0xe6, 0xdd, 0xbf, 0x36, 0xad, 0x7e, 0x47, 0x83, 0xd5,
	0x38, 0x24, 0xfd, 0xde, 0xf5, 0x0a, 0xcc, 0x1a, 0xd8, 0xe0, 0xce, 0xe7, 0x32, 0xcb, 0x98, 0x01,
	0x5b, 0x1a, 0xc5, 0x44, 0x96, 0xb1, 0x86, 0xe6, 0xa5, 0x57, 0xe5, 0x74, 0x7f, 0xf4, 0xc8, 0xad,
	0x7f, 0x91, 0x63, 0xa1, 0x15, 0xc2, 0x26, 0x3b, 0x1d, 0x91, 0xab, 0x75, 0x55, 0xd8, 0xf1, 0x59,
	0xeb, 0xfe, 0xce, 0x80, 0xf1, 0x6d, 0x3e, 0xf2, 0x9a, 0x6d, 0xba, 0x4e, 0x18, 0xbd, 0xb7, 0x1a,
	0xbf, 0x43, 0x86, 0xbd, 0x1b, 0x4d, 0x65, 0xb5, 0xc6, 0xe0, 0xd6, 0xd4, 0x55, 0x85, 0x9b, 0x65,
	0xbd, 0xf4, 0xd6, 0xa9, 0xdd, 0x9c, 0xb4, 0x0f, 0xa1, 0x29, 0x6f, 0xeb, 0xd1, 0xe9, 0x4d, 0xe2,
	0xcb, 0xb8, 0xe3, 0xbb, 0x80, 0x2f, 0xe3, 0xd1, 0xbb, 0xab, 0xf1, 0x90, 0x0c, 0x7a, 0x3b, 0x87,
	0x60, 0x99, 0x69, 0x9c, 0xe8, 0xac, 0xe5, 0x9a, 0x00, 0xc4, 0x8d, 0xe8, 0x5a, 0xd3, 0x97, 0xf1,
	0xe0, 0x97, 0x16, 0xb9, 0x59, 0x47, 0x9f, 0x54, 0x53, 0x33, 0x7e, 0x7c, 0x44, 0xbf, 0x25, 0x3b,
	0xff, 0x70, 0x88, 0x3e, 0xd8, 0x7e, 0xec, 0xcd, 0x2e, 0x05, 0xfd, 0x0b, 0xa2, 0x6b, 0x4b, 0xba,
	0xf4, 0xbb, 0xdf, 0xff, 0xfc, 0xde, 0xbf, 0x4e, 0xc9, 0xd9, 0x10, 0xd0, 0x94, 0x5c, 0xfb, 0xdb,
	0x01, 0xda, 0xdb, 0xce, 0xf7, 0x5f, 0x9b, 0x82, 0xdb, 0xbc, 0x9e, 0x58, 0xbe, 0x1e, 0x67, 0x3e,
	0xa9, 0xc6, 0xb9, 0x7b, 0xc7, 0x89, 0xbc, 0x4a, 0x5f, 0x39, 0x13, 0x09, 0xbf, 0x96, 0xf1, 0x37,
	0xc1, 0x6f, 0xde, 0x6a, 0xfc, 0xb3, 0x47, 0x4f, 0xc7, 0x07, 0xd2, 0x3e, 0x3a, 0x9d, 0xcc, 0xc1,
	0x2c, 0x6d, 0x2a, 0x55, 0x72, 0x02, 0x66, 0x2e, 0x23, 0x60, 0x31, 0x60, 0x64, 0xa4, 0x1b, 0x6a,
	0xd6, 0xef, 0xb3, 0x45, 0x2a, 0xa3, 0x94, 0x61, 0xaa, 0xcb, 0x2c, 0x66, 0x4a, 0x5b, 0x36, 0x03,
	0x56, 0x22, 0xc4, 0x4c, 0x2a, 0x56, 0x64, 0x22, 0x02, 0xa6, 0x4f, 0x99, 0x4d, 0x81, 0xc5, 0x3a,
	0x2a, 0x73, 0x50, 0xf5, 0x15, 0x60, 0x91, 0xce, 0xab, 0xc5, 0xbd, 0xe0, 0x09, 0x79, 0xf3, 0x63,
	0xa9, 0x62, 0xa6, 0x4b, 0xcb, 0x72, 0x6d, 0x80, 0x89, 0x59, 0xf5, 0x3b, 0x89, 0x52, 0xdd, 0x28,
	0x52, 0x9e, 0x5a, 0x5b, 0xe0, 0x28, 0x0c, 0x13, 0x69, 0xd3, 0x72, 0xc6, 0x23, 0x9d, 0x87, 0x89,
	0x29, 0xa2, 0x3e, 0x44, 0x1a, 0x97, 0x68, 0xa1, 0x59, 0x26, 0xc2, 0xc2, 0x42, 0x2c, 0x0f, 0xfe,
	0x68, 0xad, 0xc6, 0x3f, 0xb5, 0xe8, 0xaf, 0x1e, 0x69, 0x3f, 0xaa, 0x7c, 0x6a, 0x3a, 0xcc, 0xaa,
	0x3e, 0xde, 0xfe, 0x2c, 0x05, 0x96, 0x18, 0x10, 0x16, 0xd0, 0x56, 0x11, 0x06, 0x73, 0x30, 0xf7,
	0xba, 0x1f, 0x90, 0xcb, 0x0e, 0x4b, 0x3b, 0xe7, 0x88, 0x39, 0xb7, 0x83, 0x20, 0xfe, 0x4a, 0x2a,
	0x2b, 0x0c, 0xf2, 0x2f, 0x33, 0x31, 0x97, 0x0a, 0x3f, 0x4c, 0x72, 0x21, 0xb3, 0x0a, 0xd1, 0x3b,
	0x26, 0xaf, 0x1d, 0x9c, 0x7c, 0xc4, 0xf6, 0xfb, 0x0f, 0x33, 0x51, 0x22, 0xb0, 0x63, 0x19, 0x41,
	0x75, 0x83, 0x86, 0xdb, 0xd8, 0xdc, 0x03, 0x32, 0xcb, 0xf4, 0x2c, 0xcc, 0x05, 0x5a, 0x30, 0xe1,
	0xf1, 0xd1, 0xc3, 0xc9, 0xa7, 0x27, 0x13, 0x6e, 0x5f, 0xd8, 0xc1, 0xe5, 0x3d, 0xbe, 0xcb, 0x77,
	0x83, 0x1b, 0xd5, 0x6b, 0x55, 0x77, 0xba, 0xd2, 0xf0, 0xfd, 0x4b, 0x83, 0xb6, 0x28, 0x8a, 0x4c,
	0x46, 0xce, 0xc3, 0xf0, 0x0b, 0xd4, 0x6a, 0xb4, 0x11, 0x99, 0x3e, 0x26, 0xad, 0xe1, 0xee, 0x3e,
	0x3d, 0x22, 0x87, 0xf5, 0xac, 0x43, 0xcc, 0x16, 0x29, 0x28, 0xd7, 0x84, 0x12, 0xc1, 0xb0, 0x58,
	0x03, 0xba, 0x56, 0xa5, 0x62, 0x0e, 0xac, 0x00, 0x93, 0x4b, 0xc4, 0xaa, 0x29, 0x56, 0x33, 0x11,
	0x45, 0x80, 0xe8, 0xb0, 0x06, 0x50, 0x97, 0x26, 0x02, 0x3e, 0x7d, 0xbf, 0x62, 0x1c, 0xd2, 0x21,
	0xe9, 0x6d, 0x32, 0xae, 0x51, 0x67, 0xac, 0xf0, 0x42, 0xa2, 0xe5, 0xf4, 0x0a, 0xb9, 0xf4, 0x83,
	0xef, 0x5d, 0x35, 0xbb, 0xa4, 0xfd, 0xc9, 0x59, 0x5f, 0x6b, 0x93, 0xef, 0xfe, 0x9f, 0x2d, 0x4f,
	0xdf, 0x38, 0xff, 0x5d, 0x44, 0x5b, 0xce, 0x30, 0x9c, 0xef, 0xcd, 0xae, 0xb8, 0x51, 0xde, 0xff,
	0x2b, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xe7, 0xa5, 0x18, 0xfa, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SignupQueryAPIClient is the client API for SignupQueryAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SignupQueryAPIClient interface {
	ListSignups(ctx context.Context, in *ListSignupsRequest, opts ...grpc.CallOption) (*ListSignupsResponse, error)
	GetSignup(ctx context.Context, in *GetSignupRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type signupQueryAPIClient struct {
	cc *grpc.ClientConn
}

func NewSignupQueryAPIClient(cc *grpc.ClientConn) SignupQueryAPIClient {
	return &signupQueryAPIClient{cc}
}

func (c *signupQueryAPIClient) ListSignups(ctx context.Context, in *ListSignupsRequest, opts ...grpc.CallOption) (*ListSignupsResponse, error) {
	out := new(ListSignupsResponse)
	err := c.cc.Invoke(ctx, "/oswee.identity.signup.v1.SignupQueryAPI/ListSignups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signupQueryAPIClient) GetSignup(ctx context.Context, in *GetSignupRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/oswee.identity.signup.v1.SignupQueryAPI/GetSignup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignupQueryAPIServer is the server API for SignupQueryAPI service.
type SignupQueryAPIServer interface {
	ListSignups(context.Context, *ListSignupsRequest) (*ListSignupsResponse, error)
	GetSignup(context.Context, *GetSignupRequest) (*empty.Empty, error)
}

// UnimplementedSignupQueryAPIServer can be embedded to have forward compatible implementations.
type UnimplementedSignupQueryAPIServer struct {
}

func (*UnimplementedSignupQueryAPIServer) ListSignups(ctx context.Context, req *ListSignupsRequest) (*ListSignupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSignups not implemented")
}
func (*UnimplementedSignupQueryAPIServer) GetSignup(ctx context.Context, req *GetSignupRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSignup not implemented")
}

func RegisterSignupQueryAPIServer(s *grpc.Server, srv SignupQueryAPIServer) {
	s.RegisterService(&_SignupQueryAPI_serviceDesc, srv)
}

func _SignupQueryAPI_ListSignups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSignupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupQueryAPIServer).ListSignups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oswee.identity.signup.v1.SignupQueryAPI/ListSignups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupQueryAPIServer).ListSignups(ctx, req.(*ListSignupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignupQueryAPI_GetSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignupQueryAPIServer).GetSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oswee.identity.signup.v1.SignupQueryAPI/GetSignup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignupQueryAPIServer).GetSignup(ctx, req.(*GetSignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SignupQueryAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "oswee.identity.signup.v1.SignupQueryAPI",
	HandlerType: (*SignupQueryAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSignups",
			Handler:    _SignupQueryAPI_ListSignups_Handler,
		},
		{
			MethodName: "GetSignup",
			Handler:    _SignupQueryAPI_GetSignup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oswee/identity/signup/proto/v1/signup_query_api.proto",
}