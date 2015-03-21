package main

import "fmt"

func main() {
	array := [5]float64{1, 2, 3, 4, 5}

	var total float64 = 0
	x := array[0:2]

	for _, value := range array {
		total += value + x[1]
	}

	fmt.Println(total / float64(len(array)))
}
