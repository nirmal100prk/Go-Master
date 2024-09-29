//message

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: permission/permission.proto

package permission

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	SamplePermissionService_GetPermissionDataForClient_FullMethodName = "/sampleproto.SamplePermissionService/GetPermissionDataForClient"
)

// SamplePermissionServiceClient is the client API for SamplePermissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SamplePermissionServiceClient interface {
	// client3
	GetPermissionDataForClient(ctx context.Context, in *PermissionDataRequest, opts ...grpc.CallOption) (*PermissionDataResponse, error)
}

type samplePermissionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSamplePermissionServiceClient(cc grpc.ClientConnInterface) SamplePermissionServiceClient {
	return &samplePermissionServiceClient{cc}
}

func (c *samplePermissionServiceClient) GetPermissionDataForClient(ctx context.Context, in *PermissionDataRequest, opts ...grpc.CallOption) (*PermissionDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PermissionDataResponse)
	err := c.cc.Invoke(ctx, SamplePermissionService_GetPermissionDataForClient_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SamplePermissionServiceServer is the server API for SamplePermissionService service.
// All implementations must embed UnimplementedSamplePermissionServiceServer
// for forward compatibility
type SamplePermissionServiceServer interface {
	// client3
	GetPermissionDataForClient(context.Context, *PermissionDataRequest) (*PermissionDataResponse, error)
	mustEmbedUnimplementedSamplePermissionServiceServer()
}

// UnimplementedSamplePermissionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSamplePermissionServiceServer struct {
}

func (UnimplementedSamplePermissionServiceServer) GetPermissionDataForClient(context.Context, *PermissionDataRequest) (*PermissionDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermissionDataForClient not implemented")
}
func (UnimplementedSamplePermissionServiceServer) mustEmbedUnimplementedSamplePermissionServiceServer() {
}

// UnsafeSamplePermissionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SamplePermissionServiceServer will
// result in compilation errors.
type UnsafeSamplePermissionServiceServer interface {
	mustEmbedUnimplementedSamplePermissionServiceServer()
}

func RegisterSamplePermissionServiceServer(s grpc.ServiceRegistrar, srv SamplePermissionServiceServer) {
	s.RegisterService(&SamplePermissionService_ServiceDesc, srv)
}

func _SamplePermissionService_GetPermissionDataForClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SamplePermissionServiceServer).GetPermissionDataForClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SamplePermissionService_GetPermissionDataForClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SamplePermissionServiceServer).GetPermissionDataForClient(ctx, req.(*PermissionDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SamplePermissionService_ServiceDesc is the grpc.ServiceDesc for SamplePermissionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SamplePermissionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sampleproto.SamplePermissionService",
	HandlerType: (*SamplePermissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPermissionDataForClient",
			Handler:    _SamplePermissionService_GetPermissionDataForClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "permission/permission.proto",
}
