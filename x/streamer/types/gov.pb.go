// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymension/streamer/gov.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
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

type CreateStreamProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// distribute_to show which lock the stream should distribute to by time
	DistributeTo string `protobuf:"bytes,3,opt,name=distribute_to,json=distributeTo,proto3" json:"distribute_to,omitempty"`
	// coins are coin(s) to be distributed by the stream
	Coins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
	// start_time is the distribution start time
	StartTime            time.Time `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"timestamp"`
	DistrEpochIdentifier string    `protobuf:"bytes,6,opt,name=distr_epoch_identifier,json=distrEpochIdentifier,proto3" json:"distr_epoch_identifier,omitempty" yaml:"distr_epoch_identifier"`
	// num_epochs_paid_over is the number of epochs distribution will be completed
	// over
	NumEpochsPaidOver uint64 `protobuf:"varint,7,opt,name=num_epochs_paid_over,json=numEpochsPaidOver,proto3" json:"num_epochs_paid_over,omitempty"`
}

func (m *CreateStreamProposal) Reset()      { *m = CreateStreamProposal{} }
func (*CreateStreamProposal) ProtoMessage() {}
func (*CreateStreamProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_6da8318c7fb35093, []int{0}
}
func (m *CreateStreamProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateStreamProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateStreamProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateStreamProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateStreamProposal.Merge(m, src)
}
func (m *CreateStreamProposal) XXX_Size() int {
	return m.Size()
}
func (m *CreateStreamProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateStreamProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CreateStreamProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateStreamProposal)(nil), "dymensionxyz.dymension.streamer.CreateStreamProposal")
}

func init() { proto.RegisterFile("dymension/streamer/gov.proto", fileDescriptor_6da8318c7fb35093) }

var fileDescriptor_6da8318c7fb35093 = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x3d, 0x6f, 0xd3, 0x40,
	0x18, 0xb6, 0x49, 0x53, 0xd4, 0x4b, 0x91, 0xc0, 0xb2, 0x90, 0x89, 0x8a, 0xcf, 0x84, 0xc5, 0x0b,
	0x77, 0x34, 0x6c, 0x1d, 0x53, 0x31, 0x20, 0x21, 0x51, 0x99, 0x4a, 0x95, 0x58, 0xac, 0xb3, 0x7d,
	0x75, 0x4f, 0xc4, 0x7e, 0xad, 0xbb, 0x73, 0xd4, 0xf0, 0x0b, 0x18, 0x3b, 0x32, 0x66, 0xe6, 0x8f,
	0xd0, 0xb1, 0x23, 0x53, 0x8a, 0x92, 0x85, 0xb9, 0xbf, 0x00, 0xf9, 0x9c, 0xaf, 0x81, 0xc9, 0x7e,
	0x3e, 0xde, 0xf7, 0xee, 0x79, 0x74, 0xe8, 0x28, 0x9b, 0x16, 0xbc, 0x54, 0x02, 0x4a, 0xaa, 0xb4,
	0xe4, 0xac, 0xe0, 0x92, 0xe6, 0x30, 0x21, 0x95, 0x04, 0x0d, 0x0e, 0xde, 0xa8, 0xd7, 0xd3, 0x6f,
	0x64, 0x03, 0xc8, 0xda, 0xda, 0x77, 0x73, 0xc8, 0xc1, 0x78, 0x69, 0xf3, 0xd7, 0x8e, 0xf5, 0xfd,
	0x14, 0x54, 0x01, 0x8a, 0x26, 0x4c, 0x71, 0x3a, 0x39, 0x4e, 0xb8, 0x66, 0xc7, 0x34, 0x05, 0x51,
	0xae, 0x74, 0x9c, 0x03, 0xe4, 0x63, 0x4e, 0x0d, 0x4a, 0xea, 0x4b, 0xaa, 0x45, 0xc1, 0x95, 0x66,
	0x45, 0xd5, 0x1a, 0x06, 0xbf, 0x3a, 0xc8, 0x3d, 0x95, 0x9c, 0x69, 0xfe, 0xd9, 0x9c, 0x74, 0x26,
	0xa1, 0x02, 0xc5, 0xc6, 0x8e, 0x8b, 0xba, 0x5a, 0xe8, 0x31, 0xf7, 0xec, 0xc0, 0x0e, 0x0f, 0xa2,
	0x16, 0x38, 0x01, 0xea, 0x65, 0x5c, 0xa5, 0x52, 0x54, 0x5a, 0x40, 0xe9, 0x3d, 0x32, 0xda, 0x2e,
	0xe5, 0xbc, 0x46, 0x4f, 0x32, 0xa1, 0xb4, 0x14, 0x49, 0xad, 0x79, 0xac, 0xc1, 0xeb, 0x18, 0xcf,
	0xe1, 0x96, 0x3c, 0x07, 0x87, 0xa1, 0x6e, 0x73, 0x49, 0xe5, 0xed, 0x05, 0x9d, 0xb0, 0x37, 0x7c,
	0x41, 0xda, 0x18, 0xa4, 0x89, 0x41, 0x56, 0x31, 0xc8, 0x29, 0x88, 0x72, 0xf4, 0xf6, 0x76, 0x8e,
	0xad, 0x9f, 0xf7, 0x38, 0xcc, 0x85, 0xbe, 0xaa, 0x13, 0x92, 0x42, 0x41, 0x57, 0x99, 0xdb, 0xcf,
	0x1b, 0x95, 0x7d, 0xa5, 0x7a, 0x5a, 0x71, 0x65, 0x06, 0x54, 0xd4, 0x6e, 0x76, 0x2e, 0x10, 0x52,
	0x9a, 0x49, 0x1d, 0x37, 0x89, 0xbd, 0x6e, 0x60, 0x87, 0xbd, 0x61, 0x9f, 0xb4, 0x75, 0x90, 0x75,
	0x1d, 0xe4, 0x7c, 0x5d, 0xc7, 0xe8, 0xa8, 0x39, 0xe8, 0x61, 0x8e, 0x9f, 0x4e, 0x59, 0x31, 0x3e,
	0x19, 0x6c, 0x7a, 0x1a, 0xdc, 0xdc, 0x63, 0x3b, 0x3a, 0x30, 0xbb, 0x1a, 0xb7, 0x73, 0x81, 0x9e,
	0x9b, 0x2c, 0x31, 0xaf, 0x20, 0xbd, 0x8a, 0x45, 0xc6, 0x4b, 0x2d, 0x2e, 0x05, 0x97, 0xde, 0x7e,
	0x93, 0x74, 0xf4, 0xea, 0x61, 0x8e, 0x5f, 0xb6, 0x4b, 0xfe, 0xef, 0x1b, 0x44, 0xae, 0x11, 0xde,
	0x37, 0xfc, 0x87, 0x0d, 0xed, 0x50, 0xe4, 0x96, 0x75, 0xd1, 0xda, 0x55, 0x5c, 0x31, 0x91, 0xc5,
	0x30, 0xe1, 0xd2, 0x7b, 0x1c, 0xd8, 0xe1, 0x5e, 0xf4, 0xac, 0xac, 0x0b, 0x33, 0xa1, 0xce, 0x98,
	0xc8, 0x3e, 0x4d, 0xb8, 0x3c, 0x39, 0xfc, 0x3e, 0xc3, 0xd6, 0x8f, 0x19, 0xb6, 0xfe, 0xce, 0xb0,
	0x3d, 0xfa, 0x78, 0xbb, 0xf0, 0xed, 0xbb, 0x85, 0x6f, 0xff, 0x59, 0xf8, 0xf6, 0xcd, 0xd2, 0xb7,
	0xee, 0x96, 0xbe, 0xf5, 0x7b, 0xe9, 0x5b, 0x5f, 0x86, 0x3b, 0xdd, 0xed, 0x3e, 0xb3, 0x2d, 0xa0,
	0xd7, 0xdb, 0x37, 0x69, 0xba, 0x4c, 0xf6, 0x4d, 0x45, 0xef, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff,
	0xa6, 0x47, 0x4b, 0xda, 0xb6, 0x02, 0x00, 0x00,
}

