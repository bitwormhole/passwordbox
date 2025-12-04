package com.bitwormhole.passwordbox.app.core.keybase;

import java.security.KeyPairGenerator;
import java.security.Provider;
import java.security.Signature;
import java.security.Signer;

import javax.crypto.Cipher;
import javax.crypto.KeyGenerator;

final class AlgorithmsCheckerInner {


    public void checkWithSelector(SignatureAlgorithmSelector sel) {
        try {
            String provider = sel.getProvider();
            String algorithm = sel.toString();
            Signature sign;

            if (provider == null) {
                sign = Signature.getInstance(algorithm);
                Provider pro = sign.getProvider();
                provider = pro.getName();
            } else {
                sign = Signature.getInstance(algorithm, provider);
                Provider pro = sign.getProvider();
                provider = pro.getName();
            }
            sel.setProvider(provider);
            sel.setSupported(true);
        } catch (Exception e) {
            //  throw new RuntimeException(e);
            sel.setSupported(false);
            sel.setError(e);
        }
    }

    public void checkWithSelector(CipherAlgorithmSelector sel) {
        try {
            String provider = sel.getProvider();
            String algorithm = sel.toString();
            Cipher cipher;
            if (provider == null) {
                cipher = Cipher.getInstance(algorithm);
                Provider pro = cipher.getProvider();
                sel.setProvider(pro.getName());
            } else {
                cipher = Cipher.getInstance(algorithm, provider);
                Provider pro = cipher.getProvider();
                sel.setProvider(pro.getName());
            }
            sel.setSupported(true);
        } catch (Exception e) {
            //       throw new RuntimeException(e);
            sel.setError(e);
            sel.setSupported(false);
        }
    }

    public void checkWithSelector(KeyPairAlgorithmSelector sel) {
        try {
            String provider = sel.getProvider();
            String algorithm = sel.algorithm();
            KeyPairGenerator kpg;
            if (provider == null) {
                kpg = KeyPairGenerator.getInstance(algorithm);
                Provider pro = kpg.getProvider();
                sel.setProvider(pro.getName());
            } else {
                kpg = KeyPairGenerator.getInstance(algorithm, provider);
                Provider pro = kpg.getProvider();
                sel.setProvider(pro.getName());
            }
            sel.setSupported(true);
        } catch (Exception e) {
            sel.setError(e);
            sel.setSupported(false);
        }
    }

    public void checkWithSelector(SecretKeyAlgorithmSelector sel) {
        try {
            String provider = sel.getProvider();
            String algorithm = sel.algorithm();
            KeyGenerator kg;

            if (provider == null) {
                kg = KeyGenerator.getInstance(algorithm);
                Provider pro = kg.getProvider();
                sel.setProvider(pro.getName());
            } else {
                kg = KeyGenerator.getInstance(algorithm, provider);
                Provider pro = kg.getProvider();
                sel.setProvider(pro.getName());
            }
            sel.setSupported(true);
        } catch (Exception e) {
            sel.setError(e);
            sel.setSupported(false);
        }
    }
}
