package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, -10}
	var min int = values[0]

	for _, i := range values {
		if i < min {
			min = i
		}
	}

	fmt.Println("smallest value is:", min)
}
