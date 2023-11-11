package main

import "fmt"

func main() {
	name := "Chau Batkunda"
	counter := 0
	increment := func() {
		name := "Esau Batkunda"
		counter++
		fmt.Println(name)
	}
	increment()
	increment()
	println(counter)
	println(name)

}
