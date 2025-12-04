package com.bitwormhole.passwordbox.app.core.keybase;

public enum SignatureOption {

    DSA,

    DSAwithSHA1,


    ECDSA,

    ECDSAwithSHA1,

    Ed25519,

    MD5withRSA,


    NONEwithDSA,
    NONEwithECDSA,
    NONEwithRSA,


    SHA1withDSA,
    SHA1withECDSA,
    SHA1withRSA,


    SHA1withRSA_PSS,
    SHA224withDSA,
    SHA224withECDSA,
    SHA224withRSA,
    SHA224withRSA_PSS,
    SHA256withDSA,
    SHA256withECDSA,
    SHA256withRSA,
    SHA256withRSA_PSS,
    SHA384withECDSA,
    SHA384withRSA,
    SHA384withRSA_PSS,
    SHA512withECDSA,
    SHA512withRSA,
    SHA512withRSA_PSS,

    ;

    public static SignatureOption[] listAll() {
        return new SignatureOption[]{

                SignatureOption.SHA256withRSA,
                SignatureOption.SHA256withRSA_PSS,
                SignatureOption.SHA256withECDSA,
                SignatureOption.SHA256withDSA,

                SignatureOption.SHA1withRSA,
                SignatureOption.SHA1withRSA_PSS,
                SignatureOption.SHA1withECDSA,
                SignatureOption.SHA1withDSA,

                SignatureOption.SHA224withRSA,
                SignatureOption.SHA224withRSA_PSS,
                SignatureOption.SHA224withECDSA,
                SignatureOption.SHA224withDSA,

                SignatureOption.SHA512withRSA,
                SignatureOption.SHA512withRSA_PSS,
                SignatureOption.SHA512withECDSA,

                SignatureOption.SHA384withRSA,
                SignatureOption.SHA384withRSA_PSS,
                SignatureOption.SHA384withECDSA,


                SignatureOption.DSA,
                SignatureOption.DSAwithSHA1,

                SignatureOption.ECDSAwithSHA1,
                SignatureOption.ECDSA,

                SignatureOption.Ed25519,
                SignatureOption.MD5withRSA,

                SignatureOption.NONEwithRSA,
                SignatureOption.NONEwithECDSA,
                SignatureOption.NONEwithDSA,
        };
    }

    public static String toString(SignatureOption opt) {
        if (opt == null) {
            opt = SignatureOption.SHA256withRSA;
        }
        return opt.name();
    }
}
