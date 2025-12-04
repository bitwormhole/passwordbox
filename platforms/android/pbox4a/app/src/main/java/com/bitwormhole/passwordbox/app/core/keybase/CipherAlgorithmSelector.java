package com.bitwormhole.passwordbox.app.core.keybase;

import androidx.annotation.NonNull;

public class CipherAlgorithmSelector {

    // 注意: pk & sk 二选一

    private KeyPairOption pk;
    private SecretKeyOption sk;

    private PaddingOption padding;
    private BlockingOption blocking;

    private String provider;
    private boolean supported;
    private Throwable error;

    public CipherAlgorithmSelector() {
    }

    @NonNull
    @Override
    public String toString() {
        return this.algorithm();
    }

    public String algorithm() {
        StringBuilder builder = new StringBuilder();
        if (this.sk != null) {
            builder.append(SecretKeyOption.toString(sk));
        } else if (this.pk != null) {
            builder.append(KeyPairOption.toString(pk));
        } else {
            builder.append("NA");
        }
        builder.append('/').append(BlockingOption.toString(blocking));
        builder.append('/').append(PaddingOption.toString(padding));
        return builder.toString();
    }

    public Throwable getError() {
        return error;
    }

    public void setError(Throwable error) {
        this.error = error;
    }

    public KeyPairOption getPk() {
        return pk;
    }

    public void setPk(KeyPairOption pk) {
        this.pk = pk;
    }

    public SecretKeyOption getSk() {
        return sk;
    }

    public void setSk(SecretKeyOption sk) {
        this.sk = sk;
    }

    public PaddingOption getPadding() {
        return padding;
    }

    public void setPadding(PaddingOption padding) {
        this.padding = padding;
    }

    public BlockingOption getBlocking() {
        return blocking;
    }

    public void setBlocking(BlockingOption blocking) {
        this.blocking = blocking;
    }


    public String getProvider() {
        return provider;
    }

    public void setProvider(String provider) {
        this.provider = provider;
    }

    public boolean isSupported() {
        return supported;
    }

    public void setSupported(boolean supported) {
        this.supported = supported;
    }
}
