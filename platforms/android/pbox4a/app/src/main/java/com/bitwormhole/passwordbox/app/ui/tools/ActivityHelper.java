package com.bitwormhole.passwordbox.app.ui.tools;

import android.app.Activity;
import android.content.Intent;

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

}
