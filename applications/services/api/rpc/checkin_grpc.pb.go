// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rpc

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CheckInServiceClient is the client API for CheckInService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CheckInServiceClient interface {
	GetCheckInById(ctx context.Context, in *wrappers.Int64Value, opts ...grpc.CallOption) (*CheckIn, error)
	GetCheckInsByAttendenceCode(ctx context.Context, in *wrappers.Int64Value, opts ...grpc.CallOption) (*CheckIns, error)
	GetCheckInsByStudentId(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*CheckIns, error)
	GetCheckInsByTime(ctx context.Context, in *TimeInterval, opts ...grpc.CallOption) (*CheckIns, error)
	GetAllCheckIns(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CheckIns, error)
	InsertCheckIn(ctx context.Context, in *CheckIn, opts ...grpc.CallOption) (*empty.Empty, error)
}

type checkInServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckInServiceClient(cc grpc.ClientConnInterface) CheckInServiceClient {
	return &checkInServiceClient{cc}
}

func (c *checkInServiceClient) GetCheckInById(ctx context.Context, in *wrappers.Int64Value, opts ...grpc.CallOption) (*CheckIn, error) {
	out := new(CheckIn)
	err := c.cc.Invoke(ctx, "/rpc.CheckInService/GetCheckInById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkInServiceClient) GetCheckInsByAttendenceCode(ctx context.Context, in *wrappers.Int64Value, opts ...grpc.CallOption) (*CheckIns, error) {
	out := new(CheckIns)
	err := c.cc.Invoke(ctx, "/rpc.CheckInService/GetCheckInsByAttendenceCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkInServiceClient) GetCheckInsByStudentId(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*CheckIns, error) {
	out := new(CheckIns)
	err := c.cc.Invoke(ctx, "/rpc.CheckInService/GetCheckInsByStudentId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkInServiceClient) GetCheckInsByTime(ctx context.Context, in *TimeInterval, opts ...grpc.CallOption) (*CheckIns, error) {
	out := new(CheckIns)
	err := c.cc.Invoke(ctx, "/rpc.CheckInService/GetCheckInsByTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkInServiceClient) GetAllCheckIns(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CheckIns, error) {
	out := new(CheckIns)
	err := c.cc.Invoke(ctx, "/rpc.CheckInService/GetAllCheckIns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkInServiceClient) InsertCheckIn(ctx context.Context, in *CheckIn, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/rpc.CheckInService/InsertCheckIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckInServiceServer is the server API for CheckInService service.
// All implementations must embed UnimplementedCheckInServiceServer
// for forward compatibility
type CheckInServiceServer interface {
	GetCheckInById(context.Context, *wrappers.Int64Value) (*CheckIn, error)
	GetCheckInsByAttendenceCode(context.Context, *wrappers.Int64Value) (*CheckIns, error)
	GetCheckInsByStudentId(context.Context, *wrappers.StringValue) (*CheckIns, error)
	GetCheckInsByTime(context.Context, *TimeInterval) (*CheckIns, error)
	GetAllCheckIns(context.Context, *empty.Empty) (*CheckIns, error)
	InsertCheckIn(context.Context, *CheckIn) (*empty.Empty, error)
	mustEmbedUnimplementedCheckInServiceServer()
}

// UnimplementedCheckInServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCheckInServiceServer struct {
}

func (UnimplementedCheckInServiceServer) GetCheckInById(context.Context, *wrappers.Int64Value) (*CheckIn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCheckInById not implemented")
}
func (UnimplementedCheckInServiceServer) GetCheckInsByAttendenceCode(context.Context, *wrappers.Int64Value) (*CheckIns, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCheckInsByAttendenceCode not implemented")
}
func (UnimplementedCheckInServiceServer) GetCheckInsByStudentId(context.Context, *wrappers.StringValue) (*CheckIns, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCheckInsByStudentId not implemented")
}
func (UnimplementedCheckInServiceServer) GetCheckInsByTime(context.Context, *TimeInterval) (*CheckIns, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCheckInsByTime not implemented")
}
func (UnimplementedCheckInServiceServer) GetAllCheckIns(context.Context, *empty.Empty) (*CheckIns, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCheckIns not implemented")
}
func (UnimplementedCheckInServiceServer) InsertCheckIn(context.Context, *CheckIn) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertCheckIn not implemented")
}
func (UnimplementedCheckInServiceServer) mustEmbedUnimplementedCheckInServiceServer() {}

// UnsafeCheckInServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CheckInServiceServer will
// result in compilation errors.
type UnsafeCheckInServiceServer interface {
	mustEmbedUnimplementedCheckInServiceServer()
}

func RegisterCheckInServiceServer(s grpc.ServiceRegistrar, srv CheckInServiceServer) {
	s.RegisterService(&CheckInService_ServiceDesc, srv)
}

func _CheckInService_GetCheckInById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.Int64Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckInServiceServer).GetCheckInById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.CheckInService/GetCheckInById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckInServiceServer).GetCheckInById(ctx, req.(*wrappers.Int64Value))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheckInService_GetCheckInsByAttendenceCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.Int64Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckInServiceServer).GetCheckInsByAttendenceCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.CheckInService/GetCheckInsByAttendenceCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckInServiceServer).GetCheckInsByAttendenceCode(ctx, req.(*wrappers.Int64Value))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheckInService_GetCheckInsByStudentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckInServiceServer).GetCheckInsByStudentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.CheckInService/GetCheckInsByStudentId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckInServiceServer).GetCheckInsByStudentId(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheckInService_GetCheckInsByTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeInterval)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckInServiceServer).GetCheckInsByTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.CheckInService/GetCheckInsByTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckInServiceServer).GetCheckInsByTime(ctx, req.(*TimeInterval))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheckInService_GetAllCheckIns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckInServiceServer).GetAllCheckIns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.CheckInService/GetAllCheckIns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckInServiceServer).GetAllCheckIns(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheckInService_InsertCheckIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckInServiceServer).InsertCheckIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.CheckInService/InsertCheckIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckInServiceServer).InsertCheckIn(ctx, req.(*CheckIn))
	}
	return interceptor(ctx, in, info, handler)
}

// CheckInService_ServiceDesc is the grpc.ServiceDesc for CheckInService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CheckInService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.CheckInService",
	HandlerType: (*CheckInServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCheckInById",
			Handler:    _CheckInService_GetCheckInById_Handler,
		},
		{
			MethodName: "GetCheckInsByAttendenceCode",
			Handler:    _CheckInService_GetCheckInsByAttendenceCode_Handler,
		},
		{
			MethodName: "GetCheckInsByStudentId",
			Handler:    _CheckInService_GetCheckInsByStudentId_Handler,
		},
		{
			MethodName: "GetCheckInsByTime",
			Handler:    _CheckInService_GetCheckInsByTime_Handler,
		},
		{
			MethodName: "GetAllCheckIns",
			Handler:    _CheckInService_GetAllCheckIns_Handler,
		},
		{
			MethodName: "InsertCheckIn",
			Handler:    _CheckInService_InsertCheckIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "checkin.proto",
}
