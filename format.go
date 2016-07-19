package logger

import (
	"encoding/json"
	"fmt"
	"time"
)

var (
	colorIndex = 0
	white      = "\033[37m"
	reset      = "\033[0m"
	bold       = "\033[1m"
	red        = "\033[31m"
	cyan       = "\033[36m"
	colors     = []string{
		"\033[34m", // blue
		"\033[32m", // green
		"\033[36m", // cyan
		"\033[33m", // yellow
		"\033[35m", // magenta
	}
)

func nextColor() string {
	colorIndex = colorIndex + 1
	return colors[colorIndex%len(colors)]
}

// Format returns either a JSON string or a pretty line for printing to terminal,
// depending on whether logger believes it's printing to stderr or to another writer.
func (l *Logger) Format(lvl logLevel, msg string, attrs *Attrs) string {
	if !ColorEnabled {
		return l.JSONFormat(lvl, msg, attrs)
	}

	return l.PrettyFormat(l.PrettyPrefix(lvl), msg, attrs)
}

type marshalStruct struct {
	Time       string `json:"time"`
	Package    string `json:"package"`
	Level      string `json:"level"`
	Msg        string `json:"msg"`
	Attributes *Attrs `json:"attributes,omitempty"`
}

// JSONFormat returns a JSON string representing the data provided and some internal state
// from the logger.
func (l *Logger) JSONFormat(lvl logLevel, msg string, attrs *Attrs) string {
	if FastJSON {
		var result string
		if attrs != nil {
			for key, val := range *attrs {
				if val, ok := val.(int); ok {
					result = fmt.Sprintf(`%s "%s": %d,`, result, key, val)
					continue
				}
				result = fmt.Sprintf(`%s "%s":"%v",`, result, key, val)
			}
			if len(*attrs) > 0 {
				result = fmt.Sprintf(`"attributes": { %s }, `, result)
			}
		}
		return fmt.Sprintf(`{ %s "time":"%s", "package":"%s", "level":"%s", "msg":"%s" }`,
			result, time.Now().Format(TimeFormat), l.Name, lvl, msg)
	}

	j, _ := json.Marshal(marshalStruct{
		time.Now().Format(TimeFormat),
		l.Name,
		lvl.String(),
		msg,
		attrs,
	})
	return string(j)
}

// PrettyFormat constructs a timestamped, named, levelled log line for a given message/attrs.
func (l *Logger) PrettyFormat(prefix, msg string, attrs *Attrs) string {
	// return fmt.Sprintf("%s %s%s%s:%s %s%s", time.Now().Format("15:04:05.000"), l.Color, l.Name, prefix, reset, msg, l.PrettyAttrs(attrs))
	return ""
}

// PrettyAttrs formats structured data provided as Attrs for printing to terminal.
func (l *Logger) PrettyAttrs(attrs *Attrs) string {
	result := ""
	empty := true

	if attrs == nil {
		return ""
	}

	for key, val := range *attrs {
		if empty == true {
			empty = false
		}

		result = fmt.Sprintf("%s %s=%v", result, key, val)
	}

	if empty == true {
		return ""
	}

	return fmt.Sprintf("%s %s", result, reset)
}

// PrettyPrefix provides a red "(!)" for Error logs only.
func (l *Logger) PrettyPrefix(lvl logLevel) string {
	// was one of the below (X vs. verbosity) gates supposed to be referring to
	// global verbosity, or local argument verbosity?
	if lvl != ErrorLevel {
		return ""
	}
	return ""
	// return fmt.Sprintf("(%s%s)", red+"!", l.Color)
}
