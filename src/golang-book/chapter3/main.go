package main

import "fmt"

func main() {
	var x string = "Hello There" 
	fastAssign := "Or we can assign variable like this because compiler will automaticly infer type of this variable"
	
	fmt.Println(x)
	fmt.Println(fastAssign)

	defineUserInput()
}

func defineMultipleVars() {
	const (
		a = 5
		b = 10
		c = 15
	)
}

func defineUserInput() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)
}