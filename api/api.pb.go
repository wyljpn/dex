// Code generated by protoc-gen-go.
// source: api/api.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api/api.proto

It has these top-level messages:
	Client
	CreateClientReq
	CreateClientResp
	DeleteClientReq
	DeleteClientResp
	Password
	CreatePasswordReq
	CreatePasswordResp
	UpdatePasswordReq
	UpdatePasswordResp
	DeletePasswordReq
	DeletePasswordResp
	ListPasswordReq
	ListPasswordResp
	VersionReq
	VersionResp
	RefreshTokenRef
	ListRefreshReq
	ListRefreshResp
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// Client represents an OAuth2 client.
type Client struct {
	Id           string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Secret       string   `protobuf:"bytes,2,opt,name=secret" json:"secret,omitempty"`
	RedirectUris []string `protobuf:"bytes,3,rep,name=redirect_uris,json=redirectUris" json:"redirect_uris,omitempty"`
	TrustedPeers []string `protobuf:"bytes,4,rep,name=trusted_peers,json=trustedPeers" json:"trusted_peers,omitempty"`
	Public       bool     `protobuf:"varint,5,opt,name=public" json:"public,omitempty"`
	Name         string   `protobuf:"bytes,6,opt,name=name" json:"name,omitempty"`
	LogoUrl      string   `protobuf:"bytes,7,opt,name=logo_url,json=logoUrl" json:"logo_url,omitempty"`
}

