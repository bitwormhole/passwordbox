
#include "pbox_app.h"

#include "pbox_module.h"
#include "pbox_logger.h"
#include "pbox_app_context.h"

#include "module_auto_test.h"
#include "module_ble.h"
#include "module_usb_hid.h"
#include "module_sd_card.h"
#include "module_wifi.h"

#include <memory.h>

////////////////////////////////////////////////////////////////////////////////

PBoxError PBoxApp_on_init(PBoxApp *app);
PBoxError PBoxApp_on_create(PBoxApp *app);
PBoxError PBoxApp_on_start(PBoxApp *app);
PBoxError PBoxApp_on_resume(PBoxApp *app);
PBoxError PBoxApp_on_run(PBoxApp *app);
PBoxError PBoxApp_on_pause(PBoxApp *app);
PBoxError PBoxApp_on_stop(PBoxApp *app);
PBoxError PBoxApp_on_destroy(PBoxApp *app);

PBoxError PBoxApp_invoke_lifecycle_func(PBoxApp *app, PBoxModuleLifePhase sel, pbox_bool reverse);

PBoxError PBoxApp_prepare_modules(PBoxApp *app);
PBoxError PBoxApp_prepare_app_context(PBoxApp *app);
PBoxError PBoxApp_prepare_inner(PBoxApp *app);

PBoxError PBoxApp_release_inner(PBoxApp *app);

////////////////////////////////////////////////////////////////////////////////

typedef struct PBoxAppInner_t
{

    PBoxApp *outer;
    PBoxAppContext *context;

    PBoxModuleManager module_manager;

    PBoxAutoTestModule mod_auto_test;
    PBoxBleModule mod_ble;
    PBoxSDCardModule mod_sd_card;
    PBoxUsbHidModule mod_usb_hid;
    PBoxWifiModule mod_wifi;

} PBoxAppInner;

void PBoxAppInner_init(PBoxAppInner *self);

void PBoxAppInner_on_create(PBoxAppInner *self);

void PBoxAppInner_on_destroy(PBoxAppInner *self);

////////////////////////////////////////////////////////////////////////////////
// app-inner

void PBoxAppInner_init(PBoxAppInner *self)
{
    if (self)
    {
        memset(self, 0, sizeof(self[0]));

        PBoxModuleManager *mm = &self->module_manager;
        PBoxModuleManager_init(mm);
        PBoxModuleManager_create(mm);
    }
}

void PBoxAppInner_on_create(PBoxAppInner *self)
{
}

void PBoxAppInner_on_destroy(PBoxAppInner *self)
{
}

////////////////////////////////////////////////////////////////////////////////
// app-life

PBoxError PBoxApp_on_init(PBoxApp *app)
{
    if (app == NIL)
    {
        return pbox_make_error(500, "PBoxApp_on_init", "param:app is nil");
    }

    memset(app, 0, sizeof(app[0]));

    return NIL;
}

PBoxError PBoxApp_on_create(PBoxApp *app)
{
    PBoxError err = NIL;
    err = PBoxApp_prepare_inner(app);
    if (err)
    {
        return err;
    }

    err = PBoxApp_prepare_modules(app);
    if (err)
    {
        return err;
    }

    err = PBoxApp_prepare_app_context(app);
    if (err)
    {
        return err;
    }

    // invoke modules

    PBoxModuleLifePhase sel = PBoxModuleLifePhaseOnCreate;
    pbox_bool reverse = NO;
    return PBoxApp_invoke_lifecycle_func(app, sel, reverse);
}

PBoxError PBoxApp_on_start(PBoxApp *app)
{

    PBoxModuleLifePhase sel = PBoxModuleLifePhaseOnStart;
    pbox_bool reverse = NO;
    return PBoxApp_invoke_lifecycle_func(app, sel, reverse);
}

PBoxError PBoxApp_on_resume(PBoxApp *app)
{
    PBoxModuleLifePhase sel = PBoxModuleLifePhaseOnResume;
    pbox_bool reverse = NO;
    return PBoxApp_invoke_lifecycle_func(app, sel, reverse);
}

PBoxError PBoxApp_on_run(PBoxApp *app)
{
    PBoxModuleLifePhase sel = PBoxModuleLifePhaseOnRun;
    pbox_bool reverse = NO;
    return PBoxApp_invoke_lifecycle_func(app, sel, reverse);
}

PBoxError PBoxApp_on_pause(PBoxApp *app)
{
    PBoxModuleLifePhase sel = PBoxModuleLifePhaseOnPause;
    pbox_bool reverse = YES;
    return PBoxApp_invoke_lifecycle_func(app, sel, reverse);
}

PBoxError PBoxApp_on_stop(PBoxApp *app)
{
    PBoxModuleLifePhase sel = PBoxModuleLifePhaseOnStop;
    pbox_bool reverse = YES;
    return PBoxApp_invoke_lifecycle_func(app, sel, reverse);
}

