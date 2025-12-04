package com.bitwormhole.passwordbox.app.core.keybase;

public enum KeyPairOption {

    DH,

    DSA,

    EC,

    RSA,

    XDH,

    ;

    public static KeyPairOption[] listAll() {
        return new KeyPairOption[]{
                KeyPairOption.RSA,
                KeyPairOption.DSA,
                KeyPairOption.DH,
                KeyPairOption.EC,
                KeyPairOption.XDH,
        };
    }

    public static String toString(KeyPairOption opt) {
        if (opt == null) {
            opt = KeyPairOption.RSA;
        }
        return opt.name();
    }
}
