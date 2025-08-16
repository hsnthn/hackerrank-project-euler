package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// https://www.hackerrank.com/contests/projecteuler/challenges/euler011/problem
func main11() {
    reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

    var grid [][]int32
    for i := 0; i < 20; i++ {
        gridRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

        var gridRow []int32
        for _, gridRowItem := range gridRowTemp {
            gridItemTemp, err := strconv.ParseInt(gridRowItem, 10, 64)
            checkError(err)
            gridItem := int32(gridItemTemp)
            gridRow = append(gridRow, gridItem)
        }

        if len(gridRow) != 20 {
            panic("Bad input")
        }

        grid = append(grid, gridRow)
    }

    // TODO: Add the solution logic here
    fmt.Println(findLargestProduct(grid))
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

func findLargestProduct(grid [][]int32) int64 {
    maxProduct := int64(0)
    for x := 0; x < 20; x++ {
        for y := 0; y < 20; y++ {
            product := left(grid, x, y)
            product = max(product, right(grid, x, y))
            product = max(product, up(grid, x, y))
            product = max(product, down(grid, x, y))
            product = max(product, diagonal(grid, x, y))
            maxProduct = max(maxProduct, product)
        }
    }
    return maxProduct
}

func left(grid [][]int32, x int, y int) int64 {
    if y+3 >= 20 {
        return 0
    }
    product := int64(1)
    for i := 0; i < 4; i++ {
        product *= int64(grid[x][y+i])
    }
    return product
}

func right(grid [][]int32, x int, y int) int64 {
    if y-3 < 0 {
        return 0
    }
    product := int64(1)
    for i := 0; i < 4; i++ {
        product *= int64(grid[x][y-i])
    }
    return product
}

func up(grid [][]int32, x int, y int) int64 {
    if x+3 >= 20 {
        return 0
    }
    product := int64(1)
    for i := 0; i < 4; i++ {
        product *= int64(grid[x+i][y])
    }
    return product
}

func down(grid [][]int32, x int, y int) int64 {
    if x-3 < 0 {
        return 0
    }
    product := int64(1)
    for i := 0; i < 4; i++ {
        product *= int64(grid[x-i][y])
    }
    return product
}

func diagonal(grid [][]int32, x int, y int) int64 {
    // Check both diagonal directions
    product1 := int64(0)
    product2 := int64(0)

    // Diagonal down-right
    if x+3 < 20 && y+3 < 20 {
        product1 = int64(1)
        for i := 0; i < 4; i++ {
            product1 *= int64(grid[x+i][y+i])
        }
    }

    // Diagonal down-left
    if x+3 < 20 && y-3 >= 0 {
        product2 = int64(1)
        for i := 0; i < 4; i++ {
            product2 *= int64(grid[x+i][y-i])
        }
    }

    return max(product1, product2)
}
