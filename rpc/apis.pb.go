// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/apis.proto

package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import block "github.com/iost-official/Go-IOS-Protocol/core/block"
import event "github.com/iost-official/Go-IOS-Protocol/core/event"
import tx "github.com/iost-official/Go-IOS-Protocol/core/tx"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type VoidReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoidReq) Reset()         { *m = VoidReq{} }
func (m *VoidReq) String() string { return proto.CompactTextString(m) }
func (*VoidReq) ProtoMessage()    {}
func (*VoidReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_apis_3a13eb2c4442489e, []int{0}
}
func (m *VoidReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoidReq.Unmarshal(m, b)
}
func (m *VoidReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoidReq.Marshal(b, m, deterministic)
}
func (dst *VoidReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoidReq.Merge(dst, src)
}
func (m *VoidReq) XXX_Size() int {
	return xxx_messageInfo_VoidReq.Size(m)
}
func (m *VoidReq) XXX_DiscardUnknown() {
	xxx_messageInfo_VoidReq.DiscardUnknown(m)
}

var xxx_messageInfo_VoidReq proto.InternalMessageInfo

type HashReq struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashReq) Reset()         { *m = HashReq{} }
func (m *HashReq) String() string { return proto.CompactTextString(m) }
func (*HashReq) ProtoMessage()    {}
func (*HashReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_apis_3a13eb2c4442489e, []int{1}
}
func (m *HashReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashReq.Unmarshal(m, b)
}
func (m *HashReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashReq.Marshal(b, m, deterministic)
}
func (dst *HashReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashReq.Merge(dst, src)
}
func (m *HashReq) XXX_Size() int {
	return xxx_messageInfo_HashReq.Size(m)
}
func (m *HashReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HashReq.DiscardUnknown(m)
}

var xxx_messageInfo_HashReq proto.InternalMessageInfo

func (m *HashReq) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type BlockByHashReq struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Complete             bool     `protobuf:"varint,2,opt,name=complete" json:"complete,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockByHashReq) Reset()         { *m = BlockByHashReq{} }
func (m *BlockByHashReq) String() string { return proto.CompactTextString(m) }
func (*BlockByHashReq) ProtoMessage()    {}
func (*BlockByHashReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_apis_3a13eb2c4442489e, []int{2}
}
func (m *BlockByHashReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockByHashReq.Unmarshal(m, b)
}
func (m *BlockByHashReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockByHashReq.Marshal(b, m, deterministic)
}
func (dst *BlockByHashReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockByHashReq.Merge(dst, src)
}
func (m *BlockByHashReq) XXX_Size() int {
	return xxx_messageInfo_BlockByHashReq.Size(m)
}
func (m *BlockByHashReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockByHashReq.DiscardUnknown(m)
}

var xxx_messageInfo_BlockByHashReq proto.InternalMessageInfo

func (m *BlockByHashReq) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *BlockByHashReq) GetComplete() bool {
	if m != nil {
		return m.Complete
	}
	return false
}

type BlockByNumReq struct {
	Num                  int64    `protobuf:"varint,1,opt,name=num" json:"num,omitempty"`
	Complete             bool     `protobuf:"varint,2,opt,name=complete" json:"complete,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockByNumReq) Reset()         { *m = BlockByNumReq{} }
func (m *BlockByNumReq) String() string { return proto.CompactTextString(m) }
func (*BlockByNumReq) ProtoMessage()    {}
func (*BlockByNumReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_apis_3a13eb2c4442489e, []int{3}
}
func (m *BlockByNumReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockByNumReq.Unmarshal(m, b)
}
func (m *BlockByNumReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockByNumReq.Marshal(b, m, deterministic)
}
func (dst *BlockByNumReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockByNumReq.Merge(dst, src)
}
func (m *BlockByNumReq) XXX_Size() int {
	return xxx_messageInfo_BlockByNumReq.Size(m)
}
func (m *BlockByNumReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockByNumReq.DiscardUnknown(m)
}

