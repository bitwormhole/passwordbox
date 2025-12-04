package com.bitwormhole.passwordbox.app.ui.tools;

import android.content.Context;
import android.util.Log;
import android.widget.Toast;

import java.util.ArrayList;
import java.util.List;

public class ErrorDisplay {

    public final static int FLAG_SHOW_LOG = 0x01;
    public final static int FLAG_SHOW_TOAST = 0x02;
    public final static int FLAG_SHOW_DIALOG = 0x04;

    public final static int FLAG_SHOW_DEFAULT = FLAG_SHOW_LOG | FLAG_SHOW_TOAST | FLAG_SHOW_DIALOG;


    public static void show(Context ctx, Throwable err, int flag) {

        if (ctx == null || err == null) {
            return;
        }
        if (flag == 0) {
            flag = FLAG_SHOW_DEFAULT;
        }

        final myErrorMessageContext emc = new myErrorMessageContext();
        final List<myErrorMessageDisplayer> list = new ArrayList<>();

        if ((flag & FLAG_SHOW_LOG) > 0) {
            list.add(new myInnerLogDisplayer());
        }
        if ((flag & FLAG_SHOW_TOAST) > 0) {
            list.add(new myInnerToastDisplayer());
        }
        if ((flag & FLAG_SHOW_DIALOG) > 0) {
            list.add(new myInnerDialogDisplayer());
        }

        emc.context = ctx;
        emc.error = err;
        emc.flag = flag;

        for (myErrorMessageDisplayer displayer : list) {
            displayer.display(emc);
        }
    }


    private static class myInnerToastDisplayer implements myErrorMessageDisplayer {
        @Override
        public void display(myErrorMessageContext ctx) {
            String msg = "Error: " + ctx.error.getMessage();
            Toast toast1 = Toast.makeText(ctx.context, msg, Toast.LENGTH_LONG);
            toast1.show();
        }
    }

    private static class myInnerDialogDisplayer implements myErrorMessageDisplayer {
        @Override
        public void display(myErrorMessageContext ctx) {

        }
    }

    private static class myInnerLogDisplayer implements myErrorMessageDisplayer {
        @Override
        public void display(myErrorMessageContext ctx) {
            String tag = this.getClass().getSimpleName();
            String msg = "handle error: ";
            Log.e(tag, msg, ctx.error);
        }
    }

    private static class myErrorMessageContext {
        Context context;
        Throwable error;
        int flag;
    }

    private interface myErrorMessageDisplayer {
        void display(myErrorMessageContext ctx);
    }
}
