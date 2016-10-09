package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func countMaxConsecutiveOnes(binarySlice []string) int {
	maxCountFound := 0
	currentCount := 0
	for _, digit := range binarySlice {
		if digit == "1" {
			currentCount += 1
		} else {
			maxCountFound = int(math.Max(float64(maxCountFound), float64(currentCount)))
			currentCount = 0
		}
	}
	maxCountFound = int(math.Max(float64(maxCountFound), float64(currentCount)))
	return maxCountFound
}

func decimalToBinary(number int) []string {
	binaryRepresentation := make([]string, 21)
	for power := 20; power >= 0; power-- {
		diff := number - int(math.Pow(float64(2), float64(power)))
		switch {
		case diff >= 0:
			binaryRepresentation[power] = "1"
			number = diff
		case diff < 0:
			binaryRepresentation[power] = "0"
		}
	}
	return binaryRepresentation
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var input int
	if scanner.Scan() {
		input, _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(decimalToBinary(input))
	fmt.Println(countMaxConsecutiveOnes(decimalToBinary(input)))
}
