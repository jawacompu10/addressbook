// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addressbook.proto

package addressbook

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AddressID struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressID) Reset()         { *m = AddressID{} }
func (m *AddressID) String() string { return proto.CompactTextString(m) }
func (*AddressID) ProtoMessage()    {}
func (*AddressID) Descriptor() ([]byte, []int) {
	return fileDescriptor_1eb1a68c9dd6d429, []int{0}
}

func (m *AddressID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressID.Unmarshal(m, b)
}
func (m *AddressID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressID.Marshal(b, m, deterministic)
}
func (m *AddressID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressID.Merge(m, src)
}
func (m *AddressID) XXX_Size() int {
	return xxx_messageInfo_AddressID.Size(m)
}
func (m *AddressID) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressID.DiscardUnknown(m)
}

var xxx_messageInfo_AddressID proto.InternalMessageInfo

func (m *AddressID) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type UserID struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserID) Reset()         { *m = UserID{} }
func (m *UserID) String() string { return proto.CompactTextString(m) }
func (*UserID) ProtoMessage()    {}
func (*UserID) Descriptor() ([]byte, []int) {
	return fileDescriptor_1eb1a68c9dd6d429, []int{1}
}

func (m *UserID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserID.Unmarshal(m, b)
}
func (m *UserID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserID.Marshal(b, m, deterministic)
}
func (m *UserID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserID.Merge(m, src)
}
func (m *UserID) XXX_Size() int {
	return xxx_messageInfo_UserID.Size(m)
}
func (m *UserID) XXX_DiscardUnknown() {
	xxx_messageInfo_UserID.DiscardUnknown(m)
}

var xxx_messageInfo_UserID proto.InternalMessageInfo

func (m *UserID) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type Address struct {
	ID                   string               `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID               string               `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Name                 string               `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	IsDefault            bool                 `protobuf:"varint,4,opt,name=IsDefault,proto3" json:"IsDefault,omitempty"`
	Note                 string               `protobuf:"bytes,5,opt,name=Note,proto3" json:"Note,omitempty"`
	Lat                  string               `protobuf:"bytes,6,opt,name=Lat,proto3" json:"Lat,omitempty"`
	Long                 string               `protobuf:"bytes,7,opt,name=Long,proto3" json:"Long,omitempty"`
	Details              map[string]string    `protobuf:"bytes,8,rep,name=Details,proto3" json:"Details,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,9,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	ModifiedAt           *timestamp.Timestamp `protobuf:"bytes,10,opt,name=ModifiedAt,proto3" json:"ModifiedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_1eb1a68c9dd6d429, []int{2}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Address) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Address) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Address) GetIsDefault() bool {
	if m != nil {
		return m.IsDefault
	}
	return false
}

func (m *Address) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func (m *Address) GetLat() string {
	if m != nil {
		return m.Lat
	}
	return ""
}

func (m *Address) GetLong() string {
	if m != nil {
		return m.Long
	}
	return ""
}

func (m *Address) GetDetails() map[string]string {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *Address) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Address) GetModifiedAt() *timestamp.Timestamp {
	if m != nil {
		return m.ModifiedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*AddressID)(nil), "addressbook.AddressID")
	proto.RegisterType((*UserID)(nil), "addressbook.UserID")
	proto.RegisterType((*Address)(nil), "addressbook.Address")
	proto.RegisterMapType((map[string]string)(nil), "addressbook.Address.DetailsEntry")
}

func init() { proto.RegisterFile("addressbook.proto", fileDescriptor_1eb1a68c9dd6d429) }

