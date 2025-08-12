package main

import "fmt"

func main11() {
	var t int
	fmt.Scanln(&t)
	for range t {
		var n int64
		fmt.Scanln(&n)
		fmt.Println(sumPrimes3(n))
	}
}

func sumPrimes3(n int64) int64 {
	divisible := make([]bool, n+1)
	var sum int64
	for i := int64(2); i <= n; i++ {
		if !divisible[i] {
			if isPrime1(int(i)) {
				sum += i
				// Mark multiples of this prime as divisible
				for j := i * i; j <= n; j += i {
					divisible[j] = true
				}
			}
		}
	}
	return sum
}

func isPrime1(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// func sumPrimes(n int64,  []bool) int64 {
//     sum := int64(0)
//     for i := int64(2); i <= n; i++ {
//         if !divisible[i] {
//             sum += i
//         }
//     }
//     return sum
// }

// func sumPrimes2(n int64, divisible []bool) int64 {
//     sum := int64(0)
//     for i := int64(2); i <= n; i++ {
//         if !divisible[i] {
//             sum += i
//             for j := i * i; j <= n; j += i {
//                 divisible[j] = true
//             }
//         }
//     }
//     return sum
// }
