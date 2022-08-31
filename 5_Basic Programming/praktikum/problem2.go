package main

import "fmt"

func main() {
	var nilai int
	var nama string

	fmt.Print("Masukkan nama siswa: ")
	fmt.Scan(&nama)

	fmt.Print("Masukkan nilai siswa: ")
	fmt.Scan(&nilai)

	if nilai >= 80 && nilai <= 100 {
		fmt.Println("Nama siswa: ", nama, "Nilai siswa: ", nilai, "Grade: A")
	} else if nilai >= 65 && nilai <= 79 {
		fmt.Println("Nama siswa: ", nama, "Nilai siswa: ", nilai, "Grade: B")
	} else if nilai >= 50 && nilai <= 64 {
		fmt.Println("Nama siswa: ", nama, "Nilai siswa: ", nilai, "Grade: C")
	} else if nilai >= 35 && nilai <= 49 {
		fmt.Println("Nama siswa: ", nama, "Nilai siswa: ", nilai, "Grade: D")
	} else if nilai == 0 && nilai <= 34 {
		fmt.Println("Nama siswa: ", nama, "Nilai siswa: ", nilai, "Grade: E")
	} else {
		fmt.Println("Nilai Invalid!")
	}

}
