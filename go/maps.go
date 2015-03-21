package main

import "fmt"

func main() {
	people := map[string]string{
		"name": "alaa",
		"age":  "25",
	}

	fmt.Println(people["name"])

	if name, ok := people["name"]; ok {
		fmt.Println(name, ok)
	} else {
		fmt.Println("hash key does not exsist!")
	}

	delete(people, "name")
	delete(people, "age")
	fmt.Println("Crowd size is:", len(people))
}
