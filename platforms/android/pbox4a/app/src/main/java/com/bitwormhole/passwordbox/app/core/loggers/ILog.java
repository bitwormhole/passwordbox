package com.bitwormhole.passwordbox.app.core.loggers;

public interface ILog {

    void log(Log l);

    void log(String format, Object... args);

    ILog withLevel(Level level);

    ILog withSource(String src);

    ILog withTag(String tag);

    void forLevel(Level level, LoggerCallback callback);

    LogContext getContext();
}
