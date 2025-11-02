#include "pbox_module.h"

#include "pbox_logger.h"

#include <memory.h>

////////////////////////////////////////////////////////////////////////////////
// module-holder

typedef struct PBoxModuleHolder_t
{

    PBoxModule *module;

} PBoxModuleHolder;

////////////////////////////////////////////////////////////////////////////////
// manager-inner

#define PBoxModuleManagerInner_capacity 8

typedef struct PBoxModuleManagerInner_t
{

    PBoxModuleManager *owner;

    pbox_uint capacity;
    pbox_uint count;

    PBoxModuleHolder holders[PBoxModuleManagerInner_capacity];

} PBoxModuleManagerInner;

void PBoxModuleManagerInner_init(PBoxModuleManagerInner *self);

void PBoxModuleManagerInner_on_destroy(PBoxModuleManagerInner *self);

////////////////////////////////////////////////////////////////////////////////

void PBoxModuleManagerInner_init(PBoxModuleManagerInner *self)
{
    if (self)
    {
        memset(self, 0, sizeof(self[0]));

        self->capacity = PBoxModuleManagerInner_capacity;
    }
}

void PBoxModuleManagerInner_on_destroy(PBoxModuleManagerInner *self)
{
    if (self)
    {
        memset(self, 0, sizeof(self[0]));
    }
}

////////////////////////////////////////////////////////////////////////////////

void PBoxModuleManager_init(PBoxModuleManager *self)
{
    if (self)
    {
        if (!PBoxModuleManager_is_ready(self))
        {
            memset(self, 0, sizeof(self[0]));
        }
    }
}

PBoxError PBoxModuleManager_create(PBoxModuleManager *self)
{
    if (self == NIL)
    {
        return pbox_make_error(500, "PBoxModuleManager_create", "self is nil");
    }

    if (PBoxModuleManager_is_ready(self))
    {
        return NIL;
    }
    else
    {
        PBoxModuleManager_init(self);
    }

    pbox_size size = sizeof(PBoxModuleManagerInner);
    PBoxModuleManagerInner *inner = malloc(size);
    PBoxModuleManagerInner_init(inner);

    self->inner = inner;
    inner->owner = self;

    return NIL;
}

PBoxError PBoxModuleManager_destroy(PBoxModuleManager *self)
{
    if (self)
    {
        PBoxModuleManagerInner *inner = self->inner;
        self->inner = NIL;
        if (inner)
        {
            PBoxModuleManagerInner_on_destroy(inner);
            free(inner);
        }
    }
    return NIL;
}

pbox_bool PBoxModuleManager_is_ready(PBoxModuleManager *self)
{
    if (!self)
    {
        return NO;
    }

    PBoxModuleManagerInner *inner = self->inner;
    if (!inner)
    {
        return NO;
    }

    PBoxModuleManager *owner = inner->owner;
    if (!owner)
    {
        return NO;
    }

    return (owner == self);
}

PBoxModuleIterator *PBoxModuleManager_iterate(PBoxModuleManager *self, PBoxModuleIterator *iter, pbox_bool reverse)
{
    if ((self == NIL) || (iter == NIL))
    {
        return NIL;
    }

    memset(iter, 0, sizeof(iter[0]));

    iter->inner = self->inner;
    iter->current = 0;
    iter->reverse = reverse;

    return iter;
}

PBoxError PBoxModuleManager_add(PBoxModuleManager *self, PBoxModule *m)
{
    if (m == NIL)
    {
        return pbox_make_error(500, "PBoxModuleManager_add", "param:module is nil");
    }

    if (!PBoxModuleManager_is_ready(self))
    {
        return pbox_make_error(500, "PBoxModuleManager_add", "manager is not ready");
    }

    PBoxModuleManagerInner *inner = self->inner;
    int i0 = 0;
    int i1 = inner->capacity;
    int idx = inner->count;
    PBoxModuleHolder *array = inner->holders;

    if ((i0 <= idx) && (idx < i1))
    {
        array[idx].module = m;
        inner->count++;
        return NIL;
    }

    return pbox_make_error(500, "PBoxModuleManager_add", "manager.array is overflow");
}

void PBoxModuleManager_add_with_eh(PBoxModuleManager *self, PBoxModule *m, PBoxErrorHolder *eh)
{
    PBoxError err = PBoxModuleManager_add(self, m);
    if (err)
    {
        PBoxErrorHolder_handle(eh, err);
    }
}

