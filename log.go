package glog

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"sync"
	"time"
)

var std = &Logger{out: os.Stderr, level: LevelInfo, depth: 3}

type Level int

const (
	LevelError Level = iota
	LevelWarning
	LevelInfo
	LevelDebug

	timeFormat = "2006-01-02 15:04:05.999 MST"
)

var (
	levelNames = []string{
		"ERRO",
		"WARN",
		"INFO",
		"DEBU",
	}
)

func (l Level) String() string {
	return levelNames[l]
}

type Logger struct {
	id     uint64
	mu     sync.Mutex
	depth  int
	level  Level
	prefix string
	out    io.Writer
	buf    []byte
}

func New(out io.Writer, prefix string) *Logger {
	return &Logger{out: out, prefix: prefix, level: LevelInfo, depth: 2}
}

func (l *Logger) formatHeader(buf *[]byte, t time.Time, fname string, level Level) {

	*buf = append(*buf, t.Format(timeFormat)...)
	*buf = append(*buf, ' ')

	if l.prefix != "" {
		*buf = append(*buf, '[')
		*buf = append(*buf, l.prefix...)
		*buf = append(*buf, ']')
		*buf = append(*buf, ' ')
	}

	if l.level == LevelDebug {
		*buf = append(*buf, fname...)
		*buf = append(*buf, ' ')
	}

	*buf = append(*buf, "->"...)
	*buf = append(*buf, ' ')

	*buf = append(*buf, level.String()...)
	*buf = append(*buf, ' ')

	*buf = append(*buf, fmt.Sprintf("%03x", l.id)...)
	*buf = append(*buf, ' ')
}

func (l *Logger) Output(calldepth int, level Level, s string) error {

	now := time.Now()

	fname := ""
	if l.level == LevelDebug {
		pc, _, _, _ := runtime.Caller(calldepth)
		fname = path.Base(runtime.FuncForPC(pc).Name())
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	l.id++
	l.buf = l.buf[:0]
	l.formatHeader(&l.buf, now, fname, level)

	l.buf = append(l.buf, s...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}

	_, err := l.out.Write(l.buf)
	return err
}

func Infoln(v ...interface{}) {
	std.Infoln(v...)
}

func (l *Logger) Infoln(v ...interface{}) {
	if l.level >= LevelInfo {
		_ = l.Output(l.depth, LevelInfo, fmt.Sprintln(v...))
	}
}

func Infof(format string, v ...interface{}) {
	std.Infof(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if l.level >= LevelInfo {
		_ = l.Output(l.depth, LevelInfo, fmt.Sprintf(format, v...))
	}
}

func Warnln(v ...interface{}) {
	std.Warnln(v...)
}

func (l *Logger) Warnln(v ...interface{}) {
	if l.level >= LevelWarning {
		_ = l.Output(l.depth, LevelWarning, fmt.Sprintln(v...))
	}
}

func Warnf(format string, v ...interface{}) {
	std.Warnf(format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.level >= LevelWarning {
		_ = l.Output(l.depth, LevelWarning, fmt.Sprintf(format, v...))
	}
}

func Erroln(v ...interface{}) {
	std.Erroln(v...)
}

func (l *Logger) Erroln(v ...interface{}) {
	if l.level >= LevelError {
		_ = l.Output(l.depth, LevelError, fmt.Sprintln(v...))
	}
}

func Errof(format string, v ...interface{}) {
	std.Errof(format, v...)
}

func (l *Logger) Errof(format string, v ...interface{}) {
	if l.level >= LevelError {
		_ = l.Output(l.depth, LevelError, fmt.Sprintf(format, v...))
	}
}

func Debuln(v ...interface{}) {
	std.Debuln(v...)
}

func (l *Logger) Debuln(v ...interface{}) {
	if l.level >= LevelDebug {
		_ = l.Output(l.depth, LevelDebug, fmt.Sprintln(v...))
	}
}

func Debuf(format string, v ...interface{}) {
	std.Debuf(format, v...)
}

func (l *Logger) Debuf(format string, v ...interface{}) {
	if l.level >= LevelDebug {
		_ = l.Output(l.depth, LevelDebug, fmt.Sprintf(format, v...))
	}
}

func ID() uint64 {
	return std.ID()
}

func (l *Logger) ID() uint64 {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.id
}

func ResetID() {
	std.ResetID()
}

func (l *Logger) ResetID() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.id = 0
}

func (l *Logger) Writer() io.Writer {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.out
}

func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

func (l *Logger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.out = w
}

func Prefix() string {
	return std.Prefix()
}

func (l *Logger) Prefix() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.prefix
}

func SetPrefix(prefix string) {
	std.SetPrefix(prefix)
}

func (l *Logger) SetPrefix(prefix string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.prefix = prefix
}

func GetLevel() string {
	return std.Level()
}

func (l *Logger) Level() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.level.String()
}

func SetLevel(level Level) {
	std.SetLevel(level)
}

func (l *Logger) SetLevel(level Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}
