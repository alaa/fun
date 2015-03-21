package main

import "fmt"

func main() {
	local := 7

	add := func(x, y int) int {
		return x + y
	}

	increment := func() int {
		local++
		return local
	}

	fmt.Println(addition(1, 2, 3, 4, 5))
	fmt.Println(add(5, 5))

	fmt.Println(increment())
	fmt.Println(increment())
}

func addition(args ...int) int {
	total := 0
	for _, arg := range args {
		total += arg
	}

	return total
}
