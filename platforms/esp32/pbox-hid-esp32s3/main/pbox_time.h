
#ifndef __pbox_time_h__
#define __pbox_time_h__

#include "pbox_types.h"

// typedef struct PBoxExample_t
// {
// } PBoxExample;

pbox_time PBoxTime_now();

pbox_time PBoxTime_add(pbox_time t, pbox_time_span span);

pbox_time_span PBoxTime_span(pbox_time t1, pbox_time t2);

pbox_time_span PBoxTime_span_abs(pbox_time t1, pbox_time t2);

#endif //  __pbox_time_h__
