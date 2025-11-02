#include "module_auto_test.h"

#include "pbox_logger.h"

#include <memory.h>

////////////////////////////////////////////////////////////////////////////////

typedef struct PBoxAutoTestModuleInner_t
{

    int foo;
    char bar;

} PBoxAutoTestModuleInner;

PBoxError PBoxAutoTestModuleInner_init(PBoxAutoTestModuleInner *self);

PBoxError PBoxAutoTestModuleInner_on_run(PBoxModuleLifeCallback *callback);

////////////////////////////////////////////////////////////////////////////////

PBoxModule *PBoxAutoTestModule_init(PBoxAutoTestModule *self)
{
    PBoxModule *m = &self->module;
    PBoxAutoTestModuleInner *inner = NIL;

    inner = malloc(sizeof(PBoxAutoTestModuleInner));
    PBoxAutoTestModuleInner_init(inner);

    m->name = "PBoxAutoTestModule";
    m->enabled = YES;

    m->on_run = PBoxAutoTestModuleInner_on_run;

    self->inner = inner;

    return m;
}

////////////////////////////////////////////////////////////////////////////////

PBoxError PBoxAutoTestModuleInner_init(PBoxAutoTestModuleInner *self)
{
    if (self)
    {
        memset(self, 0, sizeof(self[0]));
    }
    return NIL;
}

PBoxError PBoxAutoTestModuleInner_on_run(PBoxModuleLifeCallback *callback)
{

    pbox_log_warn("PBoxAutoTestModuleInner_on_run", "callback");

    return NIL;
}

////////////////////////////////////////////////////////////////////////////////
// EOF
