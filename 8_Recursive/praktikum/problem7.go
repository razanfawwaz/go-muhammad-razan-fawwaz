package main

import "fmt"

func playingDomino(cards [][]int, deck []int) interface{} {
	if len(cards) == 0 {
		return "tutup kartu"
	}

	for i := 0; i < len(cards); i++ {
		var total = 0
		var currentTotal = 0
		if cards[i][0] == deck[0] || cards[i][1] == deck[0] || cards[i][0] == deck[1] || cards[i][1] == deck[1] {
			currentTotal = cards[i][0] + cards[i][1]
			if currentTotal < total {
				total = currentTotal
			} else {
				return cards[i]
			}
		}
	}
	return "tutup kartu"
}
func main() {
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}}, []int{4, 3})) //[3,4]
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}}, []int{3, 6})) //[6,5]
	fmt.Println(playingDomino([][]int{{6, 6}, {2, 4}, {3, 6}}, []int{5, 1}))         //tutup kartu
}
