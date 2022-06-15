package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	val, _ := strconv.Atoi(sc.Text())

	N := val

	for i := 1; i * i <= N; i++ {
		if N % i == 0 {
			fmt.Println(i);
			if i != N / i {
				fmt.Println(N / i)
			}
		}
	}
}
