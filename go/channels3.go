package main

import "fmt"

func main() {
	channel1 := make(chan string)

	go generator(channel1)
	go printer(channel1)

	var input string
	fmt.Scanln(&input)
}

func generator(channel chan string) {
	for i := 0; i < 10; i++ {
		channel <- "blink blink!!"
	}
}

func printer(channel chan string) {
	for {
		message := <-channel
		fmt.Println(message)
	}
}
