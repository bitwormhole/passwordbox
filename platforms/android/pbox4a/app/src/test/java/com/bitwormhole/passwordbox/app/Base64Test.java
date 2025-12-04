package com.bitwormhole.passwordbox.app;

import com.bitwormhole.passwordbox.app.core.encoding.Base64;
import com.bitwormhole.passwordbox.app.core.loggers.Loggers;

import org.junit.Assert;
import org.junit.Test;

import java.nio.charset.StandardCharsets;

public class Base64Test {


    @Test
    public void testBase64codec() {

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

}
