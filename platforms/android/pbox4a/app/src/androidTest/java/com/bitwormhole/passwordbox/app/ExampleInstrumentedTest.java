package com.bitwormhole.passwordbox.app;

import android.content.Context;

import androidx.test.platform.app.InstrumentationRegistry;
import androidx.test.ext.junit.runners.AndroidJUnit4;

import org.junit.Assert;
import org.junit.Test;
import org.junit.runner.RunWith;

import org.junit.Assert.*;


@RunWith(AndroidJUnit4.class)
public class ExampleInstrumentedTest {

    private Context getContext() {
        return InstrumentationRegistry.getInstrumentation().getTargetContext();
    }


    @Test
    public void useAppContext() {
        Context ctx = this.getContext();
        Assert.assertEquals("com.bitwormhole.passwordbox.app", ctx.getPackageName());
    }

}
