package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {

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
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}
