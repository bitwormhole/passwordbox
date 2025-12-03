package com.bitwormhole.passwordbox.app.core.time;

import java.time.LocalDateTime;
import java.time.ZoneOffset;
import java.time.format.DateTimeFormatter;
import java.util.Date;
import java.util.TimeZone;

public class TimeFormatter {

    public static String format(Time t) {
        return formatAll(t);
    }

    public static String formatYMD(Time t) {
        DateTimeFormatter fmt = getFormatterYMD();
        LocalDateTime dt = getDateOf(t);
        return dt.format(fmt);
    }

    public static String formatHMS(Time t) {
        DateTimeFormatter fmt = getFormatterHMS();
        LocalDateTime dt = getDateOf(t);
        return dt.format(fmt);
    }

    public static String formatAll(Time t) {
        DateTimeFormatter fmt = getFormatterFull();
        LocalDateTime dt = getDateOf(t);
        return dt.format(fmt);
    }

    /// ////////////////////////////////////////////////////////////////////////////////////////////


    private static LocalDateTime getDateOf(Time t) {

        // TimeZone zone = TimeZone.getDefault();


        ZoneOffset zoff = ZoneOffset.ofHours(8);
        long ms = t.value();
        long sec = ms / 1000;
        int ns = (int) (ms % 1000) * 1000 * 1000;
        return LocalDateTime.ofEpochSecond(sec, ns, zoff);
    }


    private static DateTimeFormatter makeFormatter(String pattern) {
        return DateTimeFormatter.ofPattern(pattern);
    }


    private static DateTimeFormatter getFormatterYMD() {
        DateTimeFormatter f = formatterYMD;
        if (f == null) {
            f = makeFormatter("yyyy-MM-dd");
            formatterYMD = f;
        }
        return f;
    }

    private static DateTimeFormatter getFormatterHMS() {
        DateTimeFormatter f = formatterHMS;
        if (f == null) {
            f = makeFormatter("HH:mm:ss");
            formatterHMS = f;
        }
        return f;
    }

    private static DateTimeFormatter getFormatterFull() {
        DateTimeFormatter f = formatterFull;
        if (f == null) {
            f = makeFormatter("yyyy-MM-dd HH:mm:ss.SSS");
            formatterFull = f;
        }
        return f;
    }


    /// ////////////////////////////////////////////////////////////////////////////////////////////

    private static DateTimeFormatter formatterYMD;
    private static DateTimeFormatter formatterHMS;
    private static DateTimeFormatter formatterFull;
}
