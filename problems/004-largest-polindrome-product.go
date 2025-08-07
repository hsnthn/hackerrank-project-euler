package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	n := make([]int64, t)
	for i := 0; i < int(t); i++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n[i] = nTemp
	}

	for i := 0; i < int(t); i++ {
		fmt.Println(largestPalindromeProduct(n[i]))
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

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
