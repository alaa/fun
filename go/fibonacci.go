package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(fib(i))
	}
}

func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	x := fib(n-1) + fib(n-2)
	return x
}
