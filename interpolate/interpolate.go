package main

import (
	"fmt"
	"math"
)

func main() {
	interpolate(0, float64(0), float64(1), 4, int(math.Pow(10, 6)))
	interpolate(0, float64(0), float64(1), 10, int(math.Pow(10, 6)))
	interpolate(1, float64(0), float64(1), 4, int(math.Pow(10, 6)))
	interpolate(1, float64(0), float64(1), 10, int(math.Pow(10, 6)))
}

func interpolate(functionFlag int, leftEndPoint, rightEndPoint float64, numPoints, numErrorSamples int) {
	xVals, yVals := generateInitialConditions(functionFlag, leftEndPoint, rightEndPoint, numPoints)
	coeffs := yVals
	for j := 1; j <= numPoints; j++ {
		for i := numPoints; i >= j; i-- {
			coeffs[i] = (coeffs[i] - coeffs[i-1]) / (xVals[i] - xVals[i-j])
		}
	}
	fmt.Print("p(x) = ")
	printPolynomial(xVals, coeffs)
	fmt.Print("\nError = ", calculateError(functionFlag, xVals, coeffs, leftEndPoint, rightEndPoint, numErrorSamples), "\n\n")
}

func generateInitialConditions(functionFlag int, leftEndPoint, rightEndPoint float64, numPoints int) ([]float64, []float64) {
	var deltaX = (rightEndPoint - leftEndPoint) / float64(numPoints)
	var xVals = make([]float64, numPoints+1)
	var yVals = make([]float64, numPoints+1)
	for i := 0; i <= numPoints; i++ {
		xVals[i] = leftEndPoint + (float64(i) * deltaX)
		yVals[i] = evaluateTestFunction(functionFlag, xVals[i])
	}
	return xVals, yVals
}

func evaluateTestFunction(functionFlag int, point float64) float64 {
	switch functionFlag {
	case 0:
		return math.Exp(point)
	case 1:
		return math.Sin(float64(20) * point)
	default:
		return float64(0)
	}
}

func evaluatePolynomial(xVals, coeffs []float64, point float64) float64 {
	var degree = len(coeffs) - 1
	var result = float64(0)
	for i := 0; i <= degree; i++ {
		var product = float64(1)
		for j := 0; j < i; j++ {
			product *= point - xVals[j]
		}
		result += coeffs[i] * product
	}
	return result
}

func printPolynomial(xVals, coeffs []float64) {
	var degree = len(coeffs) - 1
	fmt.Printf("%f", coeffs[0])
	for i := 1; i <= degree; i++ {
		if coeffs[i] < 0 {
			fmt.Print(" - ")
		} else {
			fmt.Print(" + ")
		}
		fmt.Print(math.Abs(coeffs[i]))
		for j := 0; j < i; j++ {
			fmt.Print("(x - ")
			fmt.Printf("%f", xVals[j])
			fmt.Print(")")
		}
	}
}

func calculateError(functionFlag int, xVals, coeffs []float64, leftEndPoint, rightEndPoint float64, samples int) float64 {
	var deltaX = (rightEndPoint - leftEndPoint) / float64(samples)
	var error = float64(0)
	for i := 0; i <= samples; i++ {
		point := leftEndPoint + (float64(i)*deltaX)
		absoluteDiff := math.Abs(evaluateTestFunction(functionFlag, point) - evaluatePolynomial(xVals, coeffs, point))
		error = math.Max(error, absoluteDiff)
	}
	return error
}
