package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/mitchellh/go-wordwrap"
	"github.com/nsf/termbox-go"
)

var in = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("Compute a solution to Ax = b.\n")

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

	// Initialize matrix A and vector b
	A := make([][]float64, n)
	b := make([]float64, n)

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

	fmt.Println()

	// Show the user what they entered for A
	for ind, mrow := range A {
		if math.Ceil(float64(n)/float64(2)) == float64(ind+1) {
			fmt.Print("A =\t")
		} else {
			fmt.Print("\t")
		}
		for _, elem := range mrow {
			fmt.Print(elem, "\t")
		}
		fmt.Println()
	}

	fmt.Print("\n Now enter your matrix b as a row vector.\n\n")

	fmt.Print("b = ")
	bin, _ := in.ReadString('\n')
	binarr := strings.Fields(bin)
	if len(binarr) != n {
		panic("Length doesn't match! Quitting...")
	}

	for i := 0; i < n; i++ {
		if val, e := strconv.ParseFloat(binarr[i], 64); e == nil {
			b[i] = val
		}
	}

	MakeMultiplierArray(A, n)
}

func MakeMultiplierArray(mat, size float64) [][]float64 {

}
