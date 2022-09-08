package main

import "fmt"

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""
	for _, char := range s.name {
		if char >= 'a' && char <= 'z' {
			char += rune(1)
			if char > 'z' {
				char -= 26
			}
		}
		nameEncode += string(char)
	}
	s.nameEncode = nameEncode
	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""
	for _, char := range s.nameEncode {
		if char >= 'a' && char <= 'z' {
			char -= rune(1)
			if char > 'z' {
				char += 26
			}
		}
		nameDecode += string(char)
	}
	return nameDecode
}

func main() {
	var menu int
	var a = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose Menu: ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student's Name: ")
		fmt.Scan(&a.name)
		fmt.Printf("\nEncode of Student's Name %s is: %s", a.name, c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student's Encode Name: ")
		fmt.Scan(&a.nameEncode)
		fmt.Printf("\nDecode of Student's Name %s is: %s", a.nameEncode, c.Decode())
	} else {
		fmt.Println("Menu Not Found")
	}

}
