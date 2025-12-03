package com.bitwormhole.passwordbox.app.core.loggers;

public class CommonLogger implements Logger {

    private final LogContext context;

    public CommonLogger(LogContext ctx) {
        this.context = ctx;
    }

    @Override
    public void log(Log l) {
        LogWriter wtr = this.context.getWriter();
        wtr.write(l);
    }

    @Override
    public void log(String format, Object... args) {

        Log rec = new Log();
        LogMessageBuilder builder = new LogMessageBuilder();
        LogContext ctx = this.context;

        builder.setLog(rec);
        builder.setContext(ctx);
        builder.setFormat(format);
        builder.setArguments(args);

        rec.setContext(ctx);
        rec.setTime(System.currentTimeMillis());
        rec.setTag(ctx.getTag());
        rec.setLevel(ctx.getLevel());
        rec.setSource(ctx.getSource());

        rec.setMessage(builder.build());

        this.log(rec);
    }

    @Override
    public ILog withLevel(Level level) {
        this.context.setLevel(level);
        return this;
    }

    @Override
    public ILog withSource(String src) {
        this.context.setSource(src);
        return this;
    }

    @Override
    public ILog withTag(String tag) {
        this.context.setTag(tag);
        return this;
    }


    @Override
    public void forLevel(Level level, LoggerCallback callback) {
        if (level == null || callback == null) {
            return;
        }
        LogContext c1 = this.context;
        int want = Level.getOrderOf(level);
        int have = Level.getOrderOf(c1.getLevel());
        if (want < have) {
            return; // ignore this log
        }
        LogContext c2 = c1.copy();
        c2.setFacade(new CommonLogger(c2));
        c2.setLevel(level);
        callback.onCallback(c2.getFacade());
    }

    @Override
    public LogContext getContext() {
        return this.context;
    }

    @Override
    public void trace(String tag, String format, Object... args) {
        this.forLevel(Level.TRACE, (logger) -> {
            logger.withTag(tag).log(format, args);
        });
    }

    @Override
    public void debug(String tag, String format, Object... args) {
        this.forLevel(Level.DEBUG, (logger) -> {
            logger.withTag(tag).log(format, args);
        });
    }

    @Override
    public void info(String tag, String format, Object... args) {
        this.forLevel(Level.INFO, (logger) -> {
            logger.withTag(tag).log(format, args);
        });
    }

    @Override
    public void warn(String tag, String format, Object... args) {
        this.forLevel(Level.WARN, (logger) -> {
            logger.withTag(tag).log(format, args);
        });
    }

    @Override
    public void error(String tag, String format, Object... args) {
        this.forLevel(Level.ERROR, (logger) -> {
            logger.withTag(tag).log(format, args);
        });
    }

    @Override
    public void fatal(String tag, String format, Object... args) {
        this.forLevel(Level.FATAL, (logger) -> {
            logger.withTag(tag).log(format, args);
        });
    }
}
