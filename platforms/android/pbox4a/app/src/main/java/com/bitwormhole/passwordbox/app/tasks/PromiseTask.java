package com.bitwormhole.passwordbox.app.tasks;

public interface PromiseTask<T> {

    Promise<T> run(Promise<T> p);

}
