package com.bitwormhole.passwordbox.app.core.loggers;

import com.bitwormhole.passwordbox.app.core.time.Time;

import java.io.ByteArrayOutputStream;
import java.io.PrintStream;
import java.nio.charset.StandardCharsets;

public final class DefaultLogFormatter implements LogFormatter {

    private final ByteArrayOutputStream buffer;
    private final PrintStream ps;

    private DefaultLogFormatter() {
        this.buffer = new ByteArrayOutputStream();
        this.ps = new PrintStream(this.buffer);
    }


    public synchronized String format(Log log, String format, Object... args) {
        String head = this.formatHead(log);
        String body = this.formatBody(format, args);
        return head + ' ' + body;
    }

    private static String formatLevel(Level l) {
        if (l == null) {
            return "(level_undef)";
        }
        if (l == Level.INFO || l == Level.WARN) {
            return l + ".";
        }
        return l.toString();
    }

    private static String formatTime(long ms) {
        Time t = new Time(ms);
        return t.toString();
    }


    private String formatHead(Log log) {
        if (log == null) {
            return "";
        }

        Level level = log.getLevel();
        long time = log.getTime();
        StringBuilder builder = new StringBuilder();

        builder.append(formatTime(time)).append(' ');
        builder.append('[').append(formatLevel(level)).append("] ");
        builder.append('(').append(log.getSource()).append(')');
        builder.append('(').append(log.getTag()).append(')');

        return builder.toString();
    }

    private String formatBody(String format, Object... args) {
        this.ps.format(format, args);
        this.ps.flush();
        String str = this.buffer.toString(StandardCharsets.UTF_8);
        this.buffer.reset();
        return str;
    }


    private static DefaultLogFormatter theSingleInstance;

    public static DefaultLogFormatter getInstance() {
        DefaultLogFormatter i = theSingleInstance;
        if (i == null) {
            i = new DefaultLogFormatter();
            theSingleInstance = i;
        }
        return i;
    }
}
