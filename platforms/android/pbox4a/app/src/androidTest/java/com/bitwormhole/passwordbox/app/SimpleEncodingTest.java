package com.bitwormhole.passwordbox.app;

import android.content.Context;

import androidx.test.platform.app.InstrumentationRegistry;
import androidx.test.ext.junit.runners.AndroidJUnit4;

import com.bitwormhole.passwordbox.app.core.encoding.Base64;
import com.bitwormhole.passwordbox.app.core.encoding.Hex;
import com.bitwormhole.passwordbox.app.core.loggers.Loggers;
import com.bitwormhole.passwordbox.app.core.utils.Bytes;
import com.bitwormhole.passwordbox.app.core.utils.Sum;

import org.junit.Assert;
import org.junit.Test;
import org.junit.runner.RunWith;

import org.junit.Assert.*;

import java.nio.charset.StandardCharsets;


@RunWith(AndroidJUnit4.class)
public class SimpleEncodingTest {

    @Test
    public void testBase64() {

        final String tag = this.getClass().getSimpleName();
        final byte[] data0 = tag.getBytes(StandardCharsets.UTF_8);

        Base64 base1 = new Base64(data0);
        String str1 = base1.toString();
        Base64 base2 = Base64.parse(str1);
        String str2 = base2.toString();

        byte[] data1 = base1.getData();
        byte[] data2 = base2.getData();

        Assert.assertArrayEquals(data0, data1);
        Assert.assertArrayEquals(data0, data2);

        Loggers.info(tag, "base64.str1 = %s", str1);
        Loggers.info(tag, "base64.str2 = %s", str2);
    }

    @Test
    public void testHex() {

        final String tag = "" + this;
        final byte[] data0raw = tag.getBytes(StandardCharsets.UTF_8);
        final byte[] data0 = Sum.sha256sum(data0raw).toByteArray();

        Hex h1 = new Hex(data0);
        String str1 = h1.toString();

        Hex h2 = Hex.parse(str1);
        String str2 = Hex.toString(h2);

        Loggers.info(tag, "hex.str1 = %s", str1);
        Loggers.info(tag, "hex.str2 = %s", str2);
        Loggers.info(tag, "data.raw = %s", Bytes.toString(data0));

        byte[] data1 = h1.getData();
        byte[] data2 = h2.getData();

        Assert.assertEquals(str1, str2);
        Assert.assertArrayEquals(data0, data1);
        Assert.assertArrayEquals(data0, data2);
    }

    @Test
    public void testPEM() {
    }

    private Context getAppContext() {
        return InstrumentationRegistry.getInstrumentation().getTargetContext();
    }
}