var fileDescriptor_1eb1a68c9dd6d429 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0xc6, 0x2b, 0xc9, 0xff, 0x34, 0x6a, 0x8b, 0xbb, 0x35, 0x66, 0x51, 0x0b, 0x55, 0x7d, 0xd2,
	0x49, 0x2e, 0xee, 0xc5, 0xb8, 0xcd, 0x41, 0x89, 0x82, 0x11, 0x38, 0x39, 0x88, 0xf8, 0x01, 0xd6,
	0x68, 0x6c, 0x84, 0x65, 0xaf, 0x91, 0xd6, 0x01, 0x3f, 0x42, 0x20, 0x0f, 0x1d, 0xb4, 0x5a, 0x61,
	0x25, 0xc8, 0x04, 0x72, 0x1b, 0x7d, 0xf3, 0xfb, 0x56, 0x33, 0xdf, 0xc0, 0x37, 0x16, 0xc7, 0x19,
	0xe6, 0xf9, 0x8a, 0xf3, 0xad, 0x77, 0xc8, 0xb8, 0xe0, 0xc4, 0xaa, 0x49, 0xf6, 0xaf, 0x0d, 0xe7,
	0x9b, 0x14, 0xc7, 0xb2, 0xb5, 0x3a, 0xae, 0xc7, 0x22, 0xd9, 0x61, 0x2e, 0xd8, 0xee, 0x50, 0xd2,
	0xa3, 0x1f, 0x60, 0xfa, 0x25, 0x1f, 0x06, 0xe4, 0x2b, 0xe8, 0x61, 0x40, 0x35, 0x47, 0x73, 0xcd,
	0x48, 0x0f, 0x83, 0x11, 0x85, 0xce, 0x32, 0xc7, 0xac, 0xa1, 0xf3, 0x6c, 0x40, 0x57, 0xf9, 0xde,
	0xf6, 0xc8, 0xb0, 0x72, 0x51, 0x5d, 0x6a, 0xd5, 0x1b, 0x04, 0x5a, 0xf7, 0x6c, 0x87, 0xd4, 0x90,
	0xaa, 0xac, 0xc9, 0x4f, 0x30, 0xc3, 0x3c, 0xc0, 0x35, 0x3b, 0xa6, 0x82, 0xb6, 0x1c, 0xcd, 0xed,
	0x45, 0x67, 0x41, 0x3a, 0xb8, 0x40, 0xda, 0x56, 0x0e, 0x2e, 0x90, 0xf4, 0xc1, 0x58, 0x30, 0x41,
	0x3b, 0x52, 0x2a, 0xca, 0x82, 0x5a, 0xf0, 0xfd, 0x86, 0x76, 0x4b, 0xaa, 0xa8, 0xc9, 0x3f, 0xe8,
	0x06, 0x28, 0x58, 0x92, 0xe6, 0xb4, 0xe7, 0x18, 0xae, 0x35, 0xf9, 0xed, 0xd5, 0x93, 0x52, 0xa3,
	0x7b, 0x8a, 0xb9, 0xdd, 0x8b, 0xec, 0x14, 0x55, 0x0e, 0x32, 0x05, 0xf3, 0x26, 0x43, 0x26, 0x30,
	0xf6, 0x05, 0x35, 0x1d, 0xcd, 0xb5, 0x26, 0xb6, 0x57, 0x06, 0xe9, 0x55, 0x41, 0x7a, 0x0f, 0x55,
	0x90, 0xd1, 0x19, 0x26, 0x33, 0x80, 0x3b, 0x1e, 0x27, 0xeb, 0x44, 0x5a, 0xe1, 0x5d, 0x6b, 0x8d,
	0xb6, 0x67, 0xf0, 0xb9, 0x3e, 0x4e, 0xb1, 0xe8, 0x16, 0x4f, 0x2a, 0xd7, 0xa2, 0x24, 0x03, 0x68,
	0x3f, 0xb2, 0xf4, 0x88, 0x2a, 0xd7, 0xf2, 0x63, 0xa6, 0x4f, 0xb5, 0xc9, 0x93, 0x0e, 0x96, 0xda,
	0xe9, 0x9a, 0xf3, 0x2d, 0xf9, 0x0f, 0x30, 0x47, 0x51, 0x1d, 0x68, 0xd8, 0xb4, 0x7b, 0x18, 0xd8,
	0x83, 0x26, 0x7d, 0xf4, 0x89, 0xf8, 0xd0, 0x9f, 0xa3, 0x28, 0xae, 0xa6, 0x34, 0xcc, 0xc9, 0xf7,
	0x57, 0x6c, 0x79, 0xd1, 0x4b, 0x0f, 0xfc, 0xd1, 0x8a, 0x20, 0xfc, 0x38, 0xae, 0x06, 0x68, 0xe4,
	0x2e, 0xfe, 0xfe, 0x0a, 0xbe, 0x2c, 0x0f, 0x31, 0x13, 0xf8, 0x21, 0xfb, 0xaa, 0x23, 0x73, 0xfe,
	0xfb, 0x12, 0x00, 0x00, 0xff, 0xff, 0x99, 0x01, 0x30, 0xa2, 0x1b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AddressBookClient is the client API for AddressBook service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddressBookClient interface {
	// Get an address with the specified ID
	GetAddress(ctx context.Context, in *AddressID, opts ...grpc.CallOption) (*Address, error)
	// Get all addresses belonging to the specified user
	GetUserAddresses(ctx context.Context, in *UserID, opts ...grpc.CallOption) (AddressBook_GetUserAddressesClient, error)
	// Add the given address
	AddAddress(ctx context.Context, in *Address, opts ...grpc.CallOption) (*Address, error)
	// Update an existing address
	UpdateAddress(ctx context.Context, in *Address, opts ...grpc.CallOption) (*Address, error)
}

