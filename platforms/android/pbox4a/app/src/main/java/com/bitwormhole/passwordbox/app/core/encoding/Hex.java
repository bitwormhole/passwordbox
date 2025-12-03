package com.bitwormhole.passwordbox.app.core.encoding;

import kotlin.text.HexFormat;

public final class Hex {

    private final byte[] data;

    public Hex() {
        this.data = new byte[0];
    }

    public Hex(byte[] b) {
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

       // HexFormatter.

        return "todo: no impl";
    }

    public static String toString(Hex h) {
        if (h == null) {
            return "";
        }
        return h.toString();
    }

    public static Hex parse(String str) {
        return null; // todo ...
    }

}
