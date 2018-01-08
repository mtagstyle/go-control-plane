// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/filter/http/health_check.proto

package envoy_api_v2_filter_http

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"
import google_protobuf1 "github.com/gogo/protobuf/types"
import _ "github.com/lyft/protoc-gen-validate/validate"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HealthCheck struct {
	// Specifies whether the filter operates in pass through mode or not.
	PassThroughMode *google_protobuf1.BoolValue `protobuf:"bytes,1,opt,name=pass_through_mode,json=passThroughMode" json:"pass_through_mode,omitempty"`
	// Specifies the incoming HTTP endpoint that should be considered the
	// health check endpoint. For example */healthcheck*.
	Endpoint string `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// If operating in pass through mode, the amount of time in milliseconds
	// that the filter should cache the upstream response.
	CacheTime *google_protobuf.Duration `protobuf:"bytes,3,opt,name=cache_time,json=cacheTime" json:"cache_time,omitempty"`
}

func (m *HealthCheck) Reset()                    { *m = HealthCheck{} }
func (m *HealthCheck) String() string            { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()               {}
func (*HealthCheck) Descriptor() ([]byte, []int) { return fileDescriptorHealthCheck, []int{0} }

func (m *HealthCheck) GetPassThroughMode() *google_protobuf1.BoolValue {
	if m != nil {
		return m.PassThroughMode
	}
	return nil
}

func (m *HealthCheck) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *HealthCheck) GetCacheTime() *google_protobuf.Duration {
	if m != nil {
		return m.CacheTime
	}
	return nil
}

func init() {
	proto.RegisterType((*HealthCheck)(nil), "envoy.api.v2.filter.http.HealthCheck")
}
func (m *HealthCheck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HealthCheck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.PassThroughMode != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHealthCheck(dAtA, i, uint64(m.PassThroughMode.Size()))
		n1, err := m.PassThroughMode.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Endpoint) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHealthCheck(dAtA, i, uint64(len(m.Endpoint)))
		i += copy(dAtA[i:], m.Endpoint)
	}
	if m.CacheTime != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintHealthCheck(dAtA, i, uint64(m.CacheTime.Size()))
		n2, err := m.CacheTime.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func encodeVarintHealthCheck(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HealthCheck) Size() (n int) {
	var l int
	_ = l
	if m.PassThroughMode != nil {
		l = m.PassThroughMode.Size()
		n += 1 + l + sovHealthCheck(uint64(l))
	}
	l = len(m.Endpoint)
	if l > 0 {
		n += 1 + l + sovHealthCheck(uint64(l))
	}
	if m.CacheTime != nil {
		l = m.CacheTime.Size()
		n += 1 + l + sovHealthCheck(uint64(l))
	}
	return n
}

func sovHealthCheck(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHealthCheck(x uint64) (n int) {
	return sovHealthCheck(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HealthCheck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHealthCheck
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HealthCheck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthCheck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PassThroughMode", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHealthCheck
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PassThroughMode == nil {
				m.PassThroughMode = &google_protobuf1.BoolValue{}
			}
			if err := m.PassThroughMode.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHealthCheck
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Endpoint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CacheTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHealthCheck
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CacheTime == nil {
				m.CacheTime = &google_protobuf.Duration{}
			}
			if err := m.CacheTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHealthCheck(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHealthCheck
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
func skipHealthCheck(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHealthCheck
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
					return 0, ErrIntOverflowHealthCheck
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHealthCheck
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthHealthCheck
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHealthCheck
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipHealthCheck(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthHealthCheck = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHealthCheck   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("api/filter/http/health_check.proto", fileDescriptorHealthCheck) }

var fileDescriptorHealthCheck = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xb1, 0x4e, 0xf3, 0x30,
	0x10, 0x80, 0xe5, 0xf6, 0xff, 0x11, 0x71, 0x91, 0x80, 0x2c, 0x84, 0x0e, 0x51, 0x54, 0x09, 0xa9,
	0x93, 0x2d, 0x95, 0x85, 0x39, 0x30, 0xb0, 0xb0, 0x44, 0x15, 0x6b, 0xe4, 0x26, 0xd7, 0xd8, 0xc2,
	0xc9, 0x59, 0x8e, 0x13, 0xc4, 0x2b, 0xf0, 0x48, 0x4c, 0x88, 0x89, 0x91, 0x47, 0x40, 0xd9, 0x78,
	0x0b, 0x94, 0xa4, 0x65, 0xe9, 0x66, 0xe9, 0xfb, 0xce, 0xdf, 0x1d, 0x5d, 0x08, 0xa3, 0xf8, 0x56,
	0x69, 0x07, 0x96, 0x4b, 0xe7, 0x0c, 0x97, 0x20, 0xb4, 0x93, 0x69, 0x26, 0x21, 0x7b, 0x62, 0xc6,
	0xa2, 0x43, 0x3f, 0x80, 0xaa, 0xc5, 0x17, 0x26, 0x8c, 0x62, 0xed, 0x8a, 0x8d, 0x32, 0xeb, 0xe5,
	0x79, 0x58, 0x20, 0x16, 0x1a, 0xf8, 0xe0, 0x6d, 0x9a, 0x2d, 0xcf, 0x1b, 0x2b, 0x9c, 0xc2, 0x6a,
	0x9c, 0x3c, 0xe4, 0xcf, 0x56, 0x18, 0x03, 0xb6, 0xde, 0xf1, 0x8b, 0x56, 0x68, 0x95, 0x0b, 0x07,
	0x7c, 0xff, 0x18, 0xc1, 0xe2, 0x83, 0xd0, 0xd9, 0xfd, 0xb0, 0xc9, 0x6d, 0xbf, 0x88, 0x9f, 0xd0,
	0x73, 0x23, 0xea, 0x3a, 0x75, 0xd2, 0x62, 0x53, 0xc8, 0xb4, 0xc4, 0x1c, 0x02, 0x12, 0x91, 0xe5,
	0x6c, 0x35, 0x67, 0x63, 0x84, 0xed, 0x23, 0x2c, 0x46, 0xd4, 0x8f, 0x42, 0x37, 0x10, 0xd3, 0xb7,
	0x9f, 0xf7, 0xe9, 0xff, 0x57, 0x32, 0x39, 0x23, 0xc9, 0x69, 0xff, 0xc1, 0x7a, 0x9c, 0x7f, 0xc0,
	0x1c, 0xfc, 0x2b, 0x7a, 0x0c, 0x55, 0x6e, 0x50, 0x55, 0x2e, 0x98, 0x44, 0x64, 0xe9, 0xc5, 0x5e,
	0xaf, 0xff, 0xb3, 0x93, 0x88, 0x24, 0x7f, 0xc8, 0xbf, 0xa1, 0x34, 0x13, 0x99, 0x84, 0xd4, 0xa9,
	0x12, 0x82, 0xe9, 0xd0, 0xbc, 0x3c, 0x68, 0xde, 0xed, 0x0e, 0x4f, 0xbc, 0x41, 0x5e, 0xab, 0x12,
	0xe2, 0x93, 0xcf, 0x2e, 0x24, 0x5f, 0x5d, 0x48, 0xbe, 0xbb, 0x90, 0x6c, 0x8e, 0x06, 0xf7, 0xfa,
	0x37, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x89, 0xb9, 0xe1, 0x72, 0x01, 0x00, 0x00,
}