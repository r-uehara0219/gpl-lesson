package main

import (
	"testing"
)

// -- Benchmarks --

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountWithLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountWithLoop(0x1234567890ABCDEF)
	}
}

// cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
// BenchmarkPopCount-8           	1000000000	         0.2803 ns/op
// BenchmarkPopCountWithLoop-8   	703718236	         1.665 ns/op
