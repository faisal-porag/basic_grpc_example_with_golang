// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package tprotos

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

// DetailsServiceClient is the client API for DetailsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DetailsServiceClient interface {
	GetTProto(ctx context.Context, in *TProtoRequest, opts ...grpc.CallOption) (*TProtoResponse, error)
}

type detailsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDetailsServiceClient(cc grpc.ClientConnInterface) DetailsServiceClient {
	return &detailsServiceClient{cc}
}

func (c *detailsServiceClient) GetTProto(ctx context.Context, in *TProtoRequest, opts ...grpc.CallOption) (*TProtoResponse, error) {
	out := new(TProtoResponse)
	err := c.cc.Invoke(ctx, "/tprotos.DetailsService/GetTProto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DetailsServiceServer is the server API for DetailsService service.
// All implementations must embed UnimplementedDetailsServiceServer
// for forward compatibility
type DetailsServiceServer interface {
	GetTProto(context.Context, *TProtoRequest) (*TProtoResponse, error)
	mustEmbedUnimplementedDetailsServiceServer()
}

// UnimplementedDetailsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDetailsServiceServer struct {
}

func (UnimplementedDetailsServiceServer) GetTProto(context.Context, *TProtoRequest) (*TProtoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTProto not implemented")
}
func (UnimplementedDetailsServiceServer) mustEmbedUnimplementedDetailsServiceServer() {}

// UnsafeDetailsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DetailsServiceServer will
// result in compilation errors.
type UnsafeDetailsServiceServer interface {
	mustEmbedUnimplementedDetailsServiceServer()
}

func RegisterDetailsServiceServer(s grpc.ServiceRegistrar, srv DetailsServiceServer) {
	s.RegisterService(&DetailsService_ServiceDesc, srv)
}

func _DetailsService_GetTProto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TProtoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DetailsServiceServer).GetTProto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tprotos.DetailsService/GetTProto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DetailsServiceServer).GetTProto(ctx, req.(*TProtoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DetailsService_ServiceDesc is the grpc.ServiceDesc for DetailsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DetailsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tprotos.DetailsService",
	HandlerType: (*DetailsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTProto",
			Handler:    _DetailsService_GetTProto_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tprotos/tprotos.proto",
}
