
#ifndef __pbox_errors_h__
#define __pbox_errors_h__

#include "pbox_types.h"

////////////////////////////////////////////////////////////////////////////////

typedef struct PBoxErrorInfo_t
{

    pbox_int code;

    pbox_string message;

    pbox_string source;

    PBoxError parent;

} PBoxErrorInfo, *PBoxError, *pbox_error;

pbox_error pbox_make_error(pbox_int code, pbox_string src, pbox_string msg);

////////////////////////////////////////////////////////////////////////////////

typedef struct PBoxErrorHolder_t
{
    pbox_bool use1st; // (YES: 保留第一个err; NO:保留最后一个err)
    PBoxError err;

} PBoxErrorHolder;

void PBoxErrorHolder_init(PBoxErrorHolder *self, pbox_bool use1st);

void PBoxErrorHolder_handle(PBoxErrorHolder *self, PBoxError err);

pbox_bool PBoxErrorHolder_has_error(PBoxErrorHolder *self);

PBoxError PBoxErrorHolder_get_error(PBoxErrorHolder *self);

////////////////////////////////////////////////////////////////////////////////

#endif //  __pbox_errors_h__
