// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: trading.proto

package trading_service

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

// PriceServiceClient is the client API for PriceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PriceServiceClient interface {
	OpenPosition(ctx context.Context, in *OpenPositionRequest, opts ...grpc.CallOption) (*OpenPositionResponse, error)
}

type priceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPriceServiceClient(cc grpc.ClientConnInterface) PriceServiceClient {
	return &priceServiceClient{cc}
}

func (c *priceServiceClient) OpenPosition(ctx context.Context, in *OpenPositionRequest, opts ...grpc.CallOption) (*OpenPositionResponse, error) {
	out := new(OpenPositionResponse)
	err := c.cc.Invoke(ctx, "/PriceService/OpenPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PriceServiceServer is the server API for PriceService service.
// All implementations must embed UnimplementedPriceServiceServer
// for forward compatibility
type PriceServiceServer interface {
	OpenPosition(context.Context, *OpenPositionRequest) (*OpenPositionResponse, error)
	mustEmbedUnimplementedPriceServiceServer()
}

// UnimplementedPriceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPriceServiceServer struct {
}

func (UnimplementedPriceServiceServer) OpenPosition(context.Context, *OpenPositionRequest) (*OpenPositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenPosition not implemented")
}
func (UnimplementedPriceServiceServer) mustEmbedUnimplementedPriceServiceServer() {}

// UnsafePriceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PriceServiceServer will
// result in compilation errors.
type UnsafePriceServiceServer interface {
	mustEmbedUnimplementedPriceServiceServer()
}

func RegisterPriceServiceServer(s grpc.ServiceRegistrar, srv PriceServiceServer) {
	s.RegisterService(&PriceService_ServiceDesc, srv)
}

func _PriceService_OpenPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceServiceServer).OpenPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PriceService/OpenPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceServiceServer).OpenPosition(ctx, req.(*OpenPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PriceService_ServiceDesc is the grpc.ServiceDesc for PriceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PriceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PriceService",
	HandlerType: (*PriceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenPosition",
			Handler:    _PriceService_OpenPosition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trading.proto",
}
