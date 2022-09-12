package main

import (
	"strings"
	"time"
)

func main() {
	var input = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolor magna aliqua"
	split := strings.Split(input, " ")

	for _, word := range split {
		go func(word string) {
			var letterFrequency = make(map[string]int)
			for _, letter := range word {
				letterFrequency[string(letter)]++
			}
			for letter, frequency := range letterFrequency {
				println(letter, frequency)
			}
		}(word)
		time.Sleep(2 * time.Second)
	}
}
