package main

import "fmt"

func FindMinAndMax(arr []int) string {
	if len(arr) == 0 {
		return "[]"
	}
	min := arr[0]
	max := arr[0]
	minIndex := arr[0]
	maxIndex := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
			minIndex = i
		}
		if arr[i] > max {
			max = arr[i]
			maxIndex = i
		}
	}

	return fmt.Sprintf("min: %d index: %d, max: %d index: %d", min, minIndex, max, maxIndex)
}
func main() {
	fmt.Println(FindMinAndMax([]int{5, 7, 4, -2, -1, 8}))
	fmt.Println(FindMinAndMax([]int{2, -5, -4, 22, 7, 7}))
	fmt.Println(FindMinAndMax([]int{4, 3, 9, 4, -21, 7}))
	fmt.Println(FindMinAndMax([]int{-1, 5, 6, 4, 2, 18}))
	fmt.Println(FindMinAndMax([]int{-2, 5, -7, 4, 7, -20}))
}
