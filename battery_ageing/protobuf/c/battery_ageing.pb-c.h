/* Generated by the protocol buffer compiler.  DO NOT EDIT! */
/* Generated from: battery_ageing.proto */

#ifndef PROTOBUF_C_battery_5fageing_2eproto__INCLUDED
#define PROTOBUF_C_battery_5fageing_2eproto__INCLUDED

#include <protobuf-c/protobuf-c.h>

PROTOBUF_C__BEGIN_DECLS

#if PROTOBUF_C_VERSION_NUMBER < 1000000
# error This file was generated by a newer version of protoc-c which is incompatible with your libprotobuf-c headers. Please update your headers.
#elif 1002001 < PROTOBUF_C_MIN_COMPILER_VERSION
# error This file was generated by an older version of protoc-c which is incompatible with your libprotobuf-c headers. Please regenerate this file with a newer version of protoc-c.
#endif


typedef struct _BatteryAgeingInfo BatteryAgeingInfo;
typedef struct _UploadInfo UploadInfo;
typedef struct _DischargeSetting DischargeSetting;
typedef struct _MSGBODY MSGBODY;


/* --- enums --- */

typedef enum _HEARTBEATBATTERYCODE {
  /*
   *正常成功
   */
  HEART__BEAT__BATTERY__CODE__HBB_STATUS_OK = 0,
  /*
   *其他错误
   */
  HEART__BEAT__BATTERY__CODE__HBB_OTHER_ERROR = 1,
  /*
   *温度错误
   */
  HEART__BEAT__BATTERY__CODE__HBB_TEMPERATURE_ERROR = 2,
  /*
   *电压错误
   */
  HEART__BEAT__BATTERY__CODE__HBB_VOLTAGE_ERROR = 3,
  /*
   *电流错误
   */
  HEART__BEAT__BATTERY__CODE__HBB_CURRENT_ERROR = 4,
  /*
   *电流错误
   */
  HEART__BEAT__BATTERY__CODE__HBB_CYCLECOUNT_ERROR = 5,
  /*
   *电池接口类型错误
   */
  HEART__BEAT__BATTERY__CODE__HBB_INTERFACE_ERROR = 6,
  /*
   *电池机身被拆开
   */
  HEART__BEAT__BATTERY__CODE__HBB_BODY_ERROR = 7,
  /*
   *电池充电线被破坏
   */
  HEART__BEAT__BATTERY__CODE__HBB_LINE_ERROR = 8,
  /*
   *电池机身被拆开且充电线被破坏
   */
  HEART__BEAT__BATTERY__CODE__HBB_BODY_AND_LINE_ERROR = 9
    PROTOBUF_C__FORCE_ENUM_TO_BE_INT_SIZE(HEART__BEAT__BATTERY__CODE)
} HEARTBEATBATTERYCODE;
typedef enum _CHARGESTATUS {
  CHARGE__STATUS__UNCHARGE = 1,
  CHARGE__STATUS__CHARGING = 2,
  CHARGE__STATUS__DISCHARGING = 3
    PROTOBUF_C__FORCE_ENUM_TO_BE_INT_SIZE(CHARGE__STATUS)
} CHARGESTATUS;
typedef enum _ENABLEBORROWSTATUS {
  /*
   *可借状态（电池容量达到）
   */
  ENABLE__BORROW__STATUS__ENABLE_STATUS = 0,
  /*
   *不可借状态（电池容量没有达到）
   */
  ENABLE__BORROW__STATUS__DISABLE_STATUS = 1
    PROTOBUF_C__FORCE_ENUM_TO_BE_INT_SIZE(ENABLE__BORROW__STATUS)
} ENABLEBORROWSTATUS;
typedef enum _DISCHARGECMD {
  DISCHARGE__CMD__START_DISCHARGE = 1,
  DISCHARGE__CMD__STOP_DISCHARGE = 2
    PROTOBUF_C__FORCE_ENUM_TO_BE_INT_SIZE(DISCHARGE__CMD)
} DISCHARGECMD;
typedef enum _MSGTYPE {
  MSG__TYPE__UPLOAD_INFO = 1,
  MSG__TYPE__DISCHARGE_SETTING = 2
    PROTOBUF_C__FORCE_ENUM_TO_BE_INT_SIZE(MSG__TYPE)
} MSGTYPE;

/* --- messages --- */

struct  _BatteryAgeingInfo
{
  ProtobufCMessage base;
  char *timestamp;
  char *battery_sn;
  int32_t slot_num;
  int32_t voltage;
  int32_t current;
  int32_t temprature;
  int32_t elapsed;
  int32_t discharging;
  /*
   *电源温度
   */
  int32_t _temprature;
  /*
   *电源电压
   */
  int32_t _voltage;
  /*
   *电源最大容量（mAh)
   */
  int32_t _fullchargecapacity;
  /*
   *剩余容量（mAh)
   */
  int32_t _remainingcapacity;
  /*
   *充放电电流(mA)
   */
  int32_t _averagecurrent;
  /*
   *循环次数
   */
  int32_t _cyclecount;
  /*
   *异常状态，预留，BMS内部状态
   */
  int32_t _bmssafetystatus;
  /*
   *预留，BMS内部充满标志
   */
  int32_t _bmsflags;
  /*
   *异常状态
   */
  HEARTBEATBATTERYCODE _batterystatus;
  /*
   *充放电状态
   */
  CHARGESTATUS _chargestatus;
  /*
   *可借标志(容量达到一定才能允许用户借)
   */
  ENABLEBORROWSTATUS _enablestatus;
  /*
   *卡槽状态
   */
  int32_t _slotstatus;
  /*
   *电池被拆开
   */
  int32_t _destroyed;
  /*
   *带bms
   */
  int32_t _hasbms;
  /*
   *电池电量百分比
   */
  int32_t _radio;
};
#define BATTERY_AGEING_INFO__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&battery_ageing_info__descriptor) \
    , NULL, NULL, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 }


