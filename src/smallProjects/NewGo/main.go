package main

import "fmt"

type Circle struct {
    x string
    y float64
    r float64
}
func main() {

	c := Circle{x: "cm", y: 3.4, r: 5.5}

	fmt.Println(c)

}