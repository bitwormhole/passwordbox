
#ifndef __pbox_app_context_h__
#define __pbox_app_context_h__

#include "pbox_types.h"

typedef struct PBoxAppContext_t
{
    PBoxApp *app;

    PBoxAutoTestModule *auto_test;
    PBoxBleModule *ble;
    PBoxUsbHidModule *usb_hid;
    PBoxSDCardModule *sd_card;
    PBoxWifiModule *wifi;

} PBoxAppContext;

void PBoxAppContext_init(PBoxAppContext *self);

#endif //  __pbox_app_context_h__
