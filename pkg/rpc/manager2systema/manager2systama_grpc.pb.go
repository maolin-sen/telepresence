// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package manager2systema

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SystemACRUDClient is the client API for SystemACRUD service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemACRUDClient interface {
	// CreateDomain requires that the manager authenticate using an
	// end-user's access token, to perform the action on behalf of that
	// user.
	CreateDomain(ctx context.Context, in *CreateDomainRequest, opts ...grpc.CallOption) (*CreateDomainResponse, error)
	// RemoveDomain removes a domain that was previously created by the
	// same manager using CreateDomain.  The manager can take this
	// action itself, not on behalf of the user that created the domain,
	// so this requires that the manager authenticate itself, but does
	// not require an end-user's token.
	RemoveDomain(ctx context.Context, in *RemoveDomainRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type systemACRUDClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemACRUDClient(cc grpc.ClientConnInterface) SystemACRUDClient {
	return &systemACRUDClient{cc}
}

func (c *systemACRUDClient) CreateDomain(ctx context.Context, in *CreateDomainRequest, opts ...grpc.CallOption) (*CreateDomainResponse, error) {
	out := new(CreateDomainResponse)
	err := c.cc.Invoke(ctx, "/telepresence.manager2systema.SystemACRUD/CreateDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemACRUDClient) RemoveDomain(ctx context.Context, in *RemoveDomainRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/telepresence.manager2systema.SystemACRUD/RemoveDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemACRUDServer is the server API for SystemACRUD service.
// All implementations must embed UnimplementedSystemACRUDServer
// for forward compatibility
type SystemACRUDServer interface {
	// CreateDomain requires that the manager authenticate using an
	// end-user's access token, to perform the action on behalf of that
	// user.
	CreateDomain(context.Context, *CreateDomainRequest) (*CreateDomainResponse, error)
	// RemoveDomain removes a domain that was previously created by the
	// same manager using CreateDomain.  The manager can take this
	// action itself, not on behalf of the user that created the domain,
	// so this requires that the manager authenticate itself, but does
	// not require an end-user's token.
	RemoveDomain(context.Context, *RemoveDomainRequest) (*empty.Empty, error)
	mustEmbedUnimplementedSystemACRUDServer()
}

// UnimplementedSystemACRUDServer must be embedded to have forward compatible implementations.
type UnimplementedSystemACRUDServer struct {
}

func (UnimplementedSystemACRUDServer) CreateDomain(context.Context, *CreateDomainRequest) (*CreateDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDomain not implemented")
}
func (UnimplementedSystemACRUDServer) RemoveDomain(context.Context, *RemoveDomainRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDomain not implemented")
}
func (UnimplementedSystemACRUDServer) mustEmbedUnimplementedSystemACRUDServer() {}

// UnsafeSystemACRUDServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemACRUDServer will
// result in compilation errors.
type UnsafeSystemACRUDServer interface {
	mustEmbedUnimplementedSystemACRUDServer()
}

func RegisterSystemACRUDServer(s grpc.ServiceRegistrar, srv SystemACRUDServer) {
	s.RegisterService(&_SystemACRUD_serviceDesc, srv)
}

func _SystemACRUD_CreateDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemACRUDServer).CreateDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.manager2systema.SystemACRUD/CreateDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemACRUDServer).CreateDomain(ctx, req.(*CreateDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemACRUD_RemoveDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemACRUDServer).RemoveDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.manager2systema.SystemACRUD/RemoveDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemACRUDServer).RemoveDomain(ctx, req.(*RemoveDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SystemACRUD_serviceDesc = grpc.ServiceDesc{
	ServiceName: "telepresence.manager2systema.SystemACRUD",
	HandlerType: (*SystemACRUDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDomain",
			Handler:    _SystemACRUD_CreateDomain_Handler,
		},
		{
			MethodName: "RemoveDomain",
			Handler:    _SystemACRUD_RemoveDomain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/manager2systema/manager2systama.proto",
}

// SystemAProxyClient is the client API for SystemAProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemAProxyClient interface {
	// ReverseConnection establishes a stream that is used for System A
	// to send gRPC requests back to the manager.  This requires that
	// the manager authenticate itself, but does not require an
	// end-user's token.
	ReverseConnection(ctx context.Context, opts ...grpc.CallOption) (SystemAProxy_ReverseConnectionClient, error)
}

type systemAProxyClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemAProxyClient(cc grpc.ClientConnInterface) SystemAProxyClient {
	return &systemAProxyClient{cc}
}

func (c *systemAProxyClient) ReverseConnection(ctx context.Context, opts ...grpc.CallOption) (SystemAProxy_ReverseConnectionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SystemAProxy_serviceDesc.Streams[0], "/telepresence.manager2systema.SystemAProxy/ReverseConnection", opts...)
	if err != nil {
		return nil, err
	}
	x := &systemAProxyReverseConnectionClient{stream}
	return x, nil
}

type SystemAProxy_ReverseConnectionClient interface {
	Send(*Chunk) error
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type systemAProxyReverseConnectionClient struct {
	grpc.ClientStream
}

func (x *systemAProxyReverseConnectionClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *systemAProxyReverseConnectionClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SystemAProxyServer is the server API for SystemAProxy service.
// All implementations must embed UnimplementedSystemAProxyServer
// for forward compatibility
type SystemAProxyServer interface {
	// ReverseConnection establishes a stream that is used for System A
	// to send gRPC requests back to the manager.  This requires that
	// the manager authenticate itself, but does not require an
	// end-user's token.
	ReverseConnection(SystemAProxy_ReverseConnectionServer) error
	mustEmbedUnimplementedSystemAProxyServer()
}

// UnimplementedSystemAProxyServer must be embedded to have forward compatible implementations.
type UnimplementedSystemAProxyServer struct {
}

func (UnimplementedSystemAProxyServer) ReverseConnection(SystemAProxy_ReverseConnectionServer) error {
	return status.Errorf(codes.Unimplemented, "method ReverseConnection not implemented")
}
func (UnimplementedSystemAProxyServer) mustEmbedUnimplementedSystemAProxyServer() {}

// UnsafeSystemAProxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemAProxyServer will
// result in compilation errors.
type UnsafeSystemAProxyServer interface {
	mustEmbedUnimplementedSystemAProxyServer()
}

func RegisterSystemAProxyServer(s grpc.ServiceRegistrar, srv SystemAProxyServer) {
	s.RegisterService(&_SystemAProxy_serviceDesc, srv)
}

func _SystemAProxy_ReverseConnection_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SystemAProxyServer).ReverseConnection(&systemAProxyReverseConnectionServer{stream})
}

type SystemAProxy_ReverseConnectionServer interface {
	Send(*Chunk) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type systemAProxyReverseConnectionServer struct {
	grpc.ServerStream
}

func (x *systemAProxyReverseConnectionServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

func (x *systemAProxyReverseConnectionServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SystemAProxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "telepresence.manager2systema.SystemAProxy",
	HandlerType: (*SystemAProxyServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReverseConnection",
			Handler:       _SystemAProxy_ReverseConnection_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "rpc/manager2systema/manager2systama.proto",
}
