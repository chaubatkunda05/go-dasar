package main

func getGoodBy(name string) string {
	return "Good bye " + name
}

func main() {
	goodBy := getGoodBy
	println(goodBy("Golang"))
}
