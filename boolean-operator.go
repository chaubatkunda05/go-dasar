package main

import "fmt"

func main() {
	var (
		ujian        = 80
		nilaiAbsensi = 80

		lulusUjian = ujian >= 80
		lulusAbsen = nilaiAbsensi >= 80

		lulus = lulusUjian && lulusAbsen
	)

	fmt.Println(lulus)
}
