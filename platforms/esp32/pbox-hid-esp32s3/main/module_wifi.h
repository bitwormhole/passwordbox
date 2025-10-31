
#ifndef __pbox_wifi_h__
#define __pbox_wifi_h__

#include "pbox_module.h"

typedef struct PBoxWifiModule_t
{

    PBoxModule module;

} PBoxWifiModule;

PBoxModule *PBoxWifiModule_init(PBoxWifiModule *self);

#endif //  __pbox_wifi_h__
