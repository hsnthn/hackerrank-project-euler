package main

import (
	"fmt"
	"strconv"
)

func main4() {
	var t int
	fmt.Scanln(&t)
	for range t {
		var n int64
		fmt.Scanln(&n)
		fmt.Println(largestPalindromeProduct(n))
	}
}

func largestPalindromeProduct(n int64) int64 {
	maxPalindrome := int64(0)

	for i := 999; i >= 100; i-- {
		for j := i; j >= 100; j-- {
			product := int64(i * j)
			if product >= n {
				continue
			}
			if isPalindrome(strconv.FormatInt(product, 10)) && product > maxPalindrome {
				maxPalindrome = product
			}
		}
	}

	return maxPalindrome
}

func isPalindrome(n string) bool {
	i, j := 0, len(n)-1
	for i < j {
		if n[i] != n[j] {
			return false
		}
		i++
		j--
	}
	return true
}
