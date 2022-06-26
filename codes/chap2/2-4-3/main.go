package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var ary [2]int
	for i := 0; i < 2; i++ {
		sc.Scan()
		val, _ := strconv.Atoi(sc.Text())
		ary[i] = val
	}
	N, S := ary[0], ary[1]

	var cnt int
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			if (i + j <= S) { cnt++ }
		}
	}

	fmt.Println(cnt)
}
