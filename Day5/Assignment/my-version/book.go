package main

import (
	"fmt"
)

type Book struct {
	Name   string
	Number int
	Author string
	Date   string
	Next   *Book
}

func (b *Book) Add(name string, number int, author, date string) {

	book_input := &Book{
		Name:   name,
		Number: number,
		Author: author,
		Date:   date,
	}
	if b.Next == nil {
		b = book_input
		return
	}
	b.Next = book_input
	return
}

func (b *Book) Show() {
	for b.Next != nil {
		fmt.Println(b)
		b = b.Next
	}
}
