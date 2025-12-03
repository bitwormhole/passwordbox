package com.bitwormhole.passwordbox.app;

import android.util.Log;

import com.bitwormhole.passwordbox.app.core.encoding.Hex;
import com.bitwormhole.passwordbox.app.core.loggers.Loggers;

import org.junit.Assert;
import org.junit.Test;

import org.junit.Assert.*;


public class HexTest {

    @Test
    public void testHexCodec() {

        byte[] data = {1, 2, 3, 4, 5, 6, 7, 8, 9, 0};

        Hex h1 = new Hex(data);
        String str1 = h1.toString();

        Hex h2 = Hex.parse(str1);
        String str2 = Hex.toString(h2);

        final String tag = "" + this;
        Loggers.info(tag, "hex.str1 = %s", str1);
        Loggers.info(tag, "hex.str2 = %s", str2);

        Assert.assertEquals(str1, str2);
    }

}
