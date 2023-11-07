package main

import "fmt"

func main() {
	name := "Chau Batkunda"

	// if statement
	if name == "Chau Batkunda" {
		fmt.Println("Hello Chau Batkunda")
	} else if name == "Julio Batkunda" {
		fmt.Println("Hello Julio Batkunda")
	} else {
		fmt.Println("Hello Stranger")
	}
	// if short statement
	if length := len(name); length > 5 {
		fmt.Println(length)
	}

}
