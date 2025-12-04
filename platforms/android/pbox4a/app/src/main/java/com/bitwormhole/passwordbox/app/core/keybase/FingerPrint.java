package com.bitwormhole.passwordbox.app.core.keybase;

import androidx.annotation.NonNull;

import com.bitwormhole.passwordbox.app.core.encoding.Hex;

import java.security.PublicKey;

public final class FingerPrint {

    private final Hex hex;

    public FingerPrint(byte[] b) {
        this.hex = new Hex(b);
    }

    @NonNull
    @Override
    public String toString() {
        return this.hex.toString();
    }


}
