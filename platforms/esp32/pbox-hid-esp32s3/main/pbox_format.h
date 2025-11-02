
#ifndef __pbox_format_h__
#define __pbox_format_h__

#include "pbox_types.h"
#include "pbox_strings.h"

// typedef struct PBoxExample_t
// { } PBoxExample;

pbox_string pbox_format(PBoxStringBuffer *dst, pbox_string fmt, ...);

#endif //  __pbox_format_h__
