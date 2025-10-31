
#ifndef __pbox_ble_h__
#define __pbox_ble_h__

#include "pbox_module.h"

typedef struct PBoxBleModule_t
{

    PBoxModule module;

} PBoxBleModule;

PBoxModule *PBoxBleModule_init(PBoxBleModule *self);

#endif //  __pbox_ble_h__
