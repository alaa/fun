package main

import "fmt"

func main() {
	var input string

	fmt.Println("what is your name? ")

	fmt.Scanln(&input)
	fmt.Println("your name is: ", input)
}
