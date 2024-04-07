package main

import "testing"

func BenchmarkWrong(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		add(1000000000, 1000000001)
	}
}
