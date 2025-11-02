
#ifndef __pbox_types_h__
#define __pbox_types_h__

// simple types

// int

typedef int pbox_int;

typedef char pbox_int8;
typedef short pbox_int16;
typedef long pbox_int32;
typedef long long pbox_int64;

// uint

typedef unsigned int pbox_uint;

typedef unsigned char pbox_uint8;
typedef unsigned short pbox_uint16;
typedef unsigned long pbox_uint32;
typedef unsigned long long pbox_uint64;

// float

// others

typedef unsigned char pbox_bool; // 表示 1 个布尔值， （0:false ; 非0:true ）

typedef unsigned char pbox_byte; // 表示 1 字节

typedef char pbox_char; // 表示 1 字符

typedef int pbox_size; // 表示数据占用内存的大小，单位是字节

typedef int pbox_count; // 表示数据单元的个数， 单位是‘个’

// time

typedef pbox_int64 pbox_time;      // 时间戳 （unit:ms）
typedef pbox_int64 pbox_time_span; // 时间长度（unit:ms）

// array ptr

typedef const unsigned char *pbox_bytes;

typedef const char *pbox_string;

////////////////////////////////////////////////////////////////////////////////

// virtual ptr

typedef struct PBoxErrorInfo_t *PBoxError, *pbox_error;

typedef struct PBoxByteBuffer_t PBoxByteBuffer;

typedef struct PBoxStringHolder_t PBoxStringHolder;

typedef struct PBoxAutoTestModule_t PBoxAutoTestModule;
typedef struct PBoxBleModule_t PBoxBleModule;
typedef struct PBoxUsbHidModule_t PBoxUsbHidModule;
typedef struct PBoxSDCardModule_t PBoxSDCardModule;
typedef struct PBoxWifiModule_t PBoxWifiModule;

typedef struct PBoxAppContext_t PBoxAppContext;
typedef struct PBoxApp_t PBoxApp;

////////////////////////////////////////////////////////////////////////////////

// const

#define YES 1
#define NO 0
#define NIL 0

#define PBOX_MIN(a, b) (((a) < (b)) ? (a) : (b))
#define PBOX_MAX(a, b) (((a) > (b)) ? (a) : (b))

////////////////////////////////////////////////////////////////////////////////

#endif //  __pbox_types_h__
