// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ova_link_api

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

// LinkAPIClient is the client API for LinkAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LinkAPIClient interface {
	CreateLink(ctx context.Context, in *CreateLinkRequest, opts ...grpc.CallOption) (*CreateLinkResponse, error)
	DescribeLink(ctx context.Context, in *DescribeLinkRequest, opts ...grpc.CallOption) (*DescribeLinkResponse, error)
	ListLink(ctx context.Context, in *ListLinkRequest, opts ...grpc.CallOption) (*ListLinkResponse, error)
	DeleteLink(ctx context.Context, in *DeleteLinkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type linkAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewLinkAPIClient(cc grpc.ClientConnInterface) LinkAPIClient {
	return &linkAPIClient{cc}
}

func (c *linkAPIClient) CreateLink(ctx context.Context, in *CreateLinkRequest, opts ...grpc.CallOption) (*CreateLinkResponse, error) {
	out := new(CreateLinkResponse)
	err := c.cc.Invoke(ctx, "/ova.link.api.LinkAPI/CreateLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkAPIClient) DescribeLink(ctx context.Context, in *DescribeLinkRequest, opts ...grpc.CallOption) (*DescribeLinkResponse, error) {
	out := new(DescribeLinkResponse)
	err := c.cc.Invoke(ctx, "/ova.link.api.LinkAPI/DescribeLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkAPIClient) ListLink(ctx context.Context, in *ListLinkRequest, opts ...grpc.CallOption) (*ListLinkResponse, error) {
	out := new(ListLinkResponse)
	err := c.cc.Invoke(ctx, "/ova.link.api.LinkAPI/ListLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkAPIClient) DeleteLink(ctx context.Context, in *DeleteLinkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ova.link.api.LinkAPI/DeleteLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkAPIServer is the server API for LinkAPI service.
// All implementations must embed UnimplementedLinkAPIServer
// for forward compatibility
type LinkAPIServer interface {
	CreateLink(context.Context, *CreateLinkRequest) (*CreateLinkResponse, error)
	DescribeLink(context.Context, *DescribeLinkRequest) (*DescribeLinkResponse, error)
	ListLink(context.Context, *ListLinkRequest) (*ListLinkResponse, error)
	DeleteLink(context.Context, *DeleteLinkRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedLinkAPIServer()
}

// UnimplementedLinkAPIServer must be embedded to have forward compatible implementations.
type UnimplementedLinkAPIServer struct {
}

func (UnimplementedLinkAPIServer) CreateLink(context.Context, *CreateLinkRequest) (*CreateLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLink not implemented")
}
func (UnimplementedLinkAPIServer) DescribeLink(context.Context, *DescribeLinkRequest) (*DescribeLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeLink not implemented")
}
func (UnimplementedLinkAPIServer) ListLink(context.Context, *ListLinkRequest) (*ListLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLink not implemented")
}
func (UnimplementedLinkAPIServer) DeleteLink(context.Context, *DeleteLinkRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLink not implemented")
}
func (UnimplementedLinkAPIServer) mustEmbedUnimplementedLinkAPIServer() {}

// UnsafeLinkAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LinkAPIServer will
// result in compilation errors.
type UnsafeLinkAPIServer interface {
	mustEmbedUnimplementedLinkAPIServer()
}

func RegisterLinkAPIServer(s grpc.ServiceRegistrar, srv LinkAPIServer) {
	s.RegisterService(&LinkAPI_ServiceDesc, srv)
}

func _LinkAPI_CreateLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkAPIServer).CreateLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.link.api.LinkAPI/CreateLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkAPIServer).CreateLink(ctx, req.(*CreateLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkAPI_DescribeLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkAPIServer).DescribeLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.link.api.LinkAPI/DescribeLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkAPIServer).DescribeLink(ctx, req.(*DescribeLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkAPI_ListLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkAPIServer).ListLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.link.api.LinkAPI/ListLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkAPIServer).ListLink(ctx, req.(*ListLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkAPI_DeleteLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkAPIServer).DeleteLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.link.api.LinkAPI/DeleteLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkAPIServer).DeleteLink(ctx, req.(*DeleteLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LinkAPI_ServiceDesc is the grpc.ServiceDesc for LinkAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LinkAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ova.link.api.LinkAPI",
	HandlerType: (*LinkAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLink",
			Handler:    _LinkAPI_CreateLink_Handler,
		},
		{
			MethodName: "DescribeLink",
			Handler:    _LinkAPI_DescribeLink_Handler,
		},
		{
			MethodName: "ListLink",
			Handler:    _LinkAPI_ListLink_Handler,
		},
		{
			MethodName: "DeleteLink",
			Handler:    _LinkAPI_DeleteLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "link.proto",
}
