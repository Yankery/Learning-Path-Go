package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestCalculatedSquares(t *testing.T) {
	n := 10
	expected := ""
	for i := 1; i <= n; i++ {
		square := i * i
		expected += fmt.Sprintf("%d ", square)
	}
	expected = strings.TrimSpace(expected)

	squares := CalculateSquaresBitwise(n)
	result := strings.Trim(fmt.Sprint(squares), "[]")

	if result != expected {
		t.Errorf("Unexpected output: got %q, want %q", result, expected)
	}
}

func BenchmarkCalculatedSquares(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateSquares(10)
	}
}

func BenchmarkCalculateSquaresAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateSquaresAppend(10)
	}
}

func BenchmarkCalculateSquaresEq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateSquaresEq(10)
	}
}

func BenchmarkCalculateSquaresMath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateSquaresMath(10)
	}
}

func BenchmarkCalculateSquaresBitwise(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateSquaresBitwise(10)
	}
}
