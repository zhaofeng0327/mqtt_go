// Code generated by protoc-gen-go. DO NOT EDIT.
// source: battery_ageing.proto

package battery_ageing

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type HEART_BEAT_BATTERY_CODE int32

const (
	HEART_BEAT_BATTERY_CODE_HBB_STATUS_OK           HEART_BEAT_BATTERY_CODE = 0
	HEART_BEAT_BATTERY_CODE_HBB_OTHER_ERROR         HEART_BEAT_BATTERY_CODE = 1
	HEART_BEAT_BATTERY_CODE_HBB_TEMPERATURE_ERROR   HEART_BEAT_BATTERY_CODE = 2
	HEART_BEAT_BATTERY_CODE_HBB_VOLTAGE_ERROR       HEART_BEAT_BATTERY_CODE = 3
	HEART_BEAT_BATTERY_CODE_HBB_CURRENT_ERROR       HEART_BEAT_BATTERY_CODE = 4
	HEART_BEAT_BATTERY_CODE_HBB_CYCLECOUNT_ERROR    HEART_BEAT_BATTERY_CODE = 5
	HEART_BEAT_BATTERY_CODE_HBB_INTERFACE_ERROR     HEART_BEAT_BATTERY_CODE = 6
	HEART_BEAT_BATTERY_CODE_HBB_BODY_ERROR          HEART_BEAT_BATTERY_CODE = 7
	HEART_BEAT_BATTERY_CODE_HBB_LINE_ERROR          HEART_BEAT_BATTERY_CODE = 8
	HEART_BEAT_BATTERY_CODE_HBB_BODY_AND_LINE_ERROR HEART_BEAT_BATTERY_CODE = 9
)

var HEART_BEAT_BATTERY_CODE_name = map[int32]string{
	0: "HBB_STATUS_OK",
	1: "HBB_OTHER_ERROR",
	2: "HBB_TEMPERATURE_ERROR",
	3: "HBB_VOLTAGE_ERROR",
	4: "HBB_CURRENT_ERROR",
	5: "HBB_CYCLECOUNT_ERROR",
	6: "HBB_INTERFACE_ERROR",
	7: "HBB_BODY_ERROR",
	8: "HBB_LINE_ERROR",
	9: "HBB_BODY_AND_LINE_ERROR",
}

var HEART_BEAT_BATTERY_CODE_value = map[string]int32{
	"HBB_STATUS_OK":           0,
	"HBB_OTHER_ERROR":         1,
	"HBB_TEMPERATURE_ERROR":   2,
	"HBB_VOLTAGE_ERROR":       3,
	"HBB_CURRENT_ERROR":       4,
	"HBB_CYCLECOUNT_ERROR":    5,
	"HBB_INTERFACE_ERROR":     6,
	"HBB_BODY_ERROR":          7,
	"HBB_LINE_ERROR":          8,
	"HBB_BODY_AND_LINE_ERROR": 9,
}

func (x HEART_BEAT_BATTERY_CODE) Enum() *HEART_BEAT_BATTERY_CODE {
	p := new(HEART_BEAT_BATTERY_CODE)
	*p = x
	return p
}

func (x HEART_BEAT_BATTERY_CODE) String() string {
	return proto.EnumName(HEART_BEAT_BATTERY_CODE_name, int32(x))
}

func (x *HEART_BEAT_BATTERY_CODE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HEART_BEAT_BATTERY_CODE_value, data, "HEART_BEAT_BATTERY_CODE")
	if err != nil {
		return err
	}
	*x = HEART_BEAT_BATTERY_CODE(value)
	return nil
}

func (HEART_BEAT_BATTERY_CODE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_23ea8d3492947939, []int{0}
}

type CHARGE_STATUS int32

const (
	CHARGE_STATUS_UNCHARGE    CHARGE_STATUS = 1
	CHARGE_STATUS_CHARGING    CHARGE_STATUS = 2
	CHARGE_STATUS_DISCHARGING CHARGE_STATUS = 3
)

var CHARGE_STATUS_name = map[int32]string{
	1: "UNCHARGE",
	2: "CHARGING",
	3: "DISCHARGING",
}

var CHARGE_STATUS_value = map[string]int32{
	"UNCHARGE":    1,
	"CHARGING":    2,
	"DISCHARGING": 3,
}

func (x CHARGE_STATUS) Enum() *CHARGE_STATUS {
	p := new(CHARGE_STATUS)
	*p = x
	return p
}

