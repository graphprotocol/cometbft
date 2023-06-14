// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cometbft/consensus/v2/wal.proto

package v2

import (
	fmt "fmt"
	v11 "github.com/cometbft/cometbft/api/cometbft/consensus/v1"
	v1 "github.com/cometbft/cometbft/api/cometbft/types/v1"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/cosmos/gogoproto/types"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
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

// MsgInfo are msgs from the reactor which may update the state
type MsgInfo struct {
	Msg    Message `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg"`
	PeerID string  `protobuf:"bytes,2,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
}

func (m *MsgInfo) Reset()         { *m = MsgInfo{} }
func (m *MsgInfo) String() string { return proto.CompactTextString(m) }
func (*MsgInfo) ProtoMessage()    {}
func (*MsgInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_604b77aec4785fe5, []int{0}
}
func (m *MsgInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInfo.Merge(m, src)
}
func (m *MsgInfo) XXX_Size() int {
	return m.Size()
}
func (m *MsgInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInfo proto.InternalMessageInfo

func (m *MsgInfo) GetMsg() Message {
	if m != nil {
		return m.Msg
	}
	return Message{}
}

func (m *MsgInfo) GetPeerID() string {
	if m != nil {
		return m.PeerID
	}
	return ""
}

type WALMessage struct {
	// Types that are valid to be assigned to Sum:
	//	*WALMessage_EventDataRoundState
	//	*WALMessage_MsgInfo
	//	*WALMessage_TimeoutInfo
	//	*WALMessage_EndHeight
	Sum isWALMessage_Sum `protobuf_oneof:"sum"`
}

func (m *WALMessage) Reset()         { *m = WALMessage{} }
func (m *WALMessage) String() string { return proto.CompactTextString(m) }
func (*WALMessage) ProtoMessage()    {}
func (*WALMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_604b77aec4785fe5, []int{1}
}
func (m *WALMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WALMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WALMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WALMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WALMessage.Merge(m, src)
}
func (m *WALMessage) XXX_Size() int {
	return m.Size()
}
func (m *WALMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_WALMessage.DiscardUnknown(m)
}

var xxx_messageInfo_WALMessage proto.InternalMessageInfo

type isWALMessage_Sum interface {
	isWALMessage_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type WALMessage_EventDataRoundState struct {
	EventDataRoundState *v1.EventDataRoundState `protobuf:"bytes,1,opt,name=event_data_round_state,json=eventDataRoundState,proto3,oneof" json:"event_data_round_state,omitempty"`
}
type WALMessage_MsgInfo struct {
	MsgInfo *MsgInfo `protobuf:"bytes,2,opt,name=msg_info,json=msgInfo,proto3,oneof" json:"msg_info,omitempty"`
}
type WALMessage_TimeoutInfo struct {
	TimeoutInfo *v11.TimeoutInfo `protobuf:"bytes,3,opt,name=timeout_info,json=timeoutInfo,proto3,oneof" json:"timeout_info,omitempty"`
}
type WALMessage_EndHeight struct {
	EndHeight *v11.EndHeight `protobuf:"bytes,4,opt,name=end_height,json=endHeight,proto3,oneof" json:"end_height,omitempty"`
}

func (*WALMessage_EventDataRoundState) isWALMessage_Sum() {}
func (*WALMessage_MsgInfo) isWALMessage_Sum()             {}
func (*WALMessage_TimeoutInfo) isWALMessage_Sum()         {}
func (*WALMessage_EndHeight) isWALMessage_Sum()           {}

func (m *WALMessage) GetSum() isWALMessage_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *WALMessage) GetEventDataRoundState() *v1.EventDataRoundState {
	if x, ok := m.GetSum().(*WALMessage_EventDataRoundState); ok {
		return x.EventDataRoundState
	}
	return nil
}

func (m *WALMessage) GetMsgInfo() *MsgInfo {
	if x, ok := m.GetSum().(*WALMessage_MsgInfo); ok {
		return x.MsgInfo
	}
	return nil
}

func (m *WALMessage) GetTimeoutInfo() *v11.TimeoutInfo {
	if x, ok := m.GetSum().(*WALMessage_TimeoutInfo); ok {
		return x.TimeoutInfo
	}
	return nil
}

func (m *WALMessage) GetEndHeight() *v11.EndHeight {
	if x, ok := m.GetSum().(*WALMessage_EndHeight); ok {
		return x.EndHeight
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*WALMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*WALMessage_EventDataRoundState)(nil),
		(*WALMessage_MsgInfo)(nil),
		(*WALMessage_TimeoutInfo)(nil),
		(*WALMessage_EndHeight)(nil),
	}
}

// TimedWALMessage wraps WALMessage and adds Time for debugging purposes.
type TimedWALMessage struct {
	Time time.Time   `protobuf:"bytes,1,opt,name=time,proto3,stdtime" json:"time"`
	Msg  *WALMessage `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *TimedWALMessage) Reset()         { *m = TimedWALMessage{} }
func (m *TimedWALMessage) String() string { return proto.CompactTextString(m) }
func (*TimedWALMessage) ProtoMessage()    {}
func (*TimedWALMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_604b77aec4785fe5, []int{2}
}
func (m *TimedWALMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TimedWALMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TimedWALMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TimedWALMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimedWALMessage.Merge(m, src)
}
func (m *TimedWALMessage) XXX_Size() int {
	return m.Size()
}
func (m *TimedWALMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TimedWALMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TimedWALMessage proto.InternalMessageInfo

func (m *TimedWALMessage) GetTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func (m *TimedWALMessage) GetMsg() *WALMessage {
	if m != nil {
		return m.Msg
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgInfo)(nil), "cometbft.consensus.v2.MsgInfo")
	proto.RegisterType((*WALMessage)(nil), "cometbft.consensus.v2.WALMessage")
	proto.RegisterType((*TimedWALMessage)(nil), "cometbft.consensus.v2.TimedWALMessage")
}

func init() { proto.RegisterFile("cometbft/consensus/v2/wal.proto", fileDescriptor_604b77aec4785fe5) }

var fileDescriptor_604b77aec4785fe5 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x5d, 0x6b, 0xdb, 0x30,
	0x14, 0xb5, 0x9b, 0x2c, 0x69, 0xd5, 0xc1, 0xc0, 0xfb, 0x20, 0xe4, 0xc1, 0x4e, 0x33, 0x18, 0x7d,
	0x92, 0x48, 0x0a, 0x65, 0xb0, 0xa7, 0x9a, 0x96, 0x25, 0xb0, 0x42, 0xf1, 0x06, 0x83, 0xc1, 0x30,
	0x4a, 0x7d, 0xad, 0x18, 0x6a, 0xc9, 0x44, 0x72, 0xc6, 0xde, 0xf6, 0x13, 0xfa, 0x7f, 0xf6, 0x07,
	0xfa, 0xd8, 0xc7, 0x3d, 0x75, 0x23, 0xf9, 0x23, 0x43, 0x1f, 0x69, 0x0a, 0x73, 0xf7, 0x26, 0x5d,
	0x9d, 0x73, 0xee, 0xbd, 0xe7, 0x08, 0x45, 0x97, 0xa2, 0x04, 0x35, 0xcb, 0x15, 0xb9, 0x14, 0x5c,
	0x02, 0x97, 0xb5, 0x24, 0xcb, 0x31, 0xf9, 0x46, 0xaf, 0x70, 0xb5, 0x10, 0x4a, 0x04, 0x2f, 0x37,
	0x00, 0x7c, 0x0f, 0xc0, 0xcb, 0x71, 0xbf, 0x91, 0x37, 0xda, 0xf2, 0xfa, 0x07, 0xcd, 0xc2, 0xea,
	0x7b, 0x05, 0xd2, 0x41, 0xc2, 0x7b, 0x88, 0xa9, 0x6a, 0x3e, 0x2c, 0x81, 0xab, 0xcd, 0xfb, 0x0b,
	0x26, 0x98, 0x30, 0x47, 0xa2, 0x4f, 0xae, 0x1a, 0x31, 0x21, 0xd8, 0x15, 0x10, 0x73, 0x9b, 0xd5,
	0x39, 0x51, 0x45, 0x09, 0x52, 0xd1, 0xb2, 0xb2, 0x80, 0x61, 0x8e, 0xba, 0xe7, 0x92, 0x4d, 0x79,
	0x2e, 0x82, 0x63, 0xd4, 0x2a, 0x25, 0xeb, 0xf9, 0x03, 0xff, 0x70, 0x7f, 0x1c, 0xe2, 0xc6, 0x55,
	0xf0, 0x39, 0x48, 0x49, 0x19, 0xc4, 0xed, 0x9b, 0xbb, 0xc8, 0x4b, 0x34, 0x21, 0x78, 0x8d, 0xba,
	0x15, 0xc0, 0x22, 0x2d, 0xb2, 0xde, 0xce, 0xc0, 0x3f, 0xdc, 0x8b, 0xd1, 0xea, 0x2e, 0xea, 0x5c,
	0x00, 0x2c, 0xa6, 0xa7, 0x49, 0x47, 0x3f, 0x4d, 0xb3, 0xe1, 0xcf, 0x1d, 0x84, 0x3e, 0x9f, 0x7c,
	0x70, 0xf4, 0xe0, 0x2b, 0x7a, 0x65, 0xa6, 0x4f, 0x33, 0xaa, 0x68, 0xba, 0x10, 0x35, 0xcf, 0x52,
	0xa9, 0xa8, 0x02, 0xd7, 0xfe, 0xcd, 0xb6, 0xbd, 0x35, 0x61, 0x39, 0xc2, 0x67, 0x9a, 0x70, 0x4a,
	0x15, 0x4d, 0x34, 0xfc, 0xa3, 0x46, 0x4f, 0xbc, 0xe4, 0x39, 0xfc, 0x5b, 0x0e, 0xde, 0xa1, 0xdd,
	0x52, 0xb2, 0xb4, 0xe0, 0xb9, 0x30, 0x33, 0xfd, 0x67, 0x1f, 0xbb, 0xfc, 0xc4, 0x4b, 0xba, 0xa5,
	0xf3, 0xe1, 0x3d, 0x7a, 0xaa, 0x5d, 0x12, 0xb5, 0xb2, 0x02, 0x2d, 0x23, 0x30, 0x6c, 0x14, 0x18,
	0xe1, 0x4f, 0x16, 0xea, 0x44, 0xf6, 0xd5, 0xf6, 0x1a, 0x9c, 0x20, 0x04, 0x3c, 0x4b, 0xe7, 0x50,
	0xb0, 0xb9, 0xea, 0xb5, 0x8d, 0xcc, 0xe0, 0x11, 0x99, 0x33, 0x9e, 0x4d, 0x0c, 0x6e, 0xe2, 0x25,
	0x7b, 0xb0, 0xb9, 0xc4, 0x4f, 0x50, 0x4b, 0xd6, 0xe5, 0xf0, 0x87, 0x8f, 0x9e, 0xe9, 0x46, 0xd9,
	0x03, 0x0b, 0xdf, 0xa2, 0xb6, 0x6e, 0xe6, 0x0c, 0xeb, 0x63, 0x9b, 0x34, 0xde, 0x24, 0x6d, 0x06,
	0x33, 0x49, 0xc7, 0xbb, 0x3a, 0xab, 0xeb, 0xdf, 0x91, 0x9f, 0x18, 0x46, 0x70, 0x64, 0x83, 0xb6,
	0xc6, 0x1c, 0x3c, 0x62, 0xcc, 0xb6, 0x93, 0x49, 0x39, 0xbe, 0xb8, 0x59, 0x85, 0xfe, 0xed, 0x2a,
	0xf4, 0xff, 0xac, 0x42, 0xff, 0x7a, 0x1d, 0x7a, 0xb7, 0xeb, 0xd0, 0xfb, 0xb5, 0x0e, 0xbd, 0x2f,
	0xc7, 0xac, 0x50, 0xf3, 0x7a, 0xa6, 0x75, 0xc8, 0x83, 0x7f, 0xec, 0x0e, 0xb4, 0x2a, 0x48, 0xe3,
	0xef, 0x9e, 0x75, 0xcc, 0xa8, 0x47, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xfa, 0xee, 0xe2,
	0x56, 0x03, 0x00, 0x00,
}

