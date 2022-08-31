package main

import "fmt"

func playWithAsterik(n int) {
	for i := 0; i < n; i++ {
		for k := 0; k <= n-i; k++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}

		fmt.Println()
	}
}

func main() {
	playWithAsterik(5)
}
