# Golog

Sample leveled logging in Go.

## Installation

`go get github.com/yakumioto/golog`

## Quick Start

```go
SetLevel(LevelDebug)
SetPrefix("golog")
SetOutput(os.Stdout)
Infoln("test")
Warnln("test")
Erroln("test")
Debuln("test")

// ---- output ----
//2019-09-03 14:29:34.759 CST [golog] golog.TestExample -> INFO 001 test
//2019-09-03 14:29:34.759 CST [golog] golog.TestExample -> WARN 002 test
//2019-09-03 14:29:34.759 CST [golog] golog.TestExample -> ERRO 003 test
//2019-09-03 14:29:34.759 CST [golog] golog.TestExample -> DEBU 004 test
```