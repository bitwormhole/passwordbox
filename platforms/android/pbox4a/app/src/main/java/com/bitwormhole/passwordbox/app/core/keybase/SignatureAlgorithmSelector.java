package com.bitwormhole.passwordbox.app.core.keybase;

import androidx.annotation.NonNull;

public class SignatureAlgorithmSelector {

    private String provider;
    private SignatureOption option;
    private boolean supported;
    private Throwable error;

    public SignatureAlgorithmSelector() {
    }

    @NonNull
    @Override
    public String toString() {
        return this.algorithm();
    }

    public String algorithm(){
        String str = SignatureOption.toString(this.option);
        if (str == null) {
            str = "";
        }
        return str;
    }

    public SignatureOption getOption() {
        return option;
    }

    public void setOption(SignatureOption option) {
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

    public String getProvider() {
        return provider;
    }

    public void setProvider(String provider) {
        this.provider = provider;
    }
}
