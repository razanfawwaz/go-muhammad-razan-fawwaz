package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	// create array merge without duplicate data from arrayA and arrayB
	var arrayMerge []string
	for _, valueA := range arrayA {
		arrayMerge = append(arrayMerge, valueA)
	}
	for _, valueB := range arrayB {
		var isDuplicate = false
		for _, valueA := range arrayA {
			if valueB == valueA {
				isDuplicate = true
			}
		}
		if !isDuplicate {
			arrayMerge = append(arrayMerge, valueB)
		}
	}
	return arrayMerge
}

func main() {
	fmt.Println(ArrayMerge([]string{"steve", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
}
