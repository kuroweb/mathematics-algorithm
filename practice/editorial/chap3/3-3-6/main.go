package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	N, _ := strconv.Atoi(sc.Text())
	A := make([]int, N + 1)

	list := []int{10000, 50000, 80000}
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= N; i++ {
		num := rand.Intn(len(list))
		A[i] = list[num]
	}

	fmt.Println(("N: " + strconv.Itoa(N) + ", " + "A: " + arrayToString(A, " ")))

	// 等差数列の和の計算方法を利用
	cnt := make([]int, 100000)
	for i := 1; i <= 99999; i++ { cnt[i] = 0 }
	for i := 1; i <= N; i++ { cnt[A[i]] += 1 }

	answer := 0
	for i := 1; i <= 49999; i++ { answer += cnt[i] * cnt[100000 - i] }
	answer += cnt[50000] * (cnt[50000] - 1) / 2 // 中央では同じアドレスを参照するため 2で割る

	fmt.Println(answer)
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
