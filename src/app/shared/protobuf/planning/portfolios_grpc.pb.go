// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: src/app/shared/protobuf/planning/portfolios.proto

package portfolios

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

// PortfoliosClient is the client API for Portfolios service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortfoliosClient interface {
	Create(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Update(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Read(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type portfoliosClient struct {
	cc grpc.ClientConnInterface
}

func NewPortfoliosClient(cc grpc.ClientConnInterface) PortfoliosClient {
	return &portfoliosClient{cc}
}

func (c *portfoliosClient) Create(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Portfolios/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfoliosClient) Update(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Portfolios/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfoliosClient) Read(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Portfolios/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfoliosClient) Delete(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Portfolios/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortfoliosServer is the server API for Portfolios service.
// All implementations must embed UnimplementedPortfoliosServer
// for forward compatibility
type PortfoliosServer interface {
	Create(context.Context, *Request) (*Response, error)
	Update(context.Context, *Request) (*Response, error)
	Read(context.Context, *Request) (*Response, error)
	Delete(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedPortfoliosServer()
}

// UnimplementedPortfoliosServer must be embedded to have forward compatible implementations.
type UnimplementedPortfoliosServer struct {
}

func (UnimplementedPortfoliosServer) Create(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPortfoliosServer) Update(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPortfoliosServer) Read(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedPortfoliosServer) Delete(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPortfoliosServer) mustEmbedUnimplementedPortfoliosServer() {}

// UnsafePortfoliosServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortfoliosServer will
// result in compilation errors.
type UnsafePortfoliosServer interface {
	mustEmbedUnimplementedPortfoliosServer()
}

func RegisterPortfoliosServer(s grpc.ServiceRegistrar, srv PortfoliosServer) {
	s.RegisterService(&Portfolios_ServiceDesc, srv)
}

func _Portfolios_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfoliosServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Portfolios/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfoliosServer).Create(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Portfolios_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfoliosServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Portfolios/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfoliosServer).Update(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Portfolios_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfoliosServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Portfolios/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfoliosServer).Read(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Portfolios_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfoliosServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Portfolios/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfoliosServer).Delete(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Portfolios_ServiceDesc is the grpc.ServiceDesc for Portfolios service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Portfolios_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Portfolios",
	HandlerType: (*PortfoliosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Portfolios_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Portfolios_Update_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Portfolios_Read_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Portfolios_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/app/shared/protobuf/planning/portfolios.proto",
}