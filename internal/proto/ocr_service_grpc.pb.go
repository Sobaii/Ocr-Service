// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: proto/ocr_service.proto

package proto

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
	OcrService_TestConnection_FullMethodName  = "/ocr_service.OcrService/TestConnection"
	OcrService_SearchFileData_FullMethodName  = "/ocr_service.OcrService/SearchFileData"
	OcrService_ExtractFileData_FullMethodName = "/ocr_service.OcrService/ExtractFileData"
)

// OcrServiceClient is the client API for OcrService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OcrServiceClient interface {
	TestConnection(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error)
	SearchFileData(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*Expenses, error)
	ExtractFileData(ctx context.Context, in *ExtractRequest, opts ...grpc.CallOption) (*ExtractResponse, error)
}

type ocrServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOcrServiceClient(cc grpc.ClientConnInterface) OcrServiceClient {
	return &ocrServiceClient{cc}
}

func (c *ocrServiceClient) TestConnection(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TestResponse)
	err := c.cc.Invoke(ctx, OcrService_TestConnection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocrServiceClient) SearchFileData(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*Expenses, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Expenses)
	err := c.cc.Invoke(ctx, OcrService_SearchFileData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocrServiceClient) ExtractFileData(ctx context.Context, in *ExtractRequest, opts ...grpc.CallOption) (*ExtractResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExtractResponse)
	err := c.cc.Invoke(ctx, OcrService_ExtractFileData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OcrServiceServer is the server API for OcrService service.
// All implementations must embed UnimplementedOcrServiceServer
// for forward compatibility
type OcrServiceServer interface {
	TestConnection(context.Context, *TestRequest) (*TestResponse, error)
	SearchFileData(context.Context, *SearchRequest) (*Expenses, error)
	ExtractFileData(context.Context, *ExtractRequest) (*ExtractResponse, error)
	mustEmbedUnimplementedOcrServiceServer()
}

// UnimplementedOcrServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOcrServiceServer struct {
}

func (UnimplementedOcrServiceServer) TestConnection(context.Context, *TestRequest) (*TestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestConnection not implemented")
}
func (UnimplementedOcrServiceServer) SearchFileData(context.Context, *SearchRequest) (*Expenses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchFileData not implemented")
}
func (UnimplementedOcrServiceServer) ExtractFileData(context.Context, *ExtractRequest) (*ExtractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExtractFileData not implemented")
}
func (UnimplementedOcrServiceServer) mustEmbedUnimplementedOcrServiceServer() {}

// UnsafeOcrServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OcrServiceServer will
// result in compilation errors.
type UnsafeOcrServiceServer interface {
	mustEmbedUnimplementedOcrServiceServer()
}

func RegisterOcrServiceServer(s grpc.ServiceRegistrar, srv OcrServiceServer) {
	s.RegisterService(&OcrService_ServiceDesc, srv)
}

func _OcrService_TestConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcrServiceServer).TestConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OcrService_TestConnection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcrServiceServer).TestConnection(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcrService_SearchFileData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcrServiceServer).SearchFileData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OcrService_SearchFileData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcrServiceServer).SearchFileData(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcrService_ExtractFileData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcrServiceServer).ExtractFileData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OcrService_ExtractFileData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcrServiceServer).ExtractFileData(ctx, req.(*ExtractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OcrService_ServiceDesc is the grpc.ServiceDesc for OcrService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OcrService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ocr_service.OcrService",
	HandlerType: (*OcrServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestConnection",
			Handler:    _OcrService_TestConnection_Handler,
		},
		{
			MethodName: "SearchFileData",
			Handler:    _OcrService_SearchFileData_Handler,
		},
		{
			MethodName: "ExtractFileData",
			Handler:    _OcrService_ExtractFileData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ocr_service.proto",
}
