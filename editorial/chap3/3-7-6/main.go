package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// N: int16   => 夏休みの素数
// A: []int16 => i日目に勉強すると向上する学力の数値

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// N
	fmt.Println("Input: N")
	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())

	// A
	A := make([]int, N + 1)
	for i := 1; i <= N; i++ {
		fmt.Println("Input: A")
		sc.Scan()
		A[i], _ = strconv.Atoi(sc.Text())
	}

	// dp初期化
	dp1 := make([]int, N + 1)
	dp2 := make([]int, N + 1)

	// 動的計画法
	for i := 1; i <= N; i++ {
		dp1[i] = dp2[i - 1] + A[i]
		dp2[i] = Max(dp1[i - 1], dp2[i - 1])
	}

	answer := Max(dp1[N], dp2[N])
	fmt.Println(answer)
}

func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
