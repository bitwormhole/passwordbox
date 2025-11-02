
#ifndef __pbox_strings_h__
#define __pbox_strings_h__

#include "pbox_types.h"
#include "pbox_array.h"
#include "pbox_errors.h"

typedef struct PBoxStringHolder_t
{

    PBoxArray array;

    pbox_string string;

} PBoxStringHolder, PBoxStringBuilder, PBoxStringBuffer, PBoxString;

void PBoxStringBuffer_init(PBoxStringBuffer *self);

PBoxError PBoxStringBuffer_create(PBoxStringBuffer *self, pbox_size len);

#endif //  __pbox_strings_h__
