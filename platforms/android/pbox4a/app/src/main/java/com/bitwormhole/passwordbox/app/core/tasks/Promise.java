package com.bitwormhole.passwordbox.app.core.tasks;

import android.content.Context;

import java.util.concurrent.Executor;

public final class Promise<T> {

    private final PromiseContext<T> pc;

    public Promise(Context ctx) {
        this.pc = new PromiseContext<>();
        this.pc.setContext(ctx);
    }

    public Promise(PromiseContext<T> pContext) {
        if (pContext == null) {
            pContext = new PromiseContext<>();
        }
        this.pc = pContext;
    }


    public PromiseContext<T> getPromiseContext() {
        return this.pc;
    }

    public T getResult() {
        return this.pc.getResult();
    }

    public void setResult(T result) {
        this.pc.setResult(result);
    }

    public Throwable getError() {
        return this.pc.getError();
    }

    public void setError(Throwable error) {
        this.pc.setError(error);
    }

    public Promise<T> Then(PromiseThenCallback<T> callback) {
        PromiseCallbackChain<T> chain = this.pc.getChain();
        chain.addThenCallback(callback);
        return this;
    }

    public Promise<T> Progress(PromiseProgressCallback<T> callback) {
        PromiseCallbackChain<T> chain = this.pc.getChain();
        chain.addProgressCallback(callback);
        return this;
    }

    public Promise<T> Catch(PromiseCatchCallback<T> callback) {
        PromiseCallbackChain<T> chain = this.pc.getChain();
        chain.addCatchCallback(callback);
        return this;
    }

    public Promise<T> Finally(PromiseFinallyCallback<T> callback) {
        PromiseCallbackChain<T> chain = this.pc.getChain();
        chain.addFinallyCallback(callback);
        return this;
    }

    public static <T> Promise<T> reject(Context ctx, Throwable err) {
        PromiseBuilder<T> builder = new PromiseBuilder<>(ctx);
        builder.setError(err);
        return builder.build();
    }

    public static <T> Promise<T> resolve(Context ctx, T res) {
        PromiseBuilder<T> builder = new PromiseBuilder<>(ctx);
        builder.setResult(res);
        return builder.build();
    }

    public static <T> boolean hasError(Promise<T> p) {
        if (p == null) {
            return false;
        }
        return (p.getError() != null);
    }

    public static <T> boolean hasResult(Promise<T> p) {
        if (p == null) {
            return false;
        }
        return (p.getResult() != null) && (p.getError() == null);
    }


    public void start() {
        this.pc.start();
    }

    public void postProgress(T payload) {
        PromiseCallbackChain<T> chain = this.pc.getChain();
        Promise<T> tmp = new Promise<>(this.pc);
        tmp.setResult(payload);
        chain.postProgress(tmp);
    }
}
