package com.bitwormhole.passwordbox.app.core.keybase;

import androidx.annotation.NonNull;

public class SecretKeyAlgorithmSelector {

    private String provider;
    private SecretKeyOption option;
    private boolean supported;
    private Throwable error;

    public SecretKeyAlgorithmSelector() {
    }

    @NonNull
    @Override
    public String toString() {
        return this.algorithm();
    }

    public String algorithm() {
        return SecretKeyOption.toString(this.option);
    }

    public String getProvider() {
        return provider;
    }

    public void setProvider(String provider) {
        this.provider = provider;
    }

    public SecretKeyOption getOption() {
        return option;
    }

    public void setOption(SecretKeyOption option) {
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
