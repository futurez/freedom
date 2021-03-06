// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fmessage/eproto/proto_msg.eproto

package pb3

import (
	fmt "fmt"
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

type ProtoMessage struct {
	MsgId uint32 `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
	Code  int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Body  []byte `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *ProtoMessage) Reset()         { *m = ProtoMessage{} }
func (m *ProtoMessage) String() string { return proto.CompactTextString(m) }
func (*ProtoMessage) ProtoMessage()    {}
func (*ProtoMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4ba7cc524ddf831, []int{0}
}
func (m *ProtoMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProtoMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProtoMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProtoMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoMessage.Merge(m, src)
}
func (m *ProtoMessage) XXX_Size() int {
	return m.Size()
}
func (m *ProtoMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoMessage proto.InternalMessageInfo

func (m *ProtoMessage) GetMsgId() uint32 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *ProtoMessage) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ProtoMessage) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*ProtoMessage)(nil), "pb3.ProtoMessage")
}

func init() { proto.RegisterFile("fmessage/eproto/proto_msg.eproto", fileDescriptor_b4ba7cc524ddf831) }

var fileDescriptor_b4ba7cc524ddf831 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcb, 0x4d, 0x2d,
	0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x87, 0x90, 0xf1, 0xb9, 0xc5, 0xe9,
	0x7a, 0x60, 0x96, 0x10, 0x73, 0x41, 0x92, 0xb1, 0x92, 0x0f, 0x17, 0x4f, 0x00, 0x88, 0xe7, 0x0b,
	0x51, 0x2a, 0x24, 0xc2, 0xc5, 0x9a, 0x5b, 0x9c, 0xee, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0x1b, 0x04, 0xe1, 0x08, 0x09, 0x71, 0xb1, 0x24, 0xe7, 0xa7, 0xa4, 0x4a, 0x30, 0x29, 0x30, 0x6a,
	0xb0, 0x06, 0x81, 0xd9, 0x20, 0xb1, 0xa4, 0xfc, 0x94, 0x4a, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x9e,
	0x20, 0x30, 0xdb, 0x49, 0xe2, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92,
	0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0x92, 0xd8,
	0xc0, 0x76, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xf7, 0x05, 0x51, 0x95, 0x00, 0x00,
	0x00,
}

func (m *ProtoMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProtoMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProtoMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintProtoMsg(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Code != 0 {
		i = encodeVarintProtoMsg(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x10
	}
	if m.MsgId != 0 {
		i = encodeVarintProtoMsg(dAtA, i, uint64(m.MsgId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProtoMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovProtoMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProtoMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MsgId != 0 {
		n += 1 + sovProtoMsg(uint64(m.MsgId))
	}
	if m.Code != 0 {
		n += 1 + sovProtoMsg(uint64(m.Code))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovProtoMsg(uint64(l))
	}
	return n
}

func sovProtoMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProtoMsg(x uint64) (n int) {
	return sovProtoMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProtoMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtoMsg
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
			return fmt.Errorf("eproto: ProtoMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("eproto: ProtoMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("eproto: wrong wireType = %d for field MsgId", wireType)
			}
			m.MsgId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtoMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MsgId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("eproto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtoMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("eproto: wrong wireType = %d for field Body", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtoMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProtoMsg
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProtoMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = append(m.Body[:0], dAtA[iNdEx:postIndex]...)
			if m.Body == nil {
				m.Body = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtoMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProtoMsg
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
func skipProtoMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProtoMsg
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
					return 0, ErrIntOverflowProtoMsg
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
					return 0, ErrIntOverflowProtoMsg
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
				return 0, ErrInvalidLengthProtoMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProtoMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("eproto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProtoMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProtoMsg        = fmt.Errorf("eproto: negative length found during unmarshaling")
	ErrIntOverflowProtoMsg          = fmt.Errorf("eproto: integer overflow")
	ErrUnexpectedEndOfGroupProtoMsg = fmt.Errorf("eproto: unexpected end of group")
)
