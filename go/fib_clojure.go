package main

import "fmt"

func fibonacci() func(i int) int {
	return func(i int) int {
		f := fibonacci()
		if i <= 1 {
			return 1
		}
		return f(i-1) + f(i-2)
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 50; i++ {
		fmt.Println(f(i))
	}
}
