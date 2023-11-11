package main

import "fmt"

type sayHelloCustomer struct {
	Name, Address string
	Age           int
}

func (customerHello sayHelloCustomer) sayHello() {
	fmt.Println("Hello", customerHello.Name)
}

func main() {
	chau := sayHelloCustomer{
		Name:    "Chau",
		Address: "Hanoi",
		Age:     20,
	}

	fmt.Println(chau)

}
