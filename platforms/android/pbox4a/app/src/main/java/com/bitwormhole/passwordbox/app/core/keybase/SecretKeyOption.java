package com.bitwormhole.passwordbox.app.core.keybase;

public enum SecretKeyOption {

    AES, AES_128, AES_256,


    ARC4,

    BLOWFISH,

    ChaCha20,

    DES,

    DESede,

    HmacMD5,


    HmacSHA1, HmacSHA224, HmacSHA256, HmacSHA384, HmacSHA512,

    ;

    public static SecretKeyOption[] listAll() {
        return new SecretKeyOption[]{
                SecretKeyOption.AES,
                SecretKeyOption.AES_128,
                SecretKeyOption.AES_256,
                SecretKeyOption.ARC4,
                SecretKeyOption.BLOWFISH,
                SecretKeyOption.ChaCha20,
                SecretKeyOption.DES,
                SecretKeyOption.DESede,
                SecretKeyOption.HmacMD5,
                SecretKeyOption.HmacSHA1,
                SecretKeyOption.HmacSHA224,
                SecretKeyOption.HmacSHA256,
                SecretKeyOption.HmacSHA384,
                SecretKeyOption.HmacSHA512,
        };
    }

    public static String toString(SecretKeyOption opt) {
        if (opt == null) {
            opt = SecretKeyOption.AES;
        }
        return opt.name();
    }
}
