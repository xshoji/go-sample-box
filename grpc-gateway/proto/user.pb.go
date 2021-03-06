// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type UserServiceResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Count                int64    `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Users                []*User  `protobuf:"bytes,4,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserServiceResponse) Reset()         { *m = UserServiceResponse{} }
func (m *UserServiceResponse) String() string { return proto.CompactTextString(m) }
func (*UserServiceResponse) ProtoMessage()    {}
func (*UserServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{0}
}

func (m *UserServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserServiceResponse.Unmarshal(m, b)
}
func (m *UserServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserServiceResponse.Marshal(b, m, deterministic)
}
func (m *UserServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserServiceResponse.Merge(m, src)
}
func (m *UserServiceResponse) XXX_Size() int {
	return xxx_messageInfo_UserServiceResponse.Size(m)
}
func (m *UserServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserServiceResponse proto.InternalMessageInfo

func (m *UserServiceResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UserServiceResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UserServiceResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *UserServiceResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type User struct {
	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to AgeOptional:
	//	*User_Age
	AgeOptional          isUser_AgeOptional `protobuf_oneof:"age_optional"`
	Tasks                []*Task            `protobuf:"bytes,4,rep,name=tasks,proto3" json:"tasks,omitempty"`
	ClearTasks           bool               `protobuf:"varint,5,opt,name=clearTasks,proto3" json:"clearTasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isUser_AgeOptional interface {
	isUser_AgeOptional()
}

type User_Age struct {
	Age int64 `protobuf:"varint,3,opt,name=age,proto3,oneof"`
}

func (*User_Age) isUser_AgeOptional() {}

func (m *User) GetAgeOptional() isUser_AgeOptional {
	if m != nil {
		return m.AgeOptional
	}
	return nil
}

func (m *User) GetAge() int64 {
	if x, ok := m.GetAgeOptional().(*User_Age); ok {
		return x.Age
	}
	return 0
}

func (m *User) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *User) GetClearTasks() bool {
	if m != nil {
		return m.ClearTasks
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*User) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*User_Age)(nil),
	}
}

type UserServiceSelector struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserServiceSelector) Reset()         { *m = UserServiceSelector{} }
func (m *UserServiceSelector) String() string { return proto.CompactTextString(m) }
func (*UserServiceSelector) ProtoMessage()    {}
func (*UserServiceSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{2}
}

func (m *UserServiceSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserServiceSelector.Unmarshal(m, b)
}
func (m *UserServiceSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserServiceSelector.Marshal(b, m, deterministic)
}
func (m *UserServiceSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserServiceSelector.Merge(m, src)
}
func (m *UserServiceSelector) XXX_Size() int {
	return xxx_messageInfo_UserServiceSelector.Size(m)
}
func (m *UserServiceSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_UserServiceSelector.DiscardUnknown(m)
}

var xxx_messageInfo_UserServiceSelector proto.InternalMessageInfo

func (m *UserServiceSelector) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*UserServiceResponse)(nil), "proto.UserServiceResponse")
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*UserServiceSelector)(nil), "proto.UserServiceSelector")
}

func init() { proto.RegisterFile("proto/user.proto", fileDescriptor_d570e3e37e5899c5) }

var fileDescriptor_d570e3e37e5899c5 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xdd, 0x8a, 0x95, 0x50,
	0x14, 0xce, 0xe3, 0xd1, 0x69, 0x96, 0x21, 0xd3, 0x9a, 0xe1, 0x60, 0x16, 0x61, 0x42, 0x20, 0x73,
	0xa1, 0x34, 0xdd, 0x75, 0xd7, 0x1f, 0x75, 0xd1, 0x45, 0xec, 0x99, 0x89, 0xee, 0x62, 0x8f, 0xae,
	0x44, 0xc6, 0xe3, 0x16, 0xf7, 0x76, 0x20, 0x22, 0x88, 0xde, 0x20, 0x7a, 0x89, 0xde, 0xa7, 0x57,
	0xe8, 0x41, 0x62, 0x6f, 0xb5, 0xcc, 0x62, 0x20, 0xba, 0x72, 0x7f, 0xeb, 0xe7, 0xfb, 0xd6, 0xb7,
	0x5c, 0xb0, 0xd7, 0x76, 0x42, 0x89, 0xac, 0x97, 0xd4, 0xa5, 0xe6, 0x89, 0x8e, 0xf9, 0x84, 0xb7,
	0x4a, 0x21, 0xca, 0x9a, 0x32, 0xde, 0x56, 0x19, 0x6f, 0x1a, 0xa1, 0xb8, 0xaa, 0x44, 0x23, 0x87,
	0xa2, 0xf0, 0xe6, 0x98, 0x35, 0xe8, 0xac, 0x7f, 0x9b, 0xd1, 0xb6, 0x55, 0xef, 0xc6, 0xe4, 0xc8,
	0xa9, 0xb8, 0x3c, 0x1f, 0x22, 0xf1, 0x47, 0x0b, 0xf6, 0x4f, 0x25, 0x75, 0xc7, 0xd4, 0x5d, 0x54,
	0x39, 0x31, 0x92, 0xad, 0x68, 0x24, 0xe1, 0x06, 0x5c, 0xa9, 0xb8, 0xea, 0x65, 0x60, 0x45, 0x56,
	0x62, 0xb3, 0x11, 0x61, 0x00, 0x3b, 0x5b, 0x92, 0x92, 0x97, 0x14, 0xac, 0x22, 0x2b, 0xd9, 0x65,
	0x13, 0xc4, 0x03, 0x70, 0x72, 0xd1, 0x37, 0x2a, 0xb0, 0x4d, 0xc3, 0x00, 0xf0, 0x0e, 0x38, 0xda,
	0x81, 0x0c, 0xd6, 0x91, 0x9d, 0x78, 0x47, 0xde, 0x20, 0x9b, 0x6a, 0x49, 0x36, 0x64, 0xe2, 0xcf,
	0x16, 0xac, 0x35, 0x46, 0x1f, 0x56, 0x55, 0x31, 0xea, 0xad, 0xaa, 0x02, 0x11, 0xd6, 0x0d, 0xdf,
	0x4e, 0x42, 0xe6, 0x8d, 0x08, 0xb6, 0xd6, 0x36, 0x1a, 0xcf, 0xaf, 0x30, 0x0d, 0xb4, 0x86, 0x76,
	0xb4, 0xd4, 0x38, 0xe1, 0xf2, 0x9c, 0x0d, 0x19, 0xbc, 0x0d, 0x90, 0xd7, 0xc4, 0xbb, 0x13, 0x53,
	0xe7, 0x44, 0x56, 0x72, 0x95, 0xcd, 0x22, 0x8f, 0x7c, 0xb8, 0xc6, 0x4b, 0x7a, 0x23, 0x5a, 0xbd,
	0x4a, 0x5e, 0xc7, 0x77, 0x7f, 0xdb, 0xca, 0x31, 0xd5, 0x94, 0x2b, 0xf1, 0xc7, 0x84, 0x47, 0x5f,
	0x6d, 0xf0, 0x66, 0x75, 0xf8, 0x0c, 0xdc, 0xc7, 0x1d, 0x71, 0x45, 0x38, 0x37, 0x1a, 0x86, 0x33,
	0xb0, 0x58, 0x74, 0x7c, 0xf0, 0xe9, 0xdb, 0xf7, 0x2f, 0x2b, 0x3f, 0xde, 0xcd, 0x2e, 0xee, 0x99,
	0x9f, 0x2d, 0x1f, 0x58, 0x87, 0xf8, 0x0a, 0xd6, 0x8c, 0x78, 0x81, 0x7f, 0xe9, 0x9c, 0x86, 0xb9,
	0x94, 0x75, 0x63, 0x58, 0xf7, 0xd0, 0xff, 0xc9, 0x9a, 0xbd, 0xaf, 0x8a, 0x0f, 0xf8, 0x12, 0x76,
	0x34, 0xef, 0xc3, 0xba, 0xc6, 0x4d, 0x3a, 0x5c, 0x4a, 0x3a, 0x5d, 0x4a, 0xfa, 0x54, 0x5f, 0xca,
	0xa5, 0xb4, 0xd7, 0x0d, 0xad, 0x87, 0xbf, 0x86, 0xc5, 0x17, 0xe0, 0x9e, 0xb6, 0xc5, 0x3f, 0x59,
	0xbe, 0x61, 0x58, 0xf6, 0xe3, 0xc5, 0x70, 0xda, 0xf7, 0x6b, 0x70, 0x9f, 0x50, 0x4d, 0x8a, 0xfe,
	0xd7, 0xf9, 0xe1, 0x82, 0xfc, 0xcc, 0x35, 0x2d, 0xf7, 0x7f, 0x04, 0x00, 0x00, 0xff, 0xff, 0x77,
	0xe4, 0x25, 0x17, 0x57, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Create(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserServiceResponse, error)
	Read(ctx context.Context, in *UserServiceSelector, opts ...grpc.CallOption) (*UserServiceResponse, error)
	ReadAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserServiceResponse, error)
	Update(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserServiceResponse, error)
	Delete(ctx context.Context, in *UserServiceSelector, opts ...grpc.CallOption) (*UserServiceResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Create(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserServiceResponse, error) {
	out := new(UserServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Read(ctx context.Context, in *UserServiceSelector, opts ...grpc.CallOption) (*UserServiceResponse, error) {
	out := new(UserServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ReadAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserServiceResponse, error) {
	out := new(UserServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/ReadAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserServiceResponse, error) {
	out := new(UserServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *UserServiceSelector, opts ...grpc.CallOption) (*UserServiceResponse, error) {
	out := new(UserServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Create(context.Context, *User) (*UserServiceResponse, error)
	Read(context.Context, *UserServiceSelector) (*UserServiceResponse, error)
	ReadAll(context.Context, *empty.Empty) (*UserServiceResponse, error)
	Update(context.Context, *User) (*UserServiceResponse, error)
	Delete(context.Context, *UserServiceSelector) (*UserServiceResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) Create(ctx context.Context, req *User) (*UserServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedUserServiceServer) Read(ctx context.Context, req *UserServiceSelector) (*UserServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedUserServiceServer) ReadAll(ctx context.Context, req *empty.Empty) (*UserServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAll not implemented")
}
func (*UnimplementedUserServiceServer) Update(ctx context.Context, req *User) (*UserServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedUserServiceServer) Delete(ctx context.Context, req *UserServiceSelector) (*UserServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserServiceSelector)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Read(ctx, req.(*UserServiceSelector))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/ReadAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ReadAll(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserServiceSelector)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*UserServiceSelector))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _UserService_Read_Handler,
		},
		{
			MethodName: "ReadAll",
			Handler:    _UserService_ReadAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user.proto",
}
