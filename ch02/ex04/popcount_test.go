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

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(0x1234567890ABCDEF)
	}
}

// cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
// BenchmarkPopCount-8             	1000000000	         0.2670 ns/op
// BenchmarkPopCountByShifting-8   	47825630	        24.74 ns/op
