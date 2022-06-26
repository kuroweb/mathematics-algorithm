package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	N, _ := strconv.Atoi(sc.Text())
	P := make([]int, N + 1)
	Q := make([]int, N + 1)
	b_list := []int{1, 2, 3, 4, 5}
	r_list := []int{10, 20, 30, 40, 50}
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= N; i++ {
		num := rand.Intn(len(b_list))
		P[i] = b_list[num]
	}

	for i := 1; i <= N; i++ {
		num := rand.Intn(len(r_list))
		Q[i] = r_list[num]
	}

	answer := 0.0
	for i := 1; i <= N; i++ {
		answer += float64(P[i]) / float64(Q[i])
	}

	fmt.Println(answer)
}
