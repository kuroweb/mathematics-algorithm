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
	answer := 0.0

	for i := N; i >= 1; i-- {
		answer += 1.0 * float64(N) / float64(i)
	}

	fmt.Println(answer)
}
