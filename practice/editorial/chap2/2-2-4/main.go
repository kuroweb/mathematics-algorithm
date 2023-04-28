package main

// N個の整数のmodを算出

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
    sc := bufio.NewScanner(os.Stdin)

    sc.Scan()
    N, _ := strconv.Atoi(sc.Text())

    A := make([]int, N + 1)
    for i := 1; i <= N; i++ {
        sc.Scan()
        val, _ := strconv.Atoi(sc.Text())
        A[i] = val
    }

    result := 0
    for i := 1; i <= N; i++ {
        result += A[i]
    }

    fmt.Println(result % 100)
}
