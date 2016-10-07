package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/mitchellh/go-wordwrap"
	"github.com/nsf/termbox-go"
)

var inputBuffer = bufio.NewReader(os.Stdin)
var instructions = "Find the LU-Factorization of a matrix A\n\n" +
	"Please enter your matrix A.\n" +
	"This is accomplished by entering one row at a time, separating numbers with a space. " +
	"When finished with a row, hit the [RETURN] key. " +
	"This program takes only square matrices for A. " +
	"It will count the number of elements in row 1 and automatically " +
	"stop taking input after you have entered that many rows. " +
	"For example:\n\n row 1: 1 2 3 [RETURN] \n row 2: 4 5 6 [RETURN] \n row 3: 7 8 9 [RETURN]\n\n"

func main() {
	// Get terminal width
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	w, _ := termbox.Size()
	termbox.Close()
	width := uint(math.Min(float64(w), float64(80)))

	// Print the instructions for the user
	fmt.Print(wordwrap.WrapString(instructions, width))

	// Read the first row and calculate its size n. A will be an n x n matrix.
	fmt.Print("row 1: ")
	rowString, _ := inputBuffer.ReadString('\n')
	r1Array := strings.Fields(rowString)
	n := len(r1Array)

	// Initialize matrices A and LU
	A := make([][]float64, n)
	LU := Zeros(n)

	// Add the first row to A
	for _, elementString := range r1Array {
		if elementFloat, err := strconv.ParseFloat(elementString, 64); err == nil {
			A[0] = append(A[0], elementFloat)
		}
	}
	// Now add the rest
	for i := 1; i < n; i++ {
		fmt.Print("row ", i+1, ": ")
		rowString, _ = inputBuffer.ReadString('\n')
		if len(strings.Fields(rowString)) != n {
			panic("Invalid row! Quitting...")
		}
		for _, elementString := range strings.Fields(rowString) {
			if elementFloat, err := strconv.ParseFloat(elementString, 64); err == nil {
				A[i] = append(A[i], elementFloat)
			}
		}
	}

	fmt.Println()

	// Show the user what they entered for A
	PrintArray(A, n, "A")

	// If the calculation takes a long time, the spinner will indicate
	// that the program is still running
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " Calculating decomposition..."
	s.Start()

	// Calculate LU
	for rowIndex, rowData := range A {
		for colIndex, element := range rowData {
			var sum float64
			var newVal float64
			if rowIndex <= colIndex {
				for k := 0; k < rowIndex; k++ {
					sum += LU[rowIndex][k] * LU[k][colIndex]
				}
				newVal = element - sum
			} else {
				for k := 0; k < colIndex; k++ {
					sum += LU[rowIndex][k] * LU[k][colIndex]
				}
				newVal = (element - sum) / LU[colIndex][colIndex]
			}
			LU[rowIndex][colIndex] = newVal
		}
	}

	s.Stop()

	// Print the answer on the screen
	fmt.Println("The resulting LU-factorization is combined into a single array:\n")
	PrintArray(LU, n, "LU")
}

func Zeros(size int) [][]float64 {
	arr := make([][]float64, size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			arr[i] = append(arr[i], 0)
		}
	}
	return arr
}

func PrintArray(arr [][]float64, size int, name string) {
	for rowIndex, rowData := range arr {
		if math.Ceil(float64(size)/float64(2)) == float64(rowIndex+1) {
			fmt.Print(name, " =\t")
		} else {
			fmt.Print("\t")
		}
		for _, element := range rowData {
			fmt.Print(element, "\t")
		}
		fmt.Println()
	}
}
