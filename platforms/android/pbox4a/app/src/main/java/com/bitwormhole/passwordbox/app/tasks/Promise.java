package com.bitwormhole.passwordbox.app.tasks;

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

    public Promise<T> Then(PromiseCallback<T> callback) {
        PromiseCallbackHolder<T> holder = new PromiseCallbackHolder<>(callback);
        PromiseCallbackChain<T> chain = this.pc.getChain();
        chain.add(holder.handleThen());
        return this;
    }

    public Promise<T> Catch(PromiseCallback<T> callback) {
        PromiseCallbackHolder<T> holder = new PromiseCallbackHolder<>(callback);
        PromiseCallbackChain<T> chain = this.pc.getChain();
        chain.add(holder.handleCatch());
        return this;
    }

    public Promise<T> Finally(PromiseCallback<T> callback) {
        PromiseCallbackHolder<T> holder = new PromiseCallbackHolder<>(callback);
        PromiseCallbackChain<T> chain = this.pc.getChain();
        chain.add(holder.handleFinally());
        return this;
    }

    public void start() {
        this.pc.start();
    }
}
