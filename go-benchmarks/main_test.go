package main

import "testing"

func addBoth(a, b int) int {
	return a + b
}

func Benchmark(b *testing.B) {
	for b.Loop() {
		addBoth(1, 2)
	}
}
