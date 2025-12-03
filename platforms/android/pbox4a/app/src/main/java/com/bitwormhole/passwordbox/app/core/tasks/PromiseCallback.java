package com.bitwormhole.passwordbox.app.core.tasks;

public interface PromiseCallback<T> {

    Promise<T> handle(Promise<T> p);

}
