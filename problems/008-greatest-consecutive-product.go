package main

import (
	"fmt"
	"strconv"
)

func main8() {
	var t int
	fmt.Scanln(&t)
	for range t {
		var n uint64
		var k int
		fmt.Scanln(&n, &k)
		var number string
		fmt.Scanln(&number)
		fmt.Println(greatestConsecutiveProduct(number, k))
	}
}

func greatestConsecutiveProduct(numberStr string, k int) uint64 {
	maxProduct := uint64(0)
	l := 0
	r := k - 1

	product := uint64(1)
	for i := 0; i < k; i++ {
		digitInt, _ := strconv.Atoi(string(numberStr[i]))
		product *= uint64(digitInt)
	}
	maxProduct = product

	for r+1 < len(numberStr) {
		leftDigit, _ := strconv.Atoi(string(numberStr[l]))
		rightDigit, _ := strconv.Atoi(string(numberStr[r+1]))

		if leftDigit == 0 {
			// Recalculate window because we can't divide by 0
			product = 1
			for i := l + 1; i <= r+1; i++ {
				d := numberStr[i] - '0'
				product *= uint64(d)
			}
		} else {
			product /= uint64(leftDigit)
			product *= uint64(rightDigit)
		}

		if product > maxProduct {
			maxProduct = product
		}

		l++
		r++
	}

	return maxProduct
}
