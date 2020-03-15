package glog

import (
	"os"
)

func ExampleLogger() {
	l := New(os.Stderr, "glog")
	l.SetLevel(LevelDebug)

	l.Infof("test example")
	l.Warnf("test example")
	l.Errof("test example")
	l.Debuf("test example")

	l.ResetID()
	// Output:
	//2020-03-16 00:18:39.087 CST [glog] ExampleLogger -> INFO 001 test example
	//2020-03-16 00:18:39.087 CST [glog] ExampleLogger -> WARN 002 test example
	//2020-03-16 00:18:39.087 CST [glog] ExampleLogger -> ERRO 003 test example
	//2020-03-16 00:18:39.087 CST [glog] ExampleLogger -> DEBU 004 test example
}
