package com.bitwormhole.passwordbox.app.tasks;

public interface PromiseCallback<T> {

    Promise<T> handle(Promise<T> p);

}