type addressBookClient struct {
	cc grpc.ClientConnInterface
}

func NewAddressBookClient(cc grpc.ClientConnInterface) AddressBookClient {
	return &addressBookClient{cc}
}

func (c *addressBookClient) GetAddress(ctx context.Context, in *AddressID, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, "/addressbook.AddressBook/GetAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressBookClient) GetUserAddresses(ctx context.Context, in *UserID, opts ...grpc.CallOption) (AddressBook_GetUserAddressesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AddressBook_serviceDesc.Streams[0], "/addressbook.AddressBook/GetUserAddresses", opts...)
	if err != nil {
		return nil, err
	}
	x := &addressBookGetUserAddressesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AddressBook_GetUserAddressesClient interface {
	Recv() (*Address, error)
	grpc.ClientStream
}

type addressBookGetUserAddressesClient struct {
	grpc.ClientStream
}

func (x *addressBookGetUserAddressesClient) Recv() (*Address, error) {
	m := new(Address)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *addressBookClient) AddAddress(ctx context.Context, in *Address, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, "/addressbook.AddressBook/AddAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressBookClient) UpdateAddress(ctx context.Context, in *Address, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, "/addressbook.AddressBook/UpdateAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressBookServer is the server API for AddressBook service.
type AddressBookServer interface {
	// Get an address with the specified ID
	GetAddress(context.Context, *AddressID) (*Address, error)
	// Get all addresses belonging to the specified user
	GetUserAddresses(*UserID, AddressBook_GetUserAddressesServer) error
	// Add the given address
	AddAddress(context.Context, *Address) (*Address, error)
	// Update an existing address
	UpdateAddress(context.Context, *Address) (*Address, error)
}

// UnimplementedAddressBookServer can be embedded to have forward compatible implementations.
type UnimplementedAddressBookServer struct {
}

func (*UnimplementedAddressBookServer) GetAddress(ctx context.Context, req *AddressID) (*Address, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddress not implemented")
}
func (*UnimplementedAddressBookServer) GetUserAddresses(req *UserID, srv AddressBook_GetUserAddressesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUserAddresses not implemented")
}
func (*UnimplementedAddressBookServer) AddAddress(ctx context.Context, req *Address) (*Address, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAddress not implemented")
}
func (*UnimplementedAddressBookServer) UpdateAddress(ctx context.Context, req *Address) (*Address, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAddress not implemented")
}

func RegisterAddressBookServer(s *grpc.Server, srv AddressBookServer) {
	s.RegisterService(&_AddressBook_serviceDesc, srv)
}

func _AddressBook_GetAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressBookServer).GetAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addressbook.AddressBook/GetAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressBookServer).GetAddress(ctx, req.(*AddressID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressBook_GetUserAddresses_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserID)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AddressBookServer).GetUserAddresses(m, &addressBookGetUserAddressesServer{stream})
}

type AddressBook_GetUserAddressesServer interface {
	Send(*Address) error
	grpc.ServerStream
}

type addressBookGetUserAddressesServer struct {
	grpc.ServerStream
}

func (x *addressBookGetUserAddressesServer) Send(m *Address) error {
	return x.ServerStream.SendMsg(m)
}

func _AddressBook_AddAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Address)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressBookServer).AddAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addressbook.AddressBook/AddAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressBookServer).AddAddress(ctx, req.(*Address))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressBook_UpdateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Address)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressBookServer).UpdateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addressbook.AddressBook/UpdateAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressBookServer).UpdateAddress(ctx, req.(*Address))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddressBook_serviceDesc = grpc.ServiceDesc{
	ServiceName: "addressbook.AddressBook",
	HandlerType: (*AddressBookServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAddress",
			Handler:    _AddressBook_GetAddress_Handler,
		},
		{
			MethodName: "AddAddress",
			Handler:    _AddressBook_AddAddress_Handler,
		},
		{
			MethodName: "UpdateAddress",
			Handler:    _AddressBook_UpdateAddress_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUserAddresses",
			Handler:       _AddressBook_GetUserAddresses_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "addressbook.proto",
}
