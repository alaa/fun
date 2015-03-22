package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	go func() { c <- "ping" }()

	msg := <-c
	fmt.Println(msg)
}
