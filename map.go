package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Chau batkunda",
		"address": "Malang",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
}
