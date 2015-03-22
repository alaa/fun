package main

import "fmt"

func main() {
	c := make(chan string, 2)

	c <- "buffer1"
	c <- "buff2"

	fmt.Println(<-c)
	fmt.Println(<-c)
}
