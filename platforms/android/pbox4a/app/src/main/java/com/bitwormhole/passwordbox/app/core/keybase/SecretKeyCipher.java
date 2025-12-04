package com.bitwormhole.passwordbox.app.core.keybase;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.security.SecureRandom;

import javax.crypto.Cipher;
import javax.crypto.CipherOutputStream;
import javax.crypto.SecretKey;

public final class SecretKeyCipher {

    private SecretKeyCipher() {
    }

    public static void decrypt(SecretKeyCiphering ciphering) {

        prepare(ciphering);

        Cipher cipher = ciphering.getCipher();
        SecureRandom rand = ciphering.getRandom();
        SecretKey key = ciphering.getKey();
        byte[] input = ciphering.getEncrypted();
        byte[] iv = ciphering.getIv();

        try {
            if (rand == null) {
                cipher.init(Cipher.DECRYPT_MODE, key);
            } else {
                cipher.init(Cipher.DECRYPT_MODE, key, rand);
            }
            byte[] output = doCipher(cipher, input);
            ciphering.setPlain(output);
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

    public static void encrypt(SecretKeyCiphering ciphering) {

        prepare(ciphering);

        Cipher cipher = ciphering.getCipher();
        SecureRandom rand = ciphering.getRandom();
        SecretKey key = ciphering.getKey();
        byte[] input = ciphering.getPlain();
        byte[] iv = ciphering.getIv();

        try {
            if (rand == null) {
                cipher.init(Cipher.ENCRYPT_MODE, key);
            } else {
                cipher.init(Cipher.ENCRYPT_MODE, key, rand);
            }
            byte[] output = doCipher(cipher, input);
            ciphering.setEncrypted(output);
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }


    /// ////////////////////////////////////////////////////////////////////////////////////////////


    private static byte[] doCipher(Cipher cipher, byte[] input) throws IOException {
        ByteArrayOutputStream output = new ByteArrayOutputStream();
        CipherOutputStream cos = new CipherOutputStream(output, cipher);
        cos.write(input);
        cos.flush();
        return output.toByteArray();
    }


    private static void prepare(SecretKeyCiphering ciphering) {

        Cipher cipher = ciphering.getCipher();
        SecretKey key = ciphering.getKey();
        SecretKeyHolder holder = ciphering.getKeyHolder();
        SecureRandom rand = ciphering.getRandom();
        PaddingOption padding = ciphering.getPadding();
        BlockingOption blocking = ciphering.getBlocking();

        if (padding == null) {
            padding = PaddingOption.PKCS5Padding;
            ciphering.setPadding(padding);
        }

        if (blocking == null) {
            blocking = BlockingOption.ECB;
            ciphering.setBlocking(blocking);
        }

        if (key == null) {
            key = holder.getKey();
            ciphering.setKey(key);
        }

        if (rand == null) {
            rand = RandomAgent.getDefaultRandom();
            ciphering.setRandom(rand);
        }

        if (cipher == null) {
            cipher = holder.getCipher(ciphering);
            ciphering.setCipher(cipher);
        }
    }


}
