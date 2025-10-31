

#include "pbox_app_context.h"

#include <memory.h>

void PBoxAppContext_init(PBoxAppContext *self)
{
    if (self)
    {
        memset(self, 0, sizeof(self[0]));
    }
}
