package glog

import (
	"bytes"
	"os"
	"testing"
	"time"
)

func TestOutput(t *testing.T) {
	const testString = "test"
	var b bytes.Buffer
	l := New(&b, "")
	l.Infoln(testString)
	now := time.Now().Format(timeFormat)
	expect := now + " -> " + "INFO" + " 001 " + testString + "\n"
	if b.String() != expect {
		t.Errorf("log output should match %q is %q", expect, b.String())
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
}

func TestErrorLevel(t *testing.T) {
	const testString = "test"
	var b bytes.Buffer
	l := New(&b, "glog")
	l.SetLevel(LevelError)

	l.Debuln(testString)
	if b.String() != "" {
		t.Errorf("log error level should match %q is %q", "", b.String())
	}

	b.Reset()

	l.Infoln(testString)
	if b.String() != "" {
		t.Errorf("log error level should match %q is %q", "", b.String())
	}

	b.Reset()

	l.Warnln(testString)
	if b.String() != "" {
		t.Errorf("log error level should match %q is %q", "", b.String())
	}

	b.Reset()

	l.Errof(testString)
	now := time.Now().Format(timeFormat)
	expect := now + " [glog]" + " -> " + "ERRO" + " 001 " + testString + "\n"
	if b.String() != expect {
		t.Errorf("log error level should match %q is %q", expect, b.String())
	}
}

func TestWarningLevel(t *testing.T) {
	const testString = "test"
	var b bytes.Buffer
	l := New(&b, "glog")
	l.SetLevel(LevelWarning)

	l.Debuln(testString)
	if b.String() != "" {
		t.Errorf("log warning level should match %q is %q", "", b.String())
	}

	b.Reset()

	l.Infoln(testString)
	if b.String() != "" {
		t.Errorf("log warning level should match %q is %q", "", b.String())
	}

	b.Reset()

	l.Warnf(testString)
	now := time.Now().Format(timeFormat)
	expect := now + " [glog]" + " -> " + "WARN" + " 001 " + testString + "\n"
	if b.String() != expect {
		t.Errorf("log warning level should match %q is %q\n", expect, b.String())
	}

	b.Reset()

	l.Erroln(testString)
	now = time.Now().Format(timeFormat)
	expect = now + " [glog]" + " -> " + "ERRO" + " 002 " + testString + "\n"
	if b.String() != expect {
		t.Errorf("log warning level should match %q is %q\n", expect, b.String())
	}
}

func TestInfoLevel(t *testing.T) {
	const testString = "test"
	var b bytes.Buffer
	l := New(&b, "glog")
	l.SetLevel(LevelInfo)

	l.Debuln(testString)
	if b.String() != "" {
		t.Errorf("log warning level should match %q is %q", "", b.String())
	}

	b.Reset()

	l.Infof(testString)
	now := time.Now().Format(timeFormat)
	expect := now + " [glog]" + " -> " + "INFO" + " 001 " + testString + "\n"
	if b.String() != expect {
		t.Errorf("log info level should match %q is %q", expect, b.String())
	}

	b.Reset()

	l.Warnln(testString)
	now = time.Now().Format(timeFormat)
	expect = now + " [glog]" + " -> " + "WARN" + " 002 " + testString + "\n"
	if b.String() != expect {
		t.Errorf("log info level should match %q is %q\n", expect, b.String())
	}

	b.Reset()

	l.Erroln(testString)
	now = time.Now().Format(timeFormat)
	expect = now + " [glog]" + " -> " + "ERRO" + " 003 " + testString + "\n"
	if b.String() != expect {
		t.Errorf("log info level should match %q is %q\n", expect, b.String())
	}
}

func TestDebugLevel(t *testing.T) {
	const testString = "test"
	var b bytes.Buffer
	SetOutput(&b)
	SetPrefix("glog")
	SetLevel(LevelDebug)

	Debuf(testString)
	now := time.Now().Format(timeFormat)
	expect := now + " [glog] " + "glog.TestDebugLevel" + " -> " + "DEBU" + " 001 " + testString + "\n"
	if b.String() != expect {
		t.Errorf("log debug level should match %q is %q", expect, b.String())
	}

	b.Reset()

	Debuln(testString)
	now = time.Now().Format(timeFormat)
	expect = now + " [glog] " + "glog.TestDebugLevel" + " -> " + "DEBU" + " 002 " + testString + "\n"
	if b.String() != expect {
		t.Errorf("log debug level should match %q is %q", expect, b.String())
	}

	b.Reset()

	Infof(testString)
	now = time.Now().Format(timeFormat)
	expect = now + " [glog] " + "glog.TestDebugLevel" + " -> " + "INFO" + " 003 " + testString + "\n"
	if b.String() != expect {
		t.Errorf("log debug level should match %q is %q", expect, b.String())
	}

	b.Reset()

	Infoln(testString)
	now = time.Now().Format(timeFormat)
	expect = now + " [glog] " + "glog.TestDebugLevel" + " -> " + "INFO" + " 004 " + testString + "\n"
	if b.String() != expect {
		t.Errorf("log debug level should match %q is %q", expect, b.String())
	}

	b.Reset()

	Warnf(testString)
	now = time.Now().Format(timeFormat)
	expect = now + " [glog] " + "glog.TestDebugLevel" + " -> " + "WARN" + " 005 " + testString + "\n"
	if b.String() != expect {
		t.Errorf("log debug level should match %q is %q\n", expect, b.String())
	}

	b.Reset()

	Warnln(testString)
	now = time.Now().Format(timeFormat)
	expect = now + " [glog] " + "glog.TestDebugLevel" + " -> " + "WARN" + " 006 " + testString + "\n"
	if b.String() != expect {
		t.Errorf("log debug level should match %q is %q\n", expect, b.String())
	}

	b.Reset()

	Errof(testString)
	now = time.Now().Format(timeFormat)
	expect = now + " [glog] " + "glog.TestDebugLevel" + " -> " + "ERRO" + " 007 " + testString + "\n"
	if b.String() != expect {
		t.Errorf("log debug level should match %q is %q\n", expect, b.String())
	}

	b.Reset()

	Erroln(testString)
	now = time.Now().Format(timeFormat)
	expect = now + " [glog] " + "glog.TestDebugLevel" + " -> " + "ERRO" + " 008 " + testString + "\n"
	if b.String() != expect {
		t.Errorf("log debug level should match %q is %q\n", expect, b.String())
	}
}

func TestLoggerID(t *testing.T) {
	const testString = "test"

	ResetID()

	for i := 0; i < 100000; i++ {
		Infoln(testString)
	}

	if ID() != 100000 {
		t.Errorf("log id should match %v is %v", 100000, ID())
	}
}

func BenchmarkInfo(b *testing.B) {
	const testString = "test"
	var buf bytes.Buffer
	l := New(&buf, "")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Infoln(testString)
	}
}

func BenchmarkDebu(b *testing.B) {
	const testString = "test"
	var buf bytes.Buffer
	l := New(&buf, "")
	l.SetLevel(LevelDebug)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Infoln(testString)
	}
}

func TestExample(t *testing.T) {
	SetLevel(LevelDebug)
	SetPrefix("glog")
	SetOutput(os.Stdout)
	Infoln("test")
	Warnln("test")
	Erroln("test")
	Debuln("test")
}
