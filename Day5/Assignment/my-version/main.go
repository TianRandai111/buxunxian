package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Add_Book() {
	// judge := true
	fmt.Println("输入书籍的书名，数量，作者，出版日期")
	// for judge {
	B_Input := bufio.NewReader(os.Stdin)
	B_Receives, _ := B_Input.ReadString('\n')
	B_Split := strings.Split(B_Receives, " ")
	fmt.Println(B_Split)
	// if B_Split[0] == "exit" {
	// 	judge = false
	// }
	B_Split1, _ = strconv.Atoi(B_Split[1])
	fmt.Println(B_Split)
	Book_Info.Add(B_Split[0], B_Split1, B_Split[2], B_Split[3])

	// }
	return
}

var (
	B_Split1  int
	Book_Info Book
)

func main() {
	judge := true
	var choose string
	for judge {
		fmt.Println("1.输入书籍信息;\n2.查询书籍信息;\n3.输入学生信息;\n4.查询学生借书信息;\n5.查询那些书被借走;\n6.exit退出。")
		fmt.Scanf("%s\n", &choose)

		switch choose {
		case "1":
			Add_Book()
		case "2":
			Book_Info.Show()
		case "3":
			fmt.Println("3")
		case "4":
			fmt.Println("4")
		case "5":
			fmt.Println("5")
		case "6":
			judge = false
		default:
			fmt.Println("请输入正确的数字")
		}
	}
}
