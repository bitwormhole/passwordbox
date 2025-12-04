package com.bitwormhole.passwordbox.app.core.keybase;

public enum PaddingOption {

    NoPadding,

    PKCS1Padding,

    PKCS5Padding,

    OAEPPadding,

    OAEPwithSHA1andMGF1Padding,

    OAEPwithSHA256andMGF1Padding,

    ISO10126Padding,
    ;

    public static PaddingOption[] listAll() {
        return new PaddingOption[]{
                PaddingOption.NoPadding,
                PaddingOption.PKCS1Padding,
                PaddingOption.PKCS5Padding,
                PaddingOption.OAEPPadding,
                PaddingOption.OAEPwithSHA1andMGF1Padding,
                PaddingOption.OAEPwithSHA256andMGF1Padding,
                PaddingOption.ISO10126Padding,
        };
    }

    public static String toString(PaddingOption opt) {
        if (opt == null) {
            opt = PaddingOption.NoPadding;
        }
        return opt.name();
    }
}
