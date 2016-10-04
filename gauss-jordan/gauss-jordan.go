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

var in = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("Find the L-U decomposition of a matrix A.\n")

	// Get terminal width
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	w, _ := termbox.Size()
	termbox.Close()
	width := uint(math.Min(float64(w), float64(80)))

	instructions := "Please enter your matrix A.\n" +
		"This is accomplished by entering one row at a time, separating numbers with a space. " +
		"When finished with a row, hit the [RETURN] key. " +
		"This program takes only square matrices for A. " +
		"It will count the number of elements in row 1 and automatically " +
		"stop taking input after you have entered that many rows. " +
		"For example:\n\n row 1: 1 2 3 [RETURN] \n row 2: 4 5 6 [RETURN] \n row 3: 7 8 9 [RETURN]\n\n"
	fmt.Print(wordwrap.WrapString(instructions, width))

	fmt.Print("row 1: ")
	row, _ := in.ReadString('\n')
	r1 := strings.Fields(row)
	n := len(r1)

	// Initialize matrix A
	A := make([][]float64, n)
	b := make([]float64, n)
	x := make([]float64, n)
	y := make([]float64, n)
	LU := Zeros(n)

	// Add the first row
	for _, num := range r1 {
		if val, e := strconv.ParseFloat(num, 64); e == nil {
			A[0] = append(A[0], val)
		}
	}
	// Now add the rest
	for i := 1; i < n; i++ {
		fmt.Print("row ")
		fmt.Print(i + 1)
		fmt.Print(": ")
		row, _ = in.ReadString('\n')
		if len(strings.Fields(row)) != n {
			panic("Invalid row! Quitting...")
		}
		for _, num := range strings.Fields(row) {
			if val, e := strconv.ParseFloat(num, 64); e == nil {
				A[i] = append(A[i], val)
			}
		}
	}

	fmt.Print("\n\nNow enter your vector b.\nb = ")

	temp, _ := in.ReadString('\n')
	if len(strings.Fields(temp)) != n {
		panic("b has the wrong size!")
	}
	for i, num := range strings.Fields(temp) {
		if val, e := strconv.ParseFloat(num, 64); e == nil {
			b[i] = val
		}
	}

	// Show the user what they entered for A
	PrintArray(A, n, "A") // PrintArray(arr, size, name)
	fmt.Print("b = ", b, "\n")

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " Calculating decomposition..."
	s.Start()

	for row, data := range A {
		for col, elem := range data {
			var sum float64
			var newVal float64
			if row <= col {
				for i := 0; i < row; i++ {
					sum += LU[row][i] * LU[i][col]
				}
				newVal = elem - sum
			} else {
				for i := 0; i < col; i++ {
					sum += LU[row][i] * LU[i][col]
				}
				newVal = (elem - sum) / LU[col][col]
			}
			LU[row][col] = newVal
		}
	}

	// Solve Ly = b for y
	for i := 0; i < n; i++ {
		var sum float64
		sum = 0.0
		for j := 0; j < i; j++ {
			sum = sum + LU[i][j]*b[j]
		}
		y[i] = b[i] - sum
	}
	// Solve Ux = y for x
	for i := n - 1; i >= 0; i-- {
		var sum float64
		sum = 0.0
		for j := n - 1; j > i; j-- {
			sum = sum + LU[i][j]*x[j]
		}
		x[i] = (y[i] - sum) / LU[i][i]
	}
	s.Stop()
	fmt.Print("x: ", x)
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
	for ind, mrow := range arr {
		if math.Ceil(float64(size)/float64(2)) == float64(ind+1) {
			fmt.Print(name, " =\t")
		} else {
			fmt.Print("\t")
		}
		for _, elem := range mrow {
			fmt.Print(elem, "\t")
		}
		fmt.Println()
	}
}
