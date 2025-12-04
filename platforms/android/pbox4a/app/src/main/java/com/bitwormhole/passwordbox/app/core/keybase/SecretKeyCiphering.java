package com.bitwormhole.passwordbox.app.core.keybase;

import java.security.SecureRandom;

import javax.crypto.Cipher;
import javax.crypto.SecretKey;

public class SecretKeyCiphering {

    private SecretKey key;
    private SecretKeyHolder keyHolder;
    private Cipher cipher;
    private SecureRandom random;

    private BlockingOption blocking;
    private PaddingOption padding;
    private byte[] iv;
    private byte[] plain;
    private byte[] encrypted;

    public SecretKeyCiphering() {
    }

    public SecureRandom getRandom() {
        return random;
    }

    public void setRandom(SecureRandom random) {
        this.random = random;
    }

    public SecretKey getKey() {
        return key;
    }

    public void setKey(SecretKey key) {
        this.key = key;
    }

    public BlockingOption getBlocking() {
        return blocking;
    }

    public void setBlocking(BlockingOption blocking) {
        this.blocking = blocking;
    }

    public PaddingOption getPadding() {
        return padding;
    }

    public void setPadding(PaddingOption padding) {
        this.padding = padding;
    }

    public SecretKeyHolder getKeyHolder() {
        return keyHolder;
    }

    public void setKeyHolder(SecretKeyHolder keyHolder) {
        this.keyHolder = keyHolder;
    }

    public Cipher getCipher() {
        return cipher;
    }

    public void setCipher(Cipher cipher) {
        this.cipher = cipher;
    }

    public byte[] getPlain() {
        return plain;
    }

    public void setPlain(byte[] plain) {
        this.plain = plain;
    }

    public byte[] getEncrypted() {
        return encrypted;
    }

    public void setEncrypted(byte[] encrypted) {
        this.encrypted = encrypted;
    }

    public byte[] getIv() {
        return iv;
    }

    public void setIv(byte[] iv) {
        this.iv = iv;
    }
}
