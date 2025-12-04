package com.bitwormhole.passwordbox.app;

import android.util.Log;

import androidx.test.ext.junit.runners.AndroidJUnit4;

import com.bitwormhole.passwordbox.app.core.keybase.AlgorithmsChecker;
import com.bitwormhole.passwordbox.app.core.keybase.CipherAlgorithmSelector;
import com.bitwormhole.passwordbox.app.core.keybase.SignatureAlgorithmSelector;

import org.junit.Test;
import org.junit.runner.RunWith;

import java.util.ArrayList;
import java.util.List;


@RunWith(AndroidJUnit4.class)
public class CipherAlgorithmsTest {

    @Test
    public void testAndroidCipherAlgorithms() {

        String provider0 = null;
        String provider1 = null;


        final String tag = "testAndroidCipherAlgorithms";
        AlgorithmsChecker checker = new AlgorithmsChecker();
        List<CipherAlgorithmSelector> results = checker.checkCipherAlgorithms(provider0);
        StringBuilder msg = new StringBuilder();

        List<String> listOK = new ArrayList<>();
        List<String> listErr = new ArrayList<>();


        for (CipherAlgorithmSelector sel : results) {

            provider1 = sel.getProvider();

            msg.setLength(0);
            msg.append("[check algorithm_type:cipher");
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
