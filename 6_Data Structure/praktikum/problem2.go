package main

import (
	"fmt"
)

func munculSekali(num string) []int {
	// create array of int from num, if the data is duplicate, then remove it only shows the not duplicate data
	for _, value := range num {
		var data []string
		data = append(data, string(value))
	}

	return nil
}

func main() {
	fmt.Println(munculSekali("1234123"))
}
