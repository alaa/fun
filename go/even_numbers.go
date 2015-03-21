package main

import "fmt"

func main() {
	hops := 10
	nextEven := evenGenerator()
	for i := 0; i <= hops; i++ {
		fmt.Println(nextEven())
	}
}

func evenGenerator() func() int {
	i := 0
	return func() (ret int) {
		ret = i
		i += 2
		return ret
	}
}
