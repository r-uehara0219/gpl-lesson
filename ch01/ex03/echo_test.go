package main

import (
	"testing"
)

func BenchmarkEcho(b *testing.B) {
	args := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	for i := 0; i < b.N; i++ {
		echo(args)
	}
}

func BenchmarkSlowEcho(b *testing.B) {
	args := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	for i := 0; i < b.N; i++ {
		slow_echo(args)
	}
}
