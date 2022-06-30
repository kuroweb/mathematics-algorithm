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
		if i == 1 { dp[i] = int(math.Abs(float64(H[i - 1] - H[i]))) }
		if i >= 2 {
			v1 := dp[i - 1] + int(math.Abs(float64(H[i - 1] - H[i])))
			v2 := dp[i - 2] + int(math.Abs(float64(H[i - 2] - H[i])))
			dp[i] = int(math.Min(float64(v1), float64(v2)))
		}
	}

	fmt.Println("dp: ", dp[N - 1])
}
