package com.bitwormhole.passwordbox.app.core.loggers;

public interface Logger extends ILog {

    void trace(String tag, String format, Object... args);

    void debug(String tag, String format, Object... args);

    void info(String tag, String format, Object... args);

    void warn(String tag, String format, Object... args);

    void error(String tag, String format, Object... args);

    void fatal(String tag, String format, Object... args);

}
