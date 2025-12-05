package com.bitwormhole.passwordbox.app.core.tasks;

import android.content.Context;

import java.util.concurrent.Executor;

public final class PromiseContext<T> {

    private Executor backgroundExecutor;
    private Executor foregroundExecutor;
    private long foregroundTID; // the foreground-thread-id
    private Context context;
    private T result;
    private boolean abort;
    private Throwable error;


    private PromiseCallbackChain<T> chain;
    private PromiseTaskList<T> tasks;

    public PromiseContext() {
    }

    public PromiseCallbackChain<T> getChain() {
        return chain;
    }

    public void setChain(PromiseCallbackChain<T> chain) {
        this.chain = chain;
    }

    public Throwable getError() {
        return error;
    }

    public void setError(Throwable error) {
        this.error = error;
    }

    public T getResult() {
        return result;
    }

    public void setResult(T result) {
        this.result = result;
    }

    public Context getContext() {
        return context;
    }

    public void setContext(Context context) {
        this.context = context;
    }

    public Executor getForegroundExecutor() {
        return foregroundExecutor;
    }

    public void setForegroundExecutor(Executor foregroundExecutor) {
        this.foregroundExecutor = foregroundExecutor;
    }

    public Executor getBackgroundExecutor() {
        return backgroundExecutor;
    }

    public void setBackgroundExecutor(Executor backgroundExecutor) {
        this.backgroundExecutor = backgroundExecutor;
    }

    public PromiseTaskList<T> getTasks() {
        return tasks;
    }

    public void setTasks(PromiseTaskList<T> tasks) {
        this.tasks = tasks;
    }


    public long getForegroundTID() {
        return foregroundTID;
    }

    public void setForegroundTID(long foregroundTID) {
        this.foregroundTID = foregroundTID;
    }

    public boolean isAbort() {
        return abort;
    }

    public void setAbort(boolean abort) {
        this.abort = abort;
    }

    public void start() {
        this.foregroundTID = Thread.currentThread().getId();
        this.tasks.execute();
    }
}
