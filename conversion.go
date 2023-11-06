package main

func main() {
	var a int = 10
	var b float32 = 10.5
	var c float64 = 10.5
	var d byte = 10
	var e rune = 10
	var f int32 = 10
	var g int64 = 10

	var name string = "Hello World"
	var n byte = name[0]
	var s string = string(n)
	println(s)

	// int to float
	println(float32(a))
	println(float64(a))

	// float to int
	println(int(b))
	println(int(c))

	// byte to int
	println(int(d))

	// rune to int
	println(int(e))

	// int32 to int
	println(int(f))

	// int64 to int
	println(int(g))
}
