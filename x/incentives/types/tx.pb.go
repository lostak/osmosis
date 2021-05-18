// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/incentives/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	types "github.com/osmosis-labs/osmosis/x/lockup/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgCreatePot struct {
	IsPerpetual bool `protobuf:"varint,1,opt,name=is_perpetual,json=isPerpetual,proto3" json:"is_perpetual,omitempty"`
	// distribution incentives by third party
	Owner        string                                   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	DistributeTo types.QueryCondition                     `protobuf:"bytes,3,opt,name=distribute_to,json=distributeTo,proto3" json:"distribute_to"`
	Coins        github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
	// distribution start time
	StartTime         time.Time `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"timestamp"`
	NumEpochsPaidOver uint64    `protobuf:"varint,6,opt,name=num_epochs_paid_over,json=numEpochsPaidOver,proto3" json:"num_epochs_paid_over,omitempty"`
}

func (m *MsgCreatePot) Reset()         { *m = MsgCreatePot{} }
func (m *MsgCreatePot) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePot) ProtoMessage()    {}
func (*MsgCreatePot) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ea120e22291556e, []int{0}
}
func (m *MsgCreatePot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePot.Merge(m, src)
}
func (m *MsgCreatePot) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePot) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePot.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePot proto.InternalMessageInfo

func (m *MsgCreatePot) GetIsPerpetual() bool {
	if m != nil {
		return m.IsPerpetual
	}
	return false
}

func (m *MsgCreatePot) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgCreatePot) GetDistributeTo() types.QueryCondition {
	if m != nil {
		return m.DistributeTo
	}
	return types.QueryCondition{}
}

func (m *MsgCreatePot) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

func (m *MsgCreatePot) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func (m *MsgCreatePot) GetNumEpochsPaidOver() uint64 {
	if m != nil {
		return m.NumEpochsPaidOver
	}
	return 0
}

type MsgCreatePotResponse struct {
}

