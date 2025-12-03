package com.bitwormhole.passwordbox.app.core.loggers;

public class Log {

    private long time;
    private Level level;
    private String source;
    private String tag;
    private String message;

    // private LogMessageBuilder messageBuilder;
    private LogContext context;

    public Log() {
    }

    public Level getLevel() {
        return level;
    }

    public void setLevel(Level level) {
        this.level = level;
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

    public long getTime() {
        return time;
    }

    public void setTime(long time) {
        this.time = time;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    /*
    public LogMessageBuilder getMessageBuilder() {
        return messageBuilder;
    }

    public void setMessageBuilder(LogMessageBuilder messageBuilder) {
        this.messageBuilder = messageBuilder;
    }
    *
     */

    public LogContext getContext() {
        return context;
    }

    public void setContext(LogContext context) {
        this.context = context;
    }
}
