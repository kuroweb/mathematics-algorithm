package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// N: int16   => カード枚数
// S: int16   => 合計値の期待値
// A: []int16 => カードに書かれた整数を配列に格納したもの

func main() {
	sc := bufio.NewScanner(os.Stdin)
	
	// N
	fmt.Println("Input: N")
	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())

	// S
	fmt.Println("Input: S")
	sc.Scan()
	S, _ := strconv.Atoi(sc.Text())

	// A
	A := make([]int, N + 1)
	for i := 1; i <= N; i++ {
		fmt.Println("Input: A")
		sc.Scan()
		A[i], _ = strconv.Atoi(sc.Text())
	}

	fmt.Println("N: ", N, "S: ", S, "A: ", A)

	// dp初期化
	dp := make([][]bool, N + 1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]bool, S + 1)
	}
	dp[0][0] = true

	// 動的計画法
	for i := 1; i <= N; i++ {
		for j := 0; j <= S; j++ {
			if j < A[i] {
				// j < A[i] のとき、カードi を選べない
				dp[i][j] = dp[i - 1][j]
			} else {
				// j >= A[i] のとき、選ぶ / 選ばない 両方の選択肢がある
				if dp[i - 1][j] == true || dp[i - 1][j - A[i]] == true {
					dp[i][j] = true
				} else {
					dp[i][j] = false
				}
			}
		}
	}

	if dp[N][S] == true {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
