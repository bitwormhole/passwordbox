
#ifndef __pbox_app_h__
#define __pbox_app_h__

#include "pbox_errors.h"

typedef struct PBoxApp_t PBoxApp;

typedef struct PBoxAppInner_t PBoxAppInner;

////////////////////////////////////////////////////////////////////////////////

typedef struct PBoxApp_t
{

    PBoxAppInner *inner;

} PBoxApp;

////////////////////////////////////////////////////////////////////////////////

PBoxError PBoxApp_run(PBoxApp *app);

#endif //  __pbox_app_h__
