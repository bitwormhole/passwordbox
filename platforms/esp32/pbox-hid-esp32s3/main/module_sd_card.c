#include "module_sd_card.h"

PBoxModule *PBoxSDCardModule_init(PBoxSDCardModule *self)
{
    PBoxModule *m = &self->module;

    m->name = "PBoxSDCardModule";
    m->enabled = YES;

    return m;
}
