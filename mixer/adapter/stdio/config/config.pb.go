// Code generated by protoc-gen-gogo.
// source: mixer/adapter/stdio/config/config.proto
// DO NOT EDIT!

/*
	Package config is a generated protocol buffer package.

	It is generated from these files:
		mixer/adapter/stdio/config/config.proto

	It has these top-level messages:
		Params
*/
package config

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strconv "strconv"

import strings "strings"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Stream is used to select between different log output sinks.
type Params_Stream int32

const (
	STDOUT Params_Stream = 0
	STDERR Params_Stream = 1
)

var Params_Stream_name = map[int32]string{
	0: "STDOUT",
	1: "STDERR",
}
var Params_Stream_value = map[string]int32{
	"STDOUT": 0,
	"STDERR": 1,
}

func (Params_Stream) EnumDescriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0, 0} }

// Importance level for individual items output by this adapter.
type Params_Level int32

const (
	INFO    Params_Level = 0
	WARNING Params_Level = 1
	ERROR   Params_Level = 2
)

var Params_Level_name = map[int32]string{
	0: "INFO",
	1: "WARNING",
	2: "ERROR",
}
var Params_Level_value = map[string]int32{
	"INFO":    0,
	"WARNING": 1,
	"ERROR":   2,
}

func (Params_Level) EnumDescriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0, 1} }

type Params struct {
	// Selects which standard stream to write to for log entries.
	// STDERR is the default Stream.
	LogStream Params_Stream `protobuf:"varint,1,opt,name=log_stream,json=logStream,proto3,enum=adapter.stdio.config.Params_Stream" json:"log_stream,omitempty"`
	// Maps from severity strings as specified in LogEntry instances to
	// the set of levels supported by this adapter.
	SeverityLevels map[string]Params_Level `protobuf:"bytes,2,rep,name=severity_levels,json=severityLevels" json:"severity_levels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=adapter.stdio.config.Params_Level"`
	// The level to assign to metrics being output.
	MetricLevel Params_Level `protobuf:"varint,3,opt,name=metric_level,json=metricLevel,proto3,enum=adapter.stdio.config.Params_Level" json:"metric_level,omitempty"`
	// Whether to output a console-friendly or json-friendly format
	OutputAsJson bool `protobuf:"varint,4,opt,name=output_as_json,json=outputAsJson,proto3" json:"output_as_json,omitempty"`
}

