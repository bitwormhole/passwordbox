#include "module_usb_hid.h"

PBoxModule *PBoxUsbHidModule_init(PBoxUsbHidModule *self)
{
    PBoxModule *m = &self->module;

    m->name = "PBoxUsbHidModule";
    m->enabled = YES;

    return m;
}
