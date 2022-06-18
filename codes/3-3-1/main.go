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
		A[i] = 150 + i
	}

	var answer int
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			for k := 1; k <= N; k++ {
				for l := 1; l <= N; l++ {
					for m := 1; m <= N; m++ {
						if A[i] + A[j] + A[k] + A[l] + A[m] == 1000 {
							answer++
						}
					}
				}
			}
		}
	}

	fmt.Println(answer)
}
