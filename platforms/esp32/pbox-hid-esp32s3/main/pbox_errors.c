

#include "pbox_errors.h"

#include <memory.h>

#define PBoxErrorPool_capacity 8

////////////////////////////////////////////////////////////////////////////////
// holder

typedef struct PBoxErrorHolder2_t
{

    PBoxErrorInfo info;

} PBoxErrorHolder2;

////////////////////////////////////////////////////////////////////////////////
// pool

typedef struct PBoxErrorPool_t
{

    PBoxErrorHolder2 holders[PBoxErrorPool_capacity];

    pbox_uint count;
    pbox_uint capacity;

} PBoxErrorPool;

PBoxErrorPool *PBoxErrorPool_get_instance();

void PBoxErrorPool_init(PBoxErrorPool *self);

PBoxErrorHolder2 *PBoxErrorPool_get_next_holder(PBoxErrorPool *self);

////////////////////////////////////////////////////////////////////////////////
// impl: pool

static PBoxErrorPool *PBoxErrorPool_the_instance = NIL;

PBoxErrorPool *PBoxErrorPool_get_instance()
{
    PBoxErrorPool *pool = PBoxErrorPool_the_instance;
    if (pool == NIL)
    {
        pbox_size size = sizeof(PBoxErrorPool);
        pool = malloc(size);
        PBoxErrorPool_init(pool);
        PBoxErrorPool_the_instance = pool;
    }
    return pool;
}

void PBoxErrorPool_init(PBoxErrorPool *self)
{
    if (self)
    {
        memset(self, 0, sizeof(self[0]));
        self->capacity = PBoxErrorPool_capacity;
    }
}

PBoxErrorHolder2 *PBoxErrorPool_get_next_holder(PBoxErrorPool *self)
{
    if (self)
    {
        self->count++;
        pbox_uint index = self->count;
        pbox_uint capacity = self->capacity;
        PBoxErrorHolder2 *holder = self->holders + (index % capacity);
        return holder;
    }
    return NIL;
}

////////////////////////////////////////////////////////////////////////////////
// impl: PBoxErrorHolder

void PBoxErrorHolder_init(PBoxErrorHolder *self, pbox_bool use1st)
{
    if (self)
    {
        memset(self, 0, sizeof(self[0]));
        self->use1st = use1st;
    }
}

void PBoxErrorHolder_handle(PBoxErrorHolder *self, PBoxError err)
{
    if (self == NIL || err == NIL)
    {
        return;
    }

    if (self->use1st && self->err)
    {
        return; // 已经存在一个错误
    }

    self->err = err;
}

pbox_bool PBoxErrorHolder_has_error(PBoxErrorHolder *self)
{
    PBoxError err = PBoxErrorHolder_get_error(self);
    return (err ? YES : NO);
}

PBoxError PBoxErrorHolder_get_error(PBoxErrorHolder *self)
{
    if (self)
    {
        return self->err;
    }
    return NIL;
}

////////////////////////////////////////////////////////////////////////////////
// api: error

pbox_error pbox_make_error(pbox_int code, pbox_string src, pbox_string msg)
{
    PBoxErrorPool *pool = PBoxErrorPool_get_instance();
    PBoxErrorHolder2 *holder = PBoxErrorPool_get_next_holder(pool);

    if (holder)
    {
        PBoxError err = &holder->info;
        err->code = code;
        err->message = msg;
        err->source = src;
        err->parent = NIL;
        return err;
    }

    return NIL;
}

////////////////////////////////////////////////////////////////////////////////
// EOF
