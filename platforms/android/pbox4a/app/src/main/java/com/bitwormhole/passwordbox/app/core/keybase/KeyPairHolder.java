package com.bitwormhole.passwordbox.app.core.keybase;

import java.security.KeyPair;

public interface KeyPairHolder {

    KeyPairHolder load();

    KeyPairHolder generate();

    boolean exists();

    KeyPair getPair();
}
