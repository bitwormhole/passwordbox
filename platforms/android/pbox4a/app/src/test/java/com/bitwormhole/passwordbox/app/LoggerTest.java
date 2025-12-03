package com.bitwormhole.passwordbox.app;

import com.bitwormhole.passwordbox.app.core.loggers.Level;
import com.bitwormhole.passwordbox.app.core.loggers.Logger;
import com.bitwormhole.passwordbox.app.core.loggers.LoggerFactory;
import com.bitwormhole.passwordbox.app.core.loggers.Loggers;

import org.junit.Assert;
import org.junit.Test;

public class LoggerTest {


    @Test
    public void testLogger() {

        Logger l = LoggerFactory.getLogger();
        String tag = "test";

        int i32 = 666;
        double f64 = 3.14159;
        Object obj = new Object();

        l.trace(tag, "%s", true);
        l.debug(tag, "%s / %d", true, i32);
        l.info(tag, "%s / %d / %f", true, i32, f64);
        l.warn(tag, "%s / %d / %f / %c  ", true, i32, f64, 'a');
        l.error(tag, "%s / %d / %f / %c / %s ", true, i32, f64, 'a', "bbb");
        l.fatal(tag, "%s / %d / %f / %c / %s / %s", true, i32, f64, 'a', "bbb", obj);

    }

    @Test
    public void testTimer() {

        final int step = 100;
        final int timeout = 5000;
        Logger l = LoggerFactory.getLogger();
        String tag = "testTimer";


        for (int ttl = timeout; ttl > 0; ttl -= step) {
            long now = System.currentTimeMillis();
            l.info(tag, "now = %d", now);
            sleep(step);
        }
    }

    @Test
    public void testCallback() {
        Loggers.forLevel(Level.WARN, (logger) -> {
            logger.withTag("testCallback").log("%s = %d (%f)", "key.name", 110, 110.23333);
        });
    }

    private static void sleep(long ms) {
        try {
            Thread.sleep(ms);
        } catch (InterruptedException e) {
            throw new RuntimeException(e);
        }
    }
}
