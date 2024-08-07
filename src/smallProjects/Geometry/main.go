package main

import "fmt"

type shape interface {
	getArea() float64
} 

type triangle struct {
	height float64
	base float64
}
type square struct {
	sideLength float64
}

func main() {
	sq := square{sideLength: 5}
	tr := triangle{height: 5, base: 3}

	printArea(sq)
	printArea(tr)

}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func(s square) getArea() float64 {
	return s.sideLength * s.sideLength 
}

func(t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}