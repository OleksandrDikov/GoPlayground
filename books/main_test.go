package main

import (
	"fmt"
	"testing"
)

func TestBookToString_FormatsBookInfoAsString(t *testing.T) {
	input := Book{
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}

	want := fmt.Sprintf("%v by %v - %v copies", input.Title, input.Author, input.Copies)
	got := BookToString(input)

	if want != got {
		t.Fatal("String not the same!")
	}
}
