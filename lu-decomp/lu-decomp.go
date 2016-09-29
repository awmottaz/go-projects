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

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " Calculating decomposition..."
	s.Start()

	for row, data := range A {
		for col, elem := range data {
			// Insert code here
		}
	}

	s.Stop()
}