func (m *MsgCreatePotResponse) Reset()         { *m = MsgCreatePotResponse{} }
func (m *MsgCreatePotResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePotResponse) ProtoMessage()    {}
func (*MsgCreatePotResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ea120e22291556e, []int{1}
}
func (m *MsgCreatePotResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePotResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePotResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePotResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePotResponse.Merge(m, src)
}
func (m *MsgCreatePotResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePotResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePotResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePotResponse proto.InternalMessageInfo

type MsgAddToPot struct {
	Owner   string                                   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	PotId   uint64                                   `protobuf:"varint,2,opt,name=pot_id,json=potId,proto3" json:"pot_id,omitempty"`
	Rewards github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=rewards,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"rewards"`
}

func (m *MsgAddToPot) Reset()         { *m = MsgAddToPot{} }
func (m *MsgAddToPot) String() string { return proto.CompactTextString(m) }
func (*MsgAddToPot) ProtoMessage()    {}
func (*MsgAddToPot) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ea120e22291556e, []int{2}
}
func (m *MsgAddToPot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddToPot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddToPot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddToPot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddToPot.Merge(m, src)
}
func (m *MsgAddToPot) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddToPot) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddToPot.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddToPot proto.InternalMessageInfo

func (m *MsgAddToPot) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgAddToPot) GetPotId() uint64 {
	if m != nil {
		return m.PotId
	}
	return 0
}

func (m *MsgAddToPot) GetRewards() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Rewards
	}
	return nil
}

type MsgAddToPotResponse struct {
}

func (m *MsgAddToPotResponse) Reset()         { *m = MsgAddToPotResponse{} }
func (m *MsgAddToPotResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAddToPotResponse) ProtoMessage()    {}
func (*MsgAddToPotResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ea120e22291556e, []int{3}
}
func (m *MsgAddToPotResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddToPotResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddToPotResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddToPotResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddToPotResponse.Merge(m, src)
}
func (m *MsgAddToPotResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddToPotResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddToPotResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddToPotResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreatePot)(nil), "osmosis.incentives.MsgCreatePot")
	proto.RegisterType((*MsgCreatePotResponse)(nil), "osmosis.incentives.MsgCreatePotResponse")
	proto.RegisterType((*MsgAddToPot)(nil), "osmosis.incentives.MsgAddToPot")
	proto.RegisterType((*MsgAddToPotResponse)(nil), "osmosis.incentives.MsgAddToPotResponse")
}

func init() { proto.RegisterFile("osmosis/incentives/tx.proto", fileDescriptor_8ea120e22291556e) }

var fileDescriptor_8ea120e22291556e = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x4e, 0xdb, 0x4c,
	0x14, 0xcd, 0x7c, 0x49, 0xf8, 0x60, 0x92, 0x4a, 0xd4, 0x85, 0xca, 0xa4, 0xc8, 0x4e, 0xbd, 0x68,
	0xbd, 0x61, 0xa6, 0xa4, 0xbb, 0xee, 0x1a, 0xd4, 0x05, 0x52, 0x51, 0xa9, 0x15, 0x09, 0xa9, 0x1b,
	0x6b, 0x6c, 0x4f, 0xcd, 0x88, 0xd8, 0x77, 0xe4, 0x19, 0x07, 0x78, 0x0b, 0xfa, 0x1a, 0xdd, 0xb6,
	0x0f, 0xc1, 0x92, 0x65, 0x57, 0x50, 0xc1, 0x1b, 0xf0, 0x04, 0x95, 0xff, 0x48, 0xaa, 0xfe, 0xb0,
	0xe9, 0x6a, 0x32, 0x73, 0xce, 0x3d, 0x73, 0xef, 0x39, 0x13, 0xe3, 0x27, 0xa0, 0x12, 0x50, 0x42,
	0x51, 0x91, 0x86, 0x3c, 0xd5, 0x62, 0xc6, 0x15, 0xd5, 0x27, 0x44, 0x66, 0xa0, 0xc1, 0x30, 0x6a,
	0x90, 0xcc, 0xc1, 0xc1, 0x5a, 0x0c, 0x31, 0x94, 0x30, 0x2d, 0x7e, 0x55, 0xcc, 0x81, 0x1d, 0x03,
	0xc4, 0x53, 0x4e, 0xcb, 0x5d, 0x90, 0x7f, 0xa4, 0x5a, 0x24, 0x5c, 0x69, 0x96, 0xc8, 0x9a, 0x60,
	0x85, 0xa5, 0x16, 0x0d, 0x98, 0xe2, 0x74, 0xb6, 0x1d, 0x70, 0xcd, 0xb6, 0x69, 0x08, 0x22, 0xad,
	0xf1, 0xcd, 0xdf, 0xf4, 0x21, 0x41, 0xd7, 0xe8, 0x46, 0x83, 0x4e, 0x21, 0x3c, 0xca, 0x65, 0xb9,
	0x54, 0x90, 0xf3, 0xa9, 0x8d, 0xfb, 0x7b, 0x2a, 0xde, 0xc9, 0x38, 0xd3, 0x7c, 0x1f, 0xb4, 0xf1,
	0x14, 0xf7, 0x85, 0xf2, 0x25, 0xcf, 0x24, 0xd7, 0x39, 0x9b, 0x9a, 0x68, 0x88, 0xdc, 0x65, 0xaf,
	0x27, 0xd4, 0x7e, 0x73, 0x64, 0x3c, 0xc3, 0x5d, 0x38, 0x4e, 0x79, 0x66, 0xfe, 0x37, 0x44, 0xee,
	0xca, 0x78, 0xf5, 0xf6, 0xd2, 0xee, 0x9f, 0xb2, 0x64, 0xfa, 0xca, 0x29, 0x8f, 0x1d, 0xaf, 0x82,
	0x8d, 0x5d, 0xfc, 0x20, 0x12, 0x4a, 0x67, 0x22, 0xc8, 0x35, 0xf7, 0x35, 0x98, 0xed, 0x21, 0x72,
	0x7b, 0x23, 0x8b, 0x34, 0xbe, 0x54, 0xed, 0x90, 0xf7, 0x39, 0xcf, 0x4e, 0x77, 0x20, 0x8d, 0x84,
	0x16, 0x90, 0x8e, 0x3b, 0xe7, 0x97, 0x76, 0xcb, 0xeb, 0xcf, 0x4b, 0x27, 0x60, 0x30, 0xdc, 0x2d,
	0xa6, 0x55, 0x66, 0x67, 0xd8, 0x76, 0x7b, 0xa3, 0x0d, 0x52, 0xf9, 0x41, 0x0a, 0x3f, 0x48, 0xed,
	0x07, 0xd9, 0x01, 0x91, 0x8e, 0x5f, 0x14, 0xd5, 0x9f, 0xaf, 0x6c, 0x37, 0x16, 0xfa, 0x30, 0x0f,
	0x48, 0x08, 0x09, 0xad, 0xcd, 0xab, 0x96, 0x2d, 0x15, 0x1d, 0x51, 0x7d, 0x2a, 0xb9, 0x2a, 0x0b,
	0x94, 0x57, 0x29, 0x1b, 0x07, 0x18, 0x2b, 0xcd, 0x32, 0xed, 0x17, 0xde, 0x9b, 0xdd, 0xb2, 0xd5,
	0x01, 0xa9, 0x82, 0x21, 0x4d, 0x30, 0x64, 0xd2, 0x04, 0x33, 0xde, 0x2c, 0x2e, 0xba, 0xbd, 0xb4,
	0x57, 0xab, 0xd1, 0xef, 0x12, 0x73, 0xce, 0xae, 0x6c, 0xe4, 0xad, 0x94, 0x5a, 0x05, 0xdb, 0xa0,
	0x78, 0x2d, 0xcd, 0x13, 0x9f, 0x4b, 0x08, 0x0f, 0x95, 0x2f, 0x99, 0x88, 0x7c, 0x98, 0xf1, 0xcc,
	0x5c, 0x1a, 0x22, 0xb7, 0xe3, 0x3d, 0x4c, 0xf3, 0xe4, 0x4d, 0x09, 0xed, 0x33, 0x11, 0xbd, 0x9b,
	0xf1, 0xcc, 0x79, 0x8c, 0xd7, 0x16, 0x23, 0xf1, 0xb8, 0x92, 0x90, 0x2a, 0xee, 0x7c, 0x41, 0xb8,
	0xb7, 0xa7, 0xe2, 0xd7, 0x51, 0x34, 0x81, 0x22, 0xaa, 0xbb, 0x1c, 0xd0, 0xdf, 0x73, 0x58, 0xc7,
	0x4b, 0x12, 0xb4, 0x2f, 0xa2, 0x32, 0xb0, 0x8e, 0xd7, 0x95, 0xa0, 0x77, 0x23, 0x83, 0xe3, 0xff,
	0x33, 0x7e, 0xcc, 0xb2, 0x48, 0x99, 0xed, 0x7f, 0xef, 0x6a, 0xa3, 0xed, 0xac, 0xe3, 0x47, 0x0b,
	0x4d, 0x37, 0xc3, 0x8c, 0xbe, 0x22, 0xdc, 0xde, 0x53, 0xb1, 0x71, 0x80, 0x57, 0xe6, 0x8f, 0x6f,
	0x48, 0x7e, 0xfd, 0xcb, 0x90, 0x45, 0x2f, 0x06, 0xee, 0x7d, 0x8c, 0xe6, 0x02, 0x63, 0x82, 0x97,
	0xef, 0x9c, 0xb2, 0xff, 0x50, 0xd5, 0x10, 0x06, 0xcf, 0xef, 0x21, 0x34, 0xaa, 0xe3, 0xb7, 0xe7,
	0xd7, 0x16, 0xba, 0xb8, 0xb6, 0xd0, 0xf7, 0x6b, 0x0b, 0x9d, 0xdd, 0x58, 0xad, 0x8b, 0x1b, 0xab,
	0xf5, 0xed, 0xc6, 0x6a, 0x7d, 0x18, 0x2d, 0x58, 0x53, 0x8b, 0x6d, 0x4d, 0x59, 0xa0, 0x9a, 0x0d,
	0x3d, 0xf9, 0xe9, 0x23, 0x51, 0x58, 0x15, 0x2c, 0x95, 0xef, 0xea, 0xe5, 0x8f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x1b, 0xd9, 0x18, 0x9c, 0x47, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreatePot(ctx context.Context, in *MsgCreatePot, opts ...grpc.CallOption) (*MsgCreatePotResponse, error)
	AddToPot(ctx context.Context, in *MsgAddToPot, opts ...grpc.CallOption) (*MsgAddToPotResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreatePot(ctx context.Context, in *MsgCreatePot, opts ...grpc.CallOption) (*MsgCreatePotResponse, error) {
	out := new(MsgCreatePotResponse)
	err := c.cc.Invoke(ctx, "/osmosis.incentives.Msg/CreatePot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddToPot(ctx context.Context, in *MsgAddToPot, opts ...grpc.CallOption) (*MsgAddToPotResponse, error) {
	out := new(MsgAddToPotResponse)
	err := c.cc.Invoke(ctx, "/osmosis.incentives.Msg/AddToPot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreatePot(context.Context, *MsgCreatePot) (*MsgCreatePotResponse, error)
	AddToPot(context.Context, *MsgAddToPot) (*MsgAddToPotResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreatePot(ctx context.Context, req *MsgCreatePot) (*MsgCreatePotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePot not implemented")
}
func (*UnimplementedMsgServer) AddToPot(ctx context.Context, req *MsgAddToPot) (*MsgAddToPotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToPot not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreatePot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.incentives.Msg/CreatePot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePot(ctx, req.(*MsgCreatePot))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddToPot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddToPot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddToPot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.incentives.Msg/AddToPot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddToPot(ctx, req.(*MsgAddToPot))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "osmosis.incentives.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePot",
			Handler:    _Msg_CreatePot_Handler,
		},
		{
			MethodName: "AddToPot",
			Handler:    _Msg_AddToPot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "osmosis/incentives/tx.proto",
}

func (m *MsgCreatePot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NumEpochsPaidOver != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.NumEpochsPaidOver))
		i--
		dAtA[i] = 0x30
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTx(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size, err := m.DistributeTo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.IsPerpetual {
		i--
		if m.IsPerpetual {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreatePotResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePotResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePotResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgAddToPot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddToPot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddToPot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for iNdEx := len(m.Rewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.PotId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.PotId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAddToPotResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddToPotResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddToPotResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreatePot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsPerpetual {
		n += 2
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.DistributeTo.Size()
	n += 1 + l + sovTx(uint64(l))
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovTx(uint64(l))
	if m.NumEpochsPaidOver != 0 {
		n += 1 + sovTx(uint64(m.NumEpochsPaidOver))
	}
	return n
}

func (m *MsgCreatePotResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgAddToPot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.PotId != 0 {
		n += 1 + sovTx(uint64(m.PotId))
	}
	if len(m.Rewards) > 0 {
		for _, e := range m.Rewards {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgAddToPotResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreatePot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreatePot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsPerpetual", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsPerpetual = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributeTo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DistributeTo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types1.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumEpochsPaidOver", wireType)
			}
			m.NumEpochsPaidOver = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumEpochsPaidOver |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreatePotResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreatePotResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePotResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgAddToPot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgAddToPot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddToPot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PotId", wireType)
			}
			m.PotId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PotId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rewards = append(m.Rewards, types1.Coin{})
			if err := m.Rewards[len(m.Rewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgAddToPotResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgAddToPotResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddToPotResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)