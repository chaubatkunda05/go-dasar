package main

func getFullName2() (firstName, lastName string) {
	firstName = "Chau"
	lastName = "Batkunda"
	return
}

func main() {
	f, l := getFullName2()
	println(f, l)

}
