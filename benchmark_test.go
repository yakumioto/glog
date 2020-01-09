package glog

import (
	"bytes"
	"log"
	"testing"
)

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
