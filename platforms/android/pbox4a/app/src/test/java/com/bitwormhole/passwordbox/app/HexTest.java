package com.bitwormhole.passwordbox.app;

import android.util.Log;

import com.bitwormhole.passwordbox.app.core.encoding.Hex;
import com.bitwormhole.passwordbox.app.core.loggers.Loggers;
import com.bitwormhole.passwordbox.app.core.utils.Bytes;
import com.bitwormhole.passwordbox.app.core.utils.Sum;

import org.junit.Assert;
import org.junit.Test;

import org.junit.Assert.*;

import java.nio.charset.StandardCharsets;


public class HexTest {

    @Test
    public void testHexCodec() {

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

}
