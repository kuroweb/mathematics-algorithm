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
	A, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	B, _ := strconv.Atoi(sc.Text())

	fmt.Println(GCD(A, B))
}

func GCD(A, B int) int {
	if B == 0 { return A }

	return GCD(B, A % B)
}
