package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	fmt.Println("this is book_input", book_input)
	fmt.Println("this is one b", b)
	if b.Next == nil {
		b = book_input
		fmt.Println("this is two b", b)
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

func main() {
	var book_info Book
	for 
	a := bufio.NewReader(os.Stdin)
	b, _ := a.ReadString('\n')
	c := strings.Split(b, " ")
	c1, _ := strconv.Atoi(c[1])
	if c[0] == "exit" {
		goto exit
	}
	book_info.Add(c[0], c1, c[2], c[3])
	book_info.Show()
exit:
}
