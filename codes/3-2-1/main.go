package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var ary [2]int
	for i := 0; i < 2; i++ {
		sc.Scan()
		val, _ := strconv.Atoi(sc.Text())
		ary[i] = val
	}

	A, B := ary[0], ary[1]

	for i := 1; i <= Min(A, B); i++ {
		if A % i == 0 && B % i == 0 { fmt.Println(i) }
	}
}

func Min(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
