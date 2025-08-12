package main

import "fmt"

func main6() {
	var t int
	fmt.Scanln(&t)
	for range t {
		var n uint64
		fmt.Scanln(&n)
		fmt.Println(calculateDifference(n))
	}
}

func calculateDifference(n uint64) uint64 {
	// Sum of first n numbers: n(n+1)/2
	sumOfNumbers := n * (n + 1) / 2

	// Sum of squares of first n numbers: n(n+1)(2n+1)/6
	sumOfSquares := n * (n + 1) * (2*n + 1) / 6

	// Difference: (sum of numbers)^2 - sum of squares
	return sumOfNumbers*sumOfNumbers - sumOfSquares
}
