package com.bitwormhole.passwordbox.app;

import android.content.Context;
import android.util.Log;

import androidx.test.ext.junit.runners.AndroidJUnit4;
import androidx.test.platform.app.InstrumentationRegistry;

import com.bitwormhole.passwordbox.app.core.keybase.FingerPrint;
import com.bitwormhole.passwordbox.app.core.keybase.KeyPairAgent;
import com.bitwormhole.passwordbox.app.core.keybase.KeyPairHolder;
import com.bitwormhole.passwordbox.app.core.keybase.PublicKeys;
import com.bitwormhole.passwordbox.app.core.keybase.RandomAgent;
import com.bitwormhole.passwordbox.app.core.keybase.SecretKeyAgent;
import com.bitwormhole.passwordbox.app.core.keybase.SecretKeyCipher;
import com.bitwormhole.passwordbox.app.core.keybase.SecretKeyCiphering;
import com.bitwormhole.passwordbox.app.core.keybase.SecretKeyHolder;
import com.bitwormhole.passwordbox.app.core.utils.Bytes;

import org.junit.Assert;
import org.junit.Test;
import org.junit.runner.RunWith;

import java.security.KeyPair;
import java.security.SecureRandom;

import javax.crypto.SecretKey;


@RunWith(AndroidJUnit4.class)
public class RootSecretKeyTest {

    private static final String tag0 = "testRootSecretKey0";
    private static final String tag1 = "testRootSecretKey1";


    @Test
    public void testRootSecretKey0() {

        // gen | load key
        SecretKeyHolder holder = this.getSecretKey0();


        // print-key-info

        SecretKey key = holder.getKey();
        byte[] raw = key.getEncoded();
        String alg = key.getAlgorithm();
        String fmt = key.getFormat();
        String key_enc = Bytes.toString(raw);

        Log.i(tag0, "algorithm = " + alg);
        Log.i(tag0, "format    = " + fmt);
        Log.i(tag0, "encoded   = " + key_enc);


        // encrypt
        byte[] plain0 = this.makeRandomData(1024 * 64);

        SecretKeyCiphering ciphering = new SecretKeyCiphering();
        ciphering.setKeyHolder(holder);
        ciphering.setPlain(plain0);

        SecretKeyCipher.encrypt(ciphering);

        byte[] encrypted = ciphering.getEncrypted();


        // re-load key
        holder = this.getSecretKey0();
        holder = holder.load();

        // decrypt
        SecretKeyCiphering ciphering0 = ciphering;
        SecretKeyCiphering ciphering1 = new SecretKeyCiphering();
        ciphering = ciphering1;

        ciphering1.setKeyHolder(holder);
        ciphering1.setPadding(ciphering0.getPadding());
        ciphering1.setBlocking(ciphering0.getBlocking());
        ciphering1.setEncrypted(encrypted);

        SecretKeyCipher.decrypt(ciphering);
        byte[] plain1 = ciphering.getPlain();


        // check
        Assert.assertArrayEquals(plain0, plain1);
    }


    private SecretKeyHolder getSecretKey0() {
        SecretKeyHolder holder, sk0;
        sk0 = SecretKeyAgent.getSecretKey0();
        // load | gen
        if (sk0.exists()) {
            Log.i(tag0, "load a older key-pair");
            holder = sk0.load();
        } else {
            Log.i(tag0, "generate a new key-pair");
            holder = sk0.generate();
        }
        return holder;
    }

    private byte[] makeRandomData(int size) {
        byte[] data = new byte[size];
        SecureRandom sr = RandomAgent.getDefaultRandom();
        sr.nextBytes(data);
        return data;
    }

    @Test
    public void testRootSecretKey1() {
        // todo ...
    }


    private Context getContext() {
        return InstrumentationRegistry.getInstrumentation().getTargetContext();
    }
}
