package com.bitwormhole.passwordbox.app.core.loggers;

public final class Loggers {

    private Loggers() {
    }

    public static void trace(String tag, String format, Object... args) {
        getLogger().trace(tag, format, args);
    }

    public static void debug(String tag, String format, Object... args) {
        getLogger().debug(tag, format, args);
    }

    public static void info(String tag, String format, Object... args) {
        getLogger().info(tag, format, args);
    }

    public static void warn(String tag, String format, Object... args) {
        getLogger().warn(tag, format, args);
    }

    public static void error(String tag, String format, Object... args) {
        getLogger().error(tag, format, args);
    }

    public static void fatal(String tag, String format, Object... args) {
        getLogger().fatal(tag, format, args);
    }

    public static void forLevel(Level level, LoggerCallback callback) {
        getLogger().forLevel(level, callback);
    }


    /// ////////////////////////////////////////////////////////////////////////////////////////////

    private static Logger theLogger;

    private static Logger getLogger() {
        Logger l = theLogger;
        if (l == null) {
            l = LoggerFactory.getLogger();
            theLogger = l;
        }
        return l;
    }
}
