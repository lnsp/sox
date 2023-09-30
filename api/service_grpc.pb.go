// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// SoxClient is the client API for Sox service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SoxClient interface {
	CreateMachine(ctx context.Context, in *CreateMachineRequest, opts ...grpc.CallOption) (*CreateMachineResponse, error)
	ListMachines(ctx context.Context, in *ListMachinesRequest, opts ...grpc.CallOption) (*ListMachinesResponse, error)
	GetMachineDetails(ctx context.Context, in *GetMachineDetailsRequest, opts ...grpc.CallOption) (*GetMachineDetailsResponse, error)
	DeleteMachine(ctx context.Context, in *DeleteMachineRequest, opts ...grpc.CallOption) (*DeleteMachineResponse, error)
	TriggerMachine(ctx context.Context, in *TriggerMachineRequest, opts ...grpc.CallOption) (*TriggerMachineResponse, error)
	CreateSSHKey(ctx context.Context, in *CreateSSHKeyRequest, opts ...grpc.CallOption) (*CreateSSHKeyResponse, error)
	ListSSHKeys(ctx context.Context, in *ListSSHKeysRequest, opts ...grpc.CallOption) (*ListSSHKeysResponse, error)
	DeleteSSHKey(ctx context.Context, in *DeleteSSHKeyRequest, opts ...grpc.CallOption) (*DeleteSSHKeyResponse, error)
	ListImages(ctx context.Context, in *ListImagesRequest, opts ...grpc.CallOption) (*ListImagesResponse, error)
	ListNetworks(ctx context.Context, in *ListNetworksRequest, opts ...grpc.CallOption) (*ListNetworksResponse, error)
	CreateNetwork(ctx context.Context, in *CreateNetworkRequest, opts ...grpc.CallOption) (*CreateNetworkResponse, error)
	ListActivities(ctx context.Context, in *ListActivitiesRequest, opts ...grpc.CallOption) (*ListActivitiesResponse, error)
}

type soxClient struct {
	cc grpc.ClientConnInterface
}

func NewSoxClient(cc grpc.ClientConnInterface) SoxClient {
	return &soxClient{cc}
}

