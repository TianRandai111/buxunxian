package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	Id     int
	Name   string
	Number int
	Author string
	Time   string
	Exit   string
	Next   *Book
}

func Add_Book(b *Book) {
	fmt.Println("请输入：序号，书名，数量，作者，出版日期，如果输入结束请在出版日期后输入exit")
	var tail = b

	Pd := true
	for Pd {
		temp_book := bufio.NewReader(os.Stdin)
		temp_book_ele, err := temp_book.ReadString('\n')
		if err != nil {
			fmt.Println("赋值错误")
			Pd = false
		}
		temp_book_Split := strings.Split(temp_book_ele, ",")
		a, _ := strconv.Atoi(temp_book_Split[0])
		b, _ := strconv.Atoi(temp_book_Split[2])

		private_book := &Book{
			Id:     a,
			Name:   temp_book_Split[1],
			Number: b,
			Author: temp_book_Split[3],
			Time:   temp_book_Split[4],
			Exit:   temp_book_Split[5],
		}
		// private_book.Id, _ = strconv.Atoi(temp_book_Split[0])
		// private_book.Name = temp_book_Split[1]
		// private_book.Number, _ = strconv.Atoi(temp_book_Split[2])
		// private_book.Author = temp_book_Split[3]
		// private_book.Time = temp_book_Split[4]
		// private_book.Exit = temp_book_Split[5]
		//fmt.Printf("序号：%d\n书名：%s\n副本数：%d\n作者：%s\n日期：%s", private_book.Id, private_book.Name, private_book.Number, private_book.Author, private_book.Time)
		//fmt.Println(private_book)
		tail.Next = private_book
		tail = private_book
		if private_book.Exit != "exit" {
			continue
		} else if private_book.Exit == "exit" {
			Pd = false
		}
	}
	return
}

func Show_book(b *Book) {
	for b.Next != nil {
		fmt.Printf("序号：%d\n书名：%s\n副本数：%d\n作者：%s\n日期：%s\n", b.Id, b.Name, b.Number, b.Author, b.Time)
		fmt.Println("------------------------------")
		b = b.Next
	}
	return
}

func main() {
	var Book Book
	tmp_pd := true
	for tmp_pd {
		var v, i string
		fmt.Println("1.添加书籍信息\n2.查询书籍信息\n3.插入学生信息\n4查询已借走的书籍")
		fmt.Scanf("%s%s", &v, &i)
		switch v {
		case "1":
			Add_Book(&Book)
			fmt.Println("---", Book)
		case "2":
			Show_book(&Book)
		case "exit":
			tmp_pd = false
		}
	}
}
