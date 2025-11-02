

#include "pbox_array.h"

#include "pbox_logger.h"
#include "esp_log.h"

#include <memory.h>

typedef struct PBoxArrayEntityNode_t PBoxArrayEntityNode;

////////////////////////////////////////////////////////////////////////////////
// type: PBoxArrayManagerInner

typedef struct PBoxArrayManagerInner_t
{

    PBoxArrayManager *outer;

    PBoxArrayEntityNode *all;

} PBoxArrayManagerInner;

PBoxArrayManagerInner *PBoxArrayManagerInner_create_new_instance();

////////////////////////////////////////////////////////////////////////////////
//  type : PBoxArrayEntityNode

typedef struct PBoxArrayEntityNode_t
{
    pbox_size total_size;

    PBoxArrayEntityNode *full_chain_prev;
    PBoxArrayEntityNode *full_chain_next;
    PBoxArrayEntityNode *side_chain_prev;
    PBoxArrayEntityNode *side_chain_next;

    void *head;
    void *body;

    PBoxArrayEntity entity;

} PBoxArrayEntityNode;

PBoxArrayEntityNode *PBoxArrayEntityNode_new(pbox_size unit_size, pbox_count count);

void PBoxArrayEntityNode_delete(PBoxArrayEntityNode *node);

////////////////////////////////////////////////////////////////////////////////
// impl: PBoxArrayManager

static PBoxArrayManager *the_global_array_manager = NIL;

PBoxArrayEntity *PBoxArrayManager_create_new_entity(PBoxArrayManager *self, PBoxArrayEntity *params)
{

    return NIL;
}

PBoxArrayManager *PBoxArrayManager_get_instance()
{
    PBoxArrayManager *inst = the_global_array_manager;
    if (inst == NIL)
    {
        PBoxArrayManagerInner *inner = PBoxArrayManagerInner_create_new_instance();
        inst = inner->outer;
        the_global_array_manager = inst;
    }
    return inst;
}

////////////////////////////////////////////////////////////////////////////////
//  impl  : PBoxArrayEntityNode

PBoxArrayEntityNode *PBoxArrayEntityNode_new(pbox_size unit_size, pbox_count count)
{
    if (unit_size < 1)
    {
        unit_size = 1;
    }
    if (count < 1)
    {
        count = 1;
    }

    pbox_size size1 = sizeof(PBoxArrayEntityNode); // head-size
    pbox_size size2 = unit_size * count;           // capacity
    pbox_size size3 = size1 + size2;               // total-size

    void *p = malloc(size3);
    if (p == NIL)
    {
        return NIL;
    }

    PBoxArrayEntityNode *node = p;
    memset(p, 0, size3);

    node->head = p;
    node->body = node->entity.data;
    node->total_size = size3;

    node->entity.capacity = size2;
    node->entity.unit_size = unit_size;
    node->entity.count = 0;
    node->entity.count_max = count;
    node->entity.owner = NIL;
    node->entity.manager = NIL;

    pbox_bytes addr = node->body;
    ESP_LOGD(PBOX_LOG_TAG, "PBoxArrayEntityNode_new: %d bytes, @%d", size2, addr);

    return node;
}

void PBoxArrayEntityNode_delete(PBoxArrayEntityNode *node)
{
    if (node == NIL)
    {
        return;
    }
    void *head1 = node->head;
    void *body1 = node->body;
    void *head2 = node;
    void *body2 = node->entity.data;
    if ((head1 == head2) && (body1 == body2))
    {

        pbox_size capacity = node->entity.capacity;
        pbox_bytes addr = node->body;
        ESP_LOGD(PBOX_LOG_TAG, "PBoxArrayEntityNode_delete: %d bytes, @%d", capacity, addr);

        pbox_size size = sizeof(PBoxArrayEntityNode);
        memset(node, 0, size);
        free(node);
    }
}

////////////////////////////////////////////////////////////////////////////////
// impl: PBoxArrayManagerInner

PBoxArrayManagerInner *PBoxArrayManagerInner_create_new_instance()
{
    PBoxArrayManagerInner *inner = NIL;
    PBoxArrayManager *outer = NIL;
    pbox_size inner_size = sizeof(PBoxArrayManagerInner);
    pbox_size outer_size = sizeof(PBoxArrayManager);

    inner = malloc(inner_size);
    outer = malloc(outer_size);
    if (inner == NIL || outer == NIL)
    {
        ESP_LOGE(PBOX_LOG_TAG, "PBoxArrayManagerInner_create_new_instance: malloc() return NIL");
        return NIL;
    }
    memset(inner, 0, inner_size);
    memset(outer, 0, outer_size);

    inner->outer = outer;
    outer->inner = inner;
    return inner;
}

////////////////////////////////////////////////////////////////////////////////
// impl : PBoxArrayEntity

void PBoxArrayEntity_init_params(PBoxArrayEntity *self) {}

////////////////////////////////////////////////////////////////////////////////
