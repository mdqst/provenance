// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/msgfees/v1/proposals.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// AddMsgFeeProposal defines a governance proposal to add additional msg based fee
type AddMsgFeeProposal struct {
	Title         string     `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string     `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	MsgTypeUrl    string     `protobuf:"bytes,3,opt,name=msg_type_url,json=msgTypeUrl,proto3" json:"msg_type_url,omitempty"`
	AdditionalFee types.Coin `protobuf:"bytes,4,opt,name=additional_fee,json=additionalFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"additional_fee" yaml:"additional_fee"`
}

func (m *AddMsgFeeProposal) Reset()      { *m = AddMsgFeeProposal{} }
func (*AddMsgFeeProposal) ProtoMessage() {}
func (*AddMsgFeeProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e168825d6c34a4, []int{0}
}
func (m *AddMsgFeeProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddMsgFeeProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddMsgFeeProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddMsgFeeProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMsgFeeProposal.Merge(m, src)
}
func (m *AddMsgFeeProposal) XXX_Size() int {
	return m.Size()
}
func (m *AddMsgFeeProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMsgFeeProposal.DiscardUnknown(m)
}

var xxx_messageInfo_AddMsgFeeProposal proto.InternalMessageInfo

func (m *AddMsgFeeProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *AddMsgFeeProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AddMsgFeeProposal) GetMsgTypeUrl() string {
	if m != nil {
		return m.MsgTypeUrl
	}
	return ""
}

func (m *AddMsgFeeProposal) GetAdditionalFee() types.Coin {
	if m != nil {
		return m.AdditionalFee
	}
	return types.Coin{}
}

// UpdateMsgFeeProposal defines a governance proposal to update a current msg based fee
type UpdateMsgFeeProposal struct {
	Title         string     `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string     `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	MsgTypeUrl    string     `protobuf:"bytes,3,opt,name=msg_type_url,json=msgTypeUrl,proto3" json:"msg_type_url,omitempty"`
	AdditionalFee types.Coin `protobuf:"bytes,4,opt,name=additional_fee,json=additionalFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"additional_fee" yaml:"additional_fee"`
}

func (m *UpdateMsgFeeProposal) Reset()      { *m = UpdateMsgFeeProposal{} }
func (*UpdateMsgFeeProposal) ProtoMessage() {}
func (*UpdateMsgFeeProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e168825d6c34a4, []int{1}
}
func (m *UpdateMsgFeeProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateMsgFeeProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateMsgFeeProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateMsgFeeProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMsgFeeProposal.Merge(m, src)
}
func (m *UpdateMsgFeeProposal) XXX_Size() int {
	return m.Size()
}
func (m *UpdateMsgFeeProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMsgFeeProposal.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMsgFeeProposal proto.InternalMessageInfo

func (m *UpdateMsgFeeProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UpdateMsgFeeProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateMsgFeeProposal) GetMsgTypeUrl() string {
	if m != nil {
		return m.MsgTypeUrl
	}
	return ""
}

func (m *UpdateMsgFeeProposal) GetAdditionalFee() types.Coin {
	if m != nil {
		return m.AdditionalFee
	}
	return types.Coin{}
}

// RemoveMsgFeeProposal defines a governance proposal to delete a current msg based fee
type RemoveMsgFeeProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	MsgTypeUrl  string `protobuf:"bytes,3,opt,name=msg_type_url,json=msgTypeUrl,proto3" json:"msg_type_url,omitempty"`
}

func (m *RemoveMsgFeeProposal) Reset()      { *m = RemoveMsgFeeProposal{} }
func (*RemoveMsgFeeProposal) ProtoMessage() {}
func (*RemoveMsgFeeProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e168825d6c34a4, []int{2}
}
func (m *RemoveMsgFeeProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RemoveMsgFeeProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RemoveMsgFeeProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RemoveMsgFeeProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveMsgFeeProposal.Merge(m, src)
}
func (m *RemoveMsgFeeProposal) XXX_Size() int {
	return m.Size()
}
func (m *RemoveMsgFeeProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveMsgFeeProposal.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveMsgFeeProposal proto.InternalMessageInfo

func (m *RemoveMsgFeeProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *RemoveMsgFeeProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RemoveMsgFeeProposal) GetMsgTypeUrl() string {
	if m != nil {
		return m.MsgTypeUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*AddMsgFeeProposal)(nil), "provenance.msgfees.v1.AddMsgFeeProposal")
	proto.RegisterType((*UpdateMsgFeeProposal)(nil), "provenance.msgfees.v1.UpdateMsgFeeProposal")
	proto.RegisterType((*RemoveMsgFeeProposal)(nil), "provenance.msgfees.v1.RemoveMsgFeeProposal")
}

