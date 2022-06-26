package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var N int
var A []int

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	N, _ := strconv.Atoi(sc.Text())

	A = make([]int, N + 1)
	for i := 1; i <= N; i++ {
		A[i] = i
	}

	fmt.Println(N, A)

	answer := solve(1, N + 1)
	fmt.Println(answer)
}

func solve(l, r int) int {
	if r - l == 1 { return A[l] }

	m := (l + r) / 2
	s1 := solve(l, m)
	s2 := solve(m, r)
	return s1 + s2
}
