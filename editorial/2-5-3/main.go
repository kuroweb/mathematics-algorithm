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
	var Answer int32 = 1

	for i := int32(1); i <= N; i++ {
		Answer = Answer * i
	}

	fmt.Println(Answer)
}