var xxx_messageInfo_BlockByNumReq proto.InternalMessageInfo

func (m *BlockByNumReq) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *BlockByNumReq) GetComplete() bool {
	if m != nil {
		return m.Complete
	}
	return false
}

type GetBalanceReq struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBalanceReq) Reset()         { *m = GetBalanceReq{} }
func (m *GetBalanceReq) String() string { return proto.CompactTextString(m) }
func (*GetBalanceReq) ProtoMessage()    {}
func (*GetBalanceReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_apis_3a13eb2c4442489e, []int{4}
}
func (m *GetBalanceReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBalanceReq.Unmarshal(m, b)
}
func (m *GetBalanceReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBalanceReq.Marshal(b, m, deterministic)
}
func (dst *GetBalanceReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBalanceReq.Merge(dst, src)
}
func (m *GetBalanceReq) XXX_Size() int {
	return xxx_messageInfo_GetBalanceReq.Size(m)
}
func (m *GetBalanceReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBalanceReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetBalanceReq proto.InternalMessageInfo

func (m *GetBalanceReq) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type GetStateReq struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStateReq) Reset()         { *m = GetStateReq{} }
func (m *GetStateReq) String() string { return proto.CompactTextString(m) }
func (*GetStateReq) ProtoMessage()    {}
func (*GetStateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_apis_3a13eb2c4442489e, []int{5}
}
func (m *GetStateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStateReq.Unmarshal(m, b)
}
func (m *GetStateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStateReq.Marshal(b, m, deterministic)
}
func (dst *GetStateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStateReq.Merge(dst, src)
}
func (m *GetStateReq) XXX_Size() int {
	return xxx_messageInfo_GetStateReq.Size(m)
}
func (m *GetStateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStateReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetStateReq proto.InternalMessageInfo

func (m *GetStateReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type RawTxReq struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RawTxReq) Reset()         { *m = RawTxReq{} }
func (m *RawTxReq) String() string { return proto.CompactTextString(m) }
func (*RawTxReq) ProtoMessage()    {}
func (*RawTxReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_apis_3a13eb2c4442489e, []int{6}
}
func (m *RawTxReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RawTxReq.Unmarshal(m, b)
}
func (m *RawTxReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RawTxReq.Marshal(b, m, deterministic)
}
func (dst *RawTxReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawTxReq.Merge(dst, src)
}
func (m *RawTxReq) XXX_Size() int {
	return xxx_messageInfo_RawTxReq.Size(m)
}
func (m *RawTxReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RawTxReq.DiscardUnknown(m)
}

var xxx_messageInfo_RawTxReq proto.InternalMessageInfo

func (m *RawTxReq) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type SubscribeReq struct {
	Topics               []event.Event_Topic `protobuf:"varint,1,rep,packed,name=topics,enum=event.Event_Topic" json:"topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SubscribeReq) Reset()         { *m = SubscribeReq{} }
func (m *SubscribeReq) String() string { return proto.CompactTextString(m) }
func (*SubscribeReq) ProtoMessage()    {}
func (*SubscribeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_apis_3a13eb2c4442489e, []int{7}
}
func (m *SubscribeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeReq.Unmarshal(m, b)
}
func (m *SubscribeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeReq.Marshal(b, m, deterministic)
}
func (dst *SubscribeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeReq.Merge(dst, src)
}
func (m *SubscribeReq) XXX_Size() int {
	return xxx_messageInfo_SubscribeReq.Size(m)
}
func (m *SubscribeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeReq.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeReq proto.InternalMessageInfo

func (m *SubscribeReq) GetTopics() []event.Event_Topic {
	if m != nil {
		return m.Topics
	}
	return nil
}

type HeightRes struct {
	Height               int64    `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeightRes) Reset()         { *m = HeightRes{} }
func (m *HeightRes) String() string { return proto.CompactTextString(m) }
func (*HeightRes) ProtoMessage()    {}
func (*HeightRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_apis_3a13eb2c4442489e, []int{8}
}
func (m *HeightRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeightRes.Unmarshal(m, b)
}
func (m *HeightRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeightRes.Marshal(b, m, deterministic)
}
func (dst *HeightRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeightRes.Merge(dst, src)
}
func (m *HeightRes) XXX_Size() int {
	return xxx_messageInfo_HeightRes.Size(m)
}
func (m *HeightRes) XXX_DiscardUnknown() {
	xxx_messageInfo_HeightRes.DiscardUnknown(m)
}

var xxx_messageInfo_HeightRes proto.InternalMessageInfo

func (m *HeightRes) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type GetBalanceRes struct {
	Balance              int64    `protobuf:"varint,1,opt,name=balance" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBalanceRes) Reset()         { *m = GetBalanceRes{} }
func (m *GetBalanceRes) String() string { return proto.CompactTextString(m) }
func (*GetBalanceRes) ProtoMessage()    {}
func (*GetBalanceRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_apis_3a13eb2c4442489e, []int{9}
}
func (m *GetBalanceRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBalanceRes.Unmarshal(m, b)
}
func (m *GetBalanceRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBalanceRes.Marshal(b, m, deterministic)
}
func (dst *GetBalanceRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBalanceRes.Merge(dst, src)
}
func (m *GetBalanceRes) XXX_Size() int {
	return xxx_messageInfo_GetBalanceRes.Size(m)
}
func (m *GetBalanceRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBalanceRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetBalanceRes proto.InternalMessageInfo

func (m *GetBalanceRes) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

type GetStateRes struct {
	Value                string   `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStateRes) Reset()         { *m = GetStateRes{} }
func (m *GetStateRes) String() string { return proto.CompactTextString(m) }
func (*GetStateRes) ProtoMessage()    {}
func (*GetStateRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_apis_3a13eb2c4442489e, []int{10}
}
func (m *GetStateRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStateRes.Unmarshal(m, b)
}
func (m *GetStateRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStateRes.Marshal(b, m, deterministic)
}
func (dst *GetStateRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStateRes.Merge(dst, src)
}
func (m *GetStateRes) XXX_Size() int {
	return xxx_messageInfo_GetStateRes.Size(m)
}
func (m *GetStateRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStateRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetStateRes proto.InternalMessageInfo

func (m *GetStateRes) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SendRawTxRes struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendRawTxRes) Reset()         { *m = SendRawTxRes{} }
func (m *SendRawTxRes) String() string { return proto.CompactTextString(m) }
func (*SendRawTxRes) ProtoMessage()    {}
func (*SendRawTxRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_apis_3a13eb2c4442489e, []int{11}
}
func (m *SendRawTxRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRawTxRes.Unmarshal(m, b)
}
func (m *SendRawTxRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRawTxRes.Marshal(b, m, deterministic)
}
func (dst *SendRawTxRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRawTxRes.Merge(dst, src)
}
func (m *SendRawTxRes) XXX_Size() int {
	return xxx_messageInfo_SendRawTxRes.Size(m)
}
func (m *SendRawTxRes) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRawTxRes.DiscardUnknown(m)
}

