package kisslog

import (
	"bytes"
	"encoding/json"
	"path"
	"runtime"
	"strconv"
	"time"
)

func (l *Logger) format(lvl logLevel, msg string, attrs *Attrs) string {
	if JSONOutput {
		return l.formatJSON(lvl, msg, attrs)
	}

	return l.formatPretty(lvl, msg, attrs)
}

type marshalStruct struct {
	Time       string `json:"time"`
	Package    string `json:"package"`
	Level      string `json:"level"`
	Trace      string `json:"trace,omitempty"`
	Msg        string `json:"msg"`
	Attributes *Attrs `json:"attributes,omitempty"`
}

func (l *Logger) formatJSON(lvl logLevel, msg string, attrs *Attrs) string {
	j, _ := json.Marshal(marshalStruct{
		time.Now().Format(TimeFormat),
		l.Name,
		lvl.String(),
		trace(),
		msg,
		attrs,
	})
	return string(j)
}

func (l *Logger) formatPretty(lvl logLevel, msg string, attrs *Attrs) string {
	b := bytes.Buffer{}
	if TimeFormat != "" {
		b.WriteRune('[')
		b.WriteString(time.Now().Format(TimeFormat))
		b.WriteRune(']')
	}
	b.WriteString(lvl.String())
	b.WriteString(trace())
	b.WriteString(l.Name)
	b.WriteString(": ")
	b.WriteString(msg)
	b.WriteString(attrs.Pretty())

	return b.String()
}

func trace() string {
	if !TraceFile {
		return ""
	}
	// We are 5 functions from where the original Info, Debug or Error was
	// called.
	pc, f, line, _ := runtime.Caller(5)
	name := runtime.FuncForPC(pc).Name()
	name = path.Base(name) // only use package.funcname
	f = path.Base(f)       // only use filename

	return "[" + name + ":" + f + ":" + strconv.Itoa(line) + "] "
}
