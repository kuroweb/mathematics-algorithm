package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputs := strings.Split(scanner.Text(), " ")
	values := []int{}

	for _, i := range inputs {
		j, _ := strconv.Atoi(i)
		values = append(values, j)
	}
	
	for _, i := range values {
		fmt.Println(i * 2)
	}
}
