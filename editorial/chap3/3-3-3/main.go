package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc :=bufio.NewScanner(os.Stdin)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())

	conb_n, conb_r := 1, 1
	for i := n; i > n - r; i-- { conb_n *= i }
	for i := 1; i <= r; i++ { conb_r *= i }

	fmt.Println(conb_n / conb_r)
}
