package com.bitwormhole.passwordbox.app.core.loggers;

public class LogMessageBuilder {

    private LogContext context;
    private Log log;
    private String format;
    private Object[] arguments;

    public LogMessageBuilder() {
    }

    public LogMessageBuilder(LogContext ctx, Log l) {
        this.context = ctx;
        this.log = l;
    }

    public LogMessageBuilder(LogContext ctx, Log l, String fmt, Object[] args) {
        this.context = ctx;
        this.log = l;
        this.format = fmt;
        this.arguments = args;
    }


    public LogContext getContext() {
        return context;
    }

    public void setContext(LogContext context) {
        this.context = context;
    }

    public Log getLog() {
        return log;
    }

    public void setLog(Log log) {
        this.log = log;
    }

    public String getFormat() {
        return format;
    }

    public void setFormat(String format) {
        this.format = format;
    }


    public Object[] getArguments() {
        return arguments;
    }

    public void setArguments(Object[] arguments) {
        this.arguments = arguments;
    }

    public String build() {
        LogFormatter formatter = this.context.getFormatter();
        return formatter.format(this.log, this.format, this.arguments);
    }
}
