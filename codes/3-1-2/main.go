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

	N := val
	Answer := false

	for i := 2; i * i <= N; i++ {
		if N % i == 0 {
			Answer = true
		}
	}

	fmt.Println(Answer)
}
