
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
    PBoxArrayManager *manager;
    PBoxArrayHolder *owner;

    pbox_size capacity;   // 总大小
    pbox_size unit_size;  // 每个条目的大小
    pbox_count count;     // 当前包含的条目数量
    pbox_count count_max; // 数组中可包含条目的最大数量

    pbox_byte data[1];

} PBoxArrayEntity;

typedef struct PBoxArrayHolder_t
{
    PBoxArrayManager *manager;
    PBoxArrayEntity *entity;

    void *head; // 指向数据缓冲区的开头

} PBoxArrayHolder, PBoxArray;

typedef struct PBoxArrayManager_t
{

    PBoxArrayManagerInner *inner;

} PBoxArrayManager;

////////////////////////////////////////////////////////////////////////////////

void PBoxArray_init(PBoxArray *self);

PBoxError PBoxArray_create(PBoxArray *self, pbox_size unit, pbox_uint cnt_max);
PBoxError PBoxArray_destroy(PBoxArray *self);
pbox_bool PBoxArray_is_ready(PBoxArray *self);

pbox_uint PBoxArray_get_count(PBoxArray *self);
pbox_uint PBoxArray_get_count_max(PBoxArray *self);
pbox_size PBoxArray_get_capacity(PBoxArray *self);
pbox_size PBoxArray_get_unit_size(PBoxArray *self);

////////////////////////////////////////////////////////////////////////////////

void PBoxArrayManager_init(PBoxArrayManager *self);
PBoxError PBoxArrayManager_create(PBoxArrayManager *self);
PBoxError PBoxArrayManager_destroy(PBoxArrayManager *self);
pbox_bool PBoxArrayManager_is_ready(PBoxArrayManager *self);
void PBoxArrayManager_run_gc(PBoxArrayManager *self);

PBoxArrayEntity *PBoxArrayManager_create_new_entity(PBoxArrayManager *self, PBoxArrayEntity *params);

PBoxArrayManager *PBoxArrayManager_get_instance();

////////////////////////////////////////////////////////////////////////////////

void PBoxArrayEntity_init_params(PBoxArrayEntity *self);

////////////////////////////////////////////////////////////////////////////////

#endif //  __pbox_array_h__
