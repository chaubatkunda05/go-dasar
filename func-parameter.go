package main

func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	println("Hello", nameFiltered)
} // func sayHelloWithFilter

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
} // func spamFilter

func main() {
	sayHelloWithFilter("Chau", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}
