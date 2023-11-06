package main

import "fmt"

func main() {
	const name = "John"
	const age = 30
	const isCool = true
	const height = 1.75
	const weight = 70.5
	const (
		monday = iota + 1
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println(name, age, isCool, height, weight)

}
