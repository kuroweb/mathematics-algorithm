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

	// 各点の座標を取得
	//// 点a
	sc.Scan()
	ax, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	ay, _ := strconv.Atoi(sc.Text())

	//// 点b
	sc.Scan()
	bx, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	by, _ := strconv.Atoi(sc.Text())

	//// 点c
	sc.Scan()
	cx, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	cy, _ := strconv.Atoi(sc.Text())

	// 各ベクトルの成分表示を算出
	//// BA
	BAx := ax - bx; BAy := ay - by
	//// BC
	BCx := cx - bx; BCy := cy - by
	//// CA
	CAx := ax - cx; CAy := ay - cy
	//// CB
	CBx := bx - cx; CBy := by - cy

	// 座標パターンを検証
	pattern := 2
	if BAx * BCx + BAy * BCy < 0 {
		pattern = 1
	}
	if CAx * CBx + CAy * CBy < 0 {
		pattern = 3
	}

	answer := 0.0
	if pattern == 1 {
		answer = math.Sqrt(float64(BAx * BAx + BAy * BAy))
	}
	if pattern == 3 {
		answer = math.Sqrt(float64(CAx * CAx + CAy * CAy))
	}
	if pattern == 2 {
		S := math.Abs(float64(BAx * BCy - BAy * BCx))
		BCLength := math.Sqrt(float64(BCx * BCx + BCy * BCy))
		answer = S / BCLength
	}

	fmt.Println(answer)
}
