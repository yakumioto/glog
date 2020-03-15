package glog

import (
	"os"
	"testing"
)

func TestExample(t *testing.T) {
	l := New(os.Stderr, "glog")
	l.SetLevel(LevelDebug)

	l.Infof("test example")
	l.Warnf("test example")
	l.Errof("test example")
	l.Debuf("test example")

	l.ResetID()
}
