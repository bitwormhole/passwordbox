package com.bitwormhole.passwordbox.app.core.utils;

import androidx.annotation.NonNull;

import com.bitwormhole.passwordbox.app.core.encoding.Hex;

import java.io.ByteArrayOutputStream;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;

public final class Sum {

    private final Hex hex;

    private Sum(byte[] b) {
        this.hex = new Hex(b);
    }

    @NonNull
    @Override
    public String toString() {
        return this.hex.toString();
    }

    public byte[] toByteArray() {
        ByteArrayOutputStream out = new ByteArrayOutputStream();
        this.hex.getData(out);
        return out.toByteArray();
    }

    /// ////////////////////////////////////////////////////////////////////////////////////////////


    private static class myInnerSumComputer {

        MessageDigest md;

        void init(String algorithm) {
            try {
                this.md = MessageDigest.getInstance(algorithm);
            } catch (NoSuchAlgorithmException e) {
                throw new RuntimeException(e);
            }
        }

        Sum compute(byte[] data) {
            if (data != null) {
                this.md.update(data);
            }
            byte[] h = this.md.digest();
            return new Sum(h);
        }
    }


    /// ////////////////////////////////////////////////////////////////////////////////////////////


    public static Sum sha1sum(byte[] data) {
        myInnerSumComputer comp = new myInnerSumComputer();
        comp.init("sha-1");
        return comp.compute(data);
    }

    public static Sum sha256sum(byte[] data) {
        myInnerSumComputer comp = new myInnerSumComputer();
        comp.init("sha-256");
        return comp.compute(data);
    }

}
