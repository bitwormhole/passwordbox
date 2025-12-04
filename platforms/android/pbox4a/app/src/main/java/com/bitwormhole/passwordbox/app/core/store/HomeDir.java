package com.bitwormhole.passwordbox.app.core.store;

import android.content.Context;

import java.io.File;

public final class HomeDir {

    private final File home;

    private HomeDir(Context c) {
        File dir = c.getDataDir();
        this.home = new File(dir, ".pbox");
    }

    public static HomeDir getInstance(Context ctx) {
        return new HomeDir(ctx);
    }

    public File getHomeDir() {
        return this.home;
    }

    public File getDataDir() {
        return new File(this.home, "data");
    }


    public File getConfigDir() {
        return new File(this.home, "etc");
    }

    public File getTmpDir() {
        return new File(this.home, "tmp");
    }

    public File getBanksDir() {
        return new File(this.home, "banks");
    }
}
