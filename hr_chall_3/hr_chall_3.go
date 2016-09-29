package main

import (
	"fmt"
	"math"
)

var mealCost, tipAmount, taxAmount float64

func Round(val float64) (newVal int) {
	var round float64
	pow := math.Pow(10, float64(0))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= 0.5 {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = int(round / pow)
	return
}

func main() {
	_, _ = fmt.Scanf("%f", &mealCost)
	_, _ = fmt.Scanf("%f", &tipAmount)
	_, _ = fmt.Scanf("%f", &taxAmount)

	tip := mealCost * (tipAmount / 100.0)
	tax := mealCost * (taxAmount / 100.0)
	totalCost := mealCost + tip + tax
	fmt.Println(tip, tax, mealCost)
	fmt.Printf("The total meal cost is %d dollars", Round(totalCost))
}
