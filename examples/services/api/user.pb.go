// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type UserRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UserRsp struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRsp) Reset()         { *m = UserRsp{} }
func (m *UserRsp) String() string { return proto.CompactTextString(m) }
func (*UserRsp) ProtoMessage()    {}
func (*UserRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRsp.Unmarshal(m, b)
}
func (m *UserRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRsp.Marshal(b, m, deterministic)
}
func (m *UserRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRsp.Merge(m, src)
}
func (m *UserRsp) XXX_Size() int {
	return xxx_messageInfo_UserRsp.Size(m)
}
func (m *UserRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRsp.DiscardUnknown(m)
}

var xxx_messageInfo_UserRsp proto.InternalMessageInfo

func (m *UserRsp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UserInfoMsg struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoMsg) Reset()         { *m = UserInfoMsg{} }
func (m *UserInfoMsg) String() string { return proto.CompactTextString(m) }
func (*UserInfoMsg) ProtoMessage()    {}
func (*UserInfoMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *UserInfoMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoMsg.Unmarshal(m, b)
}
func (m *UserInfoMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoMsg.Marshal(b, m, deterministic)
}
func (m *UserInfoMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoMsg.Merge(m, src)
}
func (m *UserInfoMsg) XXX_Size() int {
	return xxx_messageInfo_UserInfoMsg.Size(m)
}
func (m *UserInfoMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoMsg.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoMsg proto.InternalMessageInfo

func (m *UserInfoMsg) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfoMsg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type StatusMsg struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusMsg) Reset()         { *m = StatusMsg{} }
func (m *StatusMsg) String() string { return proto.CompactTextString(m) }
func (*StatusMsg) ProtoMessage()    {}
func (*StatusMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *StatusMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusMsg.Unmarshal(m, b)
}
func (m *StatusMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusMsg.Marshal(b, m, deterministic)
}
func (m *StatusMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusMsg.Merge(m, src)
}
func (m *StatusMsg) XXX_Size() int {
	return xxx_messageInfo_StatusMsg.Size(m)
}
func (m *StatusMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusMsg.DiscardUnknown(m)
}

var xxx_messageInfo_StatusMsg proto.InternalMessageInfo

func (m *StatusMsg) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "api.UserRequest")
	proto.RegisterType((*UserRsp)(nil), "api.UserRsp")
	proto.RegisterType((*UserInfoMsg)(nil), "api.UserInfoMsg")
	proto.RegisterType((*StatusMsg)(nil), "api.StatusMsg")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x92, 0xe5, 0xe2, 0x0e,
	0x2d, 0x4e, 0x2d, 0x0a, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe3, 0x62, 0xca, 0x4c,
	0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62, 0xca, 0x4c, 0x51, 0x92, 0xe5, 0x62, 0x07, 0x4b,
	0x17, 0x17, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x82, 0x25, 0x39, 0x83, 0xc0, 0x6c,
	0x25, 0x43, 0x88, 0x6e, 0xcf, 0xbc, 0xb4, 0x7c, 0xdf, 0xe2, 0x74, 0x74, 0xdd, 0x70, 0x2d, 0x4c,
	0x48, 0x5a, 0xe4, 0xb9, 0x38, 0x83, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0x41, 0x1a, 0x84, 0xb8, 0x58,
	0x92, 0xf3, 0x53, 0x52, 0xa1, 0x5a, 0xc0, 0x6c, 0xa3, 0x03, 0x8c, 0x5c, 0x6c, 0xae, 0x79, 0x25,
	0x99, 0x25, 0x95, 0x42, 0x9a, 0x5c, 0xec, 0xee, 0xa9, 0x25, 0x20, 0x1b, 0x84, 0x04, 0xf4, 0x12,
	0x0b, 0x32, 0xf5, 0x90, 0x9c, 0x2a, 0xc5, 0x83, 0x10, 0x29, 0x2e, 0x10, 0xd2, 0xe6, 0x62, 0x77,
	0x4c, 0x49, 0x41, 0x53, 0x0a, 0x75, 0x97, 0x14, 0x1f, 0x58, 0x04, 0x61, 0xad, 0x1e, 0x17, 0x57,
	0x68, 0x41, 0x4a, 0x62, 0x49, 0x2a, 0xf1, 0xea, 0x5d, 0x52, 0x73, 0x52, 0x31, 0xd4, 0xc3, 0x9c,
	0x82, 0xa6, 0x3e, 0x89, 0x0d, 0x1c, 0xc0, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xbd,
	0xc2, 0xca, 0x6e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EntityClient is the client API for Entity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EntityClient interface {
	GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserRsp, error)
	AddUser(ctx context.Context, in *UserInfoMsg, opts ...grpc.CallOption) (*StatusMsg, error)
	UpdateUser(ctx context.Context, in *UserInfoMsg, opts ...grpc.CallOption) (*StatusMsg, error)
	DeleteUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*StatusMsg, error)
}

type entityClient struct {
	cc *grpc.ClientConn
}

func NewEntityClient(cc *grpc.ClientConn) EntityClient {
	return &entityClient{cc}
}

func (c *entityClient) GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserRsp, error) {
	out := new(UserRsp)
	err := c.cc.Invoke(ctx, "/api.Entity/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityClient) AddUser(ctx context.Context, in *UserInfoMsg, opts ...grpc.CallOption) (*StatusMsg, error) {
	out := new(StatusMsg)
	err := c.cc.Invoke(ctx, "/api.Entity/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityClient) UpdateUser(ctx context.Context, in *UserInfoMsg, opts ...grpc.CallOption) (*StatusMsg, error) {
	out := new(StatusMsg)
	err := c.cc.Invoke(ctx, "/api.Entity/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityClient) DeleteUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*StatusMsg, error) {
	out := new(StatusMsg)
	err := c.cc.Invoke(ctx, "/api.Entity/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntityServer is the server API for Entity service.
type EntityServer interface {
	GetUser(context.Context, *UserRequest) (*UserRsp, error)
	AddUser(context.Context, *UserInfoMsg) (*StatusMsg, error)
	UpdateUser(context.Context, *UserInfoMsg) (*StatusMsg, error)
	DeleteUser(context.Context, *UserRequest) (*StatusMsg, error)
}

func RegisterEntityServer(s *grpc.Server, srv EntityServer) {
	s.RegisterService(&_Entity_serviceDesc, srv)
}

func _Entity_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Entity/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityServer).GetUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Entity_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Entity/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityServer).AddUser(ctx, req.(*UserInfoMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Entity_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Entity/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityServer).UpdateUser(ctx, req.(*UserInfoMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Entity_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Entity/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityServer).DeleteUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Entity_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Entity",
	HandlerType: (*EntityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _Entity_GetUser_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _Entity_AddUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Entity_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Entity_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
