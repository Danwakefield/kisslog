package kisslog

import (
	"io"
	"os"
	"strings"
	"syscall"

	"github.com/azer/is-terminal"
)

var (
	out        io.Writer
	allEnabled = true

	enabledLoggers = map[string]bool{}
	JSONOutput     = false
	LogLevel       = InfoLevel
	TimeFormat     = "15:04:05"
	TraceFile      = true
)

func init() {
	out = os.Stderr
	parseEnvVars()
}

func parseEnvVars() {
	switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
	case "debug":
		LogLevel = DebugLevel
	case "info":
		LogLevel = InfoLevel
	case "error":
		LogLevel = ErrorLevel
	case "disable":
		LogLevel = DisableLevel
	default:
		LogLevel = InfoLevel
	}

	switch strings.ToLower(os.Getenv("LOG_TRACE")) {
	case "enable", "on", "true", "1":
		TraceFile = true
	case "disable", "off", "false", "0":
		TraceFile = false
	default:
		TraceFile = true
	}

	switch strings.ToLower(os.Getenv("LOG_JSON")) {
	case "enable", "on", "true", "1":
		JSONOutput = true
	case "disable", "off", "false", "0":
		JSONOutput = false
	default:
		JSONOutput = !isterminal.IsTerminal(syscall.Stderr)
	}

	if f, exists := os.LookupEnv("LOG_TIMEFORMAT"); exists {
		TimeFormat = f
	}

	if list, exists := os.LookupEnv("LOG_ENABLED"); exists {
		allEnabled = false
		for _, name := range strings.Split(list, ",") {
			enabledLoggers[name] = true
		}
	}
}

func EnableLogger(name string) {
	allEnabled = false
	enabledLoggers[name] = true
}

func DisableLogger(name string) {
	delete(enabledLoggers, name)
}

func SetOutput(w io.Writer) {
	out = w
}
