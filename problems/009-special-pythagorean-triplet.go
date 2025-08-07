package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

    tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    t := int32(tTemp)

    n := make([]int, t)
    for tItr := 0; tItr < int(t); tItr++ {
        nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        n[tItr] = int(nTemp)
    }

    for _, number := range n {
        fmt.Println(findSpecialPythagorienTriplet(number))
    }
}

func findSpecialPythagorienTriplet(n int) int {
    upper := int(float64(n) * (1 - math.Sqrt2/2))

    for x := upper; x > 0; x-- {
        numerator := n * (2*x - n)
        denominator := 2*x - 2*n

        if denominator == 0 || numerator%denominator != 0 {
            continue
        }

        y := numerator / denominator
        z := n - x - y

        return x * y * z
    }

    return -1
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
