package main

import "fmt"

func primeNumber(number int) bool {
	var j = 0
	for i := 1; i <= number; i++ {
		if number%i == 0 && number%number == 0 {
			j += 1
		}
	}
	if j == 2 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(primeNumber(11))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))
}
