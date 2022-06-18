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

	list := []int{1, 2, 3}
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= N; i++ {
		num := rand.Intn(len(list))
		A[i] = list[num]
	}

	fmt.Println(("N: " + strconv.Itoa(N) + ", " + "A: " + arrayToString(A, " ")))

	x, y, z := 0, 0, 0
	for i := 1; i <= N; i ++ {
		if A[i] == 1 { x++ }
		if A[i] == 2 { y++ }
		if A[i] == 3 { z++ }
	}

	fmt.Println(x, y, z)
	fmt.Println(x * (x - 1) / 2 + y * (y - 1) / 2 + z * (z - 1) / 2)
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
