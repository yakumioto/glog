# Glog

[![Github Action](https://github.com/yakumioto/glog/workflows/glog/badge.svg)](https://github.com/yakumioto/glog/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/yakumioto/glog)](https://goreportcard.com/report/github.com/yakumioto/glog)
[![codecov](https://codecov.io/gh/yakumioto/glog/branch/master/graph/badge.svg)](https://codecov.io/gh/yakumioto/glog)


Simple level log package in Go.

## Installation

`go get github.com/yakumioto/glog`

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