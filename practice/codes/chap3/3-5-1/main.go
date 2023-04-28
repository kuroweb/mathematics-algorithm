package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	N := 10000
	M := 0

	for i := 1; i <= N; i++ {
		px := random(0, 1)
		py := random(0, 1)

		if px * px + py * py <= 1.0 {
			M++
		}
	}

	fmt.Println(4.0 * float64(M) / float64(N))
}

func random(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()*(max-min) + min
}
