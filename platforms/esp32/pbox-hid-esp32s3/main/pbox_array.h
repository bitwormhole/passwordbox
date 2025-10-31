
#ifndef __pbox_array_h__
#define __pbox_array_h__

#include "pbox_types.h"
#include "pbox_errors.h"

typedef struct PBoxArrayEntity_t PBoxArrayEntity;

typedef struct PBoxArrayHolder_t PBoxArrayHolder;

typedef struct PBoxArrayManager_t PBoxArrayManager;

typedef struct PBoxArrayManagerInner_t PBoxArrayManagerInner;

////////////////////////////////////////////////////////////////////////////////

typedef struct PBoxArrayEntity_t
{
    PBoxArrayHolder *owner;

    PBoxSize capacity;  // 总大小
    PBoxSize unit_size; // 每个条目的大小
    PBoxUint count;     // 当前包含的条目数量
    PBoxUint count_max; // 数组中可包含条目的最大数量

    char data[0];

} PBoxArrayEntity;

typedef struct PBoxArrayHolder_t
{
    PBoxArrayEntity *entity;
    PBoxArrayManager *manager;

    void *head; // 指向数据缓冲区的开头

} PBoxArrayHolder, PBoxArray;

typedef struct PBoxArrayManager_t
{

    PBoxArrayManagerInner *inner;

} PBoxArrayManager;

////////////////////////////////////////////////////////////////////////////////

void PBoxArray_init(PBoxArray *self);

PBoxError PBoxArray_create(PBoxArray *self, PBoxSize unit, PBoxUint cnt_max);
PBoxError PBoxArray_destroy(PBoxArray *self);
PBoxBool PBoxArray_is_ready(PBoxArray *self);

PBoxUint PBoxArray_get_count(PBoxArray *self);
PBoxUint PBoxArray_get_count_max(PBoxArray *self);
PBoxSize PBoxArray_get_capacity(PBoxArray *self);
PBoxSize PBoxArray_get_unit_size(PBoxArray *self);

////////////////////////////////////////////////////////////////////////////////

void PBoxArrayManager_init(PBoxArrayManager *self);
PBoxError PBoxArrayManager_create(PBoxArrayManager *self);
PBoxError PBoxArrayManager_destroy(PBoxArrayManager *self);
PBoxBool PBoxArrayManager_is_ready(PBoxArrayManager *self);
void PBoxArrayManager_run_gc(PBoxArrayManager *self);

PBoxArrayEntity *PBoxArrayManager_create_new_entity(PBoxArrayManager *self, PBoxArrayEntity *params);

////////////////////////////////////////////////////////////////////////////////

#endif //  __pbox_array_h__
