
#ifndef __pbox_types_h__
#define __pbox_types_h__

// simple types

// int

typedef int PBoxInt;

typedef char PBoxInt8;
typedef short PBoxInt16;
typedef long PBoxInt32;
typedef long long PBoxInt64;

// uint

typedef unsigned int PBoxUint;

typedef unsigned char PBoxUint8;
typedef unsigned short PBoxUint16;
typedef unsigned long PBoxUint32;
typedef unsigned long long PBoxUint64;

// float

// others

typedef unsigned char PBoxBool; // 表示 1 个布尔值， （0:false ; 非0:true ）

typedef unsigned char PBoxByte; // 表示 1 字节

typedef char PBoxChar; // 表示 1 字符

typedef int PBoxSize; // 表示数据占用内存的大小，单位是字节

// time

typedef PBoxInt64 PBoxTime;     // 时间戳 （unit:ms）
typedef PBoxInt64 PBoxTimeSpan; // 时间长度（unit:ms）

// array ptr

typedef const unsigned char *PBoxBytes;

typedef const char *PBoxString;

////////////////////////////////////////////////////////////////////////////////

// virtual ptr

typedef struct PBoxByteBuffer_t PBoxByteBuffer;

typedef struct PBoxStringHolder_t PBoxStringHolder;

typedef struct PBoxError_t *PBoxError;

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

////////////////////////////////////////////////////////////////////////////////

#endif //  __pbox_types_h__
