package main

import "fmt"

func primeNumber(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}

	}
	return true
}

func main() {
	fmt.Println(primeNumber(1000000007))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))
}