PBoxError PBoxApp_on_destroy(PBoxApp *app)
{
    // invoke modules
    PBoxModuleLifePhase sel = PBoxModuleLifePhaseOnDestroy;
    pbox_bool reverse = YES;
    PBoxError err = PBoxApp_invoke_lifecycle_func(app, sel, reverse);
    if (err)
    {
    }

    return NIL;
}

PBoxError PBoxApp_invoke_lifecycle_func(PBoxApp *app, PBoxModuleLifePhase sel, pbox_bool reverse)
{
    PBoxModuleManager *mm = &app->inner->module_manager;
    PBoxAppContext *ac = app->inner->context;

    PBoxModuleIterator iter1;
    PBoxModuleIterator *iter2;

    iter2 = PBoxModuleManager_iterate(mm, &iter1, reverse);

    while (PBoxModuleIterator_has_more(iter2))
    {
        PBoxModule *mod = PBoxModuleIterator_next(iter2);
        PBoxError err = PBoxModule_invoke_lifecycle_func(mod, ac, sel);
        if (err)
        {
            return err;
        }
    }

    return NIL;
}

PBoxError PBoxApp_prepare_modules(PBoxApp *app)
{
    PBoxAppInner *inner = app->inner;
    PBoxModuleManager *mm = &inner->module_manager;
    PBoxErrorHolder eh;

    PBoxErrorHolder_init(&eh, YES);
    PBoxModuleManager_init(mm);
    PBoxModuleManager_create(mm);

    PBoxAutoTestModule *mod_auto_test = &inner->mod_auto_test;
    PBoxBleModule *mod_ble = &inner->mod_ble;
    PBoxSDCardModule *mod_sd_card = &inner->mod_sd_card;
    PBoxUsbHidModule *mod_usb_hid = &inner->mod_usb_hid;
    PBoxWifiModule *mod_wifi = &inner->mod_wifi;

    PBoxModule *m1 = PBoxAutoTestModule_init(mod_auto_test);
    PBoxModule *m2 = PBoxBleModule_init(mod_ble);
    PBoxModule *m3 = PBoxSDCardModule_init(mod_sd_card);
    PBoxModule *m4 = PBoxUsbHidModule_init(mod_usb_hid);
    PBoxModule *m5 = PBoxWifiModule_init(mod_wifi);

    PBoxModuleManager_add_with_eh(mm, m1, &eh);
    PBoxModuleManager_add_with_eh(mm, m2, &eh);
    PBoxModuleManager_add_with_eh(mm, m3, &eh);
    PBoxModuleManager_add_with_eh(mm, m4, &eh);
    PBoxModuleManager_add_with_eh(mm, m5, &eh);

    // PBoxModuleManager_add_with_eh(mm, m5, &eh);
    // PBoxModuleManager_add_with_eh(mm, m5, &eh);
    // PBoxModuleManager_add_with_eh(mm, m5, &eh);
    // PBoxModuleManager_add_with_eh(mm, m5, &eh);

    return PBoxErrorHolder_get_error(&eh);
}

PBoxError PBoxApp_prepare_app_context(PBoxApp *app)
{

    PBoxAppContext *ac = NIL;
    PBoxAppInner *inner = app->inner;

    ac = malloc(sizeof(PBoxAppContext));
    PBoxAppContext_init(ac);

    inner->context = ac;

    ac->app = app;
    ac->auto_test = &inner->mod_auto_test;
    ac->ble = &inner->mod_ble;
    ac->sd_card = &inner->mod_sd_card;
    ac->usb_hid = &inner->mod_usb_hid;
    ac->wifi = &inner->mod_wifi;

    return NIL;
}

PBoxError PBoxApp_prepare_inner(PBoxApp *app)
{
    PBoxAppInner *inner = NIL;
    inner = malloc(sizeof(inner[0]));
    PBoxAppInner_init(inner);
    inner->outer = app;
    app->inner = inner;
    PBoxAppInner_on_create(inner);
    return NIL;
}

PBoxError PBoxApp_release_inner(PBoxApp *app)
{
    PBoxAppInner *inner = app->inner;
    app->inner = NIL;
    if (inner)
    {
        PBoxAppInner_on_destroy(inner);
        free(inner);
    }
    return NIL;
}

////////////////////////////////////////////////////////////////////////////////
// app-main

PBoxError PBoxApp_run(PBoxApp *app)
{
    if (app == NIL)
    {
        return pbox_make_error(500, "PBoxApp_run", "param:app is nil");
    }

    PBoxError err = PBoxApp_on_init(app);
    if (err)
    {
        return err;
    }

    err = PBoxApp_on_create(app);
    if (err)
    {
        return err;
    }

    err = PBoxApp_on_start(app);
    if (err)
    {
        return err;
    }

    err = PBoxApp_on_resume(app);
    if (err)
    {
        return err;
    }

    err = PBoxApp_on_run(app);
    if (err)
    {
        return err;
    }

    err = PBoxApp_on_pause(app);
    if (err)
    {
        return err;
    }

    err = PBoxApp_on_stop(app);
    if (err)
    {
        return err;
    }

    err = PBoxApp_on_destroy(app);
    if (err)
    {
        return err;
    }

    return NIL;
}

////////////////////////////////////////////////////////////////////////////////
// EOF
