package main

import (
	"fmt"
	"main/math"
)

func main() {
	xs := []float64{1, 2, 3, 4, 5, 6}
	avg := math.Average(xs)
	fmt.Println(avg)
}