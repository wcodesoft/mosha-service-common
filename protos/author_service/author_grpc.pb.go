// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: proto/author.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AuthorService_GetAuthor_FullMethodName    = "/authorservice.AuthorService/GetAuthor"
	AuthorService_ListAuthors_FullMethodName  = "/authorservice.AuthorService/ListAuthors"
	AuthorService_CreateAuthor_FullMethodName = "/authorservice.AuthorService/CreateAuthor"
	AuthorService_UpdateAuthor_FullMethodName = "/authorservice.AuthorService/UpdateAuthor"
	AuthorService_DeleteAuthor_FullMethodName = "/authorservice.AuthorService/DeleteAuthor"
)

// AuthorServiceClient is the client API for AuthorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorServiceClient interface {
	// GetAuthor returns an author by id
	GetAuthor(ctx context.Context, in *GetAuthorRequest, opts ...grpc.CallOption) (*Author, error)
	// ListAuthors returns a list of authors
	ListAuthors(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListAuthorsResponse, error)
	// CreateAuthor creates a new author
	CreateAuthor(ctx context.Context, in *CreateAuthorRequest, opts ...grpc.CallOption) (*Author, error)
	// UpdateAuthor updates an existing author
	UpdateAuthor(ctx context.Context, in *UpdateAuthorRequest, opts ...grpc.CallOption) (*Author, error)
	// DeleteAuthor deletes an author by id
	DeleteAuthor(ctx context.Context, in *DeleteAuthorRequest, opts ...grpc.CallOption) (*DeleteAuthorResponse, error)
}

type authorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorServiceClient(cc grpc.ClientConnInterface) AuthorServiceClient {
	return &authorServiceClient{cc}
}

func (c *authorServiceClient) GetAuthor(ctx context.Context, in *GetAuthorRequest, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, AuthorService_GetAuthor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) ListAuthors(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListAuthorsResponse, error) {
	out := new(ListAuthorsResponse)
	err := c.cc.Invoke(ctx, AuthorService_ListAuthors_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) CreateAuthor(ctx context.Context, in *CreateAuthorRequest, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, AuthorService_CreateAuthor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) UpdateAuthor(ctx context.Context, in *UpdateAuthorRequest, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, AuthorService_UpdateAuthor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) DeleteAuthor(ctx context.Context, in *DeleteAuthorRequest, opts ...grpc.CallOption) (*DeleteAuthorResponse, error) {
	out := new(DeleteAuthorResponse)
	err := c.cc.Invoke(ctx, AuthorService_DeleteAuthor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorServiceServer is the server API for AuthorService service.
// All implementations must embed UnimplementedAuthorServiceServer
// for forward compatibility
type AuthorServiceServer interface {
	// GetAuthor returns an author by id
	GetAuthor(context.Context, *GetAuthorRequest) (*Author, error)
	// ListAuthors returns a list of authors
	ListAuthors(context.Context, *emptypb.Empty) (*ListAuthorsResponse, error)
	// CreateAuthor creates a new author
	CreateAuthor(context.Context, *CreateAuthorRequest) (*Author, error)
	// UpdateAuthor updates an existing author
	UpdateAuthor(context.Context, *UpdateAuthorRequest) (*Author, error)
	// DeleteAuthor deletes an author by id
	DeleteAuthor(context.Context, *DeleteAuthorRequest) (*DeleteAuthorResponse, error)
	mustEmbedUnimplementedAuthorServiceServer()
}

// UnimplementedAuthorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorServiceServer struct {
}

func (UnimplementedAuthorServiceServer) GetAuthor(context.Context, *GetAuthorRequest) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthor not implemented")
}
func (UnimplementedAuthorServiceServer) ListAuthors(context.Context, *emptypb.Empty) (*ListAuthorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAuthors not implemented")
}
func (UnimplementedAuthorServiceServer) CreateAuthor(context.Context, *CreateAuthorRequest) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuthor not implemented")
}
func (UnimplementedAuthorServiceServer) UpdateAuthor(context.Context, *UpdateAuthorRequest) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuthor not implemented")
}
func (UnimplementedAuthorServiceServer) DeleteAuthor(context.Context, *DeleteAuthorRequest) (*DeleteAuthorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAuthor not implemented")
}
func (UnimplementedAuthorServiceServer) mustEmbedUnimplementedAuthorServiceServer() {}

// UnsafeAuthorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorServiceServer will
// result in compilation errors.
type UnsafeAuthorServiceServer interface {
	mustEmbedUnimplementedAuthorServiceServer()
}

func RegisterAuthorServiceServer(s grpc.ServiceRegistrar, srv AuthorServiceServer) {
	s.RegisterService(&AuthorService_ServiceDesc, srv)
}

func _AuthorService_GetAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).GetAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorService_GetAuthor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).GetAuthor(ctx, req.(*GetAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorService_ListAuthors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).ListAuthors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorService_ListAuthors_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).ListAuthors(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorService_CreateAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).CreateAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorService_CreateAuthor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).CreateAuthor(ctx, req.(*CreateAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorService_UpdateAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).UpdateAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorService_UpdateAuthor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).UpdateAuthor(ctx, req.(*UpdateAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorService_DeleteAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).DeleteAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorService_DeleteAuthor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).DeleteAuthor(ctx, req.(*DeleteAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthorService_ServiceDesc is the grpc.ServiceDesc for AuthorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authorservice.AuthorService",
	HandlerType: (*AuthorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuthor",
			Handler:    _AuthorService_GetAuthor_Handler,
		},
		{
			MethodName: "ListAuthors",
			Handler:    _AuthorService_ListAuthors_Handler,
		},
		{
			MethodName: "CreateAuthor",
			Handler:    _AuthorService_CreateAuthor_Handler,
		},
		{
			MethodName: "UpdateAuthor",
			Handler:    _AuthorService_UpdateAuthor_Handler,
		},
		{
			MethodName: "DeleteAuthor",
			Handler:    _AuthorService_DeleteAuthor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/author.proto",
}
