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

	n := make([]int32, t)
	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n[tItr] = int32(nTemp)
	}

	sum := make([]int64, t)
	for tItr := 0; tItr < int(t); tItr++ {
		sum[tItr] = int64(sumMultiples(n[tItr]))
	}

	for _, s := range sum {
		fmt.Println(s)
	}
}

func sumMultiples(n int32) int64 {
	multiples3 := int64((n - 1) / 3)
	multiples5 := int64((n - 1) / 5)
	multiples15 := int64((n - 1) / 15)
	return 3*(multiples3*(multiples3+1)/2) + 5*(multiples5*(multiples5+1)/2) - 15*(multiples15*(multiples15+1)/2)
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