func (m *MsgInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PeerID) > 0 {
		i -= len(m.PeerID)
		copy(dAtA[i:], m.PeerID)
		i = encodeVarintWal(dAtA, i, uint64(len(m.PeerID)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Msg.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintWal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *WALMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WALMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WALMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *WALMessage_EventDataRoundState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WALMessage_EventDataRoundState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.EventDataRoundState != nil {
		{
			size, err := m.EventDataRoundState.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *WALMessage_MsgInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WALMessage_MsgInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.MsgInfo != nil {
		{
			size, err := m.MsgInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *WALMessage_TimeoutInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WALMessage_TimeoutInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.TimeoutInfo != nil {
		{
			size, err := m.TimeoutInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *WALMessage_EndHeight) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WALMessage_EndHeight) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.EndHeight != nil {
		{
			size, err := m.EndHeight.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *TimedWALMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimedWALMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TimedWALMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Msg != nil {
		{
			size, err := m.Msg.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	n7, err7 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time):])
	if err7 != nil {
		return 0, err7
	}
	i -= n7
	i = encodeVarintWal(dAtA, i, uint64(n7))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintWal(dAtA []byte, offset int, v uint64) int {
	offset -= sovWal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Msg.Size()
	n += 1 + l + sovWal(uint64(l))
	l = len(m.PeerID)
	if l > 0 {
		n += 1 + l + sovWal(uint64(l))
	}
	return n
}

func (m *WALMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *WALMessage_EventDataRoundState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EventDataRoundState != nil {
		l = m.EventDataRoundState.Size()
		n += 1 + l + sovWal(uint64(l))
	}
	return n
}
func (m *WALMessage_MsgInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MsgInfo != nil {
		l = m.MsgInfo.Size()
		n += 1 + l + sovWal(uint64(l))
	}
	return n
}
func (m *WALMessage_TimeoutInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TimeoutInfo != nil {
		l = m.TimeoutInfo.Size()
		n += 1 + l + sovWal(uint64(l))
	}
	return n
}
func (m *WALMessage_EndHeight) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EndHeight != nil {
		l = m.EndHeight.Size()
		n += 1 + l + sovWal(uint64(l))
	}
	return n
}
func (m *TimedWALMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovWal(uint64(l))
	if m.Msg != nil {
		l = m.Msg.Size()
		n += 1 + l + sovWal(uint64(l))
	}
	return n
}

func sovWal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWal(x uint64) (n int) {
	return sovWal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWal
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
			return fmt.Errorf("proto: MsgInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
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
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Msg.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
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
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWal
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
func (m *WALMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWal
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
			return fmt.Errorf("proto: WALMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WALMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventDataRoundState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
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
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &v1.EventDataRoundState{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &WALMessage_EventDataRoundState{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
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
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &MsgInfo{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &WALMessage_MsgInfo{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
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
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &v11.TimeoutInfo{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &WALMessage_TimeoutInfo{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
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
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &v11.EndHeight{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &WALMessage_EndHeight{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWal
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
func (m *TimedWALMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWal
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
			return fmt.Errorf("proto: TimedWALMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimedWALMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
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
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
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
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Msg == nil {
				m.Msg = &WALMessage{}
			}
			if err := m.Msg.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWal
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
func skipWal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWal
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
					return 0, ErrIntOverflowWal
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
					return 0, ErrIntOverflowWal
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
				return 0, ErrInvalidLengthWal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWal = fmt.Errorf("proto: unexpected end of group")
)