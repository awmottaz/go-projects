/*

This program is a minimal example to show how to pass functions as variables to other functions

*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var val = evaluatePassedMathFunctionAtPoint(math.Exp, float64(3))
	fmt.Printf("%f", val)
}

func evaluatePassedMathFunctionAtPoint(f func(float64) float64, point float64) float64 {
	return f(point)
}