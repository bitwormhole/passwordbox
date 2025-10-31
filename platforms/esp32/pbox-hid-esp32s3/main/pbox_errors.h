
#ifndef __pbox_errors_h__
#define __pbox_errors_h__

#include "pbox_types.h"

////////////////////////////////////////////////////////////////////////////////

typedef struct PBoxError_t
{

    PBoxInt code;

    PBoxString message;

    PBoxString source;

    PBoxError parent;

} PBoxErrorInfo, *PBoxError;

PBoxError PBoxError_make(PBoxInt code, PBoxString src, PBoxString msg);

////////////////////////////////////////////////////////////////////////////////

typedef struct PBoxErrorHolder_t
{
    PBoxBool use1st; // (YES: 保留第一个err; NO:保留最后一个err)
    PBoxError err;

} PBoxErrorHolder;

void PBoxErrorHolder_init(PBoxErrorHolder *self, PBoxBool use1st);

void PBoxErrorHolder_handle(PBoxErrorHolder *self, PBoxError err);

PBoxBool PBoxErrorHolder_has_error(PBoxErrorHolder *self);

PBoxError PBoxErrorHolder_get_error(PBoxErrorHolder *self);

////////////////////////////////////////////////////////////////////////////////

#endif //  __pbox_errors_h__