func (m *Params) Reset()                    { *m = Params{} }
func (*Params) ProtoMessage()               {}
func (*Params) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func init() {
	proto.RegisterType((*Params)(nil), "adapter.stdio.config.Params")
	proto.RegisterEnum("adapter.stdio.config.Params_Stream", Params_Stream_name, Params_Stream_value)
	proto.RegisterEnum("adapter.stdio.config.Params_Level", Params_Level_name, Params_Level_value)
}
func (x Params_Stream) String() string {
	s, ok := Params_Stream_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x Params_Level) String() string {
	s, ok := Params_Level_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.LogStream != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.LogStream))
	}
	if len(m.SeverityLevels) > 0 {
		for k, _ := range m.SeverityLevels {
			dAtA[i] = 0x12
			i++
			v := m.SeverityLevels[k]
			mapSize := 1 + len(k) + sovConfig(uint64(len(k))) + 1 + sovConfig(uint64(v))
			i = encodeVarintConfig(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintConfig(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x10
			i++
			i = encodeVarintConfig(dAtA, i, uint64(v))
		}
	}
	if m.MetricLevel != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.MetricLevel))
	}
	if m.OutputAsJson {
		dAtA[i] = 0x20
		i++
		if m.OutputAsJson {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeFixed64Config(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Config(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Params) Size() (n int) {
	var l int
	_ = l
	if m.LogStream != 0 {
		n += 1 + sovConfig(uint64(m.LogStream))
	}
	if len(m.SeverityLevels) > 0 {
		for k, v := range m.SeverityLevels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovConfig(uint64(len(k))) + 1 + sovConfig(uint64(v))
			n += mapEntrySize + 1 + sovConfig(uint64(mapEntrySize))
		}
	}
	if m.MetricLevel != 0 {
		n += 1 + sovConfig(uint64(m.MetricLevel))
	}
	if m.OutputAsJson {
		n += 2
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Params) String() string {
	if this == nil {
		return "nil"
	}
	keysForSeverityLevels := make([]string, 0, len(this.SeverityLevels))
	for k, _ := range this.SeverityLevels {
		keysForSeverityLevels = append(keysForSeverityLevels, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForSeverityLevels)
	mapStringForSeverityLevels := "map[string]Params_Level{"
	for _, k := range keysForSeverityLevels {
		mapStringForSeverityLevels += fmt.Sprintf("%v: %v,", k, this.SeverityLevels[k])
	}
	mapStringForSeverityLevels += "}"
	s := strings.Join([]string{`&Params{`,
		`LogStream:` + fmt.Sprintf("%v", this.LogStream) + `,`,
		`SeverityLevels:` + mapStringForSeverityLevels + `,`,
		`MetricLevel:` + fmt.Sprintf("%v", this.MetricLevel) + `,`,
		`OutputAsJson:` + fmt.Sprintf("%v", this.OutputAsJson) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringConfig(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogStream", wireType)
			}
			m.LogStream = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LogStream |= (Params_Stream(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SeverityLevels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthConfig
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.SeverityLevels == nil {
				m.SeverityLevels = make(map[string]Params_Level)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var mapvalue Params_Level
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					mapvalue |= (Params_Level(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.SeverityLevels[mapkey] = mapvalue
			} else {
				var mapvalue Params_Level
				m.SeverityLevels[mapkey] = mapvalue
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetricLevel", wireType)
			}
			m.MetricLevel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MetricLevel |= (Params_Level(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutputAsJson", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OutputAsJson = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
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
				next, err := skipConfig(dAtA[start:])
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
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("mixer/adapter/stdio/config/config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6e, 0xd3, 0x30,
	0x1c, 0xc6, 0xed, 0x64, 0x0d, 0xeb, 0xbf, 0x53, 0x89, 0xcc, 0x0e, 0xd1, 0x0e, 0x56, 0x14, 0x90,
	0x28, 0x97, 0x04, 0x8d, 0x4b, 0xc5, 0x6d, 0x13, 0x01, 0x0d, 0xa1, 0x16, 0x79, 0x43, 0x08, 0x2e,
	0x51, 0xd8, 0x4c, 0x14, 0x48, 0xe2, 0xca, 0x76, 0x2b, 0x7a, 0xe3, 0x11, 0x78, 0x0c, 0x1e, 0xa5,
	0xc7, 0x1d, 0x39, 0x92, 0x70, 0xe1, 0xb8, 0x47, 0x40, 0xb3, 0x73, 0x41, 0xaa, 0xaa, 0x9d, 0xfc,
	0xf9, 0xef, 0xef, 0xf7, 0xe9, 0xd3, 0x5f, 0x86, 0xc7, 0x75, 0xf9, 0x8d, 0xcb, 0x24, 0xbf, 0xca,
	0x17, 0x9a, 0xcb, 0x44, 0xe9, 0xab, 0x52, 0x24, 0x97, 0xa2, 0xf9, 0x5c, 0x16, 0xfd, 0x11, 0x2f,
	0xa4, 0xd0, 0x82, 0x1c, 0xf6, 0x96, 0xd8, 0x58, 0x62, 0xfb, 0x76, 0x74, 0x58, 0x88, 0x42, 0x18,
	0x43, 0x72, 0xab, 0xac, 0x37, 0xda, 0xb8, 0xe0, 0xbd, 0xcd, 0x65, 0x5e, 0x2b, 0x72, 0x0a, 0x50,
	0x89, 0x22, 0x53, 0x5a, 0xf2, 0xbc, 0x0e, 0x70, 0x88, 0x27, 0xe3, 0xe3, 0x87, 0xf1, 0xb6, 0xac,
	0xd8, 0x12, 0xf1, 0xb9, 0xb1, 0xb2, 0x61, 0x25, 0x0a, 0x2b, 0xc9, 0x07, 0xb8, 0xaf, 0xf8, 0x8a,
	0xcb, 0x52, 0xaf, 0xb3, 0x8a, 0xaf, 0x78, 0xa5, 0x02, 0x27, 0x74, 0x27, 0xa3, 0xe3, 0xa7, 0xbb,
	0x83, 0x7a, 0xe6, 0x8d, 0x41, 0xd2, 0x46, 0xcb, 0x35, 0x1b, 0xab, 0xff, 0x86, 0x24, 0x85, 0x83,
	0x9a, 0x6b, 0x59, 0x5e, 0xda, 0xe0, 0xc0, 0x35, 0x05, 0xa3, 0x9d, 0xb9, 0x06, 0x65, 0x23, 0xcb,
	0x99, 0x0b, 0x79, 0x04, 0x63, 0xb1, 0xd4, 0x8b, 0xa5, 0xce, 0x72, 0x95, 0x7d, 0x51, 0xa2, 0x09,
	0xf6, 0x42, 0x3c, 0xd9, 0x67, 0x07, 0x76, 0x7a, 0xa2, 0x5e, 0x2b, 0xd1, 0x1c, 0x71, 0x78, 0xb0,
	0xa5, 0x13, 0xf1, 0xc1, 0xfd, 0xca, 0xd7, 0x66, 0x37, 0x43, 0x76, 0x2b, 0xc9, 0x14, 0x06, 0xab,
	0xbc, 0x5a, 0xf2, 0xc0, 0xb9, 0x73, 0x1d, 0x0b, 0x3c, 0x77, 0xa6, 0x38, 0x0a, 0xc1, 0xeb, 0x17,
	0x07, 0xe0, 0x9d, 0x5f, 0xbc, 0x98, 0xbf, 0xbb, 0xf0, 0x51, 0xaf, 0x53, 0xc6, 0x7c, 0x1c, 0x3d,
	0x81, 0x81, 0xed, 0xbd, 0x0f, 0x7b, 0x67, 0xb3, 0x97, 0x73, 0x1f, 0x91, 0x11, 0xdc, 0x7b, 0x7f,
	0xc2, 0x66, 0x67, 0xb3, 0x57, 0x3e, 0x26, 0x43, 0x18, 0xa4, 0x8c, 0xcd, 0x99, 0xef, 0x9c, 0x4e,
	0x37, 0x2d, 0x45, 0xd7, 0x2d, 0x45, 0xbf, 0x5a, 0x8a, 0x6e, 0x5a, 0x8a, 0xbe, 0x77, 0x14, 0xff,
	0xec, 0x28, 0xda, 0x74, 0x14, 0x5f, 0x77, 0x14, 0xff, 0xee, 0x28, 0xfe, 0xdb, 0x51, 0x74, 0xd3,
	0x51, 0xfc, 0xe3, 0x0f, 0x45, 0x1f, 0x3d, 0x5b, 0xef, 0x93, 0x67, 0xfe, 0xc2, 0xb3, 0x7f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xa8, 0x48, 0x88, 0x88, 0x62, 0x02, 0x00, 0x00,
}
