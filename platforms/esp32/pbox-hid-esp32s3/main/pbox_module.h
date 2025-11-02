
#ifndef __pbox_module_h__
#define __pbox_module_h__

#include "pbox_errors.h"

////////////////////////////////////////////////////////////////////////////////

typedef struct PBoxModule_t PBoxModule;

typedef struct PBoxModuleManager_t PBoxModuleManager;

typedef struct PBoxModuleLifeCallback_t PBoxModuleLifeCallback;

typedef struct PBoxModuleManagerInner_t PBoxModuleManagerInner;

typedef PBoxError (*PBoxModuleLifeFunc)(PBoxModuleLifeCallback *callback);

////////////////////////////////////////////////////////////////////////////////

typedef enum PBoxModuleLifePhase_t
{

    PBoxModuleLifePhaseOnMin,
    PBoxModuleLifePhaseOnInit,

    PBoxModuleLifePhaseOnCreate,
    PBoxModuleLifePhaseOnStart,
    PBoxModuleLifePhaseOnResume,
    PBoxModuleLifePhaseOnRun,
    PBoxModuleLifePhaseOnPause,
    PBoxModuleLifePhaseOnStop,
    PBoxModuleLifePhaseOnDestroy,

    PBoxModuleLifePhaseOnMax,

} PBoxModuleLifePhase;

pbox_string PBoxModuleLifePhase_stringify(PBoxModuleLifePhase sel);

////////////////////////////////////////////////////////////////////////////////

typedef struct PBoxModule_t
{

    pbox_string name;
    pbox_bool enabled;

    PBoxModuleLifeFunc on_create;
    PBoxModuleLifeFunc on_start;
    PBoxModuleLifeFunc on_resume;
    PBoxModuleLifeFunc on_run;
    PBoxModuleLifeFunc on_pause;
    PBoxModuleLifeFunc on_stop;
    PBoxModuleLifeFunc on_destroy;

} PBoxModule;

typedef struct PBoxModuleLifeCallback_t
{

    PBoxModuleLifePhase phase;

    PBoxAppContext *context;

    PBoxModule *module;

} PBoxModuleLifeCallback;

typedef struct PBoxModuleIterator_t
{
    PBoxModuleManagerInner *inner;

    int current;

    pbox_bool reverse;

} PBoxModuleIterator;

typedef struct PBoxModuleManager_t
{

    PBoxModuleManagerInner *inner;

} PBoxModuleManager;

////////////////////////////////////////////////////////////////////////////////
// func of PBoxModule

PBoxModuleLifeFunc PBoxModule_get_lifecycle_func(PBoxModule *self, PBoxModuleLifePhase sel);

PBoxError PBoxModule_invoke_lifecycle_func(PBoxModule *self, PBoxAppContext *ac, PBoxModuleLifePhase sel);

////////////////////////////////////////////////////////////////////////////////
// func of PBoxModuleIterator

PBoxModule *PBoxModuleIterator_next(PBoxModuleIterator *self);

pbox_bool PBoxModuleIterator_has_more(PBoxModuleIterator *self);

////////////////////////////////////////////////////////////////////////////////
// func of PBoxModuleManager

void PBoxModuleManager_init(PBoxModuleManager *self);

PBoxError PBoxModuleManager_create(PBoxModuleManager *self);

PBoxError PBoxModuleManager_destroy(PBoxModuleManager *self);

pbox_bool PBoxModuleManager_is_ready(PBoxModuleManager *self);

PBoxError PBoxModuleManager_add(PBoxModuleManager *self, PBoxModule *m);

void PBoxModuleManager_add_with_eh(PBoxModuleManager *self, PBoxModule *m, PBoxErrorHolder *eh);

PBoxModuleIterator *PBoxModuleManager_iterate(PBoxModuleManager *self, PBoxModuleIterator *iter, pbox_bool reverse);

////////////////////////////////////////////////////////////////////////////////

#endif //  __pbox_module_h__