////////////////////////////////////////////////////////////////////////////////
// impl : PBoxModuleIterator

PBoxModule *PBoxModuleIterator_next(PBoxModuleIterator *self)
{
    if (self == NIL)
    {
        return NO;
    }
    PBoxModuleManagerInner *inner = self->inner;
    PBoxModuleManager *man = inner->owner;
    if (!PBoxModuleManager_is_ready(man))
    {
        return NO;
    }

    const int idx_begin = 0;
    const int idx_end = inner->count - 1;
    int idx = self->current;
    PBoxModuleHolder *const holders = inner->holders;

    if (self->reverse)
    {
        for (; idx_begin <= idx && idx <= idx_end; idx--)
        {
            PBoxModuleHolder *const h = holders + idx;
            if (h->module)
            {
                self->current = idx - 1;
                return h->module;
            }
        }
    }
    else
    {
        for (; idx_begin <= idx && idx <= idx_end; idx++)
        {
            PBoxModuleHolder *const h = holders + idx;
            if (h->module)
            {
                self->current = idx + 1;
                return h->module;
            }
        }
    }

    self->current = idx;
    return NIL;
}

pbox_bool PBoxModuleIterator_has_more(PBoxModuleIterator *self)
{
    if (self == NIL)
    {
        return NO;
    }
    PBoxModuleManagerInner *inner = self->inner;
    PBoxModuleManager *man = inner->owner;
    if (!PBoxModuleManager_is_ready(man))
    {
        return NO;
    }

    int idx_begin = 0;
    int idx_end = inner->count - 1;
    int idx = self->current;

    return ((idx_begin <= idx) && (idx <= idx_end));
}

////////////////////////////////////////////////////////////////////////////////
// for PBoxModule

PBoxModuleLifeFunc PBoxModule_get_lifecycle_func(PBoxModule *self, PBoxModuleLifePhase sel)
{
    if (self == NIL)
    {
        return NIL;
    }
    PBoxModuleLifeFunc fn = NIL;
    switch (sel)
    {
    case PBoxModuleLifePhaseOnCreate:
        fn = self->on_create;
        break;
    case PBoxModuleLifePhaseOnStart:
        fn = self->on_start;
        break;
    case PBoxModuleLifePhaseOnResume:
        fn = self->on_resume;
        break;
    case PBoxModuleLifePhaseOnRun:
        fn = self->on_run;
        break;
    case PBoxModuleLifePhaseOnPause:
        fn = self->on_pause;
        break;
    case PBoxModuleLifePhaseOnStop:
        fn = self->on_stop;
        break;
    case PBoxModuleLifePhaseOnDestroy:
        fn = self->on_destroy;
        break;
    default:
        break;
    }
    return fn;
}

PBoxError PBoxModule_invoke_lifecycle_func(PBoxModule *self, PBoxAppContext *ac, PBoxModuleLifePhase sel)
{
    PBoxModuleLifeFunc fn = PBoxModule_get_lifecycle_func(self, sel);
    PBoxModuleLifeCallback callback;

    if (fn && ac)
    {
        if (!self->enabled)
        {
            return NIL;
        }
        pbox_string sel_str = PBoxModuleLifePhase_stringify(sel);
        pbox_string mod_name = self->name;
        pbox_log_info("PBoxModule_invoke_lifecycle_func", "invoke: %s@%s", sel_str, mod_name);

        memset(&callback, 0, sizeof(callback));
        callback.context = ac;
        callback.module = self;
        callback.phase = sel;

        return fn(&callback);
    }
    return NIL;
}

pbox_string PBoxModuleLifePhase_stringify(PBoxModuleLifePhase sel)
{
    pbox_string str = "";
    switch (sel)
    {
    case PBoxModuleLifePhaseOnCreate:
        str = "on_create";
        break;
    case PBoxModuleLifePhaseOnStart:
        str = "on_start";
        break;
    case PBoxModuleLifePhaseOnResume:
        str = "on_resume";
        break;
    case PBoxModuleLifePhaseOnRun:
        str = "on_run";
        break;
    case PBoxModuleLifePhaseOnPause:
        str = "on_pause";
        break;
    case PBoxModuleLifePhaseOnStop:
        str = "on_stop";
        break;
    case PBoxModuleLifePhaseOnDestroy:
        str = "on_destroy";
        break;
    default:
        break;
    }
    return str;
}

////////////////////////////////////////////////////////////////////////////////
// EOF
