// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package accountpb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountServiceClient interface {
	// Returnes a list of accounts
	// uses pagination
	ListAccounts(ctx context.Context, in *ListAccountsRequest, opts ...grpc.CallOption) (*ListAccountsResponse, error)
	// Counts resources based on specified filter and returnes the result
	GetAccountsCount(ctx context.Context, in *GetAccountsCountRequest, opts ...grpc.CallOption) (*GetAccountsCountResponse, error)
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*Account, error)
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*Account, error)
	UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*Account, error)
	DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type accountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServiceClient(cc grpc.ClientConnInterface) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) ListAccounts(ctx context.Context, in *ListAccountsRequest, opts ...grpc.CallOption) (*ListAccountsResponse, error) {
	out := new(ListAccountsResponse)
	err := c.cc.Invoke(ctx, "/xixo.account.v1.AccountService/ListAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccountsCount(ctx context.Context, in *GetAccountsCountRequest, opts ...grpc.CallOption) (*GetAccountsCountResponse, error) {
	out := new(GetAccountsCountResponse)
	err := c.cc.Invoke(ctx, "/xixo.account.v1.AccountService/GetAccountsCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/xixo.account.v1.AccountService/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/xixo.account.v1.AccountService/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/xixo.account.v1.AccountService/UpdateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/xixo.account.v1.AccountService/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
// All implementations must embed UnimplementedAccountServiceServer
// for forward compatibility
type AccountServiceServer interface {
	// Returnes a list of accounts
	// uses pagination
	ListAccounts(context.Context, *ListAccountsRequest) (*ListAccountsResponse, error)
	// Counts resources based on specified filter and returnes the result
	GetAccountsCount(context.Context, *GetAccountsCountRequest) (*GetAccountsCountResponse, error)
	GetAccount(context.Context, *GetAccountRequest) (*Account, error)
	CreateAccount(context.Context, *CreateAccountRequest) (*Account, error)
	UpdateAccount(context.Context, *UpdateAccountRequest) (*Account, error)
	DeleteAccount(context.Context, *DeleteAccountRequest) (*empty.Empty, error)
	mustEmbedUnimplementedAccountServiceServer()
}

// UnimplementedAccountServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (UnimplementedAccountServiceServer) ListAccounts(context.Context, *ListAccountsRequest) (*ListAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccounts not implemented")
}
func (UnimplementedAccountServiceServer) GetAccountsCount(context.Context, *GetAccountsCountRequest) (*GetAccountsCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountsCount not implemented")
}
func (UnimplementedAccountServiceServer) GetAccount(context.Context, *GetAccountRequest) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedAccountServiceServer) CreateAccount(context.Context, *CreateAccountRequest) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedAccountServiceServer) UpdateAccount(context.Context, *UpdateAccountRequest) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedAccountServiceServer) DeleteAccount(context.Context, *DeleteAccountRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedAccountServiceServer) mustEmbedUnimplementedAccountServiceServer() {}

// UnsafeAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServiceServer will
// result in compilation errors.
type UnsafeAccountServiceServer interface {
	mustEmbedUnimplementedAccountServiceServer()
}

func RegisterAccountServiceServer(s grpc.ServiceRegistrar, srv AccountServiceServer) {
	s.RegisterService(&_AccountService_serviceDesc, srv)
}

func _AccountService_ListAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).ListAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xixo.account.v1.AccountService/ListAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).ListAccounts(ctx, req.(*ListAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetAccountsCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountsCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccountsCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xixo.account.v1.AccountService/GetAccountsCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccountsCount(ctx, req.(*GetAccountsCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xixo.account.v1.AccountService/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xixo.account.v1.AccountService/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xixo.account.v1.AccountService/UpdateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).UpdateAccount(ctx, req.(*UpdateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xixo.account.v1.AccountService/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).DeleteAccount(ctx, req.(*DeleteAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "xixo.account.v1.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAccounts",
			Handler:    _AccountService_ListAccounts_Handler,
		},
		{
			MethodName: "GetAccountsCount",
			Handler:    _AccountService_GetAccountsCount_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _AccountService_GetAccount_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _AccountService_CreateAccount_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _AccountService_UpdateAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _AccountService_DeleteAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accountpb/account.proto",
}
