package com.bitwormhole.passwordbox.app.core.keybase;

import java.security.NoSuchAlgorithmException;
import java.security.SecureRandom;

public final class RandomAgent {

    private RandomAgent() {
    }

    public static SecureRandom getDefaultRandom() {
        try {
            String alg = "SHA1PRNG";
            // String alg = "SHA256PRNG";
            return SecureRandom.getInstance(alg);
        } catch (NoSuchAlgorithmException e) {
            throw new RuntimeException(e);
        }
    }

}
