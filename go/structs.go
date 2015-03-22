package main

import "fmt"

func main() {
	shape := Rectangle{1, 1, 1, 1}
	fmt.Println(shape.area())
}

type Rectangle struct {
	x1, x2, x3, x4 float64
}

func (r *Rectangle) area() (area float64) {
	area = r.x1 + r.x2 + r.x3 + r.x4
	return
}
