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
	B := make([]int, N + 1)
	R := make([]int, N + 1)
	b_list := []int{1, 2, 3, 4, 5}
	r_list := []int{10, 20, 30, 40, 50}
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= N; i++ {
		num := rand.Intn(len(b_list))
		B[i] = b_list[num]
	}

	for i := 1; i <= N; i++ {
		num := rand.Intn(len(r_list))
		R[i] = r_list[num]
	}

	b_exp, r_exp := 0.0, 0.0
	for i := 1; i <= N; i++ {
		b_exp += float64(B[i]) / float64(N)
		r_exp += float64(R[i]) / float64(N)
	}

	fmt.Println(b_exp + r_exp)
}
