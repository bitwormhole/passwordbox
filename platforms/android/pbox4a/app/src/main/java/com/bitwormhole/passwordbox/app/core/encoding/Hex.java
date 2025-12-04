package com.bitwormhole.passwordbox.app.core.encoding;

import androidx.annotation.NonNull;

import java.io.ByteArrayOutputStream;
import java.io.IOException;

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

    public void getData(ByteArrayOutputStream out) {
        try {
            out.write(this.data);
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    @NonNull
    @Override
    public String toString() {
        StringBuilder out = new StringBuilder();
        theCODEC.encode(this.data, out);
        return out.toString();
    }


    public static String toString(Hex h) {
        if (h == null) {
            return "";
        }
        return h.toString();
    }

    public static Hex parse(String str) {
        ByteArrayOutputStream out = new ByteArrayOutputStream();
        theCODEC.decode(str, out);
        return new Hex(out.toByteArray());
    }


    /// ////////////////////////////////////////////////////////////////////////////////////////////


    private static class myInnerCODEC {

        private final char[] char_set_16 = "0123456789abcdef".toCharArray();

        void encode4bits(int b, StringBuilder dst) {
            dst.append(char_set_16[b & 0x0f]);
        }

        void encode(byte[] src, StringBuilder dst) {
            if (src == null) {
                return;
            }
            for (byte b : src) {
                this.encode4bits(b >> 4, dst);
                this.encode4bits(b, dst);
            }
        }

        void decode(String src, ByteArrayOutputStream dst) {
            if (src == null) {
                return;
            }
            final char[] array = src.toCharArray();
            final int len = array.length;
            int hex, tmp;
            hex = tmp = 0;
            for (int index = 0; index < len; index++) {
                hex = this.valueOfHex(array[index]);
                if ((index & 0x01) == 0) {
                    tmp = hex << 4;
                } else {
                    dst.write(tmp | hex);
                }
            }
        }

        private int valueOfHex(char ch) {
            if (('0' <= ch) && (ch <= '9')) {
                return (ch - '0');
            } else if (('a' <= ch) && (ch <= 'f')) {
                return (ch - 'a') + 0x0a;
            } else if (('A' <= ch) && (ch <= 'F')) {
                return (ch - 'A') + 0x0a;
            }
            throw new NumberFormatException("bad hex char '" + ch + "'");
        }
    }

    private final static myInnerCODEC theCODEC = new myInnerCODEC();
}
