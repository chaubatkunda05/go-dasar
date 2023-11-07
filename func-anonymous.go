package main

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		println("You are blocked", name)
	} else {
		println("Welcome", name)
	}
}

//func blackList1(name string) bool {
//	return name == "admin"
//}

func main() {

	blackList := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blackList)
	registerUser("Chau", blackList)

}
