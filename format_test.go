package logger

import (
	"io/ioutil"
	"testing"
)

var (
	l     = New("A")
	msg   = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	attrs = Attrs{
		"A": 1, "B": 2, "C": 3, "D": 4, "E": 5,
		"F": 6, "G": 7, "H": 8, "I": 9, "J": 0,
		"K": msg, "L": msg, "M": msg,
		"A1": 1, "B1": 2, "C1": 3, "D1": 4, "E1": 5,
		"F1": 6, "G1": 7, "H1": 8, "I1": 9, "J1": 0,
		"K1": msg, "L1": msg, "M1": msg,
		"A2": 1, "B2": 2, "C2": 3, "D2": 4, "E2": 5,
		"F2": 6, "G2": 7, "H2": 8, "I2": 9, "J2": 0,
		"K2": msg, "L2": msg, "M2": msg,
	}
)

func init() {
	SetOutput(ioutil.Discard)
}

func BenchmarkFastJson(b *testing.B) {
	FastJSON = true
	for i := 0; i < b.N; i++ {
		l.Info(msg, msg, msg, msg, attrs)
	}
}

func BenchmarkSafeJson(b *testing.B) {
	FastJSON = false
	for i := 0; i < b.N; i++ {
		l.Info(msg, msg, msg, msg, attrs)
	}
}
