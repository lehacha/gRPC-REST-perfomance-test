// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.2
// source: TextProcessor.proto

package workspace

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

// TextProcessorClient is the client API for TextProcessor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TextProcessorClient interface {
	Process(ctx context.Context, in *ProcessRequest, opts ...grpc.CallOption) (*ProcessResponse, error)
}

type textProcessorClient struct {
	cc grpc.ClientConnInterface
}

func NewTextProcessorClient(cc grpc.ClientConnInterface) TextProcessorClient {
	return &textProcessorClient{cc}
}

func (c *textProcessorClient) Process(ctx context.Context, in *ProcessRequest, opts ...grpc.CallOption) (*ProcessResponse, error) {
	out := new(ProcessResponse)
	err := c.cc.Invoke(ctx, "/TextProcessor.TextProcessor/Process", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TextProcessorServer is the server API for TextProcessor service.
// All implementations must embed UnimplementedTextProcessorServer
// for forward compatibility
type TextProcessorServer interface {
	Process(context.Context, *ProcessRequest) (*ProcessResponse, error)
	mustEmbedUnimplementedTextProcessorServer()
}

// UnimplementedTextProcessorServer must be embedded to have forward compatible implementations.
type UnimplementedTextProcessorServer struct {
}

func (UnimplementedTextProcessorServer) Process(context.Context, *ProcessRequest) (*ProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Process not implemented")
}
func (UnimplementedTextProcessorServer) mustEmbedUnimplementedTextProcessorServer() {}

// UnsafeTextProcessorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TextProcessorServer will
// result in compilation errors.
type UnsafeTextProcessorServer interface {
	mustEmbedUnimplementedTextProcessorServer()
}

func RegisterTextProcessorServer(s grpc.ServiceRegistrar, srv TextProcessorServer) {
	s.RegisterService(&TextProcessor_ServiceDesc, srv)
}

func _TextProcessor_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextProcessorServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TextProcessor.TextProcessor/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextProcessorServer).Process(ctx, req.(*ProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TextProcessor_ServiceDesc is the grpc.ServiceDesc for TextProcessor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TextProcessor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TextProcessor.TextProcessor",
	HandlerType: (*TextProcessorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Process",
			Handler:    _TextProcessor_Process_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "TextProcessor.proto",
}
