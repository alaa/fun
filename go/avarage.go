package main

import "fmt"

func main() {
	set := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(avarage(set))
}

func avarage(array []int) int {
	total := 0
	for i := range array {
		total += array[i]
	}
	return total / len(array)
}
