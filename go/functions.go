package main

import "fmt"

func main() {
	arr := []float64{1, 2, 3, 4, 5}
	fmt.Println(avarage(arr))
}

func avarage(arr []float64) float64 {
	total := 0.0

	for _, i := range arr {
		total += i
	}

	return (total / float64(len(arr)))
}
