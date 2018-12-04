// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tool_service.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ToolService service

type ToolServiceClient interface {
	// 创建或者修改数据库
	DbCreateOrModify(ctx context.Context, in *DbCreateOrModifyReq, opts ...grpc.CallOption) (*DbCreateOrModifyRsp, error)
	// 数据库详情
	DbDetailByID(ctx context.Context, in *DbDetailByIDReq, opts ...grpc.CallOption) (*DbDetailByIDRsp, error)
	// 数据库分页
	DbPage(ctx context.Context, in *DbPageReq, opts ...grpc.CallOption) (*DbPageRsp, error)
	// 所有数据库
	DbAllName(ctx context.Context, in *DbAllNameReq, opts ...grpc.CallOption) (*DbAllNameRsp, error)
	// 所有数据库
	DbConnect(ctx context.Context, in *DbConnectReq, opts ...grpc.CallOption) (*DbConnectRsp, error)
	// 表模型
	DbTableModel(ctx context.Context, in *DbTableModelReq, opts ...grpc.CallOption) (*DbTableModelRsp, error)
	// 删除
	DbDel(ctx context.Context, in *DbDelReq, opts ...grpc.CallOption) (*DbDelRsp, error)
}

type toolServiceClient struct {
	cc *grpc.ClientConn
}

func NewToolServiceClient(cc *grpc.ClientConn) ToolServiceClient {
	return &toolServiceClient{cc}
}

