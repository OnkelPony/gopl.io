package main

import (
	"os"
	"strings"
	"testing"
)

func BenchmarkEcho(b *testing.B) {
	os.Args = strings.Split("banik vole pyčo", " ")
	for i := 0; i < b.N; i++ {
		Echo()
	}
}
