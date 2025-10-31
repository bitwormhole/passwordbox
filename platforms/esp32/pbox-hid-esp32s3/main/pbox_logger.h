
#ifndef __pbox_logger_h__
#define __pbox_logger_h__

#include "pbox_types.h"

typedef int PBoxLogLevel;

const static PBoxLogLevel PBoxLogMin = 0;
const static PBoxLogLevel PBoxLogTrace = 1;
const static PBoxLogLevel PBoxLogDebug = 2;
const static PBoxLogLevel PBoxLogInfo = 3;
const static PBoxLogLevel PBoxLogWarn = 4;
const static PBoxLogLevel PBoxLogError = 5;
const static PBoxLogLevel PBoxLogFatal = 6;
const static PBoxLogLevel PBoxLogMax = 7;

typedef struct PBoxLogItem_t
{
    PBoxLogLevel level;

    PBoxTime time;

    PBoxString message;

    PBoxString source;

    PBoxString tag;

} PBoxLogItem;

void PBoxLogger_push(PBoxLogItem *item);

void PBoxLogger_trace(PBoxString tag, PBoxString fmt, ...);
void PBoxLogger_debug(PBoxString tag, PBoxString fmt, ...);
void PBoxLogger_info(PBoxString tag, PBoxString fmt, ...);
void PBoxLogger_warn(PBoxString tag, PBoxString fmt, ...);
void PBoxLogger_error(PBoxString tag, PBoxString fmt, ...);
void PBoxLogger_fatal(PBoxString tag, PBoxString fmt, ...);

#endif //  __pbox_logger_h__
