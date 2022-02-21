// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package echo

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

// EchoClient is the client API for Echo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EchoClient interface {
	// BidirectionalStreaming
	BidirectionalStreamingEcho(ctx context.Context, opts ...grpc.CallOption) (Echo_BidirectionalStreamingEchoClient, error)
}

type echoClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoClient(cc grpc.ClientConnInterface) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) BidirectionalStreamingEcho(ctx context.Context, opts ...grpc.CallOption) (Echo_BidirectionalStreamingEchoClient, error) {
	stream, err := c.cc.NewStream(ctx, &Echo_ServiceDesc.Streams[0], "/echo.Echo/BidirectionalStreamingEcho", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoBidirectionalStreamingEchoClient{stream}
	return x, nil
}

type Echo_BidirectionalStreamingEchoClient interface {
	Send(*EchoRequest) error
	Recv() (*EchoResponse, error)
	grpc.ClientStream
}

type echoBidirectionalStreamingEchoClient struct {
	grpc.ClientStream
}

func (x *echoBidirectionalStreamingEchoClient) Send(m *EchoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoBidirectionalStreamingEchoClient) Recv() (*EchoResponse, error) {
	m := new(EchoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EchoServer is the server API for Echo service.
// All implementations must embed UnimplementedEchoServer
// for forward compatibility
type EchoServer interface {
	// BidirectionalStreaming
	BidirectionalStreamingEcho(Echo_BidirectionalStreamingEchoServer) error
	mustEmbedUnimplementedEchoServer()
}

// UnimplementedEchoServer must be embedded to have forward compatible implementations.
type UnimplementedEchoServer struct {
}

func (UnimplementedEchoServer) BidirectionalStreamingEcho(Echo_BidirectionalStreamingEchoServer) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStreamingEcho not implemented")
}
func (UnimplementedEchoServer) mustEmbedUnimplementedEchoServer() {}

// UnsafeEchoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EchoServer will
// result in compilation errors.
type UnsafeEchoServer interface {
	mustEmbedUnimplementedEchoServer()
}

func RegisterEchoServer(s grpc.ServiceRegistrar, srv EchoServer) {
	s.RegisterService(&Echo_ServiceDesc, srv)
}

func _Echo_BidirectionalStreamingEcho_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServer).BidirectionalStreamingEcho(&echoBidirectionalStreamingEchoServer{stream})
}

type Echo_BidirectionalStreamingEchoServer interface {
	Send(*EchoResponse) error
	Recv() (*EchoRequest, error)
	grpc.ServerStream
}

type echoBidirectionalStreamingEchoServer struct {
	grpc.ServerStream
}

func (x *echoBidirectionalStreamingEchoServer) Send(m *EchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoBidirectionalStreamingEchoServer) Recv() (*EchoRequest, error) {
	m := new(EchoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Echo_ServiceDesc is the grpc.ServiceDesc for Echo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Echo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "echo.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BidirectionalStreamingEcho",
			Handler:       _Echo_BidirectionalStreamingEcho_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "stream/bidirectionalStream/echo/echo.proto",
}