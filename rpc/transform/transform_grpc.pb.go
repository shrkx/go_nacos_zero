// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: transform.proto

package transform

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

// TransformClient is the client API for Transform service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransformClient interface {
	GetShortUrl(ctx context.Context, in *GetShortUrlRequest, opts ...grpc.CallOption) (*GetShortUrlResponse, error)
	GetLongUrl(ctx context.Context, in *GetLongUrlRequest, opts ...grpc.CallOption) (*GetLongUrlResponse, error)
}

type transformClient struct {
	cc grpc.ClientConnInterface
}

func NewTransformClient(cc grpc.ClientConnInterface) TransformClient {
	return &transformClient{cc}
}

func (c *transformClient) GetShortUrl(ctx context.Context, in *GetShortUrlRequest, opts ...grpc.CallOption) (*GetShortUrlResponse, error) {
	out := new(GetShortUrlResponse)
	err := c.cc.Invoke(ctx, "/transform.Transform/GetShortUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transformClient) GetLongUrl(ctx context.Context, in *GetLongUrlRequest, opts ...grpc.CallOption) (*GetLongUrlResponse, error) {
	out := new(GetLongUrlResponse)
	err := c.cc.Invoke(ctx, "/transform.Transform/GetLongUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransformServer is the server API for Transform service.
// All implementations must embed UnimplementedTransformServer
// for forward compatibility
type TransformServer interface {
	GetShortUrl(context.Context, *GetShortUrlRequest) (*GetShortUrlResponse, error)
	GetLongUrl(context.Context, *GetLongUrlRequest) (*GetLongUrlResponse, error)
	mustEmbedUnimplementedTransformServer()
}

// UnimplementedTransformServer must be embedded to have forward compatible implementations.
type UnimplementedTransformServer struct {
}

func (UnimplementedTransformServer) GetShortUrl(context.Context, *GetShortUrlRequest) (*GetShortUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShortUrl not implemented")
}
func (UnimplementedTransformServer) GetLongUrl(context.Context, *GetLongUrlRequest) (*GetLongUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLongUrl not implemented")
}
func (UnimplementedTransformServer) mustEmbedUnimplementedTransformServer() {}

// UnsafeTransformServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransformServer will
// result in compilation errors.
type UnsafeTransformServer interface {
	mustEmbedUnimplementedTransformServer()
}

func RegisterTransformServer(s grpc.ServiceRegistrar, srv TransformServer) {
	s.RegisterService(&Transform_ServiceDesc, srv)
}

func _Transform_GetShortUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShortUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransformServer).GetShortUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transform.Transform/GetShortUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransformServer).GetShortUrl(ctx, req.(*GetShortUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transform_GetLongUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLongUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransformServer).GetLongUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transform.Transform/GetLongUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransformServer).GetLongUrl(ctx, req.(*GetLongUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Transform_ServiceDesc is the grpc.ServiceDesc for Transform service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Transform_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transform.Transform",
	HandlerType: (*TransformServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShortUrl",
			Handler:    _Transform_GetShortUrl_Handler,
		},
		{
			MethodName: "GetLongUrl",
			Handler:    _Transform_GetLongUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transform.proto",
}