var xxx_messageInfo_SendRawTxRes proto.InternalMessageInfo

func (m *SendRawTxRes) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type GasRes struct {
	Gas                  uint64   `protobuf:"varint,1,opt,name=gas" json:"gas,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GasRes) Reset()         { *m = GasRes{} }
func (m *GasRes) String() string { return proto.CompactTextString(m) }
func (*GasRes) ProtoMessage()    {}
func (*GasRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_apis_3a13eb2c4442489e, []int{12}
}
func (m *GasRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GasRes.Unmarshal(m, b)
}
func (m *GasRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GasRes.Marshal(b, m, deterministic)
}
func (dst *GasRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GasRes.Merge(dst, src)
}
func (m *GasRes) XXX_Size() int {
	return xxx_messageInfo_GasRes.Size(m)
}
func (m *GasRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GasRes.DiscardUnknown(m)
}

var xxx_messageInfo_GasRes proto.InternalMessageInfo

func (m *GasRes) GetGas() uint64 {
	if m != nil {
		return m.Gas
	}
	return 0
}

type BlockInfo struct {
	Head                 *block.BlockHead `protobuf:"bytes,1,opt,name=head" json:"head,omitempty"`
	Txs                  []*tx.TxRaw      `protobuf:"bytes,2,rep,name=txs" json:"txs,omitempty"`
	Txhash               [][]byte         `protobuf:"bytes,3,rep,name=txhash,proto3" json:"txhash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BlockInfo) Reset()         { *m = BlockInfo{} }
func (m *BlockInfo) String() string { return proto.CompactTextString(m) }
func (*BlockInfo) ProtoMessage()    {}
func (*BlockInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_apis_3a13eb2c4442489e, []int{13}
}
func (m *BlockInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockInfo.Unmarshal(m, b)
}
func (m *BlockInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockInfo.Marshal(b, m, deterministic)
}
func (dst *BlockInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockInfo.Merge(dst, src)
}
func (m *BlockInfo) XXX_Size() int {
	return xxx_messageInfo_BlockInfo.Size(m)
}
func (m *BlockInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BlockInfo proto.InternalMessageInfo

func (m *BlockInfo) GetHead() *block.BlockHead {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *BlockInfo) GetTxs() []*tx.TxRaw {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *BlockInfo) GetTxhash() [][]byte {
	if m != nil {
		return m.Txhash
	}
	return nil
}

type SubscribeRes struct {
	Ev                   *event.Event `protobuf:"bytes,1,opt,name=ev" json:"ev,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SubscribeRes) Reset()         { *m = SubscribeRes{} }
func (m *SubscribeRes) String() string { return proto.CompactTextString(m) }
func (*SubscribeRes) ProtoMessage()    {}
func (*SubscribeRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_apis_3a13eb2c4442489e, []int{14}
}
func (m *SubscribeRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRes.Unmarshal(m, b)
}
func (m *SubscribeRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRes.Marshal(b, m, deterministic)
}
func (dst *SubscribeRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRes.Merge(dst, src)
}
func (m *SubscribeRes) XXX_Size() int {
	return xxx_messageInfo_SubscribeRes.Size(m)
}
func (m *SubscribeRes) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRes.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRes proto.InternalMessageInfo

func (m *SubscribeRes) GetEv() *event.Event {
	if m != nil {
		return m.Ev
	}
	return nil
}

func init() {
	proto.RegisterType((*VoidReq)(nil), "rpc.VoidReq")
	proto.RegisterType((*HashReq)(nil), "rpc.HashReq")
	proto.RegisterType((*BlockByHashReq)(nil), "rpc.BlockByHashReq")
	proto.RegisterType((*BlockByNumReq)(nil), "rpc.BlockByNumReq")
	proto.RegisterType((*GetBalanceReq)(nil), "rpc.GetBalanceReq")
	proto.RegisterType((*GetStateReq)(nil), "rpc.GetStateReq")
	proto.RegisterType((*RawTxReq)(nil), "rpc.RawTxReq")
	proto.RegisterType((*SubscribeReq)(nil), "rpc.SubscribeReq")
	proto.RegisterType((*HeightRes)(nil), "rpc.HeightRes")
	proto.RegisterType((*GetBalanceRes)(nil), "rpc.GetBalanceRes")
	proto.RegisterType((*GetStateRes)(nil), "rpc.GetStateRes")
	proto.RegisterType((*SendRawTxRes)(nil), "rpc.SendRawTxRes")
	proto.RegisterType((*GasRes)(nil), "rpc.GasRes")
	proto.RegisterType((*BlockInfo)(nil), "rpc.BlockInfo")
	proto.RegisterType((*SubscribeRes)(nil), "rpc.SubscribeRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ApisClient is the client API for Apis service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApisClient interface {
	GetHeight(ctx context.Context, in *VoidReq, opts ...grpc.CallOption) (*HeightRes, error)
	GetTxByHash(ctx context.Context, in *HashReq, opts ...grpc.CallOption) (*tx.TxRaw, error)
	GetBlockByHash(ctx context.Context, in *BlockByHashReq, opts ...grpc.CallOption) (*BlockInfo, error)
	GetBlockByNum(ctx context.Context, in *BlockByNumReq, opts ...grpc.CallOption) (*BlockInfo, error)
	GetBalance(ctx context.Context, in *GetBalanceReq, opts ...grpc.CallOption) (*GetBalanceRes, error)
	GetState(ctx context.Context, in *GetStateReq, opts ...grpc.CallOption) (*GetStateRes, error)
	SendRawTx(ctx context.Context, in *RawTxReq, opts ...grpc.CallOption) (*SendRawTxRes, error)
	EstimateGas(ctx context.Context, in *RawTxReq, opts ...grpc.CallOption) (*GasRes, error)
	Subscribe(ctx context.Context, in *SubscribeReq, opts ...grpc.CallOption) (Apis_SubscribeClient, error)
}

type apisClient struct {
	cc *grpc.ClientConn
}

func NewApisClient(cc *grpc.ClientConn) ApisClient {
	return &apisClient{cc}
}

func (c *apisClient) GetHeight(ctx context.Context, in *VoidReq, opts ...grpc.CallOption) (*HeightRes, error) {
	out := new(HeightRes)
	err := c.cc.Invoke(ctx, "/rpc.Apis/GetHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apisClient) GetTxByHash(ctx context.Context, in *HashReq, opts ...grpc.CallOption) (*tx.TxRaw, error) {
	out := new(tx.TxRaw)
	err := c.cc.Invoke(ctx, "/rpc.Apis/GetTxByHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apisClient) GetBlockByHash(ctx context.Context, in *BlockByHashReq, opts ...grpc.CallOption) (*BlockInfo, error) {
	out := new(BlockInfo)
	err := c.cc.Invoke(ctx, "/rpc.Apis/GetBlockByHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apisClient) GetBlockByNum(ctx context.Context, in *BlockByNumReq, opts ...grpc.CallOption) (*BlockInfo, error) {
	out := new(BlockInfo)
	err := c.cc.Invoke(ctx, "/rpc.Apis/getBlockByNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apisClient) GetBalance(ctx context.Context, in *GetBalanceReq, opts ...grpc.CallOption) (*GetBalanceRes, error) {
	out := new(GetBalanceRes)
	err := c.cc.Invoke(ctx, "/rpc.Apis/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apisClient) GetState(ctx context.Context, in *GetStateReq, opts ...grpc.CallOption) (*GetStateRes, error) {
	out := new(GetStateRes)
	err := c.cc.Invoke(ctx, "/rpc.Apis/GetState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apisClient) SendRawTx(ctx context.Context, in *RawTxReq, opts ...grpc.CallOption) (*SendRawTxRes, error) {
	out := new(SendRawTxRes)
	err := c.cc.Invoke(ctx, "/rpc.Apis/SendRawTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apisClient) EstimateGas(ctx context.Context, in *RawTxReq, opts ...grpc.CallOption) (*GasRes, error) {
	out := new(GasRes)
	err := c.cc.Invoke(ctx, "/rpc.Apis/EstimateGas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apisClient) Subscribe(ctx context.Context, in *SubscribeReq, opts ...grpc.CallOption) (Apis_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Apis_serviceDesc.Streams[0], "/rpc.Apis/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &apisSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Apis_SubscribeClient interface {
	Recv() (*SubscribeRes, error)
	grpc.ClientStream
}

type apisSubscribeClient struct {
	grpc.ClientStream
}

func (x *apisSubscribeClient) Recv() (*SubscribeRes, error) {
	m := new(SubscribeRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Apis service

type ApisServer interface {
	GetHeight(context.Context, *VoidReq) (*HeightRes, error)
	GetTxByHash(context.Context, *HashReq) (*tx.TxRaw, error)
	GetBlockByHash(context.Context, *BlockByHashReq) (*BlockInfo, error)
	GetBlockByNum(context.Context, *BlockByNumReq) (*BlockInfo, error)
	GetBalance(context.Context, *GetBalanceReq) (*GetBalanceRes, error)
	GetState(context.Context, *GetStateReq) (*GetStateRes, error)
	SendRawTx(context.Context, *RawTxReq) (*SendRawTxRes, error)
	EstimateGas(context.Context, *RawTxReq) (*GasRes, error)
	Subscribe(*SubscribeReq, Apis_SubscribeServer) error
}

func RegisterApisServer(s *grpc.Server, srv ApisServer) {
	s.RegisterService(&_Apis_serviceDesc, srv)
}

func _Apis_GetHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApisServer).GetHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Apis/GetHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApisServer).GetHeight(ctx, req.(*VoidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apis_GetTxByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApisServer).GetTxByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Apis/GetTxByHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApisServer).GetTxByHash(ctx, req.(*HashReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apis_GetBlockByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockByHashReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApisServer).GetBlockByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Apis/GetBlockByHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApisServer).GetBlockByHash(ctx, req.(*BlockByHashReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apis_GetBlockByNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockByNumReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApisServer).GetBlockByNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Apis/GetBlockByNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApisServer).GetBlockByNum(ctx, req.(*BlockByNumReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apis_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApisServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Apis/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApisServer).GetBalance(ctx, req.(*GetBalanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apis_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApisServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Apis/GetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApisServer).GetState(ctx, req.(*GetStateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apis_SendRawTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawTxReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApisServer).SendRawTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Apis/SendRawTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApisServer).SendRawTx(ctx, req.(*RawTxReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apis_EstimateGas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawTxReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApisServer).EstimateGas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Apis/EstimateGas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApisServer).EstimateGas(ctx, req.(*RawTxReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apis_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ApisServer).Subscribe(m, &apisSubscribeServer{stream})
}

type Apis_SubscribeServer interface {
	Send(*SubscribeRes) error
	grpc.ServerStream
}

type apisSubscribeServer struct {
	grpc.ServerStream
}

func (x *apisSubscribeServer) Send(m *SubscribeRes) error {
	return x.ServerStream.SendMsg(m)
}

var _Apis_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Apis",
	HandlerType: (*ApisServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHeight",
			Handler:    _Apis_GetHeight_Handler,
		},
		{
			MethodName: "GetTxByHash",
			Handler:    _Apis_GetTxByHash_Handler,
		},
		{
			MethodName: "GetBlockByHash",
			Handler:    _Apis_GetBlockByHash_Handler,
		},
		{
			MethodName: "getBlockByNum",
			Handler:    _Apis_GetBlockByNum_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _Apis_GetBalance_Handler,
		},
		{
			MethodName: "GetState",
			Handler:    _Apis_GetState_Handler,
		},
		{
			MethodName: "SendRawTx",
			Handler:    _Apis_SendRawTx_Handler,
		},
		{
			MethodName: "EstimateGas",
			Handler:    _Apis_EstimateGas_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Apis_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpc/apis.proto",
}

func init() { proto.RegisterFile("rpc/apis.proto", fileDescriptor_apis_3a13eb2c4442489e) }

var fileDescriptor_apis_3a13eb2c4442489e = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x4d, 0x6f, 0xda, 0x4c,
	0x10, 0x06, 0x9b, 0x97, 0xe0, 0xc1, 0x58, 0xbc, 0xdb, 0x2a, 0x42, 0xee, 0x47, 0xd0, 0xa6, 0x07,
	0xb7, 0xa9, 0x4c, 0x44, 0xd5, 0x56, 0xaa, 0x54, 0xa9, 0x8d, 0x12, 0x91, 0x5c, 0x72, 0x58, 0x50,
	0xef, 0x8b, 0xd9, 0x60, 0x2b, 0x80, 0x1d, 0xef, 0x42, 0x9c, 0xff, 0xd8, 0x1f, 0x55, 0xed, 0x87,
	0x0d, 0x24, 0x51, 0x2e, 0xab, 0x79, 0x66, 0xe6, 0x99, 0xdd, 0x99, 0x79, 0x6c, 0xf0, 0xf2, 0x2c,
	0x1a, 0xd0, 0x2c, 0xe1, 0x61, 0x96, 0xa7, 0x22, 0x45, 0x76, 0x9e, 0x45, 0x7e, 0x37, 0x4a, 0x73,
	0x36, 0x10, 0xc5, 0x40, 0x14, 0xda, 0xed, 0x1f, 0x2a, 0xcf, 0x74, 0x91, 0x46, 0xb7, 0xfa, 0xdc,
	0xf3, 0xb3, 0x0d, 0x5b, 0x09, 0x7d, 0x6a, 0x3f, 0x76, 0xe0, 0xe0, 0x4f, 0x9a, 0xcc, 0x08, 0xbb,
	0xc3, 0xef, 0xe0, 0xe0, 0x92, 0xf2, 0x98, 0xb0, 0x3b, 0x84, 0xa0, 0x11, 0x53, 0x1e, 0xf7, 0xea,
	0xfd, 0x7a, 0xe0, 0x12, 0x65, 0xe3, 0x5f, 0xe0, 0x9d, 0xc9, 0x82, 0x67, 0x0f, 0x2f, 0x64, 0x21,
	0x1f, 0x5a, 0x51, 0xba, 0xcc, 0x16, 0x4c, 0xb0, 0x9e, 0xd5, 0xaf, 0x07, 0x2d, 0x52, 0x61, 0xfc,
	0x13, 0x3a, 0xa6, 0xc2, 0xf5, 0x7a, 0x29, 0x0b, 0x74, 0xc1, 0x5e, 0xad, 0x97, 0x8a, 0x6f, 0x13,
	0x69, 0xbe, 0x48, 0x3f, 0x82, 0xce, 0x88, 0x89, 0x33, 0xba, 0xa0, 0xab, 0x88, 0x49, 0xba, 0x07,
	0xd6, 0xd5, 0xb9, 0x62, 0x3b, 0xc4, 0xba, 0x3a, 0xc7, 0x47, 0xd0, 0x1e, 0x31, 0x31, 0x16, 0x54,
	0x30, 0x53, 0xfd, 0x96, 0x3d, 0x98, 0xb8, 0x34, 0xf1, 0x7b, 0x68, 0x11, 0x7a, 0x3f, 0x29, 0xcc,
	0xe3, 0x67, 0x54, 0xd0, 0xf2, 0xf1, 0xd2, 0xc6, 0x3f, 0xc0, 0x1d, 0xaf, 0xa7, 0x3c, 0xca, 0x93,
	0xa9, 0xaa, 0xf0, 0x09, 0x9a, 0x22, 0xcd, 0x92, 0x88, 0xf7, 0xea, 0x7d, 0x3b, 0xf0, 0x86, 0x28,
	0xd4, 0xa3, 0xbb, 0x50, 0xe7, 0x44, 0x86, 0x88, 0xc9, 0xc0, 0xc7, 0xe0, 0x5c, 0xb2, 0x64, 0x1e,
	0x0b, 0xc2, 0x38, 0x3a, 0x84, 0x66, 0xac, 0x80, 0xe9, 0xcd, 0x20, 0xfc, 0x71, 0xbf, 0x05, 0x8e,
	0x7a, 0x70, 0x30, 0xd5, 0xc8, 0x64, 0x96, 0x10, 0x1f, 0xef, 0x36, 0xc3, 0xd1, 0x6b, 0xf8, 0x6f,
	0x43, 0x17, 0x6b, 0x66, 0xda, 0xd1, 0x00, 0x63, 0x70, 0xc7, 0x6c, 0x35, 0x33, 0x4d, 0xf1, 0x67,
	0xf7, 0xe6, 0x43, 0x73, 0x44, 0xb9, 0x8c, 0x76, 0xc1, 0x9e, 0x53, 0xae, 0x82, 0x0d, 0x22, 0x4d,
	0x7c, 0x03, 0x8e, 0xda, 0xc8, 0xd5, 0xea, 0x26, 0x45, 0x1f, 0xa0, 0x11, 0x33, 0x3a, 0x53, 0xf1,
	0xf6, 0xb0, 0x1b, 0x6a, 0xf9, 0xa8, 0xf8, 0x25, 0xa3, 0x33, 0xa2, 0xa2, 0xe8, 0x0d, 0xd8, 0xa2,
	0xe0, 0x3d, 0xab, 0x6f, 0x07, 0xed, 0xa1, 0x13, 0x8a, 0x22, 0x9c, 0x14, 0x84, 0xde, 0x13, 0xe9,
	0x95, 0x7d, 0x8b, 0x42, 0xbd, 0xc0, 0xee, 0xdb, 0x81, 0x4b, 0x0c, 0xc2, 0x9f, 0xf7, 0x06, 0xcb,
	0xd1, 0x5b, 0xb0, 0xd8, 0xc6, 0x5c, 0xe4, 0xee, 0x0e, 0x95, 0x58, 0x6c, 0x33, 0xfc, 0x6b, 0x43,
	0xe3, 0x77, 0x96, 0x70, 0x74, 0x02, 0xce, 0x88, 0x09, 0x3d, 0x56, 0xe4, 0x86, 0x79, 0x16, 0x85,
	0x46, 0xac, 0xbe, 0xa7, 0x50, 0x35, 0x71, 0x5c, 0x43, 0x81, 0x1a, 0xd8, 0xa4, 0xd0, 0xfa, 0x34,
	0xe9, 0x46, 0xaa, 0xfe, 0xf6, 0xa1, 0xb8, 0x86, 0xbe, 0x83, 0x27, 0xb7, 0xb0, 0x15, 0x33, 0x7a,
	0xa5, 0x92, 0xf7, 0xe5, 0x6d, 0xae, 0xa8, 0xe6, 0x83, 0x6b, 0xe8, 0x2b, 0x74, 0xe6, 0x15, 0xf1,
	0x7a, 0xbd, 0x44, 0x68, 0x97, 0xa7, 0x45, 0xfd, 0x0c, 0xed, 0x1b, 0xc0, 0x76, 0xeb, 0x86, 0xb3,
	0xa7, 0x64, 0xff, 0xa9, 0x4f, 0x76, 0x74, 0x0a, 0xad, 0x52, 0x02, 0xa8, 0x5b, 0x66, 0x94, 0xf2,
	0xf6, 0x1f, 0x7b, 0x24, 0x63, 0x00, 0x4e, 0xa5, 0x07, 0xd4, 0x51, 0x09, 0xa5, 0xe0, 0xfd, 0xff,
	0x15, 0xdc, 0x95, 0x0b, 0xae, 0xa1, 0x13, 0x68, 0x5f, 0x70, 0x91, 0x2c, 0xa9, 0x60, 0x23, 0xca,
	0x1f, 0x53, 0xda, 0xfa, 0x0a, 0xa5, 0x1e, 0xd5, 0xbe, 0x53, 0x6d, 0x11, 0x99, 0x72, 0x3b, 0x9f,
	0x8b, 0xff, 0xc4, 0xc5, 0x71, 0xed, 0xb4, 0x3e, 0x6d, 0xaa, 0x3f, 0xcd, 0x97, 0x7f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xe0, 0x89, 0xe4, 0x30, 0xc2, 0x04, 0x00, 0x00,
}
