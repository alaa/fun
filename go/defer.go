package main

import "fmt"

func main() {
	defer second()
	first()
}

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}
