package main

import "fmt"

func cetakTabelPerkalian(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			if j*i < 10 {
				fmt.Print("   ", j*i)
			} else {
				fmt.Print("  ", j*i)
			}
		}
		fmt.Println()
	}
}

func main() {
	cetakTabelPerkalian(9)
}