func (c *toolServiceClient) DbCreateOrModify(ctx context.Context, in *DbCreateOrModifyReq, opts ...grpc.CallOption) (*DbCreateOrModifyRsp, error) {
	out := new(DbCreateOrModifyRsp)
	err := grpc.Invoke(ctx, "/pb.ToolService/DbCreateOrModify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toolServiceClient) DbDetailByID(ctx context.Context, in *DbDetailByIDReq, opts ...grpc.CallOption) (*DbDetailByIDRsp, error) {
	out := new(DbDetailByIDRsp)
	err := grpc.Invoke(ctx, "/pb.ToolService/DbDetailByID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toolServiceClient) DbPage(ctx context.Context, in *DbPageReq, opts ...grpc.CallOption) (*DbPageRsp, error) {
	out := new(DbPageRsp)
	err := grpc.Invoke(ctx, "/pb.ToolService/DbPage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toolServiceClient) DbAllName(ctx context.Context, in *DbAllNameReq, opts ...grpc.CallOption) (*DbAllNameRsp, error) {
	out := new(DbAllNameRsp)
	err := grpc.Invoke(ctx, "/pb.ToolService/DbAllName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toolServiceClient) DbConnect(ctx context.Context, in *DbConnectReq, opts ...grpc.CallOption) (*DbConnectRsp, error) {
	out := new(DbConnectRsp)
	err := grpc.Invoke(ctx, "/pb.ToolService/DbConnect", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toolServiceClient) DbTableModel(ctx context.Context, in *DbTableModelReq, opts ...grpc.CallOption) (*DbTableModelRsp, error) {
	out := new(DbTableModelRsp)
	err := grpc.Invoke(ctx, "/pb.ToolService/DbTableModel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toolServiceClient) DbDel(ctx context.Context, in *DbDelReq, opts ...grpc.CallOption) (*DbDelRsp, error) {
	out := new(DbDelRsp)
	err := grpc.Invoke(ctx, "/pb.ToolService/DbDel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ToolService service

type ToolServiceServer interface {
	// 创建或者修改数据库
	DbCreateOrModify(context.Context, *DbCreateOrModifyReq) (*DbCreateOrModifyRsp, error)
	// 数据库详情
	DbDetailByID(context.Context, *DbDetailByIDReq) (*DbDetailByIDRsp, error)
	// 数据库分页
	DbPage(context.Context, *DbPageReq) (*DbPageRsp, error)
	// 所有数据库
	DbAllName(context.Context, *DbAllNameReq) (*DbAllNameRsp, error)
	// 所有数据库
	DbConnect(context.Context, *DbConnectReq) (*DbConnectRsp, error)
	// 表模型
	DbTableModel(context.Context, *DbTableModelReq) (*DbTableModelRsp, error)
	// 删除
	DbDel(context.Context, *DbDelReq) (*DbDelRsp, error)
}

func RegisterToolServiceServer(s *grpc.Server, srv ToolServiceServer) {
	s.RegisterService(&_ToolService_serviceDesc, srv)
}

func _ToolService_DbCreateOrModify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DbCreateOrModifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToolServiceServer).DbCreateOrModify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ToolService/DbCreateOrModify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToolServiceServer).DbCreateOrModify(ctx, req.(*DbCreateOrModifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToolService_DbDetailByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DbDetailByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToolServiceServer).DbDetailByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ToolService/DbDetailByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToolServiceServer).DbDetailByID(ctx, req.(*DbDetailByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToolService_DbPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DbPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToolServiceServer).DbPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ToolService/DbPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToolServiceServer).DbPage(ctx, req.(*DbPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToolService_DbAllName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DbAllNameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToolServiceServer).DbAllName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ToolService/DbAllName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToolServiceServer).DbAllName(ctx, req.(*DbAllNameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToolService_DbConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DbConnectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToolServiceServer).DbConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ToolService/DbConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToolServiceServer).DbConnect(ctx, req.(*DbConnectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToolService_DbTableModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DbTableModelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToolServiceServer).DbTableModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ToolService/DbTableModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToolServiceServer).DbTableModel(ctx, req.(*DbTableModelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToolService_DbDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DbDelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToolServiceServer).DbDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ToolService/DbDel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToolServiceServer).DbDel(ctx, req.(*DbDelReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ToolService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ToolService",
	HandlerType: (*ToolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DbCreateOrModify",
			Handler:    _ToolService_DbCreateOrModify_Handler,
		},
		{
			MethodName: "DbDetailByID",
			Handler:    _ToolService_DbDetailByID_Handler,
		},
		{
			MethodName: "DbPage",
			Handler:    _ToolService_DbPage_Handler,
		},
		{
			MethodName: "DbAllName",
			Handler:    _ToolService_DbAllName_Handler,
		},
		{
			MethodName: "DbConnect",
			Handler:    _ToolService_DbConnect_Handler,
		},
		{
			MethodName: "DbTableModel",
			Handler:    _ToolService_DbTableModel_Handler,
		},
		{
			MethodName: "DbDel",
			Handler:    _ToolService_DbDel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tool_service.proto",
}

func init() { proto.RegisterFile("tool_service.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4b, 0x43, 0x31,
	0x10, 0xc6, 0xa1, 0x60, 0xc1, 0xb3, 0x42, 0x39, 0x07, 0xa1, 0xab, 0xe0, 0xf8, 0x04, 0x5d, 0x5c,
	0xb5, 0x59, 0x1c, 0xaa, 0xa2, 0xdd, 0x25, 0xd7, 0x9e, 0xf2, 0xe0, 0xec, 0x9d, 0x49, 0x10, 0xfa,
	0x8f, 0x3b, 0x4b, 0x9a, 0x86, 0x86, 0xa7, 0x5b, 0xbe, 0xdf, 0x7d, 0xdf, 0xf0, 0x0b, 0x60, 0x52,
	0x95, 0xb7, 0xc8, 0xe1, 0xbb, 0x5f, 0x71, 0x67, 0x41, 0x93, 0xe2, 0xc8, 0x68, 0x06, 0x99, 0x97,
	0x7c, 0xfd, 0x33, 0x82, 0x93, 0xa5, 0xaa, 0xbc, 0x96, 0x16, 0x3a, 0x98, 0x3a, 0x9a, 0x07, 0xf6,
	0x89, 0x9f, 0xc2, 0x42, 0xd7, 0xfd, 0xfb, 0x16, 0xcf, 0x3b, 0xa3, 0x6e, 0x48, 0x5f, 0xf8, 0x6b,
	0xf6, 0xff, 0x21, 0x1a, 0xde, 0xc2, 0xc4, 0x91, 0xe3, 0xe4, 0x7b, 0xb9, 0xdf, 0x3e, 0x38, 0x3c,
	0x2b, 0xc5, 0x03, 0xc9, 0xeb, 0xbf, 0x30, 0x1a, 0x5e, 0xc2, 0xd8, 0xd1, 0xb3, 0xff, 0x60, 0x3c,
	0x2d, 0xe7, 0xfc, 0xce, 0xed, 0x36, 0x46, 0xc3, 0x2b, 0x38, 0x76, 0x74, 0x27, 0xf2, 0xe8, 0x3f,
	0x19, 0xa7, 0xe5, 0xb6, 0x8f, 0xb9, 0x3d, 0x20, 0x75, 0x30, 0xd7, 0xcd, 0x86, 0x57, 0xa9, 0x0e,
	0xf6, 0xb1, 0x19, 0x54, 0x52, 0x1d, 0x96, 0x9e, 0x84, 0x17, 0xba, 0x66, 0xa9, 0x0e, 0x07, 0xd2,
	0x38, 0x34, 0x30, 0x1a, 0x5e, 0xc0, 0x51, 0xd6, 0x12, 0x9c, 0x54, 0xc3, 0x5d, 0xb7, 0x49, 0xd1,
	0x68, 0xbc, 0xfb, 0xff, 0x9b, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x49, 0x92, 0x33, 0xa5,
	0x01, 0x00, 0x00,
}
