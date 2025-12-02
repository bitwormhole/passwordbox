package com.bitwormhole.passwordbox.app.tasks;

import android.app.Activity;

import java.util.concurrent.Executor;

final class ActivityForegroundExecutor implements Executor {

    private final Activity activity;

    public ActivityForegroundExecutor(Activity a) {
        this.activity = a;
    }

    @Override
    public void execute(Runnable command) {
        this.activity.runOnUiThread(command);
    }
}
