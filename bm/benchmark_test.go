package main

import (
	"testing"
)

func BenchmarkHandleWithType(b *testing.B) {
	flag := 1
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		HandleType(flag)	
        }
}

