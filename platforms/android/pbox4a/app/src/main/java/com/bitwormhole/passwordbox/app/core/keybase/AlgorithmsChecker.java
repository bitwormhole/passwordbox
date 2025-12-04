package com.bitwormhole.passwordbox.app.core.keybase;

import java.util.ArrayList;
import java.util.List;

public final class AlgorithmsChecker {


    private final AlgorithmsCheckerInner inner = new AlgorithmsCheckerInner();

    public BlockingOption[] listBlockingOptions() {
        return BlockingOption.listAll();
    }

    public PaddingOption[] listPaddingOptions() {
        return PaddingOption.listAll();
    }

    public KeyPairOption[] listKeyPairOptions() {
        return KeyPairOption.listAll();
    }

    public SecretKeyOption[] listSecretKeyOptions() {
        return SecretKeyOption.listAll();
    }

    public SignatureOption[] listSignatureOptions() {
        return SignatureOption.listAll();
    }

    /// ////////////////////////////////////////////////////////////////////////////////////////////

    // 检查 cipher 算法是否可用
    public List<CipherAlgorithmSelector> checkCipherAlgorithms(String provider) {

        SecretKeyOption[] allSKs = this.listSecretKeyOptions();
        KeyPairOption[] allPairs = this.listKeyPairOptions();
        PaddingOption[] allPaddings = this.listPaddingOptions();
        BlockingOption[] allBlocks = this.listBlockingOptions();

        List<CipherAlgorithmSelector> list = new ArrayList<>();
        for (BlockingOption blocking : allBlocks) {
            for (PaddingOption padding : allPaddings) {
                // SK
                for (SecretKeyOption sk : allSKs) {
                    CipherAlgorithmSelector sel = new CipherAlgorithmSelector();
                    sel.setBlocking(blocking);
                    sel.setPadding(padding);
                    sel.setProvider(provider);
                    sel.setSk(sk);

                    this.inner.checkWithSelector(sel);
                    list.add(sel);
                }
                // PK
                for (KeyPairOption pk : allPairs) {
                    CipherAlgorithmSelector sel = new CipherAlgorithmSelector();
                    sel.setBlocking(blocking);
                    sel.setPadding(padding);
                    sel.setProvider(provider);
                    sel.setPk(pk);

                    this.inner.checkWithSelector(sel);
                    list.add(sel);
                }
            }
        }
        return list;
    }

    // 检查 signature 算法是否可用

    public List<SignatureAlgorithmSelector> checkSignatureAlgorithms(String provider) {
        SignatureOption[] all = this.listSignatureOptions();
        List<SignatureAlgorithmSelector> list = new ArrayList<>();
        for (SignatureOption opt : all) {
            SignatureAlgorithmSelector sel = new SignatureAlgorithmSelector();
            sel.setOption(opt);
            sel.setProvider(provider);
            this.inner.checkWithSelector(sel);
            list.add(sel);
        }
        return list;
    }

    // 检查 secret-key-gen 算法是否可用

    public List<SecretKeyAlgorithmSelector> checkSecretKeyAlgorithms(String provider) {
        SecretKeyOption[] src = this.listSecretKeyOptions();
        List<SecretKeyAlgorithmSelector> dst = new ArrayList<>();
        for (SecretKeyOption opt : src) {
            SecretKeyAlgorithmSelector sel = new SecretKeyAlgorithmSelector();
            sel.setOption(opt);
            sel.setProvider(provider);
            this.inner.checkWithSelector(sel);
            dst.add(sel);
        }
        return dst;
    }

    // 检查 key-pair-gen 算法是否可用

    public List<KeyPairAlgorithmSelector> checkKeyPairAlgorithms(String provider) {

        KeyPairOption[] src = this.listKeyPairOptions();
        List<KeyPairAlgorithmSelector> dst = new ArrayList<>();

        for (KeyPairOption opt : src) {
            KeyPairAlgorithmSelector sel = new KeyPairAlgorithmSelector();
            sel.setOption(opt);
            sel.setProvider(provider);
            this.inner.checkWithSelector(sel);
            dst.add(sel);
        }

        return dst;
    }
}
