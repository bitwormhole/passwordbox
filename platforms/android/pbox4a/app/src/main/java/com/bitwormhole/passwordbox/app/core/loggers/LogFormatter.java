package com.bitwormhole.passwordbox.app.core.loggers;

public interface LogFormatter {

    String format(Log l, String format, Object... args);

}
