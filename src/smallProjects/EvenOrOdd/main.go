package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range numbers {
		if v % 2 == 0 {
			fmt.Println(v, "Even")
		} else {
			fmt.Println(v, "Odd")
		}
	}
}