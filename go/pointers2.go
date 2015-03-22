package main

import "fmt"

func main() {
	ptr_value := new(int)
	one(ptr_value)
	fmt.Println(*ptr_value)

}

func one(ptr *int) {
	*ptr = 1
}
