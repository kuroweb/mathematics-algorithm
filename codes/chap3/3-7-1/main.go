package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	N, _ := strconv.Atoi(sc.Text())
	H := make([]int, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		num, _ := strconv.Atoi(sc.Text())
		H[i] = num
	}

	fmt.Println("A: ", N, "H: ", H)

	dp := make([]int, N)
	for i := 0; i < N; i++ {
		if i == 0 { dp[i] = 0 }
		if i == 1 { dp[i] = abs(H[i - 1] - H[i]) }
		if i >= 2 {
			v1 := dp[i - 1] + abs(H[i - 1] - H[i])
			v2 := dp[i - 2] + abs(H[i - 2] - H[i])
			dp[i] = min(v1, v2)
		}
	}

	fmt.Println("dp: ", dp[N - 1])
}

func abs(a int) int {
	abs := math.Abs(float64(a))
	return int(abs)
}

func min(a, b int) int {
	min := math.Min(float64(a), float64(b))
	return int(min)
}
