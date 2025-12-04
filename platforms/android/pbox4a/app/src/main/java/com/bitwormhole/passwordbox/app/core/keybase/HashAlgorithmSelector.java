package com.bitwormhole.passwordbox.app.core.keybase;

import androidx.annotation.NonNull;

public final class HashAlgorithmSelector {

    private HashOption option;
    private String provider;
    private Throwable error;
    private boolean supported;

    public HashAlgorithmSelector() {
    }

    @NonNull
    @Override
    public String toString() {
        return this.algorithm();
    }

    public String algorithm() {
        return HashOption.toString(this.option);
    }

    public HashOption getOption() {
        return option;
    }

    public void setOption(HashOption option) {
        this.option = option;
    }

    public String getProvider() {
        return provider;
    }

    public void setProvider(String provider) {
        this.provider = provider;
    }

    public Throwable getError() {
        return error;
    }

    public void setError(Throwable error) {
        this.error = error;
    }

    public boolean isSupported() {
        return supported;
    }

    public void setSupported(boolean supported) {
        this.supported = supported;
    }
}
