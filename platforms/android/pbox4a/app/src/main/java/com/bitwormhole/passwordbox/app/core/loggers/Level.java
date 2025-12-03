package com.bitwormhole.passwordbox.app.core.loggers;

public enum Level {

    MIN,

    TRACE, DEBUG, INFO, WARN, ERROR, FATAL,

    MAX,

    ;

    public static int getOrderOf(Level l) {
        if (l == null) {
            return 0;
        }
        return l.ordinal();
    }
}
