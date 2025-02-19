package main

import (
	"fmt"
	"math"
)

func CalculateSquares(n int) []int {
	squares := make([]int, n)
	for i := 1; i <= n; i++ {
		squares[i-1] = i * i
	}
	return squares
}

func CalculateSquaresAppend(n int) []int {
	squares := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		squares = append(squares, i*i)
	}
	return squares
}

func CalculateSquaresEq(n int) []int {
	squares := make([]int, n)
	square := 0
	for i := 0; i < n; i++ {
		square += 2*i + 1
		squares[i] = square
	}
	return squares
}

func CalculateSquaresMath(n int) []int {
	squares := make([]int, n)
	for i := 1; i <= n; i++ {
		squares[i-1] = int(math.Pow(float64(i), 2))
	}
	return squares
}

func CalculateSquaresBitwise(n int) []int {
	squares := make([]int, n)
	for i := 1; i <= n; i++ {
		squares[i-1] = i * i
	}
	return squares
}

func main() {
	squares := CalculateSquaresAppend(10)
	for i := 0; i < len(squares); i++ {
		fmt.Printf("The square of %d is %d\n", i+1, squares[i])
	}
}
