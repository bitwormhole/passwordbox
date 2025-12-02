package com.bitwormhole.passwordbox.app.tasks;

import java.util.concurrent.Executor;
import java.util.concurrent.Executors;

final class BackgroundThreadPool {

    private static Executor pool;

    public static Executor getExecutor() {
        Executor p = pool;
        if (p == null) {
            p = Executors.newSingleThreadExecutor();
            pool = p;
        }
        return p;
    }
}