func (this *CreateStreamProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CreateStreamProposal)
	if !ok {
		that2, ok := that.(CreateStreamProposal)
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
	if this.DistributeTo != that1.DistributeTo {
		return false
	}
	if len(this.Coins) != len(that1.Coins) {
		return false
	}
	for i := range this.Coins {
		if !this.Coins[i].Equal(&that1.Coins[i]) {
			return false
		}
	}
	if !this.StartTime.Equal(that1.StartTime) {
		return false
	}
	if this.DistrEpochIdentifier != that1.DistrEpochIdentifier {
		return false
	}
	if this.NumEpochsPaidOver != that1.NumEpochsPaidOver {
		return false
	}
	return true
}
func (m *CreateStreamProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateStreamProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateStreamProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NumEpochsPaidOver != 0 {
		i = encodeVarintGov(dAtA, i, uint64(m.NumEpochsPaidOver))
		i--
		dAtA[i] = 0x38
	}
	if len(m.DistrEpochIdentifier) > 0 {
		i -= len(m.DistrEpochIdentifier)
		copy(dAtA[i:], m.DistrEpochIdentifier)
		i = encodeVarintGov(dAtA, i, uint64(len(m.DistrEpochIdentifier)))
		i--
		dAtA[i] = 0x32
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGov(dAtA, i, uint64(n1))
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
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.DistributeTo) > 0 {
		i -= len(m.DistributeTo)
		copy(dAtA[i:], m.DistributeTo)
		i = encodeVarintGov(dAtA, i, uint64(len(m.DistributeTo)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CreateStreamProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.DistributeTo)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovGov(uint64(l))
	l = len(m.DistrEpochIdentifier)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if m.NumEpochsPaidOver != 0 {
		n += 1 + sovGov(uint64(m.NumEpochsPaidOver))
	}
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreateStreamProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: CreateStreamProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateStreamProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
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
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributeTo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DistributeTo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
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
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistrEpochIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DistrEpochIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumEpochsPaidOver", wireType)
			}
			m.NumEpochsPaidOver = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
