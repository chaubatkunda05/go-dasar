package main

func main() {
	counter := 1
	for counter <= 10 {
		println("Perulangan ke ", counter)
		counter++
	}
	slice := []string{"Rizky", "Khapidsyah", "Go", "Lang"}
	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}
}
