
#ifndef __pbox_bytes_h__
#define __pbox_bytes_h__

#include "pbox_array.h"

typedef struct PBoxByteBuffer_t
{
    PBoxByteBuffer *self; // 指向结构本身的指针，用来确定这个结构的实例是否已经初始化

    PBoxArray array;

    PBoxByte *head;

} PBoxByteBuffer;

////////////////////////////////////////////////////////////////////////////////

void PBoxByteBuffer_init(PBoxByteBuffer *self);

PBoxBool PBoxByteBuffer_has_init(PBoxByteBuffer *self);

////////////////////////////////////////////////////////////////////////////////

#endif //  __pbox_bytes_h__
