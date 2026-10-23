package books

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
}

var catalog = []Book{
	{
		Title:  "In the Company of Cheerful Ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
	},
	{
		Title:  "White Heat",
		Author: "Dominic Sandbrook",
		Copies: 2,
	},
}

func BookToString(book Book) string {
	return fmt.Sprintf("%v by %v (copies: %v)", book.Title, book.Author, book.Copies)
}

func GetAllBooks() []Book {
	return catalog
}
