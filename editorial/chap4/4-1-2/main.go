package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// N
	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())

	// x, y
	x := make([]int, N + 1)
	y := make([]int, N + 1)
	for i := 1; i <= N; i++ {
		sc.Scan()
		x[i], _ = strconv.Atoi(sc.Text())
		sc.Scan()
		y[i], _ = strconv.Atoi(sc.Text())
	}

	answer := 0.0
	for i := 1; i <= N; i++ {
		for j := i + 1; j <= N; j++ {
			dist := math.Sqrt(float64((x[i] - x[j]) * (x[i] - x[j]) + (y[i] - y[j]) * (y[i] - y[j])))
			answer = Min(answer, dist)
		}
	}

	fmt.Println(answer)
}

func Min(x, y float64) float64 {
	if x > y {
		return y
	}

	return x
}
