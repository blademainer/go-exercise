package main

import "testing"

func BenchmarkExecute(b *testing.B) {
	Init()
	for i := 0; i < b.N; i++ {
		Execute()
	}
}

