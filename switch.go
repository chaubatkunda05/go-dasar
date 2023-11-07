package main

import "fmt"

func main() {
	name := "Julio" // change this to test

	switch name {
	case "Chau Batkunda":
		fmt.Println("Welcome Chau Batkunda")
	case "Julio":
		fmt.Println("Welcome Julio")
	case "Deya":
		fmt.Println("Welcome Deya")
	default:
		fmt.Println("You are not welcome here")
	} // end switch

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Your name is long")
	case false:
		fmt.Println("Your name is short")
	}
}
