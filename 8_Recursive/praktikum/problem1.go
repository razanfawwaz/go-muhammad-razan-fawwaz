package main

import (
	"fmt"
)

func primeX(number int) int {
	var result []int
	for i := 1; i < 30; i++ {
		j := 0
		for k := 1; k < 30; k++ {
			if i%k == 0 {
				j++
			}
		}
		if j == 2 && i != 1 {
			result = append(result, i)
		}
	}
	if number == 1 {
		return 2
	} else {
		return result[number-1]
	}
}
func main() {
	fmt.Println(primeX(1))  //2
	fmt.Println(primeX(5))  //11
	fmt.Println(primeX(8))  //19
	fmt.Println(primeX(9))  //23
	fmt.Println(primeX(10)) //29
}
