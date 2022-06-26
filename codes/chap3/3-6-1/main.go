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
	sc.Scan()

	N, _ := strconv.Atoi(sc.Text())
	A := make([]int, N)
	for i := 1; i <= N; i++ {
		sc.Scan()
		num, _ := strconv.Atoi(sc.Text())
		A[i - 1] = num
	}

	fmt.Println(A)

	sort.Slice(A, func(i, j int) bool { return A[i] < A[j] })

	fmt.Println(A)
}
