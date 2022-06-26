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
	A := make([]int, N + 1)
	B := make([]int, N + 1)
	a_list := []int{1, 2, 3, 4, 5}
	b_list := []int{10, 20, 30, 40, 50}
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= N; i++ {
		num := rand.Intn(len(a_list))
		A[i] = a_list[num]
	}

	for i := 1; i <= N; i++ {
		num := rand.Intn(len(b_list))
		B[i] = b_list[num]
	}

	fmt.Println(A, B)

	a_exp, b_exp := 0.0, 0.0
	for i := 1; i <= N; i++ {
		a_exp += float64(A[i]) / float64(N)
		b_exp += float64(B[i]) / float64(N)
	}

	fmt.Println(a_exp + b_exp)
}
