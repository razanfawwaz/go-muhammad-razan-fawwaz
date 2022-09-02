package main

import (
	"fmt"
	"strconv"
)

func munculSekali(num string) []int {
	// create array of int from num, if the data is duplicate, then remove it only shows the not duplicate data
	var arrayNum []int
	for _, value := range num {
		var intValue, _ = strconv.Atoi(string(value))
		arrayNum = append(arrayNum, intValue)
	}
}

func main() {
	fmt.Println(munculSekali("1234123"))
}
