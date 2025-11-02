
#include "pbox_logger.h"
#include <memory.h>

#include "esp_log.h"

void PBoxLogger_push(PBoxLogItem *item)
{
    // todo : impl
}

void PBoxLogger_push_params(pbox_string tag, PBoxLogLevel level, pbox_time time, pbox_string msg, pbox_string src)
{
    // todo : impl

    PBoxLogItem item;
    memset(&item, 0, sizeof(item));

    item.level = level;
    item.tag = tag;
    item.message = msg;

    PBoxLogger_push(&item);
}

void pbox_log_trace(pbox_string src, pbox_string fmt, ...)
{
    // pbox_time time = 0;
    // PBoxLog_push_params(tag, PBoxLogTrace, time, msg, NIL);

    ESP_LOGV(src, "%s", fmt);
}

void pbox_log_debug(pbox_string src, pbox_string fmt, ...)
{
    ESP_LOGD(src, "%s", fmt);
}

void pbox_log_info(pbox_string src, pbox_string fmt, ...)
{
    ESP_LOGI(src, "%s", fmt);
}

void pbox_log_warn(pbox_string src, pbox_string fmt, ...)
{
    ESP_LOGW(src, "%s", fmt);
}

void pbox_log_error(pbox_string src, pbox_string fmt, ...)
{
    ESP_LOGE(src, "%s", fmt);
}

void pbox_log_fatal(pbox_string src, pbox_string fmt, ...)
{
    ESP_LOGE(src, "%s", fmt);
}
