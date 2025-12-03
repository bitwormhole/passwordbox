package com.bitwormhole.passwordbox.app.ui.tools;

import android.content.Context;

public class ErrorDisplay {

    public final static int FLAG_SHOW_LOG = 0x01;
    public final static int FLAG_SHOW_TOAST = 0x02;
    public final static int FLAG_SHOW_DIALOG = 0x04;

    public final static int FLAG_SHOW_DEFAULT = FLAG_SHOW_LOG | FLAG_SHOW_TOAST | FLAG_SHOW_DIALOG;


    public static void show(Context ctx, Throwable err, int flag) {

        // todo : no impl
    }

}
