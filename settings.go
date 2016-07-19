package logger

import (
	"io"
	"os"
)

var (
	out          io.Writer
	ColorEnabled = true
	TimeFormat   = "01-02 - 15:04:05"
	FastJSON     = false
)

func init() {
	out = os.Stderr
}

// SetOutput directs logs to a writer other than Stderr; this disables pretty pretty-printing
// and outputs JSON instead.
func SetOutput(w io.Writer) {
	out = w
}
