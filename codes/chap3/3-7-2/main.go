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
	dp := make([]int, N + 1)
	for i := 0; i < N + 1; i++ {
		if i < 2 {
			dp[i] = 1
		} else {
			dp[i] = dp[i - 1] + dp[i - 2]
		}
	}
	fmt.Println(dp[N])
}
