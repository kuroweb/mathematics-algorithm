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
	A := make([]int, N + 1)
	for i := 1; i <= N; i++ {
		sc.Scan()
		num, _ := strconv.Atoi(sc.Text())
		A[i] = num
	}

	fmt.Println(N, A)

	for i := 1; i <= N; i++ {
		Min := i
		Min_Value := A[i]
		for j := i + 1; j <= N; j++ {
			if A[j] < Min_Value {
				Min = j
				Min_Value = A[j]
			}
		}
		tmp := A[i]
		A[i] = A[Min]
		A[Min] = tmp
	}

	fmt.Println(N, A)
}
