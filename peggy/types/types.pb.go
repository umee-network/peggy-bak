// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: injective/peggy/v1/types.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// BridgeValidator represents a validator's ETH address and its power
type BridgeValidator struct {
	Power           uint64 `protobuf:"varint,1,opt,name=power,proto3" json:"power,omitempty"`
	EthereumAddress string `protobuf:"bytes,2,opt,name=ethereum_address,json=ethereumAddress,proto3" json:"ethereum_address,omitempty"`
}

func (m *BridgeValidator) Reset()         { *m = BridgeValidator{} }
func (m *BridgeValidator) String() string { return proto.CompactTextString(m) }
func (*BridgeValidator) ProtoMessage()    {}
func (*BridgeValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_b641943ad411b503, []int{0}
}
func (m *BridgeValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BridgeValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BridgeValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BridgeValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BridgeValidator.Merge(m, src)
}
func (m *BridgeValidator) XXX_Size() int {
	return m.Size()
}
func (m *BridgeValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_BridgeValidator.DiscardUnknown(m)
}

var xxx_messageInfo_BridgeValidator proto.InternalMessageInfo

func (m *BridgeValidator) GetPower() uint64 {
	if m != nil {
		return m.Power
	}
	return 0
}

func (m *BridgeValidator) GetEthereumAddress() string {
	if m != nil {
		return m.EthereumAddress
	}
	return ""
}

