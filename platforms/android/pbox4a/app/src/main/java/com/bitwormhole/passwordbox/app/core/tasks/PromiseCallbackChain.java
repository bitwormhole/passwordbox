package com.bitwormhole.passwordbox.app.core.tasks;

import android.util.Log;

public final class PromiseCallbackChain<T> {

    private Node<T> first;
    private Node<T> ending;
    private final PromiseContext<T> context;


    public PromiseCallbackChain(PromiseContext<T> pc) {
        this.context = pc;
    }

    private static class Node<T> {
        Node<T> next;
        PromiseCallbackHolder<T> holder;
    }

    private void add(final Node<T> node) {
        if (node == null) {
            return;
        }
        final Node<T> older = this.ending;
        if (older != null) {
            older.next = node;
        } else {
            this.first = node;
        }
        this.ending = node;
    }

    public void addThenCallback(PromiseThenCallback<T> callback) {
        Node<T> node = new Node<>();
        node.holder = new PromiseCallbackHolder<>(callback);
        this.add(node);
    }

    public void addCatchCallback(PromiseCatchCallback<T> callback) {
        Node<T> node = new Node<>();
        node.holder = new PromiseCallbackHolder<>(callback);
        this.add(node);
    }

    public void addFinallyCallback(PromiseFinallyCallback<T> callback) {
        Node<T> node = new Node<>();
        node.holder = new PromiseCallbackHolder<>(callback);
        this.add(node);
    }

    public void addProgressCallback(PromiseProgressCallback<T> callback) {
        Node<T> node = new Node<>();
        node.holder = new PromiseCallbackHolder<>(callback);
        this.add(node);
    }

    private class ResultPost implements Runnable {

        private Promise<T> promise;
        private boolean abort;

        @Override
        public void run() {
            for (; this.hasMore(); ) {
                Node<T> node = this.fetchNextNode();
                if (node == null) {
                    break;
                }
                if (this.abort) {
                    break;
                }
                this.invokeWithNode(node);
            }
        }


        boolean hasError(Promise<T> p) {
            if (p == null) {
                return false;
            }
            return (p.getError() != null);
        }


        private void invokeWithNode(Node<T> node) {

            if (node == null) {
                return;
            }
            final PromiseCallbackHolder<T> holder = node.holder;
            if (holder == null) {
                return;
            }
            final PromiseThenCallback<T> cb_then = holder.getCallbackThen();
            final PromiseCatchCallback<T> cb_catch = holder.getCallbackCatch();
            final PromiseFinallyCallback<T> cb_finally = holder.getCallbackFinally();
            final Promise<T> p1 = this.promise;
            Promise<T> p2 = null;

            try {
                if (!this.hasError(p1)) {
                    if (cb_then != null) {
                        p2 = cb_then.handle(p1);
                    }
                } else {
                    if (cb_catch != null) {
                        p2 = cb_catch.handle(p1);
                    }
                }

                if (cb_finally != null) {
                    p2 = cb_finally.handle(p1);
                }

                if (p2 != null) {
                    p1.setResult(p2.getResult());
                    p1.setError(p2.getError());
                }

            } catch (Exception e) {
                p1.setError(e);
            }
        }

        boolean hasMore() {
            return (first != null);
        }

        Node<T> fetchNextNode() {
            Node<T> n1 = first;
            Node<T> n2 = null;
            if (n1 == null) {
                return null;
            }
            n2 = n1.next;
            if (n2 == null) {
                first = null;
                ending = null;
            } else {
                first = n2;
            }
            return n1;
        }

        public void init(Promise<T> p) {
            this.promise = p;
        }
    }

    private class ProgressPost implements Runnable {

        private Promise<T> promise;

        @Override
        public void run() {
            Node<T> node;
            for (node = first; node != null; node = node.next) {
                this.invokeProgressHandler(node);
            }
        }

        private void invokeProgressHandler(Node<T> node) {
            if (node == null) {
                return;
            }
            PromiseCallbackHolder<T> holder = node.holder;
            if (holder == null) {
                return;
            }
            PromiseProgressCallback<T> callback = holder.getCallbackProgress();
            if (callback == null) {
                return;
            }
            try {
                callback.handle(this.promise);
            } catch (Exception e) {
                //    throw new RuntimeException(e);
                final String tag = this.getClass().getName();
                Log.e(tag, "in_progress_err: " + e.getMessage());
            }
        }

        public void init(Promise<T> p) {
            this.promise = p;
        }
    }


    public void postResult(Promise<T> p) {
        ResultPost post = new ResultPost();
        post.init(p);
        this.context.getForegroundExecutor().execute(post);
    }

    public void postProgress(Promise<T> p) {
        ProgressPost post = new ProgressPost();
        post.init(p);
        this.context.getForegroundExecutor().execute(post);
    }
}
