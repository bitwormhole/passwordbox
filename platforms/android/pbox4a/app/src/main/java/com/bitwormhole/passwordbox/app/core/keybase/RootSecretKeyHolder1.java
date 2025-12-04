package com.bitwormhole.passwordbox.app.core.keybase;

import android.content.Context;

import com.bitwormhole.passwordbox.app.core.store.HomeDir;

import java.io.File;

import javax.crypto.Cipher;
import javax.crypto.SecretKey;


/**********************************************************************
 * RootSecretKeyHolder1 提供一个可以导出的 AES 密钥
 * */


public class RootSecretKeyHolder1 implements SecretKeyHolder {

    private final File file;

    public RootSecretKeyHolder1(Context ctx, String path) {
        this.file = computeFilePath(ctx, path);
    }

    private static File computeFilePath(Context ctx, String path) {
        HomeDir home = HomeDir.getInstance(ctx);
        File dir = home.getHomeDir();
        return new File(dir, path);
    }

    @Override
    public boolean exists() {
        return false;
    }

    @Override
    public SecretKeyHolder generate() {
        return null;
    }

    @Override
    public SecretKeyHolder load() {
        return null;
    }

    @Override
    public void delete() {

    }

    @Override
    public SecretKey getKey() {
        return null;
    }

    @Override
    public Cipher getCipher(SecretKeyCiphering ciphering) {
        return null;
    }
}
