package main

import (
	"fmt"
	"math"
)

func main9() {
	var t int
	fmt.Scanln(&t)
	for range t {
		var n int
		fmt.Scanln(&n)
		fmt.Println(findSpecialPythagorienTriplet(n))
	}
}

func findSpecialPythagorienTriplet(n int) int {
	upper := int(float64(n) * (1 - math.Sqrt2/2))

	for x := upper; x > 0; x-- {
		numerator := n * (2*x - n)
		denominator := 2*x - 2*n

		if denominator == 0 || numerator%denominator != 0 {
			continue
		}

		y := numerator / denominator
		z := n - x - y

		return x * y * z
	}

	return -1
}
