
#ifndef __pbox_sd_card_h__
#define __pbox_sd_card_h__

#include "pbox_module.h"

typedef struct PBoxSDCardModule_t
{

    PBoxModule module;

} PBoxSDCardModule;

PBoxModule *PBoxSDCardModule_init(PBoxSDCardModule *self);

#endif //  __pbox_sd_card_h__
