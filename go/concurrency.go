package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 20; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}

func f(f int) {
	fmt.Println(f)
	time.Sleep(time.Millisecond * 500)
}
