package com.bitwormhole.passwordbox.app.core.keybase;

import javax.crypto.Cipher;
import javax.crypto.SecretKey;

public interface SecretKeyHolder {

    boolean exists();

    SecretKeyHolder generate();

    SecretKeyHolder load();

    void delete();

    SecretKey getKey();

    Cipher getCipher(SecretKeyCiphering ciphering);

}
