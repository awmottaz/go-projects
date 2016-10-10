package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)
var maxHourglassSum int

const arraySize = 6

var A [arraySize][arraySize]int

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	readInput()
	hourglasses := fetchHourglasses()
	fmt.Print(hourglasses)
	/*
	for _, currentHourglass := range(hourglasses) {
		maxHourglassSum = int( math.Max( float64(maxHourglassSum), float64(calculateHourglassSum(currentHourglass)) ) )
	}
	*/
}

func readInput() {
	for rowIndex := 0; rowIndex < arraySize; rowIndex++ {
		stringArray := make([]string, arraySize)
		if scanner.Scan() {
			stringArray = strings.Fields(scanner.Text())
		}
		for colIndex, entry := range stringArray {
			A[rowIndex][colIndex], _ = strconv.Atoi(entry)
		}
	}
}

func fetchHourglasses() [][][]int {
	hourglasses := make([][][]int, arraySize)
	for i := 0; i < arraySize - 2; i++ {
		for j := 0; j < arraySize - 2; j++ {
			hourglasses[i][j] = [ A[i][j] A[i+1][j] A[i+2][j] A[i+1][j+1] A[i][j+2] A[i+1][j+2] A[i+2][j+2] ]
		}
	}
	return hourglasses
}
