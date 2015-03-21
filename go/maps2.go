package main

import "fmt"

func main() {

	people := map[string]map[string]string{
		"alaa": map[string]string{
			"name":       "Alaa Qutaish",
			"occupation": "Software Engineer",
		},
		"ali": map[string]string{
			"name":       "Ali Qutaish",
			"occupation": "Business manager",
		},
	}

	if person, ok := people["alaa"]; ok {
		fmt.Println(
			"Full name:", person["name"],
			"\n",
			"Occupation:", person["occupation"],
		)
	}

}
