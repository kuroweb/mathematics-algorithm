package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var ary [3]int
	for i := 0; i < 3; i++ {
		sc.Scan()
		val, _ := strconv.Atoi(sc.Text())
		ary[i] = val
	}
	N, X, Y := ary[0], ary[1], ary[2]

	cnt := 0
	for i := 1; i <= N; i++ {
		if i % X == 0 || i % Y == 0 { cnt++ }
	}

	fmt.Println(cnt)
}
