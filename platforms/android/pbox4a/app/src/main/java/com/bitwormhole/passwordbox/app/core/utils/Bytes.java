package com.bitwormhole.passwordbox.app.core.utils;

public final class Bytes {

    private Bytes() {
    }

    public static String toString(String prefix, byte[] b) {
        return innerToString(prefix, b);
    }

    public static String toString(byte[] b) {
        return innerToString("", b);
    }

    private static String innerToString(String prefix, byte[] b) {
        StringBuilder builder = new StringBuilder();
        if (prefix != null) {
            builder.append(prefix);
        }
        builder.append('[');
        if (b == null) {
            builder.append("null").append(']');
            return builder.toString();
        }

        int count = 0;
        for (byte n : b) {
            if (count > 0) {
                builder.append(',');
            }
            String x = Integer.toHexString(n & 0xff);
            if (x.length() == 1) {
                builder.append('0');
            }
            builder.append(x);
            count++;
        }

        builder.append(']');
        return builder.toString();
    }
}
