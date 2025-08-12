package main

import "fmt"

func main5() {
	var t int
	fmt.Scanln(&t)
	for range t {
		var n int64
		fmt.Scanln(&n)
		fmt.Println(findSmallestMultiple(n))
	}
}

func findSmallestMultiple(n int64) int64 {
	// Calculate LCM of numbers from 1 to n
	result := int64(1)

	for i := int64(2); i <= n; i++ {
		result = lcm(result, i)
	}

	return result
}

// EBOB (Greatest Common Divisor)
func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// EKOK (Least Common Multiple)
func lcm(a, b int64) int64 {
	return a * b / gcd(a, b)
}
