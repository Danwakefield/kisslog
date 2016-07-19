package logger

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

var LogLevel = InfoLevel

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
	// Name by which the logger is identified when enabling or disabling it, and by envvar.
	Name string
}

func New(name string) *Logger {
	return &Logger{
		Name: name,
	}
}

func (l *Logger) IsEnabled() bool {
	return true
}

func (l *Logger) Debug(format string, v ...interface{}) {
	if LogLevel > DebugLevel || !l.IsEnabled() {
		return
	}

	v, attrs := SplitAttrs(v...)

	l.Output(DebugLevel, fmt.Sprintf(format, v...), attrs)

}

func (l *Logger) Info(format string, v ...interface{}) {
	if LogLevel > InfoLevel || !l.IsEnabled() {
		return
	}

	v, attrs := SplitAttrs(v...)

	l.Output(InfoLevel, fmt.Sprintf(format, v...), attrs)
}

// Error logs an error message using fmt. It has log-level 3, the highest level.
func (l *Logger) Error(format string, v ...interface{}) {
	if !l.IsEnabled() {
		return
	}

	v, attrs := SplitAttrs(v...)

	l.Output(ErrorLevel, fmt.Sprintf(format, v...), attrs)
}

// Output is the lower-level call delegated to by Info/Timer/Error, and can be used
// to directly write to the underlying buffer regardless of log-level.
func (l *Logger) Output(lvl logLevel, msg string, attrs *Attrs) {
	l.Write(l.Format(lvl, msg, attrs))
}

func (l *Logger) Write(log string) {
	fmt.Fprintln(out, log)
}
