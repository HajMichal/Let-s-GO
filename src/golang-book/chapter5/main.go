package main

import "fmt"

func main() {
	
	var x [5]int
	y := [5]int{1, 23, 35 ,12 ,44}

	x[4] = 100

	fmt.Println(x)

	for i, value := range y {
		fmt.Println(value)
		fmt.Println("index: ", i)
	}

	slice := make([]int, 3, 9)
	fmt.Println(len(slice))

	givenSlice := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(givenSlice[2:5])
	
	arr := []int{
		48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17,
	}
	
	lowestNum := arr[0]

	for j := 0; j < len(arr); j++ {
		if arr[j] < lowestNum {
			lowestNum = arr[j]
		}
	}
	fmt.Println(lowestNum)
}