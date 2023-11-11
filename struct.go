package main

import "fmt"

type CustomerPremium struct {
	Name, Address string
	Age           int
}

func main() {
	chau := CustomerPremium{
		Name:    "Chau Batkunda",
		Address: "Malang",
		Age:     30,
	}

	fmt.Println(chau)

}
