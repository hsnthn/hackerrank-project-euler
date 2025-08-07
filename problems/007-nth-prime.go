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

	t, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	n := make([]int64, t)
	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n[tItr] = nTemp
	}

	for _, index := range n {
		fmt.Println(nthPrime(index))
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
