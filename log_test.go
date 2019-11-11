package glog

import (
	"bytes"
	"log"
	"os"
	"testing"
	"time"
)

func TestOutput(t *testing.T) {
	const testString = "test"
	var b bytes.Buffer
	l := New(&b, "")
	l.Infof(testString)
	now := time.Now().Format(timeFormat)
	expect := "\x1b[0;0;32m" + now + " -> " + "INFO" + " 001 " + testString + "\x1b[0m\n"
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

	l.Debuf(testString)
	if b.String() != "" {
		t.Errorf("log error level should match %q is %q", "", b.String())
	}

	b.Reset()

	l.Infof(testString)
	if b.String() != "" {
		t.Errorf("log error level should match %q is %q", "", b.String())
	}

	b.Reset()

	l.Warnf(testString)
	if b.String() != "" {
		t.Errorf("log error level should match %q is %q", "", b.String())
	}

	b.Reset()

	l.Errof(testString)
	now := time.Now().Format(timeFormat)
	expect := "\x1b[0;0;31m" + now + " [glog]" + " -> " + "ERRO" + " 002 " + testString + "\x1b[0m\n"
	if b.String() != expect {
		t.Errorf("log error level should match %q is %q", expect, b.String())
	}
}

func TestWarningLevel(t *testing.T) {
	const testString = "test"
	var b bytes.Buffer
	l := New(&b, "glog")
	l.SetLevel(LevelWarning)

	l.Debuf(testString)
	if b.String() != "" {
		t.Errorf("log warning level should match %q is %q", "", b.String())
	}

	b.Reset()

	l.Infof(testString)
	if b.String() != "" {
		t.Errorf("log warning level should match %q is %q", "", b.String())
	}

	b.Reset()

	l.Warnf(testString)
	now := time.Now().Format(timeFormat)
	expect := "\x1b[0;0;33m" + now + " [glog]" + " -> " + "WARN" + " 003 " + testString + "\x1b[0m\n"
	if b.String() != expect {
		t.Errorf("log warning level should match %q is %q\n", expect, b.String())
	}

	b.Reset()

	l.Errof(testString)
	now = time.Now().Format(timeFormat)
	expect = "\x1b[0;0;31m" + now + " [glog]" + " -> " + "ERRO" + " 004 " + testString + "\x1b[0m\n"
	if b.String() != expect {
		t.Errorf("log warning level should match %q is %q\n", expect, b.String())
	}
}

func TestInfoLevel(t *testing.T) {
	const testString = "test"
	var b bytes.Buffer
	l := New(&b, "glog")
	l.SetLevel(LevelInfo)

	l.Debuf(testString)
	if b.String() != "" {
		t.Errorf("log warning level should match %q is %q", "", b.String())
	}

	b.Reset()

	l.Infof(testString)
	now := time.Now().Format(timeFormat)
	expect := "\x1b[0;0;32m" + now + " [glog]" + " -> " + "INFO" + " 005 " + testString + "\x1b[0m\n"
	if b.String() != expect {
		t.Errorf("log info level should match %q is %q", expect, b.String())
	}

	b.Reset()

	l.Warnf(testString)
	now = time.Now().Format(timeFormat)
	expect = "\x1b[0;0;33m" + now + " [glog]" + " -> " + "WARN" + " 006 " + testString + "\x1b[0m\n"
	if b.String() != expect {
		t.Errorf("log info level should match %q is %q\n", expect, b.String())
	}

	b.Reset()

	l.Errof(testString)
	now = time.Now().Format(timeFormat)
	expect = "\x1b[0;0;31m" + now + " [glog]" + " -> " + "ERRO" + " 007 " + testString + "\x1b[0m\n"
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
	expect := "\x1b[0;0;34m" + now + " [glog] " + "TestDebugLevel" + " -> " + "DEBU" + " 008 " + testString + "\x1b[0m\n"
	if b.String() != expect {
		t.Errorf("log debug level should match %q is %q", expect, b.String())
	}

	b.Reset()

	Infof(testString)
	now = time.Now().Format(timeFormat)
	expect = "\x1b[0;0;32m" + now + " [glog] " + "TestDebugLevel" + " -> " + "INFO" + " 009 " + testString + "\x1b[0m\n"
	if b.String() != expect {
		t.Errorf("log debug level should match %q is %q", expect, b.String())
	}

	b.Reset()

	Warnf(testString)
	now = time.Now().Format(timeFormat)
	expect = "\x1b[0;0;33m" + now + " [glog] " + "TestDebugLevel" + " -> " + "WARN" + " 00a " + testString + "\x1b[0m\n"
	if b.String() != expect {
		t.Errorf("log debug level should match %q is %q\n", expect, b.String())
	}

	b.Reset()

	Errof(testString)
	now = time.Now().Format(timeFormat)
	expect = "\x1b[0;0;31m" + now + " [glog] " + "TestDebugLevel" + " -> " + "ERRO" + " 00b " + testString + "\x1b[0m\n"
	if b.String() != expect {
		t.Errorf("log debug level should match %q is %q\n", expect, b.String())
	}
}

func TestLoggerID(t *testing.T) {
	const testString = "test"

	ResetID()

	for i := 0; i < 450000; i++ {
		Infof(testString)
	}

	if ID() != 450000 {
		t.Errorf("log globalID should match %v is %v", 100000, ID())
	}
}

func BenchmarkStdLog(b *testing.B) {
	const testString = "test"
	var bf bytes.Buffer
	log.SetOutput(&bf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Println(testString)
	}
}

func BenchmarkGlogInfo(b *testing.B) {
	const testString = "test"
	var buf bytes.Buffer
	l := New(&buf, "[glog]")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Infof(testString)
	}
}

func BenchmarkGlogDebu(b *testing.B) {
	const testString = "test"
	var buf bytes.Buffer
	l := New(&buf, "")
	l.SetLevel(LevelDebug)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Infof(testString)
	}
}

func TestExample(t *testing.T) {
	SetLevel(LevelDebug)
	SetPrefix("glog")
	SetOutput(os.Stdout)
	Infof("test")
	Warnf("test")
	Errof("test")
	Debuf("test")
}
