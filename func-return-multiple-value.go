package main

func getFullName() (string, string) {
	return "Chau", "Batkunda"
}

func main() {
	firstName, lastName := getFullName()
	println(firstName, lastName)

	firstName, _ = getFullName()
	println(firstName)
}
