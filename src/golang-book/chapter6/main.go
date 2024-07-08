package main

import "fmt"

func main() {
	
	arr := []int{
		48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17,
	}
	
	lowestNum := arr[0]

	for _, v := range arr {
		if v < lowestNum {
			lowestNum = v
		}
	}
	fmt.Println(lowestNum)

	fmt.Println("Average of given arr: ", average(arr))
}

func average(arr []int) float64{
	total := 0.0
	for _, v := range arr {
		total += float64(v)
	}
	return total / float64(len(arr))
}