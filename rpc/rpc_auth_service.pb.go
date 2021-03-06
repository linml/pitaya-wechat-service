// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package rpc

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type AuthRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	PhoneNo              string   `protobuf:"bytes,3,opt,name=phoneNo,proto3" json:"phoneNo,omitempty"`
	Code                 string   `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{0}
}

func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (m *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(m, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AuthRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AuthRequest) GetPhoneNo() string {
	if m != nil {
		return m.PhoneNo
	}
	return ""
}

func (m *AuthRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type TokenResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Ttl                  int64    `protobuf:"varint,2,opt,name=ttl,proto3" json:"ttl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenResponse) Reset()         { *m = TokenResponse{} }
func (m *TokenResponse) String() string { return proto.CompactTextString(m) }
func (*TokenResponse) ProtoMessage()    {}
func (*TokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{1}
}

func (m *TokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenResponse.Unmarshal(m, b)
}
func (m *TokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenResponse.Marshal(b, m, deterministic)
}
func (m *TokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenResponse.Merge(m, src)
}
func (m *TokenResponse) XXX_Size() int {
	return xxx_messageInfo_TokenResponse.Size(m)
}
func (m *TokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TokenResponse proto.InternalMessageInfo

func (m *TokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *TokenResponse) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func init() {
	proto.RegisterType((*AuthRequest)(nil), "rpc.AuthRequest")
	proto.RegisterType((*TokenResponse)(nil), "rpc.TokenResponse")
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor_17b8c58d586b62f2) }

var fileDescriptor_17b8c58d586b62f2 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x3f, 0x4b, 0xc0, 0x30,
	0x10, 0xc5, 0xa9, 0xa9, 0x7f, 0x7a, 0x22, 0x94, 0xa3, 0x43, 0xe8, 0x24, 0x9d, 0x9c, 0x0a, 0xea,
	0xe0, 0xec, 0xe8, 0xe2, 0x10, 0xfd, 0x02, 0x35, 0x3d, 0x68, 0x69, 0xcc, 0xc5, 0x24, 0xb5, 0xe0,
	0xa7, 0x97, 0xa4, 0x56, 0x74, 0x7b, 0xef, 0x77, 0xf0, 0xde, 0x3d, 0xa8, 0x27, 0x32, 0x86, 0x37,
	0xf6, 0x66, 0xec, 0x9d, 0xe7, 0xc8, 0x28, 0xbc, 0xd3, 0xdd, 0x02, 0x97, 0x8f, 0x6b, 0x9c, 0x14,
	0x7d, 0xac, 0x14, 0x22, 0x22, 0x94, 0x76, 0x78, 0x27, 0x59, 0x5c, 0x17, 0x37, 0x95, 0xca, 0x1a,
	0x5b, 0xb8, 0x70, 0x43, 0x08, 0x1b, 0xfb, 0x51, 0x9e, 0x64, 0xfe, 0xeb, 0x51, 0xc2, 0xb9, 0x9b,
	0xd8, 0xd2, 0x33, 0x4b, 0x91, 0x4f, 0x87, 0x4d, 0x49, 0x9a, 0x47, 0x92, 0xe5, 0x9e, 0x94, 0x74,
	0xf7, 0x00, 0x57, 0xaf, 0xbc, 0x90, 0x55, 0x14, 0x1c, 0xdb, 0x40, 0xd8, 0xc0, 0x69, 0x4c, 0xe0,
	0xa7, 0x6f, 0x37, 0x58, 0x83, 0x88, 0xd1, 0xe4, 0x2e, 0xa1, 0x92, 0xbc, 0x7b, 0x82, 0x26, 0x7d,
	0xc9, 0x7e, 0xfe, 0x1a, 0xe2, 0xcc, 0xf6, 0x85, 0xfc, 0xe7, 0xac, 0x09, 0x6f, 0xa1, 0x3a, 0x38,
	0x61, 0xdd, 0x7b, 0xa7, 0xfb, 0x3f, 0x6b, 0x5a, 0xcc, 0xe4, 0x5f, 0xe5, 0xdb, 0x59, 0x1e, 0x7f,
	0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xc3, 0xca, 0x5f, 0x10, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthorizationServiceClient is the client API for AuthorizationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizationServiceClient interface {
	Authorize(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*TokenResponse, error)
}

type authorizationServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizationServiceClient(cc *grpc.ClientConn) AuthorizationServiceClient {
	return &authorizationServiceClient{cc}
}

func (c *authorizationServiceClient) Authorize(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, "/rpc.AuthorizationService/Authorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServiceServer is the server API for AuthorizationService service.
type AuthorizationServiceServer interface {
	Authorize(context.Context, *AuthRequest) (*TokenResponse, error)
}

func RegisterAuthorizationServiceServer(s *grpc.Server, srv AuthorizationServiceServer) {
	s.RegisterService(&_AuthorizationService_serviceDesc, srv)
}

func _AuthorizationService_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.AuthorizationService/Authorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).Authorize(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthorizationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.AuthorizationService",
	HandlerType: (*AuthorizationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authorize",
			Handler:    _AuthorizationService_Authorize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}
