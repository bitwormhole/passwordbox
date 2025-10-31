
#ifndef __pbox_usb_hid_h__
#define __pbox_usb_hid_h__

#include "pbox_module.h"

typedef struct PBoxUsbHidModule_t
{

    PBoxModule module;

} PBoxUsbHidModule;

PBoxModule *PBoxUsbHidModule_init(PBoxUsbHidModule *self);

#endif //  __pbox_usb_hid_h__
