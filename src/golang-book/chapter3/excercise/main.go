package main

import "fmt"

func main() {

	fmt.Println("What you want to convert:")
	fmt.Println("[1] - Temperature")
	fmt.Println("[2] - Measure")

	var userChoice int
	fmt.Scanf("%d", &userChoice)

	if userChoice == 1 {
		farenheitToCelcius()
	} else if userChoice == 2 {
		feetToMeters()
	} else {
		main()
	}

}

func farenheitToCelcius() {
	fmt.Print("Enter a Temperature in Farenheit: ")

	var farenheitTemp float64
	fmt.Scanf("%f", &farenheitTemp)

	convertToCelcius := (farenheitTemp - 32) * 5/9

	fmt.Print("Given temperature converted to Celcius is: ", convertToCelcius, " C")
}

func feetToMeters() {
	feetToMeterScale := 0.3048
	fmt.Print("Enter a measure in feets: ")

	var feets float64
	fmt.Scanf("%f", &feets)

	convertToMeters := feetToMeterScale * feets

	fmt.Print("Given measure converted to meters is: ", convertToMeters, " meters")
}