package com.bitwormhole.passwordbox.app.core.tasks;

public interface PromiseTask<T> {

    Promise<T> run(Promise<T> p);

}