func (x CHARGE_STATUS) String() string {
	return proto.EnumName(CHARGE_STATUS_name, int32(x))
}

func (x *CHARGE_STATUS) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CHARGE_STATUS_value, data, "CHARGE_STATUS")
	if err != nil {
		return err
	}
	*x = CHARGE_STATUS(value)
	return nil
}

func (CHARGE_STATUS) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_23ea8d3492947939, []int{1}
}

type ENABLE_BORROW_STATUS int32

const (
	ENABLE_BORROW_STATUS_ENABLE_STATUS  ENABLE_BORROW_STATUS = 0
	ENABLE_BORROW_STATUS_DISABLE_STATUS ENABLE_BORROW_STATUS = 1
)

var ENABLE_BORROW_STATUS_name = map[int32]string{
	0: "ENABLE_STATUS",
	1: "DISABLE_STATUS",
}

var ENABLE_BORROW_STATUS_value = map[string]int32{
	"ENABLE_STATUS":  0,
	"DISABLE_STATUS": 1,
}

func (x ENABLE_BORROW_STATUS) Enum() *ENABLE_BORROW_STATUS {
	p := new(ENABLE_BORROW_STATUS)
	*p = x
	return p
}

func (x ENABLE_BORROW_STATUS) String() string {
	return proto.EnumName(ENABLE_BORROW_STATUS_name, int32(x))
}

func (x *ENABLE_BORROW_STATUS) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ENABLE_BORROW_STATUS_value, data, "ENABLE_BORROW_STATUS")
	if err != nil {
		return err
	}
	*x = ENABLE_BORROW_STATUS(value)
	return nil
}

func (ENABLE_BORROW_STATUS) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_23ea8d3492947939, []int{2}
}

