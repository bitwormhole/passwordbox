package loggers

import (
	"io"
	"log"
	"os"
)

type Logger = log.Logger

////////////////////////////////////////////////////////////////////////////////
// level

type Level int

const (
	MIN Level = 0

	TRACE Level = 1
	DEBUG Level = 2
	INFO  Level = 3
	WARN  Level = 4
	ERROR Level = 5
	FATAL Level = 6

	MAX Level = 7
)

func (l Level) String() string {
	switch l {
	case INFO:
		return "INFO."
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case WARN:
		return "WARN."
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "UNDEF"
}

////////////////////////////////////////////////////////////////////////////////
// holder

type innerLoggersHolder struct {
	trace *Logger
	debug *Logger
	info  *Logger
	warn  *Logger
	error *Logger
	fatal *Logger

	nop *Logger

	initialled bool
	levelLimit Level
}

func (inst *innerLoggersHolder) tryInit() *innerLoggersHolder {
	if !inst.initialled {
		inst.innerDoInit()
	}
	return inst
}

func (inst *innerLoggersHolder) innerDoInit() {

	err := os.Stderr
	out := os.Stdout
	nop := new(innerNopWriter)

	inst.trace = inst.innerMakeLogger(TRACE, out)
	inst.debug = inst.innerMakeLogger(DEBUG, out)
	inst.info = inst.innerMakeLogger(INFO, out)
	inst.warn = inst.innerMakeLogger(WARN, err)
	inst.error = inst.innerMakeLogger(ERROR, err)
	inst.fatal = inst.innerMakeLogger(FATAL, err)
	inst.nop = inst.innerMakeLogger(MIN, nop)

	inst.levelLimit = INFO
	inst.initialled = true
}

func (inst *innerLoggersHolder) innerMakeLogger(level Level, dst io.Writer) *Logger {

	prefix := "(" + level.String() + ") "
	l := new(Logger)
	flags := log.Ltime | log.Ldate | log.LUTC | log.Lmicroseconds

	l.SetOutput(dst)
	l.SetPrefix(prefix)
	l.SetFlags(flags)

	return l
}

func (inst *innerLoggersHolder) getLogger(level Level) *Logger {

	switch level {
	case INFO:
		return inst.info
	case DEBUG:
		return inst.debug
	case TRACE:
		return inst.trace
	case WARN:
		return inst.warn
	case ERROR:
		return inst.error
	case FATAL:
		return inst.fatal
	}
	return inst.nop
}

func (inst *innerLoggersHolder) isLevelEnabled(level Level) bool {
	return (inst.levelLimit <= level)
}

func (inst *innerLoggersHolder) setLevelEnabled(level Level) {
	inst.levelLimit = level
}

////////////////////////////////////////////////////////////////////////////////
// holder (inst)

var theLoggersHolder innerLoggersHolder

func innerGetLogger(level Level) *Logger {
	holder := innerGetLoggerHolder()
	return holder.getLogger(level)
}

func innerGetLoggerHolder() *innerLoggersHolder {
	holder := &theLoggersHolder
	return holder.tryInit()
}

////////////////////////////////////////////////////////////////////////////////

type innerNopWriter struct{}

func (i *innerNopWriter) Write(p []byte) (n int, err error) {
	// nop
	return len(p), nil
}

////////////////////////////////////////////////////////////////////////////////

func IsEnabled(level Level) bool {
	holder := innerGetLoggerHolder()
	return holder.isLevelEnabled(level)
}

func SetLevelEnabled(level Level) {
	holder := innerGetLoggerHolder()
	holder.setLevelEnabled(level)
}

////////////////////////////////////////////////////////////////////////////////

func LogT(fmt string, args ...any) {
	innerGetLogger(TRACE).Printf(fmt, args...)
}

func LogD(fmt string, args ...any) {
	innerGetLogger(DEBUG).Printf(fmt, args...)
}

func LogI(fmt string, args ...any) {
	innerGetLogger(INFO).Printf(fmt, args...)
}

func LogW(fmt string, args ...any) {
	innerGetLogger(WARN).Printf(fmt, args...)
}

func LogE(fmt string, args ...any) {
	innerGetLogger(ERROR).Printf(fmt, args...)
}

func LogF(fmt string, args ...any) {
	innerGetLogger(FATAL).Printf(fmt, args...)
}

func GetLogger(level Level) *Logger {
	return innerGetLogger(level)
}

////////////////////////////////////////////////////////////////////////////////
