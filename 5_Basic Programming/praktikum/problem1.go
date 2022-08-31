package main

import "fmt"

func main() {
	const phi = 3.14
	var T, r float64

	fmt.Print("Masukkan tinggi tabung!: ")
	fmt.Scan(&T)

	fmt.Print("Masukkan jari-jari tabung!: ")
	fmt.Scan(&r)

	var count = 2 * phi * r * (r + T)
	fmt.Println(count)
}
