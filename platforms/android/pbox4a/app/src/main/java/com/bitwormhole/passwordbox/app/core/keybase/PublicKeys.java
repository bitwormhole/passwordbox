package com.bitwormhole.passwordbox.app.core.keybase;

import com.bitwormhole.passwordbox.app.core.utils.Sum;

import java.security.PublicKey;

public class PublicKeys {

    public static FingerPrint getFingerPrint(PublicKey key) {
        return getFingerPrintSHA256(key);
    }

    public static FingerPrint getFingerPrintSHA256(PublicKey key) {
        Sum sum = Sum.sha256sum(key.getEncoded());
        return new FingerPrint(sum.toByteArray());
    }

    public static FingerPrint getFingerPrintSHA1(PublicKey key) {
        Sum sum = Sum.sha1sum(key.getEncoded());
        return new FingerPrint(sum.toByteArray());
    }

}
