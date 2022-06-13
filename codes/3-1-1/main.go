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
	Answer := false

	for i := int32(2); i <= N - 1; i++ {
		if N % i == 0 {
			Answer = true
		}
	}

	fmt.Println(Answer)
}
