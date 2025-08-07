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

	t, err := strconv.ParseInt(strings.TrimSpace(readLineDiffSum(reader)), 10, 64)
	checkErrorDiffSum(err)

	n := make([]int64, t)
	for tItr := 0; tItr < int(t); tItr++ {
		ntemp, err := strconv.ParseInt(strings.TrimSpace(readLineDiffSum(reader)), 10, 64)
		checkErrorDiffSum(err)
		n[tItr] = ntemp
	}

	for _, num := range n {
		result := calculateDifference(uint64(num))
		fmt.Println(result)
	}
}

func calculateDifference(n uint64) uint64 {
	// Sum of first n numbers: n(n+1)/2
	sumOfNumbers := n * (n + 1) / 2

	// Sum of squares of first n numbers: n(n+1)(2n+1)/6
	sumOfSquares := n * (n + 1) * (2*n + 1) / 6

	// Difference: (sum of numbers)^2 - sum of squares
	return sumOfNumbers*sumOfNumbers - sumOfSquares
}

func readLineDiffSum(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkErrorDiffSum(err error) {
	if err != nil {
		panic(err)
	}
}
