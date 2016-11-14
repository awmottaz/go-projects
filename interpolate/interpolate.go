package main

import (
	"fmt"
	"math"
)

func generateInitialConditions(leftEndPoint float64, rightEndPoint float64, numPoints int) ([]float64, []float64) {
	var deltaX = (rightEndPoint - leftEndPoint) / float64(numPoints)
	var xVals = make([]float64, numPoints+1)
	var yVals = make([]float64, numPoints+1)
	for i := 0; i <= numPoints; i++ {
		xVals[i] = leftEndPoint + float64(i)*deltaX
		yVals[i] = math.Sin(20 * xVals[i]) // Change test function here
	}
	return xVals, yVals
}

func printPolynomial(coeff []float64) {
	var degree = len(coeff) - 1
	fmt.Print(coeff[0])
	for i := 1; i <= degree; i++ {
		fmt.Print(" + ")
		fmt.Printf("%f", coeff[i])
		fmt.Print("x^", i)
	}
}

func main() {
	leftEndPoint := float64(0)
	rightEndPoint := float64(1)
	numPoints := int(6)
	xVals, yVals := generateInitialConditions(leftEndPoint, rightEndPoint, numPoints)
	// We follow the algorithm outlined on page 332
	coeff := yVals
	for j := 1; j <= numPoints; j++ {
		for i := numPoints; i >= j; i-- {
			coeff[i] = (coeff[i] - coeff[i-1]) / (xVals[i] - xVals[i-j])
		}
	}
	//fmt.Print("p(x) = ")
	printPolynomial(coeff)
}
