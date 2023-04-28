package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// N
	fmt.Println("Input: N")
	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Println("Input: A")
		sc.Scan()
		A[i], _ = strconv.Atoi(sc.Text())
	}

	// X
	fmt.Println("Input: X")
	sc.Scan()
	X, _ := strconv.Atoi(sc.Text())

	// 配列をソート
	sort.Slice(A, func(i, j int) bool { return A[i] < A[j] })

	answer := "No"
	left := 0
	right := N - 1
	for left <= right {
		mid := (left + right) / 2
		if A[mid] == X {
			answer = "Yes"
			break
		}
		if A[mid] > X {
			right = mid - 1
		}
		if A[mid] < X {
			left = mid + 1
		}
	}

	fmt.Println(answer)
}
