package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func factorial(input int) int {
	if input == 1 {
		return input
	} else {
		return input * factorial(input-1)
	}
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var input int
	if scanner.Scan() {
		input, _ = strconv.Atoi(scanner.Text())
	}
	fmt.Print(factorial(input))
}
