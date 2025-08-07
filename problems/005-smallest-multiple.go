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
	t := tTemp

	n := make([]int64, t)
	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n[tItr] = nTemp
	}

	for i := int64(0); i < t; i++ {
		fmt.Println(findSmallestMultiple(n[i]))
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
