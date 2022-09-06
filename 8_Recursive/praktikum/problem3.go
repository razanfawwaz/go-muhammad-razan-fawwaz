package main

import "fmt"

func primaSegiEmpat(high, wide, start int) string {
	var result []int
	for i := 1; i < 50; i++ {
		j := 0
		for k := 1; k < 50; k++ {
			if i%k == 0 {
				j++
			}
		}
		if j == 2 && i != 1 {
			result = append(result, i)
		}
	}
	var o int
	for n := 0; n < len(result); n++ {
		if result[n] == start {
			o = 1 + n
		}
	}
	var total int
	for l := 0; l < wide; l++ {
		for m := 0; m < high; m++ {
			fmt.Print(result[o], " ")
			total = total + result[o]
			o++
		}
		fmt.Println()
	}
	fmt.Println(total)

	return ""
}

func main() {
	fmt.Println(primaSegiEmpat(2, 3, 13))
	fmt.Println(primaSegiEmpat(5, 2, 1))
}
