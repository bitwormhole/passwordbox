
#include "module_wifi.h"

PBoxModule *PBoxWifiModule_init(PBoxWifiModule *self)
{
    if (self == NIL)
    {
        return NIL;
    }
    PBoxModule *m2 = &self->module;
    return m2;
}
