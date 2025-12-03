package com.bitwormhole.passwordbox.app.core.time;

public final class Time {

    private final long t;

    public Time(long unix) {
        this.t = unix;
    }

    public long value() {
        return this.t;
    }

    @Override
    public String toString() {
        return TimeFormatter.format(this);
    }

    public static Time now() {
        return new Time(System.currentTimeMillis());
    }
}
