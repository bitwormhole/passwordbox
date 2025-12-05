package com.bitwormhole.passwordbox.app;

import android.content.Context;
import android.util.Log;

import androidx.test.ext.junit.runners.AndroidJUnit4;
import androidx.test.platform.app.InstrumentationRegistry;

import com.bitwormhole.passwordbox.app.core.tasks.Promise;
import com.bitwormhole.passwordbox.app.core.tasks.PromiseBuilder;

import org.junit.Assert;
import org.junit.Test;
import org.junit.runner.RunWith;

import java.lang.reflect.Executable;


@RunWith(AndroidJUnit4.class)
public class PromiseTest {

    private Context getContext() {
        return InstrumentationRegistry.getInstrumentation().getTargetContext();
    }

    private class Payload {
        String foo;
        int bar;

        float progress; // 0~1
    }

    @Test
    public void testPromiseWithError() {
        Exception err = new RuntimeException("mock_err: testPromiseWithError");
        innerTestPromise(err, 3000);
    }

    @Test
    public void testPromiseWithoutError() {
        innerTestPromise(null, 2000);
    }


    private void innerTestPromise(Exception err1, long msToSleepAfterDone) {

        final Context ctx = this.getContext();
        final String tag = this.getClass().getName();
        final Payload payload = new Payload();

        PromiseBuilder<Payload> pb = new PromiseBuilder<>(ctx);
        pb.addTask((p) -> {

            Log.i(tag, "promise.run_task() : ...");

            p.postProgress(payload);

            if (err1 == null) {
                return Promise.resolve(ctx, payload);
            }
            return Promise.reject(ctx, err1);
        });

        pb.build().Progress((p) -> {

            Log.i(tag, "promise.progress() : ...");

            return null;

        }).Then((p) -> {

            Log.i(tag, "promise.then() : OK");
            return p;
        }).Catch((p) -> {
            Throwable err = p.getError();
            Log.e(tag, "promise.catch() : error:" + err.getMessage());
            return p;
        }).Finally((p) -> {
            Log.i(tag, "promise.finally() : Done");
            return p;
        }).start();

        try {
            Thread.sleep(msToSleepAfterDone);
        } catch (InterruptedException e) {
            throw new RuntimeException(e);
        }
    }
}
