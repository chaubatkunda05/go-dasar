package main

import "fmt"

func main() {
	var names [2]string
	names[0] = "Chau"
	names[1] = "Batkunda"
	//names[2] = "Batkunda Chau" // error: index out of range [2] with length 2

	fmt.Println(names[0], names[1])
	var values = [3]int{
		20,
		40,
		70,
	}
	fmt.Println(values)      // [20 40 70]
	fmt.Println(len(values)) // length of array
	fmt.Println(len(names))  // length of array
}
