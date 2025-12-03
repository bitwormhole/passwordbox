package com.bitwormhole.passwordbox.app.core.keybase;

import android.security.keystore.KeyGenParameterSpec;
import android.security.keystore.KeyProperties;

import java.security.Key;
import java.security.KeyPair;
import java.security.KeyPairGenerator;
import java.security.KeyStore;
import java.security.PrivateKey;
import java.security.PublicKey;
import java.security.cert.Certificate;

public final class KeyPairHolder {

    private KeyPair pair;
    private final String alias;

    public KeyPairHolder(String _alias) {
        this.alias = _alias;
    }

    public boolean exists() {
        try {
            KeyStore ks = KeyStore.getInstance("AndroidKeyStore");
            ks.load(null);
            return ks.containsAlias(this.alias);
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

    public KeyPairHolder generate() {
        try {
            KeyPairGenerator kpg = KeyPairGenerator.getInstance(
                    KeyProperties.KEY_ALGORITHM_RSA, "AndroidKeyStore");
            kpg.initialize(new KeyGenParameterSpec.Builder(
                    this.alias,
                    KeyProperties.PURPOSE_SIGN | KeyProperties.PURPOSE_VERIFY | KeyProperties.PURPOSE_ENCRYPT | KeyProperties.PURPOSE_DECRYPT)
                    .setDigests(KeyProperties.DIGEST_SHA256, KeyProperties.DIGEST_SHA512)
                    .setKeySize(1024 * 2)
                    .build());
            this.pair = kpg.generateKeyPair();
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
        return this;
    }

    public KeyPairHolder load() {
        try {
            KeyStore ks = KeyStore.getInstance("AndroidKeyStore");
            ks.load(null);
            Key key = ks.getKey(this.alias, null);
            Certificate cert = ks.getCertificate(this.alias);
            PrivateKey pri = (PrivateKey) key;
            PublicKey pub = cert.getPublicKey();
            this.pair = new KeyPair(pub, pri);
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
        return this;
    }

    public KeyPairHolder fetch() {
        KeyPairHolder holder = this;
        KeyPair older = holder.pair;
        if (older == null) {
            holder.load();
        }
        return holder;
    }

    public KeyPairHolder open(boolean create) {
        KeyPairHolder holder = this;
        if (!holder.exists()) {
            if (create) {
                holder.generate();
            }
        }
        return holder.fetch();
    }

    public KeyPair getPair() {
        return this.pair;
    }
}
