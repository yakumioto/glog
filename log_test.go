package glog

import (
	"bytes"
	"fmt"
	"os"
	"testing"
	"time"
)

func color(level Level) string {
	switch level {
	case levelPanic, levelFatal, LevelError:
		return "31"
	case LevelWarning:
		return "33"
	case LevelInfo:
		return "32"
	case LevelDebug:
		return "34"
	default:
		return ""
	}
}

func expectlog(level Level, prefix, fnc string, id int, msg string) string {
	headFormat := "\x1b[0;0;%sm%s"
	prefixFormat := " [%s]"
	fncFormat := " %s"
	endFormat := " -> %s %03x %s\n\x1b[0m"
	now := time.Now().Format(timeFormat)

	format := headFormat
	if prefix != "" {
		format += prefixFormat
	} else {
		format += "%s"
	}
	if fnc != "" {
		format += fncFormat
	} else {
		format += "%s"
	}
	format += endFormat

	return fmt.Sprintf(format, color(level), now, prefix, fnc, level.String(), id, msg)
}

func mastEqual(t *testing.T, expect, actual string) {
	if expect == "" && actual == "" {
		return
	}

	equalTime := expect[:29] == actual[:29]
	equalLog := expect[34:] == actual[34:]

	if !equalTime || !equalLog {
		t.Errorf("log output should match\nactual: %s\nexpect: %s", actual, expect)
	}
}

func TestPrefixAndLevelSetting(t *testing.T) {
	SetPrefix("glog")
	if Prefix() != "glog" {
		t.Errorf("log prefix should match %q is %q", "glog", Prefix())
	}

	SetLevel(LevelDebug)
	if GetLevel() != LevelDebug.String() {
		t.Errorf("log level should match %q is %q", LevelDebug.String(), GetLevel())
	}

	write := Writer()
	if write != os.Stderr {
		t.Error("log writer should match os.stderr")
	}
}

func TestOutput(t *testing.T) {
	const testString = "test"
	var actual bytes.Buffer
	SetOutput(&actual)
	Infof(testString)

	mastEqual(t, expectlog(LevelInfo, "glog", "TestOutput", 1, testString), actual.String())

	ResetID()
}

func TestMustGetLogger(t *testing.T) {
	SetPrefix("")
	SetLevel(LevelInfo)
	l := MustGetLogger("1")

	const testString = "log message"
	var actual bytes.Buffer
	l.SetOutput(&actual)

	l.Infoln(testString)
	mastEqual(t, expectlog(LevelInfo, "1", "", 1, testString), actual.String())

	actual.Reset()
	l = l.MustGetLogger("2")
	l.Infoln(testString)
	mastEqual(t, expectlog(LevelInfo, "1.2", "", 2, testString), actual.String())

	l.ResetID()
}

func TestErrorLevel(t *testing.T) {
	const testString = "log message"
	var actual bytes.Buffer
	l := New(&actual, "glog")
	l.SetLevel(LevelError)

	l.Debuf(testString)
	mastEqual(t, "", actual.String())

	l.Debuln(testString)
	mastEqual(t, "", actual.String())

	l.Infof(testString)
	mastEqual(t, "", actual.String())

	l.Infoln(testString)
	mastEqual(t, "", actual.String())

	l.Warnf(testString)
	mastEqual(t, "", actual.String())

	l.Warnln(testString)
	mastEqual(t, "", actual.String())

	l.Errof(testString)
	mastEqual(t, expectlog(LevelError, "glog", "", 1, testString), actual.String())

	actual.Reset()
	l.Erroln(testString)
	mastEqual(t, expectlog(LevelError, "glog", "", 2, testString), actual.String())

	l.ResetID()
}

func TestWarningLevel(t *testing.T) {
	const testString = "log message"
	var actual bytes.Buffer
	l := New(&actual, "glog")
	l.SetLevel(LevelWarning)

	l.Debuf(testString)
	mastEqual(t, "", actual.String())

	l.Debuln(testString)
	mastEqual(t, "", actual.String())

	l.Infof(testString)
	mastEqual(t, "", actual.String())

	l.Infoln(testString)
	mastEqual(t, "", actual.String())

	l.Warnf(testString)
	mastEqual(t, expectlog(LevelWarning, "glog", "", 1, testString), actual.String())

	actual.Reset()
	l.Warnln(testString)
	mastEqual(t, expectlog(LevelWarning, "glog", "", 2, testString), actual.String())

	actual.Reset()
	l.Errof(testString)
	mastEqual(t, expectlog(LevelError, "glog", "", 3, testString), actual.String())

	actual.Reset()
	l.Erroln(testString)
	mastEqual(t, expectlog(LevelError, "glog", "", 4, testString), actual.String())

	l.ResetID()
}

