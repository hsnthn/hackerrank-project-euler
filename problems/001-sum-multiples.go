package main

import "fmt"

func main1() {
	var t int
	fmt.Scanln(&t)
	for range t {
		var n int32
		fmt.Scanln(&n)
		fmt.Println(sumMultiples(n))
	}
}

func sumMultiples(n int32) int64 {
	multiples3 := int64((n - 1) / 3)
	multiples5 := int64((n - 1) / 5)
	multiples15 := int64((n - 1) / 15)
	return 3*(multiples3*(multiples3+1)/2) + 5*(multiples5*(multiples5+1)/2) - 15*(multiples15*(multiples15+1)/2)
}
