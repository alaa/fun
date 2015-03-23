package main

import "fmt"

func main() {
	alaa := Person{
		name:     "alaa qutaish",
		birthday: "15-11-1985",
	}

	student := Student{
		person:  alaa,
		collage: "CS",
		year:    4,
	}

	fmt.Println(student.collage)
}

type Person struct {
	name, birthday string
}

func (r *Person) to_s() string {
	return r.name + ", " + r.birthday
}

type Student struct {
	person  Person
	collage string
	year    int
}
