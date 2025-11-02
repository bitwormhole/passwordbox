
#ifndef __pbox_logger_h__
#define __pbox_logger_h__

#include "pbox_types.h"

#define PBOX_LOG_TAG "pBox"

typedef enum PBoxLogLevel_t
{

    PBoxLogLevelMin,

    PBoxLogLevelTrace,
    PBoxLogLevelDebug,
    PBoxLogLevelInfo,
    PBoxLogLevelWarn,
    PBoxLogLevelError,
    PBoxLogLevelFatal,

    PBoxLogLevelMax,

} PBoxLogLevel;

typedef struct PBoxLogItem_t
{
    PBoxLogLevel level;

    pbox_time time;

    pbox_string message;

    pbox_string source;

    pbox_string tag;

} PBoxLogItem;

void PBoxLogger_push(PBoxLogItem *item);

void pbox_log_trace(pbox_string src, pbox_string fmt, ...);
void pbox_log_debug(pbox_string src, pbox_string fmt, ...);
void pbox_log_info(pbox_string src, pbox_string fmt, ...);
void pbox_log_warn(pbox_string src, pbox_string fmt, ...);
void pbox_log_error(pbox_string src, pbox_string fmt, ...);
void pbox_log_fatal(pbox_string src, pbox_string fmt, ...);

#endif //  __pbox_logger_h__
