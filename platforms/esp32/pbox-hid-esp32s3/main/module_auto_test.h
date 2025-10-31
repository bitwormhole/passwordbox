
#ifndef __pbox_auto_test_h__
#define __pbox_auto_test_h__

#include "pbox_module.h"

typedef struct PBoxAutoTestModuleInner_t PBoxAutoTestModuleInner;

////////////////////////////////////////////////////////////////////////////////

typedef struct PBoxAutoTestModule_t
{
    PBoxModule module;

    PBoxAutoTestModuleInner *inner;

} PBoxAutoTestModule;

PBoxModule *PBoxAutoTestModule_init(PBoxAutoTestModule *self);

#endif //  __pbox_auto_test_h__
