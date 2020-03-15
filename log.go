package glog

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Level int

const (
	levelPanic Level = iota
	levelFatal
	LevelError
	LevelWarning
	LevelInfo
	LevelDebug

	timeFormat = "2006-01-02 15:04:05.000 MST"
)

var (
	id        uint64 = 0
	levelName        = [...]string{
		"PANI",
		"FATA",
		"ERRO",
		"WARN",
		"INFO",
		"DEBU",
	}
)

var std = &Logger{id: &id, out: os.Stderr, level: LevelInfo, depth: 3}

func (l Level) String() string {
	return levelName[l]
}

type Logger struct {
	id     *uint64
	mu     sync.Mutex
	depth  int
	level  Level
	prefix string
	out    io.Writer
	buf    []byte
}

func New(out io.Writer, prefix string) *Logger {
	return &Logger{id: &id, out: out, prefix: prefix, level: LevelInfo, depth: 2}
}

func (l *Logger) formatHeader(buf *[]byte, t time.Time, fname string, level Level) {
	switch level {
	case levelPanic, levelFatal, LevelError:
		*buf = append(*buf, "\x1b[0;0;31m"...)
	case LevelWarning:
		*buf = append(*buf, "\x1b[0;0;33m"...)
	case LevelInfo:
		*buf = append(*buf, "\x1b[0;0;32m"...)
	case LevelDebug:
		*buf = append(*buf, "\x1b[0;0;34m"...)
	}

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

	*buf = append(*buf, fmt.Sprintf("%03x", *l.id)...)
	*buf = append(*buf, ' ')
}

func (l *Logger) Output(calldepth int, level Level, s string) error {
	now := time.Now()

	fname := ""
	if l.level == LevelDebug {
		pc, _, _, _ := runtime.Caller(calldepth)
		fnames := strings.Split(path.Base(runtime.FuncForPC(pc).Name()), ".")
		fname = fnames[len(fnames)-1]
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	*l.id++
	l.buf = l.buf[:0]
	l.formatHeader(&l.buf, now, fname, level)

	l.buf = append(l.buf, s...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}
	l.buf = append(l.buf, "\x1b[0m"...)

	_, err := l.out.Write(l.buf)
	return err
}

func Panicln(v ...interface{})               { std.Panicln(v...) }
func Fatalln(v ...interface{})               { std.Fatalln(v...) }
func Erroln(v ...interface{})                { std.Erroln(v...) }
func Warnln(v ...interface{})                { std.Warnln(v...) }
func Infoln(v ...interface{})                { std.Infoln(v...) }
func Debuln(v ...interface{})                { std.Debuln(v...) }
func Panicf(format string, v ...interface{}) { std.Panicf(format, v...) }
func Fatalf(format string, v ...interface{}) { std.Fatalf(format, v...) }
func Errof(format string, v ...interface{})  { std.Errof(format, v...) }
func Warnf(format string, v ...interface{})  { std.Warnf(format, v...) }
func Infof(format string, v ...interface{})  { std.Infof(format, v...) }
func Debuf(format string, v ...interface{})  { std.Debuf(format, v...) }

func (l *Logger) Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	l.Output(l.depth, levelPanic, s)
	panic(s)
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.Output(l.depth, levelPanic, s)
	panic(s)
}

func (l *Logger) Fatalln(v ...interface{}) {
	l.Output(l.depth, levelFatal, fmt.Sprintln(v...))
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Output(l.depth, levelFatal, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l *Logger) Erroln(v ...interface{}) {
	if l.level >= LevelError {
		l.Output(l.depth, LevelError, fmt.Sprintln(v...))
	}
}

func (l *Logger) Errof(format string, v ...interface{}) {
	if l.level >= LevelError {
		l.Output(l.depth, LevelError, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Warnln(v ...interface{}) {
	if l.level >= LevelWarning {
		l.Output(l.depth, LevelWarning, fmt.Sprintln(v...))
	}
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.level >= LevelWarning {
		l.Output(l.depth, LevelWarning, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Infoln(v ...interface{}) {
	if l.level >= LevelInfo {
		l.Output(l.depth, LevelInfo, fmt.Sprintln(v...))
	}
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if l.level >= LevelInfo {
		l.Output(l.depth, LevelInfo, fmt.Sprintf(format, v...))
	}
}

func (l *Logger) Debuln(v ...interface{}) {
	if l.level >= LevelDebug {
		l.Output(l.depth, LevelDebug, fmt.Sprintln(v...))
	}
}

func (l *Logger) Debuf(format string, v ...interface{}) {
	if l.level >= LevelDebug {
		l.Output(l.depth, LevelDebug, fmt.Sprintf(format, v...))
	}
}

func MustGetLogger(prefix string) *Logger { return std.MustGetLogger(prefix) }
func ID() uint64                          { return std.ID() }
func ResetID()                            { std.ResetID() }
func SetOutput(w io.Writer)               { std.SetOutput(w) }
func Prefix() string                      { return std.Prefix() }
func SetPrefix(prefix string)             { std.SetPrefix(prefix) }
func Writer() io.Writer                   { return std.Writer() }
func GetLevel() string                    { return std.Level() }
func SetLevel(level Level)                { std.SetLevel(level) }

func (l *Logger) MustGetLogger(prefix string) *Logger {
	if l.prefix != "" {
		prefix = strings.Join([]string{l.prefix, prefix}, ".")
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	return &Logger{
		id:     &id,
		depth:  l.depth,
		level:  l.level,
		prefix: prefix,
		out:    l.out,
	}
}

func (l *Logger) ID() uint64 {
	l.mu.Lock()
	defer l.mu.Unlock()

	return *l.id
}

func (l *Logger) ResetID() {
	l.mu.Lock()
	defer l.mu.Unlock()

	*l.id = 0
}

func (l *Logger) Writer() io.Writer {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.out
}

func (l *Logger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.out = w
}

func (l *Logger) Prefix() string {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.prefix
}

func (l *Logger) SetPrefix(prefix string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.prefix = prefix
}

func (l *Logger) Level() string {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.level.String()
}

func (l *Logger) SetLevel(level Level) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.level = level
}
