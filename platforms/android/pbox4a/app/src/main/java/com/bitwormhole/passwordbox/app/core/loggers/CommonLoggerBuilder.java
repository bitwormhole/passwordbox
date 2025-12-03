package com.bitwormhole.passwordbox.app.core.loggers;

public final class CommonLoggerBuilder {

    private LogWriter writer;
    private Level gate; // the gate level
    private Level level; // the default level
    private String tag;
    private String source;
    private LogFormatter formatter;

    public CommonLoggerBuilder() {
    }

    public Logger build() {

        LogContext ctx = new LogContext();
        CommonLoggerBuilder b = this.prepare();
        Logger logger = new CommonLogger(ctx);

        ctx.setLevel(b.level);
        ctx.setTag(b.tag);
        ctx.setSource(b.source);
        ctx.setWriter(b.writer);
        ctx.setFacade(logger);
        ctx.setGate(b.gate);
        ctx.setFormatter(b.formatter);

        return logger;
    }

    private CommonLoggerBuilder prepare() {
        final CommonLoggerBuilder b = this;

        if (b.gate == null) {
            b.gate = Level.INFO;
        }

        if (b.level == null) {
            b.level = Level.INFO;
        }

        if (b.source == null) {
            b.source = Thread.currentThread().getName();
        }

        if (b.tag == null) {
            b.tag = "NO_TAG";
        }

        if (b.formatter == null) {
            b.formatter = DefaultLogFormatter.getInstance();
        }

        if (b.writer == null) {
            b.writer = JavaSELoggerFactory.createWriter();
        }

        return b;
    }


    public LogWriter getWriter() {
        return writer;
    }

    public void setWriter(LogWriter writer) {
        this.writer = writer;
    }


    public Level getLevel() {
        return level;
    }

    public void setLevel(Level level) {
        this.level = level;
    }

    public Level getGate() {
        return gate;
    }

    public void setGate(Level gate) {
        this.gate = gate;
    }

    public String getTag() {
        return tag;
    }

    public void setTag(String tag) {
        this.tag = tag;
    }

    public LogFormatter getFormatter() {
        return formatter;
    }

    public void setFormatter(LogFormatter formatter) {
        this.formatter = formatter;
    }

    public String getSource() {
        return source;
    }

    public void setSource(String source) {
        this.source = source;
    }
}
