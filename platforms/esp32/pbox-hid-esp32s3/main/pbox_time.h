
#ifndef __pbox_time_h__
#define __pbox_time_h__

#include "pbox_types.h"

// typedef struct PBoxExample_t
// {
// } PBoxExample;

PBoxTime PBoxTime_now();

PBoxTime PBoxTime_add(PBoxTime t, PBoxTimeSpan span);

PBoxTimeSpan PBoxTime_span(PBoxTime t1, PBoxTime t2);

PBoxTimeSpan PBoxTime_span_abs(PBoxTime t1, PBoxTime t2);

#endif //  __pbox_time_h__