func (c *soxClient) CreateMachine(ctx context.Context, in *CreateMachineRequest, opts ...grpc.CallOption) (*CreateMachineResponse, error) {
	out := new(CreateMachineResponse)
	err := c.cc.Invoke(ctx, "/sox.v1.Sox/CreateMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soxClient) ListMachines(ctx context.Context, in *ListMachinesRequest, opts ...grpc.CallOption) (*ListMachinesResponse, error) {
	out := new(ListMachinesResponse)
	err := c.cc.Invoke(ctx, "/sox.v1.Sox/ListMachines", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soxClient) GetMachineDetails(ctx context.Context, in *GetMachineDetailsRequest, opts ...grpc.CallOption) (*GetMachineDetailsResponse, error) {
	out := new(GetMachineDetailsResponse)
	err := c.cc.Invoke(ctx, "/sox.v1.Sox/GetMachineDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soxClient) DeleteMachine(ctx context.Context, in *DeleteMachineRequest, opts ...grpc.CallOption) (*DeleteMachineResponse, error) {
	out := new(DeleteMachineResponse)
	err := c.cc.Invoke(ctx, "/sox.v1.Sox/DeleteMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soxClient) TriggerMachine(ctx context.Context, in *TriggerMachineRequest, opts ...grpc.CallOption) (*TriggerMachineResponse, error) {
	out := new(TriggerMachineResponse)
	err := c.cc.Invoke(ctx, "/sox.v1.Sox/TriggerMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soxClient) CreateSSHKey(ctx context.Context, in *CreateSSHKeyRequest, opts ...grpc.CallOption) (*CreateSSHKeyResponse, error) {
	out := new(CreateSSHKeyResponse)
	err := c.cc.Invoke(ctx, "/sox.v1.Sox/CreateSSHKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soxClient) ListSSHKeys(ctx context.Context, in *ListSSHKeysRequest, opts ...grpc.CallOption) (*ListSSHKeysResponse, error) {
	out := new(ListSSHKeysResponse)
	err := c.cc.Invoke(ctx, "/sox.v1.Sox/ListSSHKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soxClient) DeleteSSHKey(ctx context.Context, in *DeleteSSHKeyRequest, opts ...grpc.CallOption) (*DeleteSSHKeyResponse, error) {
	out := new(DeleteSSHKeyResponse)
	err := c.cc.Invoke(ctx, "/sox.v1.Sox/DeleteSSHKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soxClient) ListImages(ctx context.Context, in *ListImagesRequest, opts ...grpc.CallOption) (*ListImagesResponse, error) {
	out := new(ListImagesResponse)
	err := c.cc.Invoke(ctx, "/sox.v1.Sox/ListImages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soxClient) ListNetworks(ctx context.Context, in *ListNetworksRequest, opts ...grpc.CallOption) (*ListNetworksResponse, error) {
	out := new(ListNetworksResponse)
	err := c.cc.Invoke(ctx, "/sox.v1.Sox/ListNetworks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soxClient) CreateNetwork(ctx context.Context, in *CreateNetworkRequest, opts ...grpc.CallOption) (*CreateNetworkResponse, error) {
	out := new(CreateNetworkResponse)
	err := c.cc.Invoke(ctx, "/sox.v1.Sox/CreateNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *soxClient) ListActivities(ctx context.Context, in *ListActivitiesRequest, opts ...grpc.CallOption) (*ListActivitiesResponse, error) {
	out := new(ListActivitiesResponse)
	err := c.cc.Invoke(ctx, "/sox.v1.Sox/ListActivities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SoxServer is the server API for Sox service.
// All implementations must embed UnimplementedSoxServer
// for forward compatibility
type SoxServer interface {
	CreateMachine(context.Context, *CreateMachineRequest) (*CreateMachineResponse, error)
	ListMachines(context.Context, *ListMachinesRequest) (*ListMachinesResponse, error)
	GetMachineDetails(context.Context, *GetMachineDetailsRequest) (*GetMachineDetailsResponse, error)
	DeleteMachine(context.Context, *DeleteMachineRequest) (*DeleteMachineResponse, error)
	TriggerMachine(context.Context, *TriggerMachineRequest) (*TriggerMachineResponse, error)
	CreateSSHKey(context.Context, *CreateSSHKeyRequest) (*CreateSSHKeyResponse, error)
	ListSSHKeys(context.Context, *ListSSHKeysRequest) (*ListSSHKeysResponse, error)
	DeleteSSHKey(context.Context, *DeleteSSHKeyRequest) (*DeleteSSHKeyResponse, error)
	ListImages(context.Context, *ListImagesRequest) (*ListImagesResponse, error)
	ListNetworks(context.Context, *ListNetworksRequest) (*ListNetworksResponse, error)
	CreateNetwork(context.Context, *CreateNetworkRequest) (*CreateNetworkResponse, error)
	ListActivities(context.Context, *ListActivitiesRequest) (*ListActivitiesResponse, error)
	mustEmbedUnimplementedSoxServer()
}

// UnimplementedSoxServer must be embedded to have forward compatible implementations.
type UnimplementedSoxServer struct {
}

func (UnimplementedSoxServer) CreateMachine(context.Context, *CreateMachineRequest) (*CreateMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMachine not implemented")
}
func (UnimplementedSoxServer) ListMachines(context.Context, *ListMachinesRequest) (*ListMachinesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMachines not implemented")
}
func (UnimplementedSoxServer) GetMachineDetails(context.Context, *GetMachineDetailsRequest) (*GetMachineDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMachineDetails not implemented")
}
func (UnimplementedSoxServer) DeleteMachine(context.Context, *DeleteMachineRequest) (*DeleteMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMachine not implemented")
}
func (UnimplementedSoxServer) TriggerMachine(context.Context, *TriggerMachineRequest) (*TriggerMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggerMachine not implemented")
}
func (UnimplementedSoxServer) CreateSSHKey(context.Context, *CreateSSHKeyRequest) (*CreateSSHKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSSHKey not implemented")
}
func (UnimplementedSoxServer) ListSSHKeys(context.Context, *ListSSHKeysRequest) (*ListSSHKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSSHKeys not implemented")
}
func (UnimplementedSoxServer) DeleteSSHKey(context.Context, *DeleteSSHKeyRequest) (*DeleteSSHKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSSHKey not implemented")
}
func (UnimplementedSoxServer) ListImages(context.Context, *ListImagesRequest) (*ListImagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListImages not implemented")
}
func (UnimplementedSoxServer) ListNetworks(context.Context, *ListNetworksRequest) (*ListNetworksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNetworks not implemented")
}
func (UnimplementedSoxServer) CreateNetwork(context.Context, *CreateNetworkRequest) (*CreateNetworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNetwork not implemented")
}
func (UnimplementedSoxServer) ListActivities(context.Context, *ListActivitiesRequest) (*ListActivitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListActivities not implemented")
}
func (UnimplementedSoxServer) mustEmbedUnimplementedSoxServer() {}

// UnsafeSoxServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SoxServer will
// result in compilation errors.
type UnsafeSoxServer interface {
	mustEmbedUnimplementedSoxServer()
}

func RegisterSoxServer(s grpc.ServiceRegistrar, srv SoxServer) {
	s.RegisterService(&Sox_ServiceDesc, srv)
}

func _Sox_CreateMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoxServer).CreateMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sox.v1.Sox/CreateMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoxServer).CreateMachine(ctx, req.(*CreateMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sox_ListMachines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMachinesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoxServer).ListMachines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sox.v1.Sox/ListMachines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoxServer).ListMachines(ctx, req.(*ListMachinesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sox_GetMachineDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMachineDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoxServer).GetMachineDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sox.v1.Sox/GetMachineDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoxServer).GetMachineDetails(ctx, req.(*GetMachineDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sox_DeleteMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoxServer).DeleteMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sox.v1.Sox/DeleteMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoxServer).DeleteMachine(ctx, req.(*DeleteMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sox_TriggerMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoxServer).TriggerMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sox.v1.Sox/TriggerMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoxServer).TriggerMachine(ctx, req.(*TriggerMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sox_CreateSSHKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSSHKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoxServer).CreateSSHKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sox.v1.Sox/CreateSSHKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoxServer).CreateSSHKey(ctx, req.(*CreateSSHKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sox_ListSSHKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSSHKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoxServer).ListSSHKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sox.v1.Sox/ListSSHKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoxServer).ListSSHKeys(ctx, req.(*ListSSHKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sox_DeleteSSHKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSSHKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoxServer).DeleteSSHKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sox.v1.Sox/DeleteSSHKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoxServer).DeleteSSHKey(ctx, req.(*DeleteSSHKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sox_ListImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoxServer).ListImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sox.v1.Sox/ListImages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoxServer).ListImages(ctx, req.(*ListImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sox_ListNetworks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNetworksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoxServer).ListNetworks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sox.v1.Sox/ListNetworks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoxServer).ListNetworks(ctx, req.(*ListNetworksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sox_CreateNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoxServer).CreateNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sox.v1.Sox/CreateNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoxServer).CreateNetwork(ctx, req.(*CreateNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sox_ListActivities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListActivitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoxServer).ListActivities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sox.v1.Sox/ListActivities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoxServer).ListActivities(ctx, req.(*ListActivitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sox_ServiceDesc is the grpc.ServiceDesc for Sox service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sox_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sox.v1.Sox",
	HandlerType: (*SoxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMachine",
			Handler:    _Sox_CreateMachine_Handler,
		},
		{
			MethodName: "ListMachines",
			Handler:    _Sox_ListMachines_Handler,
		},
		{
			MethodName: "GetMachineDetails",
			Handler:    _Sox_GetMachineDetails_Handler,
		},
		{
			MethodName: "DeleteMachine",
			Handler:    _Sox_DeleteMachine_Handler,
		},
		{
			MethodName: "TriggerMachine",
			Handler:    _Sox_TriggerMachine_Handler,
		},
		{
			MethodName: "CreateSSHKey",
			Handler:    _Sox_CreateSSHKey_Handler,
		},
		{
			MethodName: "ListSSHKeys",
			Handler:    _Sox_ListSSHKeys_Handler,
		},
		{
			MethodName: "DeleteSSHKey",
			Handler:    _Sox_DeleteSSHKey_Handler,
		},
		{
			MethodName: "ListImages",
			Handler:    _Sox_ListImages_Handler,
		},
		{
			MethodName: "ListNetworks",
			Handler:    _Sox_ListNetworks_Handler,
		},
		{
			MethodName: "CreateNetwork",
			Handler:    _Sox_CreateNetwork_Handler,
		},
		{
			MethodName: "ListActivities",
			Handler:    _Sox_ListActivities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
