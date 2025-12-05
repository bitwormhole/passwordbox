package com.bitwormhole.passwordbox.app.core.tasks;

public class PromiseCallbackHolder<T> {

    private final PromiseThenCallback<T> callbackThen;
    private final PromiseCatchCallback<T> callbackCatch;
    private final PromiseFinallyCallback<T> callbackFinally;
    private final PromiseProgressCallback<T> callbackProgress;

    private PromiseCallbackHolder() {
        this.callbackProgress = null;
        this.callbackThen = null;
        this.callbackCatch = null;
        this.callbackFinally = null;
    }


    public PromiseCallbackHolder(PromiseThenCallback<T> c) {
        this.callbackThen = c;

        this.callbackProgress = null;
        //  this.callbackThen = null ;
        this.callbackCatch = null;
        this.callbackFinally = null;
    }

    public PromiseCallbackHolder(PromiseProgressCallback<T> c) {
        this.callbackProgress = c;

        //  this.callbackProgress  = null ;
        this.callbackThen = null;
        this.callbackCatch = null;
        this.callbackFinally = null;
    }

    public PromiseCallbackHolder(PromiseCatchCallback<T> c) {
        this.callbackCatch = c;

        this.callbackProgress = null;
        this.callbackThen = null;
        //   this.callbackCatch  = null ;
        this.callbackFinally = null;
    }

    public PromiseCallbackHolder(PromiseFinallyCallback<T> c) {
        this.callbackFinally = c;

        this.callbackProgress = null;
        this.callbackThen = null;
        this.callbackCatch = null;
        //     this.callbackFinally  = null ;
    }

    public PromiseThenCallback<T> getCallbackThen() {
        return callbackThen;
    }

    public PromiseCatchCallback<T> getCallbackCatch() {
        return callbackCatch;
    }

    public PromiseFinallyCallback<T> getCallbackFinally() {
        return callbackFinally;
    }

    public PromiseProgressCallback<T> getCallbackProgress() {
        return callbackProgress;
    }
}
