package main

import "fmt"

func main() {

	numberArr := [5]float64{3, 5, 3, 4, 10}

	
	fmt.Println(sum(numberArr[:]))
	fmt.Println(isEven(int(numberArr[2])))
	fmt.Println(findGreatest(1, 53 ,34 ,123 ,123 ,1 ,231 ,33))
	
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	
	x := 1.5
	square(&x)
	fmt.Println(x)

	firstVal := 55
	secondVal := 33

	fmt.Println("First: ", firstVal, "Second: ", secondVal)
	swapValues(&firstVal, &secondVal)
	fmt.Println("First: ", firstVal, "Second: ", secondVal)
}

func sum(numbers []float64) float64 {
	total := 0.0
	for _, v := range numbers {
		total += v
	}
	return total
}

func isEven(number int) bool {
	if(number % 2 == 0) {
		return true
	} else {
		return false
	}
}

func findGreatest(args ...int) int {
	var biggestNumber int;
	for _, v := range args {
		if biggestNumber < v {
			biggestNumber = v
		}
	}
	return biggestNumber
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i+= 2
		return
	}
}

func square(x *float64) {
	*x = *x * *x
}

func swapValues(firstVal *int, secondVal *int ) {
	tempVal := int(*firstVal)
	*firstVal = int(*secondVal)
	*secondVal = tempVal
}