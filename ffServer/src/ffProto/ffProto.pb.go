// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ffProto.proto

/*
Package ffProto is a generated protocol buffer package.

It is generated from these files:
	ffProto.proto

It has these top-level messages:
	StAccountData
	MsgServerRegister
	MsgServerKeepAlive
	MsgPrepareLoginPlatformUniqueId
	MsgLoginPlatformUniqueId
	MsgLoginPlatformSidToken
	MsgReLogin
	MsgKick
	MsgEnterGameWorld
	MsgAgentDisConnect
	MsgKeepAlive
*/
package ffProto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MessageType int32

const (
	MessageType_ServerRegister               MessageType = 0
	MessageType_ServerKeepAlive              MessageType = 1
	MessageType_PrepareLoginPlatformUniqueId MessageType = 2
	MessageType_LoginPlatformUniqueId        MessageType = 3
	MessageType_LoginPlatformSidToken        MessageType = 4
	MessageType_ReLogin                      MessageType = 5
	MessageType_Kick                         MessageType = 6
	MessageType_EnterGameWorld               MessageType = 7
	MessageType_AgentDisConnect              MessageType = 8
	MessageType_KeepAlive                    MessageType = 9
)

var MessageType_name = map[int32]string{
	0: "ServerRegister",
	1: "ServerKeepAlive",
	2: "PrepareLoginPlatformUniqueId",
	3: "LoginPlatformUniqueId",
	4: "LoginPlatformSidToken",
	5: "ReLogin",
	6: "Kick",
	7: "EnterGameWorld",
	8: "AgentDisConnect",
	9: "KeepAlive",
}
var MessageType_value = map[string]int32{
	"ServerRegister":               0,
	"ServerKeepAlive":              1,
	"PrepareLoginPlatformUniqueId": 2,
	"LoginPlatformUniqueId":        3,
	"LoginPlatformSidToken":        4,
	"ReLogin":                      5,
	"Kick":                         6,
	"EnterGameWorld":               7,
	"AgentDisConnect":              8,
	"KeepAlive":                    9,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Generated by the msggen.py message compiler.
type StAccountData struct {
	IsFirst    bool    `protobuf:"varint,1,opt,name=isFirst" json:"isFirst,omitempty"`
	ServerTime int32   `protobuf:"varint,2,opt,name=serverTime" json:"serverTime,omitempty"`
	ServerZone int32   `protobuf:"varint,3,opt,name=serverZone" json:"serverZone,omitempty"`
	Name       string  `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	BaseKeys   []int32 `protobuf:"varint,5,rep,packed,name=baseKeys" json:"baseKeys,omitempty"`
	BaseDatas  []int32 `protobuf:"varint,6,rep,packed,name=baseDatas" json:"baseDatas,omitempty"`
}

func (m *StAccountData) Reset()                    { *m = StAccountData{} }
func (m *StAccountData) String() string            { return proto.CompactTextString(m) }
func (*StAccountData) ProtoMessage()               {}
func (*StAccountData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StAccountData) GetIsFirst() bool {
	if m != nil {
		return m.IsFirst
	}
	return false
}

func (m *StAccountData) GetServerTime() int32 {
	if m != nil {
		return m.ServerTime
	}
	return 0
}

func (m *StAccountData) GetServerZone() int32 {
	if m != nil {
		return m.ServerZone
	}
	return 0
}

func (m *StAccountData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StAccountData) GetBaseKeys() []int32 {
	if m != nil {
		return m.BaseKeys
	}
	return nil
}

func (m *StAccountData) GetBaseDatas() []int32 {
	if m != nil {
		return m.BaseDatas
	}
	return nil
}

type MsgServerRegister struct {
	ServerType string `protobuf:"bytes,1,opt,name=serverType" json:"serverType,omitempty"`
	ServerID   int32  `protobuf:"varint,2,opt,name=serverID" json:"serverID,omitempty"`
}

func (m *MsgServerRegister) Reset()                    { *m = MsgServerRegister{} }
func (m *MsgServerRegister) String() string            { return proto.CompactTextString(m) }
func (*MsgServerRegister) ProtoMessage()               {}
func (*MsgServerRegister) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MsgServerRegister) GetServerType() string {
	if m != nil {
		return m.ServerType
	}
	return ""
}

func (m *MsgServerRegister) GetServerID() int32 {
	if m != nil {
		return m.ServerID
	}
	return 0
}

type MsgServerKeepAlive struct {
}

func (m *MsgServerKeepAlive) Reset()                    { *m = MsgServerKeepAlive{} }
func (m *MsgServerKeepAlive) String() string            { return proto.CompactTextString(m) }
func (*MsgServerKeepAlive) ProtoMessage()               {}
func (*MsgServerKeepAlive) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type MsgPrepareLoginPlatformUniqueId struct {
	SubChannel        string `protobuf:"bytes,1,opt,name=subChannel" json:"subChannel,omitempty"`
	UUIDPlatformBound string `protobuf:"bytes,2,opt,name=UUIDPlatformBound" json:"UUIDPlatformBound,omitempty"`
	UUIDPlatformLogin string `protobuf:"bytes,3,opt,name=UUIDPlatformLogin" json:"UUIDPlatformLogin,omitempty"`
	RandomSalt        string `protobuf:"bytes,4,opt,name=randomSalt" json:"randomSalt,omitempty"`
	Timestamp         int32  `protobuf:"varint,5,opt,name=timestamp" json:"timestamp,omitempty"`
	Status            int32  `protobuf:"varint,6,opt,name=status" json:"status,omitempty"`
	Result            int32  `protobuf:"varint,7,opt,name=result" json:"result,omitempty"`
}

func (m *MsgPrepareLoginPlatformUniqueId) Reset()                    { *m = MsgPrepareLoginPlatformUniqueId{} }
func (m *MsgPrepareLoginPlatformUniqueId) String() string            { return proto.CompactTextString(m) }
func (*MsgPrepareLoginPlatformUniqueId) ProtoMessage()               {}
func (*MsgPrepareLoginPlatformUniqueId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MsgPrepareLoginPlatformUniqueId) GetSubChannel() string {
	if m != nil {
		return m.SubChannel
	}
	return ""
}

func (m *MsgPrepareLoginPlatformUniqueId) GetUUIDPlatformBound() string {
	if m != nil {
		return m.UUIDPlatformBound
	}
	return ""
}

func (m *MsgPrepareLoginPlatformUniqueId) GetUUIDPlatformLogin() string {
	if m != nil {
		return m.UUIDPlatformLogin
	}
	return ""
}

func (m *MsgPrepareLoginPlatformUniqueId) GetRandomSalt() string {
	if m != nil {
		return m.RandomSalt
	}
	return ""
}

func (m *MsgPrepareLoginPlatformUniqueId) GetTimestamp() int32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *MsgPrepareLoginPlatformUniqueId) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *MsgPrepareLoginPlatformUniqueId) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type MsgLoginPlatformUniqueId struct {
	TokenCustom string `protobuf:"bytes,1,opt,name=tokenCustom" json:"tokenCustom,omitempty"`
	DeviceGUID  string `protobuf:"bytes,2,opt,name=deviceGUID" json:"deviceGUID,omitempty"`
	UUIDLogin   uint64 `protobuf:"varint,3,opt,name=UUIDLogin" json:"UUIDLogin,omitempty"`
	Result      int32  `protobuf:"varint,4,opt,name=result" json:"result,omitempty"`
}

func (m *MsgLoginPlatformUniqueId) Reset()                    { *m = MsgLoginPlatformUniqueId{} }
func (m *MsgLoginPlatformUniqueId) String() string            { return proto.CompactTextString(m) }
func (*MsgLoginPlatformUniqueId) ProtoMessage()               {}
func (*MsgLoginPlatformUniqueId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MsgLoginPlatformUniqueId) GetTokenCustom() string {
	if m != nil {
		return m.TokenCustom
	}
	return ""
}

func (m *MsgLoginPlatformUniqueId) GetDeviceGUID() string {
	if m != nil {
		return m.DeviceGUID
	}
	return ""
}

func (m *MsgLoginPlatformUniqueId) GetUUIDLogin() uint64 {
	if m != nil {
		return m.UUIDLogin
	}
	return 0
}

func (m *MsgLoginPlatformUniqueId) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type MsgLoginPlatformSidToken struct {
	TokenPlatform string `protobuf:"bytes,1,opt,name=tokenPlatform" json:"tokenPlatform,omitempty"`
	DeviceGUID    string `protobuf:"bytes,2,opt,name=deviceGUID" json:"deviceGUID,omitempty"`
	UUIDLogin     uint64 `protobuf:"varint,3,opt,name=UUIDLogin" json:"UUIDLogin,omitempty"`
	Result        int32  `protobuf:"varint,4,opt,name=result" json:"result,omitempty"`
}

func (m *MsgLoginPlatformSidToken) Reset()                    { *m = MsgLoginPlatformSidToken{} }
func (m *MsgLoginPlatformSidToken) String() string            { return proto.CompactTextString(m) }
func (*MsgLoginPlatformSidToken) ProtoMessage()               {}
func (*MsgLoginPlatformSidToken) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MsgLoginPlatformSidToken) GetTokenPlatform() string {
	if m != nil {
		return m.TokenPlatform
	}
	return ""
}

func (m *MsgLoginPlatformSidToken) GetDeviceGUID() string {
	if m != nil {
		return m.DeviceGUID
	}
	return ""
}

func (m *MsgLoginPlatformSidToken) GetUUIDLogin() uint64 {
	if m != nil {
		return m.UUIDLogin
	}
	return 0
}

func (m *MsgLoginPlatformSidToken) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type MsgReLogin struct {
	CheckData  string `protobuf:"bytes,1,opt,name=checkData" json:"checkData,omitempty"`
	DeviceGUID string `protobuf:"bytes,2,opt,name=deviceGUID" json:"deviceGUID,omitempty"`
	Result     int32  `protobuf:"varint,3,opt,name=result" json:"result,omitempty"`
}

func (m *MsgReLogin) Reset()                    { *m = MsgReLogin{} }
func (m *MsgReLogin) String() string            { return proto.CompactTextString(m) }
func (*MsgReLogin) ProtoMessage()               {}
func (*MsgReLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *MsgReLogin) GetCheckData() string {
	if m != nil {
		return m.CheckData
	}
	return ""
}

func (m *MsgReLogin) GetDeviceGUID() string {
	if m != nil {
		return m.DeviceGUID
	}
	return ""
}

func (m *MsgReLogin) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type MsgKick struct {
	Result int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *MsgKick) Reset()                    { *m = MsgKick{} }
func (m *MsgKick) String() string            { return proto.CompactTextString(m) }
func (*MsgKick) ProtoMessage()               {}
func (*MsgKick) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *MsgKick) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type MsgEnterGameWorld struct {
	ServerID  int32  `protobuf:"varint,1,opt,name=serverID" json:"serverID,omitempty"`
	UUIDLogin uint64 `protobuf:"varint,2,opt,name=UUIDLogin" json:"UUIDLogin,omitempty"`
	Result    int32  `protobuf:"varint,3,opt,name=result" json:"result,omitempty"`
}

func (m *MsgEnterGameWorld) Reset()                    { *m = MsgEnterGameWorld{} }
func (m *MsgEnterGameWorld) String() string            { return proto.CompactTextString(m) }
func (*MsgEnterGameWorld) ProtoMessage()               {}
func (*MsgEnterGameWorld) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *MsgEnterGameWorld) GetServerID() int32 {
	if m != nil {
		return m.ServerID
	}
	return 0
}

func (m *MsgEnterGameWorld) GetUUIDLogin() uint64 {
	if m != nil {
		return m.UUIDLogin
	}
	return 0
}

func (m *MsgEnterGameWorld) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type MsgAgentDisConnect struct {
}

func (m *MsgAgentDisConnect) Reset()                    { *m = MsgAgentDisConnect{} }
func (m *MsgAgentDisConnect) String() string            { return proto.CompactTextString(m) }
func (*MsgAgentDisConnect) ProtoMessage()               {}
func (*MsgAgentDisConnect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type MsgKeepAlive struct {
	Number int32 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
}

func (m *MsgKeepAlive) Reset()                    { *m = MsgKeepAlive{} }
func (m *MsgKeepAlive) String() string            { return proto.CompactTextString(m) }
func (*MsgKeepAlive) ProtoMessage()               {}
func (*MsgKeepAlive) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *MsgKeepAlive) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func init() {
	proto.RegisterType((*StAccountData)(nil), "StAccountData")
	proto.RegisterType((*MsgServerRegister)(nil), "MsgServerRegister")
	proto.RegisterType((*MsgServerKeepAlive)(nil), "MsgServerKeepAlive")
	proto.RegisterType((*MsgPrepareLoginPlatformUniqueId)(nil), "MsgPrepareLoginPlatformUniqueId")
	proto.RegisterType((*MsgLoginPlatformUniqueId)(nil), "MsgLoginPlatformUniqueId")
	proto.RegisterType((*MsgLoginPlatformSidToken)(nil), "MsgLoginPlatformSidToken")
	proto.RegisterType((*MsgReLogin)(nil), "MsgReLogin")
	proto.RegisterType((*MsgKick)(nil), "MsgKick")
	proto.RegisterType((*MsgEnterGameWorld)(nil), "MsgEnterGameWorld")
	proto.RegisterType((*MsgAgentDisConnect)(nil), "MsgAgentDisConnect")
	proto.RegisterType((*MsgKeepAlive)(nil), "MsgKeepAlive")
	proto.RegisterEnum("MessageType", MessageType_name, MessageType_value)
}

func init() { proto.RegisterFile("ffProto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x86, 0x3f, 0xe7, 0xdf, 0xd3, 0x2f, 0xe0, 0x2e, 0x3f, 0x32, 0xa8, 0x82, 0x10, 0x21, 0x14,
	0x21, 0xc4, 0x09, 0x57, 0x50, 0x1a, 0xa8, 0xaa, 0x62, 0x51, 0x39, 0x8d, 0x90, 0x38, 0xdb, 0xd8,
	0x53, 0x77, 0x55, 0x7b, 0x37, 0xec, 0xae, 0x2b, 0xf5, 0x1a, 0x38, 0xe3, 0x80, 0x5b, 0xe1, 0x6a,
	0xb8, 0x17, 0xb4, 0x6b, 0x3b, 0xb6, 0xd3, 0x00, 0x47, 0x9c, 0x44, 0x9e, 0x77, 0xc6, 0x3b, 0xcf,
	0x8c, 0xf7, 0x0d, 0x8c, 0x2f, 0x2e, 0xce, 0xa4, 0xd0, 0xe2, 0xf5, 0xda, 0xfc, 0x4e, 0x7f, 0x38,
	0x30, 0x5e, 0xe8, 0xc3, 0x28, 0x12, 0x39, 0xd7, 0x73, 0xaa, 0x29, 0xf1, 0x61, 0xc8, 0xd4, 0x7b,
	0x26, 0x95, 0xf6, 0x9d, 0x89, 0x33, 0x1b, 0x85, 0x55, 0x48, 0x9e, 0x00, 0x28, 0x94, 0xd7, 0x28,
	0xcf, 0x59, 0x86, 0x7e, 0x67, 0xe2, 0xcc, 0xfa, 0x61, 0x43, 0xa9, 0xf3, 0x9f, 0x05, 0x47, 0xbf,
	0xdb, 0xcc, 0x1b, 0x85, 0x10, 0xe8, 0x71, 0x9a, 0xa1, 0xdf, 0x9b, 0x38, 0x33, 0x37, 0xb4, 0xcf,
	0xe4, 0x31, 0x8c, 0x56, 0x54, 0xe1, 0x29, 0xde, 0x28, 0xbf, 0x3f, 0xe9, 0xce, 0xfa, 0xe1, 0x26,
	0x26, 0x07, 0xe0, 0x9a, 0x67, 0x43, 0xa5, 0xfc, 0x81, 0x4d, 0xd6, 0xc2, 0xf4, 0x23, 0xec, 0x07,
	0x2a, 0x59, 0xd8, 0xe3, 0x43, 0x4c, 0x98, 0xd2, 0x28, 0x1b, 0x88, 0x37, 0x6b, 0xb4, 0xfc, 0x6e,
	0xd8, 0x50, 0x4c, 0xbb, 0x22, 0x3a, 0x99, 0x97, 0x03, 0x6c, 0xe2, 0xe9, 0x7d, 0x20, 0x9b, 0x03,
	0x4f, 0x11, 0xd7, 0x87, 0x29, 0xbb, 0xc6, 0xe9, 0xd7, 0x0e, 0x3c, 0x0d, 0x54, 0x72, 0x26, 0x71,
	0x4d, 0x25, 0x7e, 0x10, 0x09, 0xe3, 0x67, 0x29, 0xd5, 0x17, 0x42, 0x66, 0x4b, 0xce, 0xbe, 0xe4,
	0x78, 0x12, 0xdb, 0xae, 0xf9, 0xea, 0xe8, 0x92, 0x72, 0x8e, 0xe9, 0xa6, 0xeb, 0x46, 0x21, 0xaf,
	0x60, 0x7f, 0xb9, 0x3c, 0x99, 0x57, 0xef, 0xbd, 0x15, 0x39, 0x8f, 0x6d, 0x7b, 0x37, 0xbc, 0x9d,
	0xd8, 0xae, 0xb6, 0x2d, 0xed, 0x36, 0xb7, 0xaa, 0x6d, 0xc2, 0xf4, 0x96, 0x94, 0xc7, 0x22, 0x5b,
	0xd0, 0x54, 0x97, 0xab, 0x6d, 0x28, 0x66, 0x89, 0x9a, 0x65, 0xa8, 0x34, 0xcd, 0xd6, 0x7e, 0xdf,
	0x8e, 0x5c, 0x0b, 0xe4, 0x21, 0x0c, 0x94, 0xa6, 0x3a, 0x37, 0xfb, 0x35, 0xa9, 0x32, 0x32, 0xba,
	0x44, 0x95, 0xa7, 0xda, 0x1f, 0x16, 0x7a, 0x11, 0x4d, 0xbf, 0x39, 0xe0, 0x07, 0x2a, 0xd9, 0xbd,
	0x86, 0x09, 0xec, 0x69, 0x71, 0x85, 0xfc, 0x28, 0x57, 0x5a, 0x64, 0xe5, 0x1e, 0x9a, 0x92, 0x81,
	0x8d, 0xf1, 0x9a, 0x45, 0x78, 0xbc, 0x2c, 0x3f, 0x80, 0x1b, 0x36, 0x14, 0x03, 0x6b, 0x26, 0xac,
	0x47, 0xee, 0x85, 0xb5, 0xd0, 0x80, 0xea, 0xb5, 0xa0, 0xbe, 0xef, 0x80, 0x5a, 0xb0, 0xf8, 0xdc,
	0x34, 0x26, 0xcf, 0x61, 0x6c, 0x09, 0xaa, 0x44, 0x89, 0xd5, 0x16, 0xff, 0x11, 0xd8, 0x0a, 0x20,
	0x50, 0x49, 0x58, 0xdc, 0x1a, 0x73, 0x46, 0x74, 0x89, 0xd1, 0x95, 0xb9, 0xbe, 0x25, 0x45, 0x2d,
	0xfc, 0x95, 0xa0, 0xee, 0xd1, 0x6d, 0xf5, 0x78, 0x06, 0xc3, 0x40, 0x25, 0xa7, 0x2c, 0xba, 0x6a,
	0x94, 0x38, 0xad, 0x12, 0xb4, 0x4e, 0x79, 0xc7, 0x35, 0xca, 0x63, 0x9a, 0xe1, 0x27, 0x21, 0xd3,
	0xb8, 0xe5, 0x04, 0xa7, 0xed, 0x84, 0xf6, 0xb4, 0x9d, 0xdf, 0x4f, 0xdb, 0x26, 0x29, 0xfc, 0x73,
	0x98, 0x20, 0xd7, 0x73, 0xa6, 0x8e, 0x04, 0xe7, 0x18, 0xe9, 0xe9, 0x0b, 0xf8, 0xdf, 0xf0, 0x55,
	0x7e, 0x32, 0x6f, 0xf3, 0x3c, 0x5b, 0xa1, 0xac, 0x20, 0x8b, 0xe8, 0xe5, 0x4f, 0x07, 0xf6, 0x02,
	0x54, 0x8a, 0x26, 0x68, 0x9d, 0x4a, 0xe0, 0x4e, 0xdb, 0xdb, 0xde, 0x7f, 0xe4, 0x1e, 0xdc, 0xdd,
	0xb2, 0xa7, 0xe7, 0x90, 0x09, 0x1c, 0xfc, 0xc9, 0x9c, 0x5e, 0x87, 0x3c, 0x82, 0x07, 0xbb, 0x53,
	0xdd, 0x5b, 0xa9, 0xea, 0xda, 0x78, 0x3d, 0xb2, 0x07, 0xc3, 0xf2, 0xcb, 0x79, 0x7d, 0x32, 0x82,
	0x9e, 0x59, 0xb1, 0x37, 0x30, 0x5c, 0xed, 0x4d, 0x7a, 0x43, 0xc3, 0xb5, 0x35, 0xb6, 0x37, 0x22,
	0x63, 0x70, 0x6b, 0x4c, 0x77, 0x35, 0xb0, 0xff, 0xb7, 0x6f, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x08, 0x97, 0xe9, 0xb2, 0x80, 0x05, 0x00, 0x00,
}
