package main

import "fmt"

type ops struct {
	op string
	fn func(x, y int) int
}

func main() {
	addition := func(x, y int) int { return x + y }
	substraction := func(x, y int) int { return x - y }
	division := func(x, y int) int { return x / y }
	multiplication := func(x, y int) int { return x * y }

	cal := []ops{
		{"addition", addition},
		{"substraction", substraction},
		{"division", division},
		{"multiplication", multiplication},
	}

	for _, x := range cal {
		fmt.Println(x.op, "=>", x.fn(1, 1))
	}
}
