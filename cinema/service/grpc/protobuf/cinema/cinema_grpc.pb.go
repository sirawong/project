// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cinema

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

// CinemaServiceClient is the client API for CinemaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CinemaServiceClient interface {
	GetCinemas(ctx context.Context, in *CinemaRequest, opts ...grpc.CallOption) (*CinemaReply, error)
}

type cinemaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCinemaServiceClient(cc grpc.ClientConnInterface) CinemaServiceClient {
	return &cinemaServiceClient{cc}
}

func (c *cinemaServiceClient) GetCinemas(ctx context.Context, in *CinemaRequest, opts ...grpc.CallOption) (*CinemaReply, error) {
	out := new(CinemaReply)
	err := c.cc.Invoke(ctx, "/cinema.CinemaService/GetCinemas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CinemaServiceServer is the server API for CinemaService service.
// All implementations should embed UnimplementedCinemaServiceServer
// for forward compatibility
type CinemaServiceServer interface {
	GetCinemas(context.Context, *CinemaRequest) (*CinemaReply, error)
}

// UnimplementedCinemaServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCinemaServiceServer struct {
}

func (UnimplementedCinemaServiceServer) GetCinemas(context.Context, *CinemaRequest) (*CinemaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCinemas not implemented")
}

// UnsafeCinemaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CinemaServiceServer will
// result in compilation errors.
type UnsafeCinemaServiceServer interface {
	mustEmbedUnimplementedCinemaServiceServer()
}

func RegisterCinemaServiceServer(s grpc.ServiceRegistrar, srv CinemaServiceServer) {
	s.RegisterService(&CinemaService_ServiceDesc, srv)
}

func _CinemaService_GetCinemas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CinemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CinemaServiceServer).GetCinemas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cinema.CinemaService/GetCinemas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CinemaServiceServer).GetCinemas(ctx, req.(*CinemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CinemaService_ServiceDesc is the grpc.ServiceDesc for CinemaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CinemaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cinema.CinemaService",
	HandlerType: (*CinemaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCinemas",
			Handler:    _CinemaService_GetCinemas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cinema/cinema.proto",
}
