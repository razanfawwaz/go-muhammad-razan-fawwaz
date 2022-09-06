package main

import (
	"fmt"
)

type pair struct {
	items string
	count int
}

func munculSekali(num string) []int {
	// create array of int from num, if the data is duplicate, then remove it only shows the not duplicate data
	var data []string
	for _, value := range num {
		data = append(data, string(value))

	}
	var result []pair
	for i := 0; i < len(data); i++ {
		var count = 0
		for j := 0; j < len(data); j++ {
			if data[i] == data[j] {
				count++
			}
		}
		result = append(result, pair{data[i], count})
		for k := 0; k < len(result); k++ {
			for l := k + 1; l < len(result); l++ {
				if result[k].items == result[l].items {
					result = append(result[:l], result[l+1:]...)
				}
			}
			if result[k].count != 1 {
				result = append(result[:k], result[k+1:]...)
			}
		}
	}
	fmt.Println(result)
	return nil
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
