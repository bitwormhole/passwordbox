package com.bitwormhole.passwordbox.app.core.loggers;

public final class LogContext {

    private LogWriter writer;
    private String source;
    private String tag;
    private Level level;
    private Level gate;
    private Logger facade;
    private LogFormatter formatter;

    public LogContext() {
    }

    private LogContext(LogContext c1) {
        this.facade = c1.facade;
        this.formatter = c1.formatter;
        this.gate = c1.gate;
        this.level = c1.level;
        this.source = c1.source;
        this.tag = c1.tag;
        this.writer = c1.writer;
    }

    public Level getGate() {
        return gate;
    }

    public void setGate(Level gate) {
        this.gate = gate;
    }

    public LogFormatter getFormatter() {
        return formatter;
    }

    public void setFormatter(LogFormatter formatter) {
        this.formatter = formatter;
    }

    public Logger getFacade() {
        return facade;
    }

    public void setFacade(Logger facade) {
        this.facade = facade;
    }

    public LogWriter getWriter() {
        return writer;
    }

    public void setWriter(LogWriter writer) {
        this.writer = writer;
    }

    public String getSource() {
        return source;
    }

    public void setSource(String source) {
        this.source = source;
    }

    public String getTag() {
        return tag;
    }

    public void setTag(String tag) {
        this.tag = tag;
    }

    public Level getLevel() {
        return level;
    }

    public void setLevel(Level level) {
        this.level = level;
    }

    public LogContext copy() {
        return new LogContext(this);
    }
}
