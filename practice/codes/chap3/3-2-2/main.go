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

	for A >= 1 && B >= 1 {
		if A < B {
			B = B % A
		} else {
			A = A % B
		}
	}
	if A >=1 { fmt.Println(A) }
	if B >=1 { fmt.Println(B) }
}
