package com.bitwormhole.passwordbox.app;

import android.content.Context;
import android.util.Log;

import androidx.test.ext.junit.runners.AndroidJUnit4;
import androidx.test.platform.app.InstrumentationRegistry;

import com.bitwormhole.passwordbox.app.core.store.HomeDir;

import org.junit.Assert;
import org.junit.Test;
import org.junit.runner.RunWith;

import java.io.File;
import java.io.IOException;


@RunWith(AndroidJUnit4.class)
public class HomeDirTest {

    @Test
    public void testHomeDir() {

        Context ctx = getContext();
        HomeDir home = HomeDir.getInstance(ctx);

        logDirPath("__home", home.getHomeDir());
        logDirPath("_banks", home.getBanksDir());
        logDirPath("config", home.getConfigDir());
        logDirPath("__data", home.getDataDir());
        logDirPath("___tmp", home.getTmpDir());

        try {
            this.createTmpFile(home);
        } catch (IOException e) {
            throw new RuntimeException(e);
        }

    }

    private void createTmpFile(HomeDir home) throws IOException {
        File tmp_dir = home.getTmpDir();
        if (!tmp_dir.exists()) {
            tmp_dir.mkdirs();
        }
        File tmp_file = File.createTempFile("unit", ".tmp", tmp_dir);
        tmp_file.createNewFile();
    }


    private void logDirPath(String name, File path) {
        String msg = name + ".path = " + path;
        Log.i("logDirPath", msg);
    }

    private Context getContext() {
        return InstrumentationRegistry.getInstrumentation().getTargetContext();
    }

}
