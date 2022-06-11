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
	var res int

	for _, i := range inputs {
		j, _ := strconv.Atoi(i)
		res += j
	}
	
	fmt.Println(res)
}