func (m *Client) Reset()                    { *m = Client{} }
func (m *Client) String() string            { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()               {}
func (*Client) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// CreateClientReq is a request to make a client.
type CreateClientReq struct {
	Client *Client `protobuf:"bytes,1,opt,name=client" json:"client,omitempty"`
}

func (m *CreateClientReq) Reset()                    { *m = CreateClientReq{} }
func (m *CreateClientReq) String() string            { return proto.CompactTextString(m) }
func (*CreateClientReq) ProtoMessage()               {}
func (*CreateClientReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateClientReq) GetClient() *Client {
	if m != nil {
		return m.Client
	}
	return nil
}

// CreateClientResp returns the response from creating a client.
type CreateClientResp struct {
	AlreadyExists bool    `protobuf:"varint,1,opt,name=already_exists,json=alreadyExists" json:"already_exists,omitempty"`
	Client        *Client `protobuf:"bytes,2,opt,name=client" json:"client,omitempty"`
}

func (m *CreateClientResp) Reset()                    { *m = CreateClientResp{} }
func (m *CreateClientResp) String() string            { return proto.CompactTextString(m) }
func (*CreateClientResp) ProtoMessage()               {}
func (*CreateClientResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateClientResp) GetClient() *Client {
	if m != nil {
		return m.Client
	}
	return nil
}

// DeleteClientReq is a request to delete a client.
type DeleteClientReq struct {
	// The ID of the client.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteClientReq) Reset()                    { *m = DeleteClientReq{} }
func (m *DeleteClientReq) String() string            { return proto.CompactTextString(m) }
func (*DeleteClientReq) ProtoMessage()               {}
func (*DeleteClientReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// DeleteClientResp determines if the.
type DeleteClientResp struct {
	NotFound bool `protobuf:"varint,1,opt,name=not_found,json=notFound" json:"not_found,omitempty"`
}

func (m *DeleteClientResp) Reset()                    { *m = DeleteClientResp{} }
func (m *DeleteClientResp) String() string            { return proto.CompactTextString(m) }
func (*DeleteClientResp) ProtoMessage()               {}
func (*DeleteClientResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// Password is an email for password mapping managed by the storage.
type Password struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	// Currently we do not accept plain text passwords. Could be an option in the future.
	Hash     []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username" json:"username,omitempty"`
	UserId   string `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *Password) Reset()                    { *m = Password{} }
func (m *Password) String() string            { return proto.CompactTextString(m) }
func (*Password) ProtoMessage()               {}
func (*Password) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// CreatePasswordReq is a request to make a password.
type CreatePasswordReq struct {
	Password *Password `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
}

func (m *CreatePasswordReq) Reset()                    { *m = CreatePasswordReq{} }
func (m *CreatePasswordReq) String() string            { return proto.CompactTextString(m) }
func (*CreatePasswordReq) ProtoMessage()               {}
func (*CreatePasswordReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CreatePasswordReq) GetPassword() *Password {
	if m != nil {
		return m.Password
	}
	return nil
}

// CreatePasswordResp returns the response from creating a password.
type CreatePasswordResp struct {
	AlreadyExists bool `protobuf:"varint,1,opt,name=already_exists,json=alreadyExists" json:"already_exists,omitempty"`
}

func (m *CreatePasswordResp) Reset()                    { *m = CreatePasswordResp{} }
func (m *CreatePasswordResp) String() string            { return proto.CompactTextString(m) }
func (*CreatePasswordResp) ProtoMessage()               {}
func (*CreatePasswordResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// UpdatePasswordReq is a request to modify an existing password.
type UpdatePasswordReq struct {
	// The email used to lookup the password. This field cannot be modified
	Email       string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	NewHash     []byte `protobuf:"bytes,2,opt,name=new_hash,json=newHash,proto3" json:"new_hash,omitempty"`
	NewUsername string `protobuf:"bytes,3,opt,name=new_username,json=newUsername" json:"new_username,omitempty"`
}

func (m *UpdatePasswordReq) Reset()                    { *m = UpdatePasswordReq{} }
func (m *UpdatePasswordReq) String() string            { return proto.CompactTextString(m) }
func (*UpdatePasswordReq) ProtoMessage()               {}
func (*UpdatePasswordReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// UpdatePasswordResp returns the response from modifying an existing password.
type UpdatePasswordResp struct {
	NotFound bool `protobuf:"varint,1,opt,name=not_found,json=notFound" json:"not_found,omitempty"`
}

func (m *UpdatePasswordResp) Reset()                    { *m = UpdatePasswordResp{} }
func (m *UpdatePasswordResp) String() string            { return proto.CompactTextString(m) }
func (*UpdatePasswordResp) ProtoMessage()               {}
func (*UpdatePasswordResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// DeletePasswordReq is a request to delete a password.
type DeletePasswordReq struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
}

func (m *DeletePasswordReq) Reset()                    { *m = DeletePasswordReq{} }
func (m *DeletePasswordReq) String() string            { return proto.CompactTextString(m) }
func (*DeletePasswordReq) ProtoMessage()               {}
func (*DeletePasswordReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

// DeletePasswordResp returns the response from deleting a password.
type DeletePasswordResp struct {
	NotFound bool `protobuf:"varint,1,opt,name=not_found,json=notFound" json:"not_found,omitempty"`
}

func (m *DeletePasswordResp) Reset()                    { *m = DeletePasswordResp{} }
func (m *DeletePasswordResp) String() string            { return proto.CompactTextString(m) }
func (*DeletePasswordResp) ProtoMessage()               {}
func (*DeletePasswordResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

// ListPasswordReq is a request to enumerate passwords.
type ListPasswordReq struct {
}

func (m *ListPasswordReq) Reset()                    { *m = ListPasswordReq{} }
func (m *ListPasswordReq) String() string            { return proto.CompactTextString(m) }
func (*ListPasswordReq) ProtoMessage()               {}
func (*ListPasswordReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

// ListPasswordResp returns a list of passwords.
type ListPasswordResp struct {
	Passwords []*Password `protobuf:"bytes,1,rep,name=passwords" json:"passwords,omitempty"`
}

func (m *ListPasswordResp) Reset()                    { *m = ListPasswordResp{} }
func (m *ListPasswordResp) String() string            { return proto.CompactTextString(m) }
func (*ListPasswordResp) ProtoMessage()               {}
func (*ListPasswordResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *ListPasswordResp) GetPasswords() []*Password {
	if m != nil {
		return m.Passwords
	}
	return nil
}

// VersionReq is a request to fetch version info.
type VersionReq struct {
}

func (m *VersionReq) Reset()                    { *m = VersionReq{} }
func (m *VersionReq) String() string            { return proto.CompactTextString(m) }
func (*VersionReq) ProtoMessage()               {}
func (*VersionReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

// VersionResp holds the version info of components.
type VersionResp struct {
	// Semantic version of the server.
	Server string `protobuf:"bytes,1,opt,name=server" json:"server,omitempty"`
	// Numeric version of the API. It increases everytime a new call is added to the API.
	// Clients should use this info to determine if the server supports specific features.
	Api int32 `protobuf:"varint,2,opt,name=api" json:"api,omitempty"`
}

func (m *VersionResp) Reset()                    { *m = VersionResp{} }
func (m *VersionResp) String() string            { return proto.CompactTextString(m) }
func (*VersionResp) ProtoMessage()               {}
func (*VersionResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

// RefreshTokenRef contains the metadata for a refresh token that is managed by the storage.
type RefreshTokenRef struct {
	// ID of the refresh token.
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	ClientId  string `protobuf:"bytes,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	CreatedAt string `protobuf:"bytes,3,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	LastUsed  string `protobuf:"bytes,4,opt,name=last_used,json=lastUsed" json:"last_used,omitempty"`
}

func (m *RefreshTokenRef) Reset()                    { *m = RefreshTokenRef{} }
func (m *RefreshTokenRef) String() string            { return proto.CompactTextString(m) }
func (*RefreshTokenRef) ProtoMessage()               {}
func (*RefreshTokenRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

// ListRefreshReq is a request to enumerate the refresh tokens of a user.
type ListRefreshReq struct {
	// The "sub" claim returned in the ID Token.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *ListRefreshReq) Reset()                    { *m = ListRefreshReq{} }
func (m *ListRefreshReq) String() string            { return proto.CompactTextString(m) }
func (*ListRefreshReq) ProtoMessage()               {}
func (*ListRefreshReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

// ListRefreshResp returns a list of refresh tokens for a user.
type ListRefreshResp struct {
	RefreshTokens []*RefreshTokenRef `protobuf:"bytes,1,rep,name=refresh_tokens,json=refreshTokens" json:"refresh_tokens,omitempty"`
}

func (m *ListRefreshResp) Reset()                    { *m = ListRefreshResp{} }
func (m *ListRefreshResp) String() string            { return proto.CompactTextString(m) }
func (*ListRefreshResp) ProtoMessage()               {}
func (*ListRefreshResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *ListRefreshResp) GetRefreshTokens() []*RefreshTokenRef {
	if m != nil {
		return m.RefreshTokens
	}
	return nil
}

func init() {
	proto.RegisterType((*Client)(nil), "api.Client")
	proto.RegisterType((*CreateClientReq)(nil), "api.CreateClientReq")
	proto.RegisterType((*CreateClientResp)(nil), "api.CreateClientResp")
	proto.RegisterType((*DeleteClientReq)(nil), "api.DeleteClientReq")
	proto.RegisterType((*DeleteClientResp)(nil), "api.DeleteClientResp")
	proto.RegisterType((*Password)(nil), "api.Password")
	proto.RegisterType((*CreatePasswordReq)(nil), "api.CreatePasswordReq")
	proto.RegisterType((*CreatePasswordResp)(nil), "api.CreatePasswordResp")
	proto.RegisterType((*UpdatePasswordReq)(nil), "api.UpdatePasswordReq")
	proto.RegisterType((*UpdatePasswordResp)(nil), "api.UpdatePasswordResp")
	proto.RegisterType((*DeletePasswordReq)(nil), "api.DeletePasswordReq")
	proto.RegisterType((*DeletePasswordResp)(nil), "api.DeletePasswordResp")
	proto.RegisterType((*ListPasswordReq)(nil), "api.ListPasswordReq")
	proto.RegisterType((*ListPasswordResp)(nil), "api.ListPasswordResp")
	proto.RegisterType((*VersionReq)(nil), "api.VersionReq")
	proto.RegisterType((*VersionResp)(nil), "api.VersionResp")
	proto.RegisterType((*RefreshTokenRef)(nil), "api.RefreshTokenRef")
	proto.RegisterType((*ListRefreshReq)(nil), "api.ListRefreshReq")
	proto.RegisterType((*ListRefreshResp)(nil), "api.ListRefreshResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Dex service

type DexClient interface {
	// CreateClient creates a client.
	CreateClient(ctx context.Context, in *CreateClientReq, opts ...grpc.CallOption) (*CreateClientResp, error)
	// DeleteClient deletes the provided client.
	DeleteClient(ctx context.Context, in *DeleteClientReq, opts ...grpc.CallOption) (*DeleteClientResp, error)
	// CreatePassword creates a password.
	CreatePassword(ctx context.Context, in *CreatePasswordReq, opts ...grpc.CallOption) (*CreatePasswordResp, error)
	// UpdatePassword modifies existing password.
	UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*UpdatePasswordResp, error)
	// DeletePassword deletes the password.
	DeletePassword(ctx context.Context, in *DeletePasswordReq, opts ...grpc.CallOption) (*DeletePasswordResp, error)
	// ListPassword lists all password entries.
	ListPasswords(ctx context.Context, in *ListPasswordReq, opts ...grpc.CallOption) (*ListPasswordResp, error)
	// GetVersion returns version information of the server.
	GetVersion(ctx context.Context, in *VersionReq, opts ...grpc.CallOption) (*VersionResp, error)
	// ListRefresh lists all the refresh token entries for a particular user.
	ListRefresh(ctx context.Context, in *ListRefreshReq, opts ...grpc.CallOption) (*ListRefreshResp, error)
}

type dexClient struct {
	cc *grpc.ClientConn
}

func NewDexClient(cc *grpc.ClientConn) DexClient {
	return &dexClient{cc}
}

func (c *dexClient) CreateClient(ctx context.Context, in *CreateClientReq, opts ...grpc.CallOption) (*CreateClientResp, error) {
	out := new(CreateClientResp)
	err := grpc.Invoke(ctx, "/api.Dex/CreateClient", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) DeleteClient(ctx context.Context, in *DeleteClientReq, opts ...grpc.CallOption) (*DeleteClientResp, error) {
	out := new(DeleteClientResp)
	err := grpc.Invoke(ctx, "/api.Dex/DeleteClient", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) CreatePassword(ctx context.Context, in *CreatePasswordReq, opts ...grpc.CallOption) (*CreatePasswordResp, error) {
	out := new(CreatePasswordResp)
	err := grpc.Invoke(ctx, "/api.Dex/CreatePassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*UpdatePasswordResp, error) {
	out := new(UpdatePasswordResp)
	err := grpc.Invoke(ctx, "/api.Dex/UpdatePassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) DeletePassword(ctx context.Context, in *DeletePasswordReq, opts ...grpc.CallOption) (*DeletePasswordResp, error) {
	out := new(DeletePasswordResp)
	err := grpc.Invoke(ctx, "/api.Dex/DeletePassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) ListPasswords(ctx context.Context, in *ListPasswordReq, opts ...grpc.CallOption) (*ListPasswordResp, error) {
	out := new(ListPasswordResp)
	err := grpc.Invoke(ctx, "/api.Dex/ListPasswords", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) GetVersion(ctx context.Context, in *VersionReq, opts ...grpc.CallOption) (*VersionResp, error) {
	out := new(VersionResp)
	err := grpc.Invoke(ctx, "/api.Dex/GetVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dexClient) ListRefresh(ctx context.Context, in *ListRefreshReq, opts ...grpc.CallOption) (*ListRefreshResp, error) {
	out := new(ListRefreshResp)
	err := grpc.Invoke(ctx, "/api.Dex/ListRefresh", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Dex service

type DexServer interface {
	// CreateClient creates a client.
	CreateClient(context.Context, *CreateClientReq) (*CreateClientResp, error)
	// DeleteClient deletes the provided client.
	DeleteClient(context.Context, *DeleteClientReq) (*DeleteClientResp, error)
	// CreatePassword creates a password.
	CreatePassword(context.Context, *CreatePasswordReq) (*CreatePasswordResp, error)
	// UpdatePassword modifies existing password.
	UpdatePassword(context.Context, *UpdatePasswordReq) (*UpdatePasswordResp, error)
	// DeletePassword deletes the password.
	DeletePassword(context.Context, *DeletePasswordReq) (*DeletePasswordResp, error)
	// ListPassword lists all password entries.
	ListPasswords(context.Context, *ListPasswordReq) (*ListPasswordResp, error)
	// GetVersion returns version information of the server.
	GetVersion(context.Context, *VersionReq) (*VersionResp, error)
	// ListRefresh lists all the refresh token entries for a particular user.
	ListRefresh(context.Context, *ListRefreshReq) (*ListRefreshResp, error)
}

func RegisterDexServer(s *grpc.Server, srv DexServer) {
	s.RegisterService(&_Dex_serviceDesc, srv)
}

func _Dex_CreateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).CreateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Dex/CreateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).CreateClient(ctx, req.(*CreateClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_DeleteClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).DeleteClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Dex/DeleteClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).DeleteClient(ctx, req.(*DeleteClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_CreatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).CreatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Dex/CreatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).CreatePassword(ctx, req.(*CreatePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Dex/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).UpdatePassword(ctx, req.(*UpdatePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_DeletePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).DeletePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Dex/DeletePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).DeletePassword(ctx, req.(*DeletePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_ListPasswords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).ListPasswords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Dex/ListPasswords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).ListPasswords(ctx, req.(*ListPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Dex/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).GetVersion(ctx, req.(*VersionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dex_ListRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRefreshReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DexServer).ListRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Dex/ListRefresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DexServer).ListRefresh(ctx, req.(*ListRefreshReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dex_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Dex",
	HandlerType: (*DexServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClient",
			Handler:    _Dex_CreateClient_Handler,
		},
		{
			MethodName: "DeleteClient",
			Handler:    _Dex_DeleteClient_Handler,
		},
		{
			MethodName: "CreatePassword",
			Handler:    _Dex_CreatePassword_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _Dex_UpdatePassword_Handler,
		},
		{
			MethodName: "DeletePassword",
			Handler:    _Dex_DeletePassword_Handler,
		},
		{
			MethodName: "ListPasswords",
			Handler:    _Dex_ListPasswords_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _Dex_GetVersion_Handler,
		},
		{
			MethodName: "ListRefresh",
			Handler:    _Dex_ListRefresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("api/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 742 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x55, 0x6b, 0x4b, 0x1c, 0x4b,
	0x10, 0xdd, 0x97, 0xbb, 0xb3, 0xb5, 0xef, 0xbe, 0x5e, 0x1d, 0x57, 0x2e, 0x68, 0xcb, 0x05, 0xe5,
	0x82, 0xa2, 0x17, 0x12, 0x88, 0xc4, 0x10, 0x34, 0x0f, 0x21, 0x04, 0x19, 0xb2, 0xf9, 0x98, 0x61,
	0xdc, 0x29, 0xb5, 0xc9, 0x38, 0x33, 0xe9, 0xee, 0xcd, 0x9a, 0x7c, 0xcb, 0x4f, 0xcb, 0x3f, 0x0b,
	0xfd, 0xd8, 0x75, 0x1e, 0x06, 0xf3, 0x6d, 0xea, 0x74, 0xd7, 0xa9, 0xae, 0xd3, 0xa7, 0x6b, 0xa0,
	0x17, 0xa4, 0xec, 0x20, 0x48, 0xd9, 0x7e, 0xca, 0x13, 0x99, 0x90, 0x7a, 0x90, 0x32, 0xfa, 0xb3,
	0x0a, 0xcd, 0xd3, 0x88, 0x61, 0x2c, 0x49, 0x1f, 0x6a, 0x2c, 0x74, 0xab, 0x5b, 0xd5, 0xdd, 0xb6,
	0x57, 0x63, 0x21, 0x59, 0x83, 0xa6, 0xc0, 0x29, 0x47, 0xe9, 0xd6, 0x34, 0x66, 0x23, 0xb2, 0x03,
	0x3d, 0x8e, 0x21, 0xe3, 0x38, 0x95, 0xfe, 0x8c, 0x33, 0xe1, 0xd6, 0xb7, 0xea, 0xbb, 0x6d, 0xaf,
	0xbb, 0x00, 0x27, 0x9c, 0x09, 0xb5, 0x49, 0xf2, 0x99, 0x90, 0x18, 0xfa, 0x29, 0x22, 0x17, 0x6e,
	0xc3, 0x6c, 0xb2, 0xe0, 0x85, 0xc2, 0x54, 0x85, 0x74, 0x76, 0x19, 0xb1, 0xa9, 0xbb, 0xb2, 0x55,
	0xdd, 0x75, 0x3c, 0x1b, 0x11, 0x02, 0x8d, 0x38, 0xb8, 0x45, 0xb7, 0xa9, 0xeb, 0xea, 0x6f, 0xb2,
	0x01, 0x4e, 0x94, 0x5c, 0x27, 0xfe, 0x8c, 0x47, 0x6e, 0x4b, 0xe3, 0x2d, 0x15, 0x4f, 0x78, 0x44,
	0x9f, 0xc0, 0xe0, 0x94, 0x63, 0x20, 0xd1, 0x34, 0xe2, 0xe1, 0x17, 0xb2, 0x03, 0xcd, 0xa9, 0x0e,
	0x74, 0x3f, 0x9d, 0xa3, 0xce, 0xbe, 0xea, 0xdb, 0xae, 0xdb, 0x25, 0xfa, 0x09, 0x86, 0xf9, 0x3c,
	0x91, 0x92, 0x7f, 0xa1, 0x1f, 0x44, 0x1c, 0x83, 0xf0, 0x9b, 0x8f, 0x77, 0x4c, 0x48, 0xa1, 0x09,
	0x1c, 0xaf, 0x67, 0xd1, 0x57, 0x1a, 0xcc, 0xf0, 0xd7, 0x7e, 0xcf, 0xbf, 0x0d, 0x83, 0x33, 0x8c,
	0x30, 0x7b, 0xae, 0x82, 0xc6, 0xf4, 0x00, 0x86, 0xf9, 0x2d, 0x22, 0x25, 0x9b, 0xd0, 0x8e, 0x13,
	0xe9, 0x5f, 0x25, 0xb3, 0x38, 0xb4, 0xd5, 0x9d, 0x38, 0x91, 0xaf, 0x55, 0x4c, 0x19, 0x38, 0x17,
	0x81, 0x10, 0xf3, 0x84, 0x87, 0x64, 0x15, 0x56, 0xf0, 0x36, 0x60, 0x91, 0xe5, 0x33, 0x81, 0x12,
	0xef, 0x26, 0x10, 0x37, 0xfa, 0x60, 0x5d, 0x4f, 0x7f, 0x93, 0x31, 0x38, 0x33, 0x81, 0x5c, 0x8b,
	0x5a, 0xd7, 0x9b, 0x97, 0x31, 0x59, 0x87, 0x96, 0xfa, 0xf6, 0x59, 0xe8, 0x36, 0xcc, 0x3d, 0xab,
	0xf0, 0x3c, 0xa4, 0x27, 0x30, 0x32, 0xf2, 0x2c, 0x0a, 0xaa, 0x06, 0xf6, 0xc0, 0x49, 0x6d, 0x68,
	0xa5, 0xed, 0xe9, 0xd6, 0x97, 0x7b, 0x96, 0xcb, 0xf4, 0x18, 0x48, 0x31, 0xff, 0x8f, 0x05, 0xa6,
	0xd7, 0x30, 0x9a, 0xa4, 0x61, 0xa1, 0xf8, 0xc3, 0x0d, 0x6f, 0x80, 0x13, 0xe3, 0xdc, 0xcf, 0x34,
	0xdd, 0x8a, 0x71, 0xfe, 0x56, 0xf5, 0xbd, 0x0d, 0x5d, 0xb5, 0x54, 0xe8, 0xbd, 0x13, 0xe3, 0x7c,
	0x62, 0x21, 0x7a, 0x08, 0xa4, 0x58, 0xe8, 0xb1, 0x3b, 0xd8, 0x83, 0x91, 0xb9, 0xb4, 0x47, 0xcf,
	0xa6, 0xd8, 0x8b, 0x5b, 0x1f, 0x63, 0x1f, 0xc1, 0xe0, 0x1d, 0x13, 0x32, 0xc3, 0x4d, 0x5f, 0xc0,
	0x30, 0x0f, 0x89, 0x94, 0xfc, 0x07, 0xed, 0x85, 0xd2, 0x4a, 0xc2, 0x7a, 0xf9, 0x26, 0xee, 0xd7,
	0x69, 0x17, 0xe0, 0x23, 0x72, 0xc1, 0x92, 0x58, 0xd1, 0x3d, 0x85, 0xce, 0x32, 0x12, 0xa9, 0x79,
	0xe7, 0xfc, 0x2b, 0x72, 0x7b, 0x74, 0x1b, 0x91, 0x21, 0xa8, 0x09, 0xa1, 0x25, 0x5d, 0xf1, 0xf4,
	0xb0, 0xf8, 0x0e, 0x03, 0x0f, 0xaf, 0x38, 0x8a, 0x9b, 0x0f, 0xc9, 0x67, 0x8c, 0x3d, 0xbc, 0x2a,
	0x0d, 0x8d, 0x4d, 0x68, 0x1b, 0xf7, 0x2b, 0x3f, 0x99, 0xb9, 0xe1, 0x18, 0xe0, 0x3c, 0x24, 0xff,
	0x00, 0x4c, 0xb5, 0x23, 0x42, 0x3f, 0x90, 0xf6, 0x32, 0xda, 0x16, 0x79, 0x29, 0x55, 0x6e, 0x14,
	0x08, 0xa9, 0xae, 0x6b, 0xe1, 0x45, 0x47, 0x01, 0x13, 0x81, 0x4a, 0xf4, 0xbe, 0xd2, 0xc0, 0xd6,
	0x57, 0x8a, 0x67, 0x8c, 0x5b, 0xcd, 0x19, 0xf7, 0xbd, 0x51, 0x70, 0xb9, 0x55, 0xa4, 0xe4, 0x18,
	0xfa, 0xdc, 0x84, 0xbe, 0x54, 0x47, 0x5f, 0x48, 0xb6, 0xaa, 0x25, 0x2b, 0x34, 0xe5, 0xf5, 0x78,
	0x06, 0x10, 0x47, 0x3f, 0x1a, 0x50, 0x3f, 0xc3, 0x3b, 0xf2, 0x1c, 0xba, 0xd9, 0x79, 0x41, 0x4c,
	0x72, 0x61, 0xf4, 0x8c, 0xff, 0x7e, 0x00, 0x15, 0x29, 0xad, 0xa8, 0xf4, 0xec, 0x5b, 0xb7, 0xe9,
	0x85, 0x09, 0x61, 0xd3, 0x8b, 0x43, 0x81, 0x56, 0xc8, 0x29, 0xf4, 0xf3, 0xcf, 0x89, 0xac, 0x65,
	0x2a, 0x65, 0xec, 0x32, 0x5e, 0x7f, 0x10, 0x5f, 0x90, 0xe4, 0xdd, 0x6e, 0x49, 0x4a, 0x6f, 0xcd,
	0x92, 0x94, 0x9f, 0x86, 0x21, 0xc9, 0x9b, 0xda, 0x92, 0x94, 0x1e, 0x85, 0x25, 0x29, 0xbf, 0x00,
	0x5a, 0x21, 0x27, 0xd0, 0xcb, 0x7a, 0x5a, 0x58, 0x39, 0x0a, 0xd6, 0xb7, 0x72, 0x14, 0xdd, 0x4f,
	0x2b, 0xe4, 0x10, 0xe0, 0x0d, 0x4a, 0xeb, 0x63, 0x32, 0xd0, 0xdb, 0xee, 0x3d, 0x3e, 0x1e, 0xe6,
	0x01, 0x9d, 0xf2, 0x0c, 0x3a, 0x19, 0x5f, 0x90, 0xbf, 0x96, 0xd4, 0xf7, 0xa6, 0x1a, 0xaf, 0x96,
	0x41, 0x95, 0x7b, 0xd9, 0xd4, 0xff, 0xcc, 0xff, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x57, 0x96,
	0xbe, 0xf0, 0x44, 0x07, 0x00, 0x00,
}