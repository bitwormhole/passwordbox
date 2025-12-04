package com.bitwormhole.passwordbox.app;

import android.util.Log;

import androidx.test.ext.junit.runners.AndroidJUnit4;

import com.bitwormhole.passwordbox.app.core.keybase.AlgorithmsChecker;
import com.bitwormhole.passwordbox.app.core.keybase.CipherAlgorithmSelector;
import com.bitwormhole.passwordbox.app.core.keybase.KeyPairAlgorithmSelector;

import org.junit.Test;
import org.junit.runner.RunWith;

import java.util.ArrayList;
import java.util.List;


@RunWith(AndroidJUnit4.class)
public class PublicKeyAlgorithmsTest {


    @Test
    public void testPublicKeyWithNull() {
        this.playForPublicKey(null);
    }

    @Test
    public void testPublicKeyWithBC() {
        this.playForPublicKey("BC");
    }

    @Test
    public void testPublicKeyWithAoSSL() {
        this.playForPublicKey("AndroidOpenSSL");
    }

    @Test
    public void testPublicKeyWithAKS() {
        this.playForPublicKey("AndroidKeyStore");
    }


    /// ////////////////////////////////////////////////////////////////////////////////////////////


    public void playForPublicKey(String provider0) {

        String provider1 = null;
        final String tag = "playTestForPublicKey";
        AlgorithmsChecker checker = new AlgorithmsChecker();
        List<KeyPairAlgorithmSelector> results = checker.checkKeyPairAlgorithms(provider0);

        StringBuilder msg = new StringBuilder();
        List<String> listOK = new ArrayList<>();
        List<String> listErr = new ArrayList<>();

        for (KeyPairAlgorithmSelector sel : results) {

            provider1 = sel.getProvider();
            msg.setLength(0);

            msg.append("[check algorithm_type:key-pair");
            msg.append(" raw_provider:").append(provider0);
            msg.append(" provider:").append(provider1);
            msg.append(" name:").append(sel);
            msg.append(" supported:").append(sel.isSupported());
            msg.append(']');

            if (sel.isSupported()) {
                listOK.add(msg.toString());
            } else {
                msg.append(sel.getError());
                listErr.add(msg.toString());
            }
        }

        for (String line : listOK) {
            Log.i(tag, line);
        }
        for (String line : listErr) {
            Log.e(tag, line);
        }
    }
}
