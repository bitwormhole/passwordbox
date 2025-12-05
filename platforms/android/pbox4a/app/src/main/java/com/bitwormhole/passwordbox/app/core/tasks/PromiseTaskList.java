package com.bitwormhole.passwordbox.app.core.tasks;

public class PromiseTaskList<T> {

    private Node<T> first;
    private Node<T> ending;
    private final PromiseContext<T> context;

    public PromiseTaskList(PromiseContext<T> pc) {
        this.context = pc;
    }

    private static class Node<T> {
        PromiseTask<T> task;
        Node<T> next;
    }

    public void addTask(PromiseTask<T> task) {
        final Node<T> older = this.ending;
        final Node<T> node = new Node<>();
        node.task = task;
        if (older != null) {
            older.next = node;
        } else {
            this.first = node;
        }
        this.ending = node;
    }


    private class Runner implements Runnable {

        @Override
        public void run() {
            Promise<T> p1 = new Promise<>(context);
            try {
                Throwable err = this.execAll(p1);
                if (err == null) {
                    this.onThen(p1);
                } else {
                    this.onCatch(p1, err);
                }
            } catch (Exception e) {
                this.onCatch(p1, e);
            } finally {
                this.onFinally(p1);
            }
        }

        private Throwable execAll(Promise<T> p1) throws Exception {

            Promise<T> p2 = p1;

            p1 = new Promise<>(context);

            for (; ; ) {
                if (!this.hasMore()) {
                    break;
                }
                Node<T> node = this.fetchNextTask();
                p2 = node.task.handle(p1);
                if (p2 == null) {
                    p2 = p1;
                }
                Throwable err = p2.getError();
                T res = p2.getResult();
                if (err != null) {
                    return err;
                }
                if (res != null) {
                    p1.setResult(res);
                }
            }

            return null;
        }

        private void onThen(Promise<T> p) {
            p.setError(null);
            context.getChain().postResult(p);
        }

        private void onCatch(Promise<T> p, Throwable err) {
            p.setError(err);
            context.getChain().postResult(p);
        }

        private void onFinally(Promise<T> p) {
            context.getChain().postResult(p);
        }

        void start() {
            context.getBackgroundExecutor().execute(this);
        }

        void startAuto() {
            long curr = Thread.currentThread().getId();
            long fore = context.getForegroundTID();
            if (curr == fore) {
                this.start();
            } else {
                this.run();
            }
        }

        boolean hasMore() {
            return (first != null);
        }

        Node<T> fetchNextTask() {
            Node<T> next1 = first;
            Node<T> next2 = null;
            if (next1 != null) {
                next2 = next1.next;
                if (next2 == null) {
                    first = null;
                    ending = null;
                } else {
                    first = next2;
                }
            }
            return next1;
        }
    }

    public void execute() {
        Runner r1 = new Runner();
        r1.startAuto();
    }
}
