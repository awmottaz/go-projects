package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

const arraySize = 6

var maxHourglassSum = int(-9 * arraySize * arraySize)

var A [arraySize][arraySize]int

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	readInput()
	hourglasses := fetchHourglasses()
	for _, currentHourglass := range hourglasses {
		currentHourglassSum := calculateHourglassSum(currentHourglass)
		maxHourglassSum = int(math.Max(float64(maxHourglassSum), float64(currentHourglassSum)))
	}
	fmt.Print(maxHourglassSum)
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

func fetchHourglasses() [][]int {
	numHourglasses := (arraySize - 2) * (arraySize - 2)
	hourglasses := make([][]int, numHourglasses)
	counter := 0
	for i := 0; i < arraySize-2; i++ {
		for j := 0; j < arraySize-2; j++ {
			hourglasses[counter] = []int{A[i][j], A[i][j+1], A[i][j+2], A[i+1][j+1], A[i+2][j], A[i+2][j+1], A[i+2][j+2]}
			counter++
		}
	}
	return hourglasses
}

func calculateHourglassSum(hourglass []int) int {
	sum := 0
	for _, element := range hourglass {
		sum += element
	}
	return sum
}
