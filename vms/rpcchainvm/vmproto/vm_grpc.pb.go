// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package vmproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VMClient is the client API for VM service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VMClient interface {
	Initialize(ctx context.Context, in *InitializeRequest, opts ...grpc.CallOption) (*InitializeResponse, error)
	Bootstrapping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Bootstrapped(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Shutdown(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateHandlers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CreateHandlersResponse, error)
	CreateStaticHandlers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CreateStaticHandlersResponse, error)
	BuildBlock(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BuildBlockResponse, error)
	ParseBlock(ctx context.Context, in *ParseBlockRequest, opts ...grpc.CallOption) (*ParseBlockResponse, error)
	GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*GetBlockResponse, error)
	SetPreference(ctx context.Context, in *SetPreferenceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Health(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HealthResponse, error)
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
	AppRequest(ctx context.Context, in *AppRequestMsg, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AppRequestFailed(ctx context.Context, in *AppRequestFailedMsg, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AppResponse(ctx context.Context, in *AppResponseMsg, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AppGossip(ctx context.Context, in *AppGossipMsg, opts ...grpc.CallOption) (*emptypb.Empty, error)
	BlockVerify(ctx context.Context, in *BlockVerifyRequest, opts ...grpc.CallOption) (*BlockVerifyResponse, error)
	BlockAccept(ctx context.Context, in *BlockAcceptRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	BlockReject(ctx context.Context, in *BlockRejectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type vMClient struct {
	cc grpc.ClientConnInterface
}

func NewVMClient(cc grpc.ClientConnInterface) VMClient {
	return &vMClient{cc}
}

func (c *vMClient) Initialize(ctx context.Context, in *InitializeRequest, opts ...grpc.CallOption) (*InitializeResponse, error) {
	out := new(InitializeResponse)
	err := c.cc.Invoke(ctx, "/vmproto.VM/Initialize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMClient) Bootstrapping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/vmproto.VM/Bootstrapping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMClient) Bootstrapped(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/vmproto.VM/Bootstrapped", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMClient) Shutdown(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/vmproto.VM/Shutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMClient) CreateHandlers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CreateHandlersResponse, error) {
	out := new(CreateHandlersResponse)
	err := c.cc.Invoke(ctx, "/vmproto.VM/CreateHandlers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMClient) CreateStaticHandlers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CreateStaticHandlersResponse, error) {
	out := new(CreateStaticHandlersResponse)
	err := c.cc.Invoke(ctx, "/vmproto.VM/CreateStaticHandlers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMClient) BuildBlock(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BuildBlockResponse, error) {
	out := new(BuildBlockResponse)
	err := c.cc.Invoke(ctx, "/vmproto.VM/BuildBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMClient) ParseBlock(ctx context.Context, in *ParseBlockRequest, opts ...grpc.CallOption) (*ParseBlockResponse, error) {
	out := new(ParseBlockResponse)
	err := c.cc.Invoke(ctx, "/vmproto.VM/ParseBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMClient) GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*GetBlockResponse, error) {
	out := new(GetBlockResponse)
	err := c.cc.Invoke(ctx, "/vmproto.VM/GetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMClient) SetPreference(ctx context.Context, in *SetPreferenceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/vmproto.VM/SetPreference", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMClient) Health(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/vmproto.VM/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/vmproto.VM/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMClient) AppRequest(ctx context.Context, in *AppRequestMsg, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/vmproto.VM/AppRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMClient) AppRequestFailed(ctx context.Context, in *AppRequestFailedMsg, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/vmproto.VM/AppRequestFailed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMClient) AppResponse(ctx context.Context, in *AppResponseMsg, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/vmproto.VM/AppResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMClient) AppGossip(ctx context.Context, in *AppGossipMsg, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/vmproto.VM/AppGossip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMClient) BlockVerify(ctx context.Context, in *BlockVerifyRequest, opts ...grpc.CallOption) (*BlockVerifyResponse, error) {
	out := new(BlockVerifyResponse)
	err := c.cc.Invoke(ctx, "/vmproto.VM/BlockVerify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMClient) BlockAccept(ctx context.Context, in *BlockAcceptRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/vmproto.VM/BlockAccept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMClient) BlockReject(ctx context.Context, in *BlockRejectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/vmproto.VM/BlockReject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VMServer is the server API for VM service.
// All implementations must embed UnimplementedVMServer
// for forward compatibility
type VMServer interface {
	Initialize(context.Context, *InitializeRequest) (*InitializeResponse, error)
	Bootstrapping(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	Bootstrapped(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	Shutdown(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	CreateHandlers(context.Context, *emptypb.Empty) (*CreateHandlersResponse, error)
	CreateStaticHandlers(context.Context, *emptypb.Empty) (*CreateStaticHandlersResponse, error)
	BuildBlock(context.Context, *emptypb.Empty) (*BuildBlockResponse, error)
	ParseBlock(context.Context, *ParseBlockRequest) (*ParseBlockResponse, error)
	GetBlock(context.Context, *GetBlockRequest) (*GetBlockResponse, error)
	SetPreference(context.Context, *SetPreferenceRequest) (*emptypb.Empty, error)
	Health(context.Context, *emptypb.Empty) (*HealthResponse, error)
	Version(context.Context, *emptypb.Empty) (*VersionResponse, error)
	AppRequest(context.Context, *AppRequestMsg) (*emptypb.Empty, error)
	AppRequestFailed(context.Context, *AppRequestFailedMsg) (*emptypb.Empty, error)
	AppResponse(context.Context, *AppResponseMsg) (*emptypb.Empty, error)
	AppGossip(context.Context, *AppGossipMsg) (*emptypb.Empty, error)
	BlockVerify(context.Context, *BlockVerifyRequest) (*BlockVerifyResponse, error)
	BlockAccept(context.Context, *BlockAcceptRequest) (*emptypb.Empty, error)
	BlockReject(context.Context, *BlockRejectRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedVMServer()
}

// UnimplementedVMServer must be embedded to have forward compatible implementations.
type UnimplementedVMServer struct {
}

func (UnimplementedVMServer) Initialize(context.Context, *InitializeRequest) (*InitializeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Initialize not implemented")
}
func (UnimplementedVMServer) Bootstrapping(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bootstrapping not implemented")
}
func (UnimplementedVMServer) Bootstrapped(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bootstrapped not implemented")
}
func (UnimplementedVMServer) Shutdown(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shutdown not implemented")
}
func (UnimplementedVMServer) CreateHandlers(context.Context, *emptypb.Empty) (*CreateHandlersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHandlers not implemented")
}
func (UnimplementedVMServer) CreateStaticHandlers(context.Context, *emptypb.Empty) (*CreateStaticHandlersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStaticHandlers not implemented")
}
func (UnimplementedVMServer) BuildBlock(context.Context, *emptypb.Empty) (*BuildBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildBlock not implemented")
}
func (UnimplementedVMServer) ParseBlock(context.Context, *ParseBlockRequest) (*ParseBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseBlock not implemented")
}
func (UnimplementedVMServer) GetBlock(context.Context, *GetBlockRequest) (*GetBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlock not implemented")
}
func (UnimplementedVMServer) SetPreference(context.Context, *SetPreferenceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPreference not implemented")
}
func (UnimplementedVMServer) Health(context.Context, *emptypb.Empty) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedVMServer) Version(context.Context, *emptypb.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedVMServer) AppRequest(context.Context, *AppRequestMsg) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppRequest not implemented")
}
func (UnimplementedVMServer) AppRequestFailed(context.Context, *AppRequestFailedMsg) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppRequestFailed not implemented")
}
func (UnimplementedVMServer) AppResponse(context.Context, *AppResponseMsg) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppResponse not implemented")
}
func (UnimplementedVMServer) AppGossip(context.Context, *AppGossipMsg) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppGossip not implemented")
}
func (UnimplementedVMServer) BlockVerify(context.Context, *BlockVerifyRequest) (*BlockVerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockVerify not implemented")
}
func (UnimplementedVMServer) BlockAccept(context.Context, *BlockAcceptRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockAccept not implemented")
}
func (UnimplementedVMServer) BlockReject(context.Context, *BlockRejectRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockReject not implemented")
}
func (UnimplementedVMServer) mustEmbedUnimplementedVMServer() {}

// UnsafeVMServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VMServer will
// result in compilation errors.
type UnsafeVMServer interface {
	mustEmbedUnimplementedVMServer()
}

func RegisterVMServer(s grpc.ServiceRegistrar, srv VMServer) {
	s.RegisterService(&VM_ServiceDesc, srv)
}

func _VM_Initialize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitializeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).Initialize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vmproto.VM/Initialize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).Initialize(ctx, req.(*InitializeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VM_Bootstrapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).Bootstrapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vmproto.VM/Bootstrapping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).Bootstrapping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VM_Bootstrapped_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).Bootstrapped(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vmproto.VM/Bootstrapped",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).Bootstrapped(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VM_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vmproto.VM/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).Shutdown(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VM_CreateHandlers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).CreateHandlers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vmproto.VM/CreateHandlers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).CreateHandlers(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VM_CreateStaticHandlers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).CreateStaticHandlers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vmproto.VM/CreateStaticHandlers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).CreateStaticHandlers(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VM_BuildBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).BuildBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vmproto.VM/BuildBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).BuildBlock(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VM_ParseBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).ParseBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vmproto.VM/ParseBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).ParseBlock(ctx, req.(*ParseBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VM_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vmproto.VM/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).GetBlock(ctx, req.(*GetBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VM_SetPreference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPreferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).SetPreference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vmproto.VM/SetPreference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).SetPreference(ctx, req.(*SetPreferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VM_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vmproto.VM/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).Health(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VM_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vmproto.VM/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VM_AppRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppRequestMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).AppRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vmproto.VM/AppRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).AppRequest(ctx, req.(*AppRequestMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _VM_AppRequestFailed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppRequestFailedMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).AppRequestFailed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vmproto.VM/AppRequestFailed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).AppRequestFailed(ctx, req.(*AppRequestFailedMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _VM_AppResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppResponseMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).AppResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vmproto.VM/AppResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).AppResponse(ctx, req.(*AppResponseMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _VM_AppGossip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppGossipMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).AppGossip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vmproto.VM/AppGossip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).AppGossip(ctx, req.(*AppGossipMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _VM_BlockVerify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockVerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).BlockVerify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vmproto.VM/BlockVerify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).BlockVerify(ctx, req.(*BlockVerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VM_BlockAccept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockAcceptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).BlockAccept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vmproto.VM/BlockAccept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).BlockAccept(ctx, req.(*BlockAcceptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VM_BlockReject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockRejectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).BlockReject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vmproto.VM/BlockReject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).BlockReject(ctx, req.(*BlockRejectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VM_ServiceDesc is the grpc.ServiceDesc for VM service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VM_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vmproto.VM",
	HandlerType: (*VMServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Initialize",
			Handler:    _VM_Initialize_Handler,
		},
		{
			MethodName: "Bootstrapping",
			Handler:    _VM_Bootstrapping_Handler,
		},
		{
			MethodName: "Bootstrapped",
			Handler:    _VM_Bootstrapped_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _VM_Shutdown_Handler,
		},
		{
			MethodName: "CreateHandlers",
			Handler:    _VM_CreateHandlers_Handler,
		},
		{
			MethodName: "CreateStaticHandlers",
			Handler:    _VM_CreateStaticHandlers_Handler,
		},
		{
			MethodName: "BuildBlock",
			Handler:    _VM_BuildBlock_Handler,
		},
		{
			MethodName: "ParseBlock",
			Handler:    _VM_ParseBlock_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _VM_GetBlock_Handler,
		},
		{
			MethodName: "SetPreference",
			Handler:    _VM_SetPreference_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _VM_Health_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _VM_Version_Handler,
		},
		{
			MethodName: "AppRequest",
			Handler:    _VM_AppRequest_Handler,
		},
		{
			MethodName: "AppRequestFailed",
			Handler:    _VM_AppRequestFailed_Handler,
		},
		{
			MethodName: "AppResponse",
			Handler:    _VM_AppResponse_Handler,
		},
		{
			MethodName: "AppGossip",
			Handler:    _VM_AppGossip_Handler,
		},
		{
			MethodName: "BlockVerify",
			Handler:    _VM_BlockVerify_Handler,
		},
		{
			MethodName: "BlockAccept",
			Handler:    _VM_BlockAccept_Handler,
		},
		{
			MethodName: "BlockReject",
			Handler:    _VM_BlockReject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vm.proto",
}
