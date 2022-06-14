package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var ary [2]int64
	for i := 0; i < 2; i++ {
		sc.Scan()
		val, _ := strconv.Atoi(sc.Text())
		ary[i] = int64(val)
	}

	A, B := ary[0], ary[1]

	for i := int64(1); i <= Min(A, B); i++ {
		if A % i == 0 && B % i == 0 { fmt.Println(i) }
	}
}

func Min(x int64, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
