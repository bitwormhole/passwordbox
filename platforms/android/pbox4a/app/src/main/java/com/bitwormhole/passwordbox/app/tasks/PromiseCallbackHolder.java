package com.bitwormhole.passwordbox.app.tasks;

public class PromiseCallbackHolder<T> {

    private PromiseCallback<T> callback;

    public PromiseCallbackHolder(PromiseCallback<T> cb) {
        this.callback = cb;
    }

    public PromiseCallback<T> getCallback() {
        return callback;
    }

    public void setCallback(PromiseCallback<T> callback) {
        this.callback = callback;
    }

    public PromiseCallbackHolder<T> handleThen() {
        // todo
        return this;
    }

    public PromiseCallbackHolder<T> handleCatch() {
        // todo
        return this;
    }

    public PromiseCallbackHolder<T> handleFinally() {
        // todo
        return this;
    }

}
