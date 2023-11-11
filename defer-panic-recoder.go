package main

import "fmt"

func longging() {
	fmt.Println("Selesai Memanggul function")

}

func runApplication(value int) {
	defer longging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	runApplication(1)
}
