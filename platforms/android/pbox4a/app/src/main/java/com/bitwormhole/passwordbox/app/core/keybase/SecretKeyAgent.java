package com.bitwormhole.passwordbox.app.core.keybase;

import android.content.Context;

public final class SecretKeyAgent {

    public static SecretKeyHolder getSecretKey0() {
        return new RootSecretKeyHolder0("root_sk_0");
    }

    public static SecretKeyHolder getSecretKey1(Context ctx) {
        return new RootSecretKeyHolder1(ctx, "etc/root_sk_1");
    }

}
