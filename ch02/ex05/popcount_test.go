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

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(0x1234567890ABCDEF)
	}
}

// cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
// BenchmarkPopCount-8             	1000000000	         0.2887 ns/op
// BenchmarkPopCountByClearing-8   	79573148	        14.06 ns/op
