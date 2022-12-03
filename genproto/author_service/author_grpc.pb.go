// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: author.proto

package author_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthorServiceClient is the client API for AuthorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorServiceClient interface {
	CreateAuthor(ctx context.Context, in *CreateAuthorRequest, opts ...grpc.CallOption) (*CreateAuthorResponse, error)
	GetAuthor(ctx context.Context, in *GetAuthorRequest, opts ...grpc.CallOption) (*GetAuthorResponse, error)
	GetAuthorById(ctx context.Context, in *GetAuthorByIdRequest, opts ...grpc.CallOption) (*Author, error)
	UpdateAuthor(ctx context.Context, in *UpdateAuthorRequest, opts ...grpc.CallOption) (*UpdateAuthorResponse, error)
	DeleteAuthor(ctx context.Context, in *DeleteAuthorRequest, opts ...grpc.CallOption) (*DeleteAuthorResponse, error)
}

type authorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorServiceClient(cc grpc.ClientConnInterface) AuthorServiceClient {
	return &authorServiceClient{cc}
}

func (c *authorServiceClient) CreateAuthor(ctx context.Context, in *CreateAuthorRequest, opts ...grpc.CallOption) (*CreateAuthorResponse, error) {
	out := new(CreateAuthorResponse)
	err := c.cc.Invoke(ctx, "/genproto.AuthorService/CreateAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) GetAuthor(ctx context.Context, in *GetAuthorRequest, opts ...grpc.CallOption) (*GetAuthorResponse, error) {
	out := new(GetAuthorResponse)
	err := c.cc.Invoke(ctx, "/genproto.AuthorService/GetAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) GetAuthorById(ctx context.Context, in *GetAuthorByIdRequest, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, "/genproto.AuthorService/GetAuthorById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) UpdateAuthor(ctx context.Context, in *UpdateAuthorRequest, opts ...grpc.CallOption) (*UpdateAuthorResponse, error) {
	out := new(UpdateAuthorResponse)
	err := c.cc.Invoke(ctx, "/genproto.AuthorService/UpdateAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) DeleteAuthor(ctx context.Context, in *DeleteAuthorRequest, opts ...grpc.CallOption) (*DeleteAuthorResponse, error) {
	out := new(DeleteAuthorResponse)
	err := c.cc.Invoke(ctx, "/genproto.AuthorService/DeleteAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorServiceServer is the server API for AuthorService service.
// All implementations must embed UnimplementedAuthorServiceServer
// for forward compatibility
type AuthorServiceServer interface {
	CreateAuthor(context.Context, *CreateAuthorRequest) (*CreateAuthorResponse, error)
	GetAuthor(context.Context, *GetAuthorRequest) (*GetAuthorResponse, error)
	GetAuthorById(context.Context, *GetAuthorByIdRequest) (*Author, error)
	UpdateAuthor(context.Context, *UpdateAuthorRequest) (*UpdateAuthorResponse, error)
	DeleteAuthor(context.Context, *DeleteAuthorRequest) (*DeleteAuthorResponse, error)
	mustEmbedUnimplementedAuthorServiceServer()
}

// UnimplementedAuthorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorServiceServer struct {
}

func (UnimplementedAuthorServiceServer) CreateAuthor(context.Context, *CreateAuthorRequest) (*CreateAuthorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuthor not implemented")
}
func (UnimplementedAuthorServiceServer) GetAuthor(context.Context, *GetAuthorRequest) (*GetAuthorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthor not implemented")
}
func (UnimplementedAuthorServiceServer) GetAuthorById(context.Context, *GetAuthorByIdRequest) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthorById not implemented")
}
func (UnimplementedAuthorServiceServer) UpdateAuthor(context.Context, *UpdateAuthorRequest) (*UpdateAuthorResponse, error) {
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
		FullMethod: "/genproto.AuthorService/CreateAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).CreateAuthor(ctx, req.(*CreateAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
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
		FullMethod: "/genproto.AuthorService/GetAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).GetAuthor(ctx, req.(*GetAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorService_GetAuthorById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthorByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).GetAuthorById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.AuthorService/GetAuthorById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).GetAuthorById(ctx, req.(*GetAuthorByIdRequest))
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
		FullMethod: "/genproto.AuthorService/UpdateAuthor",
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
		FullMethod: "/genproto.AuthorService/DeleteAuthor",
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
	ServiceName: "genproto.AuthorService",
	HandlerType: (*AuthorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAuthor",
			Handler:    _AuthorService_CreateAuthor_Handler,
		},
		{
			MethodName: "GetAuthor",
			Handler:    _AuthorService_GetAuthor_Handler,
		},
		{
			MethodName: "GetAuthorById",
			Handler:    _AuthorService_GetAuthorById_Handler,
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
	Metadata: "author.proto",
}