func TestInfoLevel(t *testing.T) {
	const testString = "log message"
	var actual bytes.Buffer
	l := New(&actual, "glog")
	l.SetLevel(LevelInfo)

	l.Debuf(testString)
	mastEqual(t, "", actual.String())

	l.Debuln(testString)
	mastEqual(t, "", actual.String())

	l.Infof(testString)
	mastEqual(t, expectlog(LevelInfo, "glog", "", 1, testString), actual.String())

	actual.Reset()
	l.Infoln(testString)
	mastEqual(t, expectlog(LevelInfo, "glog", "", 2, testString), actual.String())

	actual.Reset()
	l.Warnf(testString)
	mastEqual(t, expectlog(LevelWarning, "glog", "", 3, testString), actual.String())

	actual.Reset()
	l.Warnln(testString)
	mastEqual(t, expectlog(LevelWarning, "glog", "", 4, testString), actual.String())

	actual.Reset()
	l.Errof(testString)
	mastEqual(t, expectlog(LevelError, "glog", "", 5, testString), actual.String())

	actual.Reset()
	l.Erroln(testString)
	mastEqual(t, expectlog(LevelError, "glog", "", 6, testString), actual.String())

	l.ResetID()
}

func TestDebugLevel(t *testing.T) {
	const testString = "log message"
	var actual bytes.Buffer
	l := New(&actual, "glog")
	l.SetLevel(LevelDebug)

	l.Debuf(testString)
	mastEqual(t, expectlog(LevelDebug, "glog", "TestDebugLevel", 1, testString), actual.String())

	actual.Reset()
	l.Debuln(testString)
	mastEqual(t, expectlog(LevelDebug, "glog", "TestDebugLevel", 2, testString), actual.String())

	actual.Reset()
	l.Infof(testString)
	mastEqual(t, expectlog(LevelInfo, "glog", "TestDebugLevel", 3, testString), actual.String())

	actual.Reset()
	l.Infoln(testString)
	mastEqual(t, expectlog(LevelInfo, "glog", "TestDebugLevel", 4, testString), actual.String())

	actual.Reset()
	l.Warnf(testString)
	mastEqual(t, expectlog(LevelWarning, "glog", "TestDebugLevel", 5, testString), actual.String())

	actual.Reset()
	l.Warnln(testString)
	mastEqual(t, expectlog(LevelWarning, "glog", "TestDebugLevel", 6, testString), actual.String())

	actual.Reset()
	l.Errof(testString)
	mastEqual(t, expectlog(LevelError, "glog", "TestDebugLevel", 7, testString), actual.String())

	actual.Reset()
	l.Erroln(testString)
	mastEqual(t, expectlog(LevelError, "glog", "TestDebugLevel", 8, testString), actual.String())

	l.ResetID()
}

func TestLoggerID(t *testing.T) {
	np, err := os.Open(os.DevNull)
	if err != nil {
		t.Fatalf("open %s error: %s", os.DevNull, err)
	}

	const testString = "log message"
	l := New(np, "glog")
	l.SetLevel(LevelDebug)

	for i := 0; i < 1234; i++ {
		l.Infof(testString)
	}

	if ID() != 1234 {
		t.Errorf("log id should match %v is %v", 1234, ID())
	}

	l.ResetID()
}

func TestStdLog(t *testing.T) {
	const testString = "log message"
	var actual bytes.Buffer
	SetOutput(&actual)
	SetLevel(LevelDebug)

	Infof(testString)
	mastEqual(t, expectlog(LevelInfo, "", "TestStdLog", 1, testString), actual.String())
	actual.Reset()
	Infoln(testString)
	mastEqual(t, expectlog(LevelInfo, "", "TestStdLog", 2, testString), actual.String())
	actual.Reset()

	Warnf(testString)
	mastEqual(t, expectlog(LevelWarning, "", "TestStdLog", 3, testString), actual.String())
	actual.Reset()
	Warnln(testString)
	mastEqual(t, expectlog(LevelWarning, "", "TestStdLog", 4, testString), actual.String())
	actual.Reset()

	Errof(testString)
	mastEqual(t, expectlog(LevelError, "", "TestStdLog", 5, testString), actual.String())
	actual.Reset()
	Erroln(testString)
	mastEqual(t, expectlog(LevelError, "", "TestStdLog", 6, testString), actual.String())
	actual.Reset()

	Debuf(testString)
	mastEqual(t, expectlog(LevelDebug, "", "TestStdLog", 7, testString), actual.String())
	actual.Reset()
	Debuln(testString)
	mastEqual(t, expectlog(LevelDebug, "", "TestStdLog", 8, testString), actual.String())
	actual.Reset()

	ResetID()
}
