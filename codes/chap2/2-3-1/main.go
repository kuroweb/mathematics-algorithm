package main

import "fmt"

var cnt = 1000

func main() {
	fmt.Println(func1())
	fmt.Println(func2(500))
	fmt.Println(func2(500))
}

func func1() int {
	return 2021
}

func func2(pos int) int {
	cnt += 1
	return cnt + pos
}
