package main

import (
	"fmt"
)

func msort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	needle := len(array) / 2
	left := msort(array[0:needle])
	right := msort(array[needle:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	set := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(set, right...)
		}

		if len(right) == 0 {
			return append(set, left...)
		}

		if left[0] <= right[0] {
			set = append(set, left[0])
			left = left[1:]

		} else {
			set = append(set, right[0])
			right = right[1:]
		}
	}
	return set
}

func main() {
	unsorted_list := []int{56, 6, 6, 4, 3, 3, 4, 65, 7, 1, 4, 3, 4, 5, 6, 7}
	fmt.Printf("unsorted list %v \n", unsorted_list)
	fmt.Printf("sorted list %v \n", msort(unsorted_list))
}
