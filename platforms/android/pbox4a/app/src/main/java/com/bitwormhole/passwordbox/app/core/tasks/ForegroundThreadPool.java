package com.bitwormhole.passwordbox.app.core.tasks;

import java.util.concurrent.Executor;
import java.util.concurrent.Executors;

final class ForegroundThreadPool {

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
