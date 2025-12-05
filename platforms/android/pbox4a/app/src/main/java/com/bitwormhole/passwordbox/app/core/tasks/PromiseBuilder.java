package com.bitwormhole.passwordbox.app.core.tasks;

import android.app.Activity;
import android.content.Context;

import java.util.concurrent.Executor;

public final class PromiseBuilder<T> {

    private Executor backgroundExecutor;
    private Executor foregroundExecutor;
    private PromiseContext<T> context;
    private PromiseTaskList<T> tasks;
    private Throwable error;
    private T result;

    public PromiseBuilder(Context ctx) {

        PromiseContext<T> pc = new PromiseContext<>();
        PromiseTaskList<T> tasklist = new PromiseTaskList<>(pc);
        PromiseCallbackChain<T> chain = new PromiseCallbackChain<>(pc);

        pc.setContext(ctx);
        pc.setChain(chain);
        pc.setTasks(tasklist);

        this.context = pc;
        this.tasks = tasklist;
    }

    public PromiseBuilder<T> reset() {
        this.backgroundExecutor = null;
        this.foregroundExecutor = null;
        this.tasks = null;
        return this;
    }


    public PromiseBuilder<T> addTask(PromiseTask<T> task) {
        PromiseTaskList<T> list = this.tasks;
        if (list == null) {
            list = new PromiseTaskList<>(this.context);
            this.tasks = list;
        }
        list.addTask(task);
        return this;
    }

    public Promise<T> build() {

        innerDefaultGetter<T> getter = new innerDefaultGetter<>(this);
        PromiseContext<T> pc = getter.getPC();

        pc.setContext(getter.getContext());
        pc.setBackgroundExecutor(getter.getBgExecutor());
        pc.setForegroundExecutor(getter.getFgExecutor());
        pc.setTasks(getter.getTaskList());
        pc.setError(this.error);
        pc.setResult(this.result);

        return new Promise<>(pc);
    }


    private final static class innerDefaultGetter<T> {

        final PromiseBuilder<T> builder;

        innerDefaultGetter(PromiseBuilder<T> b) {
            this.builder = b;
        }

        Context getContext() {
            Context ctx = builder.getContext();
            if (ctx == null) {
                ctx = this.loadContext();
                builder.setContext(ctx);
            }
            return ctx;
        }

        Executor getFgExecutor() {
            Executor fg = builder.getForegroundExecutor();
            if (fg == null) {
                fg = this.loadFgExecutor();
                builder.setForegroundExecutor(fg);
            }
            return fg;
        }

        Executor getBgExecutor() {
            Executor bg = builder.getBackgroundExecutor();
            if (bg == null) {
                bg = this.loadBgExecutor();
                builder.setBackgroundExecutor(bg);
            }
            return bg;
        }

        PromiseContext<T> getPC() {
            PromiseContext<T> pc = builder.context;
            if (pc == null) {
                pc = this.loadPC();
                builder.context = pc;
            }
            return pc;
        }

        PromiseTaskList<T> getTaskList() {
            PromiseTaskList<T> list = builder.getTasks();
            if (list == null) {
                list = this.loadTaskList();
                builder.setTasks(list);
            }
            return list;
        }

        Context loadContext() {
            throw new RuntimeException("no default context");
        }

        Executor loadFgExecutor() {
            Context ctx = this.getContext();
            if (ctx instanceof Activity) {
                Activity act = (Activity) ctx;
                return new ActivityForegroundExecutor(act);
            } else {
                return ForegroundThreadPool.getExecutor();
            }
        }

        Executor loadBgExecutor() {
            return BackgroundThreadPool.getExecutor();
        }

        PromiseContext<T> loadPC() {
            return new PromiseContext<>();
        }


        PromiseTaskList<T> loadTaskList() {
            PromiseContext<T> pc = this.getPC();
            return new PromiseTaskList<>(pc);
        }
    }

    public PromiseTaskList<T> getTasks() {
        return tasks;
    }

    public PromiseBuilder<T> setTasks(PromiseTaskList<T> tasks) {
        this.tasks = tasks;
        return this;
    }

    public Context getContext() {
        return context.getContext();
    }

    public PromiseBuilder<T> setContext(Context context) {
        this.context.setContext(context);
        return this;
    }


    public Throwable getError() {
        return error;
    }

    public PromiseBuilder<T> setError(Throwable error) {
        this.error = error;
        return this;
    }

    public T getResult() {
        return result;
    }

    public PromiseBuilder<T> setResult(T result) {
        this.result = result;
        return this;
    }

    public Executor getForegroundExecutor() {
        return foregroundExecutor;
    }

    public PromiseBuilder<T> setForegroundExecutor(Executor foregroundExecutor) {
        this.foregroundExecutor = foregroundExecutor;
        return this;
    }

    public Executor getBackgroundExecutor() {
        return backgroundExecutor;
    }

    public PromiseBuilder<T> setBackgroundExecutor(Executor backgroundExecutor) {
        this.backgroundExecutor = backgroundExecutor;
        return this;
    }
}
