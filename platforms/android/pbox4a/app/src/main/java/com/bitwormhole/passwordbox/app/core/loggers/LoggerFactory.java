package com.bitwormhole.passwordbox.app.core.loggers;

public abstract class LoggerFactory {

    public abstract Logger createNewLogger();


    public static LoggerFactory getInstance() {
        return LoggerFactoryAgent.getFactory();
    }


    public static Logger getLogger() {
        LoggerFactory f = getInstance();
        return f.createNewLogger();
    }

}
