// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tool.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DbInfo struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name"`
	Ip       string `protobuf:"bytes,2,opt,name=ip" json:"ip"`
	Port     int32  `protobuf:"varint,3,opt,name=port" json:"port"`
	UserName string `protobuf:"bytes,4,opt,name=userName" json:"userName"`
	Password string `protobuf:"bytes,5,opt,name=password" json:"password"`
}

func (m *DbInfo) Reset()                    { *m = DbInfo{} }
func (m *DbInfo) String() string            { return proto.CompactTextString(m) }
func (*DbInfo) ProtoMessage()               {}
func (*DbInfo) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *DbInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DbInfo) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *DbInfo) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *DbInfo) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *DbInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// 创建数据库
type CreateReq struct {
	Db *DbInfo `protobuf:"bytes,1,opt,name=db" json:"db"`
}

func (m *CreateReq) Reset()                    { *m = CreateReq{} }
func (m *CreateReq) String() string            { return proto.CompactTextString(m) }
func (*CreateReq) ProtoMessage()               {}
func (*CreateReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *CreateReq) GetDb() *DbInfo {
	if m != nil {
		return m.Db
	}
	return nil
}

// 创建数据库
type CreateRsp struct {
}

func (m *CreateRsp) Reset()                    { *m = CreateRsp{} }
func (m *CreateRsp) String() string            { return proto.CompactTextString(m) }
func (*CreateRsp) ProtoMessage()               {}
func (*CreateRsp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func init() {
	proto.RegisterType((*DbInfo)(nil), "pb.DbInfo")
	proto.RegisterType((*CreateReq)(nil), "pb.CreateReq")
	proto.RegisterType((*CreateRsp)(nil), "pb.CreateRsp")
}

func init() { proto.RegisterFile("tool.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0xb1, 0xce, 0x82, 0x40,
	0x10, 0x84, 0xc3, 0xfd, 0x40, 0x7e, 0x96, 0xc4, 0xe2, 0xaa, 0x0b, 0x15, 0xa1, 0x91, 0x8a, 0x42,
	0x1f, 0x41, 0x1b, 0x1b, 0x8b, 0x7b, 0x03, 0x36, 0x9c, 0x09, 0x89, 0xb2, 0xeb, 0xdd, 0x19, 0x7d,
	0x7c, 0x73, 0x4b, 0xa4, 0x9b, 0xd9, 0xf9, 0x36, 0x33, 0x00, 0x91, 0xe8, 0x3e, 0xb0, 0xa7, 0x48,
	0x5a, 0x31, 0x76, 0x1f, 0x28, 0xcf, 0x78, 0x59, 0x6e, 0xa4, 0x35, 0xe4, 0xcb, 0xf8, 0x70, 0x26,
	0x6b, 0xb3, 0xbe, 0xb2, 0xa2, 0xf5, 0x0e, 0xd4, 0xcc, 0x46, 0xc9, 0x45, 0xcd, 0x9c, 0x18, 0x26,
	0x1f, 0xcd, 0x5f, 0x9b, 0xf5, 0x85, 0x15, 0xad, 0x1b, 0xf8, 0x7f, 0x05, 0xe7, 0xaf, 0xe9, 0x37,
	0x17, 0x72, 0xf3, 0x29, 0xe3, 0x31, 0x84, 0x37, 0xf9, 0xc9, 0x14, 0x6b, 0xf6, 0xf3, 0xdd, 0x1e,
	0xaa, 0x93, 0x77, 0x63, 0x74, 0xd6, 0x3d, 0x75, 0x03, 0x6a, 0x42, 0xa9, 0xae, 0x0f, 0x30, 0x30,
	0x0e, 0xeb, 0x28, 0xab, 0x26, 0xec, 0xea, 0x0d, 0x0c, 0x8c, 0xa5, 0x4c, 0x3f, 0x7e, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x81, 0xc0, 0xbc, 0x27, 0xc8, 0x00, 0x00, 0x00,
}
