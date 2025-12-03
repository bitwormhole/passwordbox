package com.bitwormhole.passwordbox.app.core.loggers;

import java.io.PrintStream;

public final class JavaSELoggerFactory extends LoggerFactory {

    private JavaSELoggerFactory() {
    }

    @Override
    public Logger createNewLogger() {
        CommonLoggerBuilder builder = new CommonLoggerBuilder();
        builder.setTag("NO_TAG");
        builder.setSource("NO_SRC");
        builder.setWriter(new myInnerLogWriter());
        builder.setGate(Level.DEBUG);
        return builder.build();
    }

    private static boolean isLogReady(Log l) {
        if (l == null) {
            return false;
        }
        return (l.getMessage() != null);
    }

    private static class myInnerLogWriter implements LogWriter {
        @Override
        public void write(Log l) {
            if (!isLogReady(l)) {
                return;
            }

            String msg = l.getMessage();
            int o1 = Level.getOrderOf(l.getLevel());
            int o2 = Level.getOrderOf(Level.WARN);

            PrintStream ps = System.out;
            if (o1 >= o2) {
                ps = System.err;
            }
            ps.println(msg);
        }
    }

    public static LoggerFactory getFactory() {
        return new JavaSELoggerFactory();
    }

    public static LogWriter createWriter() {
        return new myInnerLogWriter();
    }
}
