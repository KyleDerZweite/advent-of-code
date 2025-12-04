package main

import (
	"testing"
)

func BenchmarkSolve(b *testing.B) {
	instructions := parseInput("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solve(50, instructions)
	}
}

func BenchmarkParseAndSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		instructions := parseInput("input.txt")
		solve(50, instructions)
	}
}
