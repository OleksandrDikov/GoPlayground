package books_test

import (
	"books"
	"fmt"
	"testing"
)

func TestBookToString_FormatsBookInfoAsString(t *testing.T) {
	input := books.Book{
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}

	want := fmt.Sprintf("%v by %v (copies: %v)", input.Title, input.Author, input.Copies)
	got := books.BookToString(input)

	if want != got {
		t.Fatalf("want %q, got %q", want, got)
	}
}
