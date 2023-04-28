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

	N, _ := strconv.Atoi(sc.Text())

	A := make([]int, N + 1)
	C := make([]int, N + 1)
	for i := 1; i <= N; i++ {
		sc.Scan()
		num, _ := strconv.Atoi(sc.Text())
		A[i] = num
	}

	MergeSort(A, C, 1, N + 1)

	fmt.Println(A, C)

	for i := 1; i <= N; i++ {
		fmt.Println(A[i])
	}
}

func MergeSort(A []int, C []int, l int, r int) {
	if r - l == 1 { return }

	// 分割統治法
	m := (l + r) / 2
	MergeSort(A, C, l, m)
	MergeSort(A, C, m, r)

	// Merge
	c1, c2, cnt := l, m, 0
	for c1 != m || c2 != r {
		if c1 == m {
			C[cnt] = A[c2]
			c2++
		} else if c2 == r {
			C[cnt] = A[c1]
			c1++
		} else {
			if A[c1] <= A[c2] {
				C[cnt] = A[c1]
				c1++
			} else {
				C[cnt] = A[c2]
				c2++
			}
		}
		cnt++
	}
	for i := 0; i < cnt; i++ { A[l + i] = C[i] }
}
