package com.bitwormhole.passwordbox.app.core.keybase;

import android.security.keystore.KeyGenParameterSpec;
import android.security.keystore.KeyProperties;

import java.security.KeyStore;
import java.security.KeyStoreException;

import javax.crypto.Cipher;
import javax.crypto.KeyGenerator;
import javax.crypto.SecretKey;


/**********************************************************************
 * RootSecretKeyHolder0 提供一个由  AndroidKeyStore 托管的 AES 密钥
 * */


public class RootSecretKeyHolder0 implements SecretKeyHolder {

    private final String alias;

    private SecretKey key;

    public RootSecretKeyHolder0(String _alias) {
        this.alias = _alias;
    }

    @Override
    public boolean exists() {
        try {
            KeyStore ks = KeyStore.getInstance("AndroidKeyStore");
            ks.load(null);
            return ks.containsAlias(this.alias);
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

    @Override
    public SecretKeyHolder generate() {
        try {
            KeyGenParameterSpec.Builder spec_builder = new KeyGenParameterSpec.Builder(this.alias,
                    KeyProperties.PURPOSE_ENCRYPT | KeyProperties.PURPOSE_DECRYPT);
            spec_builder.setKeySize(256);
            spec_builder.setDigests(KeyProperties.DIGEST_SHA256, KeyProperties.DIGEST_SHA512);

            KeyGenerator kg = KeyGenerator.getInstance(KeyProperties.KEY_ALGORITHM_AES, "AndroidKeyStore");
            kg.init(spec_builder.build());
            this.key = kg.generateKey();
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
        return this;
    }

    @Override
    public SecretKeyHolder load() {
        try {
            KeyStore ks = KeyStore.getInstance("AndroidKeyStore");
            ks.load(null);
            KeyStore.Entry entry1 = ks.getEntry(this.alias, null);
            KeyStore.SecretKeyEntry entry2 = (KeyStore.SecretKeyEntry) entry1;
            this.key = entry2.getSecretKey();
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
        return this;
    }

    @Override
    public void delete() {
        // nop
    }

    @Override
    public SecretKey getKey() {
        return this.key;
    }

    @Override
    public Cipher getCipher(SecretKeyCiphering ciphering) {
        String alg1 = this.key.getAlgorithm();
        String blocking = ciphering.getBlocking() + "";
        String padding = ciphering.getPadding() + "";
        try {
            String alg2 = alg1 + '/' + blocking + '/' + padding;
            return Cipher.getInstance(alg2, "AndroidKeyStore");
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }
}
