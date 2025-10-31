
#ifndef __pbox_strings_h__
#define __pbox_strings_h__

#include "pbox_types.h"
#include "pbox_array.h"
#include "pbox_errors.h"

typedef struct PBoxStringHolder_t
{

    PBoxArray array;

    PBoxString string;

} PBoxStringHolder;

void PBoxStringHolder_init(PBoxStringHolder *self);
PBoxError PBoxStringHolder_create(PBoxStringHolder *self, PBoxSize len);

#endif //  __pbox_strings_h__
