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

	list := []int{100, 200, 300, 400}
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= N; i++ {
		num := rand.Intn(len(list))
		A[i] = list[num]
	}

	fmt.Println(("N: " + strconv.Itoa(N) + ", " + "A: " + arrayToString(A, " ")))

	a, b, c, d := 0, 0, 0, 0
	for i := 1; i <= N; i++ {
		if A[i] == 100 { a++ }
		if A[i] == 200 { b++ }
		if A[i] == 300 { c++ }
		if A[i] == 400 { d++ }
	}

	fmt.Println(a, b, c, d)
	fmt.Println(a * d + b * c)
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
