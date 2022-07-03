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

	// 商品個数
	fmt.Println("input: N")
	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())

	// 許容最大重量
	fmt.Println("input: W")
	sc.Scan()
	W, _ := strconv.Atoi(sc.Text())

	// 商品の重量・価値
	w := make([]int, N + 1)
	v := make([]int, N + 1)
	for i := 1; i <= N; i++ {
		fmt.Println("input: weight")
		sc.Scan()
		num, _ := strconv.Atoi(sc.Text())

		w[i] = num
		v[i] = num
	}

	// 結果の入れ物を初期化
	inf := math.MaxInt16
	dp := make([][]int, N + 1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, W + 1)
	}
	for i := 1; i <= W; i++ {
		dp[0][i] = -inf
	}

	// 動的計画法
	for i := 1; i <= N; i++ {
		for j := 0; j <= W; j++ {
			if j < w[i] {
				// j < w[i] のとき、方法B は選べない
				dp[i][j] = dp[i - 1][j]
			} else {
				//  j >= w[i] のとき、方法A・方法B どうちらも選べる
				dp[i][j] = Max(dp[i - 1][j], dp[i - 1][j - w[i]] + v[i])
			}
		}
	}

	answer := 0
	for i := 0; i <= W; i++ {
		answer = Max(answer, dp[N][i])
	}
	fmt.Println(answer)
}

func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
