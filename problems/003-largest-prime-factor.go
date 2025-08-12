package main

import "fmt"

func main3() {
	var t int
	fmt.Scanln(&t)
	for range t {
		var n int64
		fmt.Scanln(&n)
		fmt.Println(findLargestNum(n))
	}
}

func findLargestNum(n int64) int64 {
	for i := int64(2); i*i <= n; i++ {
		if i*i == n {
			return i
		}
		for n%i == 0 {
			n = n / i
		}
	}

	return n
}
