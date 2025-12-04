package com.bitwormhole.passwordbox.app;

import android.util.Log;

import androidx.test.ext.junit.runners.AndroidJUnit4;

import com.bitwormhole.passwordbox.app.core.keybase.AlgorithmsChecker;
import com.bitwormhole.passwordbox.app.core.keybase.SignatureAlgorithmSelector;

import org.junit.Test;
import org.junit.runner.RunWith;

import java.util.List;


@RunWith(AndroidJUnit4.class)
public class SignatureAlgorithmsTest {

    @Test
    public void testAndroidSignatureAlgorithms() {

        String provider = null;
        final String tag = "testAndroidSignatureAlgorithms";
        AlgorithmsChecker checker = new AlgorithmsChecker();
        List<SignatureAlgorithmSelector> results = checker.checkSignatureAlgorithms(provider);
        StringBuilder msg = new StringBuilder();

        for (SignatureAlgorithmSelector sel : results) {

            provider = sel.getProvider();

            msg.setLength(0);
            msg.append("[check algorithm_type:signature");
            msg.append(" provider:").append(provider);
            msg.append(" name:").append(sel);
            msg.append(" supported:").append(sel.isSupported());
            msg.append(']');

            if (sel.isSupported()) {
                Log.i(tag, msg.toString());
            } else {
                msg.append(sel.getError());
                Log.e(tag, msg.toString());
            }
        }
    }
}