struct  _UploadInfo
{
  ProtobufCMessage base;
  size_t n_battery_ageing_info;
  BatteryAgeingInfo **battery_ageing_info;
};
#define UPLOAD_INFO__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&upload_info__descriptor) \
    , 0,NULL }


struct  _DischargeSetting
{
  ProtobufCMessage base;
  DISCHARGECMD cmd;
  int32_t level;
};
#define DISCHARGE_SETTING__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&discharge_setting__descriptor) \
    , 0, 0 }


struct  _MSGBODY
{
  ProtobufCMessage base;
  MSGTYPE type;
  UploadInfo *up_info;
  DischargeSetting *dis_setting;
};
#define MSG__BODY__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&msg__body__descriptor) \
    , 0, NULL, NULL }


/* BatteryAgeingInfo methods */
void   battery_ageing_info__init
                     (BatteryAgeingInfo         *message);
size_t battery_ageing_info__get_packed_size
                     (const BatteryAgeingInfo   *message);
size_t battery_ageing_info__pack
                     (const BatteryAgeingInfo   *message,
                      uint8_t             *out);
size_t battery_ageing_info__pack_to_buffer
                     (const BatteryAgeingInfo   *message,
                      ProtobufCBuffer     *buffer);
BatteryAgeingInfo *
       battery_ageing_info__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   battery_ageing_info__free_unpacked
                     (BatteryAgeingInfo *message,
                      ProtobufCAllocator *allocator);
/* UploadInfo methods */
void   upload_info__init
                     (UploadInfo         *message);
size_t upload_info__get_packed_size
                     (const UploadInfo   *message);
size_t upload_info__pack
                     (const UploadInfo   *message,
                      uint8_t             *out);
size_t upload_info__pack_to_buffer
                     (const UploadInfo   *message,
                      ProtobufCBuffer     *buffer);
UploadInfo *
       upload_info__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   upload_info__free_unpacked
                     (UploadInfo *message,
                      ProtobufCAllocator *allocator);
/* DischargeSetting methods */
void   discharge_setting__init
                     (DischargeSetting         *message);
size_t discharge_setting__get_packed_size
                     (const DischargeSetting   *message);
size_t discharge_setting__pack
                     (const DischargeSetting   *message,
                      uint8_t             *out);
size_t discharge_setting__pack_to_buffer
                     (const DischargeSetting   *message,
                      ProtobufCBuffer     *buffer);
DischargeSetting *
       discharge_setting__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   discharge_setting__free_unpacked
                     (DischargeSetting *message,
                      ProtobufCAllocator *allocator);
/* MSGBODY methods */
void   msg__body__init
                     (MSGBODY         *message);
size_t msg__body__get_packed_size
                     (const MSGBODY   *message);
size_t msg__body__pack
                     (const MSGBODY   *message,
                      uint8_t             *out);
size_t msg__body__pack_to_buffer
                     (const MSGBODY   *message,
                      ProtobufCBuffer     *buffer);
MSGBODY *
       msg__body__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   msg__body__free_unpacked
                     (MSGBODY *message,
                      ProtobufCAllocator *allocator);
/* --- per-message closures --- */

typedef void (*BatteryAgeingInfo_Closure)
                 (const BatteryAgeingInfo *message,
                  void *closure_data);
typedef void (*UploadInfo_Closure)
                 (const UploadInfo *message,
                  void *closure_data);
typedef void (*DischargeSetting_Closure)
                 (const DischargeSetting *message,
                  void *closure_data);
typedef void (*MSGBODY_Closure)
                 (const MSGBODY *message,
                  void *closure_data);

/* --- services --- */


/* --- descriptors --- */

extern const ProtobufCEnumDescriptor    heart__beat__battery__code__descriptor;
extern const ProtobufCEnumDescriptor    charge__status__descriptor;
extern const ProtobufCEnumDescriptor    enable__borrow__status__descriptor;
extern const ProtobufCEnumDescriptor    discharge__cmd__descriptor;
extern const ProtobufCEnumDescriptor    msg__type__descriptor;
extern const ProtobufCMessageDescriptor battery_ageing_info__descriptor;
extern const ProtobufCMessageDescriptor upload_info__descriptor;
extern const ProtobufCMessageDescriptor discharge_setting__descriptor;
extern const ProtobufCMessageDescriptor msg__body__descriptor;

PROTOBUF_C__END_DECLS


#endif  /* PROTOBUF_C_battery_5fageing_2eproto__INCLUDED */
