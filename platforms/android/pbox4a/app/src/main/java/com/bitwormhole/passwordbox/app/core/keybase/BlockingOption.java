package com.bitwormhole.passwordbox.app.core.keybase;

public enum BlockingOption {

    NONE,

    CBC,

    CFB,

    CTR,

    CTS,

    ECB,

    OFB,

    GCM,

    GCM_SIV,
    ;

    public static BlockingOption[] listAll() {
        return new BlockingOption[]{
                BlockingOption.NONE,
                BlockingOption.ECB,
                BlockingOption.CBC,
                BlockingOption.CFB,
                BlockingOption.CTR,
                BlockingOption.CTS,
                BlockingOption.OFB,
                BlockingOption.GCM,
                BlockingOption.GCM_SIV,
        };
    }

    public static String toString(BlockingOption opt) {
        if (opt == null) {
            opt = BlockingOption.NONE;
        }
        return opt.name();
    }
}
