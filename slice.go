package main

import "fmt"

func main() {
	var months = [...]string{
		1:  "January", // 0 is empty
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December", // 13 is empty
	} // array
	var summer = months[6:9] // [June July August]

	fmt.Println(summer)
	fmt.Println(len(summer)) // 3
	fmt.Println(cap(summer)) // 7

	var slice2 = months[10:] // [October November December]
	fmt.Println(slice2)      // [October November December]
}