func init() {
	proto.RegisterFile("provenance/msgfees/v1/proposals.proto", fileDescriptor_a2e168825d6c34a4)
}

var fileDescriptor_a2e168825d6c34a4 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x93, 0xbf, 0x0b, 0xd3, 0x40,
	0x14, 0xc7, 0x73, 0xf5, 0x07, 0x7a, 0x55, 0xc1, 0x90, 0x42, 0xec, 0x90, 0x84, 0x80, 0xd0, 0xa5,
	0x77, 0xc4, 0x6e, 0xdd, 0xac, 0x50, 0x70, 0x10, 0x4a, 0xb1, 0x8b, 0x4b, 0xb9, 0x24, 0xaf, 0xf1,
	0x30, 0xc9, 0x1d, 0xb9, 0x6b, 0xb0, 0xee, 0x82, 0xe0, 0xe2, 0xe8, 0xd8, 0xd9, 0xbf, 0xa4, 0x63,
	0x47, 0xa7, 0x2a, 0xed, 0xe2, 0xec, 0x5f, 0x20, 0xf9, 0x81, 0xa9, 0xb8, 0xbb, 0x38, 0xdd, 0xbd,
	0x7b, 0x9f, 0x77, 0xef, 0xfb, 0x1e, 0xef, 0xe1, 0xc7, 0xb2, 0x10, 0x25, 0xe4, 0x2c, 0x8f, 0x80,
	0x66, 0x2a, 0xd9, 0x00, 0x28, 0x5a, 0x06, 0x54, 0x16, 0x42, 0x0a, 0xc5, 0x52, 0x45, 0x64, 0x21,
	0xb4, 0x30, 0x07, 0x1d, 0x46, 0x5a, 0x8c, 0x94, 0xc1, 0xd0, 0x4a, 0x44, 0x22, 0x6a, 0x82, 0x56,
	0xb7, 0x06, 0x1e, 0x3a, 0x91, 0x50, 0x99, 0x50, 0x34, 0x64, 0x0a, 0x68, 0x19, 0x84, 0xa0, 0x59,
	0x40, 0x23, 0xc1, 0xf3, 0xc6, 0xef, 0xbf, 0xef, 0xe1, 0x87, 0x4f, 0xe3, 0xf8, 0x85, 0x4a, 0xe6,
	0x00, 0x8b, 0x36, 0x93, 0x69, 0xe1, 0x5b, 0x9a, 0xeb, 0x14, 0x6c, 0xe4, 0xa1, 0xd1, 0xdd, 0x65,
	0x63, 0x98, 0x1e, 0xee, 0xc7, 0xa0, 0xa2, 0x82, 0x4b, 0xcd, 0x45, 0x6e, 0xf7, 0x6a, 0xdf, 0xf5,
	0x93, 0xe9, 0xe1, 0x7b, 0x99, 0x4a, 0xd6, 0x7a, 0x27, 0x61, 0xbd, 0x2d, 0x52, 0xfb, 0x46, 0x8d,
	0xe0, 0x4c, 0x25, 0x2f, 0x77, 0x12, 0x56, 0x45, 0x6a, 0x7e, 0x44, 0xf8, 0x01, 0x8b, 0x63, 0x5e,
	0xe1, 0x2c, 0x5d, 0x6f, 0x00, 0xec, 0x9b, 0x1e, 0x1a, 0xf5, 0x9f, 0x3c, 0x22, 0x8d, 0x52, 0x52,
	0x29, 0x25, 0xad, 0x52, 0xf2, 0x4c, 0xf0, 0x7c, 0xf6, 0xfc, 0x70, 0x72, 0x8d, 0x9f, 0x27, 0x77,
	0xb0, 0x63, 0x59, 0x3a, 0xf5, 0xff, 0x0c, 0xf7, 0xbf, 0x7c, 0x73, 0x47, 0x09, 0xd7, 0xaf, 0xb7,
	0x21, 0x89, 0x44, 0x46, 0xdb, 0x7a, 0x9b, 0x63, 0xac, 0xe2, 0x37, 0xb4, 0x52, 0xa3, 0xea, 0x9f,
	0xd4, 0xf2, 0x7e, 0x17, 0x3c, 0x07, 0x98, 0xde, 0xf9, 0xbc, 0x77, 0x8d, 0x1f, 0x7b, 0x17, 0xf9,
	0x1f, 0x7a, 0xd8, 0x5a, 0xc9, 0x98, 0x69, 0xf8, 0xef, 0x5b, 0xf1, 0x0e, 0x5b, 0x4b, 0xc8, 0x44,
	0xf9, 0xcf, 0x3a, 0xd1, 0xe5, 0x9e, 0xf1, 0xc3, 0xd9, 0x41, 0xc7, 0xb3, 0x83, 0xbe, 0x9f, 0x1d,
	0xf4, 0xe9, 0xe2, 0x18, 0xc7, 0x8b, 0x63, 0x7c, 0xbd, 0x38, 0x06, 0xb6, 0x79, 0x3d, 0xd2, 0x7f,
	0x0f, 0xfe, 0x02, 0xbd, 0x9a, 0x5c, 0x15, 0xdd, 0x31, 0x63, 0x2e, 0xae, 0x2c, 0xfa, 0xf6, 0xf7,
	0x4e, 0xd5, 0x5d, 0x08, 0x6f, 0xd7, 0x0b, 0x30, 0xf9, 0x15, 0x00, 0x00, 0xff, 0xff, 0x36, 0x1f,
	0x37, 0xce, 0x76, 0x03, 0x00, 0x00,
}

