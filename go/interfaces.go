package main

import "fmt"

type Shape interface {
	area() string
}

type Circle struct {
}

type Rectangle struct {
}

func (r Circle) area() string {
	return "circle area"
}

func (r Rectangle) area() string {
	return "rectangle area"
}

func main() {
	shapes := []Shape{Circle{}, Rectangle{}}

	for _, shape := range shapes {
		fmt.Println(shape.area())
	}
}
