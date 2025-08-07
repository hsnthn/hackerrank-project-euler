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
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    t := int32(tTemp)

    n := make([]int, t)
    for i := 0; i < int(t); i++ {
        nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        n[i] = int(nTemp)
    }
    
    sum := make([]int, t)
    for i := 0; i < int(t); i++{
        sum[i] = sumEvenFibNums(n[i])
    } 
        
    for _, s := range sum {
        fmt.Println(s)
    }
}

func sumEvenFibNums(n int) int {
    if n == 0 || n == 1 {
        return n
    }

    a := 0
    b := 1
    sum := a + b
    evenSum := 0
    for i := 1; sum < n; i++ {
        if sum%2 == 0 {
            evenSum += sum
        }
        sum = a + b
        a = b
        b = sum
    }
    return evenSum
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
