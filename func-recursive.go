package main

import "fmt"

func factorialLoop(n int) int {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	return f
}

func factorialRecursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorialRecursive(n-1)
}

func main() {
	loop := factorialLoop(5)
	fmt.Println(loop)

	recursive := factorialRecursive(5)
	fmt.Println(recursive)
}
