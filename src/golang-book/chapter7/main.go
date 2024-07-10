package main

import (
	"fmt"
	"math"
)


type Rect struct {
	x, y, z float64
}
type Square struct {
	x float64
}
type Shape interface {
	capacity() float64
}
type MultiShape struct {
	shapes []Shape
}

func main() {

	multishape := MultiShape{
		shapes: []Shape{
			&Rect{x:3,y:4,z:2},
			&Square{x: 5},
		},
	}
	// define a shapes: rectangle and square
	
	printCapacity(multishape.shapes...)
}
// Define a method to structs
func (rect *Rect) capacity() float64 {
	return rect.x * rect.y * rect.z
}
func (square *Square) capacity() float64 {
	return math.Sqrt(square.x) * square.x
}

func printCapacity(shapes ...Shape) {
	for _, s := range shapes {
		fmt.Println(s.capacity())
	}
}