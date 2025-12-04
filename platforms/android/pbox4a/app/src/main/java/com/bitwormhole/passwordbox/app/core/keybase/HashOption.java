package com.bitwormhole.passwordbox.app.core.keybase;

public enum HashOption {

    MD5,

    SHA1, SHA256, SHA512,

    ;

    public static HashOption[] listAll() {
        return new HashOption[]{
                HashOption.SHA256,
                HashOption.SHA1,
                HashOption.SHA512,
                HashOption.MD5,
        };
    }

    public static String toString(HashOption opt) {
        if (opt == null) {
            opt = HashOption.SHA256;
        }
        return opt.name();
    }
}
