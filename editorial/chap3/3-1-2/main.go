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
	val, _ := strconv.Atoi(sc.Text())

	N := int32(val)

	for i := int32(2); i * i <= N; i++ {
		for N % i == 0 {
			N /= i
			fmt.Println(i)
		}
	}

	if N >= 2 {
		fmt.Println(N)
	}
}
