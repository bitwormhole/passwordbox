package com.bitwormhole.passwordbox.app.core.keybase;

public final class KeyPairAgent {

    public static KeyPairHolder getRoot() {
        return new RootKeyPairHolder("root_pub");
    }

}
