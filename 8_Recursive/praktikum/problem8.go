package main

import (
	"fmt"
)

type pair struct {
	items string
	count int
}

func MostAppearItem(items []string) []pair {
	if len(items) == 0 {
		return []pair{}
	}
	var result []pair
	for i := 0; i < len(items); i++ {
		var count = 0
		for j := 0; j < len(items); j++ {
			if items[i] == items[j] {
				count++
			}
		}
		result = append(result, pair{items[i], count})
		//remove duplicates
		for k := 0; k < len(result); k++ {
			for l := k + 1; l < len(result); l++ {
				if result[k].items == result[l].items {
					result = append(result[:l], result[l+1:]...)
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"})) //[{a 3} {b 2}]
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))

}
