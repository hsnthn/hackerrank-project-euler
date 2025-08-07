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
    n := make([]uint64, t)
    k := make([]int, t)
    number := make([]string, t)
    for tItr := 0; tItr < int(t); tItr++ {
        firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
        checkError(err)
        n[tItr] = uint64(nTemp)

        kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
        checkError(err)
        k[tItr] = int(kTemp)

        number[tItr] = readLine(reader)
    }
    for i := 0; i < int(t); i++ {
        fmt.Println(greatestConsecutiveProduct(number[i], k[i]))
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
