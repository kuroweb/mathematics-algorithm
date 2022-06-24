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

	fmt.Println(fact(N))
}

func fact (N int) int {
	if N == 1 { return 1 }

	return fact(N - 1) * N
}
