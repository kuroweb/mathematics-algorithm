package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	val, _ := strconv.Atoi(sc.Text())
	N := val

	A := make([]int, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		val, _ := strconv.Atoi(sc.Text())
		A[i] = val
	}

	R := LCM(A[0], A[1])
	for i := 2; i < N; i++ {
		R = LCM(R, A[i])
	}

	fmt.Println(R)
}

func GCD(A int, B int) int {
	for A >= 1 && B >= 1 {
		if A < B {
			B %= A
		} else {
			A %= B
		}
	}
	if A >= 1 {
		return A
	}
	return B
}

func LCM(A int, B int) int {
	return A / GCD(A, B) * B
}
