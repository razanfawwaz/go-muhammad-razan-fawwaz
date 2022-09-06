package main

import "fmt"

func fibbonaci(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return fibbonaci(n-1) + fibbonaci(n-2)
	}
}

func main() {
	fmt.Println(fibbonaci(0))  //1
	fmt.Println(fibbonaci(2))  //1
	fmt.Println(fibbonaci(9))  //2
	fmt.Println(fibbonaci(10)) //3
	fmt.Println(fibbonaci(12)) //5
}
