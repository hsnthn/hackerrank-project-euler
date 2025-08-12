package main

import "fmt"

func main2() {
	var t int
	fmt.Scanln(&t)
	for range t {
		var n int
		fmt.Scanln(&n)
		fmt.Println(sumEvenFibNums(n))
	}
}

func sumEvenFibNums(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	a := 0
	b := 1
	sum := a + b
	evenSum := 0
	for i := 1; sum < n; i++ {
		if sum%2 == 0 {
			evenSum += sum
		}
		sum = a + b
		a = b
		b = sum
	}
	return evenSum
}