// Valset is the Ethereum Bridge Multsig Set, each peggy validator also
// maintains an ETH key to sign messages, these are used to check signatures on
// ETH because of the significant gas savings
type Valset struct {
	Nonce        uint64                                 `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Members      []*BridgeValidator                     `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	Height       uint64                                 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	RewardAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=reward_amount,json=rewardAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"reward_amount"`
	// the reward token in it's Ethereum hex address representation
	RewardToken string `protobuf:"bytes,5,opt,name=reward_token,json=rewardToken,proto3" json:"reward_token,omitempty"`
}

func (m *Valset) Reset()         { *m = Valset{} }
func (m *Valset) String() string { return proto.CompactTextString(m) }
func (*Valset) ProtoMessage()    {}
func (*Valset) Descriptor() ([]byte, []int) {
	return fileDescriptor_b641943ad411b503, []int{1}
}
func (m *Valset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Valset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Valset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Valset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Valset.Merge(m, src)
}
func (m *Valset) XXX_Size() int {
	return m.Size()
}
func (m *Valset) XXX_DiscardUnknown() {
	xxx_messageInfo_Valset.DiscardUnknown(m)
}

var xxx_messageInfo_Valset proto.InternalMessageInfo

func (m *Valset) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Valset) GetMembers() []*BridgeValidator {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *Valset) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Valset) GetRewardToken() string {
	if m != nil {
		return m.RewardToken
	}
	return ""
}

// LastObservedEthereumBlockHeight stores the last observed
// Ethereum block height along with the Cosmos block height that
// it was observed at. These two numbers can be used to project
// outward and always produce batches with timeouts in the future
// even if no Ethereum block height has been relayed for a long time
type LastObservedEthereumBlockHeight struct {
	CosmosBlockHeight   uint64 `protobuf:"varint,1,opt,name=cosmos_block_height,json=cosmosBlockHeight,proto3" json:"cosmos_block_height,omitempty"`
	EthereumBlockHeight uint64 `protobuf:"varint,2,opt,name=ethereum_block_height,json=ethereumBlockHeight,proto3" json:"ethereum_block_height,omitempty"`
}

func (m *LastObservedEthereumBlockHeight) Reset()         { *m = LastObservedEthereumBlockHeight{} }
func (m *LastObservedEthereumBlockHeight) String() string { return proto.CompactTextString(m) }
func (*LastObservedEthereumBlockHeight) ProtoMessage()    {}
func (*LastObservedEthereumBlockHeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_b641943ad411b503, []int{2}
}
func (m *LastObservedEthereumBlockHeight) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LastObservedEthereumBlockHeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LastObservedEthereumBlockHeight.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LastObservedEthereumBlockHeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LastObservedEthereumBlockHeight.Merge(m, src)
}
func (m *LastObservedEthereumBlockHeight) XXX_Size() int {
	return m.Size()
}
func (m *LastObservedEthereumBlockHeight) XXX_DiscardUnknown() {
	xxx_messageInfo_LastObservedEthereumBlockHeight.DiscardUnknown(m)
}

var xxx_messageInfo_LastObservedEthereumBlockHeight proto.InternalMessageInfo

func (m *LastObservedEthereumBlockHeight) GetCosmosBlockHeight() uint64 {
	if m != nil {
		return m.CosmosBlockHeight
	}
	return 0
}

func (m *LastObservedEthereumBlockHeight) GetEthereumBlockHeight() uint64 {
	if m != nil {
		return m.EthereumBlockHeight
	}
	return 0
}

// LastClaimEvent stores last claim event details of validator.
type LastClaimEvent struct {
	EthereumEventNonce  uint64 `protobuf:"varint,1,opt,name=ethereum_event_nonce,json=ethereumEventNonce,proto3" json:"ethereum_event_nonce,omitempty"`
	EthereumEventHeight uint64 `protobuf:"varint,2,opt,name=ethereum_event_height,json=ethereumEventHeight,proto3" json:"ethereum_event_height,omitempty"`
}

func (m *LastClaimEvent) Reset()         { *m = LastClaimEvent{} }
func (m *LastClaimEvent) String() string { return proto.CompactTextString(m) }
func (*LastClaimEvent) ProtoMessage()    {}
func (*LastClaimEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_b641943ad411b503, []int{3}
}
func (m *LastClaimEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LastClaimEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LastClaimEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LastClaimEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LastClaimEvent.Merge(m, src)
}
func (m *LastClaimEvent) XXX_Size() int {
	return m.Size()
}
func (m *LastClaimEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_LastClaimEvent.DiscardUnknown(m)
}

var xxx_messageInfo_LastClaimEvent proto.InternalMessageInfo

func (m *LastClaimEvent) GetEthereumEventNonce() uint64 {
	if m != nil {
		return m.EthereumEventNonce
	}
	return 0
}

func (m *LastClaimEvent) GetEthereumEventHeight() uint64 {
	if m != nil {
		return m.EthereumEventHeight
	}
	return 0
}

// This records the relationship between an ERC20 token and the denom
// of the corresponding Cosmos originated asset
type ERC20ToDenom struct {
	Erc20 string `protobuf:"bytes,1,opt,name=erc20,proto3" json:"erc20,omitempty"`
	Denom string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *ERC20ToDenom) Reset()         { *m = ERC20ToDenom{} }
func (m *ERC20ToDenom) String() string { return proto.CompactTextString(m) }
func (*ERC20ToDenom) ProtoMessage()    {}
func (*ERC20ToDenom) Descriptor() ([]byte, []int) {
	return fileDescriptor_b641943ad411b503, []int{4}
}
func (m *ERC20ToDenom) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ERC20ToDenom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ERC20ToDenom.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ERC20ToDenom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ERC20ToDenom.Merge(m, src)
}
func (m *ERC20ToDenom) XXX_Size() int {
	return m.Size()
}
func (m *ERC20ToDenom) XXX_DiscardUnknown() {
	xxx_messageInfo_ERC20ToDenom.DiscardUnknown(m)
}

var xxx_messageInfo_ERC20ToDenom proto.InternalMessageInfo

func (m *ERC20ToDenom) GetErc20() string {
	if m != nil {
		return m.Erc20
	}
	return ""
}

func (m *ERC20ToDenom) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func init() {
	proto.RegisterType((*BridgeValidator)(nil), "injective.peggy.v1.BridgeValidator")
	proto.RegisterType((*Valset)(nil), "injective.peggy.v1.Valset")
	proto.RegisterType((*LastObservedEthereumBlockHeight)(nil), "injective.peggy.v1.LastObservedEthereumBlockHeight")
	proto.RegisterType((*LastClaimEvent)(nil), "injective.peggy.v1.LastClaimEvent")
	proto.RegisterType((*ERC20ToDenom)(nil), "injective.peggy.v1.ERC20ToDenom")
}

func init() { proto.RegisterFile("injective/peggy/v1/types.proto", fileDescriptor_b641943ad411b503) }

var fileDescriptor_b641943ad411b503 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x6e, 0xf6, 0xa3, 0x68, 0xde, 0x60, 0xe0, 0x15, 0x54, 0xed, 0x90, 0x8e, 0x22, 0xa1, 0x72,
	0x58, 0xdc, 0x96, 0x1b, 0x12, 0x87, 0x75, 0x54, 0x62, 0x62, 0x02, 0x29, 0x4c, 0x3b, 0x70, 0x89,
	0x9c, 0xe4, 0x29, 0x09, 0x4d, 0xec, 0xca, 0x76, 0x33, 0xed, 0x0f, 0xe0, 0xce, 0x9f, 0xb5, 0xe3,
	0x8e, 0x88, 0xc3, 0x84, 0xda, 0x33, 0xff, 0x03, 0xb2, 0x9d, 0x54, 0x81, 0x71, 0x8a, 0xbf, 0xef,
	0xf9, 0x7d, 0xdf, 0xf3, 0x7b, 0x2f, 0xc8, 0xcd, 0xd8, 0x57, 0x88, 0x54, 0x56, 0x02, 0x99, 0x43,
	0x92, 0x5c, 0x93, 0x72, 0x44, 0xd4, 0xf5, 0x1c, 0xa4, 0x37, 0x17, 0x5c, 0x71, 0x8c, 0xd7, 0x71,
	0xcf, 0xc4, 0xbd, 0x72, 0x74, 0xe8, 0x46, 0x5c, 0x16, 0x5c, 0x92, 0x90, 0x4a, 0x20, 0xe5, 0x28,
	0x04, 0x45, 0x47, 0x24, 0xe2, 0x19, 0xb3, 0x39, 0x87, 0x9d, 0x84, 0x27, 0xdc, 0x1c, 0x89, 0x3e,
	0x59, 0xb6, 0xef, 0xa3, 0xfd, 0x89, 0xc8, 0xe2, 0x04, 0x2e, 0x69, 0x9e, 0xc5, 0x54, 0x71, 0x81,
	0x3b, 0x68, 0x7b, 0xce, 0xaf, 0x40, 0x74, 0x9d, 0x23, 0x67, 0xb0, 0xe5, 0x5b, 0x80, 0x5f, 0xa1,
	0xc7, 0xa0, 0x52, 0x10, 0xb0, 0x28, 0x02, 0x1a, 0xc7, 0x02, 0xa4, 0xec, 0x6e, 0x1c, 0x39, 0x83,
	0x1d, 0x7f, 0xbf, 0xe6, 0x4f, 0x2c, 0xdd, 0xff, 0xed, 0xa0, 0xf6, 0x25, 0xcd, 0x25, 0x28, 0xad,
	0xc5, 0x38, 0x8b, 0xa0, 0xd6, 0x32, 0x00, 0xbf, 0x45, 0x0f, 0x0a, 0x28, 0x42, 0x10, 0x5a, 0x62,
	0x73, 0xb0, 0x3b, 0x7e, 0xe1, 0xdd, 0x7f, 0x90, 0xf7, 0x4f, 0x5d, 0x7e, 0x9d, 0x83, 0x9f, 0xa1,
	0x76, 0x0a, 0x59, 0x92, 0xaa, 0xee, 0xa6, 0x51, 0xad, 0x10, 0xfe, 0x8c, 0x1e, 0x0a, 0xb8, 0xa2,
	0x22, 0x0e, 0x68, 0xc1, 0x17, 0x4c, 0x75, 0xb7, 0x74, 0x7d, 0x13, 0xef, 0xe6, 0xae, 0xd7, 0xfa,
	0x79, 0xd7, 0x7b, 0x99, 0x64, 0x2a, 0x5d, 0x84, 0x5e, 0xc4, 0x0b, 0x52, 0xf5, 0xca, 0x7e, 0x8e,
	0x65, 0x3c, 0xab, 0xda, 0x7b, 0xc6, 0x94, 0xbf, 0x67, 0x45, 0x4e, 0x8c, 0x06, 0x7e, 0x8e, 0x2a,
	0x1c, 0x28, 0x3e, 0x03, 0xd6, 0xdd, 0x36, 0x6f, 0xde, 0xb5, 0xdc, 0x85, 0xa6, 0xfa, 0xdf, 0x1c,
	0xd4, 0x3b, 0xa7, 0x52, 0x7d, 0x0a, 0x25, 0x88, 0x12, 0xe2, 0x69, 0xd5, 0x8f, 0x49, 0xce, 0xa3,
	0xd9, 0x7b, 0x5b, 0x9b, 0x87, 0x0e, 0xac, 0x59, 0x10, 0x6a, 0x36, 0xa8, 0x1e, 0x60, 0xdb, 0xf2,
	0xc4, 0x86, 0x9a, 0xf7, 0xc7, 0xe8, 0xe9, 0xba, 0xdd, 0x7f, 0x65, 0x6c, 0x98, 0x8c, 0x03, 0xb8,
	0xef, 0xd1, 0x2f, 0xd1, 0x23, 0x5d, 0xc6, 0x69, 0x4e, 0xb3, 0x62, 0x5a, 0x02, 0x53, 0x78, 0x88,
	0x3a, 0x6b, 0x15, 0xd0, 0x4c, 0xd0, 0x9c, 0x06, 0xae, 0x63, 0xe6, 0xf2, 0x47, 0x33, 0x9a, 0xa6,
	0xaf, 0xcd, 0xf8, 0xbf, 0xaf, 0x49, 0xa9, 0x7c, 0xdf, 0xa0, 0xbd, 0xa9, 0x7f, 0x3a, 0x1e, 0x5e,
	0xf0, 0x77, 0xc0, 0x78, 0xa1, 0x87, 0x0e, 0x22, 0x1a, 0x0f, 0x8d, 0xcd, 0x8e, 0x6f, 0x81, 0x66,
	0x63, 0x1d, 0xae, 0xb6, 0xc6, 0x82, 0x09, 0xdc, 0x2c, 0x5d, 0xe7, 0x76, 0xe9, 0x3a, 0xbf, 0x96,
	0xae, 0xf3, 0x7d, 0xe5, 0xb6, 0x6e, 0x57, 0x6e, 0xeb, 0xc7, 0xca, 0x6d, 0x7d, 0xf9, 0xd0, 0x18,
	0xd7, 0x59, 0xbd, 0x1d, 0xe7, 0x34, 0x94, 0x64, 0xbd, 0x2b, 0xc7, 0x11, 0x17, 0xd0, 0x84, 0x29,
	0xcd, 0x18, 0x29, 0x78, 0xbc, 0xc8, 0x41, 0x56, 0x7f, 0x8e, 0x99, 0x6b, 0xd8, 0x36, 0xdb, 0xfe,
	0xfa, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xdb, 0x16, 0x10, 0x59, 0x03, 0x00, 0x00,
}

func (m *BridgeValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BridgeValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BridgeValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EthereumAddress) > 0 {
		i -= len(m.EthereumAddress)
		copy(dAtA[i:], m.EthereumAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.EthereumAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.Power != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Valset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Valset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Valset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RewardToken) > 0 {
		i -= len(m.RewardToken)
		copy(dAtA[i:], m.RewardToken)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.RewardToken)))
		i--
		dAtA[i] = 0x2a
	}
	{
		size := m.RewardAmount.Size()
		i -= size
		if _, err := m.RewardAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Members[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Nonce != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LastObservedEthereumBlockHeight) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LastObservedEthereumBlockHeight) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LastObservedEthereumBlockHeight) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EthereumBlockHeight != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.EthereumBlockHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.CosmosBlockHeight != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.CosmosBlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LastClaimEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LastClaimEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LastClaimEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EthereumEventHeight != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.EthereumEventHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.EthereumEventNonce != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.EthereumEventNonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ERC20ToDenom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ERC20ToDenom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ERC20ToDenom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Erc20) > 0 {
		i -= len(m.Erc20)
		copy(dAtA[i:], m.Erc20)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Erc20)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BridgeValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Power != 0 {
		n += 1 + sovTypes(uint64(m.Power))
	}
	l = len(m.EthereumAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Valset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nonce != 0 {
		n += 1 + sovTypes(uint64(m.Nonce))
	}
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.Height != 0 {
		n += 1 + sovTypes(uint64(m.Height))
	}
	l = m.RewardAmount.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.RewardToken)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *LastObservedEthereumBlockHeight) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CosmosBlockHeight != 0 {
		n += 1 + sovTypes(uint64(m.CosmosBlockHeight))
	}
	if m.EthereumBlockHeight != 0 {
		n += 1 + sovTypes(uint64(m.EthereumBlockHeight))
	}
	return n
}

func (m *LastClaimEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EthereumEventNonce != 0 {
		n += 1 + sovTypes(uint64(m.EthereumEventNonce))
	}
	if m.EthereumEventHeight != 0 {
		n += 1 + sovTypes(uint64(m.EthereumEventHeight))
	}
	return n
}

func (m *ERC20ToDenom) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Erc20)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BridgeValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: BridgeValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BridgeValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthereumAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthereumAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Valset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Valset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Valset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, &BridgeValidator{})
			if err := m.Members[len(m.Members)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *LastObservedEthereumBlockHeight) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: LastObservedEthereumBlockHeight: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LastObservedEthereumBlockHeight: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CosmosBlockHeight", wireType)
			}
			m.CosmosBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CosmosBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthereumBlockHeight", wireType)
			}
			m.EthereumBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EthereumBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *LastClaimEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: LastClaimEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LastClaimEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthereumEventNonce", wireType)
			}
			m.EthereumEventNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EthereumEventNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthereumEventHeight", wireType)
			}
			m.EthereumEventHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EthereumEventHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ERC20ToDenom) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ERC20ToDenom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ERC20ToDenom: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Erc20", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Erc20 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)