// Code generated by protoc-gen-go.
// source: auth.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	auth.proto
	users.proto

It has these top-level messages:
	LoginRequest
	LoginResponse
	RefreshRequest
	RefreshResponse
	Page
	User
	GetAllResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

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

type LoginRequest struct {
	Username     string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	PasswordHash string `protobuf:"bytes,2,opt,name=password_hash,json=passwordHash" json:"password_hash,omitempty"`
	DeviceName   string `protobuf:"bytes,3,opt,name=device_name,json=deviceName" json:"device_name,omitempty"`
	IsMobile     bool   `protobuf:"varint,4,opt,name=is_mobile,json=isMobile" json:"is_mobile,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPasswordHash() string {
	if m != nil {
		return m.PasswordHash
	}
	return ""
}

func (m *LoginRequest) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *LoginRequest) GetIsMobile() bool {
	if m != nil {
		return m.IsMobile
	}
	return false
}

type LoginResponse struct {
	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *LoginResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type RefreshRequest struct {
	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
}

func (m *RefreshRequest) Reset()                    { *m = RefreshRequest{} }
func (m *RefreshRequest) String() string            { return proto.CompactTextString(m) }
func (*RefreshRequest) ProtoMessage()               {}
func (*RefreshRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RefreshRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *RefreshRequest) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type RefreshResponse struct {
	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
}

func (m *RefreshResponse) Reset()                    { *m = RefreshResponse{} }
func (m *RefreshResponse) String() string            { return proto.CompactTextString(m) }
func (*RefreshResponse) ProtoMessage()               {}
func (*RefreshResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RefreshResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *RefreshResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "pb.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "pb.LoginResponse")
	proto.RegisterType((*RefreshRequest)(nil), "pb.RefreshRequest")
	proto.RegisterType((*RefreshResponse)(nil), "pb.RefreshResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Auth service

type AuthClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/pb.Auth/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error) {
	out := new(RefreshResponse)
	err := grpc.Invoke(ctx, "/pb.Auth/Refresh", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Auth/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Auth/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Refresh(ctx, req.(*RefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _Auth_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x92, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x95, 0x52, 0x20, 0xbd, 0xb6, 0xfc, 0x71, 0x97, 0x2a, 0x20, 0x51, 0xc2, 0x52, 0x31,
	0x34, 0x02, 0x36, 0xb6, 0x4e, 0x30, 0x00, 0x43, 0x84, 0x04, 0x4c, 0x91, 0xd3, 0x1e, 0x8d, 0x45,
	0x6b, 0x9b, 0x9c, 0x53, 0x76, 0x76, 0x16, 0xf8, 0x68, 0x7c, 0x05, 0x3e, 0x08, 0x8a, 0xe3, 0x56,
	0xed, 0xde, 0xf1, 0xfd, 0xf4, 0xfc, 0xfc, 0x7c, 0x3e, 0x00, 0x5e, 0x98, 0x6c, 0xa0, 0x73, 0x65,
	0x14, 0xab, 0xe9, 0x34, 0x38, 0x9e, 0x28, 0x35, 0x99, 0x62, 0xc4, 0xb5, 0x88, 0xb8, 0x94, 0xca,
	0x70, 0x23, 0x94, 0xa4, 0xca, 0x11, 0x7e, 0x79, 0xd0, 0xba, 0x53, 0x13, 0x21, 0x63, 0x7c, 0x2f,
	0x90, 0x0c, 0x0b, 0xc0, 0x2f, 0x08, 0x73, 0xc9, 0x67, 0xd8, 0xf5, 0x7a, 0x5e, 0xbf, 0x11, 0x2f,
	0x35, 0x3b, 0x83, 0xb6, 0xe6, 0x44, 0x1f, 0x2a, 0x1f, 0x27, 0x19, 0xa7, 0xac, 0x5b, 0xb3, 0x86,
	0xd6, 0x02, 0xde, 0x72, 0xca, 0xd8, 0x09, 0x34, 0xc7, 0x38, 0x17, 0x23, 0x4c, 0x6c, 0xc6, 0x96,
	0xb5, 0x40, 0x85, 0x1e, 0xca, 0x94, 0x23, 0x68, 0x08, 0x4a, 0x66, 0x2a, 0x15, 0x53, 0xec, 0xd6,
	0x7b, 0x5e, 0xdf, 0x8f, 0x7d, 0x41, 0xf7, 0x56, 0x87, 0x4f, 0xd0, 0x76, 0x75, 0x48, 0x2b, 0x49,
	0xc8, 0x4e, 0xa1, 0xc5, 0x47, 0x23, 0x24, 0x4a, 0x8c, 0x7a, 0x43, 0xe9, 0x3a, 0x35, 0x2b, 0xf6,
	0x58, 0xa2, 0xb2, 0x56, 0x8e, 0xaf, 0x39, 0x52, 0xe6, 0x3c, 0xae, 0x96, 0x83, 0xd6, 0x14, 0x3e,
	0xc3, 0x5e, 0x5c, 0xe9, 0xc5, 0x4b, 0x37, 0x95, 0xfc, 0x02, 0xfb, 0xcb, 0xe4, 0xcd, 0x96, 0xbe,
	0xfc, 0xf6, 0xa0, 0x3e, 0x2c, 0x4c, 0xc6, 0x86, 0xb0, 0x6d, 0xc7, 0xc2, 0x0e, 0x06, 0x3a, 0x1d,
	0xac, 0x7e, 0x58, 0x70, 0xb8, 0x42, 0xaa, 0xeb, 0xc3, 0xce, 0xe7, 0xef, 0xdf, 0x4f, 0xad, 0x1d,
	0xfa, 0xd1, 0xfc, 0x22, 0x2a, 0xd7, 0xe1, 0xda, 0x3b, 0x67, 0x37, 0xb0, 0xeb, 0x6a, 0x32, 0x56,
	0x1e, 0x59, 0x9f, 0x46, 0xd0, 0x59, 0x63, 0xeb, 0x41, 0xc1, 0x6a, 0x50, 0xba, 0x63, 0x37, 0xe7,
	0xea, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x32, 0x33, 0x24, 0x3a, 0x69, 0x02, 0x00, 0x00,
}