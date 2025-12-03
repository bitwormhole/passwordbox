package com.bitwormhole.passwordbox.app.core.loggers;

public final class LoggerFactoryAgent {

    private static LoggerFactory factory;

    public static void setFactory(LoggerFactory f) {
        if (f == null) {
            return;
        }
        LoggerFactoryAgent.factory = f;
    }

    public static LoggerFactory getFactory() {
        LoggerFactory f = LoggerFactoryAgent.factory;
        if (f == null) {
            f = JavaSELoggerFactory.getFactory();
            LoggerFactoryAgent.factory = f;
        }
        return f;
    }
}
