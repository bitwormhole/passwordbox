#include "module_ble.h"

PBoxModule *PBoxBleModule_init(PBoxBleModule *self)
{
    PBoxModule *m = &self->module;

    m->name = "PBoxBleModule";
    m->enabled = YES;

    return m;
}