type BatteryAgeingInfo struct {
	Timestamp            *string                  `protobuf:"bytes,1,req,name=timestamp" json:"timestamp,omitempty"`
	BatterySn            *string                  `protobuf:"bytes,2,req,name=battery_sn,json=batterySn" json:"battery_sn,omitempty"`
	SlotNum              *int32                   `protobuf:"varint,3,req,name=slot_num,json=slotNum" json:"slot_num,omitempty"`
	Voltage              *int32                   `protobuf:"varint,4,req,name=voltage" json:"voltage,omitempty"`
	Current              *int32                   `protobuf:"varint,5,req,name=current" json:"current,omitempty"`
	Temprature           *int32                   `protobuf:"varint,6,req,name=temprature" json:"temprature,omitempty"`
	Elapsed              *int32                   `protobuf:"varint,7,req,name=elapsed" json:"elapsed,omitempty"`
	Discharging          *int32                   `protobuf:"varint,8,req,name=discharging" json:"discharging,omitempty"`
	XTemprature          *int32                   `protobuf:"varint,9,req,name=_temprature,json=Temprature" json:"_temprature,omitempty"`
	XVoltage             *int32                   `protobuf:"varint,10,req,name=_voltage,json=Voltage" json:"_voltage,omitempty"`
	XFullChargecapacity  *int32                   `protobuf:"varint,11,req,name=_fullChargecapacity,json=FullChargecapacity" json:"_fullChargecapacity,omitempty"`
	XRemainingcapacity   *int32                   `protobuf:"varint,12,req,name=_remainingcapacity,json=Remainingcapacity" json:"_remainingcapacity,omitempty"`
	XAveragecurrent      *int32                   `protobuf:"varint,13,req,name=_averagecurrent,json=Averagecurrent" json:"_averagecurrent,omitempty"`
	XCyclecount          *int32                   `protobuf:"varint,14,req,name=_cyclecount,json=Cyclecount" json:"_cyclecount,omitempty"`
	XBmssafetyStatus     *int32                   `protobuf:"varint,15,req,name=_bmssafetyStatus,json=BmssafetyStatus" json:"_bmssafetyStatus,omitempty"`
	XBmsflags            *int32                   `protobuf:"varint,16,req,name=_bmsflags,json=Bmsflags" json:"_bmsflags,omitempty"`
	XBatterystatus       *HEART_BEAT_BATTERY_CODE `protobuf:"varint,17,req,name=_batterystatus,json=Batterystatus,enum=HEART_BEAT_BATTERY_CODE" json:"_batterystatus,omitempty"`
	XChargestatus        *CHARGE_STATUS           `protobuf:"varint,18,req,name=_chargestatus,json=Chargestatus,enum=CHARGE_STATUS" json:"_chargestatus,omitempty"`
	XEnablestatus        *ENABLE_BORROW_STATUS    `protobuf:"varint,19,req,name=_enablestatus,json=Enablestatus,enum=ENABLE_BORROW_STATUS" json:"_enablestatus,omitempty"`
	XSlotstatus          *int32                   `protobuf:"varint,20,req,name=_slotstatus,json=Slotstatus" json:"_slotstatus,omitempty"`
	XDestroyed           *int32                   `protobuf:"varint,21,req,name=_destroyed,json=Destroyed" json:"_destroyed,omitempty"`
	XHasbms              *int32                   `protobuf:"varint,22,req,name=_hasbms,json=Hasbms" json:"_hasbms,omitempty"`
	XRadio               *int32                   `protobuf:"varint,23,req,name=_radio,json=Radio" json:"_radio,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *BatteryAgeingInfo) Reset()         { *m = BatteryAgeingInfo{} }
func (m *BatteryAgeingInfo) String() string { return proto.CompactTextString(m) }
func (*BatteryAgeingInfo) ProtoMessage()    {}
func (*BatteryAgeingInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_23ea8d3492947939, []int{0}
}

func (m *BatteryAgeingInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatteryAgeingInfo.Unmarshal(m, b)
}
func (m *BatteryAgeingInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatteryAgeingInfo.Marshal(b, m, deterministic)
}
func (m *BatteryAgeingInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatteryAgeingInfo.Merge(m, src)
}
func (m *BatteryAgeingInfo) XXX_Size() int {
	return xxx_messageInfo_BatteryAgeingInfo.Size(m)
}
func (m *BatteryAgeingInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BatteryAgeingInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BatteryAgeingInfo proto.InternalMessageInfo

func (m *BatteryAgeingInfo) GetTimestamp() string {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return ""
}

func (m *BatteryAgeingInfo) GetBatterySn() string {
	if m != nil && m.BatterySn != nil {
		return *m.BatterySn
	}
	return ""
}

func (m *BatteryAgeingInfo) GetSlotNum() int32 {
	if m != nil && m.SlotNum != nil {
		return *m.SlotNum
	}
	return 0
}

func (m *BatteryAgeingInfo) GetVoltage() int32 {
	if m != nil && m.Voltage != nil {
		return *m.Voltage
	}
	return 0
}

func (m *BatteryAgeingInfo) GetCurrent() int32 {
	if m != nil && m.Current != nil {
		return *m.Current
	}
	return 0
}

func (m *BatteryAgeingInfo) GetTemprature() int32 {
	if m != nil && m.Temprature != nil {
		return *m.Temprature
	}
	return 0
}

func (m *BatteryAgeingInfo) GetElapsed() int32 {
	if m != nil && m.Elapsed != nil {
		return *m.Elapsed
	}
	return 0
}

func (m *BatteryAgeingInfo) GetDischarging() int32 {
	if m != nil && m.Discharging != nil {
		return *m.Discharging
	}
	return 0
}

func (m *BatteryAgeingInfo) GetXTemprature() int32 {
	if m != nil && m.XTemprature != nil {
		return *m.XTemprature
	}
	return 0
}

func (m *BatteryAgeingInfo) GetXVoltage() int32 {
	if m != nil && m.XVoltage != nil {
		return *m.XVoltage
	}
	return 0
}

func (m *BatteryAgeingInfo) GetXFullChargecapacity() int32 {
	if m != nil && m.XFullChargecapacity != nil {
		return *m.XFullChargecapacity
	}
	return 0
}

func (m *BatteryAgeingInfo) GetXRemainingcapacity() int32 {
	if m != nil && m.XRemainingcapacity != nil {
		return *m.XRemainingcapacity
	}
	return 0
}

func (m *BatteryAgeingInfo) GetXAveragecurrent() int32 {
	if m != nil && m.XAveragecurrent != nil {
		return *m.XAveragecurrent
	}
	return 0
}

func (m *BatteryAgeingInfo) GetXCyclecount() int32 {
	if m != nil && m.XCyclecount != nil {
		return *m.XCyclecount
	}
	return 0
}

func (m *BatteryAgeingInfo) GetXBmssafetyStatus() int32 {
	if m != nil && m.XBmssafetyStatus != nil {
		return *m.XBmssafetyStatus
	}
	return 0
}

func (m *BatteryAgeingInfo) GetXBmsflags() int32 {
	if m != nil && m.XBmsflags != nil {
		return *m.XBmsflags
	}
	return 0
}

func (m *BatteryAgeingInfo) GetXBatterystatus() HEART_BEAT_BATTERY_CODE {
	if m != nil && m.XBatterystatus != nil {
		return *m.XBatterystatus
	}
	return HEART_BEAT_BATTERY_CODE_HBB_STATUS_OK
}

func (m *BatteryAgeingInfo) GetXChargestatus() CHARGE_STATUS {
	if m != nil && m.XChargestatus != nil {
		return *m.XChargestatus
	}
	return CHARGE_STATUS_UNCHARGE
}

func (m *BatteryAgeingInfo) GetXEnablestatus() ENABLE_BORROW_STATUS {
	if m != nil && m.XEnablestatus != nil {
		return *m.XEnablestatus
	}
	return ENABLE_BORROW_STATUS_ENABLE_STATUS
}

func (m *BatteryAgeingInfo) GetXSlotstatus() int32 {
	if m != nil && m.XSlotstatus != nil {
		return *m.XSlotstatus
	}
	return 0
}

func (m *BatteryAgeingInfo) GetXDestroyed() int32 {
	if m != nil && m.XDestroyed != nil {
		return *m.XDestroyed
	}
	return 0
}

func (m *BatteryAgeingInfo) GetXHasbms() int32 {
	if m != nil && m.XHasbms != nil {
		return *m.XHasbms
	}
	return 0
}

func (m *BatteryAgeingInfo) GetXRadio() int32 {
	if m != nil && m.XRadio != nil {
		return *m.XRadio
	}
	return 0
}

type UploadInfo struct {
	BatteryAgeingInfo    []*BatteryAgeingInfo `protobuf:"bytes,1,rep,name=battery_ageing_info,json=batteryAgeingInfo" json:"battery_ageing_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UploadInfo) Reset()         { *m = UploadInfo{} }
func (m *UploadInfo) String() string { return proto.CompactTextString(m) }
func (*UploadInfo) ProtoMessage()    {}
func (*UploadInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_23ea8d3492947939, []int{1}
}

func (m *UploadInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadInfo.Unmarshal(m, b)
}
func (m *UploadInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadInfo.Marshal(b, m, deterministic)
}
func (m *UploadInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadInfo.Merge(m, src)
}
func (m *UploadInfo) XXX_Size() int {
	return xxx_messageInfo_UploadInfo.Size(m)
}
func (m *UploadInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UploadInfo proto.InternalMessageInfo

func (m *UploadInfo) GetBatteryAgeingInfo() []*BatteryAgeingInfo {
	if m != nil {
		return m.BatteryAgeingInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*BatteryAgeingInfo)(nil), "BatteryAgeingInfo")
	proto.RegisterType((*UploadInfo)(nil), "UploadInfo")
	proto.RegisterEnum("HEART_BEAT_BATTERY_CODE", HEART_BEAT_BATTERY_CODE_name, HEART_BEAT_BATTERY_CODE_value)
	proto.RegisterEnum("CHARGE_STATUS", CHARGE_STATUS_name, CHARGE_STATUS_value)
	proto.RegisterEnum("ENABLE_BORROW_STATUS", ENABLE_BORROW_STATUS_name, ENABLE_BORROW_STATUS_value)
}

func init() { proto.RegisterFile("battery_ageing.proto", fileDescriptor_23ea8d3492947939) }

var fileDescriptor_23ea8d3492947939 = []byte{
	// 712 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x5b, 0x8f, 0xda, 0x46,
	0x14, 0xc7, 0x03, 0x1b, 0x6e, 0x87, 0xe5, 0x36, 0x40, 0x99, 0x55, 0x7a, 0x41, 0x79, 0x29, 0x5d,
	0xa9, 0x54, 0x4a, 0xdf, 0xaa, 0x56, 0x95, 0x6d, 0x26, 0x0b, 0x2a, 0xb5, 0xa3, 0xc1, 0xa4, 0xda,
	0xa7, 0xd1, 0x00, 0x83, 0x63, 0xc9, 0x17, 0x64, 0x0f, 0x91, 0x78, 0xed, 0xd7, 0xeb, 0x97, 0xaa,
	0x66, 0x6c, 0x13, 0x93, 0xb4, 0x6f, 0x9c, 0xdf, 0xef, 0x7f, 0x8e, 0x3c, 0x87, 0x19, 0x18, 0xed,
	0xb8, 0x94, 0x22, 0xb9, 0x30, 0xee, 0x09, 0x3f, 0xf2, 0xe6, 0xa7, 0x24, 0x96, 0xf1, 0xeb, 0x7f,
	0xea, 0x30, 0x30, 0x33, 0x61, 0x68, 0xbe, 0x8a, 0x8e, 0x31, 0xfa, 0x1a, 0x5a, 0xd2, 0x0f, 0x45,
	0x2a, 0x79, 0x78, 0xc2, 0x95, 0x69, 0x75, 0xd6, 0xa2, 0x9f, 0x00, 0xfa, 0x06, 0xa0, 0x98, 0x95,
	0x46, 0xb8, 0x9a, 0xe9, 0x9c, 0x6c, 0x22, 0xf4, 0x00, 0xcd, 0x34, 0x88, 0x25, 0x8b, 0xce, 0x21,
	0xbe, 0x9b, 0x56, 0x67, 0x35, 0xda, 0x50, 0xb5, 0x7d, 0x0e, 0x11, 0x86, 0xc6, 0xc7, 0x38, 0x90,
	0xdc, 0x13, 0xf8, 0x65, 0x66, 0xf2, 0x52, 0x99, 0xfd, 0x39, 0x49, 0x44, 0x24, 0x71, 0x2d, 0x33,
	0x79, 0x89, 0xbe, 0x05, 0x90, 0x22, 0x3c, 0x25, 0x5c, 0x9e, 0x13, 0x81, 0xeb, 0x5a, 0x96, 0x88,
	0xea, 0x14, 0x01, 0x3f, 0xa5, 0xe2, 0x80, 0x1b, 0x59, 0x67, 0x5e, 0xa2, 0x29, 0xb4, 0x0f, 0x7e,
	0xba, 0xff, 0xc0, 0x13, 0xcf, 0x8f, 0x3c, 0xdc, 0xd4, 0xb6, 0x8c, 0xd0, 0x77, 0xd0, 0x66, 0xa5,
	0xe1, 0xad, 0x6c, 0xb8, 0xfb, 0x69, 0xf8, 0x03, 0x34, 0x59, 0xf1, 0xc5, 0x90, 0x4d, 0x7f, 0x9f,
	0x7f, 0xf1, 0x4f, 0x30, 0x64, 0xc7, 0x73, 0x10, 0x58, 0x6a, 0x98, 0xd8, 0xf3, 0x13, 0xdf, 0xfb,
	0xf2, 0x82, 0xdb, 0x3a, 0x85, 0xde, 0x7e, 0x61, 0xd0, 0x8f, 0x80, 0x58, 0x22, 0x42, 0xee, 0x47,
	0x7e, 0xe4, 0x5d, 0xf3, 0xf7, 0x3a, 0x3f, 0xa0, 0x9f, 0x0b, 0xf4, 0x3d, 0xf4, 0x18, 0xff, 0x28,
	0x12, 0xee, 0x89, 0x62, 0x33, 0x1d, 0x9d, 0xed, 0x1a, 0x37, 0x54, 0x1f, 0x62, 0x7f, 0xd9, 0x07,
	0x62, 0x1f, 0x9f, 0x23, 0x89, 0xbb, 0xd9, 0x21, 0xac, 0x2b, 0x41, 0x3f, 0x40, 0x9f, 0xed, 0xc2,
	0x34, 0xe5, 0x47, 0x21, 0x2f, 0x1b, 0xc9, 0xe5, 0x39, 0xc5, 0x3d, 0x9d, 0xea, 0x99, 0xb7, 0x18,
	0xbd, 0x82, 0x96, 0x8a, 0x1e, 0x03, 0xee, 0xa5, 0xb8, 0xaf, 0x33, 0x4d, 0x33, 0xaf, 0xd1, 0xef,
	0xd0, 0x65, 0xf9, 0xdf, 0x9c, 0x66, 0x53, 0x06, 0xd3, 0xea, 0xac, 0xfb, 0x06, 0xcf, 0x97, 0xc4,
	0xa0, 0x2e, 0x33, 0x89, 0xe1, 0x32, 0xd3, 0x70, 0x5d, 0x42, 0x9f, 0x99, 0xe5, 0x2c, 0x08, 0xed,
	0x98, 0xe5, 0x38, 0xfa, 0x19, 0x3a, 0x4c, 0xef, 0x5e, 0xe4, 0xfd, 0x48, 0xf7, 0x77, 0xe7, 0xd6,
	0xd2, 0xa0, 0x4f, 0x84, 0x6d, 0x5c, 0xc3, 0xdd, 0x6e, 0xe8, 0xbd, 0x55, 0xca, 0xa0, 0x5f, 0xa0,
	0xc3, 0x44, 0xc4, 0x77, 0x41, 0xd1, 0x34, 0xd4, 0x4d, 0xe3, 0x39, 0xb1, 0x0d, 0x73, 0x4d, 0x98,
	0xe9, 0x50, 0xea, 0xfc, 0x75, 0xed, 0x25, 0xa5, 0xa8, 0x5e, 0x8d, 0xba, 0x7b, 0x79, 0xe7, 0x28,
	0x5b, 0xcd, 0xe6, 0x4a, 0xd4, 0x55, 0x66, 0x07, 0x91, 0xca, 0x24, 0xbe, 0x88, 0x03, 0x1e, 0x6b,
	0xdf, 0x5a, 0x14, 0x00, 0x4d, 0xa0, 0xc1, 0x3e, 0xf0, 0x74, 0x17, 0xa6, 0xf8, 0x2b, 0xed, 0xea,
	0x4b, 0x5d, 0xa1, 0x31, 0xd4, 0x59, 0xc2, 0x0f, 0x7e, 0x8c, 0x27, 0x9a, 0xd7, 0xa8, 0x2a, 0x5e,
	0xbf, 0x03, 0xd8, 0x9e, 0x82, 0x98, 0x1f, 0xf4, 0x2b, 0x32, 0x61, 0x78, 0xfb, 0xe6, 0x98, 0x1f,
	0x1d, 0x63, 0x5c, 0x99, 0xde, 0xcd, 0xda, 0x6f, 0xd0, 0xfc, 0x8b, 0x67, 0x47, 0x07, 0xbb, 0xcf,
	0xd1, 0xe3, 0xdf, 0x55, 0x98, 0xfc, 0xcf, 0x76, 0xd1, 0x00, 0x3a, 0x4b, 0xd3, 0xcc, 0x4f, 0xce,
	0x9c, 0x3f, 0xfa, 0x2f, 0xd0, 0x10, 0x7a, 0x0a, 0x39, 0xee, 0x92, 0x50, 0x46, 0x28, 0x75, 0x68,
	0xbf, 0x82, 0x1e, 0x60, 0xac, 0xa0, 0x4b, 0xfe, 0x7c, 0x47, 0xa8, 0xe1, 0x6e, 0x29, 0xc9, 0x55,
	0x15, 0x8d, 0x61, 0xa0, 0xd4, 0x7b, 0x67, 0xed, 0x1a, 0x4f, 0x05, 0xbe, 0x2b, 0xb0, 0xb5, 0xa5,
	0x94, 0xd8, 0x6e, 0x8e, 0x5f, 0x22, 0x0c, 0x23, 0x8d, 0x9f, 0xad, 0x35, 0xb1, 0x9c, 0xed, 0xd5,
	0xd4, 0xd0, 0x04, 0x86, 0xca, 0xac, 0x6c, 0x97, 0xd0, 0xb7, 0x86, 0x55, 0x4c, 0xaa, 0x23, 0x04,
	0x5d, 0x25, 0x4c, 0x67, 0xf1, 0x9c, 0xb3, 0x46, 0xc1, 0xd6, 0x2b, 0xbb, 0xc8, 0x35, 0xd1, 0x2b,
	0x98, 0x5c, 0x73, 0x86, 0xbd, 0x28, 0xcb, 0xd6, 0xe3, 0xaf, 0xd0, 0xb9, 0xb9, 0x21, 0xe8, 0x1e,
	0x9a, 0x5b, 0x3b, 0x43, 0xfd, 0x8a, 0xaa, 0xf4, 0xef, 0x95, 0xfd, 0xd4, 0xaf, 0xa2, 0x1e, 0xb4,
	0x17, 0xab, 0xcd, 0x15, 0xdc, 0x3d, 0xfe, 0x06, 0xa3, 0xff, 0xba, 0x2a, 0x6a, 0x7d, 0x39, 0xcf,
	0x40, 0xff, 0x85, 0xfa, 0xb2, 0xc5, 0x6a, 0x53, 0x66, 0x95, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xd5, 0xcd, 0xde, 0x81, 0x38, 0x05, 0x00, 0x00,
}
