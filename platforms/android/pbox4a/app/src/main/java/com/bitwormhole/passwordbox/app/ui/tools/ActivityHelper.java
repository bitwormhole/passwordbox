package com.bitwormhole.passwordbox.app.ui.tools;

import android.app.Activity;
import android.content.Context;
import android.content.Intent;
import android.view.View;

public final class ActivityHelper {

    private final Activity activity;

    public ActivityHelper(Activity a) {
        this.activity = a;
    }

    public void setupButtonToOpenActivity(int id, Class<?> activity_class) {
        this.activity.findViewById(id).setOnClickListener((v) -> {
            Activity ctx = this.activity;
            Intent i = new Intent(ctx, activity_class);
            ctx.startActivity(i);
        });
    }


    public void setupButtonToOpenActivity(int id, View.OnClickListener li) {
        myInnerOnClickListenerWrapper li2 = new myInnerOnClickListenerWrapper(li);
        this.activity.findViewById(id).setOnClickListener(li2);
    }


    private class myInnerOnClickListenerWrapper implements View.OnClickListener {

        final View.OnClickListener targetListener;

        public myInnerOnClickListenerWrapper(View.OnClickListener li) {
            this.targetListener = li;
        }

        @Override
        public void onClick(View v) {
            try {
                this.targetListener.onClick(v);
            } catch (Exception e) {
                final Context ctx = ActivityHelper.this.activity;
                ErrorDisplay.show(ctx, e, 0);
            }
        }
    }

}
