package com.bitwormhole.passwordbox.app.core.encoding;

public final class Base64 {

    private final byte[] data;

    public Base64() {
        this.data = new byte[0];
    }

    public Base64(byte[] b) {
        if (b == null) {
            b = new byte[0];
        }
        this.data = b;
    }

    public byte[] getData() {
        return data;
    }

    @Override
    public String toString() {
        return "todo: no impl";
    }

    public static Base64 parse(String str) {

        return null; // todo ...
    }

}
