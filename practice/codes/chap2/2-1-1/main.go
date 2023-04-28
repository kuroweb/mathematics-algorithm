package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	orange_size, _ := strconv.Atoi(scanner.Text())
	apple_size := 5

	fmt.Println(orange_size + apple_size)
}
