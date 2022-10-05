package main

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

var books []string

func getBooks() {
	fmt.Println("===")
	fmt.Println(strings.ReplaceAll(strings.Join(books, "),("), ",", "\n"))
	fmt.Println("===")

}

func createBook(title string, price string, category string) {
	id := uuid.New()
	books = append(books, `ID: `+id.String()+`,Title: `+title+`,Price: `+price+`,Category: `+category)
	fmt.Println("Success create the book!")
}

func updateBook(id string, title string, price string, category string) {
	for i, book := range books {
		if strings.Contains(book, id) {
			books[i] = `ID: ` + id + `, Title: ` + title + `, Price: ` + price + `, Category: ` + category
		}
	}
	fmt.Println("Success update the book!")
}

func deleteBook(id string) {
	for i, book := range books {
		if strings.Contains(book, id) {
			books = append(books[:i], books[i+1:]...)
		}
	}
	fmt.Println("Success delete the book!")
}

func main() {
	fmt.Println("Welcome to Book App")
	// create menu using switch case
	for {
		fmt.Println("")
		fmt.Println("1. Get all books")
		fmt.Println("2. Create a book")
		fmt.Println("3. Update a book")
		fmt.Println("4. Delete a book")
		fmt.Println("5. Exit")
		var input int
		fmt.Print("Choose menu: ")
		fmt.Scan(&input)
		switch input {
		case 1:
			getBooks()
		case 2:
			var title, price, category string
			fmt.Print("Title: ")
			fmt.Scan(&title)
			fmt.Print("Price: ")
			fmt.Scan(&price)
			fmt.Print("Category: ")
			fmt.Scan(&category)
			createBook(title, price, category)
		case 3:
			var id, title, price, category string
			fmt.Print("ID: ")
			fmt.Scan(&id)
			fmt.Print("Title: ")
			fmt.Scan(&title)
			fmt.Print("Price: ")
			fmt.Scan(&price)
			fmt.Print("Category: ")
			fmt.Scan(&category)
			updateBook(id, title, price, category)
		case 4:
			var id string
			fmt.Print("ID: ")
			fmt.Scan(&id)
			deleteBook(id)
		case 5:
			return
		}
	}
}
