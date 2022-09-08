package main

import "fmt"

func caesar(offset int, input string) string {
	var result string
	offset = offset % 26
	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			char += rune(offset)
			if char > 'z' {
				char -= 26
			}
		}
		result += string(char)
	}
	return result
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))

}
