package kisslog

import (
	"fmt"
)

type logLevel int

const (
	DebugLevel logLevel = iota
	InfoLevel
	ErrorLevel
	DisableLevel
)

func (l logLevel) String() string {
	switch l {
	case DebugLevel:
		return "[DEBUG]"
	case InfoLevel:
		return "[INFO ]"
	case ErrorLevel:
		return "[ERROR]"
	}
	return ""
}

type Logger struct {
	Name        string
	LoggerLevel logLevel
}

func New(name string) *Logger {
	return &Logger{
		Name: name,
	}
}

func (l *Logger) Debug(format string, v ...interface{}) {
	if LogLevel > DebugLevel ||
		l.LoggerLevel > DebugLevel ||
		(!allEnabled || enabledLoggers[l.Name]) {
		return
	}

	v, attrs := splitAttrs(v...)
	l.output(DebugLevel, fmt.Sprintf(format, v...), attrs)
}

func (l *Logger) Info(format string, v ...interface{}) {
	if LogLevel > InfoLevel ||
		l.LoggerLevel > InfoLevel ||
		(!allEnabled || enabledLoggers[l.Name]) {
		return
	}

	v, attrs := splitAttrs(v...)
	l.output(InfoLevel, fmt.Sprintf(format, v...), attrs)
}

func (l *Logger) Error(format string, v ...interface{}) {
	if LogLevel > ErrorLevel ||
		l.LoggerLevel > ErrorLevel ||
		(!allEnabled && enabledLoggers[l.Name]) {
		return
	}

	v, attrs := splitAttrs(v...)
	l.output(ErrorLevel, fmt.Sprintf(format, v...), attrs)
}

func (l *Logger) output(lvl logLevel, msg string, attrs *Attrs) {
	l.write(l.format(lvl, msg, attrs))
}

func (l *Logger) write(log string) {
	fmt.Fprintln(out, log)
}
