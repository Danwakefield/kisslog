package kisslog

import (
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"github.com/azer/is-terminal"
)

var (
	out            io.Writer
	allEnabled     = true
	enabledLoggers = map[string]bool{}
	enabledMutex   sync.RWMutex

	JSONOutput = false
	LogLevel   = InfoLevel
	TimeFormat = "15:04:05"
	TraceFile  = true
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

	v, _ := strconv.ParseBool(os.Getenv("LOG_TRACE"))
	TraceFile = v

	v, err := strconv.ParseBool(os.Getenv("LOG_JSON"))
	if err != nil {
		JSONOutput = !isterminal.IsTerminal(syscall.Stderr)
	} else {
		JSONOutput = v
	}

	if f, exists := os.LookupEnv("LOG_TIMEFORMAT"); exists {
		TimeFormat = f
	}

	if list, exists := os.LookupEnv("LOG_ENABLED"); exists {
		enabledMutex.Lock()
		allEnabled = false
		for _, name := range strings.Split(list, ",") {
			enabledLoggers[name] = true
		}
		enabledMutex.Unlock()
	}
}

func EnableLogger(name string) {
	enabledMutex.Lock()

	allEnabled = false
	enabledLoggers[name] = true

	enabledMutex.Unlock()
}

func DisableLogger(name string) {
	enabledMutex.Lock()
	delete(enabledLoggers, name)
	enabledMutex.Unlock()
}

func SetOutput(w io.Writer) {
	out = w
}
