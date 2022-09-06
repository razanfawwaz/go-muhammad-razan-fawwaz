package main

import (
	"fmt"
	"sort"
)

func MaximumBuyProduct(money int, productPrice []int) {
	if len(productPrice) == 0 {
		fmt.Println("0")
	}
	sort.Ints(productPrice)
	var maxProduct = 0
	for i := 0; i < len(productPrice); i++ {
		if money-productPrice[i] >= 0 {
			maxProduct++
			money -= productPrice[i]
		} else {
			break
		}
	}
	fmt.Println(maxProduct)
}

func main() {
	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})
	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000})
	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})
	MaximumBuyProduct(4000, []int{7500, 3000, 2000})
	MaximumBuyProduct(0, []int{10000, 30000})
}
