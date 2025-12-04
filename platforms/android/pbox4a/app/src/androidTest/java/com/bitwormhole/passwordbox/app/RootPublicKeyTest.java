package com.bitwormhole.passwordbox.app;

import android.content.Context;
import android.util.Log;

import androidx.test.ext.junit.runners.AndroidJUnit4;
import androidx.test.platform.app.InstrumentationRegistry;

import com.bitwormhole.passwordbox.app.core.keybase.FingerPrint;
import com.bitwormhole.passwordbox.app.core.keybase.KeyPairAgent;
import com.bitwormhole.passwordbox.app.core.keybase.KeyPairHolder;
import com.bitwormhole.passwordbox.app.core.keybase.PublicKeys;
import com.bitwormhole.passwordbox.app.core.keybase.SecretKeyAgent;
import com.bitwormhole.passwordbox.app.core.keybase.SecretKeyHolder;
import com.bitwormhole.passwordbox.app.core.utils.Bytes;

import org.junit.Test;
import org.junit.runner.RunWith;

import java.security.KeyPair;

import javax.crypto.SecretKey;


@RunWith(AndroidJUnit4.class)
public class RootPublicKeyTest {

    @Test
    public void testRootKeyPair() {

        String tag = this.getClass().getSimpleName();
        String msg = "";

        KeyPairHolder root = KeyPairAgent.getRoot();
        KeyPairHolder holder;

        if (root.exists()) {
            Log.i(tag, "load a existed key-pair");
            holder = root.load();
        } else {
            Log.i(tag, "generate a new key-pair");
            holder = root.generate();
        }

        KeyPair pair = holder.getPair();
        FingerPrint fp = PublicKeys.getFingerPrint(pair.getPublic());

        msg = "public_key.finger-print = " + fp;
        Log.i(tag, msg);
    }

    private Context getContext() {
        return InstrumentationRegistry.getInstrumentation().getTargetContext();
    }
}
