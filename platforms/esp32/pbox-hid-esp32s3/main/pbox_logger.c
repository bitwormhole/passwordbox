
#include "pbox_logger.h"
#include <memory.h>

#include "esp_log.h"

void PBoxLogger_push(PBoxLogItem *item)
{
    // todo : impl
}

void PBoxLogger_push_params(PBoxString tag, PBoxLogLevel level, PBoxTime time, PBoxString msg, PBoxString src)
{
    // todo : impl

    PBoxLogItem item;
    memset(&item, 0, sizeof(item));

    item.level = level;
    item.tag = tag;
    item.message = msg;

    PBoxLogger_push(&item);
}

void PBoxLogger_trace(PBoxString tag, PBoxString fmt, ...)
{
    // PBoxTime time = 0;
    // PBoxLog_push_params(tag, PBoxLogTrace, time, msg, NIL);

    ESP_LOGV(tag, "%s", fmt);
}

void PBoxLogger_debug(PBoxString tag, PBoxString fmt, ...)
{
    ESP_LOGD(tag, "%s", fmt);
}

void PBoxLogger_info(PBoxString tag, PBoxString fmt, ...)
{
    ESP_LOGI(tag, "%s", fmt);
}

void PBoxLogger_warn(PBoxString tag, PBoxString fmt, ...)
{
    ESP_LOGW(tag, "%s", fmt);
}

void PBoxLogger_error(PBoxString tag, PBoxString fmt, ...)
{
    ESP_LOGE(tag, "%s", fmt);
}

void PBoxLogger_fatal(PBoxString tag, PBoxString fmt, ...)
{
    ESP_LOGE(tag, "%s", fmt);
}