func (this *AddMsgFeeProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AddMsgFeeProposal)
	if !ok {
		that2, ok := that.(AddMsgFeeProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.MsgTypeUrl != that1.MsgTypeUrl {
		return false
	}
	if !this.AdditionalFee.Equal(&that1.AdditionalFee) {
		return false
	}
	return true
}
func (this *UpdateMsgFeeProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpdateMsgFeeProposal)
	if !ok {
		that2, ok := that.(UpdateMsgFeeProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.MsgTypeUrl != that1.MsgTypeUrl {
		return false
	}
	if !this.AdditionalFee.Equal(&that1.AdditionalFee) {
		return false
	}
	return true
}
func (this *RemoveMsgFeeProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RemoveMsgFeeProposal)
	if !ok {
		that2, ok := that.(RemoveMsgFeeProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.MsgTypeUrl != that1.MsgTypeUrl {
		return false
	}
	return true
}
func (m *AddMsgFeeProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddMsgFeeProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddMsgFeeProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.AdditionalFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProposals(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.MsgTypeUrl) > 0 {
		i -= len(m.MsgTypeUrl)
		copy(dAtA[i:], m.MsgTypeUrl)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.MsgTypeUrl)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpdateMsgFeeProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateMsgFeeProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateMsgFeeProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.AdditionalFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProposals(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.MsgTypeUrl) > 0 {
		i -= len(m.MsgTypeUrl)
		copy(dAtA[i:], m.MsgTypeUrl)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.MsgTypeUrl)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RemoveMsgFeeProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoveMsgFeeProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RemoveMsgFeeProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MsgTypeUrl) > 0 {
		i -= len(m.MsgTypeUrl)
		copy(dAtA[i:], m.MsgTypeUrl)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.MsgTypeUrl)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposals(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposals(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AddMsgFeeProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.MsgTypeUrl)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = m.AdditionalFee.Size()
	n += 1 + l + sovProposals(uint64(l))
	return n
}

func (m *UpdateMsgFeeProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.MsgTypeUrl)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = m.AdditionalFee.Size()
	n += 1 + l + sovProposals(uint64(l))
	return n
}

func (m *RemoveMsgFeeProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.MsgTypeUrl)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	return n
}

func sovProposals(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposals(x uint64) (n int) {
	return sovProposals(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AddMsgFeeProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposals
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
			return fmt.Errorf("proto: AddMsgFeeProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddMsgFeeProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgTypeUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgTypeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AdditionalFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposals(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposals
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
func (m *UpdateMsgFeeProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposals
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
			return fmt.Errorf("proto: UpdateMsgFeeProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateMsgFeeProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgTypeUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgTypeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AdditionalFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposals(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposals
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
func (m *RemoveMsgFeeProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposals
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
			return fmt.Errorf("proto: RemoveMsgFeeProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoveMsgFeeProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgTypeUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgTypeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposals(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposals
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
func skipProposals(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposals
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
					return 0, ErrIntOverflowProposals
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
					return 0, ErrIntOverflowProposals
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
				return 0, ErrInvalidLengthProposals
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposals
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposals
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposals        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposals          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposals = fmt.Errorf("proto: unexpected end of group")
)
