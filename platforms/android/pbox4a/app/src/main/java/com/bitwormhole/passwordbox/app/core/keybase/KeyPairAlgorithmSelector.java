package com.bitwormhole.passwordbox.app.core.keybase;

import androidx.annotation.NonNull;

public class KeyPairAlgorithmSelector {

    private String provider;
    private KeyPairOption option;

    private boolean supported;
    private Throwable error;

    public KeyPairAlgorithmSelector() {
    }

    public String algorithm() {
        return KeyPairOption.toString(this.option);
    }

    @NonNull
    @Override
    public String toString() {
        return this.algorithm();
    }

    public String getProvider() {
        return provider;
    }

    public void setProvider(String provider) {
        this.provider = provider;
    }

    public KeyPairOption getOption() {
        return option;
    }

    public void setOption(KeyPairOption option) {
        this.option = option;
    }

    public boolean isSupported() {
        return supported;
    }

    public void setSupported(boolean supported) {
        this.supported = supported;
    }

    public Throwable getError() {
        return error;
    }

    public void setError(Throwable error) {
        this.error = error;
    }
}
