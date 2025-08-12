package main

import "fmt"

func main7() {
	var t int
	fmt.Scanln(&t)
	for range t {
		var n int64
		fmt.Scanln(&n)
		fmt.Println(nthPrime(n))
	}
}

func nthPrime(index int64) int64 {
	var result int64
	for maybePrime := int64(2); index > 0; maybePrime++ {
		if isPrime(maybePrime) {
			index--
			result = maybePrime
		}
	}
	return result
}

func isPrime(maybePrime int64) bool {
	if maybePrime < 2 {
		return false
	}
	if maybePrime == 2 {
		return true
	}
	if maybePrime%2 == 0 {
		return false
	}
	for i := int64(2); i < maybePrime; i++ {
		if maybePrime%i == 0 {
			return false
		}
	}
	return true
}
